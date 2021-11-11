package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

const SpellIDCL6 int32 = 25442
var ChainLightningCooldownID = core.NewCooldownID()

// Returns a cast object for Chain Lightning with as many fields precomputed as possible.
func (shaman *Shaman) newChainLightningTemplate(sim *core.Simulation, isLightningOverload bool) core.DirectCastAction {
	spellTemplate := shaman.newElectricSpellTemplate(
		"Chain Lightning",
		core.ActionID{
			SpellID: SpellIDCL6,
			CooldownID: ChainLightningCooldownID,
		},
		760.0,
		time.Millisecond*2000,
		isLightningOverload)
	spellTemplate.Cast.Cooldown = time.Second*6

	hitInput := core.DirectCastDamageInput{
		MinBaseDamage: 734,
		MaxBaseDamage: 838,
		SpellCoefficient: 0.651,
		DamageMultiplier: 1,
	}
	shaman.applyElectricSpellHitInputModifiers(&hitInput, isLightningOverload)

	numHits := core.MinInt32(3, sim.GetNumTargets())
	hitInputs := make([]core.DirectCastDamageInput, 0, numHits)
	hitInputs = append(hitInputs, hitInput)
	for i := int32(1); i < numHits; i++ {
		bounceHit := hitInputs[i - 1] // Makes a copy

		if shaman.HasAura(Tidefury2PcAuraID) {
			bounceHit.DamageMultiplier *= 0.83
		} else {
			bounceHit.DamageMultiplier *= 0.7
		}

		hitInputs = append(hitInputs, bounceHit)
	}
	spellTemplate.HitInputs = hitInputs

	spellTemplate.HitResults = make([]core.DirectCastDamageResult, numHits)

	if !isLightningOverload && shaman.Talents.LightningOverload > 0 {
		lightningOverloadChance := float64(shaman.Talents.LightningOverload) * 0.04 / 3
		spellTemplate.OnSpellHit = func(sim *core.Simulation, cast *core.Cast, result *core.DirectCastDamageResult) {
			if shaman.Talents.ElementalFocus && result.Crit {
				shaman.ElementalFocusStacks = 2
			}

			if sim.RandomFloat("CL Lightning Overload") < lightningOverloadChance {
				overloadAction := shaman.NewChainLightning(sim, result.Target, true)
				overloadAction.Act(sim)
			}
		}
	} else {
		spellTemplate.OnSpellHit = func(sim *core.Simulation, cast *core.Cast, result *core.DirectCastDamageResult) {
			if shaman.Talents.ElementalFocus && result.Crit {
				shaman.ElementalFocusStacks = 2
			}
		}
	}

	return spellTemplate
}

func (shaman *Shaman) NewChainLightning(sim *core.Simulation, target *core.Target, isLightningOverload bool) *core.DirectCastAction {
	var spell *core.DirectCastAction

	// Initialize cast from precomputed template.
	if isLightningOverload {
		spell = &shaman.electricSpellLO
		*spell = shaman.chainLightningLOTemplate
		spell.HitInputs = shaman.clHitInputs
		copy(spell.HitInputs, shaman.chainLightningLOTemplate.HitInputs)
	} else {
		spell = &shaman.electricSpell
		*spell = shaman.chainLightningTemplate
		spell.HitInputs = shaman.clHitInputs
		copy(spell.HitInputs, shaman.chainLightningTemplate.HitInputs)
	}

	// Set dynamic fields, i.e. the stuff we couldn't precompute.

	// Set the targets.
	spell.HitInputs[0].Target = target
	curTargetIndex := target.Index
	numHits := core.MinInt32(3, sim.GetNumTargets())
	for i := int32(1); i < numHits; i++ {
		// Pick targets by index, incrementally.
		newTargetIndex := (curTargetIndex + 1) % sim.GetNumTargets()
		spell.HitInputs[i].Target = sim.GetTarget(newTargetIndex)
		curTargetIndex = newTargetIndex
	}

	shaman.applyElectricSpellInitModifiers(spell)

	spell.Init(sim)

	return spell
}
