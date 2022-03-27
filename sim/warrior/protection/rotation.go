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
			war.NewShieldSlam(sim, target).Cast(sim)
		} else if war.CanRevenge(sim) {
			war.NewRevenge(sim, target).Cast(sim)
		} else if war.ShouldShout(sim) {
			war.CastShout(sim)
		} else if (war.Rotation.ThunderClap == proto.ProtectionWarrior_Rotation_ThunderClapOnCD || (war.Rotation.ThunderClap == proto.ProtectionWarrior_Rotation_ThunderClapMaintain && target.RemainingAuraDuration(sim, core.ThunderClapDebuffID) < time.Second*2)) && war.CanThunderClap(sim) {
			war.NewThunderClap(sim).Cast(sim)
		} else if (war.Rotation.DemoShout == proto.ProtectionWarrior_Rotation_DemoShoutFiller || (war.Rotation.DemoShout == proto.ProtectionWarrior_Rotation_DemoShoutMaintain && target.RemainingAuraDuration(sim, core.DemoralizingShoutDebuffID) < time.Second*2)) && war.CanDemoralizingShout(sim) {
			war.NewDemoralizingShout(sim).Cast(sim)
		} else if war.CanDevastate(sim) {
			war.NewDevastate(sim, target).Cast(sim)
		} else if war.CanSunderArmor(sim, target) {
			war.NewSunderArmor(sim, target).Cast(sim)
		}
	}

	if war.CurrentRage() >= float64(war.Rotation.HeroicStrikeThreshold) && war.CanHeroicStrike(sim) {
		war.QueueHeroicStrike(sim)
	}
}
