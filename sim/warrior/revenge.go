package warrior

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (warrior *Warrior) registerRevengeSpell(cdTimer *core.Timer) {
	cost := 5.0 - float64(warrior.Talents.FocusedRage)
	refundAmount := cost * 0.8

	warrior.Revenge = warrior.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 30357},
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
				Duration: time.Second * 5,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask: core.ProcMaskMeleeMHSpecial,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			FlatThreatBonus:  200,

			BaseDamage:     core.BaseDamageConfigRoll(414, 506),
			OutcomeApplier: warrior.OutcomeFuncMeleeSpecialHitAndCrit(warrior.critMultiplier(true)),

			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					warrior.AddRage(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
				}
			},
		}),
	})
}

func (warrior *Warrior) CanRevenge(sim *core.Simulation) bool {
	return warrior.StanceMatches(DefensiveStance) && warrior.revengeTriggered && warrior.CurrentRage() >= warrior.Revenge.DefaultCast.Cost && warrior.Revenge.IsReady(sim)
}
