package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Totem Item IDs
const (
	TotemOfAncestralGuidance = 32330
	TotemOfStorms            = 23199
	TotemOfThePulsingEarth   = 29389
	TotemOfTheVoid           = 28248
)

const (
	CastTagLightningOverload int32 = 1 // This could be value or bitflag if we ended up needing multiple flags at the same time.
)

const SpellIDLB12 int32 = 25449

type LightningBolt struct {
	shaman *Shaman

	isLightningOverload bool
}

func (lb LightningBolt) GetActionID() core.ActionID {
	return core.ActionID{
		SpellID: SpellIDLB12,
	}
}

func (lb LightningBolt) GetName() string {
	if lb.isLightningOverload {
		return "Lightning Bolt (LO)"
	} else {
		return "Lightning Bolt"
	}
}

func (lb LightningBolt) GetTag() int32 {
	if lb.isLightningOverload {
		return CastTagLightningOverload
	} else {
		return 0
	}
}

func (lb LightningBolt) GetAgent() core.Agent {
	return lb.shaman
}

func (lb LightningBolt) GetBaseManaCost() float64 {
	if lb.shaman.Equip[items.ItemSlotRanged].ID == TotemOfThePulsingEarth {
		return 300 - 27
	} else {
		return 300
	}
}

func (lb LightningBolt) GetSpellSchool() stats.Stat {
	return stats.NatureSpellPower
}

func (lb LightningBolt) GetCooldown() time.Duration {
	return 0
}

func (lb LightningBolt) GetCastInput(sim *core.Simulation, cast *core.DirectCastAction) core.DirectCastInput {
	input := core.DirectCastInput{
		ManaCost: lb.GetBaseManaCost(),
		CastTime: time.Millisecond*2500,
	}

	if lb.isLightningOverload {
		input.ManaCost = 0
		input.CastTime = 0
	} else if lb.shaman.Talents.LightningMastery > 0 {
		// Convection applies against the base cost of the spell.
		input.ManaCost -= lb.GetBaseManaCost() * lb.shaman.convectionBonus

		input.CastTime -= time.Millisecond * 100 * time.Duration(lb.shaman.Talents.LightningMastery)
	}

	if lb.shaman.elementalFocusStacks > 0 {
		// TODO: This should subtract 40% of base cost
		input.ManaCost *= .6 // reduced by 40%
	}

	return input
}

func (lb LightningBolt) GetHitInputs(sim *core.Simulation, cast *core.DirectCastAction) []core.DirectCastDamageInput{
	hitInput := core.DirectCastDamageInput{
		MinBaseDamage: 571,
		MaxBaseDamage: 652,
		SpellCoefficient: 0.794,
		CritMultiplier: 1.5,
		DamageMultiplier: 1,
	}

	hitInput.DamageMultiplier *= lb.shaman.concussionBonus
	if lb.isLightningOverload {
		hitInput.DamageMultiplier *= 0.5
	}

	if lb.shaman.EquippedMetaGem(core.ChaoticSkyfireDiamond) {
		hitInput.CritMultiplier *= 1.03
	}
	if lb.shaman.Talents.ElementalFury {
		hitInput.CritMultiplier *= 2
		hitInput.CritMultiplier -= 1
	}

	hitInput.BonusHit += float64(lb.shaman.Talents.ElementalPrecision) * 0.02
	hitInput.BonusHit += float64(lb.shaman.Talents.NaturesGuidance) * 0.01
	hitInput.BonusCrit += float64(lb.shaman.Talents.TidalMastery) * 0.01
	hitInput.BonusCrit += float64(lb.shaman.Talents.CallOfThunder) * 0.01

	if lb.shaman.Equip[items.ItemSlotRanged].ID == TotemOfStorms {
		hitInput.BonusSpellPower += 33
	} else if lb.shaman.Equip[items.ItemSlotRanged].ID == TotemOfTheVoid {
		hitInput.BonusSpellPower += 55
	} else if lb.shaman.Equip[items.ItemSlotRanged].ID == TotemOfAncestralGuidance {
		hitInput.BonusSpellPower += 85
	}

	return []core.DirectCastDamageInput{hitInput}
}

func (lb LightningBolt) OnCastComplete(sim *core.Simulation, cast *core.DirectCastAction) {
	if !lb.isLightningOverload && lb.shaman.elementalFocusStacks > 0 {
		lb.shaman.elementalFocusStacks--
	}
}
func (lb LightningBolt) OnSpellHit(sim *core.Simulation, cast *core.DirectCastAction, result *core.DirectCastDamageResult) {
	if result.Crit {
		lb.shaman.elementalFocusStacks = 2
	}

	if !lb.isLightningOverload {
		lightningOverloadChance := float64(lb.shaman.Talents.LightningOverload) * 0.04
		if sim.Rando.Float64("LO") < lightningOverloadChance {
			overloadAction := NewLightningBolt(sim, lb.shaman, true)
			overloadAction.Act(sim)
		}
	}
}
func (lb LightningBolt) OnSpellMiss(sim *core.Simulation, cast *core.DirectCastAction) {
}

func NewLightningBolt(sim *core.Simulation, shaman *Shaman, isLightningOverload bool) *core.DirectCastAction {
	return core.NewDirectCastAction(
		sim,
		LightningBolt{
			shaman: shaman,
			isLightningOverload: isLightningOverload,
		})
}
