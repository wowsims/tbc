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

func (shaman *Shaman) NewElectricSpell(name string, actionID core.ActionID, baseManaCost float64, baseCastTime time.Duration, isLightningOverload bool) *core.DirectCastAction {
	spell := &shaman.electricSpell
	if isLightningOverload {
		spell = &shaman.electricSpellLO
	}

	// Clear data from previous use of object
	*spell = core.DirectCastAction{}

	spell.Cast.Name = name
	spell.Cast.ActionID = actionID
	spell.Cast.Character = shaman.GetCharacter()
	spell.Cast.BaseManaCost = baseManaCost
	spell.Cast.ManaCost = baseManaCost
	spell.Cast.CastTime = baseCastTime
	spell.Cast.SpellSchool = stats.NatureSpellPower
	spell.Cast.CritMultiplier = 1.5

	if isLightningOverload {
		spell.Cast.Name += " (LO)"
		spell.Cast.Tag = CastTagLightningOverload
		spell.Cast.CastTime = 0
		spell.Cast.ManaCost = 0
		spell.Cast.IgnoreCooldowns = true
		spell.Cast.IgnoreManaCost = true
	} else if shaman.Talents.LightningMastery > 0 {
		// Convection applies against the base cost of the spell.
		spell.Cast.ManaCost -= spell.BaseManaCost * shaman.convectionBonus

		spell.Cast.CastTime -= time.Millisecond * 100 * time.Duration(shaman.Talents.LightningMastery)
	}

	if shaman.ElementalFocusStacks > 0 {
		// TODO: This should subtract 40% of base cost
		spell.Cast.ManaCost *= .6 // reduced by 40%
	}

	if shaman.Talents.ElementalFury {
		spell.Cast.CritMultiplier = 2
	}

	spell.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		shaman.OnElectricSpellCastComplete(sim, cast, isLightningOverload)
	}
	spell.OnSpellMiss = func(sim *core.Simulation, cast *core.Cast) {}

	return spell
}

func (shaman *Shaman) ApplyElectricSpellHitInputModifiers(hitInput *core.DirectCastDamageInput, isLightningOverload bool) {
	hitInput.DamageMultiplier *= shaman.concussionBonus
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

func (shaman *Shaman) OnElectricSpellCastComplete(sim *core.Simulation, cast *core.Cast, isLightningOverload bool) {
	if !isLightningOverload && shaman.ElementalFocusStacks > 0 {
		shaman.ElementalFocusStacks--
	}
}
func (shaman *Shaman) OnElectricSpellHit(sim *core.Simulation, cast *core.Cast, result *core.DirectCastDamageResult) {
	if shaman.Talents.ElementalFocus && result.Crit {
		shaman.ElementalFocusStacks = 2
	}
}
