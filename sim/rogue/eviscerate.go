package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var EviscerateActionID = core.ActionID{SpellID: 26865}

func (rogue *Rogue) makeEviscerateDamageCalcFn(sim *core.Simulation, numPoints int) core.MeleeDamageCalculator {
	base := 60.0 + 185.0*float64(numPoints)
	if ItemSetDeathmantle.CharacterHasSetBonus(&rogue.Character, 2) {
		base += 40.0 * float64(numPoints)
	}

	return func(attackPower float64, bonusWeaponDamage float64) float64 {
		roll := sim.RandomFloat("Eviscerate") * 120.0
		return base + roll + (attackPower*0.03)*float64(numPoints) + bonusWeaponDamage
	}
}

func (rogue *Rogue) newEviscerateTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	rogue.eviscerateEnergyCost = 35
	if ItemSetAssassination.CharacterHasSetBonus(&rogue.Character, 4) {
		rogue.eviscerateEnergyCost -= 10
	}

	rogue.eviscerateDamageCalcs = []core.MeleeDamageCalculator{
		nil,
		rogue.makeEviscerateDamageCalcFn(sim, 1),
		rogue.makeEviscerateDamageCalcFn(sim, 2),
		rogue.makeEviscerateDamageCalcFn(sim, 3),
		rogue.makeEviscerateDamageCalcFn(sim, 4),
		rogue.makeEviscerateDamageCalcFn(sim, 5),
	}

	finishingMoveEffects := rogue.makeFinishingMoveEffectApplier(sim)
	refundAmount := 0.4 * float64(rogue.Talents.QuickRecovery)

	ability := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            EviscerateActionID,
				Character:           &rogue.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				GCD:                 time.Second * 1,
				Cost: core.ResourceCost{
					Type:  stats.Energy,
					Value: rogue.eviscerateEnergyCost,
				},
				CritMultiplier: rogue.critMultiplier(true, false),
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ProcMask:               core.ProcMaskMeleeMHSpecial,
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
				OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
					if spellEffect.Landed() {
						numPoints := rogue.comboPoints
						rogue.SpendComboPoints(sim)
						finishingMoveEffects(sim, numPoints)
					} else {
						if refundAmount > 0 {
							rogue.AddEnergy(sim, spellCast.Cost.Value*refundAmount, core.ActionID{SpellID: 31245})
						}
					}
				},
			},
			WeaponInput: core.WeaponDamageInput{},
		},
	}

	ability.Effect.StaticDamageMultiplier *= 1 + 0.05*float64(rogue.Talents.ImprovedEviscerate)
	ability.Effect.StaticDamageMultiplier *= 1 + 0.02*float64(rogue.Talents.Aggression)
	if rogue.Talents.SurpriseAttacks {
		ability.Effect.CannotBeDodged = true
	}

	return core.NewSimpleSpellTemplate(ability)
}

func (rogue *Rogue) NewEviscerate(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	if rogue.comboPoints == 0 {
		panic("Eviscerate requires combo points!")
	}

	ev := &rogue.eviscerate
	rogue.eviscerateTemplate.Apply(ev)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	ev.ActionID.Tag = rogue.comboPoints
	ev.Effect.WeaponInput.CalculateDamage = rogue.eviscerateDamageCalcs[rogue.comboPoints]
	ev.Effect.Target = target
	if rogue.deathmantle4pcProc {
		ev.Cost.Value = 0
		rogue.deathmantle4pcProc = false
	}

	return ev
}
