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

	NaturesGrace bool // when true next spellcast is 0.5s faster
	RebirthUsed  bool

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

	FaerieFireSpell        core.SimpleSpell
	faerieFireCastTemplate core.SimpleSpellTemplate

	HurricaneSpell        core.SimpleSpell
	hurricaneCastTemplate core.SimpleSpellTemplate
}

type SelfBuffs struct {
	Omen bool

	InnervateTarget proto.RaidTarget
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
	druid.faerieFireCastTemplate = druid.newFaerieFireTemplate(sim)
	druid.hurricaneCastTemplate = druid.newHurricaneTemplate(sim)
}

func (druid *Druid) Reset(sim *core.Simulation) {
	druid.RebirthUsed = false
}

func (druid *Druid) Act(sim *core.Simulation) time.Duration {
	return core.NeverExpires // does nothing
}

func New(char core.Character, selfBuffs SelfBuffs, talents proto.DruidTalents) *Druid {
	druid := &Druid{
		Character:   char,
		SelfBuffs:   selfBuffs,
		Talents:     talents,
		RebirthUsed: false,
	}
	druid.EnableManaBar()

	druid.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Intellect,
		ModifiedStat: stats.SpellCrit,
		Modifier: func(intellect float64, spellCrit float64) float64 {
			return spellCrit + (intellect/79.4)*core.SpellCritRatingPerCritChance
		},
	})

	druid.registerInnervateCD()

	return druid
}

func init() {
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTauren, Class: proto.Class_ClassDruid}] = stats.Stats{
		stats.Health:    3434,
		stats.Strength:  81,
		stats.Agility:   65,
		stats.Stamina:   85,
		stats.Intellect: 115,
		stats.Spirit:    135,
		stats.Mana:      2370,
		stats.SpellCrit: 40.66, // 3.29% chance to crit shown on naked character screen
		// 4498 health shown on naked character (would include tauren bonus)
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceNightElf, Class: proto.Class_ClassDruid}] = stats.Stats{
		stats.Health:    3434,
		stats.Strength:  73,
		stats.Agility:   75,
		stats.Stamina:   82,
		stats.Intellect: 120,
		stats.Spirit:    133,
		stats.Mana:      2370,
		stats.SpellCrit: 40.60, // 3.35% chance to crit shown on naked character screen
		// 4254 health shown on naked character
	}
}

// Agent is a generic way to access underlying druid on any of the agents (for example balance druid.)
type Agent interface {
	GetDruid() *Druid
}
