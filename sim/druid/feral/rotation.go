package feral

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/druid"
)

func (cat *FeralDruid) OnGCDReady(sim *core.Simulation) {
	cat.doRotation(sim)
}

func (cat *FeralDruid) doRotation(sim *core.Simulation) {
	if !cat.Form.Matches(druid.Cat) {
		cat.Powershift.Cast(sim, nil)
		return
	}

	if cat.readyToShift {
		cat.innervateOrShift(sim)
		return
	}

	energy := cat.CurrentEnergy()
	comboPoints := cat.ComboPoints()

	if cat.shouldRip(sim, energy, comboPoints) {
		cat.Rip.Cast(sim, cat.CurrentTarget)
	} else if !cat.MangleAura.IsActive() && cat.CanMangleCat() {
		cat.Mangle.Cast(sim, cat.CurrentTarget)
	} else if cat.CanShred() {
		cat.Shred.Cast(sim, cat.CurrentTarget)
	} else if cat.PseudoStats.InFrontOfTarget && cat.CanMangleCat() {
		cat.Mangle.Cast(sim, cat.CurrentTarget)
	}
}

func (cat *FeralDruid) innervateOrShift(sim *core.Simulation) {
	cat.waitingForTick = false

	// If we have just now decided to shift, then we do not execute the shift immediately, but instead trigger an input delay for realism.
	if !cat.readyToShift {
		cat.readyToShift = true
		return
	}

	cat.readyToShift = false

	// Logic for Innervate and Haste Pot usage will go here. For now we just execute simple powershifts without bundling any caster form CDs.
	cat.Powershift.Cast(sim, nil)
}

func (cat *FeralDruid) shouldRip(sim *core.Simulation, energy float64, comboPoints int32) bool {
	canPrimaryRip := cat.Rotation.FinishingMove == proto.FeralDruid_Rotation_Rip && energy > cat.Rip.DefaultCast.Cost
	canWeaveRip := cat.Rotation.FinishingMove == proto.FeralDruid_Rotation_Bite && cat.Rotation.Ripweave && (energy >= 52) && !cat.PseudoStats.NoCost
	nearFightEnd := sim.GetRemainingDuration() < time.Second*10

	return (canPrimaryRip || canWeaveRip) && (comboPoints >= cat.Rotation.RipMinComboPoints) && !cat.RipDot.IsActive() && !nearFightEnd
}
