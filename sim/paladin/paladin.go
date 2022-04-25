package paladin

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

type Paladin struct {
	core.Character

	Talents proto.PaladinTalents

	CurrentSeal      *core.Aura
	CurrentJudgement *core.Aura

	Consecration           *core.Spell
	CrusaderStrike         *core.Spell
	Exorcism               *core.Spell
	JudgementOfBlood       *core.Spell
	JudgementOfTheCrusader *core.Spell
	JudgementOfWisdom      *core.Spell
	SealOfBlood            *core.Spell
	SealOfCommand          *core.Spell
	SealOfTheCrusader      *core.Spell
	SealOfWisdom           *core.Spell

	ConsecrationDot *core.Dot

	JudgementOfTheCrusaderAura *core.Aura
	JudgementOfWisdomAura      *core.Aura
	SealOfBloodAura            *core.Aura
	SealOfCommandAura          *core.Aura
	SealOfTheCrusaderAura      *core.Aura
	SealOfWisdomAura           *core.Aura
}

// Implemented by each Paladin spec.
type PaladinAgent interface {
	core.Agent

	// The Paladin controlled by this Agent.
	GetPaladin() *Paladin
}

func (paladin *Paladin) GetCharacter() *core.Character {
	return &paladin.Character
}

func (paladin *Paladin) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
}

func (paladin *Paladin) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {
}

func (paladin *Paladin) Init(sim *core.Simulation) {
	paladin.registerCrusaderStrikeSpell(sim)
	paladin.registerExorcismSpell(sim)
	paladin.registerJudgements(sim)
}

func (paladin *Paladin) Reset(sim *core.Simulation) {
	paladin.CurrentSeal = nil
	paladin.CurrentJudgement = nil
}

func (paladin *Paladin) OnAutoAttack(sim *core.Simulation, spell *core.Spell) {
}

// maybe need to add stat dependencies
func NewPaladin(character core.Character, talents proto.PaladinTalents) *Paladin {
	paladin := &Paladin{
		Character: character,
		Talents:   talents,
	}

	paladin.EnableManaBar()

	// Add paladin stat dependencies
	paladin.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Intellect,
		ModifiedStat: stats.SpellCrit,
		Modifier: func(intellect float64, spellCrit float64) float64 {
			return spellCrit + (intellect/80)*core.SpellCritRatingPerCritChance
		},
	})

	paladin.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Strength,
		ModifiedStat: stats.AttackPower,
		Modifier: func(strength float64, attackPower float64) float64 {
			return attackPower + strength*2
		},
	})

	paladin.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Agility,
		ModifiedStat: stats.MeleeCrit,
		Modifier: func(agility float64, meleeCrit float64) float64 {
			return meleeCrit + (agility/25)*core.MeleeCritRatingPerCritChance
		},
	})

	paladin.setupSealOfBlood()
	paladin.setupSealOfTheCrusader()
	paladin.setupSealOfWisdom()

	paladin.registerAvengingWrathCD()

	paladin.setupJudgementRefresh()

	return paladin
}

func init() {
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceBloodElf, Class: proto.Class_ClassPaladin}] = stats.Stats{
		stats.Health:      3197,
		stats.Stamina:     118,
		stats.Intellect:   87,
		stats.Mana:        2953,
		stats.Spirit:      88,
		stats.Strength:    123,
		stats.AttackPower: 190,
		stats.Agility:     79,
		stats.MeleeCrit:   14.35,
		stats.SpellCrit:   73.69,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDraenei, Class: proto.Class_ClassPaladin}] = stats.Stats{
		stats.Health:      3197,
		stats.Stamina:     119,
		stats.Intellect:   84,
		stats.Mana:        2953,
		stats.Spirit:      91,
		stats.Strength:    127,
		stats.AttackPower: 190,
		stats.Agility:     74,
		stats.MeleeCrit:   14.35,
		stats.SpellCrit:   73.69,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceHuman, Class: proto.Class_ClassPaladin}] = stats.Stats{
		stats.Health:      3197,
		stats.Stamina:     120,
		stats.Intellect:   83,
		stats.Mana:        2953,
		stats.Spirit:      97,
		stats.Strength:    126,
		stats.AttackPower: 190,
		stats.Agility:     77,
		stats.MeleeCrit:   14.35,
		stats.SpellCrit:   73.69,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDwarf, Class: proto.Class_ClassPaladin}] = stats.Stats{
		stats.Health:      3197,
		stats.Stamina:     123,
		stats.Intellect:   82,
		stats.Mana:        2953,
		stats.Spirit:      88,
		stats.Strength:    128,
		stats.AttackPower: 190,
		stats.Agility:     73,
		stats.MeleeCrit:   14.35,
		stats.SpellCrit:   73.69,
	}
}
