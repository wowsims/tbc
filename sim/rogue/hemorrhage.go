package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var HemorrhageActionID = core.ActionID{SpellID: 26864}
var HemorrhageDebuffID = core.NewDebuffID()
var HemorrhageEnergyCost = 35.0

func (rogue *Rogue) newHemorrhageTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
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
	ability := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            HemorrhageActionID,
				Character:           &rogue.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				GCD:                 time.Second * 1,
				BaseCost: core.ResourceCost{
					Type:  stats.Energy,
					Value: HemorrhageEnergyCost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Energy,
					Value: HemorrhageEnergyCost,
				},
				CritMultiplier: rogue.critMultiplier(true, true),
				SpellExtras:    SpellFlagBuilder,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ProcMask:               core.ProcMaskMeleeMHSpecial,
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
				OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
					if spellEffect.Landed() {
						rogue.AddComboPoints(sim, 1, HemorrhageActionID)

						aura := hemoAura
						aura.Expires = sim.CurrentTime + hemoDuration
						spellEffect.Target.ReplaceAura(sim, aura)
					} else {
						rogue.AddEnergy(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
					}
				},
			},
			WeaponInput: core.WeaponDamageInput{
				Normalized:       true,
				DamageMultiplier: 1.1,
			},
		},
	}

	if ItemSetSlayers.CharacterHasSetBonus(&rogue.Character, 4) {
		ability.Effect.StaticDamageMultiplier *= 1.06
	}

	ability.Effect.WeaponInput.DamageMultiplier += 0.01 * float64(rogue.Talents.SinisterCalling)

	return core.NewSimpleSpellTemplate(ability)
}

func (rogue *Rogue) NewHemorrhage(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	hm := &rogue.hemorrhage
	rogue.hemorrhageTemplate.Apply(hm)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	hm.Effect.Target = target

	return hm
}
