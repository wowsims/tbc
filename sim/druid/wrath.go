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

func (druid *Druid) registerWrathSpell(sim *core.Simulation) {
	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    core.ActionID{SpellID: SpellIDWrath},
				Character:   &druid.Character,
				SpellSchool: core.SpellSchoolNature,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 255,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 255,
				},
				CastTime: time.Millisecond * 2000,
				GCD:      core.GCDDefault,
			},
		},
	}

	template.CastTime -= time.Millisecond * 100 * time.Duration(druid.Talents.StarlightWrath)
	template.Cost.Value -= template.BaseCost.Value * 0.03 * float64(druid.Talents.Moonglow)

	// This seems to be unaffected by wrath of cenarius.
	bonusFlatDamage := core.TernaryFloat64(druid.Equip[items.ItemSlotRanged].ID == IdolAvenger, 25*0.571, 0)

	druid.Wrath = druid.RegisterSpell(core.SpellConfig{
		Template: template,
		ModifyCast: func(sim *core.Simulation, target *core.Target, instance *core.SimpleSpell) {
			druid.applyNaturesGrace(&instance.SpellCast)
			druid.applyNaturesSwiftness(&instance.SpellCast)
		},
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			BonusSpellCritRating: float64(druid.Talents.FocusedStarlight) * 2 * core.SpellCritRatingPerCritChance, // 2% crit per point
			DamageMultiplier:     1 + 0.02*float64(druid.Talents.Moonfury),
			ThreatMultiplier:     1,

			BaseDamage:     core.BaseDamageConfigMagic(383+bonusFlatDamage, 432+bonusFlatDamage, 0.571+0.02*float64(druid.Talents.WrathOfCenarius)),
			OutcomeApplier: core.OutcomeFuncMagicHitAndCrit(druid.SpellCritMultiplier(1, 0.2*float64(druid.Talents.Vengeance))),
			OnSpellHit:     druid.applyOnHitTalents,
		}),
	})
}
