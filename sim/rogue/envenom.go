package rogue

import (
	"github.com/wowsims/tbc/sim/core"
)

var EnvenomActionID = core.ActionID{SpellID: 32684}

func (rogue *Rogue) newEnvenomTemplate(_ *core.Simulation) core.SimpleSpellTemplate {
	rogue.envenomEnergyCost = 35
	if ItemSetAssassination.CharacterHasSetBonus(&rogue.Character, 4) {
		rogue.envenomEnergyCost -= 10
	}

	refundAmount := 0.4 * float64(rogue.Talents.QuickRecovery)

	ability := rogue.newAbility(EnvenomActionID, rogue.envenomEnergyCost, SpellFlagFinisher|core.SpellExtrasIgnoreResists, core.ProcMaskMeleeMHSpecial)
	ability.SpellCast.SpellSchool = core.SpellSchoolNature
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		if spellEffect.Landed() {
			rogue.ApplyFinisher(sim, spellCast.ActionID)
		} else {
			if refundAmount > 0 {
				rogue.AddEnergy(sim, spellCast.Cost.Value*refundAmount, core.ActionID{SpellID: 31245})
			}
		}
	}

	basePerComboPoint := 180.0
	if ItemSetDeathmantle.CharacterHasSetBonus(&rogue.Character, 2) {
		basePerComboPoint += 40
	}
	ability.Effect.BaseDamage = core.BaseDamageConfig{
		Calculator: func(sim *core.Simulation, hitEffect *core.SpellHitEffect, spellCast *core.SpellCast) float64 {
			comboPoints := rogue.ComboPoints()
			base := basePerComboPoint * float64(comboPoints)
			return base + (hitEffect.MeleeAttackPower(spellCast)*0.03)*float64(comboPoints)
		},
		TargetSpellCoefficient: 0,
	}

	// cp. backstab
	ability.Effect.DamageMultiplier += 0.04 * float64(rogue.Talents.VilePoisons)
	if rogue.Talents.SurpriseAttacks {
		ability.SpellExtras |= core.SpellExtrasCannotBeDodged
	}

	return core.NewSimpleSpellTemplate(ability)
}

func (rogue *Rogue) NewEnvenom(_ *core.Simulation, target *core.Target) *core.SimpleSpell {
	comboPoints := rogue.ComboPoints()
	if comboPoints == 0 {
		panic("Envenom requires combo points!")
	}

	ev := &rogue.envenom
	rogue.envenomTemplate.Apply(ev)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	ev.ActionID.Tag = comboPoints
	ev.Effect.Target = target
	if rogue.deathmantle4pcProc {
		ev.Cost.Value = 0
		rogue.deathmantle4pcProc = false
	}

	return ev
}
