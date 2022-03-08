package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

var HemorrhageActionID = core.ActionID{SpellID: 26864}
var HemorrhageDebuffID = core.NewDebuffID()
var HemorrhageEnergyCost = 35.0

func (rogue *Rogue) newHemorrhageTemplate(_ *core.Simulation) core.SimpleSpellTemplate {
	hemoDuration := time.Second * 15
	hemoAura := core.Aura{
		ID:       HemorrhageDebuffID,
		ActionID: HemorrhageActionID,
		Stacks:   10,
	}
	hemoAura.OnBeforeSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellHitEffect) {
		if spellCast.SpellSchool != core.SpellSchoolPhysical {
			return
		}
		spellEffect.DirectInput.FlatDamageBonus += 42

		stacks := spellEffect.Target.NumStacks(HemorrhageDebuffID) - 1
		if stacks == 0 {
			spellEffect.Target.RemoveAura(sim, HemorrhageDebuffID)
		} else {
			aura := hemoAura
			aura.Stacks = stacks
			aura.Expires = sim.CurrentTime + hemoDuration
			spellEffect.Target.ReplaceAura(sim, aura)
		}
	}

	refundAmount := HemorrhageEnergyCost * 0.8

	ability := rogue.newAbility(HemorrhageActionID, HemorrhageEnergyCost, SpellFlagBuilder, core.ProcMaskMeleeMHSpecial)
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		if spellEffect.Landed() {
			rogue.AddComboPoints(sim, 1, HemorrhageActionID)

			aura := hemoAura
			aura.Expires = sim.CurrentTime + hemoDuration
			spellEffect.Target.ReplaceAura(sim, aura)
		} else {
			rogue.AddEnergy(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
		}
	}
	ability.Effect.WeaponInput = core.WeaponDamageInput{
		Normalized:       true,
		DamageMultiplier: 1.1,
	}

	// cp. backstab
	if ItemSetSlayers.CharacterHasSetBonus(&rogue.Character, 4) {
		ability.Effect.StaticDamageMultiplier += 0.06
	}

	ability.Effect.WeaponInput.DamageMultiplier += 0.01 * float64(rogue.Talents.SinisterCalling)

	return core.NewSimpleSpellTemplate(ability)
}

func (rogue *Rogue) NewHemorrhage(_ *core.Simulation, target *core.Target) *core.SimpleSpell {
	hm := &rogue.hemorrhage
	rogue.hemorrhageTemplate.Apply(hm)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	hm.Effect.Target = target

	return hm
}
