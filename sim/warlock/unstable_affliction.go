package warlock

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDUnstableAff3 int32 = 30405

var UnstableAff3ActionID = core.ActionID{SpellID: SpellIDUnstableAff3}

func (warlock *Warlock) registerUnstableAffSpell(sim *core.Simulation) {
	baseCost := 400.0
	warlock.UnstableAff = warlock.RegisterSpell(core.SpellConfig{
		ActionID:     UnstableAff3ActionID,
		SpellSchool:  core.SpellSchoolShadow,
		ResourceType: stats.Mana,
		BaseCost:     baseCost,
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost:     baseCost,
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond * 1500,
			},
		},
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			OutcomeApplier: warlock.OutcomeFuncMagicHit(),
			OnSpellHit:     applyDotOnLanded(&warlock.UnstableAffDot),
		}),
	})

	target := sim.GetPrimaryTarget()
	spellCoefficient := 0.2
	warlock.UnstableAffDot = core.NewDot(core.Dot{
		Spell: warlock.UnstableAff,
		Aura: target.RegisterAura(core.Aura{
			Label:    "unstableaff-" + strconv.Itoa(int(warlock.Index)),
			ActionID: UnstableAff3ActionID,
		}),
		NumberOfTicks: 6,
		TickLength:    time.Second * 3,
		TickEffects: core.TickFuncSnapshot(target, core.SpellEffect{
			DamageMultiplier: 1 * (1 + 0.02*float64(warlock.Talents.ShadowMastery)),
			ThreatMultiplier: 1 - 0.05*float64(warlock.Talents.ImprovedDrainSoul),
			BaseDamage:       core.BaseDamageConfigMagicNoRoll(1050/6, spellCoefficient),
			OutcomeApplier:   warlock.OutcomeFuncTick(),
			IsPeriodic:       true,
		}),
	})
}
