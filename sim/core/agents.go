package core

import (
	"time"
)

type PlayerAgent struct {
	Agent
	*Player
}

type Agent interface {
	// Any pre-start buffs to apply to the raid/party/self
	BuffUp(*Simulation, *Party)

	// Returns this Agent to its initial state.
	Reset(newsim *Simulation)

	// Returns the action this Agent would like to take next.
	ChooseAction(*Simulation, *Party) AgentAction

	// This will be invoked if the chosen action is actually executed, so the Agent can update its state.
	OnActionAccepted(*Simulation, AgentAction)

	// OnSpellHit is used by class agents to customize casts before actually applying the damage.
	OnSpellHit(*Simulation, *Player, *Cast)
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
