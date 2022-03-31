package rogue

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

var BackstabActionID = core.ActionID{SpellID: 26863}
var BackstabEnergyCost = 60.0

func (rogue *Rogue) registerBackstabSpell(_ *core.Simulation) {
	refundAmount := BackstabEnergyCost * 0.8

	ability := rogue.newAbility(BackstabActionID, BackstabEnergyCost, SpellFlagBuilder, core.ProcMaskMeleeMHSpecial)
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		if spellEffect.Landed() {
			rogue.AddComboPoints(sim, 1, BackstabActionID)
		} else {
			rogue.AddEnergy(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
		}
	}
	ability.Effect.BaseDamage = core.BaseDamageConfigMeleeWeapon(core.MainHand, true, 170, 1.5+0.01*float64(rogue.Talents.SinisterCalling), true)

	// all these use "Apply Aura: Modifies Damage/Healing Done", and stack additively (up to 142%)
	ability.Effect.DamageMultiplier += 0.02 * float64(rogue.Talents.Aggression)
	if rogue.Talents.SurpriseAttacks {
		ability.Effect.DamageMultiplier += 0.1
	}

	ability.Effect.DamageMultiplier += 0.04 * float64(rogue.Talents.Opportunity)

	if ItemSetSlayers.CharacterHasSetBonus(&rogue.Character, 4) {
		ability.Effect.DamageMultiplier += 0.06
	}

	ability.Effect.BonusCritRating += 10 * core.MeleeCritRatingPerCritChance * float64(rogue.Talents.PuncturingWounds)

	rogue.Backstab = rogue.RegisterSpell(core.SpellConfig{
		Template:   ability,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}
