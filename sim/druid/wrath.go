package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Starfire spell IDs
const SpellIDWrath int32 = 26985

var WrathActionID = core.ActionID{SpellID: SpellIDWrath}

const IdolAvenger int32 = 31025

func (druid *Druid) registerWrathSpell(sim *core.Simulation) {
	baseCost := 255.0

	// This seems to be unaffected by wrath of cenarius.
	bonusFlatDamage := core.TernaryFloat64(druid.Equip[items.ItemSlotRanged].ID == IdolAvenger, 25*0.571, 0)

	druid.Wrath = druid.RegisterSpell(core.SpellConfig{
		ActionID:    WrathActionID,
		SpellSchool: core.SpellSchoolNature,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost:     baseCost * (1 - 0.03*float64(druid.Talents.Moonglow)),
				GCD:      core.GCDDefault,
				CastTime: time.Second*2 - (time.Millisecond * 100 * time.Duration(druid.Talents.StarlightWrath)),
			},

			ModifyCast: func(_ *core.Simulation, _ *core.Spell, cast *core.Cast) {
				druid.applyNaturesGrace(cast)
				druid.applyNaturesSwiftness(cast)
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			BonusSpellCritRating: float64(druid.Talents.FocusedStarlight) * 2 * core.SpellCritRatingPerCritChance, // 2% crit per point
			DamageMultiplier:     1 + 0.02*float64(druid.Talents.Moonfury),
			ThreatMultiplier:     1,

			BaseDamage:     core.BaseDamageConfigMagic(383+bonusFlatDamage, 432+bonusFlatDamage, 0.571+0.02*float64(druid.Talents.WrathOfCenarius)),
			OutcomeApplier: core.OutcomeFuncMagicHitAndCrit(druid.SpellCritMultiplier(1, 0.2*float64(druid.Talents.Vengeance))),
		}),
	})
}
