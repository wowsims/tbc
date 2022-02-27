package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var SinisterStrikeActionID = core.ActionID{SpellID: 26862}

func (rogue *Rogue) SinisterStrikeEnergyCost() float64 {
	return 45.0 - 2.5*float64(rogue.Talents.ImprovedSinisterStrike)
}

func (rogue *Rogue) newSinisterStrikeTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	energyCost := rogue.SinisterStrikeEnergyCost()
	refundAmount := energyCost * 0.8
	ability := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            SinisterStrikeActionID,
				Character:           &rogue.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				GCD:                 time.Second * 1,
				Cost: core.ResourceCost{
					Type:  stats.Energy,
					Value: energyCost,
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
						rogue.AddComboPoint(sim, SinisterStrikeActionID)
					} else {
						rogue.AddEnergy(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
					}
				},
			},
			WeaponInput: core.WeaponDamageInput{
				DamageMultiplier: 1,
				FlatDamageBonus:  98,
			},
		},
	}

	ability.Effect.StaticDamageMultiplier *= 1 + 0.02*float64(rogue.Talents.Aggression)
	if rogue.Talents.SurpriseAttacks {
		ability.Effect.StaticDamageMultiplier *= 1.1
	}
	if ItemSetSlayers.CharacterHasSetBonus(&rogue.Character, 4) {
		ability.Effect.StaticDamageMultiplier *= 1.06
	}

	return core.NewSimpleSpellTemplate(ability)
}

func (rogue *Rogue) NewSinisterStrike(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	ss := &rogue.sinisterStrike
	rogue.sinisterStrikeTemplate.Apply(ss)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	ss.Effect.Target = target

	return ss
}
