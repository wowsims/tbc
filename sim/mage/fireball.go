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

func (mage *Mage) registerFireballSpell(sim *core.Simulation) {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    core.ActionID{SpellID: SpellIDFireball},
				Character:   &mage.Character,
				SpellSchool: core.SpellSchoolFire,
				SpellExtras: SpellFlagMage,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 425,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 425,
				},
				CastTime: time.Millisecond * 3500,
				GCD:      core.GCDDefault,
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
			CritMultiplier:      mage.SpellCritMultiplier(1, 0.25*float64(mage.Talents.SpellPower)),
			DamageMultiplier:    mage.spellDamageMultiplier,
			ThreatMultiplier:    1 - 0.05*float64(mage.Talents.BurningSoul),
			OnSpellHit: func(sim *core.Simulation, spell *core.SimpleSpellTemplate, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					mage.FireballDot.Instance.Cancel(sim)
					mage.FireballDot.Cast(sim, spellEffect.Target)
				}
			},
			BaseDamage: core.BaseDamageConfigMagic(649, 821, 1.0+0.03*float64(mage.Talents.EmpoweredFireball)),
		},
	}

	spell.CastTime -= time.Millisecond * 100 * time.Duration(mage.Talents.ImprovedFireball)
	spell.Cost.Value -= spell.BaseCost.Value * float64(mage.Talents.Pyromaniac) * 0.01
	spell.Cost.Value *= 1 - float64(mage.Talents.ElementalPrecision)*0.01
	spell.Effect.BonusSpellHitRating += float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance
	spell.Effect.BonusSpellCritRating += float64(mage.Talents.CriticalMass) * 2 * core.SpellCritRatingPerCritChance
	spell.Effect.BonusSpellCritRating += float64(mage.Talents.Pyromaniac) * 1 * core.SpellCritRatingPerCritChance
	spell.Effect.DamageMultiplier *= 1 + 0.02*float64(mage.Talents.FirePower)

	if ItemSetTempestRegalia.CharacterHasSetBonus(&mage.Character, 4) {
		spell.Effect.DamageMultiplier *= 1.05
	}

	mage.Fireball = mage.RegisterSpell(core.SpellConfig{
		Template:   spell,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}

var FireballDotDebuffID = core.NewDebuffID()

func (mage *Mage) registerFireballDotSpell(sim *core.Simulation) {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID: core.ActionID{
					SpellID: SpellIDFireball,
					Tag:     CastTagFireballDot,
				},
				Character:   &mage.Character,
				SpellSchool: core.SpellSchoolFire,
				SpellExtras: SpellFlagMage,
			},
		},
		Effect: core.SpellEffect{
			DamageMultiplier: mage.spellDamageMultiplier,
			DotInput: core.DotDamageInput{
				NumberOfTicks:  4,
				TickLength:     time.Second * 2,
				TickBaseDamage: core.DotSnapshotFuncMagic(84/4, 0),
				DebuffID:       FireballDotDebuffID,
			},
		},
	}

	spell.Effect.DamageMultiplier *= 1 + 0.02*float64(mage.Talents.FirePower)

	if ItemSetTempestRegalia.CharacterHasSetBonus(&mage.Character, 4) {
		spell.Effect.DamageMultiplier *= 1.05
	}

	mage.FireballDot = mage.RegisterSpell(core.SpellConfig{
		Template:   spell,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}
