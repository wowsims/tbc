package rogue

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

var BackstabActionID = core.ActionID{SpellID: 26863}
var BackstabEnergyCost = 60.0

func (rogue *Rogue) newBackstabTemplate(_ *core.Simulation) core.SimpleSpellTemplate {
	refundAmount := BackstabEnergyCost * 0.8

	ability := rogue.newAbility(BackstabActionID, BackstabEnergyCost, SpellFlagBuilder, core.ProcMaskMeleeMHSpecial)
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		if spellEffect.Landed() {
			rogue.AddComboPoints(sim, 1, BackstabActionID)
		} else {
			rogue.AddEnergy(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
		}
	}
	ability.Effect.WeaponInput = core.WeaponDamageInput{
		Normalized:       true,
		FlatDamageBonus:  170,
		DamageMultiplier: 1.5,
	}
	ability.Effect.DirectInput = core.DirectDamageInput{
		SpellCoefficient: 1,
	}

	// all these use "Apply Aura: Modifies Damage/Healing Done", and stack additively (up to 142%)
	ability.Effect.StaticDamageMultiplier += 0.02 * float64(rogue.Talents.Aggression)
	if rogue.Talents.SurpriseAttacks {
		ability.Effect.StaticDamageMultiplier += 0.1
	}

	ability.Effect.StaticDamageMultiplier += 0.04 * float64(rogue.Talents.Opportunity)

	if ItemSetSlayers.CharacterHasSetBonus(&rogue.Character, 4) {
		ability.Effect.StaticDamageMultiplier += 0.06
	}

	ability.Effect.BonusCritRating += 10 * core.MeleeCritRatingPerCritChance * float64(rogue.Talents.PuncturingWounds)

	// SinisterCalling uses "Apply Aura: Modifies Effect Value", adding to the DamageMultiplier (up to 155%)
	ability.Effect.WeaponInput.DamageMultiplier += 0.01 * float64(rogue.Talents.SinisterCalling)

	return core.NewSimpleSpellTemplate(ability)
}

func (rogue *Rogue) NewBackstab(_ *core.Simulation, target *core.Target) *core.SimpleSpell {
	bs := &rogue.backstab
	rogue.backstabTemplate.Apply(bs)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	bs.Effect.Target = target

	return bs
}
