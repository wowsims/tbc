package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

type WaitAction struct {
	agent Agent

	duration time.Duration
}

func (action *WaitAction) GetActionID() ActionID {
	return ActionID{
		OtherID: proto.OtherAction_OtherActionWait,
	}
}

func (action *WaitAction) GetName() string {
	return "Wait"
}

func (action *WaitAction) GetTag() int32 {
	return 0
}

func (action *WaitAction) GetAgent() Agent {
	return action.agent
}

func (action *WaitAction) GetDuration() time.Duration {
	return action.duration
}

func (action *WaitAction) GetManaCost() float64 {
	return 0
}

func (action *WaitAction) Act(sim *Simulation) {
	if sim.Log != nil {
		sim.Log("Doing nothing for %0.1f seconds.\n", action.GetDuration())
	}
	sim.metricsAggregator.addAction(action)
}

func NewWaitAction(sim *Simulation, agent Agent, duration time.Duration) *WaitAction {
	return &WaitAction{
		agent: agent,
		duration: duration,
	}
}

// TODO: Find a better home for TLC action.
const ItemIDTLC = 28785

type LightningCapacitorCast struct {
	agent Agent
}

func (lcc LightningCapacitorCast) GetActionID() ActionID {
	return ActionID{
		ItemID: ItemIDTLC,
	}
}

func (lcc LightningCapacitorCast) GetName() string {
	return "Lightning Capacitor"
}

func (lcc LightningCapacitorCast) GetTag() int32 {
	return 0
}

func (lcc LightningCapacitorCast) GetAgent() Agent {
	return lcc.agent
}

func (lcc LightningCapacitorCast) GetBaseManaCost() float64 {
	return 0
}

func (lcc LightningCapacitorCast) GetSpellSchool() stats.Stat {
	return stats.NatureSpellPower
}

func (lcc LightningCapacitorCast) GetCooldown() time.Duration {
	return 0
}

func (lcc LightningCapacitorCast) GetCastInput(sim *Simulation, cast *DirectCastAction) DirectCastInput {
	return DirectCastInput{
		CritMultiplier: 1.5,
	}
}

func (lcc LightningCapacitorCast) GetHitInputs(sim *Simulation, cast *DirectCastAction) []DirectCastDamageInput{
	hitInput := DirectCastDamageInput{
		MinBaseDamage: 694,
		MaxBaseDamage: 807,
		DamageMultiplier: 1,
	}

	return []DirectCastDamageInput{hitInput}
}

func (lcc LightningCapacitorCast) OnCastComplete(sim *Simulation, cast *DirectCastAction) {
}
func (lcc LightningCapacitorCast) OnSpellHit(sim *Simulation, cast *DirectCastAction, result *DirectCastDamageResult) {
}
func (lcc LightningCapacitorCast) OnSpellMiss(sim *Simulation, cast *DirectCastAction) {
}

func NewLightningCapacitorCast(sim *Simulation, agent Agent) *DirectCastAction {
	return NewDirectCastAction(sim, LightningCapacitorCast{agent: agent})
}
