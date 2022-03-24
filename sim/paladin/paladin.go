package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

type Paladin struct {
	core.Character

	Talents proto.PaladinTalents

	currentSealID           core.AuraID
	currentSealExpires      time.Duration
	currentJudgementID      core.AuraID
	currentJudgementExpires time.Duration

	sealOfBlood       core.SimpleCast
	sealOfCommand     core.SimpleCast
	sealOfTheCrusader core.SimpleCast
	sealOfWisdom      core.SimpleCast

	SealOfTheCrusaderAura core.Aura
	SealOfWisdomAura      core.Aura
	SealOfCommandAura     core.Aura

	consecrationTemplate core.SimpleSpellTemplate
	ConsecrationSpell    core.SimpleSpell

	exorcismTemplate core.SimpleSpellTemplate
	exorcismSpell    core.SimpleSpell

	crusaderStrikeTemplate core.SimpleSpellTemplate
	crusaderStrikeSpell    core.SimpleSpell

	judgementOfBloodTemplate core.SimpleSpellTemplate
	judgementOfBloodSpell    core.SimpleSpell

	judgementOfTheCrusaderTemplate core.SimpleSpellTemplate
	judgementOfTheCrusaderSpell    core.SimpleSpell

	judgementOfWisdomTemplate core.SimpleSpellTemplate
	judgementOfWisdomSpell    core.SimpleSpell
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
	paladin.crusaderStrikeTemplate = paladin.newCrusaderStrikeTemplate(sim)
	paladin.judgementOfBloodTemplate = paladin.newJudgementOfBloodTemplate(sim)
	paladin.judgementOfTheCrusaderTemplate = paladin.newJudgementOfTheCrusaderTemplate(sim)
	paladin.judgementOfWisdomTemplate = paladin.newJudgementOfWisdomTemplate(sim)
	paladin.consecrationTemplate = paladin.newConsecrationTemplate(sim)
	paladin.exorcismTemplate = paladin.newExorcismTemplate(sim)
}

func (paladin *Paladin) Reset(sim *core.Simulation) {
	paladin.currentSealID = 0
	paladin.currentSealExpires = 0
	paladin.currentJudgementID = 0
	paladin.currentJudgementExpires = 0
}

func (paladin *Paladin) OnAutoAttack(sim *core.Simulation, ability *core.SimpleSpell) {
	if paladin.currentJudgementID == 0 || paladin.currentJudgementExpires >= sim.CurrentTime {
		return
	}
	paladin.currentJudgementExpires = sim.CurrentTime + JudgementDuration
	ability.Effect.Target.RefreshAura(sim, paladin.currentJudgementID)
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
	paladin.setupSealOfCommand()
	paladin.setupSealOfTheCrusader()
	paladin.setupSealOfWisdom()

	paladin.applyTalents()

	paladin.registerAvengingWrathCD()

	return paladin
}

func init() {
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceBloodElf, Class: proto.Class_ClassPaladin}] = stats.Stats{
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
