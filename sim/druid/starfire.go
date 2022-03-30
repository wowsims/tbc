package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Starfire spell IDs
const SpellIDSF8 int32 = 26986
const SpellIDSF6 int32 = 9876

// Idol IDs
const IvoryMoongoddess int32 = 27518

func (druid *Druid) newStarfireTemplate(sim *core.Simulation, rank int) core.SimpleSpellTemplate {
	baseCast := core.Cast{
		ActionID:    core.ActionID{SpellID: SpellIDSF8},
		Character:   &druid.Character,
		SpellSchool: core.SpellSchoolArcane,
		BaseCost: core.ResourceCost{
			Type:  stats.Mana,
			Value: 370,
		},
		Cost: core.ResourceCost{
			Type:  stats.Mana,
			Value: 370,
		},
		CastTime: time.Millisecond * 3500,
		GCD:      core.GCDDefault,
	}

	minBaseDamage := 550.0
	maxBaseDamage := 647.0
	spellCoefficient := 1.0
	if rank == 6 {
		baseCast.BaseCost.Value = 315
		baseCast.Cost.Value = 315
		baseCast.ActionID = core.ActionID{
			SpellID: SpellIDSF6,
		}
		minBaseDamage = 463
		maxBaseDamage = 543
		spellCoefficient = 0.99
	}
	bonusFlatDamage := 0.0
	if druid.Equip[items.ItemSlotRanged].ID == IvoryMoongoddess {
		// This seems to be unaffected by wrath of cenarius so it needs to come first.
		bonusFlatDamage += 55 * spellCoefficient
	}
	spellCoefficient += 0.04 * float64(druid.Talents.WrathOfCenarius)

	effect := core.SpellEffect{
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BaseDamage:       core.BaseDamageConfigMagic(minBaseDamage+bonusFlatDamage, maxBaseDamage+bonusFlatDamage, spellCoefficient),
	}

	if ItemSetNordrassil.CharacterHasSetBonus(&druid.Character, 4) {
		effect.BaseDamage = core.WrapBaseDamageConfig(effect.BaseDamage, func(oldCalculator core.BaseDamageCalculator) core.BaseDamageCalculator {
			return func(sim *core.Simulation, hitEffect *core.SpellEffect, spellCast *core.SpellCast) float64 {
				normalDamage := oldCalculator(sim, hitEffect, spellCast)

				// Check if moonfire/insectswarm is ticking on the target.
				// TODO: in a raid simulator we need to be able to see which dots are ticking from other druids.
				if (druid.MoonfireSpell.Effect.DotInput.IsTicking(sim) && druid.MoonfireSpell.Effect.Target == hitEffect.Target) ||
					(druid.InsectSwarmSpell.Effect.DotInput.IsTicking(sim) && druid.InsectSwarmSpell.Effect.Target == hitEffect.Target) {
					return normalDamage * 1.1
				} else {
					return normalDamage
				}
			}
		})
	}

	baseCast.Cost.Value -= baseCast.BaseCost.Value * 0.03 * float64(druid.Talents.Moonglow)
	baseCast.CastTime -= time.Millisecond * 100 * time.Duration(druid.Talents.StarlightWrath)

	effect.DamageMultiplier *= 1 + 0.02*float64(druid.Talents.Moonfury)
	effect.BonusSpellCritRating += float64(druid.Talents.FocusedStarlight) * 2 * core.SpellCritRatingPerCritChance // 2% crit per point

	if ItemSetThunderheart.CharacterHasSetBonus(&druid.Character, 4) { // Thunderheart 4p adds 5% crit to starfire
		effect.BonusSpellCritRating += 5 * core.SpellCritRatingPerCritChance
	}

	effect.OnSpellHit = druid.applyOnHitTalents
	spCast := &core.SpellCast{
		Cast:                baseCast,
		OutcomeRollCategory: core.OutcomeRollCategoryMagic,
		CritRollCategory:    core.CritRollCategoryMagical,
		CritMultiplier:      druid.SpellCritMultiplier(1, 0.2*float64(druid.Talents.Vengeance)),
	}

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: *spCast,
		Effect:    effect,
	})
}

func (druid *Druid) NewStarfire(sim *core.Simulation, target *core.Target, rank int) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	sf := &druid.starfireSpell

	if rank == 8 {
		druid.starfire8CastTemplate.Apply(sf)
	} else if rank == 6 {
		druid.starfire6CastTemplate.Apply(sf)
	}

	// Applies nature's grace cast time reduction if available.
	druid.applyNaturesGrace(&sf.SpellCast)
	druid.applyNaturesSwiftness(&sf.SpellCast)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	sf.Effect.Target = target
	sf.Init(sim)

	return sf
}
