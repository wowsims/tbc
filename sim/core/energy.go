package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
)

// Time between energy ticks.
const tickDuration = time.Second * 2

// Extra 0.2 because Blizzard
const energyPerTick = 20.2

// OnEnergyTick is called each time an energy tick occurs, after the energy has been updated.
type OnEnergyTick func(sim *Simulation)

type energyBar struct {
	character *Character

	maxEnergy     float64
	currentEnergy float64

	onEnergyTick OnEnergyTick
	tickAction   *PendingAction
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

func (eb *energyBar) AddEnergy(sim *Simulation, amount float64, actionID ActionID) {
	if amount < 0 {
		panic("Trying to add negative energy!")
	}

	newEnergy := MinFloat(eb.currentEnergy+amount, eb.maxEnergy)

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

	if sim.Log != nil {
		eb.character.Log(sim, "Spent %0.3f energy from %s (%0.3f --> %0.3f).", amount, actionID, eb.currentEnergy, newEnergy)
	}

	eb.currentEnergy = newEnergy
}

func (eb *energyBar) reset(sim *Simulation) {
	if eb.character == nil {
		return
	}

	eb.currentEnergy = eb.maxEnergy

	pa := &PendingAction{
		Name:         "Energy Tick",
		Priority:     ActionPriorityRegen,
		NextActionAt: tickDuration,
	}
	pa.OnAction = func(sim *Simulation) {
		eb.AddEnergy(sim, energyPerTick, ActionID{OtherID: proto.OtherAction_OtherActionEnergyRegen})
		eb.onEnergyTick(sim)

		pa.NextActionAt = sim.CurrentTime + tickDuration
		sim.AddPendingAction(pa)
	}
	eb.tickAction = pa
	sim.AddPendingAction(pa)
}
