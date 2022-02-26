package rogue

import (
	"github.com/wowsims/tbc/sim/core"
)

func (rogue *Rogue) OnGCDReady(sim *core.Simulation) {
	rogue.doRotation(sim)
}

func (rogue *Rogue) doRotation(sim *core.Simulation) {
	energy := rogue.CurrentEnergy()

	sndTimeRemaining := rogue.RemainingAuraDuration(sim, SliceAndDiceAuraID)
	if sndTimeRemaining <= 0 && rogue.comboPoints > 0 {
		if energy >= SliceAndDiceEnergyCost {
			rogue.castSliceAndDice()
		}
		return
	}

	if rogue.comboPoints == 5 {
		if energy >= rogue.eviscerateEnergyCost {
			rogue.NewEviscerate(sim, sim.GetPrimaryTarget()).Cast(sim)
		}
	} else if energy >= rogue.builderEnergyCost {
		rogue.newBuilder(sim, sim.GetPrimaryTarget()).Cast(sim)
	}
}
