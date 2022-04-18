package warrior

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

func (warrior *Warrior) registerBloodrageCD() {
	actionID := core.ActionID{SpellID: 2687}

	instantRage := 10.0 + 3*float64(warrior.Talents.ImprovedBloodrage)
	rageOverTime := 10.0

	brSpell := warrior.RegisterSpell(core.SpellConfig{
		ActionID: actionID,

		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    warrior.NewTimer(),
				Duration: time.Minute,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Target, _ *core.Spell) {
			// TODO: Rage over time should be done over time, not immediately.
			warrior.AddRage(sim, instantRage+rageOverTime, actionID)
		},
	})

	warrior.AddMajorCooldown(core.MajorCooldown{
		Spell: brSpell,
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return warrior.CurrentRage() < 70
		},
	})
}
