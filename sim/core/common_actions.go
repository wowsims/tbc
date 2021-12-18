package core

// Ideally everything in here could go in sim/common, but these are needed by
// core so it would create a circular dependency.

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
)

type WaitReason byte

const (
	WaitReasonNone     WaitReason = iota // unknown why we waited
	WaitReasonOOM                        // no mana to cast
	WaitReasonRotation                   // waiting on rotation
	WaitReasonOptimal                    // waiting because its more optimal than casting.
)

type WaitAction struct {
	character *Character

	duration time.Duration
	reason   WaitReason
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
	switch action.reason {
	case WaitReasonNone:
		if sim.Log != nil {
			action.character.Log(sim, "Idling for %s seconds, for no particular reason.", action.GetDuration())
		}
	case WaitReasonOOM:
		action.character.Metrics.MarkOOM(sim, action.character, action.GetDuration())
		if sim.Log != nil {
			action.character.Log(sim, "Not enough mana to cast, regenerating for %s.", action.GetDuration())
		}
	case WaitReasonRotation:
		if sim.Log != nil {
			action.character.Log(sim, "Waiting for %s due to rotation / CDs.", action.GetDuration())
		}
	case WaitReasonOptimal:
		if sim.Log != nil {
			action.character.Log(sim, "Waiting for %s because its more dps.", action.GetDuration())
		}
	}

	return true
}

func NewWaitAction(sim *Simulation, character *Character, duration time.Duration, reason WaitReason) WaitAction {
	return WaitAction{
		character: character,
		duration:  duration,
		reason:    reason,
	}
}
