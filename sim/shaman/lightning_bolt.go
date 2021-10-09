package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
)

const SpellIDLB12 int32 = 25449

type LightningBolt struct {
	ElectricSpell
}

func (lb LightningBolt) GetActionID() core.ActionID {
	return core.ActionID{
		SpellID: SpellIDLB12,
	}
}

func (lb LightningBolt) GetCooldown() time.Duration {
	return 0
}

func (lb LightningBolt) GetHitInputs(sim *core.Simulation, cast *core.DirectCastAction) []core.DirectCastDamageInput{
	hitInput := core.DirectCastDamageInput{
		MinBaseDamage: 571,
		MaxBaseDamage: 652,
		SpellCoefficient: 0.794,
		DamageMultiplier: 1,
	}

	lb.ApplyHitInputModifiers(&hitInput)

	return []core.DirectCastDamageInput{hitInput}
}

func (lb LightningBolt) OnSpellHit(sim *core.Simulation, cast *core.DirectCastAction, result *core.DirectCastDamageResult) {
	lb.OnElectricSpellHit(sim, cast, result)

	if !lb.IsLightningOverload {
		lightningOverloadChance := float64(lb.Shaman.Talents.LightningOverload) * 0.04
		if sim.Rando.Float64("LO") < lightningOverloadChance {
			overloadAction := NewLightningBolt(sim, lb.Shaman, true)
			overloadAction.Act(sim)
		}
	}
}

func NewLightningBolt(sim *core.Simulation, shaman *Shaman, IsLightningOverload bool) *core.DirectCastAction {
	baseManaCost := 300.0
	if shaman.Equip[items.ItemSlotRanged].ID == TotemOfThePulsingEarth {
		baseManaCost -= 27.0
	}

	return core.NewDirectCastAction(
		sim,
		LightningBolt{ElectricSpell{
			Shaman: shaman,
			IsLightningOverload: IsLightningOverload,
			name: "Lightning Bolt",
			baseManaCost: baseManaCost,
			baseCastTime: time.Millisecond*2500,
		}})
}
