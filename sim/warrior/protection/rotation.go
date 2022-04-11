package protection

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func (war *ProtectionWarrior) OnGCDReady(sim *core.Simulation) {
	war.doRotation(sim)
}

func (war *ProtectionWarrior) doRotation(sim *core.Simulation) {
	target := sim.GetPrimaryTarget()

	if !war.IsOnCD(core.GCDCooldownID, sim.CurrentTime) {
		if war.CanShieldSlam(sim) {
			war.ShieldSlam.Cast(sim, target)
		} else if war.CanRevenge(sim) {
			war.Revenge.Cast(sim, target)
		} else if war.ShouldShout(sim) {
			war.CastShout(sim)
		} else if (war.Rotation.ThunderClap == proto.ProtectionWarrior_Rotation_ThunderClapOnCD || (war.Rotation.ThunderClap == proto.ProtectionWarrior_Rotation_ThunderClapMaintain && war.ThunderClapAura.RemainingDuration(sim) < time.Second*2)) && war.CanThunderClap(sim) {
			war.ThunderClap.Cast(sim, target)
		} else if (war.Rotation.DemoShout == proto.ProtectionWarrior_Rotation_DemoShoutFiller || (war.Rotation.DemoShout == proto.ProtectionWarrior_Rotation_DemoShoutMaintain && war.DemoralizingShoutAura.RemainingDuration(sim) < time.Second*2)) && war.CanDemoralizingShout(sim) {
			war.DemoralizingShout.Cast(sim, target)
		} else if war.CanDevastate(sim) {
			war.Devastate.Cast(sim, target)
		} else if war.CanSunderArmor(sim, target) {
			war.SunderArmor.Cast(sim, target)
		}
	}

	if war.CurrentRage() >= float64(war.Rotation.HeroicStrikeThreshold) && war.CanHeroicStrike(sim) {
		war.QueueHeroicStrike(sim)
	}
}
