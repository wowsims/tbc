package warrior

import (
	"github.com/wowsims/tbc/sim/core"
)

func (warrior *Warrior) OnGCDReady(sim *core.Simulation) {
	warrior.doRotation(sim)
}

func (warrior *Warrior) doRotation(sim *core.Simulation) {
	if warrior.CanBloodthirst(sim) {
		warrior.NewBloodthirst(sim, sim.GetPrimaryTarget()).Cast(sim)
	} else if warrior.CanWhirlwind(sim) {
		warrior.NewWhirlwind(sim, sim.GetPrimaryTarget()).Cast(sim)
	} else if warrior.CanHeroicStrike(sim) {
		warrior.QueueHeroicStrike(sim)
	}
}
