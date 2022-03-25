package protection

import (
	"github.com/wowsims/tbc/sim/core"
)

func (war *ProtectionWarrior) OnGCDReady(sim *core.Simulation) {
	war.doRotation(sim)
}

func (war *ProtectionWarrior) doRotation(sim *core.Simulation) {
}
