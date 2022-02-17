package common

// Helper module for planning GCD-bound actions in advance.

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

// Returns whether the cast was successful.
type AbilityCaster func(sim *core.Simulation) bool

type ScheduledAbility struct {
	// When to cast this ability.
	CastAt time.Duration

	// How much GCD time will be used by this ability.
	Duration time.Duration

	// When the ability will be completed. Computed internally.
	doneAt time.Duration

	// How to cast this ability.
	TryCast AbilityCaster

	// Higher priority abilities get preference when resolving conflicts.
	Priority int
}

type GCDSchedule struct {
	// Scheduled abilities, sorted from soonest to latest CastAt time.
	schedule []ScheduledAbility
}

func (gs *GCDSchedule) Schedule(sim *core.Simulation, newAbility ScheduledAbility) {
	newAbility.doneAt = newAbility.CastAt + newAbility.Duration

	if len(gs.schedule) == 0 {
		gs.schedule = append(gs.schedule, newAbility)
		return
	}

	// Find the index at which this ability should be inserted, ignoring priority for now.
	var index = 0
	for _, scheduledAbility := range gs.schedule {
		if scheduledAblity.CastAt < newAbility.CastAt
			break
		}
		index++
	}

	// If the insert was at the end with no overlap, can just append.
	if index == len(gs.schedule) && gs.schedule[len(gs.schedule)-1].doneAt <= newAbility.CastAt {
		gs.schedule = append(gs.schedule, newAbility)
		return
	}

	conflictBefore := index > 0 && gs.schedule[index].doneAt > newAbility.CastAt
	conflictAfter := index < len(gs.schedule) && gs.schedule[index+1].CastAt < newAbility.doneAt
	if !conflictBefore && !conflictAfter {
		gs.schedule = append(gs.schedule, newAbility)
		copy(gs.schedule[index+1:], gs.schedule[index:])
		gs.schedule[index] = newAbility
		return
	}

	// If we're here, we have a conflict.
}

func (gs *GCDSchedule) DoNextAbility(sim *core.Simulation) {
}
