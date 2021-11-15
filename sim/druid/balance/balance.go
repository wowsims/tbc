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
	}
	if balanceOptions.Options.InnervateTarget != nil {
		// if targetting myself for individual sim
		// TODO: what is my player idx for raid?
		selfBuffs.Innervate = balanceOptions.Options.InnervateTarget.TargetIndex == 0
	}

	druid := druid.NewDruid(character, selfBuffs, *balanceOptions.Talents)
	return &BalanceDruid{
		Druid:           druid,
		rotationOptions: *balanceOptions.Rotation,
	}
}

type BalanceDruid struct {
	druid.Druid

	rotationOptions proto.BalanceDruid_Rotation
}

func (moonkin *BalanceDruid) Reset(sim *core.Simulation) {
	moonkin.Druid.Reset(sim)
}

func (moonkin *BalanceDruid) Act(sim *core.Simulation) time.Duration {
	// Activate shared druid behaviors
	moonkin.TryInnervate(sim)

	target := sim.GetPrimaryTarget()
	if moonkin.rotationOptions.FaerieFire && !target.HasAura(druid.FaerieFireDebuffID) {
		target.AddAura(sim, core.Aura{
			ID:      druid.FaerieFireDebuffID,
			Name:    "Faerie Fire",
			Expires: sim.CurrentTime + time.Second*40,
			// TODO: implement increased melee hit
		})
		// TODO: turn faerie fire into a real cast so we get automatic GCD
		return sim.CurrentTime + time.Millisecond*1500
	} else if moonkin.rotationOptions.InsectSwarm && !moonkin.InsectSwarmSpell.DotInput.IsTicking(sim) {
		swarm := moonkin.NewInsectSwarm(sim, target)
		success := swarm.Act(sim)
		if !success {
			regenTime := moonkin.TimeUntilManaRegen(swarm.GetManaCost())
			return sim.CurrentTime + regenTime
		}
		return sim.CurrentTime + moonkin.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime)
	} else if moonkin.rotationOptions.Moonfire && !moonkin.MoonfireSpell.DotInput.IsTicking(sim) {
		moonfire := moonkin.NewMoonfire(sim, target)
		success := moonfire.Act(sim)
		if !success {
			regenTime := moonkin.TimeUntilManaRegen(moonfire.GetManaCost())
			return sim.CurrentTime + regenTime
		}
		return sim.CurrentTime + moonkin.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime)
	}

	var spell *core.SingleTargetDirectDamageSpell
	switch moonkin.rotationOptions.PrimarySpell {
	case proto.BalanceDruid_Rotation_Starfire:
		spell = moonkin.NewStarfire(sim, target, 8)
	case proto.BalanceDruid_Rotation_Starfire6:
		spell = moonkin.NewStarfire(sim, target, 6)
	case proto.BalanceDruid_Rotation_Wrath:
		spell = moonkin.NewWrath(sim, target)
	}

	actionSuccessful := spell.Act(sim)

	if !actionSuccessful {
		regenTime := moonkin.TimeUntilManaRegen(spell.GetManaCost())
		if sim.Log != nil {
			sim.Log("Not enough mana, regenerating for %s.\n", regenTime)
		}
		return sim.CurrentTime + regenTime
	}

	return sim.CurrentTime + core.MaxDuration(
		moonkin.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime),
		spell.CastTime)
}
