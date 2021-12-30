package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDFireball int32 = 27070

func (mage *Mage) newFireballTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				Name:           "Fireball",
				CritMultiplier: 1.5 + 0.125*float64(mage.Talents.SpellPower),
				SpellSchool:    stats.FireSpellPower,
				Character:      &mage.Character,
				BaseManaCost:   425,
				ManaCost:       425,
				CastTime:       time.Millisecond * 3500,
				ActionID: core.ActionID{
					SpellID: SpellIDFireball,
				},
			},
		},
		SpellHitEffect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: mage.spellDamageMultiplier,
			},
			DirectInput: core.DirectDamageInput{
				MinBaseDamage:    649,
				MaxBaseDamage:    821,
				SpellCoefficient: 1.0,
			},
		},
	}

	spell.CastTime -= time.Millisecond * 100 * time.Duration(mage.Talents.ImprovedFireball)
	spell.ManaCost -= spell.BaseManaCost * float64(mage.Talents.Pyromaniac) * 0.01
	spell.SpellHitEffect.SpellEffect.BonusSpellHitRating += float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance
	spell.SpellHitEffect.SpellEffect.BonusSpellCritRating += float64(mage.Talents.CriticalMass) * 2 * core.SpellCritRatingPerCritChance
	spell.SpellHitEffect.SpellEffect.BonusSpellCritRating += float64(mage.Talents.Pyromaniac) * 1 * core.SpellCritRatingPerCritChance
	spell.SpellHitEffect.SpellEffect.StaticDamageMultiplier *= 1 + 0.02*float64(mage.Talents.FirePower)
	spell.SpellHitEffect.DirectInput.SpellCoefficient += 0.03 * float64(mage.Talents.EmpoweredFireball)

	if ItemSetTempestRegalia.CharacterHasSetBonus(&mage.Character, 4) {
		spell.SpellHitEffect.SpellEffect.StaticDamageMultiplier *= 1.05
	}

	// TODO: Fireball built-in dot

	return core.NewSimpleSpellTemplate(spell)
}

func (mage *Mage) NewFireball(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	fireball := &mage.fireballSpell
	mage.fireballCastTemplate.Apply(fireball)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	fireball.Target = target
	fireball.Init(sim)

	return fireball
}
