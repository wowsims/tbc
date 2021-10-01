package core

import (
	"time"
)

type Agent interface {
	// Any pre-start buffs to apply to the raid/party/self
	BuffUp(*Simulation, *Party)

	// Returns this Agent to its initial state.
	Reset(newsim *Simulation)

	// Returns the action this Agent would like to take next.
	ChooseAction(*Simulation, *Party) AgentAction

	// This will be invoked right before the chosen action is actually executed, so the Agent can update its state.
	// Note that the action may be different from the action chosen by this agent
	OnActionAccepted(*Simulation, AgentAction)

	// OnSpellHit is used by class agents to customize casts before actually applying the damage.
	OnSpellHit(*Simulation, PlayerAgent, *Cast)
}

// A single action that an Agent can take.
type AgentAction struct {
	// Exactly one of these should be set.
	Wait time.Duration // Duration to wait
	Cast *Cast
}
