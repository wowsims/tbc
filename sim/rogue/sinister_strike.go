package rogue

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

var SinisterStrikeActionID = core.ActionID{SpellID: 26862}

func (rogue *Rogue) SinisterStrikeEnergyCost() float64 {
	return []float64{45, 42, 40}[rogue.Talents.ImprovedSinisterStrike]
}

func (rogue *Rogue) registerSinisterStrikeSpell(_ *core.Simulation) {
	energyCost := rogue.SinisterStrikeEnergyCost()
	refundAmount := energyCost * 0.8
	ability := rogue.newAbility(SinisterStrikeActionID, energyCost, SpellFlagBuilder, core.ProcMaskMeleeMHSpecial)
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		if spellEffect.Landed() {
			rogue.AddComboPoints(sim, 1, SinisterStrikeActionID)
		} else {
			rogue.AddEnergy(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
		}
	}
	ability.Effect.BaseDamage = core.BaseDamageConfigMeleeWeapon(core.MainHand, true, 98, 1, true)

	// cp. backstab
	ability.Effect.DamageMultiplier += 0.02 * float64(rogue.Talents.Aggression)
	if rogue.Talents.SurpriseAttacks {
		ability.Effect.DamageMultiplier += 0.1
	}
	if ItemSetSlayers.CharacterHasSetBonus(&rogue.Character, 4) {
		ability.Effect.DamageMultiplier += 0.06
	}

	rogue.SinisterStrike = rogue.RegisterSpell(core.SpellConfig{
		Template:   ability,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}
