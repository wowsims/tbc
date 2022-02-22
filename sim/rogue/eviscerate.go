package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var EviscerateActionID = core.ActionID{SpellID: 26865}

const EviscerateEnergyCost = 35.0

func makeEviscerateDamageCalcFn(sim *core.Simulation, numPoints int) core.MeleeDamageCalculator {
	return func(attackPower float64, bonusWeaponDamage float64) float64 {
		base := 60.0 + sim.RandomFloat("Eviscerate")*120.0
		return base + (185.0+attackPower*0.03)*float64(numPoints)
	}
}

func (rogue *Rogue) newEviscerateTemplate(sim *core.Simulation) core.MeleeAbilityTemplate {
	rogue.eviscerateDamageCalcs = []core.MeleeDamageCalculator{
		nil,
		makeEviscerateDamageCalcFn(sim, 1),
		makeEviscerateDamageCalcFn(sim, 2),
		makeEviscerateDamageCalcFn(sim, 3),
		makeEviscerateDamageCalcFn(sim, 4),
		makeEviscerateDamageCalcFn(sim, 5),
	}

	ability := core.ActiveMeleeAbility{
		MeleeAbility: core.MeleeAbility{
			ActionID:            EviscerateActionID,
			Character:           &rogue.Character,
			OutcomeRollCategory: core.OutcomeRollCategorySpecial,
			CritRollCategory:    core.CritRollCategoryPhysical,
			SpellSchool:         core.SpellSchoolPhysical,
			GCD:                 time.Second * 1,
			Cost: core.ResourceCost{
				Type:  stats.Energy,
				Value: EviscerateEnergyCost,
			},
			CritMultiplier: rogue.DefaultMeleeCritMultiplier(),
		},
		Effect: core.AbilityHitEffect{
			AbilityEffect: core.AbilityEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
			},
			WeaponInput: core.WeaponDamageInput{},
		},
		OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
			if hitEffect.Landed() {
				rogue.SpendComboPoints(sim)
			}
		},
	}

	return core.NewMeleeAbilityTemplate(ability)
}

func (rogue *Rogue) NewEviscerate(sim *core.Simulation, target *core.Target) *core.ActiveMeleeAbility {
	if rogue.comboPoints == 0 {
		panic("Eviscerate requires combo points!")
	}

	ev := &rogue.eviscerate
	rogue.eviscerateTemplate.Apply(ev)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	ev.ActionID.Tag = rogue.comboPoints
	ev.Effect.WeaponInput.CalculateDamage = rogue.eviscerateDamageCalcs[rogue.comboPoints]
	ev.Effect.Target = target

	return ev
}
