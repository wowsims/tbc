package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (mage *Mage) registerFrostboltSpell() {
	baseCost := 330.0

	mage.Frostbolt = mage.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 27072},
		SpellSchool: core.SpellSchoolFrost,
		SpellExtras: SpellFlagMage | core.SpellExtrasBinary,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost *
					(1 - 0.05*float64(mage.Talents.FrostChanneling)) *
					(1 - 0.01*float64(mage.Talents.ElementalPrecision)),

				GCD:      core.GCDDefault,
				CastTime: time.Second*3 - time.Millisecond*100*time.Duration(mage.Talents.ImprovedFrostbolt),
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask:             core.ProcMaskSpellDamage,
			BonusSpellHitRating:  float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance,
			BonusSpellCritRating: float64(mage.Talents.EmpoweredFrostbolt) * 1 * core.SpellCritRatingPerCritChance,

			DamageMultiplier: mage.spellDamageMultiplier *
				(1 + 0.02*float64(mage.Talents.PiercingIce)) *
				(1 + 0.01*float64(mage.Talents.ArcticWinds)) *
				core.TernaryFloat64(ItemSetTempestRegalia.CharacterHasSetBonus(&mage.Character, 4), 1.05, 1),

			ThreatMultiplier: 1 - (0.1/3)*float64(mage.Talents.FrostChanneling),

			BaseDamage:     core.BaseDamageConfigMagic(600, 647, (3.0/3.5)*0.95+0.02*float64(mage.Talents.EmpoweredFrostbolt)),
			OutcomeApplier: mage.OutcomeFuncMagicHitAndCrit(mage.SpellCritMultiplier(1, 0.25*float64(mage.Talents.SpellPower)+0.2*float64(mage.Talents.IceShards))),
		}),
	})
}
