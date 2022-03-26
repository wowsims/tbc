package warrior

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

var ShoutCost = 10.0

func (warrior *Warrior) makeCastShoutHelper(actionID core.ActionID) func(sim *core.Simulation) {
	return func(sim *core.Simulation) {
		if warrior.IsOnCD(core.GCDCooldownID, sim.CurrentTime) {
			panic("Shout requires GCD")
		}
		if warrior.CurrentRage() < ShoutCost {
			panic("Shout requires rage")
		}

		// Actual shout effects are handled in core/buffs.go
		warrior.SetCD(core.GCDCooldownID, sim.CurrentTime+core.GCDDefault)
		warrior.SpendRage(sim, ShoutCost, actionID)
		warrior.Metrics.AddInstantCast(actionID)
		warrior.shoutExpiresAt = sim.CurrentTime + warrior.shoutDuration
	}
}

func (warrior *Warrior) makeCastShout() func(sim *core.Simulation) {
	if warrior.Shout == proto.WarriorShout_WarriorShoutBattle {
		return warrior.makeCastShoutHelper(core.ActionID{SpellID: 2048})
	} else if warrior.Shout == proto.WarriorShout_WarriorShoutCommanding {
		return warrior.makeCastShoutHelper(core.ActionID{SpellID: 469})
	} else {
		return nil
	}
}

func (warrior *Warrior) ShouldShout(sim *core.Simulation) bool {
	return warrior.Shout != proto.WarriorShout_WarriorShoutNone && warrior.CurrentRage() >= ShoutCost && sim.CurrentTime+time.Second*3 > warrior.shoutExpiresAt
}
