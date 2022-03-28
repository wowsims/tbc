package rogue

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

var SinisterStrikeActionID = core.ActionID{SpellID: 26862}

func (rogue *Rogue) SinisterStrikeEnergyCost() float64 {
	return []float64{45, 42, 40}[rogue.Talents.ImprovedSinisterStrike]
}

func (rogue *Rogue) newSinisterStrikeTemplate(_ *core.Simulation) core.SimpleSpellTemplate {
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
	ability.Effect.BaseDamage = core.BaseDamageFuncMeleeWeapon(core.MainHand, true, 98, 1, true)

	// cp. backstab
	ability.Effect.StaticDamageMultiplier += 0.02 * float64(rogue.Talents.Aggression)
	if rogue.Talents.SurpriseAttacks {
		ability.Effect.StaticDamageMultiplier += 0.1
	}
	if ItemSetSlayers.CharacterHasSetBonus(&rogue.Character, 4) {
		ability.Effect.StaticDamageMultiplier += 0.06
	}

	return core.NewSimpleSpellTemplate(ability)
}

func (rogue *Rogue) NewSinisterStrike(_ *core.Simulation, target *core.Target) *core.SimpleSpell {
	ss := &rogue.sinisterStrike
	rogue.sinisterStrikeTemplate.Apply(ss)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	ss.Effect.Target = target

	return ss
}
