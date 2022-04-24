package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
)

// Time between energy ticks.
const EnergyTickDuration = time.Millisecond * 2020

// Extra 0.2 because Blizzard
const EnergyPerTick = 20.2

// OnEnergyGain is called any time energy is increased.
type OnEnergyGain func(sim *Simulation)

type energyBar struct {
	character *Character

	maxEnergy     float64
	currentEnergy float64

	comboPoints int32

	onEnergyGain OnEnergyGain
	tickAction   *PendingAction

	// Multiplies energy regen from ticks.
	EnergyTickMultiplier float64
}

func (character *Character) EnableEnergyBar(maxEnergy float64, onEnergyGain OnEnergyGain) {
	character.energyBar = energyBar{
		character:    character,
		maxEnergy:    MaxFloat(100, maxEnergy),
		onEnergyGain: onEnergyGain,
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

func (eb *energyBar) addEnergyInternal(sim *Simulation, amount float64, actionID ActionID) {
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
func (eb *energyBar) AddEnergy(sim *Simulation, amount float64, actionID ActionID) {
	eb.addEnergyInternal(sim, amount, actionID)
	eb.onEnergyGain(sim)
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

// Gives an immediate partial energy tick and restarts the tick timer.
func (eb *energyBar) ResetEnergyTick(sim *Simulation) {
	timeSinceLastTick := sim.CurrentTime - (eb.NextEnergyTickAt() - EnergyTickDuration)
	partialTickAmount := (EnergyPerTick * eb.EnergyTickMultiplier) * (float64(timeSinceLastTick) / float64(EnergyTickDuration))

	eb.addEnergyInternal(sim, partialTickAmount, ActionID{OtherID: proto.OtherAction_OtherActionEnergyRegen})
	eb.character.TryUseCooldowns(sim)
	eb.onEnergyGain(sim)

	eb.newTickAction(sim, false)
}

func (eb *energyBar) SpendComboPoints(sim *Simulation, actionID ActionID) {
	if sim.Log != nil {
		eb.character.Log(sim, "Spent %d combo points from %s (%d --> %d).", eb.comboPoints, actionID, eb.comboPoints, 0)
	}
	eb.character.Metrics.AddResourceEvent(actionID, proto.ResourceType_ResourceTypeComboPoints, float64(-eb.comboPoints), float64(-eb.comboPoints))
	eb.comboPoints = 0
}

func (eb *energyBar) newTickAction(sim *Simulation, randomTickTime bool) {
	if eb.tickAction != nil {
		eb.tickAction.Cancel(sim)
	}

	nextTickDuration := EnergyTickDuration
	if randomTickTime {
		nextTickDuration = time.Duration(sim.RandomFloat("Energy Tick") * float64(EnergyTickDuration))
	}

	pa := &PendingAction{
		NextActionAt: sim.CurrentTime + nextTickDuration,
		Priority:     ActionPriorityRegen,
	}
	pa.OnAction = func(sim *Simulation) {
		eb.addEnergyInternal(sim, EnergyPerTick*eb.EnergyTickMultiplier, ActionID{OtherID: proto.OtherAction_OtherActionEnergyRegen})
		eb.character.TryUseCooldowns(sim)
		eb.onEnergyGain(sim)

		pa.NextActionAt = sim.CurrentTime + EnergyTickDuration
		sim.AddPendingAction(pa)
	}
	eb.tickAction = pa
	sim.AddPendingAction(pa)
}

func (eb *energyBar) reset(sim *Simulation) {
	if eb.character == nil {
		return
	}

	eb.currentEnergy = eb.maxEnergy
	eb.comboPoints = 0
	eb.EnergyTickMultiplier = 1
	eb.newTickAction(sim, true)
}
