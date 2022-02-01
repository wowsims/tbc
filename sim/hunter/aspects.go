package hunter

import (
	"github.com/wowsims/tbc/sim/core"
)

var AspectOfTheHawkActionID = core.ActionID{SpellID: 27044}
var AspectOfTheHawkAuraID = core.NewAuraID()
var ImprovedAspectOfTheHawkAuraID = core.NewAuraID()

var AspectOfTheViperActionID = core.ActionID{SpellID: 34074}
var AspectOfTheViperAuraID = core.NewAuraID()

func (hunter *Hunter) aspectOfTheHawkAura() core.Aura {
	const improvedHawkProcChance = 0.1
	improvedHawkBonus := 0.03 * float64(hunter.Talents.ImprovedAspectOfTheHawk)

	aura := core.Aura{
		ID:       AspectOfTheHawkAuraID,
		ActionID: AspectOfTheHawkActionID,
		Expires:  core.NeverExpires,
		OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
			hitEffect.BonusAttackPower += 155
		},
		OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
			if improvedHawkBonus > 0 && sim.RandomFloat("Imp Aspect of the Hawk") < improvedHawkProcChance {
				hunter.AddAura(sim, core.Aura{
					ID:       ImprovedAspectOfTheHawkAuraID,
					ActionID: AspectOfTheHawkActionID,
				})
			}
		},
	}
	return aura
}

func (hunter *Hunter) applyAspectOfTheHawk() {
	aura := hunter.aspectOfTheHawkAura()
	hunter.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return aura
	})
}

func (hunter *Hunter) newAspectOfTheHawkTemplate(sim *core.Simulation) core.SimpleCast {
	aura := hunter.aspectOfTheHawkAura()

	template := core.SimpleCast{
		Cast: core.Cast{
			ActionID:     AspectOfTheHawkActionID,
			Character:    hunter.GetCharacter(),
			BaseManaCost: 140,
			ManaCost:     140,
			GCD:          core.GCDDefault,
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				hunter.RemoveAura(sim, AspectOfTheViperAuraID)
				hunter.AddAura(sim, aura)
				hunter.aspectOfTheViper = false
			},
		},
	}

	return template
}

func (hunter *Hunter) NewAspectOfTheHawk(sim *core.Simulation) core.SimpleCast {
	return hunter.aspectOfTheHawkTemplate
}

func (hunter *Hunter) newAspectOfTheViperTemplate(sim *core.Simulation) core.SimpleCast {
	aura := core.Aura{
		ID:       AspectOfTheViperAuraID,
		ActionID: AspectOfTheViperActionID,
		Expires:  core.NeverExpires,
	}

	template := core.SimpleCast{
		Cast: core.Cast{
			ActionID:     core.ActionID{SpellID: 34074},
			Character:    hunter.GetCharacter(),
			BaseManaCost: 40,
			ManaCost:     40,
			GCD:          core.GCDDefault,
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				hunter.RemoveAura(sim, AspectOfTheHawkAuraID)
				hunter.AddAura(sim, aura)
				hunter.aspectOfTheViper = true
			},
		},
	}

	// Mana gain from viper is handled in rotation.go

	return template
}

func (hunter *Hunter) NewAspectOfTheViper(sim *core.Simulation) core.SimpleCast {
	return hunter.aspectOfTheViperTemplate
}
