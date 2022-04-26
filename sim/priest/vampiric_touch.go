package priest

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var VampiricTouchActionID = core.ActionID{SpellID: 34917}

const VampiricTouchBaseCost = 425.0

func (priest *Priest) registerVampiricTouchSpell(sim *core.Simulation) {
	priest.VampiricTouch = priest.RegisterSpell(core.SpellConfig{
		ActionID:    VampiricTouchActionID,
		SpellSchool: core.SpellSchoolShadow,

		ResourceType: stats.Mana,
		BaseCost:     VampiricTouchBaseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost:     VampiricTouchBaseCost,
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond * 1500,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			BonusSpellHitRating: float64(priest.Talents.ShadowFocus) * 2 * core.SpellHitRatingPerHitChance,
			ThreatMultiplier:    1 - 0.08*float64(priest.Talents.ShadowAffinity),
			OutcomeApplier:      priest.OutcomeFuncMagicHit(),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					priest.VampiricTouchDot.Apply(sim)
				}
			},
		}),
	})

	target := sim.GetPrimaryTarget()
	priest.VampiricTouchDot = core.NewDot(core.Dot{
		Spell: priest.VampiricTouch,
		Aura: target.RegisterAura(core.Aura{
			Label:    "VampiricTouch-" + strconv.Itoa(int(priest.Index)),
			ActionID: VampiricTouchActionID,
		}),

		NumberOfTicks: 5,
		TickLength:    time.Second * 3,

		TickEffects: core.TickFuncSnapshot(target, core.SpellEffect{
			DamageMultiplier: 1 *
				(1 + float64(priest.Talents.Darkness)*0.02) *
				core.TernaryFloat64(priest.Talents.Shadowform, 1.15, 1),
			ThreatMultiplier: 1 - 0.08*float64(priest.Talents.ShadowAffinity),
			IsPeriodic:       true,
			BaseDamage:       core.BaseDamageConfigMagicNoRoll(650/5, 0.2),
			OutcomeApplier:   priest.OutcomeFuncTick(),
		}),
	})
}
