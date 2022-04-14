package mage

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDPyroblast int32 = 33938

var PyroblastActionID = core.ActionID{SpellID: SpellIDPyroblast}

func (mage *Mage) registerPyroblastSpell(sim *core.Simulation) {
	baseCost := 500.0

	mage.Pyroblast = mage.RegisterSpell(core.SpellConfig{
		ActionID:    PyroblastActionID,
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
				CastTime: time.Second * 6,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			BonusSpellHitRating: float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance,

			BonusSpellCritRating: 0 +
				float64(mage.Talents.CriticalMass)*2*core.SpellCritRatingPerCritChance +
				float64(mage.Talents.Pyromaniac)*1*core.SpellCritRatingPerCritChance,

			DamageMultiplier: mage.spellDamageMultiplier * (1 + 0.02*float64(mage.Talents.FirePower)),

			ThreatMultiplier: 1 - 0.05*float64(mage.Talents.BurningSoul),

			BaseDamage:     core.BaseDamageConfigMagic(939, 1191, 1.15),
			OutcomeApplier: core.OutcomeFuncMagicHitAndCrit(mage.SpellCritMultiplier(1, 0.25*float64(mage.Talents.SpellPower))),

			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					mage.PyroblastDot.Apply(sim)
				}
			},
		}),
	})

	target := sim.GetPrimaryTarget()
	mage.PyroblastDot = core.NewDot(core.Dot{
		Spell: mage.Pyroblast,
		Aura: target.RegisterAura(core.Aura{
			Label:    "Pyroblast-" + strconv.Itoa(int(mage.Index)),
			ActionID: PyroblastActionID,
		}),
		NumberOfTicks: 4,
		TickLength:    time.Second * 3,
		TickEffects: core.TickFuncSnapshot(target, core.SpellEffect{
			DamageMultiplier: mage.spellDamageMultiplier * (1 + 0.02*float64(mage.Talents.FirePower)),

			ThreatMultiplier: 1 - 0.05*float64(mage.Talents.BurningSoul),

			BaseDamage:     core.BaseDamageConfigFlat(356 / 4),
			OutcomeApplier: core.OutcomeFuncTick(),
			IsPeriodic:     true,
		}),
	})
}
