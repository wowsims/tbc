package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

type PetAbilityType int

const (
	Unknown PetAbilityType = iota
	Bite
	Claw
	FireBreath
	Gore
	LightningBreath
	PoisonSpit
	Screech
)

type PetAbility struct {
	Type PetAbilityType

	ActionID core.ActionID

	// Focus cost
	Cost float64

	// 0 if no cooldown
	Cooldown time.Duration

	CooldownID core.CooldownID

	Cast func(target *core.Target)
}

// Returns whether the ability was successfully cast.
func (ability *PetAbility) TryCast(sim *core.Simulation, target *core.Target, hp *HunterPet) bool {
	if hp.currentFocus < ability.Cost {
		return false
	}
	if ability.Cooldown != 0 && hp.IsOnCD(ability.CooldownID, sim.CurrentTime) {
		return false
	}

	hp.SpendFocus(sim, ability.Cost, ability.ActionID)
	ability.Cast(target)
	return true
}

func (hp *HunterPet) NewPetAbility(sim *core.Simulation, abilityType PetAbilityType, isPrimary bool) PetAbility {
	switch abilityType {
	case Bite:
		return hp.newBite(sim, isPrimary)
	case Claw:
		return hp.newClaw(sim, isPrimary)
	case FireBreath:
		return PetAbility{}
	case Gore:
		return hp.newGore(sim, isPrimary)
	case LightningBreath:
		return hp.newLightningBreath(sim, isPrimary)
	case PoisonSpit:
		return PetAbility{}
	case Screech:
		return hp.newScreech(sim, isPrimary)
	case Unknown:
		return PetAbility{}
	default:
		panic("Invalid pet ability type")
	}
	return PetAbility{}
}

var PetPrimaryCooldownID = core.NewCooldownID()
var PetSecondaryCooldownID = core.NewCooldownID()

func (hp *HunterPet) newBite(sim *core.Simulation, isPrimary bool) PetAbility {
	actionID := core.ActionID{SpellID: 27050}
	cooldown := time.Second * 10
	if isPrimary {
		actionID.CooldownID = PetPrimaryCooldownID
	} else {
		actionID.CooldownID = PetSecondaryCooldownID
	}

	pa := PetAbility{
		ActionID:   actionID,
		Type:       Bite,
		Cost:       35,
		Cooldown:   cooldown,
		CooldownID: actionID.CooldownID,
	}

	spell := hp.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasMeleeMetrics,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
			Cooldown:    cooldown,
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask:         core.ProcMaskMeleeMHSpecial,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BaseDamage:       core.BaseDamageConfigRoll(108, 132),
			OutcomeApplier:   core.OutcomeFuncMeleeSpecialHitAndCrit(2),
		}),
	})

	pa.Cast = func(target *core.Target) {
		spell.Cast(sim, target)
	}
	return pa
}

func (hp *HunterPet) newClaw(sim *core.Simulation, isPrimary bool) PetAbility {
	actionID := core.ActionID{SpellID: 27049}
	pa := PetAbility{
		ActionID: actionID,
		Type:     Claw,
		Cost:     25,
	}

	spell := hp.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasMeleeMetrics,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask:         core.ProcMaskMeleeMHSpecial,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BaseDamage:       core.BaseDamageConfigRoll(54, 76),
			OutcomeApplier:   core.OutcomeFuncMeleeSpecialHitAndCrit(2),
		}),
	})

	pa.Cast = func(target *core.Target) {
		spell.Cast(sim, target)
	}
	return pa
}

func (hp *HunterPet) newGore(sim *core.Simulation, isPrimary bool) PetAbility {
	actionID := core.ActionID{SpellID: 35298}

	pa := PetAbility{
		ActionID: actionID,
		Type:     Gore,
		Cost:     25,
	}

	spell := hp.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasMeleeMetrics,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask:         core.ProcMaskMeleeMHSpecial,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BaseDamage: core.WrapBaseDamageConfig(core.BaseDamageConfigRoll(37, 61), func(oldCalculator core.BaseDamageCalculator) core.BaseDamageCalculator {
				return func(sim *core.Simulation, spellEffect *core.SpellEffect, spell *core.Spell) float64 {
					damage := oldCalculator(sim, spellEffect, spell)
					if sim.RandomFloat("Gore") < 0.5 {
						damage *= 2
					}
					return damage
				}
			}),
			OutcomeApplier: core.OutcomeFuncMeleeSpecialHitAndCrit(2),
		}),
	})

	pa.Cast = func(target *core.Target) {
		spell.Cast(sim, target)
	}
	return pa
}

func (hp *HunterPet) newLightningBreath(sim *core.Simulation, isPrimary bool) PetAbility {
	actionID := core.ActionID{SpellID: 25011}

	pa := PetAbility{
		ActionID: actionID,
		Type:     LightningBreath,
		Cost:     50,
	}

	spell := hp.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolNature,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BaseDamage:       core.BaseDamageConfigMagic(80, 93, 0.05),
			OutcomeApplier:   core.OutcomeFuncMagicHitAndCrit(1.5),
		}),
	})

	pa.Cast = func(target *core.Target) {
		spell.Cast(sim, target)
	}
	return pa
}

func (hp *HunterPet) newScreech(sim *core.Simulation, isPrimary bool) PetAbility {
	actionID := core.ActionID{SpellID: 27051}

	pa := PetAbility{
		ActionID: actionID,
		Type:     Screech,
		Cost:     20,
	}

	spell := hp.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasMeleeMetrics,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask:         core.ProcMaskMeleeMHSpecial,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BaseDamage:       core.BaseDamageConfigRoll(33, 61),
			OutcomeApplier:   core.OutcomeFuncMeleeSpecialHitAndCrit(2),
		}),
	})

	pa.Cast = func(target *core.Target) {
		spell.Cast(sim, target)
	}
	return pa
}
