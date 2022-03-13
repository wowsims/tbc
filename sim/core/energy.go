package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
)

// Time between energy ticks.
const EnergyTickDuration = time.Millisecond * 2020

// Extra 0.2 because Blizzard
const EnergyPerTick = 20.2

// OnEnergyTick is called each time an energy tick occurs, after the energy has been updated.
type OnEnergyTick func(sim *Simulation)

type energyBar struct {
	character *Character

	maxEnergy     float64
	currentEnergy float64

	comboPoints int32

	onEnergyTick OnEnergyTick
	tickAction   *PendingAction

	// Multiplies energy regen from ticks.
	EnergyTickMultiplier float64

	// Adds to the next tick. Can also be negative.
	NextEnergyTickAdjustment float64
}

func (character *Character) EnableEnergyBar(maxEnergy float64, onEnergyTick OnEnergyTick) {
	character.energyBar = energyBar{
		character:    character,
		maxEnergy:    MaxFloat(100, maxEnergy),
		onEnergyTick: onEnergyTick,
	}
}

func (character *Character) HasEnergyBar() bool {
	return character.energyBar.character != nil
}

func (eb *energyBar) CurrentEnergy() float64 {
	return eb.currentEnergy
}

func (eb *energyBar) NextEnergyTickAt() time.Duration {
	return eb.tickAction.NextActionAt
}

func (eb *energyBar) AddEnergy(sim *Simulation, amount float64, actionID ActionID) {
	if amount < 0 {
		panic("Trying to add negative energy!")
	}

	newEnergy := MinFloat(eb.currentEnergy+amount, eb.maxEnergy)
	eb.character.Metrics.AddResourceEvent(actionID, proto.ResourceType_ResourceTypeEnergy, amount, newEnergy-eb.currentEnergy)

	if sim.Log != nil {
		eb.character.Log(sim, "Gained %0.3f energy from %s (%0.3f --> %0.3f).", amount, actionID, eb.currentEnergy, newEnergy)
	}

	eb.currentEnergy = newEnergy
}

func (eb *energyBar) SpendEnergy(sim *Simulation, amount float64, actionID ActionID) {
	if amount < 0 {
		panic("Trying to spend negative energy!")
	}

	newEnergy := eb.currentEnergy - amount
	eb.character.Metrics.AddResourceEvent(actionID, proto.ResourceType_ResourceTypeEnergy, -amount, -amount)

	if sim.Log != nil {
		eb.character.Log(sim, "Spent %0.3f energy from %s (%0.3f --> %0.3f).", amount, actionID, eb.currentEnergy, newEnergy)
	}

	eb.currentEnergy = newEnergy
}

func (eb *energyBar) ComboPoints() int32 {
	return eb.comboPoints
}

func (eb *energyBar) AddComboPoints(sim *Simulation, pointsToAdd int32, actionID ActionID) {
	newComboPoints := MinInt32(eb.comboPoints+pointsToAdd, 5)
	eb.character.Metrics.AddResourceEvent(actionID, proto.ResourceType_ResourceTypeComboPoints, float64(pointsToAdd), float64(newComboPoints-eb.comboPoints))

	if sim.Log != nil {
		eb.character.Log(sim, "Gained %d combo points from %s (%d --> %d)", pointsToAdd, actionID, eb.comboPoints, newComboPoints)
	}

	eb.comboPoints = newComboPoints
}

func (eb *energyBar) SpendComboPoints(sim *Simulation, actionID ActionID) {
	if sim.Log != nil {
		eb.character.Log(sim, "Spent %d combo points from %s (%d --> %d).", eb.comboPoints, actionID, eb.comboPoints, 0)
	}
	eb.character.Metrics.AddResourceEvent(actionID, proto.ResourceType_ResourceTypeComboPoints, float64(-eb.comboPoints), float64(-eb.comboPoints))
	eb.comboPoints = 0
}

func (eb *energyBar) reset(sim *Simulation) {
	if eb.character == nil {
		return
	}

	eb.currentEnergy = eb.maxEnergy
	eb.comboPoints = 0
	eb.EnergyTickMultiplier = 1
	eb.NextEnergyTickAdjustment = 0

	pa := &PendingAction{
		Name:         "Energy Tick",
		Priority:     ActionPriorityRegen,
		NextActionAt: EnergyTickDuration,
	}
	pa.OnAction = func(sim *Simulation) {
		eb.AddEnergy(sim, EnergyPerTick*eb.EnergyTickMultiplier+eb.NextEnergyTickAdjustment, ActionID{OtherID: proto.OtherAction_OtherActionEnergyRegen})
		eb.NextEnergyTickAdjustment = 0
		eb.character.TryUseCooldowns(sim)
		eb.onEnergyTick(sim)

		pa.NextActionAt = sim.CurrentTime + EnergyTickDuration
		sim.AddPendingAction(pa)
	}
	eb.tickAction = pa
	sim.AddPendingAction(pa)
}
