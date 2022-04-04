package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var AspectOfTheHawkActionID = core.ActionID{SpellID: 27044}
var AspectOfTheViperActionID = core.ActionID{SpellID: 34074}

func (hunter *Hunter) newAspectOfTheHawkTemplate(sim *core.Simulation) core.SimpleCast {
	var impHawkAura *core.Aura
	const improvedHawkProcChance = 0.1
	if hunter.Talents.ImprovedAspectOfTheHawk > 0 {
		improvedHawkBonus := 1 + 0.03*float64(hunter.Talents.ImprovedAspectOfTheHawk)
		impHawkAura = hunter.GetOrRegisterAura(&core.Aura{
			Label:    "Improved Aspect of the Hawk",
			ActionID: core.ActionID{SpellID: 19556},
			Duration: time.Second * 12,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.PseudoStats.RangedSpeedMultiplier *= improvedHawkBonus
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.PseudoStats.RangedSpeedMultiplier /= improvedHawkBonus
			},
		})
	}

	hunter.AspectOfTheHawkAura = hunter.NewTemporaryStatsAuraWrapped("Aspect of the Hawk", AspectOfTheHawkActionID, stats.Stats{stats.RangedAttackPower: 155}, core.NeverExpires, func(aura *core.Aura) {

		aura.Tag = "Aspect"
		aura.Priority = 1

		oldOnGain := aura.OnGain
		aura.OnGain = func(aura *core.Aura, sim *core.Simulation) {
			oldOnGain(aura, sim)
			hunter.currentAspect = aura
		}
		aura.OnSpellHit = func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.ProcMask.Matches(core.ProcMaskRangedAuto) {
				return
			}

			if impHawkAura != nil && sim.RandomFloat("Imp Aspect of the Hawk") < improvedHawkProcChance {
				impHawkAura.Activate(sim)
			}
		}
	})

	template := core.SimpleCast{
		Cast: core.Cast{
			ActionID:  AspectOfTheHawkActionID,
			Character: hunter.GetCharacter(),
			BaseCost: core.ResourceCost{
				Type:  stats.Mana,
				Value: 140,
			},
			Cost: core.ResourceCost{
				Type:  stats.Mana,
				Value: 140,
			},
			GCD:         core.GCDDefault,
			IgnoreHaste: true, // Hunter GCD is locked at 1.5s
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				hunter.AspectOfTheHawkAura.Activate(sim)
			},
		},
	}

	return template
}

func (hunter *Hunter) NewAspectOfTheHawk(sim *core.Simulation) core.SimpleCast {
	v := hunter.aspectOfTheHawkTemplate
	v.Init(sim)
	return v
}

func (hunter *Hunter) newAspectOfTheViperTemplate(sim *core.Simulation) core.SimpleCast {
	hunter.AspectOfTheViperAura = hunter.RegisterAura(&core.Aura{
		Label:    "Aspect of the Viper",
		Tag:      "Aspect",
		ActionID: AspectOfTheViperActionID,
		Duration: core.NeverExpires,
		Priority: 1,
		// Mana gain from viper is handled in rotation.go
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			hunter.currentAspect = aura
		},
	})

	template := core.SimpleCast{
		Cast: core.Cast{
			ActionID:  core.ActionID{SpellID: 34074},
			Character: hunter.GetCharacter(),
			BaseCost: core.ResourceCost{
				Type:  stats.Mana,
				Value: 40,
			},
			Cost: core.ResourceCost{
				Type:  stats.Mana,
				Value: 40,
			},
			GCD:         core.GCDDefault,
			IgnoreHaste: true, // Hunter GCD is locked at 1.5s
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				hunter.AspectOfTheViperAura.Activate(sim)
			},
		},
	}

	return template
}

func (hunter *Hunter) NewAspectOfTheViper(sim *core.Simulation) core.SimpleCast {
	v := hunter.aspectOfTheViperTemplate
	v.Init(sim)
	return v
}

func (hunter *Hunter) applyInitialAspect() {
	hunter.RegisterResetEffect(func(sim *core.Simulation) {
		if hunter.Rotation.ViperStartManaPercent >= 1 {
			hunter.AspectOfTheViperAura.Activate(sim)
		} else {
			hunter.AspectOfTheHawkAura.Activate(sim)
		}
	})
}
