package core

import (
	"time"
)

type PendingAction struct {
	Name         string
	Priority     int
	OnAction     func(*Simulation)
	CleanUp      func(*Simulation)
	NextActionAt time.Duration

	cancelled bool
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

type ActionsQueue []*PendingAction

func (queue ActionsQueue) Len() int {
	return len(queue)
}
func (queue ActionsQueue) Less(i, j int) bool {
	return queue[i].NextActionAt < queue[j].NextActionAt ||
		(queue[i].NextActionAt == queue[j].NextActionAt && queue[i].Priority > queue[j].Priority)
}
func (queue ActionsQueue) Swap(i, j int) {
	queue[i], queue[j] = queue[j], queue[i]
}
func (queue *ActionsQueue) Push(newAction interface{}) {
	*queue = append(*queue, newAction.(*PendingAction))
}
func (queue *ActionsQueue) Pop() interface{} {
	old := *queue
	n := len(old)
	action := old[n-1]
	old[n-1] = nil
	*queue = old[0 : n-1]
	return action
}
