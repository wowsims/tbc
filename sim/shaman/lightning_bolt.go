package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
)

const SpellIDLB12 int32 = 25449

func (shaman *Shaman) NewLightningBolt(sim *core.Simulation, target *core.Target, isLightningOverload bool) *core.DirectCastAction {
	baseManaCost := 300.0
	if shaman.Equip[items.ItemSlotRanged].ID == TotemOfThePulsingEarth {
		baseManaCost -= 27.0
	}

	spell := shaman.NewElectricSpell(
		"Lightning Bolt",
		core.ActionID{
			SpellID: SpellIDLB12,
		},
		baseManaCost,
		time.Millisecond*2500,
		isLightningOverload)

	hitInput := core.DirectCastDamageInput{
		Target: target,
		MinBaseDamage: 571,
		MaxBaseDamage: 652,
		SpellCoefficient: 0.794,
		DamageMultiplier: 1,
	}
	shaman.ApplyElectricSpellHitInputModifiers(&hitInput, isLightningOverload)

	spell.HitInputs = []core.DirectCastDamageInput{hitInput}

	spell.OnSpellHit = func(sim *core.Simulation, cast *core.Cast, result *core.DirectCastDamageResult) {
		shaman.OnElectricSpellHit(sim, cast, result)

		if !isLightningOverload {
			lightningOverloadChance := float64(shaman.Talents.LightningOverload) * 0.04
			if sim.RandomFloat("LB Lightning Overload") < lightningOverloadChance {
				overloadAction := shaman.NewLightningBolt(sim, target, true)
				overloadAction.Act(sim)
			}
		}
	}

	spell.Init(sim)
	return spell
}
