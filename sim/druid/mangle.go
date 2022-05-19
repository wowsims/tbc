package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (druid *Druid) registerMangleSpell(cdTimer *core.Timer) {
	if !druid.Talents.Mangle {
		return
	}

	cost := 20.0 - float64(druid.Talents.Ferocity)
	refundAmount := cost * 0.8

	debuff := core.MangleAura(druid.CurrentTarget)

	druid.Mangle = druid.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 33987},
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
				Timer:    cdTimer,
				Duration: time.Second * 6,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask: core.ProcMaskMeleeMHSpecial,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			BaseDamage:     core.BaseDamageConfigMeleeWeapon(core.MainHand, true, 155, 1.15, true),
			OutcomeApplier: druid.OutcomeFuncMeleeSpecialHitAndCrit(druid.MeleeCritMultiplier()),

			OnSpellHitDealt: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					debuff.Activate(sim)
				} else {
					druid.AddRage(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
				}
			},
		}),
	})
}

func (druid *Druid) CanMangle(sim *core.Simulation) bool {
	return druid.Mangle != nil && druid.CurrentRage() >= druid.Mangle.DefaultCast.Cost && druid.Mangle.IsReady(sim)
}
