package priest

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDShadowWordPain int32 = 25368

var ShadowWordPainActionID = core.ActionID{SpellID: SpellIDShadowWordPain}

func (priest *Priest) registerShadowWordPainSpell(sim *core.Simulation) {
	baseCost := 575.0

	priest.ShadowWordPain = priest.RegisterSpell(core.SpellConfig{
		ActionID:    ShadowWordPainActionID,
		SpellSchool: core.SpellSchoolShadow,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost * (1 - 0.02*float64(priest.Talents.MentalAgility)),
				GCD:  core.GCDDefault,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask:            core.ProcMaskEmpty,
			BonusSpellHitRating: float64(priest.Talents.ShadowFocus) * 2 * core.SpellHitRatingPerHitChance,
			ThreatMultiplier:    1 - 0.08*float64(priest.Talents.ShadowAffinity),
			OutcomeApplier:      priest.OutcomeFuncMagicHit(),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					priest.ShadowWordPainDot.Apply(sim)
				}
			},
		}),
	})

	target := sim.GetPrimaryTarget()
	priest.ShadowWordPainDot = core.NewDot(core.Dot{
		Spell: priest.ShadowWordPain,
		Aura: target.RegisterAura(core.Aura{
			Label:    "ShadowWordPain-" + strconv.Itoa(int(priest.Index)),
			ActionID: ShadowWordPainActionID,
		}),

		NumberOfTicks: 6 +
			int(priest.Talents.ImprovedShadowWordPain) +
			core.TernaryInt(ItemSetAbsolution.CharacterHasSetBonus(&priest.Character, 2), 1, 0),
		TickLength: time.Second * 3,

		TickEffects: core.TickFuncSnapshot(target, core.SpellEffect{
			DamageMultiplier: 1 *
				(1 + float64(priest.Talents.Darkness)*0.02) *
				core.TernaryFloat64(priest.Talents.Shadowform, 1.15, 1),
			ThreatMultiplier: 1 - 0.08*float64(priest.Talents.ShadowAffinity),
			IsPeriodic:       true,
			ProcMask:         core.ProcMaskPeriodicDamage,
			BaseDamage:       core.BaseDamageConfigMagicNoRoll(1236/6, 0.183),
			OutcomeApplier:   priest.OutcomeFuncTick(),
		}),
	})
}
