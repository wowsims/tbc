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
	hemoAuraApplier := func(target *core.Target) core.Aura {
		charges := 10

		return core.Aura{
			ID:       HemorrhageDebuffID,
			ActionID: HemorrhageActionID,
			Duration: time.Second * 15,
			OnGain: func(sim *core.Simulation) {
				target.PseudoStats.BonusPhysicalDamageTaken += 42
			},
			OnExpire: func(sim *core.Simulation) {
				target.PseudoStats.BonusPhysicalDamageTaken -= 42
			},
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellCast.SpellSchool != core.SpellSchoolPhysical {
					return
				}
				if !spellEffect.Landed() || spellEffect.Damage == 0 {
					return
				}

				charges--
				if charges == 0 {
					target.RemoveAura(sim, HemorrhageDebuffID)
				}
			},
		}
	}

	refundAmount := HemorrhageEnergyCost * 0.8

	ability := rogue.newAbility(HemorrhageActionID, HemorrhageEnergyCost, SpellFlagBuilder, core.ProcMaskMeleeMHSpecial)
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		if spellEffect.Landed() {
			rogue.AddComboPoints(sim, 1, HemorrhageActionID)
			spellEffect.Target.ReplaceAura(sim, hemoAuraApplier(spellEffect.Target))
		} else {
			rogue.AddEnergy(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
		}
	}
	ability.Effect.WeaponInput = core.WeaponDamageInput{
		Normalized:       true,
		DamageMultiplier: 1.1,
	}
	ability.Effect.DirectInput = core.DirectDamageInput{
		SpellCoefficient: 1,
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
