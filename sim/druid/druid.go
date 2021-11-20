package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

type Druid struct {
	core.Character
	SelfBuffs
	Talents proto.DruidTalents

	innervateCD  time.Duration
	NaturesGrace bool // when true next spellcast is 0.5s faster

	// cached cast stuff
	starfireSpell         core.SimpleSpell
	starfire8CastTemplate core.SimpleSpellTemplate
	starfire6CastTemplate core.SimpleSpellTemplate

	MoonfireSpell        core.SimpleSpell
	moonfireCastTemplate core.SimpleSpellTemplate

	wrathSpell        core.SimpleSpell
	wrathCastTemplate core.SimpleSpellTemplate

	InsectSwarmSpell        core.SimpleSpell
	insectSwarmCastTemplate core.SimpleSpellTemplate

	malorne4p bool // cached since we need to check on every innervate
}

type SelfBuffs struct {
	Omen      bool
	Innervate bool
}

func (druid *Druid) GetCharacter() *core.Character {
	return &druid.Character
}

func (druid *Druid) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
	raidBuffs.GiftOfTheWild = core.MaxTristate(raidBuffs.GiftOfTheWild, proto.TristateEffect_TristateEffectRegular)
	if druid.Talents.ImprovedMarkOfTheWild == 5 { // probably could work on actually calculating the fraction effect later if we care.
		raidBuffs.GiftOfTheWild = proto.TristateEffect_TristateEffectImproved
	}
}

const ravenGoddessItemID = 32387

func (druid *Druid) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {
	if druid.Talents.MoonkinForm { // assume if you have moonkin talent you are using it.
		partyBuffs.MoonkinAura = core.MaxTristate(partyBuffs.MoonkinAura, proto.TristateEffect_TristateEffectRegular)
		for _, e := range druid.Equip {
			if e.ID == ravenGoddessItemID {
				partyBuffs.MoonkinAura = proto.TristateEffect_TristateEffectImproved
				break
			}
		}
	}
}

func (druid *Druid) Init(sim *core.Simulation) {
	druid.starfire8CastTemplate = druid.newStarfireTemplate(sim, 8)
	druid.starfire6CastTemplate = druid.newStarfireTemplate(sim, 6)
	druid.moonfireCastTemplate = druid.newMoonfireTemplate(sim)
	druid.wrathCastTemplate = druid.newWrathTemplate(sim)
	druid.insectSwarmCastTemplate = druid.newInsectSwarmTemplate(sim)
}

func (druid *Druid) Reset(newsim *core.Simulation) {
	// Cleanup and pending dots and casts
	druid.MoonfireSpell = core.SimpleSpell{}
	druid.InsectSwarmSpell = core.SimpleSpell{}
	druid.starfireSpell = core.SimpleSpell{}
	druid.wrathSpell = core.SimpleSpell{}

	druid.Character.Reset(newsim)
}

func (druid *Druid) Advance(sim *core.Simulation, elapsedTime time.Duration) {
	// druid should never be outside the 5s window, use combat regen.
	druid.Character.RegenManaCasting(sim, elapsedTime)
	druid.Character.Advance(sim, elapsedTime)
}

func (druid *Druid) Act(sim *core.Simulation) time.Duration {
	return core.NeverExpires // does nothing
}

func (druid *Druid) applyOnHitTalents(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
	if druid.Talents.NaturesGrace && spellEffect.Crit {
		druid.NaturesGrace = true
	}
}

func (druid *Druid) applyNaturesGrace(spellCast *core.SpellCast) {
	if druid.NaturesGrace {
		spellCast.CastTime -= time.Millisecond * 500
		// This applies on cast complete, removing the effect.
		//  if it crits, during 'onspellhit' then it will be reapplied (see func above)
		spellCast.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
			druid.NaturesGrace = false
		}
	}
}

func NewDruid(char core.Character, selfBuffs SelfBuffs, talents proto.DruidTalents) Druid {

	char.AddStat(stats.SpellHit, float64(talents.BalanceOfPower)*2*core.SpellHitRatingPerHitChance)

	char.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Intellect,
		ModifiedStat: stats.SpellCrit,
		Modifier: func(intellect float64, spellCrit float64) float64 {
			return spellCrit + (intellect/79.4)*core.SpellCritRatingPerCritChance
		},
	})

	if talents.LunarGuidance > 0 {
		bonus := (0.25 / 3) * float64(talents.LunarGuidance)
		char.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.SpellPower,
			Modifier: func(intellect float64, spellPower float64) float64 {
				return spellPower + intellect*bonus
			},
		})
	}

	if talents.Dreamstate > 0 {
		bonus := (0.1 / 3) * float64(talents.Dreamstate)
		char.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.MP5,
			Modifier: func(intellect float64, mp5 float64) float64 {
				return mp5 + intellect*bonus
			},
		})
	}

	if talents.Intensity > 0 {
		char.PseudoStats.SpiritRegenRateCasting = float64(talents.Intensity) * 0.1
	}

	return Druid{
		Character: char,
		SelfBuffs: selfBuffs,
		Talents:   talents,
		malorne4p: ItemSetMalorne.CharacterHasSetBonus(&char, 4),
	}
}

func init() {
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTauren, Class: proto.Class_ClassDruid}] = stats.Stats{
		stats.Strength:  81,
		stats.Agility:   65,
		stats.Stamina:   85,
		stats.Intellect: 115,
		stats.Spirit:    135,
		stats.Mana:      2090,  // 3815 mana shown on naked character
		stats.SpellCrit: 40.66, // 3.29% chance to crit shown on naked character screen
		// 4498 health shown on naked character (would include tauren bonus)
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceNightElf, Class: proto.Class_ClassDruid}] = stats.Stats{
		stats.Strength:  73,
		stats.Agility:   75,
		stats.Stamina:   82,
		stats.Intellect: 120,
		stats.Spirit:    133,
		stats.Mana:      2090,  // 3890 mana shown on naked character
		stats.SpellCrit: 40.60, // 3.35% chance to crit shown on naked character screen
		// 4254 health shown on naked character
	}
}

// Agent is a generic way to access underlying druid on any of the agents (for example balance druid.)
type Agent interface {
	GetDruid() *Druid
}
