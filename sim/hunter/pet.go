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

	config PetConfig

	hunterOwner *Hunter

	// Combines a few static effects.
	damageMultiplier float64

	killCommandTemplate core.MeleeAbilityTemplate
	killCommand         core.ActiveMeleeAbility

	primaryAbility   PetAbility
	secondaryAbility PetAbility
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
		config:           petConfig,
		hunterOwner:      hunter,
		damageMultiplier: petConfig.DamageMultiplier,
	}

	// Happiness
	hp.damageMultiplier *= 1.25

	hp.EnableFocusBar(1.0+0.5*float64(hunter.Talents.BestialDiscipline), func(sim *core.Simulation) {
		if !hp.IsOnCD(core.GCDCooldownID, sim.CurrentTime) {
			hp.OnGCDReady(sim)
		}
	})

	casterPenalty := 0.0
	//if petConfig.IsCaster {
	//	casterPenalty = 2
	//}

	hp.EnableAutoAttacks(hp, core.AutoAttackOptions{
		MainHand: core.Weapon{
			BaseDamageMin:  42 - casterPenalty,
			BaseDamageMax:  68 - casterPenalty,
			SwingSpeed:     2,
			SwingDuration:  time.Second * 2,
			CritMultiplier: 2,
		},
		AutoSwingMelee: true,
	})

	// Cobra reflexes
	hp.PseudoStats.MeleeSpeedMultiplier *= 1.3
	hp.AutoAttacks.ActiveMeleeAbility.Effect.DamageMultiplier *= 0.85

	hp.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Strength,
		ModifiedStat: stats.AttackPower,
		Modifier: func(strength float64, attackPower float64) float64 {
			return attackPower + strength*2
		},
	})
	hp.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Agility,
		ModifiedStat: stats.MeleeCrit,
		Modifier: func(agility float64, meleeCrit float64) float64 {
			return meleeCrit + (agility/33)*core.MeleeCritRatingPerCritChance
		},
	})

	core.ApplyPetConsumeEffects(&hp.Character, hunter.Consumes)
	hp.applyPetEffects()

	hunter.AddPet(hp)

	return hp
}

func (hp *HunterPet) GetPet() *core.Pet {
	return &hp.Pet
}

func (hp *HunterPet) Init(sim *core.Simulation) {
	hp.killCommandTemplate = hp.newKillCommandTemplate(sim)
	hp.primaryAbility = hp.NewPetAbility(sim, hp.config.PrimaryAbility, true)
	hp.secondaryAbility = hp.NewPetAbility(sim, hp.config.SecondaryAbility, false)
}

func (hp *HunterPet) Reset(sim *core.Simulation) {
	hp.focusBar.reset(sim)
	if sim.Log != nil {
		hp.Log(sim, "Pet stats: %s", hp.GetStats())
		inheritedStats := hunterPetStatInheritance(hp.hunterOwner.GetStats())
		hp.Log(sim, "Inherited stats: %s", inheritedStats)
	}
}

func (hp *HunterPet) OnGCDReady(sim *core.Simulation) {
	target := sim.GetPrimaryTarget()
	if !hp.primaryAbility.TryCast(sim, target, hp) {
		if hp.secondaryAbility.Type != Unknown {
			hp.secondaryAbility.TryCast(sim, target, hp)
		}
	}
}

// These numbers are just rough guesses based on looking at some logs.
var hunterPetBaseStats = stats.Stats{
	stats.Agility:     127,
	stats.Strength:    162,
	stats.AttackPower: -20, // Apparently pets and warriors have a AP penalty.
	stats.MeleeCrit:   1.1515 * core.MeleeCritRatingPerCritChance,
}

var hunterPetStatInheritance = func(ownerStats stats.Stats) stats.Stats {
	return stats.Stats{
		stats.Stamina:     ownerStats[stats.Stamina] * 0.3,
		stats.Armor:       ownerStats[stats.Armor] * 0.35,
		stats.AttackPower: ownerStats[stats.RangedAttackPower] * 0.22,
		stats.SpellPower:  ownerStats[stats.RangedAttackPower] * 0.128,
	}
}

var PetEffectsAuraID = core.NewAuraID()

func (hp *HunterPet) applyPetEffects() {
	hp.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: PetEffectsAuraID,
			OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				hitEffect.DamageMultiplier *= hp.damageMultiplier
			},
			OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				spellEffect.DamageMultiplier *= hp.damageMultiplier
			},
			OnBeforePeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage *float64) {
				*tickDamage *= hp.damageMultiplier
			},
		}
	})
}

type PetConfig struct {
	Name string

	DamageMultiplier float64

	PrimaryAbility   PetAbilityType
	SecondaryAbility PetAbilityType

	IsCaster bool // Caster pets have a melee penalty
}

var PetConfigs = map[proto.Hunter_Options_PetType]PetConfig{
	proto.Hunter_Options_Bat: PetConfig{
		Name:             "Bat",
		DamageMultiplier: 1.07,
		PrimaryAbility:   Bite,
		SecondaryAbility: Screech,
	},
	proto.Hunter_Options_Cat: PetConfig{
		Name:             "Cat",
		DamageMultiplier: 1.1,
		PrimaryAbility:   Bite,
		SecondaryAbility: Claw,
	},
	proto.Hunter_Options_Owl: PetConfig{
		Name:             "Owl",
		DamageMultiplier: 1.07,
		PrimaryAbility:   Claw,
		SecondaryAbility: Screech,
	},
	proto.Hunter_Options_Raptor: PetConfig{
		Name:             "Raptor",
		DamageMultiplier: 1.1,
		PrimaryAbility:   Bite,
		SecondaryAbility: Claw,
	},
	proto.Hunter_Options_Ravager: PetConfig{
		Name:             "Ravager",
		DamageMultiplier: 1.1,
		PrimaryAbility:   Bite,
		SecondaryAbility: Gore,
	},
	proto.Hunter_Options_WindSerpent: PetConfig{
		Name:             "Wind Serpent",
		DamageMultiplier: 1.07,
		PrimaryAbility:   Bite,
		SecondaryAbility: LightningBreath,
		IsCaster:         true,
	},
}
