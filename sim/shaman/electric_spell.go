package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Totem Item IDs
const (
	StormfuryTotem           = 31031
	TotemOfAncestralGuidance = 32330
	TotemOfImpact            = 27947
	TotemOfStorms            = 23199
	TotemOfThePulsingEarth   = 29389
	TotemOfTheVoid           = 28248
	TotemOfRage              = 22395
)

const (
	CastTagLightningOverload int32 = 1 // This could be value or bitflag if we ended up needing multiple flags at the same time.
)

// Mana cost numbers based on in-game testing:
//
// With 5/5 convection:
// Normal: 270, w/ EF: 150
//
// With 5/5 convection and TotPE equipped:
// Normal: 246, w/ EF: 136

// Shared precomputation logic for LB and CL.
func (shaman *Shaman) newElectricSpellCast(actionID core.ActionID, baseManaCost float64, baseCastTime time.Duration, isLightningOverload bool) core.SpellCast {
	cost := core.ResourceCost{Type: stats.Mana, Value: baseManaCost}
	spell := core.SpellCast{
		Cast: core.Cast{
			ActionID:    actionID,
			Character:   shaman.GetCharacter(),
			SpellSchool: core.SpellSchoolNature,
			BaseCost:    cost,
			Cost:        cost,
			CastTime:    baseCastTime,
			GCD:         core.GCDDefault,
			SpellExtras: SpellFlagElectric,
		},
	}

	if isLightningOverload {
		spell.ActionID.Tag = CastTagLightningOverload
		spell.CastTime = 0
		spell.GCD = 0
		spell.Cost.Value = 0
	} else if shaman.Talents.LightningMastery > 0 {
		// Convection applies against the base cost of the spell.
		spell.Cost.Value -= spell.BaseCost.Value * float64(shaman.Talents.Convection) * 0.02
		spell.CastTime -= time.Millisecond * 100 * time.Duration(shaman.Talents.LightningMastery)
	}

	if !isLightningOverload && shaman.Talents.ElementalFocus {
		spell.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
			if shaman.ElementalFocusStacks > 0 {
				shaman.ElementalFocusStacks--
			}
		}
	} else {
		spell.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {}
	}

	return spell
}

// Helper for precomputing spell effects.
func (shaman *Shaman) newElectricSpellEffect(minBaseDamage float64, maxBaseDamage float64, spellCoefficient float64, isLightningOverload bool) core.SpellEffect {
	effect := core.SpellEffect{
		OutcomeRollCategory: core.OutcomeRollCategoryMagic,
		CritRollCategory:    core.CritRollCategoryMagical,
		CritMultiplier:      shaman.DefaultSpellCritMultiplier(),
		DamageMultiplier:    1,
		ThreatMultiplier:    1,
		BaseDamage:          core.BaseDamageConfigMagic(minBaseDamage, maxBaseDamage, spellCoefficient),
	}

	if shaman.Talents.ElementalFury {
		effect.CritMultiplier = shaman.SpellCritMultiplier(1, 1)
	}

	effect.DamageMultiplier *= 1 + 0.01*float64(shaman.Talents.Concussion)
	if isLightningOverload {
		effect.DamageMultiplier *= 0.5
		effect.ThreatMultiplier = 0
	}

	effect.ThreatMultiplier *= 1 - (0.1/3)*float64(shaman.Talents.ElementalPrecision)
	effect.BonusSpellHitRating += float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance
	effect.BonusSpellCritRating += float64(shaman.Talents.TidalMastery) * 1 * core.SpellCritRatingPerCritChance
	effect.BonusSpellCritRating += float64(shaman.Talents.CallOfThunder) * 1 * core.SpellCritRatingPerCritChance

	if shaman.Equip[items.ItemSlotRanged].ID == TotemOfStorms {
		effect.BonusSpellPower += 33
	} else if shaman.Equip[items.ItemSlotRanged].ID == TotemOfTheVoid {
		effect.BonusSpellPower += 55
	} else if shaman.Equip[items.ItemSlotRanged].ID == TotemOfAncestralGuidance {
		effect.BonusSpellPower += 85
	}

	return effect
}

// Shared LB/CL logic that is dynamic, i.e. can't be precomputed.
func (shaman *Shaman) applyElectricSpellCastInitModifiers(spellCast *core.SpellCast) {
	if shaman.ElementalFocusStacks > 0 {
		// Reduces mana cost by 40%
		spellCast.Cost.Value -= spellCast.BaseCost.Value * 0.4
	}
	if shaman.HasAura(ElementalMasteryAuraID) {
		spellCast.Cost.Value = 0
	}
}
