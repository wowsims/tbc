package rogue

import (
	"github.com/wowsims/tbc/sim/core"
)

var EnvenomActionID = core.ActionID{SpellID: 32684}

func (rogue *Rogue) registerEnvenomSpell(_ *core.Simulation) {
	rogue.envenomEnergyCost = 35
	if ItemSetAssassination.CharacterHasSetBonus(&rogue.Character, 4) {
		rogue.envenomEnergyCost -= 10
	}

	refundAmount := 0.4 * float64(rogue.Talents.QuickRecovery)

	ability := rogue.newAbility(EnvenomActionID, rogue.envenomEnergyCost, SpellFlagFinisher|core.SpellExtrasIgnoreResists, core.ProcMaskMeleeMHSpecial)
	ability.SpellCast.SpellSchool = core.SpellSchoolNature
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
		if spellEffect.Landed() {
			rogue.ApplyFinisher(sim, spell.ActionID)
		} else {
			if refundAmount > 0 {
				rogue.AddEnergy(sim, spell.MostRecentCost*refundAmount, core.ActionID{SpellID: 31245})
			}
		}
	}

	basePerComboPoint := 180.0
	if ItemSetDeathmantle.CharacterHasSetBonus(&rogue.Character, 2) {
		basePerComboPoint += 40
	}
	ability.Effect.BaseDamage = core.BaseDamageConfig{
		Calculator: func(sim *core.Simulation, hitEffect *core.SpellEffect, spell *core.Spell) float64 {
			comboPoints := rogue.ComboPoints()
			base := basePerComboPoint * float64(comboPoints)
			return base + (hitEffect.MeleeAttackPower(spell.Character)*0.03)*float64(comboPoints)
		},
		TargetSpellCoefficient: 0,
	}

	// cp. backstab
	ability.Effect.DamageMultiplier += 0.04 * float64(rogue.Talents.VilePoisons)
	if rogue.Talents.SurpriseAttacks {
		ability.SpellExtras |= core.SpellExtrasCannotBeDodged
	}

	rogue.Envenom = rogue.RegisterSpell(core.SpellConfig{
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
