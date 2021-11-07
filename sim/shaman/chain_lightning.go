package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

const SpellIDCL6 int32 = 25442
var ChainLightningCooldownID = core.NewCooldownID()

type ChainLightning struct {
	ElectricSpell
}

func (cl ChainLightning) GetActionID() core.ActionID {
	return core.ActionID{
		SpellID: SpellIDCL6,
		CooldownID: ChainLightningCooldownID,
	}
}

func (cl ChainLightning) GetCooldown() time.Duration {
	return time.Second*6
}

func (cl ChainLightning) GetHitInputs(sim *core.Simulation, cast core.DirectCastAction) []core.DirectCastDamageInput{
	hitInput := core.DirectCastDamageInput{
		Target: cl.Target,
		MinBaseDamage: 734,
		MaxBaseDamage: 838,
		SpellCoefficient: 0.651,
		DamageMultiplier: 1,
	}

	cl.ApplyHitInputModifiers(&hitInput)

	numHits := core.MinInt32(3, sim.GetNumTargets())
	hitInputs := make([]core.DirectCastDamageInput, 0, numHits)
	hitInputs = append(hitInputs, hitInput)

	for i := int32(1); i < numHits; i++ {
		bounceHit := hitInputs[i - 1] // Makes a copy

		// Pick targets by index, incrementally.
		newTargetIndex := (bounceHit.Target.Index + 1) % sim.GetNumTargets()
		bounceHit.Target = sim.GetTarget(newTargetIndex)

		if cl.Shaman.HasAura(Tidefury2PcAuraID) {
			bounceHit.DamageMultiplier *= 0.83
		} else {
			bounceHit.DamageMultiplier *= 0.7
		}

		hitInputs = append(hitInputs, bounceHit)
	}

	return hitInputs
}


func (cl ChainLightning) OnSpellHit(sim *core.Simulation, cast core.DirectCastAction, result *core.DirectCastDamageResult) {
	cl.OnElectricSpellHit(sim, cast, result)

	if !cl.IsLightningOverload {
		lightningOverloadChance := float64(cl.Shaman.Talents.LightningOverload) * 0.04 / 3
		if sim.RandomFloat("LO") < lightningOverloadChance {
			overloadAction := NewChainLightning(sim, cl.Shaman, cl.Target, true)
			overloadAction.Act(sim)
		}
	}
}

func NewChainLightning(sim *core.Simulation, shaman *Shaman, target *core.Target, IsLightningOverload bool) core.DirectCastAction {
	return core.NewDirectCastAction(
		sim,
		ChainLightning{ElectricSpell{
			Shaman: shaman,
			Target: target,
			IsLightningOverload: IsLightningOverload,
			name: "Chain Lightning",
			baseManaCost: 760.0,
			baseCastTime: time.Millisecond*2000,
		}})
}
