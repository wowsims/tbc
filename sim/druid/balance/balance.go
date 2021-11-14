package balance

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/druid"
)

func RegisterBalanceDruid() {
	core.RegisterAgentFactory(proto.PlayerOptions_BalanceDruid{}, func(character core.Character, options proto.PlayerOptions, isr proto.IndividualSimRequest) core.Agent {
		return NewBalanceDruid(character, options, isr)
	})
}

func NewBalanceDruid(character core.Character, options proto.PlayerOptions, isr proto.IndividualSimRequest) *BalanceDruid {
	balanceOptions := options.GetBalanceDruid()

	selfBuffs := druid.SelfBuffs{
		Omen: balanceOptions.Options.OmenOfClarity,
		// if targetting myself for individual sim
		// TODO: what is my player idx for raid?
		Innervate: balanceOptions.Options.InnervateTarget.TargetIndex == 0,
	}

	druid := druid.NewDruid(character, selfBuffs, *balanceOptions.Talents)
	return &BalanceDruid{
		Druid:           druid,
		rotationOptions: balanceOptions.Rotation,
	}
}

type BalanceDruid struct {
	druid.Druid

	rotationOptions *proto.BalanceDruid_Rotation
}

func (moonkin *BalanceDruid) Reset(sim *core.Simulation) {
	moonkin.Druid.Reset(sim)
}

func (moonkin *BalanceDruid) Act(sim *core.Simulation) time.Duration {
	// TODO: handle all the buffs you keep up
	// target := sim.GetPrimaryTarget()

	// if moonkin.rotationOptions.FaerieFire && !target.HasAura(druid.FaerieFireAuraID) {
	// 	// TODO: add faerie fire aura
	// 	return sim.CurrentTime + moonkin.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime)
	// } else if moonkin.rotationOptions.InsectSwarm && !moonkin.InsectSwarmSpell.DotInput.IsTicking(sim) {
	// 	// TODO: add insect swarm aura
	// 	return sim.CurrentTime + moonkin.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime)
	// } else
	if moonkin.rotationOptions.Moonfire && !moonkin.MoonfireSpell.DotInput.IsTicking(sim) {
		moonfire := moonkin.NewMoonfire(sim, sim.GetPrimaryTarget())
		moonfire.Act(sim)
		return sim.CurrentTime + moonkin.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime)
	}

	if moonkin.rotationOptions.PrimarySpell == proto.BalanceDruid_Rotation_Starfire {
		starfire := moonkin.NewStarfire(sim, sim.GetPrimaryTarget(), 8)
		starfire.Act(sim)
		return sim.CurrentTime + starfire.CastTime
	} else if moonkin.rotationOptions.PrimarySpell == proto.BalanceDruid_Rotation_Starfire6 {
		starfire := moonkin.NewStarfire(sim, sim.GetPrimaryTarget(), 6)
		starfire.Act(sim)
		return sim.CurrentTime + starfire.CastTime
	}

	// default to normal druid stuff...
	return moonkin.Druid.Act(sim)
}
