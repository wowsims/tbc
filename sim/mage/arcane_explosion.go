package mage

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDArcaneExplosion int32 = 10202

var ArcaneExplosionActionID = core.ActionID{SpellID: SpellIDArcaneExplosion}

func (mage *Mage) registerArcaneExplosionSpell(sim *core.Simulation) {
	//AOECap: 10180,
	baseCost := 390.0

	mage.ArcaneExplosion = mage.RegisterSpell(core.SpellConfig{
		ActionID:    ArcaneExplosionActionID,
		SpellSchool: core.SpellSchoolArcane,
		SpellExtras: SpellFlagMage,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost,
				GCD:  core.GCDDefault,
			},
		},

		ApplyEffects: core.ApplyEffectFuncAOEDamage(sim, core.SpellEffect{
			ProcMask:             core.ProcMaskSpellDamage,
			BonusSpellHitRating:  float64(mage.Talents.ArcaneFocus) * 2 * core.SpellHitRatingPerHitChance,
			BonusSpellCritRating: float64(mage.Talents.ArcaneImpact) * 2 * core.SpellCritRatingPerCritChance,

			DamageMultiplier: mage.spellDamageMultiplier,
			ThreatMultiplier: 1 - 0.2*float64(mage.Talents.ArcaneSubtlety),

			BaseDamage:     core.BaseDamageConfigMagic(249, 270, 0.214),
			OutcomeApplier: core.OutcomeFuncMagicHitAndCrit(mage.SpellCritMultiplier(1, 0.25*float64(mage.Talents.SpellPower))),
		}),
	})
}
