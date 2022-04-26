package warrior

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var SlamActionID = core.ActionID{SpellID: 25242}

func (warrior *Warrior) registerSlamSpell(_ *core.Simulation) {
	cost := 15.0 - float64(warrior.Talents.FocusedRage)
	refundAmount := cost * 0.8

	warrior.Slam = warrior.RegisterSpell(core.SpellConfig{
		ActionID:    SlamActionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasMeleeMetrics,

		ResourceType: stats.Rage,
		BaseCost:     cost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost:     cost,
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond*1500 - time.Millisecond*500*time.Duration(warrior.Talents.ImprovedSlam),
			},
			IgnoreHaste: true,
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask: core.ProcMaskMeleeMHSpecial,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			BaseDamage:     core.BaseDamageConfigMeleeWeapon(core.MainHand, false, 140, 1, true),
			OutcomeApplier: warrior.OutcomeFuncMeleeSpecialHitAndCrit(warrior.critMultiplier(true)),

			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					warrior.AddRage(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
				}
			},
		}),
	})
}

func (warrior *Warrior) CanSlam() bool {
	return warrior.CurrentRage() >= warrior.Slam.DefaultCast.Cost
}

func (warrior *Warrior) CastSlam(sim *core.Simulation, target *core.Target) bool {
	warrior.AutoAttacks.DelayAllUntil(sim, sim.CurrentTime+warrior.Slam.DefaultCast.CastTime+warrior.AutoAttacks.MainhandSwingSpeed())
	return warrior.Slam.Cast(sim, target)
}
