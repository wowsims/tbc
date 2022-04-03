package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func (rogue *Rogue) registerThistleTeaCD() {
	if rogue.Consumes.DefaultConjured != proto.Conjured_ConjuredRogueThistleTea {
		return
	}

	actionID := core.ActionID{ItemID: 7676, CooldownID: core.ConjuredCooldownID}

	const energyRegen = 40.0
	cooldown := time.Minute * 5

	template := core.SimpleCast{
		Cast: core.Cast{
			ActionID:  actionID,
			Character: rogue.GetCharacter(),
			Cooldown:  cooldown,
			OnCastComplete: func(aura *core.Aura, sim *core.Simulation, cast *core.Cast) {
				rogue.AddEnergy(sim, energyRegen, actionID)
			},
		},
	}

	rogue.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: core.ConjuredCooldownID,
		Cooldown:   cooldown,
		Type:       core.CooldownTypeDPS,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			// Make sure we have plenty of room so we dont energy cap right after using.
			if rogue.CurrentEnergy() > 40 {
				return false
			}
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				cast := template
				cast.Init(sim)
				cast.StartCast(sim)
			}
		},
	})
}
