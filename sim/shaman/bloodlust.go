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
		Tag:        int32(shaman.Index),
	}
}

func (shaman *Shaman) registerBloodlustCD() {
	if !shaman.SelfBuffs.Bloodlust {
		return
	}
	actionID := shaman.BloodlustActionID()

	var blAuras []*core.Aura
	var bloodlustMCD *core.MajorCooldown

	baseCost := 750.0
	bloodlustSpell := shaman.RegisterSpell(core.SpellConfig{
		ActionID: actionID,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.NewCast{
				Cost: baseCost * (1 - 0.02*float64(shaman.Talents.MentalQuickness)),
				GCD:  core.GCDDefault,
			},
			Cooldown: core.BloodlustCD,

			ModifyCast: func(_ *core.Simulation, _ *core.Spell, cast *core.NewCast) {
				// Needed because of the interaction between enhance GCD scheduler and other bloodlusts.
				if !bloodlustMCD.UsesGCD {
					cast.GCD = 0
				}
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Target, _ *core.Spell) {
			for _, blAura := range blAuras {
				blAura.Activate(sim)
			}

			// All MCDs that use the GCD and have a non-zero cast time must call this.
			shaman.UpdateMajorCooldowns()
		},
	})

	shaman.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: BloodlustCooldownID,
		Cooldown:   core.BloodlustCD,
		UsesGCD:    true,
		Priority:   core.CooldownPriorityBloodlust,
		Type:       core.CooldownTypeDPS,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			if character.CurrentMana() < bloodlustSpell.DefaultCast.Cost {
				return false
			}

			// Need to check if any party member has lust, not just self, because of
			// major CD ordering issues with the shared bloodlust.
			for _, partyMember := range character.Party.Players {
				if partyMember.GetCharacter().HasActiveAuraWithTag(core.BloodlustAuraTag) {
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

			blAuras = []*core.Aura{}
			for _, partyMember := range shaman.Party.Players {
				blAuras = append(blAuras, core.BloodlustAura(partyMember.GetCharacter(), actionID.Tag))
			}

			return func(sim *core.Simulation, character *core.Character) {
				bloodlustSpell.Cast(sim, nil)
			}
		},
	})
}
