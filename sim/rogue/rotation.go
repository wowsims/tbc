package rogue

import (
	"time"

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
		rogue.tryUseDamageFinisher(sim, energy)
	} else if energy >= rogue.builderEnergyCost {
		rogue.newBuilder(sim, sim.GetPrimaryTarget()).Cast(sim)
	}
}

func (rogue *Rogue) tryUseDamageFinisher(sim *core.Simulation, energy float64) {
	if rogue.Rotation.UseRupture &&
		!rogue.rupture.IsInUse() &&
		sim.GetRemainingDuration() >= time.Second*16 &&
		(sim.GetNumTargets() == 1 || !rogue.HasAura(BladeFlurryAuraID)) {
		if energy >= RuptureEnergyCost {
			rogue.NewRupture(sim, sim.GetPrimaryTarget()).Cast(sim)
		}
		return
	}

	if energy >= rogue.eviscerateEnergyCost {
		rogue.NewEviscerate(sim, sim.GetPrimaryTarget()).Cast(sim)
	}
}
