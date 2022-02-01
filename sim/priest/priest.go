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

	VTSpell        *core.SimpleSpell
	VTSpellCasting *core.SimpleSpell
	vtCastTemplate core.SimpleSpellTemplate

	ShadowfiendSpell    core.SimpleSpell
	shadowfiendTemplate core.SimpleSpellTemplate

	DevouringPlagueSpell    core.SimpleSpell
	devouringPlagueTemplate core.SimpleSpellTemplate

	StarshardsSpell    core.SimpleSpell
	starshardsTemplate core.SimpleSpellTemplate
}

type SelfBuffs struct {
	UseShadowfiend bool
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
	priest.mindflayCastTemplate = priest.newMindflayTemplate(sim)
	priest.mindblastCastTemplate = priest.newMindBlastTemplate(sim)
	priest.swpCastTemplate = priest.newShadowWordPainTemplate(sim)
	priest.vtCastTemplate = priest.newVampiricTouchTemplate(sim)
	priest.swdCastTemplate = priest.newShadowWordDeathTemplate(sim)
	priest.shadowfiendTemplate = priest.newShadowfiendTemplate(sim)
	priest.devouringPlagueTemplate = priest.newDevouringPlagueTemplate(sim)
	priest.starshardsTemplate = priest.newStarshardsTemplate(sim)
}

func (priest *Priest) Reset(newsim *core.Simulation) {
	// These spells still need special cleanup because they're wierd.
	priest.VTSpell = &core.SimpleSpell{}
	priest.VTSpellCasting = &core.SimpleSpell{}
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
	priest.applyTalents()

	return priest
}

func init() {
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceHuman, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  39,
		stats.Agility:   45,
		stats.Stamina:   58,
		stats.Intellect: 145,
		stats.Spirit:    166,
		stats.Mana:      2620,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 1.24,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDwarf, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  41,
		stats.Agility:   41,
		stats.Stamina:   61,
		stats.Intellect: 144,
		stats.Spirit:    150,
		stats.Mana:      2620,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 1.24,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceNightElf, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  36,
		stats.Agility:   50,
		stats.Stamina:   57,
		stats.Intellect: 145,
		stats.Spirit:    151,
		stats.Mana:      2620,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 1.24,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDraenei, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  40,
		stats.Agility:   42,
		stats.Stamina:   57,
		stats.Intellect: 146,
		stats.Spirit:    153,
		stats.Mana:      2620,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 1.24,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceUndead, Class: proto.Class_ClassPriest}] = stats.Stats{
		stats.Strength:  38,
		stats.Agility:   43,
		stats.Stamina:   59,
		stats.Intellect: 143,
		stats.Spirit:    156,
		stats.Mana:      2620,
		stats.SpellCrit: core.SpellCritRatingPerCritChance * 1.24,
	}
	trollStats := stats.Stats{
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
