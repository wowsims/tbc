package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const (
	CastTagFlamestrikeDot int32 = 1
)

const SpellIDFlamestrike int32 = 27086

var FlamestrikeActionID = core.ActionID{SpellID: SpellIDFlamestrike}

func (mage *Mage) registerFlamestrikeSpell(sim *core.Simulation) {
	//AOECap: 7830,
	baseCost := 1175.0

	applyAOEDamage := core.ApplyEffectFuncAOEDamage(sim, core.SpellEffect{
		BonusSpellHitRating: float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance,

		BonusSpellCritRating: 0 +
			float64(mage.Talents.CriticalMass)*2*core.SpellCritRatingPerCritChance +
			float64(mage.Talents.Pyromaniac)*1*core.SpellCritRatingPerCritChance +
			float64(mage.Talents.ImprovedFlamestrike)*5*core.SpellCritRatingPerCritChance,

		DamageMultiplier: mage.spellDamageMultiplier * (1 + 0.02*float64(mage.Talents.FirePower)),
		ThreatMultiplier: 1 - 0.05*float64(mage.Talents.BurningSoul),

		BaseDamage:     core.BaseDamageConfigMagic(480, 585, 0.236),
		OutcomeApplier: core.OutcomeFuncMagicHitAndCrit(mage.SpellCritMultiplier(1, 0.25*float64(mage.Talents.SpellPower))),
	})

	mage.Flamestrike = mage.RegisterSpell(core.SpellConfig{
		ActionID:    FlamestrikeActionID,
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
				CastTime: time.Second * 3,
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Target, spell *core.Spell) {
			applyAOEDamage(sim, target, spell)
			mage.FlamestrikeDot.Apply(sim)
		},
	})

	mage.FlamestrikeDot = core.NewDot(core.Dot{
		Spell: mage.Flamestrike,
		Aura: mage.RegisterAura(&core.Aura{
			Label:    "Flamestrike",
			ActionID: FlamestrikeActionID,
		}),
		NumberOfTicks: 4,
		TickLength:    time.Second * 2,
		TickEffects: core.TickFuncAOESnapshot(sim, core.SpellEffect{
			DamageMultiplier: mage.spellDamageMultiplier * (1 + 0.02*float64(mage.Talents.FirePower)),

			ThreatMultiplier: 1 - 0.05*float64(mage.Talents.BurningSoul),

			BaseDamage:     core.BaseDamageConfigMagicNoRoll(106, 0.03),
			OutcomeApplier: core.OutcomeFuncTick(),
			IsPeriodic:     true,
		}),
	})
}
