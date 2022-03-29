package rogue

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

var MutilateActionID = core.ActionID{SpellID: 34413}
var MutilateMHActionID = core.ActionID{SpellID: 34419}
var MutilateOHActionID = core.ActionID{SpellID: 34418}
var MutilateEnergyCost = 60.0

func (rogue *Rogue) newMutilateTemplate(_ *core.Simulation) core.SimpleSpellTemplate {
	bonusCritRating := 5 * core.MeleeCritRatingPerCritChance * float64(rogue.Talents.PuncturingWounds)

	mhDamageAbility := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            MutilateMHActionID,
				Character:           &rogue.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				CritMultiplier:      rogue.critMultiplier(true, true),
				SpellExtras:         core.SpellExtrasAlwaysHits,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ProcMask:         core.ProcMaskMeleeMHSpecial,
				DamageMultiplier: 1,
				ThreatMultiplier: 1,
				BonusCritRating:  bonusCritRating,
			},
			BaseDamage: core.BaseDamageConfigMeleeWeapon(core.MainHand, true, 101, 1, true),
		},
	}

	// cp. backstab
	mhDamageAbility.Effect.DamageMultiplier += 0.04 * float64(rogue.Talents.Opportunity)

	if ItemSetSlayers.CharacterHasSetBonus(&rogue.Character, 4) {
		mhDamageAbility.Effect.DamageMultiplier += 0.06
	}

	ohDamageAbility := mhDamageAbility
	ohDamageAbility.SpellCast.Cast.ActionID = MutilateOHActionID
	ohDamageAbility.SpellCast.Cast.CritMultiplier = rogue.critMultiplier(false, true)
	ohDamageAbility.Effect.SpellEffect.ProcMask = core.ProcMaskMeleeOHSpecial
	ohDamageAbility.Effect.BaseDamage = core.BaseDamageConfigMeleeWeapon(core.OffHand, true, 101, 1+0.1*float64(rogue.Talents.DualWieldSpecialization), true)

	mhTemplate := core.NewSimpleSpellTemplate(mhDamageAbility)
	ohTemplate := core.NewSimpleSpellTemplate(ohDamageAbility)

	mhAtk := core.SimpleSpell{}
	ohAtk := core.SimpleSpell{}

	refundAmount := MutilateEnergyCost * 0.8
	ability := rogue.newAbility(MutilateActionID, MutilateEnergyCost, SpellFlagBuilder, core.ProcMaskMeleeMHSpecial)
	ability.SpellCast.CritRollCategory = core.CritRollCategoryNone
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		if !spellEffect.Landed() {
			rogue.AddEnergy(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
			return
		}

		rogue.AddComboPoints(sim, 2, MutilateActionID)

		mhTemplate.Apply(&mhAtk)
		mhAtk.Effect.Target = spellEffect.Target

		ohTemplate.Apply(&ohAtk)
		ohAtk.Effect.Target = spellEffect.Target

		if rogue.deadlyPoisonStacks > 0 {
			mhAtk.Effect.DamageMultiplier *= 1.5
			ohAtk.Effect.DamageMultiplier *= 1.5
		}

		// TODO: while this is the most natural handling, the oh attack might have effects
		//  from the mh attack applied
		mhAtk.Cast(sim)
		ohAtk.Cast(sim)

		// applyResultsToCast() has already been done here, so we have to update the spell statistics, too
		if mhAtk.Effect.Outcome.Matches(core.OutcomeCrit) || ohAtk.Effect.Outcome.Matches(core.OutcomeCrit) {
			spellEffect.Outcome = core.OutcomeCrit
			spellCast.Hits--
			spellCast.Crits++
		}
	}

	return core.NewSimpleSpellTemplate(ability)
}

func (rogue *Rogue) NewMutilate(_ *core.Simulation, target *core.Target) *core.SimpleSpell {
	mt := &rogue.mutilate
	rogue.mutilateTemplate.Apply(mt)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	mt.Effect.Target = target

	return mt
}
