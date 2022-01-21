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

func (mage *Mage) newFlamestrikeTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				CritMultiplier: 1.5 + 0.125*float64(mage.Talents.SpellPower),
				SpellSchool:    stats.FireSpellPower,
				Character:      &mage.Character,
				BaseManaCost:   1175,
				ManaCost:       1175,
				CastTime:       time.Second * 3,
				ActionID: core.ActionID{
					SpellID: SpellIDFlamestrike,
				},
				OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
					flamestrikeDot := mage.newFlamestrikeDot(sim)
					flamestrikeDot.Cast(sim)
				},
			},
		},
		AOECap: 7830,
	}

	baseEffect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier:       1,
			StaticDamageMultiplier: mage.spellDamageMultiplier,
		},
		DirectInput: core.DirectDamageInput{
			MinBaseDamage:    480,
			MaxBaseDamage:    585,
			SpellCoefficient: 0.236,
		},
	}

	spell.ManaCost -= spell.BaseManaCost * float64(mage.Talents.Pyromaniac) * 0.01
	baseEffect.BonusSpellHitRating += float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance
	baseEffect.BonusSpellCritRating += float64(mage.Talents.CriticalMass) * 2 * core.SpellCritRatingPerCritChance
	baseEffect.BonusSpellCritRating += float64(mage.Talents.Pyromaniac) * 1 * core.SpellCritRatingPerCritChance
	baseEffect.BonusSpellCritRating += float64(mage.Talents.ImprovedFlamestrike) * 5 * core.SpellCritRatingPerCritChance
	baseEffect.StaticDamageMultiplier *= 1 + 0.02*float64(mage.Talents.FirePower)

	numHits := sim.GetNumTargets()
	effects := make([]core.SpellHitEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)
	}
	spell.Effects = effects

	return core.NewSimpleSpellTemplate(spell)
}

func (mage *Mage) newFlamestrikeDotTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				CritMultiplier: 1.5 + 0.125*float64(mage.Talents.SpellPower),
				SpellSchool:    stats.FireSpellPower,
				Character:      &mage.Character,
				ActionID: core.ActionID{
					SpellID: SpellIDFlamestrike,
					Tag:     CastTagFlamestrikeDot,
				},
				IgnoreCooldowns: true,
				IgnoreManaCost:  true,
			},
		},
	}

	baseEffect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier:       1,
			StaticDamageMultiplier: mage.spellDamageMultiplier,
			IgnoreHitCheck:         true,
		},
		DotInput: core.DotDamageInput{
			NumberOfTicks:        4,
			TickLength:           time.Second * 2,
			TickBaseDamage:       106,
			TickSpellCoefficient: 0.03,
		},
	}

	baseEffect.BonusSpellHitRating += float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance
	baseEffect.StaticDamageMultiplier *= 1 + 0.02*float64(mage.Talents.FirePower)

	numHits := sim.GetNumTargets()
	effects := make([]core.SpellHitEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)
	}
	spell.Effects = effects

	return core.NewSimpleSpellTemplate(spell)
}

func (mage *Mage) newFlamestrikeDot(sim *core.Simulation) *core.SimpleSpell {
	// Cancel the current flamestrike dot.
	mage.flamestrikeDotSpell.Cancel(sim)

	flamestrikeDot := &mage.flamestrikeDotSpell
	mage.flamestrikeDotCastTemplate.Apply(flamestrikeDot)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	flamestrikeDot.Init(sim)

	return flamestrikeDot
}

func (mage *Mage) NewFlamestrike(sim *core.Simulation) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	flamestrike := &mage.flamestrikeSpell
	mage.flamestrikeCastTemplate.Apply(flamestrike)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	flamestrike.Init(sim)

	return flamestrike
}
