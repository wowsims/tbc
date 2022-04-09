package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDFrostbolt int32 = 27072

func (mage *Mage) registerFrostboltSpell(sim *core.Simulation) {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    core.ActionID{SpellID: SpellIDFrostbolt},
				Character:   &mage.Character,
				SpellSchool: core.SpellSchoolFrost,
				SpellExtras: SpellFlagMage | core.SpellExtrasBinary,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 330,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 330,
				},
				CastTime: time.Millisecond * 3000,
				GCD:      core.GCDDefault,
			},
		},
	}
	spell.CastTime -= time.Millisecond * 100 * time.Duration(mage.Talents.ImprovedFrostbolt)
	spell.Cost.Value -= spell.BaseCost.Value * float64(mage.Talents.FrostChanneling) * 0.05
	spell.Cost.Value *= 1 - float64(mage.Talents.ElementalPrecision)*0.01

	mage.Frostbolt = mage.RegisterSpell(core.SpellConfig{
		Template: spell,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			BonusSpellHitRating:  float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance,
			BonusSpellCritRating: float64(mage.Talents.EmpoweredFrostbolt) * 1 * core.SpellCritRatingPerCritChance,

			DamageMultiplier: mage.spellDamageMultiplier *
				(1 + 0.02*float64(mage.Talents.PiercingIce)) *
				(1 + 0.01*float64(mage.Talents.ArcticWinds)) *
				core.TernaryFloat64(ItemSetTempestRegalia.CharacterHasSetBonus(&mage.Character, 4), 1.05, 1),

			ThreatMultiplier: 1 - (0.1/3)*float64(mage.Talents.FrostChanneling),

			BaseDamage:     core.BaseDamageConfigMagic(600, 647, (3.0/3.5)*0.95+0.02*float64(mage.Talents.EmpoweredFrostbolt)),
			OutcomeApplier: core.OutcomeFuncMagicHitAndCrit(mage.SpellCritMultiplier(1, 0.25*float64(mage.Talents.SpellPower)+0.2*float64(mage.Talents.IceShards))),
		}),
	})
}
