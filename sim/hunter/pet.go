package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

type HunterPet struct {
	core.Pet

	focusBar

	// Combines a few static effects.
	damageMultiplier float64
}

func (hunter *Hunter) NewHunterPet() *HunterPet {
	if hunter.Options.PetType == proto.Hunter_Options_PetNone {
		return nil
	}
	petConfig := PetConfigs[hunter.Options.PetType]

	hp := &HunterPet{
		Pet: core.NewPet(
			petConfig.Name,
			&hunter.Character,
			hunterPetBaseStats,
			hunterPetStatInheritance,
			true,
		),
		damageMultiplier: petConfig.DamageMultiplier,
	}

	// Happiness
	hp.damageMultiplier *= 1.25

	// Cobra reflexes
	hp.PseudoStats.MeleeSpeedMultiplier *= 1.3
	hp.damageMultiplier *= 0.85

	hp.EnableFocusBar(1.0, func(sim *core.Simulation) {
		if !hp.IsOnCD(core.GCDCooldownID, sim.CurrentTime) {
			hp.OnGCDReady(sim)
		}
	})

	hp.EnableAutoAttacks(hp, core.AutoAttackOptions{
		MainHand: core.Weapon{
			BaseDamageMin: 42,
			BaseDamageMax: 68,
			SwingSpeed:    2,
			SwingDuration: time.Second * 2,
		},
		AutoSwingMelee: true,
	})

	hunter.AddPet(hp)

	return hp
}

func (hp *HunterPet) GetPet() *core.Pet {
	return &hp.Pet
}

func (hp *HunterPet) Init(sim *core.Simulation) {
}

func (hp *HunterPet) Reset(newsim *core.Simulation) {
}

func (hp *HunterPet) OnGCDReady(sim *core.Simulation) {
}

// These numbers are just rough guesses based on looking at some logs.
var hunterPetBaseStats = stats.Stats{
	stats.Agility:   127,
	stats.Strength:  162,
	stats.MeleeCrit: 1.1515 * core.MeleeCritRatingPerCritChance,
}

var hunterPetStatInheritance = func(ownerStats stats.Stats) stats.Stats {
	return stats.Stats{
		stats.Stamina:     ownerStats[stats.Stamina] * 0.3,
		stats.Armor:       ownerStats[stats.Armor] * 0.35,
		stats.AttackPower: ownerStats[stats.RangedAttackPower] * 0.22,
		stats.SpellPower:  ownerStats[stats.RangedAttackPower] * 0.125,
	}
}

type PetConfig struct {
	Name string

	DamageMultiplier float64
}

var PetConfigs = map[proto.Hunter_Options_PetType]PetConfig{
	proto.Hunter_Options_Bat: PetConfig{
		Name:             "Bat",
		DamageMultiplier: 1.07,
	},
	proto.Hunter_Options_Cat: PetConfig{
		Name:             "Cat",
		DamageMultiplier: 1.1,
	},
	proto.Hunter_Options_Owl: PetConfig{
		Name:             "Owl",
		DamageMultiplier: 1.07,
	},
	proto.Hunter_Options_Raptor: PetConfig{
		Name:             "Raptor",
		DamageMultiplier: 1.1,
	},
	proto.Hunter_Options_Ravager: PetConfig{
		Name:             "Ravager",
		DamageMultiplier: 1.1,
	},
	proto.Hunter_Options_WindSerpent: PetConfig{
		Name:             "Wind Serpent",
		DamageMultiplier: 1.07,
	},
}
