package core

import (
	"time"
)

type PlayerAgent interface {
	BuffUp(*Simulation) // Any pre-start buffs to apply to the raid.

	// Returns the action this Agent would like to take next.
	ChooseAction(*Simulation) AgentAction

	// This will be invoked if the chosen action is actually executed, so the Agent can update its state.
	OnActionAccepted(*Simulation, AgentAction)

	// Returns this Agent to its initial state.
	Reset(newsim *Simulation)
}

// A single action that an Agent can take.
type AgentAction struct {
	// Exactly one of these should be set.
	Wait time.Duration // Duration to wait
	Cast *Cast
}

func NewWaitAction(duration time.Duration) AgentAction {
	return AgentAction{
		Wait: duration,
	}
}

// func NewCastAction(sim *Simulation, sp *Spell) AgentAction {
// 	return AgentAction{
// 		Cast: NewCast(sim, sp),
// 	}
// }
