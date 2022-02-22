package shaman

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var BloodlustCooldownID = core.NewCooldownID()

func (shaman *Shaman) BloodlustActionID() core.ActionID {
	return core.ActionID{
		SpellID:    2825,
		CooldownID: BloodlustCooldownID,
		Tag:        int32(shaman.RaidIndex),
	}
}

func (shaman *Shaman) registerBloodlustCD() {
	if !shaman.SelfBuffs.Bloodlust {
		return
	}
	actionID := shaman.BloodlustActionID()

	bloodlustTemplate := core.SimpleCast{
		Cast: core.Cast{
			ActionID:  actionID,
			Character: shaman.GetCharacter(),
			BaseCost: core.ResourceCost{
				Type:  stats.Mana,
				Value: 750,
			},
			Cost: core.ResourceCost{
				Type:  stats.Mana,
				Value: 750,
			},
			GCD:      core.GCDDefault,
			Cooldown: core.BloodlustCD,
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				for _, partyMember := range shaman.Party.Players {
					core.AddBloodlustAura(sim, partyMember.GetCharacter(), actionID.Tag)
				}

				// All MCDs that use the GCD and have a non-zero cast time must call this.
				shaman.UpdateMajorCooldowns()
			},
		},
	}

	bloodlustTemplate.Cost.Value -= bloodlustTemplate.BaseCost.Value * float64(shaman.Talents.MentalQuickness) * 0.02
	manaCost := bloodlustTemplate.Cost.Value
	var bloodlustMCD *core.MajorCooldown

	shaman.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: BloodlustCooldownID,
		Cooldown:   core.BloodlustCD,
		UsesGCD:    true,
		Priority:   core.CooldownPriorityBloodlust,
		Type:       core.CooldownTypeDPS,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			if character.CurrentMana() < manaCost {
				return false
			}

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
			bloodlustMCD = shaman.GetMajorCooldown(actionID)
			return func(sim *core.Simulation, character *core.Character) {
				cast := bloodlustTemplate

				// Needed because of the interaction between enhance GCD scheduler and other bloodlusts.
				if !bloodlustMCD.UsesGCD {
					cast.GCD = 0
				}

				cast.Init(sim)
				cast.StartCast(sim)
			}
		},
	})
}
