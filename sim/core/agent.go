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

	// Called once before the first iteration, after all Agents and Targets are finalized.
	// Use this to do any precomputations that require access to Sim or Target fields.
	Init(sim *Simulation)

	// Returns this Agent to its initial state. Called before each Sim iteration
	// and once after the final iteration.
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
	// Only one of these should be set.
	SpellID int32
	ItemID  int32
	OtherID proto.OtherAction

	Tag int32

	CooldownID CooldownID // used only for tracking CDs internally
}

func (actionID ActionID) ToProto() *proto.ActionID {
	protoID := &proto.ActionID{
		Tag: actionID.Tag,
	}

	if actionID.SpellID != 0 {
		protoID.RawId = &proto.ActionID_SpellId{SpellId: actionID.SpellID}
	} else if actionID.ItemID != 0 {
		protoID.RawId = &proto.ActionID_ItemId{ItemId: actionID.ItemID}
	} else if actionID.OtherID != 0 {
		protoID.RawId = &proto.ActionID_OtherId{OtherId: actionID.OtherID}
	}

	return protoID
}

type AgentFactory func(Character, proto.Player) Agent

var agentFactories map[string]AgentFactory = make(map[string]AgentFactory)

func RegisterAgentFactory(emptyOptions interface{}, factory AgentFactory) {
	typeName := reflect.TypeOf(emptyOptions).Name()
	if _, ok := agentFactories[typeName]; ok {
		panic("Aleady registered agent factory: " + typeName)
	}
	//fmt.Printf("Registering type: %s", typeName)

	agentFactories[typeName] = factory
}

// Constructs a new Agent.
func NewAgent(player proto.Player) Agent {
	typeName := reflect.TypeOf(player.GetSpec()).Elem().Name()

	factory, ok := agentFactories[typeName]
	if !ok {
		panic("No agent factory for type: " + typeName)
	}

	character := NewCharacter(player)
	return factory(character, player)
}
