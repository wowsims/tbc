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

func (shaman *Shaman) newElectricSpellTemplate(name string, actionID core.ActionID, baseManaCost float64, baseCastTime time.Duration, isLightningOverload bool) core.DirectCastAction {
	template := core.DirectCastAction{
		Cast: core.Cast{
			Name: name,
			ActionID: actionID,
			Character: shaman.GetCharacter(),
			BaseManaCost: baseManaCost,
			ManaCost: baseManaCost,
			CastTime: baseCastTime,
			SpellSchool: stats.NatureSpellPower,
			CritMultiplier: 1.5,
		},
	}


	if isLightningOverload {
		template.Cast.Name += " (LO)"
		template.Cast.Tag = CastTagLightningOverload
		template.Cast.CastTime = 0
		template.Cast.ManaCost = 0
		template.Cast.IgnoreCooldowns = true
		template.Cast.IgnoreManaCost = true
	} else if shaman.Talents.LightningMastery > 0 {
		// Convection applies against the base cost of the spell.
		template.Cast.ManaCost -= template.BaseManaCost * float64(shaman.Talents.Convection) * 0.02

		template.Cast.CastTime -= time.Millisecond * 100 * time.Duration(shaman.Talents.LightningMastery)
	}

	if shaman.Talents.ElementalFury {
		template.Cast.CritMultiplier = 2
	}

	if !isLightningOverload && shaman.Talents.ElementalFocus {
		template.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
			if shaman.ElementalFocusStacks > 0 {
				shaman.ElementalFocusStacks--
			}
		}
	} else {
		template.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {}
	}

	template.OnSpellMiss = func(sim *core.Simulation, cast *core.Cast) {}

	return template
}

func (shaman *Shaman) applyElectricSpellHitInputModifiers(hitInput *core.DirectCastDamageInput, isLightningOverload bool) {
	hitInput.DamageMultiplier *= 1 + 0.01*float64(shaman.Talents.Concussion)
	if isLightningOverload {
		hitInput.DamageMultiplier *= 0.5
	}

	hitInput.BonusHit += float64(shaman.Talents.ElementalPrecision) * 0.02
	hitInput.BonusHit += float64(shaman.Talents.NaturesGuidance) * 0.01
	hitInput.BonusCrit += float64(shaman.Talents.TidalMastery) * 0.01
	hitInput.BonusCrit += float64(shaman.Talents.CallOfThunder) * 0.01

	if shaman.Equip[items.ItemSlotRanged].ID == TotemOfStorms {
		hitInput.BonusSpellPower += 33
	} else if shaman.Equip[items.ItemSlotRanged].ID == TotemOfTheVoid {
		hitInput.BonusSpellPower += 55
	} else if shaman.Equip[items.ItemSlotRanged].ID == TotemOfAncestralGuidance {
		hitInput.BonusSpellPower += 85
	}
}

func (shaman *Shaman) applyElectricSpellInitModifiers(spell *core.DirectCastAction) {
	if shaman.ElementalFocusStacks > 0 {
		// TODO: This should subtract 40% of base cost
		spell.Cast.ManaCost *= .6 // reduced by 40%
	}
}
