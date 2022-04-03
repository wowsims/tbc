package rogue

import (
	"github.com/wowsims/tbc/sim/core"
)

var EviscerateActionID = core.ActionID{SpellID: 26865}

func (rogue *Rogue) registerEviscerateSpell(sim *core.Simulation) {
	rogue.eviscerateEnergyCost = 35
	if ItemSetAssassination.CharacterHasSetBonus(&rogue.Character, 4) {
		rogue.eviscerateEnergyCost -= 10
	}

	refundAmount := 0.4 * float64(rogue.Talents.QuickRecovery)

	ability := rogue.newAbility(EviscerateActionID, rogue.eviscerateEnergyCost, SpellFlagFinisher, core.ProcMaskMeleeMHSpecial)
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
		if spellEffect.Landed() {
			rogue.ApplyFinisher(sim, spell.ActionID)
		} else {
			if refundAmount > 0 {
				rogue.AddEnergy(sim, spell.MostRecentCost*refundAmount, core.ActionID{SpellID: 31245})
			}
		}
	}

	basePerComboPoint := 185.0
	if ItemSetDeathmantle.CharacterHasSetBonus(&rogue.Character, 2) {
		basePerComboPoint += 40
	}
	ability.Effect.BaseDamage = core.BaseDamageConfig{
		Calculator: func(sim *core.Simulation, hitEffect *core.SpellEffect, spell *core.Spell) float64 {
			comboPoints := rogue.ComboPoints()
			base := 60.0 + basePerComboPoint*float64(comboPoints)
			roll := sim.RandomFloat("Eviscerate") * 120.0
			return base + roll + (hitEffect.MeleeAttackPower(spell.Character)*0.03)*float64(comboPoints) + hitEffect.BonusWeaponDamage(spell.Character)
		},
		TargetSpellCoefficient: 1,
	}

	// cp. backstab
	ability.Effect.DamageMultiplier += 0.05 * float64(rogue.Talents.ImprovedEviscerate)
	ability.Effect.DamageMultiplier += 0.02 * float64(rogue.Talents.Aggression)
	if rogue.Talents.SurpriseAttacks {
		ability.SpellExtras |= core.SpellExtrasCannotBeDodged
	}

	rogue.Eviscerate = rogue.RegisterSpell(core.SpellConfig{
		Template: ability,
		ModifyCast: func(sim *core.Simulation, target *core.Target, instance *core.SimpleSpell) {
			instance.Effect.Target = target
			instance.ActionID.Tag = rogue.ComboPoints()
			if rogue.deathmantle4pcProc {
				instance.Cost.Value = 0
				rogue.deathmantle4pcProc = false
			}
		},
	})
}
