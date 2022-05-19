package druid

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (druid *Druid) registerSwipeSpell() {
	cost := 20.0 - float64(druid.Talents.Ferocity)

	baseEffect := core.SpellEffect{
		ProcMask: core.ProcMaskMeleeMHSpecial,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		BaseDamage: core.BaseDamageConfig{
			Calculator: func(sim *core.Simulation, hitEffect *core.SpellEffect, spell *core.Spell) float64 {
				return 84 + 0.07*hitEffect.MeleeAttackPower(spell.Unit)
			},
			TargetSpellCoefficient: 1,
		},
		OutcomeApplier: druid.OutcomeFuncMeleeSpecialHitAndCrit(druid.MeleeCritMultiplier()),
	}

	numHits := core.MinInt32(3, druid.Env.GetNumTargets())
	effects := make([]core.SpellEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effect := baseEffect
		effect.Target = druid.Env.GetTargetUnit(i)
		effects = append(effects, effect)
	}

	druid.Swipe = druid.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 26997},
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

		ApplyEffects: core.ApplyEffectFuncDamageMultiple(effects),
	})
}

func (druid *Druid) CanSwipe(sim *core.Simulation) bool {
	return druid.CurrentRage() >= druid.Swipe.DefaultCast.Cost
}
