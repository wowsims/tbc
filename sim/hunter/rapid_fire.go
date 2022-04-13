package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var RapidFireCooldownID = core.NewCooldownID()

func (hunter *Hunter) registerRapidFireCD() {
	cooldown := time.Minute*5 - time.Minute*time.Duration(hunter.Talents.RapidKilling)
	actionID := core.ActionID{SpellID: 3045, CooldownID: RapidFireCooldownID}

	rfAura := hunter.RegisterAura(&core.Aura{
		Label:    "Rapid Fire",
		ActionID: actionID,
		Duration: time.Second * 15,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.PseudoStats.RangedSpeedMultiplier *= 1.4
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.PseudoStats.RangedSpeedMultiplier /= 1.4
		},
	})

	baseCost := 100.0
	rfSpell := hunter.RegisterSpell(core.SpellConfig{
		ActionID: actionID,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost,
			},
			Cooldown: cooldown,
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Target, _ *core.Spell) {
			rfAura.Activate(sim)
		},
	})

	hunter.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: RapidFireCooldownID,
		Cooldown:   cooldown,
		Type:       core.CooldownTypeDPS,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			// Make sure we don't reuse after a Readiness cast.
			return !rfAura.IsActive()
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				rfSpell.Cast(sim, nil)
			}
		},
	})
}
