package warrior

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var DevastateActionID = core.ActionID{SpellID: 30022}

func (warrior *Warrior) newDevastateTemplate(_ *core.Simulation) core.SimpleSpellTemplate {
	ability := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            DevastateActionID,
				Character:           &warrior.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				GCD:                 core.GCDDefault,
				IgnoreHaste:         true,
				BaseCost: core.ResourceCost{
					Type:  stats.Rage,
					Value: warrior.sunderArmorCost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Rage,
					Value: warrior.sunderArmorCost,
				},
				CritMultiplier: warrior.critMultiplier(true),
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ProcMask:               core.ProcMaskMeleeMHSpecial,
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
				FlatThreatBonus:        100,
			},
			WeaponInput: core.WeaponDamageInput{
				Normalized:       true,
				DamageMultiplier: 0.5,
			},
			DirectInput: core.DirectDamageInput{
				SpellCoefficient: 1,
			},
		},
	}

	refundAmount := warrior.sunderArmorCost * 0.8
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		if spellEffect.Landed() {
			target := spellEffect.Target
			if !target.HasAura(core.ExposeArmorDebuffID) {
				sa := &warrior.sunderArmor
				warrior.sunderArmorTemplate.Apply(sa)

				sa.Effect.Target = target
				sa.SpellExtras |= core.SpellExtrasAlwaysHits
				sa.Cost.Value = 0
				sa.BaseCost.Value = 0
				if target.NumStacks(core.SunderArmorDebuffID) == 5 {
					sa.Effect.SpellEffect.ThreatMultiplier = 0
				}

				sa.Cast(sim)
			}
		} else {
			warrior.AddRage(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
		}
	}

	return core.NewSimpleSpellTemplate(ability)
}

func (warrior *Warrior) NewDevastate(_ *core.Simulation, target *core.Target) *core.SimpleSpell {
	dv := &warrior.devastate
	warrior.devastateTemplate.Apply(dv)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	dv.Effect.Target = target

	// Bonus 35 damage / stack of sunder. Counts stacks AFTER cast but only if stacks > 0.
	saStacks := target.NumStacks(core.SunderArmorDebuffID)
	if saStacks != 0 {
		dv.Effect.WeaponInput.FlatDamageBonus = 35 * float64(core.MinInt32(saStacks+1, 5))
	}

	return dv
}

func (warrior *Warrior) CanDevastate(sim *core.Simulation) bool {
	return warrior.CurrentRage() >= warrior.sunderArmorCost
}
