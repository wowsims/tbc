package rotations

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

func NewSimpleCasterRotation(impl SimpleCasterRotationImpl) *SimpleCasterRotation {
	return &SimpleCasterRotation{
		character: impl.GetAgent().GetCharacter(),
		impl: impl,
	}
}

// Manages the caster's rotation. Assumes the following:
// - Agent uses mana for their spells.
// - Agent uses only spells (no melee attacks / wands) unless OOM.
// - Agent is assumed to be OOM if there is not enough mana to cast the spell
//     returned by ChooseAction(). If that happens, ChooseOOMAction() will be
//     called to get a backup action.
//
// SimpleCasterRotation takes care of checking for OOM and using a backup
// action if necessary.
type SimpleCasterRotation struct {
	// Cache this for performance
	character *core.Character

	impl SimpleCasterRotationImpl

	// Amount of time left to wait before the Agent is no longer considered OOM.
	oomDuration time.Duration

	// The time at which Act() was called last.
	lastActTime time.Duration
}

// Interface to implement in order to use SimpleCasterRotation.
type SimpleCasterRotationImpl interface {
	// Returns the Agent that will be using this rotation.
	GetAgent() core.Agent

	// Returns the action this rotation would like to take next.
	ChooseAction(*core.Simulation) core.AgentAction

	// Returns the action this rotation would like to take next. Only called if the Agent is OOM.
	// Second parameter is the amount of time that will be spent in the OOM state.
	ChooseOOMAction(*core.Simulation, time.Duration) core.AgentAction

	// This will be invoked right before the chosen action is actually executed,
	// so the rotation can update its state if necessary. The given action will
	// be the result from either ChooseAction() or ChooseOOMAction().
	OnActionAccepted(*core.Simulation, core.AgentAction)

	// Returns this rotation to its initial state. Called before each Sim iteration.
	Reset(*core.Simulation)
}

func (rotation *SimpleCasterRotation) Reset(sim *core.Simulation) {
	rotation.impl.Reset(sim)
}

func (rotation *SimpleCasterRotation) Act(sim *core.Simulation) time.Duration {
	if rotation.oomDuration != 0 {
		rotation.oomDuration = core.MaxDuration(0, rotation.oomDuration - (sim.CurrentTime - rotation.lastActTime))
	}
	rotation.lastActTime = sim.CurrentTime

	var newAction core.AgentAction
	if rotation.oomDuration > 0 {
		newAction = rotation.impl.ChooseOOMAction(sim, rotation.oomDuration)
	} else {
		newAction = rotation.impl.ChooseAction(sim)

		actionSuccessful := newAction.Act(sim)
		if !actionSuccessful {
			// Cast failed, we assume that means Agent is OOM.
			regenTime := rotation.character.TimeUntilManaRegen(newAction.GetManaCost())
			rotation.oomDuration = regenTime

			newAction = rotation.impl.ChooseOOMAction(sim, rotation.oomDuration)
			newAction.Act(sim)
		}
	}

	rotation.impl.OnActionAccepted(sim, newAction)
	return sim.CurrentTime + core.MaxDuration(
			rotation.character.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime),
			newAction.GetDuration())
}
