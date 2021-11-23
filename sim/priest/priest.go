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
}

func (priest *Priest) Advance(sim *core.Simulation, elapsedTime time.Duration) {
	// druid should never be outside the 5s window, use combat regen.
	priest.Character.RegenManaCasting(sim, elapsedTime)
}

func (priest *Priest) Act(sim *core.Simulation) time.Duration {
	return core.NeverExpires // does nothing
}

func (priest *Priest) applyTalentsToShadowSpell(cast *core.Cast, effect *core.SpellHitEffect) {
	if cast.ActionID.SpellID == SpellIDSWD || cast.ActionID.SpellID == SpellIDMB {
		effect.BonusSpellCritRating += float64(priest.Talents.ShadowPower) * 3 * core.SpellCritRatingPerCritChance
	}

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

var InnerFocusAuraID = core.NewAuraID()
var InnerFocusCooldownID = core.NewCooldownID()

func ApplyInnerFocus(sim *core.Simulation, priest *Priest) bool {
	priest.Metrics.AddInstantCast(core.ActionID{SpellID: 14751})
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

func init() {

	// TODO: str/agi/stm are just the base priest stats, not modified for each race yet. Not sure it matters...

	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceHuman, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  146,
		stats.Agility:   184,
		stats.Stamina:   154,
		stats.Intellect: 180,
		stats.Spirit:    135,
		stats.Mana:      2090,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 1.24,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDwarf, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  146,
		stats.Agility:   184,
		stats.Stamina:   154,
		stats.Intellect: 179,
		stats.Spirit:    134,
		stats.Mana:      2090,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 1.24,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceNightElf, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  146,
		stats.Agility:   184,
		stats.Stamina:   154,
		stats.Intellect: 180,
		stats.Spirit:    135,
		stats.Mana:      2090,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 1.24,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDraenei, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  146,
		stats.Agility:   184,
		stats.Stamina:   154,
		stats.Intellect: 180,
		stats.Spirit:    137,
		stats.Mana:      2090,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 1.24,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceUndead, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  146,
		stats.Agility:   184,
		stats.Stamina:   154,
		stats.Intellect: 178,
		stats.Spirit:    135,
		stats.Mana:      2090,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 1.24,
	}
	troll := stats.Stats{
		stats.Strength:  146,
		stats.Agility:   184,
		stats.Stamina:   154,
		stats.Intellect: 176,
		stats.Spirit:    136,
		stats.Mana:      2090,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 1.24,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll10, Class: proto.Class_ClassPriest}] = troll
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll30, Class: proto.Class_ClassPriest}] = troll
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceBloodElf, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  146,
		stats.Agility:   184,
		stats.Stamina:   154,
		stats.Intellect: 183,
		stats.Spirit:    138,
		stats.Mana:      2090,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 1.24,
	}
}

// Agent is a generic way to access underlying priest on any of the agents
type Agent interface {
	GetPriest() *Priest
}

// class Priest(talents: Map<String, Talent>, spec: Spec) : Class(talents, spec) {
//     override val baseStats: Stats = Stats(
//         agility = 184,
//         intellect = 180,
//         strength = 146,
//         stamina = 154,
//         spirit = 135
//     )
// class Dwarf : Race() {
//     override var baseStats: Stats = Stats(
//         strength = 5,
//         agility = -4,
//         stamina = 1,
//         intellect = -1,
//         spirit = -1
//     )
// class Draenei : Race() {
//     override var baseStats: Stats = Stats(
//         strength = 1,
//         agility = -3,
//         spirit = 2
//     )
// class NightElf : Race() {
//     override var baseStats: Stats = Stats(
//         strength = -4,
//         agility = 4,
//         stamina = 0,
//         intellect = 0,
//         spirit = 0
//     )
// class Troll : Race() {
//     override var baseStats: Stats = Stats(
//         strength = 1,
//         agility = 2,
//         intellect = -4,
//         spirit = 1
//     )
// class Undead : Race() {
//     override var baseStats: Stats = Stats(
//         strength = -1,
//         agility = -2,
//         stamina = 0,
//         intellect = -2,
//         spirit = 5
//     )
// class BloodElf : Race() {
//     override var baseStats: Stats = Stats(
//         strength = -3,
//         agility = 2,
//         stamina = 0,
//         intellect = 3,
//         spirit = -2
//     )
// // https://worldofwarcraft.fandom.com/et/wiki/Spell_critical_strike
//     override val baseSpellCritChance: Double = 1.24
