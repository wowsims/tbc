package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var MutilateActionID = core.ActionID{SpellID: 34413}
var MutilateEnergyCost = 60.0

func (rogue *Rogue) newMutilateTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	damageCast := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            MutilateActionID,
				Character:           &rogue.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				CritMultiplier:      rogue.critMultiplier(true, true),
				SpellExtras:         core.SpellExtrasAlwaysHits | SpellFlagBuilder,
			},
		},
	}
	baseEffect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			ProcMask:               core.ProcMaskMeleeMHSpecial,
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
			ThreatMultiplier:       1,
		},
		WeaponInput: core.WeaponDamageInput{
			Normalized:       true,
			FlatDamageBonus:  101,
			DamageMultiplier: 1,
		},
	}

	if true { // TODO: This is only from behind.
		baseEffect.StaticDamageMultiplier *= 1 + 0.04*float64(rogue.Talents.Opportunity)
	}
	if ItemSetSlayers.CharacterHasSetBonus(&rogue.Character, 4) {
		baseEffect.StaticDamageMultiplier *= 1.06
	}
	baseEffect.BonusCritRating += 5 * core.MeleeCritRatingPerCritChance * float64(rogue.Talents.PuncturingWounds)
	damageCast.Effects = []core.SpellHitEffect{
		baseEffect,
		baseEffect,
	}
	damageCast.Effects[1].ProcMask = core.ProcMaskMeleeOHSpecial

	rogue.mutilateDamageTemplate = core.NewSimpleSpellTemplate(damageCast)

	refundAmount := MutilateEnergyCost * 0.8
	ability := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            MutilateActionID,
				Character:           &rogue.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryNone,
				SpellSchool:         core.SpellSchoolPhysical,
				GCD:                 time.Second * 1,
				BaseCost: core.ResourceCost{
					Type:  stats.Energy,
					Value: MutilateEnergyCost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Energy,
					Value: MutilateEnergyCost,
				},
				SpellExtras: SpellFlagBuilder,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ProcMask:               core.ProcMaskMeleeMHSpecial,
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
				OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
					// TODO: Need to adjust cast/hit metrics because of the 2-cast behavior.
					if spellEffect.Landed() {
						rogue.AddComboPoints(sim, 2, MutilateActionID)

						spell := &rogue.mutilateDamage
						rogue.mutilateDamageTemplate.Apply(spell)
						spell.Effects[0].Target = spellEffect.Target
						spell.Effects[1].Target = spellEffect.Target
						if rogue.deadlyPoison.IsInUse() {
							spell.Effect.DamageMultiplier *= 1.5
						}
						spell.Cast(sim)
					} else {
						rogue.AddEnergy(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
					}
				},
			},
			WeaponInput: core.WeaponDamageInput{
				DamageMultiplier: 1,
			},
		},
	}

	return core.NewSimpleSpellTemplate(ability)
}

func (rogue *Rogue) NewMutilate(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	mt := &rogue.mutilate
	rogue.mutilateTemplate.Apply(mt)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	mt.Effect.Target = target

	return mt
}
