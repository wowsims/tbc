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
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    FlamestrikeActionID,
				Character:   &mage.Character,
				SpellSchool: core.SpellSchoolFire,
				SpellExtras: SpellFlagMage,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 1175,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 1175,
				},
				CastTime: time.Second * 3,
				GCD:      core.GCDDefault,
			},
		},
		AOECap: 7830,
	}
	spell.Cost.Value -= spell.BaseCost.Value * float64(mage.Talents.Pyromaniac) * 0.01
	spell.Cost.Value *= 1 - float64(mage.Talents.ElementalPrecision)*0.01

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
		Template: spell,
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
