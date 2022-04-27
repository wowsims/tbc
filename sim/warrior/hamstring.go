package warrior

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (warrior *Warrior) registerHamstringSpell() {
	cost := 10 - float64(warrior.Talents.FocusedRage)
	refundAmount := cost * 0.8

	warrior.Hamstring = warrior.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 25212},
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
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask: core.ProcMaskMeleeMHSpecial,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			BaseDamage: core.BaseDamageConfig{
				Calculator:             core.BaseDamageFuncFlat(63),
				TargetSpellCoefficient: 1,
			},
			OutcomeApplier: warrior.OutcomeFuncMeleeSpecialHitAndCrit(warrior.critMultiplier(true)),

			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					warrior.AddRage(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
				}
			},
		}),
	})
}

func (warrior *Warrior) ShouldHamstring(sim *core.Simulation) bool {
	return warrior.CurrentRage() >= warrior.Hamstring.DefaultCast.Cost &&
		(!warrior.Talents.Bloodthirst || warrior.Bloodthirst.TimeToReady(sim) > core.GCDDefault) &&
		(!warrior.Talents.MortalStrike || warrior.MortalStrike.TimeToReady(sim) > core.GCDDefault) &&
		warrior.Whirlwind.TimeToReady(sim) > core.GCDDefault
}
