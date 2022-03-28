package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var AspectOfTheHawkActionID = core.ActionID{SpellID: 27044}
var AspectOfTheHawkAuraID = core.NewAuraID()
var ImprovedAspectOfTheHawkAuraID = core.NewAuraID()

var AspectOfTheViperActionID = core.ActionID{SpellID: 34074}
var AspectOfTheViperAuraID = core.NewAuraID()

func (hunter *Hunter) aspectOfTheHawkAura(sim *core.Simulation) core.Aura {
	const improvedHawkProcChance = 0.1
	improvedHawkBonus := 1 + 0.03*float64(hunter.Talents.ImprovedAspectOfTheHawk)
	impHawkAura := core.Aura{
		ID:       ImprovedAspectOfTheHawkAuraID,
		ActionID: core.ActionID{SpellID: 19556},
		Duration: time.Second * 12,
		OnGain: func(sim *core.Simulation) {
			hunter.PseudoStats.RangedSpeedMultiplier *= improvedHawkBonus
		},
		OnExpire: func(sim *core.Simulation) {
			hunter.PseudoStats.RangedSpeedMultiplier /= improvedHawkBonus
		},
	}

	factory := hunter.NewTemporaryStatsAuraFactory(AspectOfTheHawkAuraID, AspectOfTheHawkActionID, stats.Stats{stats.RangedAttackPower: 155}, core.NeverExpires)
	aura := factory(sim)
	aura.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		if !spellEffect.ProcMask.Matches(core.ProcMaskRangedAuto) {
			return
		}

		if improvedHawkBonus > 1 && sim.RandomFloat("Imp Aspect of the Hawk") < improvedHawkProcChance {
			hunter.ReplaceAura(sim, impHawkAura)
		}
	}
	return aura
}

func (hunter *Hunter) newAspectOfTheHawkTemplate(sim *core.Simulation) core.SimpleCast {
	aura := hunter.aspectOfTheHawkAura(sim)

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
				hunter.aspectOfTheViper = false
				hunter.RemoveAuraOnNextAdvance(sim, AspectOfTheViperAuraID)
				hunter.AddAuraOnNextAdvance(sim, aura)
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

func (hunter *Hunter) aspectOfTheViperAura() core.Aura {
	aura := core.Aura{
		ID:       AspectOfTheViperAuraID,
		ActionID: AspectOfTheViperActionID,
		Duration: core.NeverExpires,
		// Mana gain from viper is handled in rotation.go
	}
	return aura
}

func (hunter *Hunter) newAspectOfTheViperTemplate(sim *core.Simulation) core.SimpleCast {
	aura := hunter.aspectOfTheViperAura()

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
				hunter.aspectOfTheViper = true
				hunter.RemoveAuraOnNextAdvance(sim, AspectOfTheHawkAuraID)
				hunter.AddAuraOnNextAdvance(sim, aura)
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
	hunter.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		if hunter.Rotation.ViperStartManaPercent >= 1 {
			return hunter.aspectOfTheViperAura()
		} else {
			return hunter.aspectOfTheHawkAura(sim)
		}
	})
}
