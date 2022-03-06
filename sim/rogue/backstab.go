package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var BackstabActionID = core.ActionID{SpellID: 26863}
var BackstabEnergyCost = 60.0

func (rogue *Rogue) newBackstabTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	refundAmount := BackstabEnergyCost * 0.8
	ability := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            BackstabActionID,
				Character:           &rogue.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				GCD:                 time.Second * 1,
				BaseCost: core.ResourceCost{
					Type:  stats.Energy,
					Value: BackstabEnergyCost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Energy,
					Value: BackstabEnergyCost,
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
						rogue.AddComboPoints(sim, 1, BackstabActionID)
					} else {
						rogue.AddEnergy(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
					}
				},
			},
			WeaponInput: core.WeaponDamageInput{
				Normalized:       true,
				FlatDamageBonus:  170,
				DamageMultiplier: 1.5,
			},
		},
	}

	ability.Effect.StaticDamageMultiplier *= 1 + 0.02*float64(rogue.Talents.Aggression)
	if rogue.Talents.SurpriseAttacks {
		ability.Effect.StaticDamageMultiplier *= 1.1
	}
	if true { // TODO: This is only from behind.
		ability.Effect.StaticDamageMultiplier *= 1 + 0.04*float64(rogue.Talents.Opportunity)
	}
	if ItemSetSlayers.CharacterHasSetBonus(&rogue.Character, 4) {
		ability.Effect.StaticDamageMultiplier *= 1.06
	}

	ability.Effect.BonusCritRating += 10 * core.MeleeCritRatingPerCritChance * float64(rogue.Talents.PuncturingWounds)
	ability.Effect.WeaponInput.DamageMultiplier += 0.01 * float64(rogue.Talents.SinisterCalling)

	return core.NewSimpleSpellTemplate(ability)
}

func (rogue *Rogue) NewBackstab(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	bs := &rogue.backstab
	rogue.backstabTemplate.Apply(bs)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	bs.Effect.Target = target

	return bs
}
