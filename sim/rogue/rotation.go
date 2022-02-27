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
	comboPoints := rogue.ComboPoints()

	sndTimeRemaining := rogue.RemainingAuraDuration(sim, SliceAndDiceAuraID)
	if sndTimeRemaining <= 0 && comboPoints > 0 {
		if energy >= SliceAndDiceEnergyCost {
			rogue.castSliceAndDice()
		}
		return
	}

	target := sim.GetPrimaryTarget()
	if comboPoints == 5 {
		if rogue.Rotation.MaintainExposeArmor && !target.HasAura(core.ExposeArmorDebuffID) {
			if energy >= ExposeArmorEnergyCost {
				rogue.NewExposeArmor(sim, target).Cast(sim)
			}
		} else {
			rogue.tryUseDamageFinisher(sim, energy)
		}
	} else if energy >= rogue.builderEnergyCost {
		rogue.newBuilder(sim, target).Cast(sim)
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
