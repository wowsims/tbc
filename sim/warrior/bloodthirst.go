package warrior

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var BloodthirstCooldownID = core.NewCooldownID()
var BloodthirstActionID = core.ActionID{SpellID: 30335, CooldownID: BloodthirstCooldownID}

func (warrior *Warrior) registerBloodthirstSpell(_ *core.Simulation) {
	warrior.bloodthirstCost = 30
	if ItemSetDestroyerBattlegear.CharacterHasSetBonus(&warrior.Character, 4) {
		warrior.bloodthirstCost -= 5
	}
	refundAmount := warrior.bloodthirstCost * 0.8

	ability := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    BloodthirstActionID,
				Character:   &warrior.Character,
				SpellSchool: core.SpellSchoolPhysical,
				GCD:         core.GCDDefault,
				Cooldown:    time.Second * 6,
				IgnoreHaste: true,
				BaseCost: core.ResourceCost{
					Type:  stats.Rage,
					Value: warrior.bloodthirstCost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Rage,
					Value: warrior.bloodthirstCost,
				},
			},
		},
	}

	warrior.Bloodthirst = warrior.RegisterSpell(core.SpellConfig{
		Template: ability,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask: core.ProcMaskMeleeMHSpecial,

			DamageMultiplier: 1 * core.TernaryFloat64(ItemSetOnslaughtBattlegear.CharacterHasSetBonus(&warrior.Character, 4), 1.05, 1),
			ThreatMultiplier: 1,

			BaseDamage: core.BaseDamageConfig{
				Calculator: func(sim *core.Simulation, hitEffect *core.SpellEffect, spell *core.Spell) float64 {
					return hitEffect.MeleeAttackPower(spell.Character) * 0.45
				},
				TargetSpellCoefficient: 0, // Doesn't scale with +damage on target?
			},
			OutcomeApplier: core.OutcomeFuncMeleeSpecialHitAndCrit(warrior.critMultiplier(true)),

			OnInit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if warrior.StanceMatches(DefensiveStance) {
					spellEffect.ThreatMultiplier *= 1 + 0.21*float64(warrior.Talents.TacticalMastery)
				}
			},
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					warrior.AddRage(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
				}
			},
		}),
	})
}

func (warrior *Warrior) CanBloodthirst(sim *core.Simulation) bool {
	return warrior.Talents.Bloodthirst && warrior.CurrentRage() >= warrior.bloodthirstCost && !warrior.IsOnCD(BloodthirstCooldownID, sim.CurrentTime)
}
