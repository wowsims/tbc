package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

type Priest struct {
	core.Character
	SelfBuffs
	Talents proto.PriestTalents

	// cached cast stuff
	// TODO: aoe multi-target situations will need multiple spells ticking for each target.
	MindFlaySpell        core.SimpleSpell
	mindflayCastTemplate core.SimpleSpellTemplate

	mindblastSpell        core.SimpleSpell
	mindblastCastTemplate core.SimpleSpellTemplate

	swdSpell        core.SimpleSpell
	swdCastTemplate core.SimpleSpellTemplate

	SWPSpell        core.SimpleSpell
	swpCastTemplate core.SimpleSpellTemplate

	VTSpell        core.SimpleSpell
	vtCastTemplate core.SimpleSpellTemplate

	ShadowfiendSpell    core.SimpleSpell
	shadowfiendTemplate core.SimpleSpellTemplate

	DevouringPlagueSpell    core.SimpleSpell
	devouringPlagueTemplate core.SimpleSpellTemplate

	// TODO: starshards
}

type SelfBuffs struct {
}

func (priest *Priest) GetCharacter() *core.Character {
	return &priest.Character
}

func (priest *Priest) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
	raidBuffs.DivineSpirit = proto.TristateEffect_TristateEffectRegular
}

func (priest *Priest) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {

}

func (priest *Priest) Init(sim *core.Simulation) {
	priest.mindflayCastTemplate = priest.newMindflayTemplate(sim)
	priest.mindblastCastTemplate = priest.newMindBlastTemplate(sim)
	priest.swpCastTemplate = priest.newSWPTemplate(sim)
	priest.vtCastTemplate = priest.newVTTemplate(sim)
	priest.swdCastTemplate = priest.newSWDTemplate(sim)
	priest.shadowfiendTemplate = priest.newShadowfiendTemplate(sim)
	priest.devouringPlagueTemplate = priest.newDevouringPlagueTemplate(sim)
}

func (priest *Priest) Reset(newsim *core.Simulation) {
	// Cleanup and pending dots and casts
	priest.mindblastSpell = core.SimpleSpell{}
	priest.swdSpell = core.SimpleSpell{}
	priest.MindFlaySpell = core.SimpleSpell{}
	priest.SWPSpell = core.SimpleSpell{}
	priest.VTSpell = core.SimpleSpell{}
	priest.ShadowfiendSpell = core.SimpleSpell{}
	priest.DevouringPlagueSpell = core.SimpleSpell{}

	priest.Character.Reset(newsim)
}

func (priest *Priest) Advance(sim *core.Simulation, elapsedTime time.Duration) {
	// druid should never be outside the 5s window, use combat regen.
	priest.Character.RegenManaCasting(sim, elapsedTime)
	priest.Character.Advance(sim, elapsedTime)
}

func (priest *Priest) Act(sim *core.Simulation) time.Duration {
	return core.NeverExpires // does nothing
}

func (priest *Priest) applyTalentsToShadowSpell(cast *core.Cast, effect *core.SpellHitEffect) {
	if cast.ActionID.SpellID == SpellIDMF || cast.ActionID.SpellID == SpellIDMB {
		cast.ManaCost -= cast.BaseManaCost * float64(priest.Talents.FocusedMind) * 0.05
	}
	if cast.SpellSchool == stats.ShadowSpellPower {
		effect.DamageMultiplier *= 1 + float64(priest.Talents.Darkness)*0.02

		if priest.Talents.Shadowform {
			effect.DamageMultiplier *= 1.15
		}

		// shadow focus gives 2% hit per level
		effect.BonusSpellHitRating += float64(priest.Talents.ShadowFocus) * 2 * core.SpellHitRatingPerHitChance
	}
}

func NewPriest(char core.Character, selfBuffs SelfBuffs, talents proto.PriestTalents) Priest {

	// char.AddStat(stats.SpellHit, float64(talents.BalanceOfPower)*2*core.SpellHitRatingPerHitChance)

	char.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Intellect,
		ModifiedStat: stats.SpellCrit,
		Modifier: func(intellect float64, spellCrit float64) float64 {
			return spellCrit + (intellect/79.4)*core.SpellCritRatingPerCritChance
		},
	})

	if talents.Meditation > 0 {
		char.PseudoStats.SpiritRegenRateCasting = float64(talents.Meditation) * 0.1
	}

	// if talents.LunarGuidance > 0 {
	// 	bonus := (0.25 / 3) * float64(talents.LunarGuidance)
	// 	char.AddStatDependency(stats.StatDependency{
	// 		SourceStat:   stats.Intellect,
	// 		ModifiedStat: stats.SpellPower,
	// 		Modifier: func(intellect float64, spellPower float64) float64 {
	// 			return spellPower + intellect*bonus
	// 		},
	// 	})
	// }

	priest := Priest{
		Character: char,
		SelfBuffs: selfBuffs,
		Talents:   talents,
	}

	return priest
}

// newVTOnTick is the OnDamage function for all priest DoTs to apply VT
func newVTOnTick(party *core.Party) core.OnDamageTick {
	return func(sim *core.Simulation, damage float64) {
		s := stats.Stats{stats.Mana: damage * 0.05}
		if sim.Log != nil {
			sim.Log("VT Regenerated %0f mana.\n", s[stats.Mana])
		}
		party.AddStats(s)
	}
}

var InnerFocusAuraID = core.NewAuraID()
var InnerFocusCooldownID = core.NewCooldownID()

func ApplyInnerFocus(sim *core.Simulation, priest *Priest) bool {
	sim.MetricsAggregator.AddInstantCast(&priest.Character, core.ActionID{SpellID: 14751})

	priest.Character.AddAura(sim, core.Aura{
		ID:      InnerFocusAuraID,
		Name:    "Inner Focus",
		Expires: core.NeverExpires,
		OnCast: func(sim *core.Simulation, cast *core.Cast) {
			cast.ManaCost = 0
		},
		OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
			spellEffect.BonusSpellCritRating += 25 * core.SpellCritRatingPerCritChance
		},
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
			// Remove the buff and put skill on CD
			priest.SetCD(InnerFocusCooldownID, sim.CurrentTime+time.Minute*3)
			priest.RemoveAura(sim, InnerFocusAuraID)
		},
	})
	return true
}

// TODO: Get Priest base stats
func init() {
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceHuman, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  81,
		stats.Agility:   65,
		stats.Stamina:   85,
		stats.Intellect: 115,
		stats.Spirit:    135,
		stats.Mana:      2090,  // 3815 mana shown on naked character
		stats.SpellCrit: 40.66, // 3.29% chance to crit shown on naked character screen
		// 4498 health shown on naked character (would include tauren bonus)
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDwarf, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  81,
		stats.Agility:   65,
		stats.Stamina:   85,
		stats.Intellect: 115,
		stats.Spirit:    135,
		stats.Mana:      2090,  // 3815 mana shown on naked character
		stats.SpellCrit: 40.66, // 3.29% chance to crit shown on naked character screen
		// 4498 health shown on naked character (would include tauren bonus)
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceNightElf, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  81,
		stats.Agility:   65,
		stats.Stamina:   85,
		stats.Intellect: 115,
		stats.Spirit:    135,
		stats.Mana:      2090,  // 3815 mana shown on naked character
		stats.SpellCrit: 40.66, // 3.29% chance to crit shown on naked character screen
		// 4498 health shown on naked character (would include tauren bonus)
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDraenei, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  81,
		stats.Agility:   65,
		stats.Stamina:   85,
		stats.Intellect: 115,
		stats.Spirit:    135,
		stats.Mana:      2090,  // 3815 mana shown on naked character
		stats.SpellCrit: 40.66, // 3.29% chance to crit shown on naked character screen
		// 4498 health shown on naked character (would include tauren bonus)
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceUndead, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  81,
		stats.Agility:   65,
		stats.Stamina:   85,
		stats.Intellect: 115,
		stats.Spirit:    135,
		stats.Mana:      2090,  // 3815 mana shown on naked character
		stats.SpellCrit: 40.66, // 3.29% chance to crit shown on naked character screen
		// 4498 health shown on naked character (would include tauren bonus)
	}
	troll := stats.Stats{
		stats.Strength:  81,
		stats.Agility:   65,
		stats.Stamina:   85,
		stats.Intellect: 115,
		stats.Spirit:    135,
		stats.Mana:      2090,  // 3815 mana shown on naked character
		stats.SpellCrit: 40.66, // 3.29% chance to crit shown on naked character screen
		// 4498 health shown on naked character (would include tauren bonus)
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll10, Class: proto.Class_ClassPriest}] = troll
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll30, Class: proto.Class_ClassPriest}] = troll
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceBloodElf, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  81,
		stats.Agility:   65,
		stats.Stamina:   85,
		stats.Intellect: 115,
		stats.Spirit:    135,
		stats.Mana:      2090,  // 3815 mana shown on naked character
		stats.SpellCrit: 40.66, // 3.29% chance to crit shown on naked character screen
		// 4498 health shown on naked character (would include tauren bonus)
	}
}

// Agent is a generic way to access underlying priest on any of the agents
type Agent interface {
	GetPriest() *Priest
}
