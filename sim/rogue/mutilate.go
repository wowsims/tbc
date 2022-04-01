package rogue

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

var MutilateActionID = core.ActionID{SpellID: 34413}
var MutilateMHActionID = core.ActionID{SpellID: 34419}
var MutilateOHActionID = core.ActionID{SpellID: 34418}
var MutilateEnergyCost = 60.0

func (rogue *Rogue) registerMutilateSpell(_ *core.Simulation) {
	bonusCritRating := 5 * core.MeleeCritRatingPerCritChance * float64(rogue.Talents.PuncturingWounds)

	mhDamageAbility := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    MutilateMHActionID,
				Character:   &rogue.Character,
				SpellSchool: core.SpellSchoolPhysical,
				SpellExtras: core.SpellExtrasAlwaysHits,
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategorySpecial,
			CritRollCategory:    core.CritRollCategoryPhysical,
			CritMultiplier:      rogue.critMultiplier(true, true),
			ProcMask:            core.ProcMaskMeleeMHSpecial,
			DamageMultiplier:    1,
			ThreatMultiplier:    1,
			BonusCritRating:     bonusCritRating,
			BaseDamage:          core.BaseDamageConfigMeleeWeapon(core.MainHand, true, 101, 1, true),
		},
	}

	// cp. backstab
	mhDamageAbility.Effect.DamageMultiplier += 0.04 * float64(rogue.Talents.Opportunity)

	if ItemSetSlayers.CharacterHasSetBonus(&rogue.Character, 4) {
		mhDamageAbility.Effect.DamageMultiplier += 0.06
	}

	ohDamageAbility := mhDamageAbility
	ohDamageAbility.SpellCast.Cast.ActionID = MutilateOHActionID
	ohDamageAbility.Effect.CritMultiplier = rogue.critMultiplier(false, true)
	ohDamageAbility.Effect.ProcMask = core.ProcMaskMeleeOHSpecial
	ohDamageAbility.Effect.BaseDamage = core.BaseDamageConfigMeleeWeapon(core.OffHand, true, 101, 1+0.1*float64(rogue.Talents.DualWieldSpecialization), true)

	modifyCast := func(sim *core.Simulation, target *core.Target, instance *core.SimpleSpell) {
		instance.Effect.Target = target
		if rogue.deadlyPoisonStacks > 0 {
			instance.Effect.DamageMultiplier *= 1.5
		}
	}
	mhHitSpell := rogue.RegisterSpell(core.SpellConfig{
		Template:   mhDamageAbility,
		ModifyCast: modifyCast,
	})
	ohHitSpell := rogue.RegisterSpell(core.SpellConfig{
		Template:   ohDamageAbility,
		ModifyCast: modifyCast,
	})

	refundAmount := MutilateEnergyCost * 0.8
	ability := rogue.newAbility(MutilateActionID, MutilateEnergyCost, SpellFlagBuilder, core.ProcMaskMeleeMHSpecial)
	ability.Effect.CritRollCategory = core.CritRollCategoryNone
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
		if !spellEffect.Landed() {
			rogue.AddEnergy(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
			return
		}

		rogue.AddComboPoints(sim, 2, MutilateActionID)

		// TODO: while this is the most natural handling, the oh attack might have effects
		//  from the mh attack applied
		mhHitSpell.Cast(sim, spellEffect.Target)
		ohHitSpell.Cast(sim, spellEffect.Target)

		if mhHitSpell.Instance.Effect.Outcome.Matches(core.OutcomeCrit) || ohHitSpell.Instance.Effect.Outcome.Matches(core.OutcomeCrit) {
			spellEffect.Outcome = core.OutcomeCrit
			rogue.Mutilate.Hits--
			rogue.Mutilate.Crits++
		}
	}

	rogue.Mutilate = rogue.RegisterSpell(core.SpellConfig{
		Template:   ability,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}
