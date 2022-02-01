package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

var RapidFireCooldownID = core.NewCooldownID()
var RapidFireAuraID = core.NewAuraID()

func (hunter *Hunter) registerRapidFireCD() {
	cooldown := time.Minute * 5
	actionID := core.ActionID{SpellID: 3045, CooldownID: RapidFireCooldownID}

	template := core.SimpleCast{
		Cast: core.Cast{
			ActionID:     actionID,
			Character:    hunter.GetCharacter(),
			Cooldown:     cooldown,
			BaseManaCost: 100,
			ManaCost:     100,
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				hunter.PseudoStats.RangedSpeedMultiplier *= 1.4
				hunter.AddAura(sim, core.Aura{
					ID:       RapidFireAuraID,
					ActionID: actionID,
					Expires:  sim.CurrentTime + time.Second*15,
					OnExpire: func(sim *core.Simulation) {
						hunter.PseudoStats.RangedSpeedMultiplier /= 1.4
					},
				})
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
