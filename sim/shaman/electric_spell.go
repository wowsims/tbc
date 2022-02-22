package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
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
	spellCast := core.SpellCast{
		Cast: core.Cast{
			ActionID:            actionID,
			Character:           shaman.GetCharacter(),
			CritRollCategory:    core.CritRollCategoryMagical,
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			SpellSchool:         core.SpellSchoolNature,
			BaseManaCost:        baseManaCost,
			ManaCost:            baseManaCost,
			CastTime:            baseCastTime,
			GCD:                 core.GCDDefault,
			CritMultiplier:      shaman.DefaultSpellCritMultiplier(),
		},
	}

	if shaman.Talents.ElementalFury {
		spellCast.CritMultiplier = shaman.SpellCritMultiplier(1, 1)
	}

	if isLightningOverload {
		spellCast.ActionID.Tag = CastTagLightningOverload
		spellCast.CastTime = 0
		spellCast.ManaCost = 0
		spellCast.GCD = 0
		spellCast.IgnoreManaCost = true
	} else if shaman.Talents.LightningMastery > 0 {
		// Convection applies against the base cost of the spell.
		spellCast.ManaCost -= spellCast.BaseManaCost * float64(shaman.Talents.Convection) * 0.02
		spellCast.CastTime -= time.Millisecond * 100 * time.Duration(shaman.Talents.LightningMastery)
	}

	if !isLightningOverload && shaman.Talents.ElementalFocus {
		spellCast.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
			if shaman.ElementalFocusStacks > 0 {
				shaman.ElementalFocusStacks--
			}
		}
	} else {
		spellCast.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {}
	}

	return spellCast
}

// Helper for precomputing spell effects.
func (shaman *Shaman) newElectricSpellEffect(minBaseDamage float64, maxBaseDamage float64, spellCoefficient float64, isLightningOverload bool) core.SpellHitEffect {
	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
			ThreatMultiplier:       1,
		},
		DirectInput: core.DirectDamageInput{
			MinBaseDamage:    minBaseDamage,
			MaxBaseDamage:    maxBaseDamage,
			SpellCoefficient: spellCoefficient,
		},
	}

	effect.SpellEffect.DamageMultiplier *= 1 + 0.01*float64(shaman.Talents.Concussion)
	if isLightningOverload {
		effect.SpellEffect.DamageMultiplier *= 0.5
		effect.SpellEffect.ThreatMultiplier = 0
	}

	effect.ThreatMultiplier *= 1 - (0.1/3)*float64(shaman.Talents.ElementalPrecision)
	effect.SpellEffect.BonusSpellHitRating += float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance
	effect.SpellEffect.BonusSpellCritRating += float64(shaman.Talents.TidalMastery) * 1 * core.SpellCritRatingPerCritChance
	effect.SpellEffect.BonusSpellCritRating += float64(shaman.Talents.CallOfThunder) * 1 * core.SpellCritRatingPerCritChance

	if shaman.Equip[items.ItemSlotRanged].ID == TotemOfStorms {
		effect.SpellEffect.BonusSpellPower += 33
	} else if shaman.Equip[items.ItemSlotRanged].ID == TotemOfTheVoid {
		effect.SpellEffect.BonusSpellPower += 55
	} else if shaman.Equip[items.ItemSlotRanged].ID == TotemOfAncestralGuidance {
		effect.SpellEffect.BonusSpellPower += 85
	}

	return effect
}

// Shared LB/CL logic that is dynamic, i.e. can't be precomputed.
func (shaman *Shaman) applyElectricSpellCastInitModifiers(spellCast *core.SpellCast) {
	if shaman.ElementalFocusStacks > 0 {
		// Reduces mana cost by 40%
		spellCast.Cast.ManaCost -= spellCast.Cast.BaseManaCost * 0.4
	}
}
