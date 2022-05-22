package feral

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func (cat *FeralDruid) OnGCDReady(sim *core.Simulation) {
	cat.doRotation(sim)
}

func (cat *FeralDruid) doRotation(sim *core.Simulation) {
	// TODO: If not in cat form, shift back.

	// TODO: If ready to shift, shift.

	energy := cat.CurrentEnergy()
	comboPoints := cat.ComboPoints()

	if cat.shouldRip(sim, energy, comboPoints) {
		cat.Rip.Cast(sim, cat.CurrentTarget)
	} else if (!cat.MangleAura.IsActive() || cat.PseudoStats.InFrontOfTarget) && cat.CanMangleCat() {
		cat.Mangle.Cast(sim, cat.CurrentTarget)
	} else {
		// TODO Use shred here instead.
		cat.Mangle.Cast(sim, cat.CurrentTarget)
	}
}

func (cat *FeralDruid) shouldRip(sim *core.Simulation, energy float64, comboPoints int32) bool {
	canPrimaryRip := cat.Rotation.FinishingMove == proto.FeralDruid_Rotation_Rip && energy > cat.Rip.DefaultCast.Cost
	canWeaveRip := cat.Rotation.FinishingMove == proto.FeralDruid_Rotation_Bite && cat.Rotation.Ripweave && (energy >= 52) && !cat.PseudoStats.NoCost
	nearFightEnd := sim.GetRemainingDuration() < time.Second*10

	return (canPrimaryRip || canWeaveRip) && (comboPoints >= cat.Rotation.RipCp) && !cat.RipDot.IsActive() && !nearFightEnd
}
