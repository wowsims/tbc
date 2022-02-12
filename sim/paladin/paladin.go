package paladin

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

type Paladin struct {
	core.Character

	Talents proto.PaladinTalents

	currentSeal core.Aura

	consecrationTemplate   core.SimpleSpellTemplate
	ConsecrationSpell      core.SimpleSpell
	crusaderStrikeTemplate core.MeleeAbilityTemplate
	crusaderStrikeSpell    core.ActiveMeleeAbility
	sealOfBlood            core.SimpleCast
	sealOfCommand          core.SimpleCast
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
	paladin.consecrationTemplate = paladin.newConsecrationTemplate(sim)
}

func (paladin *Paladin) Reset(sim *core.Simulation) {
}

// maybe need to add stat dependencies
func NewPaladin(character core.Character, talents proto.PaladinTalents) *Paladin {
	paladin := &Paladin{
		Character: character,
		Talents:   talents,
	}

	paladin.EnableManaBar()

	paladin.setupSealOfBlood()
	paladin.setupSealOfCommand()

	return paladin
}

func init() {
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceBloodElf, Class: proto.Class_ClassPaladin}] = stats.Stats{
		stats.Strength:  123,
		stats.Agility:   79,
		stats.Stamina:   118,
		stats.Intellect: 87,
		stats.Spirit:    88,
		stats.Mana:      3978, // pretty sure I need to subtract mana from the int stat

		stats.AttackPower: 120,
	}
}
