package protection

import (
	"github.com/wowsims/tbc/sim/core"
)

func (war *ProtectionWarrior) OnGCDReady(sim *core.Simulation) {
	war.doRotation(sim)
}

func (war *ProtectionWarrior) doRotation(sim *core.Simulation) {
	target := sim.GetPrimaryTarget()
	if war.IsOnCD(core.GCDCooldownID, sim.CurrentTime) {
		if war.CanHeroicStrike(sim) {
			war.QueueHeroicStrike(sim)
		}
		return
	}

	if war.CanShieldSlam(sim) {
		war.NewShieldSlam(sim, target).Cast(sim)
	} else if war.CanRevenge(sim) {
		war.NewRevenge(sim, target).Cast(sim)
	} else if war.CanDevastate(sim) {
		war.NewDevastate(sim, target).Cast(sim)
	} else if war.CanSunderArmor(sim, target) {
		war.NewSunderArmor(sim, target).Cast(sim)
	} else if war.CanHeroicStrike(sim) {
		war.QueueHeroicStrike(sim)
	}
}
