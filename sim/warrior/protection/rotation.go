package protection

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func (war *ProtectionWarrior) OnGCDReady(sim *core.Simulation) {
	war.doRotation(sim)
}

func (war *ProtectionWarrior) OnAutoAttack(sim *core.Simulation, spell *core.Spell) {
	war.tryQueueHsCleave(sim)
}

func (war *ProtectionWarrior) doRotation(sim *core.Simulation) {
	target := sim.GetPrimaryTarget()

	if war.GCD.IsReady(sim) {
		if war.CanShieldSlam(sim) {
			war.ShieldSlam.Cast(sim, target)
		} else if war.CanBloodthirst(sim) {
			war.Bloodthirst.Cast(sim, sim.GetPrimaryTarget())
		} else if war.CanMortalStrike(sim) {
			war.MortalStrike.Cast(sim, sim.GetPrimaryTarget())
		} else if war.CanRevenge(sim) {
			war.Revenge.Cast(sim, target)
		} else if war.ShouldShout(sim) {
			war.Shout.Cast(sim, nil)
		} else if (war.Rotation.ThunderClap == proto.ProtectionWarrior_Rotation_ThunderClapOnCD || (war.Rotation.ThunderClap == proto.ProtectionWarrior_Rotation_ThunderClapMaintain && war.ThunderClapAura.RemainingDuration(sim) < time.Second*2)) && war.CanThunderClap(sim) {
			war.ThunderClap.Cast(sim, target)
		} else if (war.Rotation.DemoShout == proto.ProtectionWarrior_Rotation_DemoShoutFiller || (war.Rotation.DemoShout == proto.ProtectionWarrior_Rotation_DemoShoutMaintain && war.DemoralizingShoutAura.RemainingDuration(sim) < time.Second*2)) && war.CanDemoralizingShout(sim) {
			war.DemoralizingShout.Cast(sim, target)
		} else if war.CanDevastate(sim) {
			war.Devastate.Cast(sim, target)
		} else if war.CanSunderArmor(sim) {
			war.SunderArmor.Cast(sim, target)
		}
	}

	war.tryQueueHsCleave(sim)
}

func (war *ProtectionWarrior) tryQueueHsCleave(sim *core.Simulation) {
	if war.CurrentRage() >= float64(war.Rotation.HsRageThreshold) {
		if war.Rotation.UseCleave {
			if war.CanCleave(sim) {
				war.QueueCleave(sim)
			}
		} else {
			if war.CanHeroicStrike(sim) {
				war.QueueHeroicStrike(sim)
			}
		}
	}
}
