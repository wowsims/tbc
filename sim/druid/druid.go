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

	// cached cast stuff
	starfireSpell         core.SingleTargetDirectDamageSpell
	starfire8CastTemplate core.SingleTargetDirectDamageSpellTemplate
	starfire6CastTemplate core.SingleTargetDirectDamageSpellTemplate

	MoonfireSpell        core.DamageOverTimeSpell
	moonfireCastTemplate core.DamageOverTimeSpellTemplate
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

func (druid *Druid) Reset(newsim *core.Simulation) {
}

func (druid *Druid) Act(sim *core.Simulation) time.Duration {
	return core.NeverExpires // makes the bot wait forever and do nothing.
}

func (druid *Druid) Init(sim *core.Simulation) {
	druid.starfire8CastTemplate = druid.newStarfireTemplate(sim, 8)
	druid.starfire6CastTemplate = druid.newStarfireTemplate(sim, 6)
	druid.moonfireCastTemplate = druid.newMoonfireTemplate(sim)
}

func NewDruid(char core.Character, selfBuffs SelfBuffs, talents proto.DruidTalents) Druid {
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
