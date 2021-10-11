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
	AddRaidBuffs(*proto.Buffs)
	// Updates the input Buffs to include party-wide buffs provided by this Agent.
	AddPartyBuffs(*proto.Buffs)

	// Any pre-start buffs to apply to the raid/party/self
	BuffUp(*Simulation)

	// Returns this Agent to its initial state. Called before each Sim iteration.
	Reset(newsim *Simulation)

	// Returns the action this Agent would like to take next.
	ChooseAction(*Simulation) AgentAction

	// This will be invoked right before the chosen action is actually executed, so the Agent can update its state.
	// Note that the action may be different from the action chosen by this agent
	OnActionAccepted(*Simulation, AgentAction)
}

type ActionID struct {
	SpellID    int32
	ItemID     int32
	CooldownID int32 // used only for tracking CDs internally
	OtherID    proto.OtherAction
	// Can add future id types here.
}

// A single action that an Agent can take.
type AgentAction interface {
	GetActionID() ActionID

	// For logging / debugging.
	GetName() string

	GetTag() int32

	// The Agent performing this action.
	GetAgent() Agent

	// How long this action takes to cast/channel/etc.
	// In other words, how long until another action should be chosen.
	GetDuration() time.Duration

	// TODO: Maybe change this to 'ResourceCost'
	// Amount of mana required to perform the action.
	GetManaCost() float64

	// Do the action.
	Act(sim *Simulation)
}

type AgentFactory func(*Simulation, Character, *proto.PlayerOptions) Agent

var agentFactories map[string]AgentFactory = make(map[string]AgentFactory)

func RegisterAgentFactory(emptyOptions interface{}, factory AgentFactory) {
	typeName := reflect.TypeOf(emptyOptions).Name()
	if _, ok := agentFactories[typeName]; ok {
		panic("Aleady registered agent factory: " + typeName)
	}
	//fmt.Printf("Registering type: %s", typeName)

	agentFactories[typeName] = factory
}

func NewAgent(sim *Simulation, character Character, playerOptions *proto.PlayerOptions) Agent {
	typeName := reflect.TypeOf(playerOptions.GetSpec()).Elem().Name()

	factory, ok := agentFactories[typeName]
	if !ok {
		panic("No agent factory for type: " + typeName)
	}

	return factory(sim, character, playerOptions)
}
