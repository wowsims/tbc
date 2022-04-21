package warrior

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var BloodthirstActionID = core.ActionID{SpellID: 30335}

func (warrior *Warrior) registerBloodthirstSpell(_ *core.Simulation) {
	cost := 30.0
	if ItemSetDestroyerBattlegear.CharacterHasSetBonus(&warrior.Character, 4) {
		cost -= 5
	}
	refundAmount := cost * 0.8

	warrior.Bloodthirst = warrior.RegisterSpell(core.SpellConfig{
		ActionID:    BloodthirstActionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasMeleeMetrics,

		ResourceType: stats.Rage,
		BaseCost:     cost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: cost,
				GCD:  core.GCDDefault,
			},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    warrior.NewTimer(),
				Duration: time.Second * 6,
			},
		},

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
	return warrior.Talents.Bloodthirst && warrior.CurrentRage() >= warrior.Bloodthirst.DefaultCast.Cost && warrior.Bloodthirst.IsReady(sim)
}
