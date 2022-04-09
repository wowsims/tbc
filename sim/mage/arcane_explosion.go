package mage

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDArcaneExplosion int32 = 10202

func (mage *Mage) registerArcaneExplosionSpell(sim *core.Simulation) {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    core.ActionID{SpellID: SpellIDArcaneExplosion},
				Character:   &mage.Character,
				SpellSchool: core.SpellSchoolArcane,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 390,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 390,
				},
				GCD: core.GCDDefault,
			},
		},
		AOECap: 10180,
	}

	mage.ArcaneExplosion = mage.RegisterSpell(core.SpellConfig{
		Template: spell,
		ApplyEffects: core.ApplyEffectFuncAOEDamage(sim, core.SpellEffect{
			BonusSpellHitRating:  float64(mage.Talents.ArcaneFocus) * 2 * core.SpellHitRatingPerHitChance,
			BonusSpellCritRating: float64(mage.Talents.ArcaneImpact) * 2 * core.SpellCritRatingPerCritChance,

			DamageMultiplier: mage.spellDamageMultiplier,
			ThreatMultiplier: 1 - 0.2*float64(mage.Talents.ArcaneSubtlety),

			BaseDamage:     core.BaseDamageConfigMagic(249, 270, 0.214),
			OutcomeApplier: core.OutcomeFuncMagicHitAndCrit(mage.SpellCritMultiplier(1, 0.25*float64(mage.Talents.SpellPower))),
		}),
	})
}
