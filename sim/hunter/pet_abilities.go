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
	cooldown := time.Second * 10
	ama := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            core.ActionID{SpellID: 27050},
				Character:           &hp.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				Cooldown:            cooldown,
				GCD:                 core.GCDDefault,
				IgnoreHaste:         true,
				CritMultiplier:      2,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ProcMask:               core.ProcMaskMeleeMHSpecial,
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
			},
			DirectInput: core.DirectDamageInput{
				MinBaseDamage: 108,
				MaxBaseDamage: 132,
			},
		},
	}

	pa := PetAbility{
		Type:     Bite,
		Cost:     35,
		Cooldown: cooldown,
	}

	if isPrimary {
		ama.ActionID.CooldownID = PetPrimaryCooldownID
		pa.CooldownID = PetPrimaryCooldownID
	} else {
		ama.ActionID.CooldownID = PetSecondaryCooldownID
		pa.CooldownID = PetSecondaryCooldownID
	}
	pa.ActionID = ama.ActionID

	template := core.NewSimpleSpellTemplate(ama)
	cast := core.SimpleSpell{}

	pa.Cast = func(target *core.Target) {
		template.Apply(&cast)

		// Set dynamic fields, i.e. the stuff we couldn't precompute.
		cast.Effect.Target = target

		cast.Cast(sim)
	}
	return pa
}

func (hp *HunterPet) newClaw(sim *core.Simulation, isPrimary bool) PetAbility {
	ama := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            core.ActionID{SpellID: 27049},
				Character:           &hp.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				GCD:                 core.GCDDefault,
				IgnoreHaste:         true,
				CritMultiplier:      2,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ProcMask:               core.ProcMaskMeleeMHSpecial,
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
			},
			DirectInput: core.DirectDamageInput{
				MinBaseDamage: 54,
				MaxBaseDamage: 76,
			},
		},
	}

	pa := PetAbility{
		Type: Claw,
		Cost: 25,
	}
	pa.ActionID = ama.ActionID

	template := core.NewSimpleSpellTemplate(ama)
	cast := core.SimpleSpell{}

	pa.Cast = func(target *core.Target) {
		template.Apply(&cast)

		// Set dynamic fields, i.e. the stuff we couldn't precompute.
		cast.Effect.Target = target

		cast.Cast(sim)
	}
	return pa
}

func (hp *HunterPet) newGore(sim *core.Simulation, isPrimary bool) PetAbility {
	ama := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            core.ActionID{SpellID: 35298},
				Character:           &hp.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				GCD:                 core.GCDDefault,
				IgnoreHaste:         true,
				CritMultiplier:      2,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ProcMask:               core.ProcMaskMeleeMHSpecial,
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
			},
			DirectInput: core.DirectDamageInput{
				MinBaseDamage: 37,
				MaxBaseDamage: 61,
			},
		},
	}

	pa := PetAbility{
		Type: Gore,
		Cost: 25,
	}
	pa.ActionID = ama.ActionID

	template := core.NewSimpleSpellTemplate(ama)
	cast := core.SimpleSpell{}

	pa.Cast = func(target *core.Target) {
		template.Apply(&cast)

		// Set dynamic fields, i.e. the stuff we couldn't precompute.
		cast.Effect.Target = target
		if sim.RandomFloat("Gore") < 0.5 {
			cast.Effect.DamageMultiplier *= 2
		}

		cast.Cast(sim)
	}
	return pa
}

func (hp *HunterPet) newLightningBreath(sim *core.Simulation, isPrimary bool) PetAbility {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            core.ActionID{SpellID: 25011},
				Character:           &hp.Character,
				CritRollCategory:    core.CritRollCategoryMagical,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				SpellSchool:         core.SpellSchoolNature,
				GCD:                 core.GCDDefault,
				CritMultiplier:      1.5,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
			},
			DirectInput: core.DirectDamageInput{
				MinBaseDamage:    80,
				MaxBaseDamage:    93,
				SpellCoefficient: 0.05,
			},
		},
	}

	pa := PetAbility{
		Type: LightningBreath,
		Cost: 50,
	}
	pa.ActionID = spell.ActionID

	template := core.NewSimpleSpellTemplate(spell)
	cast := core.SimpleSpell{}

	pa.Cast = func(target *core.Target) {
		template.Apply(&cast)

		// Set dynamic fields, i.e. the stuff we couldn't precompute.
		cast.Effect.Target = target

		cast.Init(sim)
		cast.Cast(sim)
	}
	return pa
}

func (hp *HunterPet) newScreech(sim *core.Simulation, isPrimary bool) PetAbility {
	ama := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            core.ActionID{SpellID: 27051},
				Character:           &hp.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				GCD:                 core.GCDDefault,
				CritMultiplier:      2,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ProcMask:               core.ProcMaskMeleeMHSpecial,
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
			},
			DirectInput: core.DirectDamageInput{
				MinBaseDamage: 33,
				MaxBaseDamage: 61,
			},
		},
	}

	pa := PetAbility{
		Type: Screech,
		Cost: 20,
	}
	pa.ActionID = ama.ActionID

	template := core.NewSimpleSpellTemplate(ama)
	cast := core.SimpleSpell{}

	pa.Cast = func(target *core.Target) {
		template.Apply(&cast)

		// Set dynamic fields, i.e. the stuff we couldn't precompute.
		cast.Effect.Target = target

		cast.Cast(sim)
	}
	return pa
}
