package core

import (
	"reflect"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
)

// Agent can be thought of as the 'Player', i.e. the thing controlling the Character.
// This is the interface implemented by each class/spec.
type Agent interface {
	// The Character controller by this Agent.
	GetCharacter() *Character

	// Updates the input Buffs to include raid-wide buffs provided by this Agent.
	AddRaidBuffs(*Buffs)
	// Updates the input Buffs to include party-wide buffs provided by this Agent.
	AddPartyBuffs(*Buffs)

	// Any pre-start buffs to apply to the raid/party/self
	BuffUp(*Simulation)

	// Returns this Agent to its initial state.
	Reset(newsim *Simulation)

	// Returns the action this Agent would like to take next.
	ChooseAction(*Simulation) AgentAction

	// This will be invoked right before the chosen action is actually executed, so the Agent can update its state.
	// Note that the action may be different from the action chosen by this agent
	OnActionAccepted(*Simulation, AgentAction)

	// OnSpellHit is used by class agents to customize casts before actually applying the damage.
	OnSpellHit(*Simulation, *Cast)
}

// A single action that an Agent can take.
type AgentAction struct {
	// Exactly one of these should be set.
	Wait time.Duration // Duration to wait
	Cast *Cast
}

type AgentFactory func(*Simulation, *Character, *proto.PlayerOptions) Agent

var agentFactories map[string]AgentFactory = make(map[string]AgentFactory)

func RegisterAgentFactory(emptyOptions interface{}, factory AgentFactory) {
	typeName := reflect.TypeOf(emptyOptions).Name()
	if _, ok := agentFactories[typeName]; ok {
		panic("Aleady registered agent factory: " + typeName)
	}
	//fmt.Printf("Registering type: %s", typeName)

	agentFactories[typeName] = factory
}

func NewAgent(sim *Simulation, character *Character, playerOptions *proto.PlayerOptions) Agent {
	typeName := reflect.TypeOf(playerOptions.GetSpec()).Elem().Name()

	factory, ok := agentFactories[typeName]
	if !ok {
		panic("No agent factory for type: " + typeName)
	}

	return factory(sim, character, playerOptions)
}
