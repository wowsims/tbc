package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDScorch int32 = 27074

var ScorchActionID = core.ActionID{SpellID: SpellIDScorch}

func (mage *Mage) registerScorchSpell(sim *core.Simulation) {
	baseCost := 180.0

	effect := core.SpellEffect{
		ProcMask:            core.ProcMaskSpellDamage,
		BonusSpellHitRating: float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance,

		BonusSpellCritRating: 0 +
			float64(mage.Talents.Incineration)*2*core.SpellCritRatingPerCritChance +
			float64(mage.Talents.CriticalMass)*2*core.SpellCritRatingPerCritChance +
			float64(mage.Talents.Pyromaniac)*1*core.SpellCritRatingPerCritChance,

		DamageMultiplier: mage.spellDamageMultiplier * (1 + 0.02*float64(mage.Talents.FirePower)),
		ThreatMultiplier: 1 - 0.05*float64(mage.Talents.BurningSoul),

		BaseDamage:     core.BaseDamageConfigMagic(305, 361, 1.5/3.5),
		OutcomeApplier: core.OutcomeFuncMagicHitAndCrit(mage.SpellCritMultiplier(1, 0.25*float64(mage.Talents.SpellPower))),
	}

	if mage.Talents.ImprovedScorch > 0 {
		mage.ScorchAura = sim.GetPrimaryTarget().GetAura(core.ImprovedScorchAuraLabel)
		if mage.ScorchAura == nil {
			mage.ScorchAura = core.ImprovedScorchAura(sim.GetPrimaryTarget(), 0)
		}

		procChance := float64(mage.Talents.ImprovedScorch) / 3.0
		effect.OnSpellHit = func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() {
				return
			}

			if procChance != 1.0 || sim.RandomFloat("Improved Scorch") > procChance {
				return
			}

			mage.ScorchAura.Activate(sim)
			mage.ScorchAura.AddStack(sim)
		}
	}

	mage.Scorch = mage.RegisterSpell(core.SpellConfig{
		ActionID:    ScorchActionID,
		SpellSchool: core.SpellSchoolFire,
		SpellExtras: SpellFlagMage,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost *
					(1 - 0.01*float64(mage.Talents.Pyromaniac)) *
					(1 - 0.01*float64(mage.Talents.ElementalPrecision)),

				GCD:      core.GCDDefault,
				CastTime: time.Millisecond * 1500,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(effect),
	})
}
