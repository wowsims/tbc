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
	hasDeathmantle2pc := ItemSetDeathmantle.CharacterHasSetBonus(&rogue.Character, 2)

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
	// Envenom doesn't have partial resists
	ability.Effect.WeaponInput.CalculateDamage = func(attackPower float64, bonusWeaponDamage float64) float64 {
		comboPoints := rogue.ComboPoints()
		base := 180 * float64(comboPoints)
		if hasDeathmantle2pc {
			base += 40 * float64(comboPoints)
		}
		return base + (attackPower*0.03)*float64(comboPoints)
	}

	// cp. backstab
	ability.Effect.StaticDamageMultiplier += 0.04 * float64(rogue.Talents.VilePoisons)
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
