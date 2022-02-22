package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

const (
	CastTagPyroblastDot int32 = 1
)

const SpellIDPyroblast int32 = 33938

func (mage *Mage) newPyroblastTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            core.ActionID{SpellID: SpellIDPyroblast},
				Character:           &mage.Character,
				CritRollCategory:    core.CritRollCategoryMagical,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				SpellSchool:         core.SpellSchoolFire,
				BaseManaCost:        500,
				ManaCost:            500,
				CastTime:            time.Millisecond * 6000,
				GCD:                 core.GCDDefault,
				CritMultiplier:      mage.SpellCritMultiplier(1, 0.25*float64(mage.Talents.SpellPower)),
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: mage.spellDamageMultiplier,
				ThreatMultiplier:       1 - 0.05*float64(mage.Talents.BurningSoul),
				OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
					pyroblastDot := mage.newPyroblastDot(sim, spellEffect.Target)
					pyroblastDot.Cast(sim)
				},
			},
			DirectInput: core.DirectDamageInput{
				MinBaseDamage:    939,
				MaxBaseDamage:    1191,
				SpellCoefficient: 1.15,
			},
		},
	}

	spell.ManaCost -= spell.BaseManaCost * float64(mage.Talents.Pyromaniac) * 0.01
	spell.ManaCost *= 1 - float64(mage.Talents.ElementalPrecision)*0.01
	spell.Effect.BonusSpellHitRating += float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance
	spell.Effect.BonusSpellCritRating += float64(mage.Talents.CriticalMass) * 2 * core.SpellCritRatingPerCritChance
	spell.Effect.BonusSpellCritRating += float64(mage.Talents.Pyromaniac) * 1 * core.SpellCritRatingPerCritChance
	spell.Effect.StaticDamageMultiplier *= 1 + 0.02*float64(mage.Talents.FirePower)

	return core.NewSimpleSpellTemplate(spell)
}

var PyroblastDotDebuffID = core.NewDebuffID()

func (mage *Mage) newPyroblastDotTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID: core.ActionID{
					SpellID: SpellIDPyroblast,
					Tag:     CastTagPyroblastDot,
				},
				Character:           &mage.Character,
				CritRollCategory:    core.CritRollCategoryMagical,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				SpellSchool:         core.SpellSchoolFire,
				IgnoreManaCost:      true,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: mage.spellDamageMultiplier,
				IgnoreHitCheck:         true,
			},
			DotInput: core.DotDamageInput{
				NumberOfTicks:        4,
				TickLength:           time.Second * 3,
				TickBaseDamage:       356 / 4,
				TickSpellCoefficient: 0,
				DebuffID:             PyroblastDotDebuffID,
			},
		},
	}

	spell.Effect.StaticDamageMultiplier *= 1 + 0.02*float64(mage.Talents.FirePower)

	return core.NewSimpleSpellTemplate(spell)
}

func (mage *Mage) newPyroblastDot(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Cancel the current pyroblast dot.
	mage.pyroblastDotSpell.Cancel(sim)

	pyroblastDot := &mage.pyroblastDotSpell
	mage.pyroblastDotCastTemplate.Apply(pyroblastDot)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	pyroblastDot.Effect.Target = target
	pyroblastDot.Init(sim)

	return pyroblastDot
}

func (mage *Mage) NewPyroblast(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	pyroblast := &mage.pyroblastSpell
	mage.pyroblastCastTemplate.Apply(pyroblast)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	pyroblast.Effect.Target = target
	pyroblast.Init(sim)

	return pyroblast
}
