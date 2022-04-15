package priest

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var PowerInfusionCooldownID = core.NewCooldownID()

func (priest *Priest) registerPowerInfusionCD() {
	if !priest.Talents.PowerInfusion {
		return
	}

	var powerInfusionTarget *core.Character
	var powerInfusionAura *core.Aura
	actionID := core.ActionID{SpellID: 10060, CooldownID: PowerInfusionCooldownID, Tag: int32(priest.Index)}
	baseCost := priest.BaseMana() * 0.16

	priest.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: PowerInfusionCooldownID,
		Cooldown:   core.PowerInfusionCD,
		Priority:   core.CooldownPriorityBloodlust,
		Type:       core.CooldownTypeMana,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			if powerInfusionTarget == nil {
				return false
			}
			if character.CurrentMana() < baseCost {
				return false
			}
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			// How can we determine the target will be able to continue casting
			// 	for the next 15s at 20% reduced mana cost? Arbitrary value until then.
			//if powerInfusionTarget.CurrentMana() < 3000 {
			//	return false
			//}
			if powerInfusionTarget.HasActiveAuraWithTag(core.BloodlustAuraTag) {
				return false
			}
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			powerInfusionTargetAgent := sim.Raid.GetPlayerFromRaidTarget(priest.SelfBuffs.PowerInfusionTarget)

			if powerInfusionTargetAgent != nil {
				powerInfusionTarget = powerInfusionTargetAgent.GetCharacter()
				powerInfusionAura = core.PowerInfusionAura(powerInfusionTarget, actionID.Tag)
			}

			piSpell := priest.GetOrRegisterSpell(core.SpellConfig{
				ActionID: actionID,

				ResourceType: stats.Mana,
				BaseCost:     baseCost,

				Cast: core.CastConfig{
					DefaultCast: core.Cast{
						Cost: baseCost,
					},
					Cooldown: core.PowerInfusionCD,
				},

				ApplyEffects: func(sim *core.Simulation, _ *core.Target, _ *core.Spell) {
					powerInfusionAura.Activate(sim)
				},
			})

			return func(sim *core.Simulation, character *core.Character) {
				piSpell.Cast(sim, nil)
			}
		},
	})
}
