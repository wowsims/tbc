package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

var HemorrhageActionID = core.ActionID{SpellID: 26864}
var HemorrhageAuraID = core.NewAuraID()
var HemorrhageEnergyCost = 35.0

func (rogue *Rogue) registerHemorrhageSpell(_ *core.Simulation) {
	hemoAura := core.Aura{
		ID:       HemorrhageAuraID,
		ActionID: HemorrhageActionID,
		Duration: time.Second * 15,
		OnGain: func(sim *core.Simulation) {
			sim.GetPrimaryTarget().PseudoStats.BonusPhysicalDamageTaken += 42
		},
		OnExpire: func(sim *core.Simulation) {
			sim.GetPrimaryTarget().PseudoStats.BonusPhysicalDamageTaken -= 42
		},
	}
	hemoAura.OnSpellHit = func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
		if spell.SpellSchool != core.SpellSchoolPhysical {
			return
		}
		if !spellEffect.Landed() || spellEffect.Damage == 0 {
			return
		}

		stacks := spellEffect.Target.NumStacks(HemorrhageAuraID) - 1
		if stacks == 0 {
			spellEffect.Target.RemoveAura(sim, HemorrhageAuraID)
		} else {
			hemoAura.Stacks = stacks
			spellEffect.Target.ReplaceAura(sim, hemoAura)
		}
	}

	refundAmount := HemorrhageEnergyCost * 0.8

	ability := rogue.newAbility(HemorrhageActionID, HemorrhageEnergyCost, SpellFlagBuilder, core.ProcMaskMeleeMHSpecial)
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
		if spellEffect.Landed() {
			rogue.AddComboPoints(sim, 1, HemorrhageActionID)

			hemoAura.Stacks = 10
			spellEffect.Target.ReplaceAura(sim, hemoAura)
		} else {
			rogue.AddEnergy(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
		}
	}
	ability.Effect.BaseDamage = core.BaseDamageConfigMeleeWeapon(core.MainHand, true, 0, 1.1+0.01*float64(rogue.Talents.SinisterCalling), true)

	// cp. backstab
	if ItemSetSlayers.CharacterHasSetBonus(&rogue.Character, 4) {
		ability.Effect.DamageMultiplier += 0.06
	}

	rogue.Hemorrhage = rogue.RegisterSpell(core.SpellConfig{
		Template:   ability,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}
