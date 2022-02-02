package core

import (
	"time"
)

const (
	ActionPriorityLow = -1
	ActionPriorityGCD = 0

	// Higher than GCD because regen can cause GCD actions (if we were waiting
	// for mana).
	ActionPriorityRegen = 1

	// Autos can cause regen (JoW, rage, energy procs, etc) so they should be
	// higher prio so that we never go backwards in the priority order.
	ActionPriorityAuto = 2

	// DOTs need to be higher than anything else so that dots can properly expire before we take other actions.
	ActionPriorityDOT = 3
)

type PendingAction struct {
	Name         string
	Priority     int
	OnAction     func(*Simulation)
	CleanUp      func(*Simulation)
	NextActionAt time.Duration

	cancelled bool
	id        int
}

func (pa *PendingAction) Cancel(sim *Simulation) {
	if pa.cancelled {
		return
	}

	if pa.CleanUp != nil {
		pa.CleanUp(sim)
		pa.CleanUp = nil
	}

	pa.cancelled = true
}

type paPool struct {
	objs  []*PendingAction
	maxid int
}

func newPAPool() *paPool {
	objs := make([]*PendingAction, 64)
	for i := range objs {
		objs[i] = &PendingAction{
			id: i + 1,
		}
	}
	return &paPool{
		objs:  objs,
		maxid: len(objs) + 1,
	}
}

func (pap *paPool) Get() *PendingAction {
	if len(pap.objs) == 0 {
		// Allocate more
		newObjs := make([]*PendingAction, 128)
		copy(newObjs, pap.objs)
		pap.objs = newObjs
		for i := range pap.objs {
			if pap.objs[i] == nil {
				pap.objs[i] = &PendingAction{
					id: pap.maxid,
				}
				pap.maxid++
			}
		}
		// panic("for now dont do this")
	}

	pa := pap.objs[len(pap.objs)-1]
	pap.objs = pap.objs[:len(pap.objs)-1]

	return pa
}

func (pap *paPool) Put(pa *PendingAction) {
	pa.cancelled = false
	pa.CleanUp = nil
	pa.Name = ""
	pa.NextActionAt = 0
	pa.OnAction = nil
	pa.Priority = 0

	pap.objs = append(pap.objs, pa)
}
