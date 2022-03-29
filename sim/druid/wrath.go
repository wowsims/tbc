package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Starfire spell IDs
const SpellIDWrath int32 = 26985

const IdolAvenger int32 = 31025

func (druid *Druid) newWrathTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	baseCast := core.Cast{
		ActionID:            core.ActionID{SpellID: SpellIDWrath},
		Character:           &druid.Character,
		CritRollCategory:    core.CritRollCategoryMagical,
		OutcomeRollCategory: core.OutcomeRollCategoryMagic,
		SpellSchool:         core.SpellSchoolNature,
		BaseCost: core.ResourceCost{
			Type:  stats.Mana,
			Value: 255,
		},
		Cost: core.ResourceCost{
			Type:  stats.Mana,
			Value: 255,
		},
		CastTime:       time.Millisecond * 2000,
		GCD:            core.GCDDefault,
		CritMultiplier: druid.SpellCritMultiplier(1, 0.2*float64(druid.Talents.Vengeance)),
	}

	spellCoefficient := 0.571
	bonusFlatDamage := 0.0
	if druid.Equip[items.ItemSlotRanged].ID == IdolAvenger {
		// This seems to be unaffected by wrath of cenarius so it needs to come first.
		bonusFlatDamage += 25 * spellCoefficient
	}
	spellCoefficient += 0.02 * float64(druid.Talents.WrathOfCenarius)

	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
		},
		BaseDamage: core.BaseDamageConfigMagic(383+bonusFlatDamage, 432+bonusFlatDamage, spellCoefficient),
	}

	baseCast.CastTime -= time.Millisecond * 100 * time.Duration(druid.Talents.StarlightWrath)
	baseCast.Cost.Value -= baseCast.BaseCost.Value * 0.03 * float64(druid.Talents.Moonglow)

	effect.BonusSpellCritRating += float64(druid.Talents.FocusedStarlight) * 2 * core.SpellCritRatingPerCritChance // 2% crit per point
	effect.DamageMultiplier *= 1 + 0.02*float64(druid.Talents.Moonfury)

	effect.OnSpellHit = druid.applyOnHitTalents
	spCast := &core.SpellCast{
		Cast: baseCast,
	}

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: *spCast,
		Effect:    effect,
	})
}

func (druid *Druid) NewWrath(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	sf := &druid.wrathSpell

	druid.wrathCastTemplate.Apply(sf)

	// Modifies the cast time.
	druid.applyNaturesGrace(&sf.SpellCast)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	sf.Effect.Target = target
	sf.Init(sim)

	return sf
}
