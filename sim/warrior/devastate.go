package warrior

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var DevastateActionID = core.ActionID{SpellID: 30022}

func (warrior *Warrior) registerDevastateSpell(_ *core.Simulation) {
	ability := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    DevastateActionID,
				Character:   &warrior.Character,
				SpellSchool: core.SpellSchoolPhysical,
				GCD:         core.GCDDefault,
				IgnoreHaste: true,
				BaseCost: core.ResourceCost{
					Type:  stats.Rage,
					Value: warrior.sunderArmorCost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Rage,
					Value: warrior.sunderArmorCost,
				},
			},
			OutcomeRollCategory: core.OutcomeRollCategorySpecial,
			CritRollCategory:    core.CritRollCategoryPhysical,
			CritMultiplier:      warrior.critMultiplier(true),
		},
		Effect: core.SpellEffect{
			ProcMask:         core.ProcMaskMeleeMHSpecial,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			FlatThreatBonus:  100,
		},
	}

	normalBaseDamage := core.BaseDamageFuncMeleeWeapon(core.MainHand, true, 0, 0.5, true)
	ability.Effect.BaseDamage = core.BaseDamageConfig{
		Calculator: func(sim *core.Simulation, hitEffect *core.SpellEffect, spellCast *core.SpellCast) float64 {
			// Bonus 35 damage / stack of sunder. Counts stacks AFTER cast but only if stacks > 0.
			sunderBonus := 0.0
			saStacks := hitEffect.Target.NumStacks(core.SunderArmorDebuffID)
			if saStacks != 0 {
				sunderBonus = 35 * float64(core.MinInt32(saStacks+1, 5))
			}

			return normalBaseDamage(sim, hitEffect, spellCast) + sunderBonus
		},
		TargetSpellCoefficient: 1,
	}

	normalSunderModifier := core.ModifyCastAssignTarget
	devastateSunderModifier := func(sim *core.Simulation, target *core.Target, instance *core.SimpleSpell) {
		instance.Effect.Target = target
		instance.SpellExtras |= core.SpellExtrasAlwaysHits
		instance.Cost.Value = 0
		instance.BaseCost.Value = 0
		instance.GCD = 0
		if target.NumStacks(core.SunderArmorDebuffID) == 5 {
			instance.Effect.ThreatMultiplier = 0
		}
	}

	refundAmount := warrior.sunderArmorCost * 0.8
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		if spellEffect.Landed() {
			target := spellEffect.Target
			if !target.HasAura(core.ExposeArmorDebuffID) {
				warrior.SunderArmor.ModifyCast = devastateSunderModifier
				warrior.SunderArmor.Cast(sim, spellEffect.Target)
				warrior.SunderArmor.ModifyCast = normalSunderModifier
			}
		} else {
			warrior.AddRage(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
		}
	}

	warrior.Devastate = warrior.RegisterSpell(core.SpellConfig{
		Template:   ability,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}

func (warrior *Warrior) CanDevastate(sim *core.Simulation) bool {
	return warrior.CurrentRage() >= warrior.sunderArmorCost
}
