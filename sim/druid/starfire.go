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

func (druid *Druid) newStarfireSpell(sim *core.Simulation, rank int) *core.Spell {
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
	template.Cost.Value -= template.BaseCost.Value * 0.03 * float64(druid.Talents.Moonglow)
	template.CastTime -= time.Millisecond * 100 * time.Duration(druid.Talents.StarlightWrath)

	// This seems to be unaffected by wrath of cenarius so it needs to come first.
	bonusFlatDamage := core.TernaryFloat64(druid.Equip[items.ItemSlotRanged].ID == IvoryMoongoddess, 55*spellCoefficient, 0)
	spellCoefficient += 0.04 * float64(druid.Talents.WrathOfCenarius)

	effect := core.SpellEffect{
		BonusSpellCritRating: (float64(druid.Talents.FocusedStarlight) * 2 * core.SpellCritRatingPerCritChance) + core.TernaryFloat64(ItemSetThunderheart.CharacterHasSetBonus(&druid.Character, 4), 5*core.SpellCritRatingPerCritChance, 0),
		DamageMultiplier:     1 + 0.02*float64(druid.Talents.Moonfury),
		ThreatMultiplier:     1,
		BaseDamage:           core.BaseDamageConfigMagic(minBaseDamage+bonusFlatDamage, maxBaseDamage+bonusFlatDamage, spellCoefficient),
		OutcomeApplier:       core.OutcomeFuncMagicHitAndCrit(druid.SpellCritMultiplier(1, 0.2*float64(druid.Talents.Vengeance))),
		OnSpellHit:           druid.applyOnHitTalents,
	}

	if ItemSetNordrassil.CharacterHasSetBonus(&druid.Character, 4) {
		effect.BaseDamage = core.WrapBaseDamageConfig(effect.BaseDamage, func(oldCalculator core.BaseDamageCalculator) core.BaseDamageCalculator {
			return func(sim *core.Simulation, hitEffect *core.SpellEffect, spell *core.Spell) float64 {
				normalDamage := oldCalculator(sim, hitEffect, spell)

				// Check if moonfire/insectswarm is ticking on the target.
				// TODO: in a raid simulator we need to be able to see which dots are ticking from other druids.
				if druid.MoonfireDot.IsActive() || druid.InsectSwarmDot.IsActive() {
					return normalDamage * 1.1
				} else {
					return normalDamage
				}
			}
		})
	}

	return druid.RegisterSpell(core.SpellConfig{
		Template: template,
		ModifyCast: func(sim *core.Simulation, target *core.Target, instance *core.SimpleSpell) {
			instance.Effect.Target = target
			druid.applyNaturesGrace(&instance.SpellCast)
			druid.applyNaturesSwiftness(&instance.SpellCast)
		},
		ApplyEffects: core.ApplyEffectFuncDirectDamage(effect),
	})
}
