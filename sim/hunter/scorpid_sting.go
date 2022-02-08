package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var ScorpidStingDebuffID = core.NewDebuffID()

func (hunter *Hunter) newScorpidStingTemplate(sim *core.Simulation) core.MeleeAbilityTemplate {
	manaCost := hunter.BaseMana() * 0.09
	actionID := core.ActionID{SpellID: 3043}

	ama := core.ActiveMeleeAbility{
		MeleeAbility: core.MeleeAbility{
			ActionID:    actionID,
			Character:   &hunter.Character,
			SpellSchool: stats.NatureSpellPower,
			GCD:         core.GCDDefault,
			Cost: core.ResourceCost{
				Type:  stats.Mana,
				Value: manaCost,
			},
		},
		Effect: core.AbilityHitEffect{
			WeaponInput: core.WeaponDamageInput{
				IsRanged: true,
			},
		},
		OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
			if !hitEffect.Landed() {
				return
			}

			hitEffect.Target.AddAura(sim, core.Aura{
				ID:       ScorpidStingDebuffID,
				ActionID: actionID,
				Expires:  sim.CurrentTime + time.Second*20,
			})
		},
	}

	ama.Cost.Value *= 1 - 0.02*float64(hunter.Talents.Efficiency)

	return core.NewMeleeAbilityTemplate(ama)
}

func (hunter *Hunter) NewScorpidSting(sim *core.Simulation, target *core.Target) *core.ActiveMeleeAbility {
	as := &hunter.scorpidSting
	hunter.scorpidStingTemplate.Apply(as)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	as.Effect.Target = target

	return as
}
