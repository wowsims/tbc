package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

const buildTimeBuffer = time.Second * 0

const (
	PlanNone = iota
	PlanOpener
	PlanExposeArmor
	PlanSliceASAP
	PlanFillBeforeEA
	PlanFillBeforeSND
	PlanMaximalSlice
)

func (rogue *Rogue) OnGCDReady(sim *core.Simulation) {
	rogue.doRotation(sim)
}

func (rogue *Rogue) doRotation(sim *core.Simulation) {
	switch rogue.plan {
	case PlanNone:
		rogue.doPlanNone(sim)
	case PlanSliceASAP:
		rogue.doPlanSliceASAP(sim)
	case PlanMaximalSlice:
		rogue.doPlanMaximalSlice(sim)
	case PlanExposeArmor:
		rogue.doPlanExposeArmor(sim)
	case PlanFillBeforeEA:
		rogue.doPlanFillBeforeEA(sim)
	case PlanFillBeforeSND:
		rogue.doPlanFillBeforeSND(sim)
	case PlanOpener:
		rogue.doPlanOpener(sim)
	}
}

// Opening rotation.
func (rogue *Rogue) doPlanOpener(sim *core.Simulation) {
	// Can add other opener logic here if we want.
	rogue.plan = PlanSliceASAP
	rogue.doPlanSliceASAP(sim)
}

// Cast SND as the next finisher, using no more builders than necessary.
func (rogue *Rogue) doPlanSliceASAP(sim *core.Simulation) {
	if rogue.doneSND {
		rogue.plan = PlanNone
		rogue.doPlanNone(sim)
		return
	}

	energy := rogue.CurrentEnergy()
	comboPoints := rogue.ComboPoints()
	target := sim.GetPrimaryTarget()
	sndTimeRemaining := rogue.RemainingAuraDuration(sim, SliceAndDiceAuraID)

	if comboPoints > 0 {
		if energy >= SliceAndDiceEnergyCost || rogue.deathmantle4pcProc {
			if rogue.canPoolEnergy(sim, energy) && sndTimeRemaining > time.Second*2 {
				return
			}
			rogue.castSliceAndDice()
			if rogue.disabledMCDs != nil {
				rogue.EnableAllCooldowns(rogue.disabledMCDs)
				rogue.disabledMCDs = nil
			}
			rogue.plan = PlanNone
		}
		return
	} else {
		if energy >= rogue.builderEnergyCost {
			rogue.CastBuilder(sim, target)
		}
	}
}

// Get the biggest Slice we can, but still leaving time for EA if necessary.
func (rogue *Rogue) doPlanMaximalSlice(sim *core.Simulation) {
	if rogue.doneSND {
		rogue.plan = PlanNone
		rogue.doPlanNone(sim)
		return
	}

	energy := rogue.CurrentEnergy()
	comboPoints := rogue.ComboPoints()
	target := sim.GetPrimaryTarget()
	sndTimeRemaining := rogue.RemainingAuraDuration(sim, SliceAndDiceAuraID)

	remainingSimDuration := sim.GetRemainingDuration()
	if rogue.sliceAndDiceDurations[comboPoints] >= remainingSimDuration {
		if energy >= SliceAndDiceEnergyCost || rogue.deathmantle4pcProc {
			if rogue.canPoolEnergy(sim, energy) && sndTimeRemaining > time.Second*2 {
				return
			}
			rogue.castSliceAndDice()
			rogue.plan = PlanNone
		}
		return
	}

	if sndTimeRemaining <= time.Second && comboPoints > 0 {
		if energy >= SliceAndDiceEnergyCost || rogue.deathmantle4pcProc {
			rogue.castSliceAndDice()
			rogue.plan = PlanNone
		}
		return
	}

	if rogue.MaintainingExpose(target) {
		eaTimeRemaining := target.RemainingAuraDuration(sim, core.ExposeArmorAuraID)
		if rogue.eaBuildTime+buildTimeBuffer > eaTimeRemaining {
			// Cast our slice and start prepping for EA.
			if comboPoints == 0 {
				rogue.plan = PlanExposeArmor
				rogue.doPlanExposeArmor(sim)
				return
			}
			if energy >= SliceAndDiceEnergyCost || rogue.deathmantle4pcProc {
				if rogue.canPoolEnergy(sim, energy) && sndTimeRemaining > time.Second*2 {
					return
				}
				rogue.castSliceAndDice()
				rogue.plan = PlanExposeArmor
				return
			}
		} else {
			if comboPoints == 5 {
				if energy >= SliceAndDiceEnergyCost || rogue.deathmantle4pcProc {
					if rogue.canPoolEnergy(sim, energy) && sndTimeRemaining > time.Second*2 {
						return
					}
					rogue.castSliceAndDice()
					rogue.plan = PlanFillBeforeEA
					return
				}
			} else if energy >= rogue.builderEnergyCost {
				rogue.CastBuilder(sim, target)
			}
		}
	} else {
		if comboPoints == 5 {
			if energy >= SliceAndDiceEnergyCost || rogue.deathmantle4pcProc {
				if rogue.canPoolEnergy(sim, energy) && sndTimeRemaining > time.Second*2 {
					return
				}
				rogue.castSliceAndDice()
				rogue.plan = PlanFillBeforeSND
				return
			}
		} else if energy >= rogue.builderEnergyCost {
			rogue.CastBuilder(sim, target)
		}
	}
}

// Build towards and cast a 5 pt Expose Armor.
func (rogue *Rogue) doPlanExposeArmor(sim *core.Simulation) {
	if rogue.doneEA {
		rogue.plan = PlanNone
		rogue.doPlanNone(sim)
		return
	}

	energy := rogue.CurrentEnergy()
	comboPoints := rogue.ComboPoints()
	target := sim.GetPrimaryTarget()

	if comboPoints == 5 {
		if energy >= ExposeArmorEnergyCost || rogue.deathmantle4pcProc {
			eaTimeRemaining := target.RemainingAuraDuration(sim, core.ExposeArmorAuraID)
			if rogue.canPoolEnergy(sim, energy) && eaTimeRemaining > time.Second*2 {
				return
			}
			rogue.ExposeArmor.Cast(sim, target)
			rogue.plan = PlanNone
		}
		return
	} else {
		if energy >= rogue.builderEnergyCost {
			rogue.CastBuilder(sim, target)
		}
	}
}

// Optional dps finisher followed by EA.
func (rogue *Rogue) doPlanFillBeforeEA(sim *core.Simulation) {
	energy := rogue.CurrentEnergy()
	comboPoints := rogue.ComboPoints()
	target := sim.GetPrimaryTarget()
	eaTimeRemaining := target.RemainingAuraDuration(sim, core.ExposeArmorAuraID)

	if rogue.eaBuildTime+buildTimeBuffer > eaTimeRemaining {
		// Cast our finisher and start prepping for EA.
		if comboPoints == 0 {
			rogue.plan = PlanExposeArmor
			rogue.doPlanExposeArmor(sim)
			return
		} else {
			if comboPoints < rogue.Rotation.MinComboPointsForDamageFinisher {
				rogue.plan = PlanExposeArmor
				return
			}
			if rogue.tryUseDamageFinisher(sim, energy, comboPoints) {
				rogue.plan = PlanExposeArmor
				return
			}
		}
	} else {
		if comboPoints == 5 {
			rogue.tryUseDamageFinisher(sim, energy, comboPoints)
		} else if energy >= rogue.builderEnergyCost {
			rogue.CastBuilder(sim, target)
		}
	}
}

// Optional dps finisher followed by SND.
func (rogue *Rogue) doPlanFillBeforeSND(sim *core.Simulation) {
	energy := rogue.CurrentEnergy()
	comboPoints := rogue.ComboPoints()
	target := sim.GetPrimaryTarget()
	sndTimeRemaining := rogue.RemainingAuraDuration(sim, SliceAndDiceAuraID)

	if !rogue.doneSND && rogue.eaBuildTime+buildTimeBuffer > sndTimeRemaining {
		// Cast our finisher and start prepping for SND.
		if comboPoints == 0 {
			rogue.plan = PlanMaximalSlice
			rogue.doPlanMaximalSlice(sim)
			return
		} else {
			if comboPoints < rogue.Rotation.MinComboPointsForDamageFinisher {
				rogue.plan = PlanMaximalSlice
				return
			}
			if rogue.tryUseDamageFinisher(sim, energy, comboPoints) {
				rogue.plan = PlanMaximalSlice
				return
			}
		}
	} else {
		if comboPoints == 5 || (comboPoints > 0 && sim.GetRemainingDuration() < time.Second*2) {
			rogue.tryUseDamageFinisher(sim, energy, comboPoints)
		} else if energy >= rogue.builderEnergyCost {
			rogue.CastBuilder(sim, target)
		}
	}
}

func (rogue *Rogue) doPlanNone(sim *core.Simulation) {
	energy := rogue.CurrentEnergy()
	if energy < 25 {
		// No ability costs < 25 energy so just wait.
		return
	}

	comboPoints := rogue.ComboPoints()
	target := sim.GetPrimaryTarget()

	if comboPoints == 0 {
		// No option other than using a builder.
		if energy >= rogue.builderEnergyCost {
			rogue.CastBuilder(sim, target)
		}
		return
	}

	sndTimeRemaining := rogue.RemainingAuraDuration(sim, SliceAndDiceAuraID)

	if !rogue.MaintainingExpose(target) {
		if rogue.doneSND || sndTimeRemaining > rogue.eaBuildTime+buildTimeBuffer {
			rogue.plan = PlanFillBeforeSND
			rogue.doPlanFillBeforeSND(sim)
		} else {
			rogue.plan = PlanMaximalSlice
			rogue.doPlanMaximalSlice(sim)
		}
		return
	}

	eaTimeRemaining := target.RemainingAuraDuration(sim, core.ExposeArmorAuraID)
	energyForEANext := rogue.builderEnergyCost*float64(5-comboPoints) + ExposeArmorEnergyCost
	eaNextBuildTime := core.MaxDuration(0, time.Duration(((energyForEANext-energy)/rogue.energyPerSecondAvg)*float64(time.Second)))
	spareTime := core.MaxDuration(0, eaTimeRemaining-eaNextBuildTime)
	if spareTime <= buildTimeBuffer {
		rogue.plan = PlanExposeArmor
		rogue.doPlanExposeArmor(sim)
		return
	}

	if sndTimeRemaining == 0 {
		rogue.plan = PlanSliceASAP
		rogue.doPlanSliceASAP(sim)
		return
	}

	if sndTimeRemaining > eaTimeRemaining {
		rogue.plan = PlanFillBeforeEA
		rogue.doPlanFillBeforeEA(sim)
		return
	}

	if rogue.doneSND {
		rogue.plan = PlanFillBeforeSND
		rogue.doPlanFillBeforeSND(sim)
		return
	}

	rogue.plan = PlanMaximalSlice
	rogue.doPlanMaximalSlice(sim)
}

func (rogue *Rogue) canPoolEnergy(sim *core.Simulation, energy float64) bool {
	return sim.GetRemainingDuration() >= time.Second*6 && energy <= 50 && (!rogue.HasAura(AdrenalineRushAuraID) || energy <= 30)
}

func (rogue *Rogue) tryUseDamageFinisher(sim *core.Simulation, energy float64, comboPoints int32) bool {
	if rogue.Rotation.UseRupture &&
		!rogue.Rupture.Instance.IsInUse() &&
		sim.GetRemainingDuration() >= rogue.RuptureDuration(comboPoints) &&
		(sim.GetNumTargets() == 1 || !rogue.HasAura(BladeFlurryAuraID)) {
		if energy >= RuptureEnergyCost || rogue.deathmantle4pcProc {
			rogue.Rupture.Cast(sim, sim.GetPrimaryTarget())
		}
		return true
	}

	if energy >= rogue.eviscerateEnergyCost || rogue.deathmantle4pcProc {
		rogue.Eviscerate.Cast(sim, sim.GetPrimaryTarget())
		return true
	}
	return false
}
