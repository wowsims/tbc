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

func (mage *Mage) registerFlamestrikeSpell(sim *core.Simulation) {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    core.ActionID{SpellID: SpellIDFlamestrike},
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
				OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
					mage.FlamestrikeDot.Instance.Cancel(sim)
					mage.FlamestrikeDot.Cast(sim, nil)
				},
			},
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
			CritMultiplier:      mage.SpellCritMultiplier(1, 0.25*float64(mage.Talents.SpellPower)),
		},
		AOECap: 7830,
	}

	baseEffect := core.SpellEffect{
		DamageMultiplier: mage.spellDamageMultiplier,
		ThreatMultiplier: 1 - 0.05*float64(mage.Talents.BurningSoul),
		BaseDamage:       core.BaseDamageConfigMagic(480, 585, 0.236),
	}

	spell.Cost.Value -= spell.BaseCost.Value * float64(mage.Talents.Pyromaniac) * 0.01
	spell.Cost.Value *= 1 - float64(mage.Talents.ElementalPrecision)*0.01
	baseEffect.BonusSpellHitRating += float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance
	baseEffect.BonusSpellCritRating += float64(mage.Talents.CriticalMass) * 2 * core.SpellCritRatingPerCritChance
	baseEffect.BonusSpellCritRating += float64(mage.Talents.Pyromaniac) * 1 * core.SpellCritRatingPerCritChance
	baseEffect.BonusSpellCritRating += float64(mage.Talents.ImprovedFlamestrike) * 5 * core.SpellCritRatingPerCritChance
	baseEffect.DamageMultiplier *= 1 + 0.02*float64(mage.Talents.FirePower)

	numHits := sim.GetNumTargets()
	effects := make([]core.SpellEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)
	}
	spell.Effects = effects

	mage.Flamestrike = mage.RegisterSpell(core.SpellConfig{
		Template: spell,
	})
}

func (mage *Mage) registerFlamestrikeDotSpell(sim *core.Simulation) {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID: core.ActionID{
					SpellID: SpellIDFlamestrike,
					Tag:     CastTagFlamestrikeDot,
				},
				Character:   &mage.Character,
				SpellSchool: core.SpellSchoolFire,
				SpellExtras: SpellFlagMage | core.SpellExtrasAlwaysHits,
			},
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
		},
	}

	baseEffect := core.SpellEffect{
		DamageMultiplier: mage.spellDamageMultiplier,
		ThreatMultiplier: 1,
		DotInput: core.DotDamageInput{
			NumberOfTicks:  4,
			TickLength:     time.Second * 2,
			TickBaseDamage: core.DotSnapshotFuncMagic(106, 0.03),
		},
	}

	baseEffect.BonusSpellHitRating += float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance
	baseEffect.DamageMultiplier *= 1 + 0.02*float64(mage.Talents.FirePower)

	numHits := sim.GetNumTargets()
	effects := make([]core.SpellEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)
	}
	spell.Effects = effects

	mage.FlamestrikeDot = mage.RegisterSpell(core.SpellConfig{
		Template:   spell,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}
