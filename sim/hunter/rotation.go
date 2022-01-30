package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

func (hunter *Hunter) OnGCDReady(sim *core.Simulation) {
	hunter.tryUseGCD(sim)
}

func (hunter *Hunter) OnManaTick(sim *core.Simulation) {
	if hunter.FinishedWaitingForManaAndGCDReady(sim) {
		hunter.tryUseGCD(sim)
	}
}

func (hunter *Hunter) tryUseGCD(sim *core.Simulation) {
	hunter.WaitUntil(sim, sim.Duration+time.Second)
}
