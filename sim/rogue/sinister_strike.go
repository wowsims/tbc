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

func (rogue *Rogue) newSinisterStrikeTemplate(sim *core.Simulation) core.MeleeAbilityTemplate {
	energyCost := rogue.SinisterStrikeEnergyCost()
	refundAmount := energyCost * 0.8
	ability := core.ActiveMeleeAbility{
		MeleeAbility: core.MeleeAbility{
			ActionID:    SinisterStrikeActionID,
			Character:   &rogue.Character,
			SpellSchool: stats.AttackPower,
			GCD:         time.Second * 1,
			Cost: core.ResourceCost{
				Type:  stats.Energy,
				Value: energyCost,
			},
			CritMultiplier: rogue.DefaultMeleeCritMultiplier(),
		},
		Effect: core.AbilityHitEffect{
			AbilityEffect: core.AbilityEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
			},
			WeaponInput: core.WeaponDamageInput{
				IsOH:             false,
				DamageMultiplier: 1,
				FlatDamageBonus:  98,
			},
		},
		OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
			if hitEffect.Landed() {
				rogue.AddComboPoint(sim)
			} else {
				rogue.AddEnergy(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
			}
		},
	}

	return core.NewMeleeAbilityTemplate(ability)
}

func (rogue *Rogue) NewSinisterStrike(sim *core.Simulation, target *core.Target) *core.ActiveMeleeAbility {
	ss := &rogue.sinisterStrike
	rogue.sinisterStrikeTemplate.Apply(ss)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	ss.Effect.Target = target

	return ss
}
