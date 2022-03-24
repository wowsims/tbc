package dps

import (
	"github.com/wowsims/tbc/sim/core"
)

func (war *DpsWarrior) OnGCDReady(sim *core.Simulation) {
	war.doRotation(sim)
}

func (war *DpsWarrior) doRotation(sim *core.Simulation) {
	if war.CanBloodthirst(sim) {
		war.NewBloodthirst(sim, sim.GetPrimaryTarget()).Cast(sim)
	} else if war.CanWhirlwind(sim) {
		war.NewWhirlwind(sim, sim.GetPrimaryTarget()).Cast(sim)
	} else if war.CanHeroicStrike(sim) {
		war.QueueHeroicStrike(sim)
	}
}
