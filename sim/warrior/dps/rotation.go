package dps

import (
	"github.com/wowsims/tbc/sim/core"
)

func (war *DpsWarrior) OnGCDReady(sim *core.Simulation) {
	war.doRotation(sim)
}

func (war *DpsWarrior) doRotation(sim *core.Simulation) {
	if war.GCD.IsReady(sim) {
		if war.ShouldRampage(sim) {
			war.Rampage.Cast(sim, nil)
		} else if war.CanExecute(sim) {
			war.Execute.Cast(sim, sim.GetPrimaryTarget())
		} else if war.CanBloodthirst(sim) {
			war.Bloodthirst.Cast(sim, sim.GetPrimaryTarget())
		} else if war.CanMortalStrike(sim) {
			war.MortalStrike.Cast(sim, sim.GetPrimaryTarget())
		} else if war.CanWhirlwind(sim) {
			war.Whirlwind.Cast(sim, sim.GetPrimaryTarget())
		} else if war.ShouldBerserkerRage(sim) {
			war.BerserkerRage.Cast(sim, nil)
		}
	}

	if war.CanHeroicStrike(sim) {
		war.QueueHeroicStrike(sim)
	} else if war.CanCleave(sim) {
		war.QueueCleave(sim)
	}
}
