package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (druid *Druid) registerDemoralizingRoarSpell() {
	cost := 10.0

	baseEffect := core.SpellEffect{
		ThreatMultiplier: 1,
		FlatThreatBonus:  56,
		OutcomeApplier:   druid.OutcomeFuncMagicHit(),
	}

	numHits := druid.Env.GetNumTargets()
	effects := make([]core.SpellEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = druid.Env.GetTargetUnit(i)

		demoRoarAura := core.DemoralizingRoarAura(effects[i].Target, druid.Talents.FeralAggression)
		if i == 0 {
			druid.DemoralizingRoarAura = demoRoarAura
		}

		effects[i].OnSpellHitDealt = func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spellEffect.Landed() {
				demoRoarAura.Activate(sim)
			}
		}
	}

	druid.DemoralizingRoar = druid.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 26998},
		SpellSchool: core.SpellSchoolPhysical,

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

func (druid *Druid) CanDemoralizingRoar(sim *core.Simulation) bool {
	return druid.CurrentRage() >= druid.DemoralizingRoar.DefaultCast.Cost
}

func (druid *Druid) ShouldDemoralizingRoar(sim *core.Simulation, filler bool, maintainOnly bool) bool {
	if !druid.CanDemoralizingRoar(sim) {
		return false
	}

	activeDebuff := druid.CurrentTarget.GetActiveAuraWithTag(core.APReductionAuraTag)
	if activeDebuff != nil && activeDebuff.Priority > druid.DemoralizingRoarAura.Priority {
		return false
	}

	if filler {
		return true
	}

	if maintainOnly {
		return activeDebuff == nil || activeDebuff.Priority < druid.DemoralizingRoarAura.Priority || activeDebuff.RemainingDuration(sim) < time.Second*2
	}

	return false
}
