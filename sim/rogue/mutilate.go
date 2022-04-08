package rogue

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

var MutilateActionID = core.ActionID{SpellID: 34413}
var MutilateMHActionID = core.ActionID{SpellID: 34419}
var MutilateOHActionID = core.ActionID{SpellID: 34418}
var MutilateEnergyCost = 60.0

func (rogue *Rogue) newMutilateHitSpell(isMH bool) *core.Spell {
	ability := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    MutilateMHActionID,
				Character:   &rogue.Character,
				SpellSchool: core.SpellSchoolPhysical,
				SpellExtras: core.SpellExtrasAlwaysHits,
			},
		},
	}
	if !isMH {
		ability.ActionID = MutilateOHActionID
	}

	effect := core.SpellEffect{
		ProcMask: core.ProcMaskMeleeMHSpecial,

		BonusCritRating: 5 * core.MeleeCritRatingPerCritChance * float64(rogue.Talents.PuncturingWounds),
		DamageMultiplier: 1 +
			0.04*float64(rogue.Talents.Opportunity) +
			core.TernaryFloat64(ItemSetSlayers.CharacterHasSetBonus(&rogue.Character, 4), 0.06, 0),
		ThreatMultiplier: 1,

		BaseDamage:     core.BaseDamageConfigMeleeWeapon(core.MainHand, true, 101, 1, true),
		OutcomeApplier: core.OutcomeFuncMeleeSpecialCritOnly(rogue.critMultiplier(isMH, true)),

		OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			rogue.AddComboPoints(sim, 1, ability.ActionID)
		},
	}
	if !isMH {
		effect.ProcMask = core.ProcMaskMeleeOHSpecial
		effect.BaseDamage = core.BaseDamageConfigMeleeWeapon(core.OffHand, true, 101, 1+0.1*float64(rogue.Talents.DualWieldSpecialization), true)
	}

	effect.BaseDamage = core.WrapBaseDamageConfig(effect.BaseDamage, func(oldCalculator core.BaseDamageCalculator) core.BaseDamageCalculator {
		return func(sim *core.Simulation, spellEffect *core.SpellEffect, spell *core.Spell) float64 {
			normalDamage := oldCalculator(sim, spellEffect, spell)
			if rogue.DeadlyPoisonDot.IsActive() {
				return normalDamage * 1.5
			} else {
				return normalDamage
			}
		}
	})

	return rogue.RegisterSpell(core.SpellConfig{
		Template:     ability,
		ModifyCast:   core.ModifyCastAssignTarget,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(effect),
	})
}

func (rogue *Rogue) registerMutilateSpell(_ *core.Simulation) {
	mhHitSpell := rogue.newMutilateHitSpell(true)
	ohHitSpell := rogue.newMutilateHitSpell(false)

	refundAmount := MutilateEnergyCost * 0.8
	ability := rogue.newAbility(MutilateActionID, MutilateEnergyCost, SpellFlagBuilder, core.ProcMaskMeleeMHSpecial)

	rogue.Mutilate = rogue.RegisterSpell(core.SpellConfig{
		Template:   ability,
		ModifyCast: core.ModifyCastAssignTarget,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask:         core.ProcMaskMeleeMHSpecial,
			ThreatMultiplier: 1,
			OutcomeApplier:   core.OutcomeFuncMeleeSpecialHit(),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
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
			},
		}),
	})
}
