package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

var HemorrhageActionID = core.ActionID{SpellID: 26864}
var HemorrhageEnergyCost = 35.0

func (rogue *Rogue) registerHemorrhageSpell(sim *core.Simulation) {
	hemoAura := sim.GetPrimaryTarget().GetOrRegisterAura(&core.Aura{
		Label:    "Hemorrhage",
		ActionID: HemorrhageActionID,
		Duration: time.Second * 15,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.PseudoStats.BonusPhysicalDamageTaken += 42
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.PseudoStats.BonusPhysicalDamageTaken -= 42
		},
		OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spell.SpellSchool != core.SpellSchoolPhysical {
				return
			}
			if !spellEffect.Landed() || spellEffect.Damage == 0 {
				return
			}

			aura.RemoveStack(sim)
		},
	})

	refundAmount := HemorrhageEnergyCost * 0.8

	ability := rogue.newAbility(HemorrhageActionID, HemorrhageEnergyCost, SpellFlagBuilder, core.ProcMaskMeleeMHSpecial)
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
		if spellEffect.Landed() {
			rogue.AddComboPoints(sim, 1, HemorrhageActionID)

			hemoAura.Activate(sim)
			hemoAura.SetStacks(sim, 10)
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
