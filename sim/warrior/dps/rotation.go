package dps

import (
	"github.com/wowsims/tbc/sim/core"
)

func (war *DpsWarrior) OnGCDReady(sim *core.Simulation) {
	war.doRotation(sim)
}

func (war *DpsWarrior) doRotation(sim *core.Simulation) {
	if war.IsOnCD(core.GCDCooldownID, sim.CurrentTime) {
		if war.CanHeroicStrike(sim) {
			war.QueueHeroicStrike(sim)
		}
		return
	}

	if war.CanBloodthirst(sim) {
		war.Bloodthirst.Cast(sim, sim.GetPrimaryTarget())
	} else if war.CanWhirlwind(sim) {
		war.Whirlwind.Cast(sim, sim.GetPrimaryTarget())
	} else if war.CanHeroicStrike(sim) {
		war.QueueHeroicStrike(sim)
	}
}
