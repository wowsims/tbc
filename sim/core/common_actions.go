package core

// Ideally everything in here could go in sim/common, but these are needed by
// core so it would create a circular dependency.

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
)

type WaitAction struct {
	character *Character

	duration time.Duration
}

func (action WaitAction) GetActionID() ActionID {
	return ActionID{
		OtherID: proto.OtherAction_OtherActionWait,
	}
}

func (action WaitAction) GetName() string {
	return "Wait"
}

func (action WaitAction) GetTag() int32 {
	return 0
}

func (action WaitAction) GetCharacter() *Character {
	return action.character
}

func (action WaitAction) GetDuration() time.Duration {
	return action.duration
}

func (action WaitAction) GetManaCost() float64 {
	return 0
}

func (action WaitAction) Cast(sim *Simulation) bool {
	if sim.Log != nil {
		sim.Log("Doing nothing for %0.1f seconds.\n", action.GetDuration())
	}
	//sim.MetricsAggregator.AddAction(action)
	return true
}

func NewWaitAction(sim *Simulation, character *Character, duration time.Duration) WaitAction {
	return WaitAction{
		character: character,
		duration:  duration,
	}
}
