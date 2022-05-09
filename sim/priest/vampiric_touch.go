package priest

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (priest *Priest) registerVampiricTouchSpell() {
	actionID := core.ActionID{SpellID: 34917}
	baseCost := 425.0

	priest.VampiricTouch = priest.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolShadow,

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
			BonusSpellHitRating: float64(priest.Talents.ShadowFocus) * 2 * core.SpellHitRatingPerHitChance,
			ThreatMultiplier:    1 - 0.08*float64(priest.Talents.ShadowAffinity),
			OutcomeApplier:      priest.OutcomeFuncMagicHit(),
			OnSpellHitDealt: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					priest.VampiricTouchDot.Apply(sim)
				}
			},
			ProcMask: core.ProcMaskEmpty,
		}),
	})

	target := priest.CurrentTarget
	priest.VampiricTouchDot = core.NewDot(core.Dot{
		Spell: priest.VampiricTouch,
		Aura: target.RegisterAura(core.Aura{
			Label:    "VampiricTouch-" + strconv.Itoa(int(priest.Index)),
			ActionID: actionID,
		}),

		NumberOfTicks: 5,
		TickLength:    time.Second * 3,

		TickEffects: core.TickFuncSnapshot(target, core.SpellEffect{
			DamageMultiplier: 1 *
				(1 + float64(priest.Talents.Darkness)*0.02) *
				core.TernaryFloat64(priest.Talents.Shadowform, 1.15, 1),
			ThreatMultiplier: 1 - 0.08*float64(priest.Talents.ShadowAffinity),
			IsPeriodic:       true,
			ProcMask:         core.ProcMaskPeriodicDamage,
			BaseDamage:       core.BaseDamageConfigMagicNoRoll(650/5, 0.2),
			OutcomeApplier:   priest.OutcomeFuncTick(),
		}),
	})
}
