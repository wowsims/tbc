package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDFrostbolt int32 = 27072

func (mage *Mage) newFrostboltTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            core.ActionID{SpellID: SpellIDFrostbolt},
				Character:           &mage.Character,
				CritRollCategory:    core.CritRollCategoryMagical,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				SpellSchool:         core.SpellSchoolFrost,
				SpellExtras:         SpellFlagMage | core.SpellExtrasBinary,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 330,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 330,
				},
				CastTime:       time.Millisecond * 3000,
				GCD:            core.GCDDefault,
				CritMultiplier: mage.SpellCritMultiplier(1, 0.25*float64(mage.Talents.SpellPower)+0.2*float64(mage.Talents.IceShards)),
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: mage.spellDamageMultiplier,
				ThreatMultiplier:       1 - (0.1/3)*float64(mage.Talents.FrostChanneling),
			},
			DirectInput: core.DirectDamageInput{
				MinBaseDamage:    600,
				MaxBaseDamage:    647,
				SpellCoefficient: (3.0 / 3.5) * 0.95,
			},
		},
	}

	spell.CastTime -= time.Millisecond * 100 * time.Duration(mage.Talents.ImprovedFrostbolt)
	spell.Cost.Value -= spell.BaseCost.Value * float64(mage.Talents.FrostChanneling) * 0.05
	spell.Cost.Value *= 1 - float64(mage.Talents.ElementalPrecision)*0.01
	spell.Effect.BonusSpellHitRating += float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance
	spell.Effect.BonusSpellCritRating += float64(mage.Talents.EmpoweredFrostbolt) * 1 * core.SpellCritRatingPerCritChance
	spell.Effect.StaticDamageMultiplier *= 1 + 0.02*float64(mage.Talents.PiercingIce)
	spell.Effect.StaticDamageMultiplier *= 1 + 0.01*float64(mage.Talents.ArcticWinds)
	spell.Effect.DirectInput.SpellCoefficient += 0.02 * float64(mage.Talents.EmpoweredFrostbolt)

	if ItemSetTempestRegalia.CharacterHasSetBonus(&mage.Character, 4) {
		spell.Effect.StaticDamageMultiplier *= 1.05
	}

	return core.NewSimpleSpellTemplate(spell)
}

func (mage *Mage) NewFrostbolt(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	frostbolt := &mage.frostboltSpell
	mage.frostboltCastTemplate.Apply(frostbolt)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	frostbolt.Effect.Target = target
	frostbolt.Init(sim)

	return frostbolt
}
