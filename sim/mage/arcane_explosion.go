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

	baseEffect := core.SpellEffect{
		OutcomeRollCategory: core.OutcomeRollCategoryMagic,
		CritRollCategory:    core.CritRollCategoryMagical,
		CritMultiplier:      mage.SpellCritMultiplier(1, 0.25*float64(mage.Talents.SpellPower)),
		DamageMultiplier:    mage.spellDamageMultiplier,
		ThreatMultiplier:    1 - 0.2*float64(mage.Talents.ArcaneSubtlety),
		BaseDamage:          core.BaseDamageConfigMagic(249, 270, 0.214),
	}
	baseEffect.BonusSpellHitRating += float64(mage.Talents.ArcaneFocus) * 2 * core.SpellHitRatingPerHitChance
	baseEffect.BonusSpellCritRating += float64(mage.Talents.ArcaneImpact) * 2 * core.SpellCritRatingPerCritChance

	numHits := sim.GetNumTargets()
	effects := make([]core.SpellEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)
	}
	spell.Effects = effects

	mage.ArcaneExplosion = mage.RegisterSpell(core.SpellConfig{
		Template: spell,
	})
}
