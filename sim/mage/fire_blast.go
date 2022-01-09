package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDFireBlast int32 = 27079

var FireBlastCooldownID = core.NewCooldownID()

func (mage *Mage) newFireBlastTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				CritMultiplier: 1.5 + 0.125*float64(mage.Talents.SpellPower),
				SpellSchool:    stats.FireSpellPower,
				Character:      &mage.Character,
				BaseManaCost:   465,
				ManaCost:       465,
				CastTime:       0,
				Cooldown:       time.Second * 8,
				ActionID: core.ActionID{
					SpellID:    SpellIDFireBlast,
					CooldownID: FireBlastCooldownID,
				},
			},
		},
		SpellHitEffect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: mage.spellDamageMultiplier,
			},
			DirectInput: core.DirectDamageInput{
				MinBaseDamage:    664,
				MaxBaseDamage:    786,
				SpellCoefficient: 1.5 / 3.5,
			},
		},
	}

	spell.CastTime -= time.Millisecond * 500 * time.Duration(mage.Talents.ImprovedFireBlast)
	spell.ManaCost -= spell.BaseManaCost * float64(mage.Talents.Pyromaniac) * 0.01
	spell.SpellHitEffect.SpellEffect.BonusSpellHitRating += float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance
	spell.SpellHitEffect.SpellEffect.BonusSpellCritRating += float64(mage.Talents.CriticalMass) * 2 * core.SpellCritRatingPerCritChance
	spell.SpellHitEffect.SpellEffect.BonusSpellCritRating += float64(mage.Talents.Pyromaniac) * 1 * core.SpellCritRatingPerCritChance
	spell.SpellHitEffect.SpellEffect.StaticDamageMultiplier *= 1 + 0.02*float64(mage.Talents.FirePower)

	return core.NewSimpleSpellTemplate(spell)
}

func (mage *Mage) NewFireBlast(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	fireBlast := &mage.fireBlastSpell
	mage.fireBlastCastTemplate.Apply(fireBlast)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	fireBlast.Target = target
	fireBlast.Init(sim)

	return fireBlast
}
