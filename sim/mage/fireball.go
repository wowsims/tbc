package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const (
	CastTagFireballDot int32 = 1
)

const SpellIDFireball int32 = 27070

func (mage *Mage) newFireballTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            core.ActionID{SpellID: SpellIDFireball},
				Character:           &mage.Character,
				CritRollCategory:    core.CritRollCategoryMagical,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				SpellSchool:         core.SpellSchoolFire,
				SpellExtras:         SpellFlagMage,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 425,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 425,
				},
				CastTime:       time.Millisecond * 3500,
				GCD:            core.GCDDefault,
				CritMultiplier: mage.SpellCritMultiplier(1, 0.25*float64(mage.Talents.SpellPower)),
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: mage.spellDamageMultiplier,
				ThreatMultiplier:       1 - 0.05*float64(mage.Talents.BurningSoul),
				OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
					if !spellEffect.Landed() {
						return
					}
					fireballDot := mage.newFireballDot(sim, spellEffect.Target)
					fireballDot.Cast(sim)
				},
			},
			BaseDamage: core.BaseDamageFuncMagic(649, 821, 1.0+0.03*float64(mage.Talents.EmpoweredFireball)),
		},
	}

	spell.CastTime -= time.Millisecond * 100 * time.Duration(mage.Talents.ImprovedFireball)
	spell.Cost.Value -= spell.BaseCost.Value * float64(mage.Talents.Pyromaniac) * 0.01
	spell.Cost.Value *= 1 - float64(mage.Talents.ElementalPrecision)*0.01
	spell.Effect.BonusSpellHitRating += float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance
	spell.Effect.BonusSpellCritRating += float64(mage.Talents.CriticalMass) * 2 * core.SpellCritRatingPerCritChance
	spell.Effect.BonusSpellCritRating += float64(mage.Talents.Pyromaniac) * 1 * core.SpellCritRatingPerCritChance
	spell.Effect.StaticDamageMultiplier *= 1 + 0.02*float64(mage.Talents.FirePower)

	if ItemSetTempestRegalia.CharacterHasSetBonus(&mage.Character, 4) {
		spell.Effect.StaticDamageMultiplier *= 1.05
	}

	return core.NewSimpleSpellTemplate(spell)
}

var FireballDotDebuffID = core.NewDebuffID()

func (mage *Mage) newFireballDotTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID: core.ActionID{
					SpellID: SpellIDFireball,
					Tag:     CastTagFireballDot,
				},
				Character:        &mage.Character,
				CritRollCategory: core.CritRollCategoryMagical,
				SpellSchool:      core.SpellSchoolFire,
				SpellExtras:      SpellFlagMage,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: mage.spellDamageMultiplier,
			},
			DotInput: core.DotDamageInput{
				NumberOfTicks:  4,
				TickLength:     time.Second * 2,
				TickBaseDamage: core.DotSnapshotFuncMagic(84/4, 0),
				DebuffID:       FireballDotDebuffID,
			},
		},
	}

	spell.Effect.StaticDamageMultiplier *= 1 + 0.02*float64(mage.Talents.FirePower)

	if ItemSetTempestRegalia.CharacterHasSetBonus(&mage.Character, 4) {
		spell.Effect.StaticDamageMultiplier *= 1.05
	}

	return core.NewSimpleSpellTemplate(spell)
}

func (mage *Mage) newFireballDot(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Cancel the current fireball dot.
	mage.fireballDotSpell.Cancel(sim)

	fireballDot := &mage.fireballDotSpell
	mage.fireballDotCastTemplate.Apply(fireballDot)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	fireballDot.Effect.Target = target
	fireballDot.Init(sim)

	return fireballDot
}

func (mage *Mage) NewFireball(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	fireball := &mage.fireballSpell
	mage.fireballCastTemplate.Apply(fireball)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	fireball.Effect.Target = target
	fireball.Init(sim)

	return fireball
}
