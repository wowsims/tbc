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

func (mage *Mage) registerPyroblastSpell(sim *core.Simulation) {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    core.ActionID{SpellID: SpellIDPyroblast},
				Character:   &mage.Character,
				SpellSchool: core.SpellSchoolFire,
				SpellExtras: SpellFlagMage,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 500,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 500,
				},
				CastTime: time.Millisecond * 6000,
				GCD:      core.GCDDefault,
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
			CritMultiplier:      mage.SpellCritMultiplier(1, 0.25*float64(mage.Talents.SpellPower)),
			DamageMultiplier:    mage.spellDamageMultiplier,
			ThreatMultiplier:    1 - 0.05*float64(mage.Talents.BurningSoul),
			BaseDamage:          core.BaseDamageConfigMagic(939, 1191, 1.15),
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					mage.PyroblastDot.Instance.Cancel(sim)
					mage.PyroblastDot.Cast(sim, spellEffect.Target)
				}
			},
		},
	}

	spell.Cost.Value -= spell.BaseCost.Value * float64(mage.Talents.Pyromaniac) * 0.01
	spell.Cost.Value *= 1 - float64(mage.Talents.ElementalPrecision)*0.01
	spell.Effect.BonusSpellHitRating += float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance
	spell.Effect.BonusSpellCritRating += float64(mage.Talents.CriticalMass) * 2 * core.SpellCritRatingPerCritChance
	spell.Effect.BonusSpellCritRating += float64(mage.Talents.Pyromaniac) * 1 * core.SpellCritRatingPerCritChance
	spell.Effect.DamageMultiplier *= 1 + 0.02*float64(mage.Talents.FirePower)

	mage.Pyroblast = mage.RegisterSpell(core.SpellConfig{
		Template:   spell,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}

var PyroblastDotDebuffID = core.NewDebuffID()

func (mage *Mage) registerPyroblastDotSpell(sim *core.Simulation) {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID: core.ActionID{
					SpellID: SpellIDPyroblast,
					Tag:     CastTagPyroblastDot,
				},
				Character:   &mage.Character,
				SpellSchool: core.SpellSchoolFire,
				SpellExtras: SpellFlagMage,
			},
		},
		Effect: core.SpellEffect{
			DamageMultiplier: mage.spellDamageMultiplier,
			ThreatMultiplier: 1,
			DotInput: core.DotDamageInput{
				NumberOfTicks:  4,
				TickLength:     time.Second * 3,
				TickBaseDamage: core.DotSnapshotFuncMagic(356/4, 0),
				DebuffID:       PyroblastDotDebuffID,
			},
		},
	}

	spell.Effect.DamageMultiplier *= 1 + 0.02*float64(mage.Talents.FirePower)

	mage.PyroblastDot = mage.RegisterSpell(core.SpellConfig{
		Template:   spell,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}
