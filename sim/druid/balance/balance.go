package balance

import (
	"log"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
	"github.com/wowsims/tbc/sim/druid"
)

func RegisterBalanceDruid() {
	core.RegisterAgentFactory(proto.PlayerOptions_BalanceDruid{}, func(character core.Character, options proto.PlayerOptions, isr proto.IndividualSimRequest) core.Agent {
		return NewBalanceDruid(character, options, isr)
	})
}

var InnervateCD = core.NewCooldownID()
var InnervateAuraID = core.NewAuraID()

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

	// TODO: implement innervate in main druid code.
	if moonkin.SelfBuffs.Innervate && moonkin.GetRemainingCD(InnervateCD, sim.CurrentTime) == 0 {
		if moonkin.GetStat(stats.Mana)/moonkin.GetInitialStat(stats.Mana) < 0.75 {
			oldRegen := moonkin.PsuedoStats.SpiritRegenRateCasting
			moonkin.PsuedoStats.SpiritRegenRateCasting = 1.0
			moonkin.PsuedoStats.ManaRegenMultiplier *= 3.0
			moonkin.AddAura(sim, core.Aura{
				ID:      InnervateAuraID,
				Name:    "Innervate",
				Expires: sim.CurrentTime + time.Second*20,
				OnExpire: func(sim *core.Simulation) {
					moonkin.PsuedoStats.SpiritRegenRateCasting = oldRegen
					moonkin.PsuedoStats.ManaRegenMultiplier /= 3.0
				},
			})
			moonkin.SetCD(InnervateCD, time.Minute*6)
		}
	}

	// if moonkin.rotationOptions.FaerieFire && !target.HasAura(druid.FaerieFireAuraID) {
	// 	// TODO: add faerie fire aura
	// 	return sim.CurrentTime + moonkin.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime)
	// } else if moonkin.rotationOptions.InsectSwarm && !moonkin.InsectSwarmSpell.DotInput.IsTicking(sim) {
	// 	// TODO: add insect swarm aura
	// 	return sim.CurrentTime + moonkin.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime)
	// } else
	if moonkin.rotationOptions.Moonfire && !moonkin.MoonfireSpell.DotInput.IsTicking(sim) {
		moonfire := moonkin.NewMoonfire(sim, sim.GetPrimaryTarget())
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
		spell = moonkin.NewStarfire(sim, sim.GetPrimaryTarget(), 8)
	case proto.BalanceDruid_Rotation_Starfire6:
		spell = moonkin.NewStarfire(sim, sim.GetPrimaryTarget(), 6)
	case proto.BalanceDruid_Rotation_Wrath:
		spell = moonkin.NewWrath(sim, sim.GetPrimaryTarget())
	}

	actionSuccessful := spell.Act(sim)

	if !actionSuccessful {
		regenTime := moonkin.TimeUntilManaRegen(spell.GetManaCost())
		if sim.Log != nil {
			sim.Log("Not enough mana, regenerating for %s.\n", regenTime)
		}
		if regenTime > time.Second*5 {
			log.Fatalf("spending more than 5 sec regen")
		}
		return sim.CurrentTime + regenTime
	}

	return sim.CurrentTime + spell.CastTime
}
