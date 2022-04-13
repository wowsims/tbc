package warrior

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

var BloodrageCooldownID = core.NewCooldownID()
var BloodrageCooldown = time.Minute

func (warrior *Warrior) registerBloodrageCD() {
	actionID := core.ActionID{SpellID: 2687, CooldownID: BloodrageCooldownID}

	instantRage := 10.0 + 3*float64(warrior.Talents.ImprovedBloodrage)
	rageOverTime := 10.0

	brSpell := warrior.RegisterSpell(core.SpellConfig{
		ActionID: actionID,

		Cast: core.CastConfig{
			Cooldown: BloodrageCooldown,
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Target, _ *core.Spell) {
			// TODO: Rage over time should be done over time, not immediately.
			warrior.AddRage(sim, instantRage+rageOverTime, actionID)
		},
	})

	warrior.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: BloodrageCooldownID,
		Cooldown:   BloodrageCooldown,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return warrior.CurrentRage() < 70
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				brSpell.Cast(sim, nil)
			}
		},
	})
}
