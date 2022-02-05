package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

var AspectOfTheHawkActionID = core.ActionID{SpellID: 27044}
var AspectOfTheHawkAuraID = core.NewAuraID()
var ImprovedAspectOfTheHawkAuraID = core.NewAuraID()

var AspectOfTheViperActionID = core.ActionID{SpellID: 34074}
var AspectOfTheViperAuraID = core.NewAuraID()

func (hunter *Hunter) aspectOfTheHawkAura() core.Aura {
	const improvedHawkProcChance = 0.1
	improvedHawkBonus := 1 + 0.03*float64(hunter.Talents.ImprovedAspectOfTheHawk)

	aura := core.Aura{
		ID:       AspectOfTheHawkAuraID,
		ActionID: AspectOfTheHawkActionID,
		Expires:  core.NeverExpires,
		OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
			hitEffect.BonusAttackPower += 155
		},
		OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
			if !hitEffect.IsWhiteHit {
				return
			}

			if improvedHawkBonus > 0 && sim.RandomFloat("Imp Aspect of the Hawk") < improvedHawkProcChance {
				hunter.PseudoStats.RangedSpeedMultiplier *= improvedHawkBonus
				hunter.AddAura(sim, core.Aura{
					ID:       ImprovedAspectOfTheHawkAuraID,
					ActionID: core.ActionID{SpellID: 19556},
					Expires:  sim.CurrentTime + time.Second*12,
					OnExpire: func(sim *core.Simulation) {
						hunter.PseudoStats.RangedSpeedMultiplier /= improvedHawkBonus
					},
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
			IgnoreHaste:  true, // Hunter GCD is locked at 1.5s
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				hunter.aspectOfTheViper = false
				hunter.RemoveAuraOnNextAdvance(sim, AspectOfTheHawkAuraID)
				hunter.AddAura(sim, aura)
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
		// Mana gain from viper is handled in rotation.go
	}

	template := core.SimpleCast{
		Cast: core.Cast{
			ActionID:     core.ActionID{SpellID: 34074},
			Character:    hunter.GetCharacter(),
			BaseManaCost: 40,
			ManaCost:     40,
			GCD:          core.GCDDefault,
			IgnoreHaste:  true, // Hunter GCD is locked at 1.5s
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				hunter.aspectOfTheViper = true
				hunter.RemoveAuraOnNextAdvance(sim, AspectOfTheHawkAuraID)
				hunter.AddAura(sim, aura)
			},
		},
	}

	return template
}

func (hunter *Hunter) NewAspectOfTheViper(sim *core.Simulation) core.SimpleCast {
	return hunter.aspectOfTheViperTemplate
}
