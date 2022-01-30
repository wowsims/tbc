package warrior

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

func (warrior *Warrior) OnGCDReady(sim *core.Simulation) {
	warrior.tryUseGCD(sim)
}

func (warrior *Warrior) tryUseGCD(sim *core.Simulation) {
	warrior.WaitUntil(sim, sim.Duration+time.Second)
}
