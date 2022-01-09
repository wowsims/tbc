package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const (
	CastTagPyroblastDot int32 = 1
)

const SpellIDPyroblast int32 = 33938

func (mage *Mage) newPyroblastTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				CritMultiplier: 1.5 + 0.125*float64(mage.Talents.SpellPower),
				SpellSchool:    stats.FireSpellPower,
				Character:      &mage.Character,
				BaseManaCost:   500,
				ManaCost:       500,
				CastTime:       time.Millisecond * 6000,
				ActionID: core.ActionID{
					SpellID: SpellIDPyroblast,
				},
			},
		},
		SpellHitEffect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: mage.spellDamageMultiplier,
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
	spell.SpellHitEffect.SpellEffect.BonusSpellHitRating += float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance
	spell.SpellHitEffect.SpellEffect.BonusSpellCritRating += float64(mage.Talents.CriticalMass) * 2 * core.SpellCritRatingPerCritChance
	spell.SpellHitEffect.SpellEffect.BonusSpellCritRating += float64(mage.Talents.Pyromaniac) * 1 * core.SpellCritRatingPerCritChance
	spell.SpellHitEffect.SpellEffect.StaticDamageMultiplier *= 1 + 0.02*float64(mage.Talents.FirePower)

	return core.NewSimpleSpellTemplate(spell)
}

var PyroblastDotDebuffID = core.NewDebuffID()

func (mage *Mage) newPyroblastDotTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				SpellSchool: stats.FireSpellPower,
				Character:   &mage.Character,
				ActionID: core.ActionID{
					SpellID: SpellIDPyroblast,
					Tag:     CastTagPyroblastDot,
				},
				IgnoreCooldowns: true,
				IgnoreManaCost:  true,
			},
		},
		SpellHitEffect: core.SpellHitEffect{
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

	spell.SpellHitEffect.SpellEffect.StaticDamageMultiplier *= 1 + 0.02*float64(mage.Talents.FirePower)

	return core.NewSimpleSpellTemplate(spell)
}

func (mage *Mage) newPyroblastDot(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Cancel the current pyroblast dot.
	mage.pyroblastDotSpell.Cancel(sim)

	pyroblastDot := &mage.pyroblastDotSpell
	mage.pyroblastDotCastTemplate.Apply(pyroblastDot)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	pyroblastDot.Target = target
	pyroblastDot.Init(sim)

	return pyroblastDot
}

func (mage *Mage) NewPyroblast(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	pyroblast := &mage.pyroblastSpell
	mage.pyroblastCastTemplate.Apply(pyroblast)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	pyroblast.Target = target
	pyroblast.Init(sim)

	return pyroblast
}
