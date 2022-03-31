package priest

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

type Priest struct {
	core.Character
	SelfBuffs
	Talents proto.PriestTalents

	SurgeOfLight bool

	Latency float64

	// cached cast stuff
	// TODO: aoe multi-target situations will need multiple spells ticking for each target.
	DevouringPlague *core.SimpleSpellTemplate
	HolyFire        *core.SimpleSpellTemplate
	MindBlast       *core.SimpleSpellTemplate
	MindFlay        []*core.SimpleSpellTemplate
	ShadowWordDeath *core.SimpleSpellTemplate
	ShadowWordPain  *core.SimpleSpellTemplate
	Shadowfiend     *core.SimpleSpellTemplate
	Smite           *core.SimpleSpellTemplate
	Starshards      *core.SimpleSpellTemplate
	VampiricTouch   *core.SimpleSpellTemplate
	VampiricTouch2  *core.SimpleSpellTemplate

	CurVTSpell  *core.SimpleSpellTemplate
	NextVTSpell *core.SimpleSpellTemplate
}

type SelfBuffs struct {
	UseShadowfiend bool

	PowerInfusionTarget proto.RaidTarget
}

func (priest *Priest) GetCharacter() *core.Character {
	return &priest.Character
}

func (priest *Priest) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
	if priest.Talents.DivineSpirit {
		ds := proto.TristateEffect_TristateEffectRegular
		if priest.Talents.ImprovedDivineSpirit == 2 {
			// TODO: handle a larger variety of IDS values.
			ds = proto.TristateEffect_TristateEffectImproved
		}
		raidBuffs.DivineSpirit = core.MaxTristate(raidBuffs.DivineSpirit, ds)
	}
}

func (priest *Priest) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {
}

func (priest *Priest) Init(sim *core.Simulation) {
	priest.registerDevouringPlagueSpell(sim)
	priest.registerHolyFireSpell(sim)
	priest.registerMindBlastSpell(sim)
	priest.registerShadowWordDeathSpell(sim)
	priest.registerShadowWordPainSpell(sim)
	priest.registerShadowfiendSpell(sim)
	priest.registerSmiteSpell(sim)
	priest.registerStarshardsSpell(sim)
	priest.registerVampiricTouchSpell(sim)

	priest.MindFlay = []*core.SimpleSpellTemplate{
		nil, // So we can use # of ticks as the index
		priest.newMindFlaySpell(sim, 1),
		priest.newMindFlaySpell(sim, 2),
		priest.newMindFlaySpell(sim, 3),
	}
}

func (priest *Priest) Reset(newsim *core.Simulation) {
	priest.CurVTSpell = priest.VampiricTouch
	priest.NextVTSpell = priest.VampiricTouch
}

func New(char core.Character, selfBuffs SelfBuffs, talents proto.PriestTalents) *Priest {
	priest := &Priest{
		Character: char,
		SelfBuffs: selfBuffs,
		Talents:   talents,
	}
	priest.EnableManaBar()

	priest.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Intellect,
		ModifiedStat: stats.SpellCrit,
		Modifier: func(intellect float64, spellCrit float64) float64 {
			return spellCrit + (intellect/80)*core.SpellCritRatingPerCritChance
		},
	})

	priest.registerShadowfiendCD()

	priest.registerPowerInfusionCD()

	return priest
}

func init() {
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceHuman, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Health:    3211,
		stats.Strength:  39,
		stats.Agility:   45,
		stats.Stamina:   58,
		stats.Intellect: 145,
		stats.Spirit:    166,
		stats.Mana:      2620,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 1.24,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDwarf, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Health:    3211,
		stats.Strength:  41,
		stats.Agility:   41,
		stats.Stamina:   61,
		stats.Intellect: 144,
		stats.Spirit:    150,
		stats.Mana:      2620,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 1.24,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceNightElf, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Health:    3211,
		stats.Strength:  36,
		stats.Agility:   50,
		stats.Stamina:   57,
		stats.Intellect: 145,
		stats.Spirit:    151,
		stats.Mana:      2620,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 1.24,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDraenei, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Health:    3211,
		stats.Strength:  40,
		stats.Agility:   42,
		stats.Stamina:   57,
		stats.Intellect: 146,
		stats.Spirit:    153,
		stats.Mana:      2620,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 1.24,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceUndead, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Health:    3211,
		stats.Strength:  38,
		stats.Agility:   43,
		stats.Stamina:   59,
		stats.Intellect: 143,
		stats.Spirit:    156,
		stats.Mana:      2620,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 1.24,
	}
	trollStats := stats.Stats{
		stats.Health:    3211,
		stats.Strength:  40,
		stats.Agility:   47,
		stats.Stamina:   59,
		stats.Intellect: 141,
		stats.Spirit:    152,
		stats.Mana:      2620,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 1.24,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll10, Class: proto.Class_ClassPriest}] = trollStats
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll30, Class: proto.Class_ClassPriest}] = trollStats
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceBloodElf, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Health:    3211,
		stats.Strength:  36,
		stats.Agility:   47,
		stats.Stamina:   57,
		stats.Intellect: 149,
		stats.Spirit:    150,
		stats.Mana:      2620,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 1.24,
	}
}

// Agent is a generic way to access underlying priest on any of the agents.
type Agent interface {
	GetPriest() *Priest
}
