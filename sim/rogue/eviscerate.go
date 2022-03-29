package rogue

import (
	"github.com/wowsims/tbc/sim/core"
)

var EviscerateActionID = core.ActionID{SpellID: 26865}

func (rogue *Rogue) newEviscerateTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	rogue.eviscerateEnergyCost = 35
	if ItemSetAssassination.CharacterHasSetBonus(&rogue.Character, 4) {
		rogue.eviscerateEnergyCost -= 10
	}

	refundAmount := 0.4 * float64(rogue.Talents.QuickRecovery)

	ability := rogue.newAbility(EviscerateActionID, rogue.eviscerateEnergyCost, SpellFlagFinisher, core.ProcMaskMeleeMHSpecial)
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		if spellEffect.Landed() {
			rogue.ApplyFinisher(sim, spellCast.ActionID)
		} else {
			if refundAmount > 0 {
				rogue.AddEnergy(sim, spellCast.Cost.Value*refundAmount, core.ActionID{SpellID: 31245})
			}
		}
	}

	basePerComboPoint := 185.0
	if ItemSetDeathmantle.CharacterHasSetBonus(&rogue.Character, 2) {
		basePerComboPoint += 40
	}
	ability.Effect.BaseDamage = core.BaseDamageConfig{
		Calculator: func(sim *core.Simulation, hitEffect *core.SpellEffect, spellCast *core.SpellCast) float64 {
			comboPoints := rogue.ComboPoints()
			base := 60.0 + basePerComboPoint*float64(comboPoints)
			roll := sim.RandomFloat("Eviscerate") * 120.0
			return base + roll + (hitEffect.MeleeAttackPower(spellCast)*0.03)*float64(comboPoints) + hitEffect.BonusWeaponDamage(spellCast)
		},
		TargetSpellCoefficient: 1,
	}

	// cp. backstab
	ability.Effect.DamageMultiplier += 0.05 * float64(rogue.Talents.ImprovedEviscerate)
	ability.Effect.DamageMultiplier += 0.02 * float64(rogue.Talents.Aggression)
	if rogue.Talents.SurpriseAttacks {
		ability.SpellExtras |= core.SpellExtrasCannotBeDodged
	}

	return core.NewSimpleSpellTemplate(ability)
}

func (rogue *Rogue) NewEviscerate(_ *core.Simulation, target *core.Target) *core.SimpleSpell {
	comboPoints := rogue.ComboPoints()
	if comboPoints == 0 {
		panic("Eviscerate requires combo points!")
	}

	ev := &rogue.eviscerate
	rogue.eviscerateTemplate.Apply(ev)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	ev.ActionID.Tag = comboPoints
	ev.Effect.Target = target
	if rogue.deathmantle4pcProc {
		ev.Cost.Value = 0
		rogue.deathmantle4pcProc = false
	}

	return ev
}
