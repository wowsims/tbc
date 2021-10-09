package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDCL6 int32 = 25442

type ChainLightning struct {
	shaman *Shaman

	isLightningOverload bool
}

func (cl ChainLightning) GetActionID() core.ActionID {
	return core.ActionID{
		SpellID: SpellIDCL6,
		CooldownID: core.MagicIDChainLightning6,
	}
}

func (cl ChainLightning) GetName() string {
	if cl.isLightningOverload {
		return "Chain Lightning (LO)"
	} else {
		return "Chain Lightning"
	}
}

func (cl ChainLightning) GetTag() int32 {
	if cl.isLightningOverload {
		return CastTagLightningOverload
	} else {
		return 0
	}
}

func (cl ChainLightning) GetAgent() core.Agent {
	return cl.shaman
}

func (cl ChainLightning) GetBaseManaCost() float64 {
	return 760
}

func (cl ChainLightning) GetSpellSchool() stats.Stat {
	return stats.NatureSpellPower
}

func (cl ChainLightning) GetCooldown() time.Duration {
	return time.Second*6
}

func (cl ChainLightning) GetCastInput(sim *core.Simulation, cast *core.DirectCastAction) core.DirectCastInput {
	input := core.DirectCastInput{
		ManaCost: cl.GetBaseManaCost(),
		CastTime: time.Millisecond*2000,
	}

	if cl.isLightningOverload {
		input.ManaCost = 0
		input.CastTime = 0
	} else if cl.shaman.Talents.LightningMastery > 0 {
		// Convection applies against the base cost of the spell.
		input.ManaCost -= cl.GetBaseManaCost() * cl.shaman.convectionBonus

		input.CastTime -= time.Millisecond * 100 * time.Duration(cl.shaman.Talents.LightningMastery)
	}

	if cl.shaman.elementalFocusStacks > 0 {
		// TODO: This should subtract 40% of base cost
		input.ManaCost *= .6 // reduced by 40%
	}

	return input
}

func (cl ChainLightning) GetHitInputs(sim *core.Simulation, cast *core.DirectCastAction) []core.DirectCastDamageInput{
	hitInput := core.DirectCastDamageInput{
		MinBaseDamage: 734,
		MaxBaseDamage: 838,
		SpellCoefficient: 0.651,
		CritMultiplier: 1.5,
		DamageMultiplier: 1,
	}

	hitInput.DamageMultiplier *= cl.shaman.concussionBonus
	if cl.isLightningOverload {
		hitInput.DamageMultiplier *= 0.5
	}

	if cl.shaman.EquippedMetaGem(core.ChaoticSkyfireDiamond) {
		hitInput.CritMultiplier *= 1.03
	}
	if cl.shaman.Talents.ElementalFury {
		hitInput.CritMultiplier *= 2
		hitInput.CritMultiplier -= 1
	}

	hitInput.BonusHit += float64(cl.shaman.Talents.ElementalPrecision) * 0.02
	hitInput.BonusHit += float64(cl.shaman.Talents.NaturesGuidance) * 0.01
	hitInput.BonusCrit += float64(cl.shaman.Talents.TidalMastery) * 0.01
	hitInput.BonusCrit += float64(cl.shaman.Talents.CallOfThunder) * 0.01

	if cl.shaman.Equip[items.ItemSlotRanged].ID == TotemOfStorms {
		hitInput.BonusSpellPower += 33
	} else if cl.shaman.Equip[items.ItemSlotRanged].ID == TotemOfTheVoid {
		hitInput.BonusSpellPower += 55
	} else if cl.shaman.Equip[items.ItemSlotRanged].ID == TotemOfAncestralGuidance {
		hitInput.BonusSpellPower += 85
	}

	numHits := core.MinInt32(3, sim.Options.Encounter.NumTargets)
	hitInputs := make([]core.DirectCastDamageInput, 0, numHits)
	hitInputs = append(hitInputs, hitInput)

	chainMultiplier := 1.0
	for i := int32(1); i < numHits; i++ {
		if cl.shaman.HasAura(core.MagicIDTidefury) {
			chainMultiplier *= 0.83
		} else {
			chainMultiplier *= 0.7
		}

		bounceHit := hitInputs[0] // Makes a copy
		bounceHit.DamageMultiplier *= chainMultiplier
		hitInputs = append(hitInputs, bounceHit)
	}

	return hitInputs
}

func (cl ChainLightning) OnCastComplete(sim *core.Simulation, cast *core.DirectCastAction) {
	if !cl.isLightningOverload && cl.shaman.elementalFocusStacks > 0 {
		cl.shaman.elementalFocusStacks--
	}
}
func (cl ChainLightning) OnSpellHit(sim *core.Simulation, cast *core.DirectCastAction, result *core.DirectCastDamageResult) {
	if !cl.isLightningOverload {
		lightningOverloadChance := float64(cl.shaman.Talents.LightningOverload) * 0.04 / 3
		if sim.Rando.Float64("LO") < lightningOverloadChance {
			overloadAction := NewChainLightning(sim, cl.shaman, true)
			overloadAction.Act(sim)
		}
	}
}
func (cl ChainLightning) OnSpellMiss(sim *core.Simulation, cast *core.DirectCastAction) {
}

func NewChainLightning(sim *core.Simulation, shaman *Shaman, isLightningOverload bool) *core.DirectCastAction {
	return core.NewDirectCastAction(
		sim,
		ChainLightning{
			shaman: shaman,
			isLightningOverload: isLightningOverload,
		})
}
