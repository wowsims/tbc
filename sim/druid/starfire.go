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

func (druid *Druid) newStarfireSpell(sim *core.Simulation, rank int) *core.SimpleSpellTemplate {
	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
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
			},
		},
	}

	minBaseDamage := 550.0
	maxBaseDamage := 647.0
	spellCoefficient := 1.0
	if rank == 6 {
		template.BaseCost.Value = 315
		template.Cost.Value = 315
		template.ActionID = core.ActionID{
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

	template.Effect = core.SpellEffect{
		OutcomeRollCategory: core.OutcomeRollCategoryMagic,
		CritRollCategory:    core.CritRollCategoryMagical,
		CritMultiplier:      druid.SpellCritMultiplier(1, 0.2*float64(druid.Talents.Vengeance)),
		DamageMultiplier:    1,
		ThreatMultiplier:    1,
		BaseDamage:          core.BaseDamageConfigMagic(minBaseDamage+bonusFlatDamage, maxBaseDamage+bonusFlatDamage, spellCoefficient),
	}

	if ItemSetNordrassil.CharacterHasSetBonus(&druid.Character, 4) {
		template.Effect.BaseDamage = core.WrapBaseDamageConfig(template.Effect.BaseDamage, func(oldCalculator core.BaseDamageCalculator) core.BaseDamageCalculator {
			return func(sim *core.Simulation, hitEffect *core.SpellEffect, spell *core.SimpleSpellTemplate) float64 {
				normalDamage := oldCalculator(sim, hitEffect, spell)

				// Check if moonfire/insectswarm is ticking on the target.
				// TODO: in a raid simulator we need to be able to see which dots are ticking from other druids.
				if (druid.Moonfire.Instance.Effect.DotInput.IsTicking(sim) && druid.Moonfire.Instance.Effect.Target == hitEffect.Target) ||
					(druid.InsectSwarm.Instance.Effect.DotInput.IsTicking(sim) && druid.InsectSwarm.Instance.Effect.Target == hitEffect.Target) {
					return normalDamage * 1.1
				} else {
					return normalDamage
				}
			}
		})
	}

	template.Cost.Value -= template.BaseCost.Value * 0.03 * float64(druid.Talents.Moonglow)
	template.CastTime -= time.Millisecond * 100 * time.Duration(druid.Talents.StarlightWrath)

	template.Effect.DamageMultiplier *= 1 + 0.02*float64(druid.Talents.Moonfury)
	template.Effect.BonusSpellCritRating += float64(druid.Talents.FocusedStarlight) * 2 * core.SpellCritRatingPerCritChance // 2% crit per point

	if ItemSetThunderheart.CharacterHasSetBonus(&druid.Character, 4) { // Thunderheart 4p adds 5% crit to starfire
		template.Effect.BonusSpellCritRating += 5 * core.SpellCritRatingPerCritChance
	}

	template.Effect.OnSpellHit = druid.applyOnHitTalents

	return druid.RegisterSpell(core.SpellConfig{
		Template: template,
		ModifyCast: func(sim *core.Simulation, target *core.Target, instance *core.SimpleSpell) {
			instance.Effect.Target = target
			druid.applyNaturesGrace(&instance.SpellCast)
			druid.applyNaturesSwiftness(&instance.SpellCast)
		},
	})
}
