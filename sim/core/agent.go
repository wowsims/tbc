package core

import (
	"reflect"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
)

// Agent can be thought of as the 'Player', i.e. the thing controlling the Character.
// This is the interface implemented by each class/spec.
type Agent interface {
	// The Character controlled by this Agent.
	GetCharacter() *Character

	// Updates the input Buffs to include raid-wide buffs provided by this Agent.
	AddRaidBuffs(raidBuffs *proto.RaidBuffs)
	// Updates the input Buffs to include party-wide buffs provided by this Agent.
	AddPartyBuffs(partyBuffs *proto.PartyBuffs)

	// Returns this Agent to its initial state. Called before each Sim iteration.
	Reset(sim *Simulation)

	// Allows the Agent to take whatever actions it wants to. This is called by
	// the main event loop. The return value determines when the main event loop
	// will call this again; it will call Act() at the time specified by the return
	// value.
	Act(sim *Simulation) time.Duration

	// Called after sim.CurrentTime is changed. Use this function to calculate
	// mana/energy regen, cooldown changes, etc.
	Advance(sim *Simulation, elapsedTime time.Duration)
}

type ActionID struct {
	SpellID    int32
	ItemID     int32
	CooldownID CooldownID // used only for tracking CDs internally
	OtherID    proto.OtherAction
	// Can add future id types here.
}

// A single action that an Agent can take.
type AgentAction interface {
	GetActionID() ActionID

	// For logging / debugging.
	GetName() string

	GetTag() int32

	// The Character performing this action.
	GetCharacter() *Character

	// How long this action takes to cast/channel/etc.
	// In other words, how long until another action should be chosen.
	GetDuration() time.Duration

	// TODO: Maybe change this to 'ResourceCost'
	// Amount of mana required to perform the action.
	GetManaCost() float64

	// Do the action. Returns whether the action was successful. An unsuccessful
	// action indicates that the prerequisites, like resource cost, were not met.
	Act(sim *Simulation) bool
}

type AgentFactory func(Character, proto.PlayerOptions, proto.IndividualSimRequest) Agent

var agentFactories map[string]AgentFactory = make(map[string]AgentFactory)

func RegisterAgentFactory(emptyOptions interface{}, factory AgentFactory) {
	typeName := reflect.TypeOf(emptyOptions).Name()
	if _, ok := agentFactories[typeName]; ok {
		panic("Aleady registered agent factory: " + typeName)
	}
	//fmt.Printf("Registering type: %s", typeName)

	agentFactories[typeName] = factory
}

// Constructs a new Agent. isr is only used for presims.
func NewAgent(player proto.Player, isr proto.IndividualSimRequest) Agent {
	if player.Options == nil {
		panic("No player options provided!")
	}

	typeName := reflect.TypeOf(player.Options.GetSpec()).Elem().Name()

	factory, ok := agentFactories[typeName]
	if !ok {
		panic("No agent factory for type: " + typeName)
	}

	character := NewCharacter(player)
	return factory(character, *player.Options, isr)
}
