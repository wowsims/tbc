package tank

import (
	"github.com/wowsims/tbc/sim/core"
)

func (bear *FeralTankDruid) OnGCDReady(sim *core.Simulation) {
	bear.doRotation(sim)
}

func (bear *FeralTankDruid) OnAutoAttack(sim *core.Simulation, spell *core.Spell) {
	bear.tryQueueMaul(sim)
}

func (bear *FeralTankDruid) doRotation(sim *core.Simulation) {
	if bear.GCD.IsReady(sim) {
		if bear.CanMangle(sim) {
			bear.Mangle.Cast(sim, bear.CurrentTarget)
		} else if bear.CanLacerate(sim) {
			bear.Lacerate.Cast(sim, bear.CurrentTarget)
		} else if bear.shouldDemoRoar(sim) {
			bear.DemoralizingRoar.Cast(sim, bear.CurrentTarget)
		} else if bear.Rotation.MaintainFaerieFire && bear.ShouldFaerieFire(sim) {
			bear.FaerieFire.Cast(sim, bear.CurrentTarget)
		}
	}

	bear.tryQueueMaul(sim)
}

func (bear *FeralTankDruid) tryQueueMaul(sim *core.Simulation) {
	if bear.ShouldQueueMaul(sim) {
		bear.QueueMaul(sim)
	}
}

func (bear *FeralTankDruid) shouldDemoRoar(sim *core.Simulation) bool {
	return bear.ShouldDemoralizingRoar(sim, false, bear.Rotation.MaintainDemoralizingRoar)
}
