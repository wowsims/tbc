package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDFireBlast int32 = 27079

var FireBlastCooldownID = core.NewCooldownID()

func (mage *Mage) registerFireBlastSpell(sim *core.Simulation) {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID: core.ActionID{
					SpellID:    SpellIDFireBlast,
					CooldownID: FireBlastCooldownID,
				},
				Character:   &mage.Character,
				SpellSchool: core.SpellSchoolFire,
				SpellExtras: SpellFlagMage,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 465,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 465,
				},
				GCD:      core.GCDDefault,
				Cooldown: time.Second * 8,
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
			CritMultiplier:      mage.SpellCritMultiplier(1, 0.25*float64(mage.Talents.SpellPower)),
			DamageMultiplier:    mage.spellDamageMultiplier,
			ThreatMultiplier:    1 - 0.05*float64(mage.Talents.BurningSoul),
			BaseDamage:          core.BaseDamageConfigMagic(664, 786, 1.5/3.5),
		},
	}

	spell.CastTime -= time.Millisecond * 500 * time.Duration(mage.Talents.ImprovedFireBlast)
	spell.Cost.Value -= spell.BaseCost.Value * float64(mage.Talents.Pyromaniac) * 0.01
	spell.Cost.Value *= 1 - float64(mage.Talents.ElementalPrecision)*0.01
	spell.Effect.BonusSpellHitRating += float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance
	spell.Effect.BonusSpellCritRating += float64(mage.Talents.CriticalMass) * 2 * core.SpellCritRatingPerCritChance
	spell.Effect.BonusSpellCritRating += float64(mage.Talents.Pyromaniac) * 1 * core.SpellCritRatingPerCritChance
	spell.Effect.DamageMultiplier *= 1 + 0.02*float64(mage.Talents.FirePower)

	mage.FireBlast = mage.RegisterSpell(core.SpellConfig{
		Template:   spell,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}
