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

	naturesGrace bool // when true next spellcast is 0.5s faster

	// cached cast stuff
	starfireSpell         core.SingleTargetDirectDamageSpell
	starfire8CastTemplate core.SingleTargetDirectDamageSpellTemplate
	starfire6CastTemplate core.SingleTargetDirectDamageSpellTemplate

	MoonfireSpell        core.DamageOverTimeSpell
	moonfireCastTemplate core.DamageOverTimeSpellTemplate

	wrathSpell        core.SingleTargetDirectDamageSpell
	wrathCastTemplate core.SingleTargetDirectDamageSpellTemplate
}

type SelfBuffs struct {
	Omen      bool
	Innervate bool
}

func (druid *Druid) GetCharacter() *core.Character {
	return &druid.Character
}

func (druid *Druid) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
	raidBuffs.GiftOfTheWild = proto.TristateEffect_TristateEffectRegular
	if druid.Talents.ImprovedMarkOfTheWild > 0 { // ya ya whatever
		raidBuffs.GiftOfTheWild = proto.TristateEffect_TristateEffectImproved
	}
}

const ravenGoddessItemID = 32387

func (druid *Druid) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {
	if druid.Talents.MoonkinForm { // assume if you have moonkin talent you are using it.
		partyBuffs.MoonkinAura = proto.TristateEffect_TristateEffectRegular
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
}

func (druid *Druid) Reset(newsim *core.Simulation) {
	druid.moonfireCastTemplate.Apply(&druid.MoonfireSpell)
	druid.Character.Reset(newsim)
}

func (druid *Druid) Advance(sim *core.Simulation, elapsedTime time.Duration) {
	druid.Character.RegenManaMP5Only(sim, elapsedTime)
	druid.Character.Advance(sim, elapsedTime)
}

func (druid *Druid) Act(sim *core.Simulation) time.Duration {
	return core.NeverExpires // makes the bot wait forever and do nothing.
}

func (druid *Druid) applyOnHitTalents(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
	if druid.Talents.NaturesGrace && spellEffect.Crit {
		druid.naturesGrace = true
	}

	if druid.Talents.WrathOfCenarius > 0 {
		if spellCast.ActionID.SpellID == SpellIDSF8 || spellCast.ActionID.SpellID == SpellIDSF6 {
			spellEffect.BonusSpellPower += (druid.GetStat(stats.SpellPower) + druid.GetStat(stats.ArcaneSpellPower)) * 0.04 * float64(druid.Talents.WrathOfCenarius)
		}

		if spellCast.ActionID.SpellID == SpellIDWrath {
			spellEffect.BonusSpellPower += (druid.GetStat(stats.SpellPower) + druid.GetStat(stats.NatureSpellPower)) * 0.02 * float64(druid.Talents.WrathOfCenarius)
		}
	}
}

func (druid *Druid) applyNaturesGrace(spellCast *core.SpellCast) {
	if druid.naturesGrace {
		spellCast.CastTime -= time.Millisecond * 500
		// This applies on cast complete, removing the effect.
		//  if it crits, during 'onspellhit' then it will be reapplied (see func above)
		spellCast.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
			druid.naturesGrace = false
		}
	}
}

func NewDruid(char core.Character, selfBuffs SelfBuffs, talents proto.DruidTalents) Druid {

	if talents.LunarGuidance > 0 {
		bonus := 0.083 * float64(talents.LunarGuidance)
		char.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.SpellPower,
			Modifier: func(intellect float64, spellPower float64) float64 {
				return spellPower + intellect*bonus
			},
		})
	}

	if talents.Dreamstate > 0 {
		bonus := 0.0333 * float64(talents.Dreamstate)
		char.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.MP5,
			Modifier: func(intellect float64, mp5 float64) float64 {
				return mp5 + intellect*bonus
			},
		})
	}

	if talents.Intensity > 0 {
		char.PsuedoStats.SpiritRegenRateCasting = float64(talents.Intensity) * 0.1
	}

	return Druid{
		Character: char,
		SelfBuffs: selfBuffs,
		Talents:   talents,
	}
}

var FaerieFireAuraID = core.NewAuraID()

func init() {
	// TODO: get the actual real base stats here.
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTauren, Class: proto.Class_ClassDruid}] = stats.Stats{
		stats.Strength:  103,
		stats.Agility:   61,
		stats.Stamina:   113,
		stats.Intellect: 109,
		stats.Spirit:    122,
		stats.Mana:      2678,
		stats.SpellCrit: 47.89,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceNightElf, Class: proto.Class_ClassDruid}] = stats.Stats{
		stats.Intellect: 104,
		stats.Mana:      2678,
		stats.Spirit:    135,
		stats.SpellCrit: 47.89,
	}
}
