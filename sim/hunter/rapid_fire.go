package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var RapidFireCooldownID = core.NewCooldownID()
var RapidFireAuraID = core.NewAuraID()

func (hunter *Hunter) registerRapidFireCD() {
	cooldown := time.Minute * 5
	actionID := core.ActionID{SpellID: 3045, CooldownID: RapidFireCooldownID}

	rfAura := core.Aura{
		ID:       RapidFireAuraID,
		ActionID: actionID,
		Duration: time.Second * 15,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			hunter.PseudoStats.RangedSpeedMultiplier *= 1.4
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			hunter.PseudoStats.RangedSpeedMultiplier /= 1.4
		},
	}

	template := core.SimpleCast{
		Cast: core.Cast{
			ActionID:  actionID,
			Character: hunter.GetCharacter(),
			Cooldown:  cooldown,
			BaseCost: core.ResourceCost{
				Type:  stats.Mana,
				Value: 100,
			},
			Cost: core.ResourceCost{
				Type:  stats.Mana,
				Value: 100,
			},
			OnCastComplete: func(aura *core.Aura, sim *core.Simulation, cast *core.Cast) {
				hunter.AddAura(sim, rfAura)
			},
		},
	}

	template.Cooldown -= time.Minute * time.Duration(hunter.Talents.RapidKilling)

	hunter.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: RapidFireCooldownID,
		Cooldown:   cooldown,
		Type:       core.CooldownTypeDPS,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			// Make sure we don't reuse after a Readiness cast.
			if character.HasAura(RapidFireAuraID) {
				return false
			}
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				cast := template
				cast.Init(sim)
				cast.StartCast(sim)
			}
		},
	})
}
