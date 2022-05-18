package tank

import (
	"github.com/wowsims/tbc/sim/core"
)

func (bear *FeralTankDruid) OnGCDReady(sim *core.Simulation) {
	bear.doRotation(sim)
}

func (bear *FeralTankDruid) doRotation(sim *core.Simulation) {
}
