package priest

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDDevouringPlague int32 = 25467

var DevouringPlagueActionID = core.ActionID{SpellID: SpellIDDevouringPlague}

func (priest *Priest) registerDevouringPlagueSpell(sim *core.Simulation) {
	baseCost := 1145.0

	priest.DevouringPlague = priest.RegisterSpell(core.SpellConfig{
		ActionID:    DevouringPlagueActionID,
		SpellSchool: core.SpellSchoolShadow,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost * (1 - 0.02*float64(priest.Talents.MentalAgility)),
				GCD:  core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    priest.NewTimer(),
				Duration: time.Minute * 3,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask:            core.ProcMaskEmpty,
			BonusSpellHitRating: float64(priest.Talents.ShadowFocus) * 2 * core.SpellHitRatingPerHitChance,
			ThreatMultiplier:    1 - 0.08*float64(priest.Talents.ShadowAffinity),
			OutcomeApplier:      priest.OutcomeFuncMagicHit(),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					priest.DevouringPlagueDot.Apply(sim)
				}
			},
		}),
	})

	target := sim.GetPrimaryTarget()
	priest.DevouringPlagueDot = core.NewDot(core.Dot{
		Spell: priest.DevouringPlague,
		Aura: target.RegisterAura(core.Aura{
			Label:    "DevouringPlague-" + strconv.Itoa(int(priest.Index)),
			ActionID: DevouringPlagueActionID,
		}),
		NumberOfTicks: 8,
		TickLength:    time.Second * 3,
		TickEffects: core.TickFuncSnapshot(target, core.SpellEffect{
			DamageMultiplier: 1 *
				(1 + float64(priest.Talents.Darkness)*0.02) *
				core.TernaryFloat64(priest.Talents.Shadowform, 1.15, 1),
			ThreatMultiplier: 1 - 0.08*float64(priest.Talents.ShadowAffinity),
			IsPeriodic:       true,
			BaseDamage:       core.BaseDamageConfigMagicNoRoll(1216/8, 0.1),
			OutcomeApplier:   priest.OutcomeFuncTick(),
			ProcMask:         core.ProcMaskPeriodicDamage,
		}),
	})
}
