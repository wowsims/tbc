package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDArcaneBlast int32 = 30451
const ArcaneBlastBaseManaCost = 195.0

func (mage *Mage) newArcaneBlastTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				Name:           "Arcane Blast",
				CritMultiplier: 1.5 + 0.125*float64(mage.Talents.SpellPower),
				SpellSchool:    stats.ArcaneSpellPower,
				Character:      &mage.Character,
				BaseManaCost:   ArcaneBlastBaseManaCost,
				ManaCost:       ArcaneBlastBaseManaCost,
				CastTime:       time.Millisecond * 2500,
				ActionID: core.ActionID{
					SpellID: SpellIDArcaneBlast,
				},
				OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
					newNumStacks := core.MinInt32(3, mage.NumStacks(ArcaneBlastAuraID)+1)
					cast.Character.ReplaceAura(sim, ArcaneBlastAura(sim, newNumStacks))
				},
			},
		},
		SpellHitEffect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: mage.spellDamageMultiplier,
			},
			DirectInput: core.DirectDamageInput{
				MinBaseDamage:    668,
				MaxBaseDamage:    772,
				SpellCoefficient: 2.5 / 3.5,
			},
		},
	}

	spell.SpellHitEffect.SpellEffect.BonusSpellHitRating += float64(mage.Talents.ArcaneFocus) * 2 * core.SpellHitRatingPerHitChance
	spell.SpellHitEffect.SpellEffect.BonusSpellCritRating += float64(mage.Talents.ArcaneImpact) * 2 * core.SpellCritRatingPerCritChance

	if ItemSetTirisfalRegalia.CharacterHasSetBonus(&mage.Character, 2) {
		spell.SpellHitEffect.SpellEffect.StaticDamageMultiplier *= 1.2
		spell.ManaCost += 0.2 * ArcaneBlastBaseManaCost
	}

	return core.NewSimpleSpellTemplate(spell)
}

func (mage *Mage) NewArcaneBlast(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	arcaneBlast := &mage.arcaneBlastSpell
	mage.arcaneBlastCastTemplate.Apply(arcaneBlast)

	numStacks := mage.NumStacks(ArcaneBlastAuraID)
	arcaneBlast.CastTime -= time.Duration(numStacks) * time.Second / 3
	arcaneBlast.ManaCost += float64(numStacks) * ArcaneBlastBaseManaCost * 0.75

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	arcaneBlast.Target = target
	arcaneBlast.Init(sim)

	return arcaneBlast
}

var ArcaneBlastAuraID = core.NewAuraID()

func ArcaneBlastAura(sim *core.Simulation, numStacks int32) core.Aura {
	return core.Aura{
		ID:      ArcaneBlastAuraID,
		Name:    "Arcane Blast",
		SpellID: 36032,
		Expires: sim.CurrentTime + time.Second*8,
		Stacks:  numStacks,
	}
}
