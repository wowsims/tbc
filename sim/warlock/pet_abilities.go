package warlock

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

type PetAbilityType byte

const (
	Unknown PetAbilityType = iota
	Cleave
	Intercept
	LashOfPain
	Firebolt
)

type PetAbility struct {
	Type     PetAbilityType
	ActionID core.ActionID

	// Mana cost
	Cost float64
	CD   core.Cooldown
	Cast func(target *core.Target)
}

// Returns whether the ability was successfully cast.
func (ability *PetAbility) TryCast(sim *core.Simulation, target *core.Target, wp *WarlockPet) bool {
	if wp.CurrentMana() < ability.Cost {
		return false
	}
	if ability.CD.Duration != 0 && !ability.CD.IsReady(sim) {
		return false
	}

	wp.SpendMana(sim, ability.Cost, ability.ActionID)
	ability.Cast(target)
	return true
}

func (wp *WarlockPet) NewPetAbility(sim *core.Simulation, abilityType PetAbilityType, isPrimary bool) PetAbility {
	switch abilityType {
	case Cleave:
		return wp.newCleave(sim)
	case Intercept:
		return wp.newIntercept(sim)
	case LashOfPain:
		return wp.newLashOfPain(sim)
	case Firebolt:
		return wp.newFirebolt(sim)
	case Unknown:
		return PetAbility{}
	default:
		panic("Invalid pet ability type")
	}
}

func (wp *WarlockPet) newIntercept(sim *core.Simulation) PetAbility {
	return PetAbility{}
}
func (wp *WarlockPet) newFirebolt(sim *core.Simulation) PetAbility {
	actionID := core.ActionID{SpellID: 27267}
	pa := PetAbility{
		ActionID: actionID,
		Cost:     190,
	}

	spell := wp.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolFire,
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond * 2000,
			},
			IgnoreHaste: true,
		},
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask:         core.ProcMaskSpellDamage,
			DamageMultiplier: 1.0,
			ThreatMultiplier: 1,
			BaseDamage:       core.BaseDamageConfigMagic(112, 127, 0.571),
			OutcomeApplier:   core.OutcomeFuncMagicHitAndCrit(2),
		}),
	})

	pa.Cast = func(target *core.Target) {
		spell.Cast(sim, target)
	}
	return pa
}
func (wp *WarlockPet) newCleave(sim *core.Simulation) PetAbility {
	actionID := core.ActionID{SpellID: 30223}
	cd := core.Cooldown{
		Timer:    wp.NewTimer(),
		Duration: time.Second * 6,
	}

	pa := PetAbility{
		ActionID: actionID,
		Cost:     295, // 10% of base
		CD:       cd,
	}

	spell := wp.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasMeleeMetrics,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
			CD:          cd,
		},
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask:         core.ProcMaskMeleeMHSpecial,
			DamageMultiplier: 1.0,
			ThreatMultiplier: 1,
			BaseDamage:       core.BaseDamageConfigMeleeWeapon(core.MainHand, false, 78, 1.0, true),
			OutcomeApplier:   core.OutcomeFuncMeleeSpecialHitAndCrit(2),
		}),
	})

	pa.Cast = func(target *core.Target) {
		spell.Cast(sim, target)
	}
	return pa
}

func (wp *WarlockPet) newLashOfPain(sim *core.Simulation) PetAbility {
	actionID := core.ActionID{SpellID: 27274}
	cd := core.Cooldown{
		Timer:    wp.NewTimer(),
		Duration: time.Second * 12,
	}

	pa := PetAbility{
		ActionID: actionID,
		Cost:     190,
		CD:       cd,
	}

	spell := wp.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolShadow,
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
			CD:          cd,
		},
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask:         core.ProcMaskSpellDamage,
			DamageMultiplier: 1.0,
			ThreatMultiplier: 1,
			BaseDamage:       core.BaseDamageConfigMagic(123, 123, 0.429),
			OutcomeApplier:   core.OutcomeFuncMagicHitAndCrit(2),
		}),
	})

	pa.Cast = func(target *core.Target) {
		spell.Cast(sim, target)
	}
	return pa
}
