package shaman

import (
	"github.com/wowsims/tbc/sim/core"
)

var BloodlustCooldownID = core.NewCooldownID()

func (shaman *Shaman) registerBloodlustCD() {
	if !shaman.SelfBuffs.Bloodlust {
		return
	}
	actionID := core.ActionID{SpellID: 2825, Tag: int32(shaman.RaidIndex)}

	shaman.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: BloodlustCooldownID,
		Cooldown:   core.BloodlustCD,
		Priority:   core.CooldownPriorityBloodlust,
		Type:       core.CooldownTypeDPS,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			// Need to check if any party member has lust, not just self, because of
			// major CD ordering issues with the shared bloodlust.
			for _, partyMember := range character.Party.Players {
				if partyMember.GetCharacter().HasAura(core.BloodlustAuraID) {
					return false
				}
			}
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				for _, partyMember := range character.Party.Players {
					core.AddBloodlustAura(sim, partyMember.GetCharacter(), actionID.Tag)
				}
				character.SetCD(BloodlustCooldownID, sim.CurrentTime+core.BloodlustCD)
				character.Metrics.AddInstantCast(actionID)
			}
		},
	})
}
