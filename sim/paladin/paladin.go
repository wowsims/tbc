package paladin

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

type Paladin struct {
	core.Character

	Talents proto.PaladinTalents
	
	// maybe I should make this a pointer instead so I can nil check for no active seal
	currentSeal core.Aura

	sealOfBlood            core.SimpleCast
	sealOfCommand          core.SimpleCast
	sealOfTheCrusader      core.SimpleCast

	crusaderStrikeTemplate core.MeleeAbilityTemplate
	crusaderStrikeSpell    core.ActiveMeleeAbility

	judgementOfBloodTemplate core.SimpleSpellTemplate
	judgementOfBloodSpell    core.SimpleSpell

	judgementOfTheCrusaderTemplate core.SimpleSpellTemplate
	judgementOfTheCrusaderSpell    core.SimpleSpell
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
}

func (paladin *Paladin) Reset(sim *core.Simulation) {
	paladin.currentSeal.Expires = sim.CurrentTime
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

	return paladin
}

func init() {
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceBloodElf, Class: proto.Class_ClassPaladin}] = stats.Stats{
		stats.Strength:  123,
		stats.Agility:   79,
		stats.Stamina:   118,
		stats.Intellect: 87,
		stats.Spirit:    88,
		stats.Mana:      2953,
		stats.SpellCrit:   139.76,
		stats.AttackPower: 190,
		stats.MeleeCrit:   77.06,
	}
}
