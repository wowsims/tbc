package priest

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var PowerInfusionCooldownID = core.NewCooldownID()

func (priest *Priest) registerPowerInfusionCD() {

	if !!priest.Talents.PowerInfusion {
		return
	}

	actionID := core.ActionID{SpellID: 10060, CooldownID: PowerInfusionCooldownID, Tag: int32(priest.RaidIndex)}
	
	baseManaCost := priest.BaseMana() * 0.16

	powerInfusionCD := core.PowerInfusionCD
	
	var powerInfusionTarget *core.Character
	
	priest.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: PowerInfusionCooldownID,
		Cooldown:   powerInfusionCD,
		UsesGCD:    true,
		Type:       core.CooldownTypeMana,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			if powerInfusionTarget == nil {
				return false
			}
			// How can we determine the target will be able to continue casting 
			// 	for the next 15s at 20% reduced mana cost? Arbitrary value until then.
			if powerInfusionTarget.CurrentMana() < 3000 {
				return false
			}
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			powerInfusionTargetAgent := sim.Raid.GetPlayerFromRaidTarget(priest.SelfBuffs.PowerInfusionTarget)
			
			if powerInfusionTargetAgent != nil {
				powerInfusionTarget = powerInfusionTargetAgent.GetCharacter()
			}
			
			castTemplate := core.SimpleCast{
				Cast: core.Cast{
					ActionID:  actionID,
					Character: priest.GetCharacter(),
					BaseCost: core.ResourceCost{
						Type:  stats.Mana,
						Value: baseManaCost,
					},
					Cost: core.ResourceCost{
						Type:  stats.Mana,
						Value: baseManaCost,
					},
					GCD:      core.GCDDefault,
					Cooldown: powerInfusionCD,
					OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
						core.AddPowerInfusionAura(sim, powerInfusionTarget, actionID.Tag)
					},
				},
			}
			
			return func(sim *core.Simulation, character *core.Character) {
				cast := castTemplate
				cast.Init(sim)
				cast.StartCast(sim)
			}
		},
	})
}
