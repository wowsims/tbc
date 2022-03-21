package warrior

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var BloodthirstCooldownID = core.NewCooldownID()
var BloodthirstActionID = core.ActionID{SpellID: 30335, CooldownID: BloodthirstCooldownID}

const BloodthirstCost = 30.0

func (warrior *Warrior) newBloodthirstTemplate(_ *core.Simulation) core.SimpleSpellTemplate {
	ability := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            BloodthirstActionID,
				Character:           &warrior.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				GCD:                 core.GCDDefault,
				Cooldown:            time.Second * 6,
				IgnoreHaste:         true,
				BaseCost: core.ResourceCost{
					Type:  stats.Rage,
					Value: BloodthirstCost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Rage,
					Value: BloodthirstCost,
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
			},
			WeaponInput: core.WeaponDamageInput{
				CalculateDamage: func(attackPower float64, bonusWeaponDamage float64) float64 {
					return attackPower * 0.45
				},
			},
		},
	}

	refundAmount := BloodthirstCost * 0.8
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		if !spellEffect.Landed() {
			warrior.AddRage(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
		}
	}

	return core.NewSimpleSpellTemplate(ability)
}

func (warrior *Warrior) NewBloodthirst(_ *core.Simulation, target *core.Target) *core.SimpleSpell {
	bt := &warrior.bloodthirst
	warrior.bloodthirstTemplate.Apply(bt)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	bt.Effect.Target = target

	return bt
}

func (warrior *Warrior) CanBloodthirst(sim *core.Simulation) bool {
	return warrior.Talents.Bloodthirst && warrior.CurrentRage() >= BloodthirstCost && !warrior.IsOnCD(BloodthirstCooldownID, sim.CurrentTime)
}
