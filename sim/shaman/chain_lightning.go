package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

const SpellIDCL6 int32 = 25442
var ChainLightningCooldownID = core.NewCooldownID()

func (shaman *Shaman) NewChainLightning(sim *core.Simulation, target *core.Target, isLightningOverload bool) core.DirectCastAction {
	cast := shaman.NewElectricCast(
		"Chain Lightning",
		core.ActionID{
			SpellID: SpellIDCL6,
			CooldownID: ChainLightningCooldownID,
		},
		760.0,
		time.Millisecond*2000,
		isLightningOverload)
	cast.Cooldown = time.Second*6

	hitInput := core.DirectCastDamageInput{
		Target: target,
		MinBaseDamage: 734,
		MaxBaseDamage: 838,
		SpellCoefficient: 0.651,
		DamageMultiplier: 1,
	}
	shaman.ApplyElectricSpellHitInputModifiers(&hitInput, isLightningOverload)

	numHits := core.MinInt32(3, sim.GetNumTargets())
	hitInputs := make([]core.DirectCastDamageInput, 0, numHits)
	hitInputs = append(hitInputs, hitInput)
	for i := int32(1); i < numHits; i++ {
		bounceHit := hitInputs[i - 1] // Makes a copy

		// Pick targets by index, incrementally.
		newTargetIndex := (bounceHit.Target.Index + 1) % sim.GetNumTargets()
		bounceHit.Target = sim.GetTarget(newTargetIndex)

		if shaman.HasAura(Tidefury2PcAuraID) {
			bounceHit.DamageMultiplier *= 0.83
		} else {
			bounceHit.DamageMultiplier *= 0.7
		}

		hitInputs = append(hitInputs, bounceHit)
	}

	return core.NewDirectCastAction(
		sim,
		cast,
		hitInputs,
		// OnCastComplete
		func(sim *core.Simulation, cast *core.Cast) {
			shaman.OnElectricSpellCastComplete(sim, cast, isLightningOverload)
		},
		// OnSpellHit
		func(sim *core.Simulation, cast *core.Cast, result *core.DirectCastDamageResult) {
			shaman.OnElectricSpellHit(sim, cast, result)

			if !isLightningOverload {
				lightningOverloadChance := float64(shaman.Talents.LightningOverload) * 0.04 / 3
				if sim.RandomFloat("CL Lightning Overload") < lightningOverloadChance {
					overloadAction := shaman.NewChainLightning(sim, hitInput.Target, true)
					overloadAction.Act(sim)
				}
			}
		},
		// OnSpellMiss
		func(sim *core.Simulation, cast *core.Cast) {
		})
}
