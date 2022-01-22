package shaman

import (
	"github.com/wowsims/tbc/sim/core"
)

var BloodlustCooldownID = core.NewCooldownID()

func (shaman *Shaman) registerBloodlustCD() {
	if !shaman.SelfBuffs.Bloodlust {
		return
	}
	actionID := core.ActionID{
		SpellID:    2825,
		CooldownID: BloodlustCooldownID,
		Tag:        int32(shaman.RaidIndex),
	}

	bloodlustTemplate := core.SimpleCast{
		Cast: core.Cast{
			ActionID:     actionID,
			Character:    shaman.GetCharacter(),
			BaseManaCost: 750,
			ManaCost:     750,
			Cooldown:     core.BloodlustCD,
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				for _, partyMember := range shaman.Party.Players {
					core.AddBloodlustAura(sim, partyMember.GetCharacter(), actionID.Tag)
				}
			},
		},
	}

	bloodlustTemplate.ManaCost -= bloodlustTemplate.BaseManaCost * float64(shaman.Talents.MentalQuickness) * 0.02
	manaCost := bloodlustTemplate.ManaCost

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
			return func(sim *core.Simulation, character *core.Character) {
				cast := bloodlustTemplate
				cast.Init(sim)
				cast.StartCast(sim)
			}
		},
	})
}
