package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDSmite int32 = 25364

var SmiteActionID = core.ActionID{SpellID: SpellIDSmite}

func (priest *Priest) registerSmiteSpell(sim *core.Simulation) {
	baseCost := 385.0

	normalOutcome := priest.OutcomeFuncMagicHitAndCrit(priest.DefaultSpellCritMultiplier())
	surgeOfLightOutcome := priest.OutcomeFuncMagicHit()

	priest.Smite = priest.RegisterSpell(core.SpellConfig{
		ActionID:    SmiteActionID,
		SpellSchool: core.SpellSchoolHoly,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost:     baseCost,
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond*2500 - time.Millisecond*100*time.Duration(priest.Talents.DivineFury),
			},
			ModifyCast: priest.applySurgeOfLight,
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask:            core.ProcMaskSpellDamage,
			BonusSpellHitRating: float64(priest.Talents.FocusedPower) * 2 * core.SpellHitRatingPerHitChance,

			BonusSpellCritRating: float64(priest.Talents.HolySpecialization) * 1 * core.SpellCritRatingPerCritChance,

			DamageMultiplier: 1 + 0.05*float64(priest.Talents.SearingLight),

			ThreatMultiplier: 1 - 0.04*float64(priest.Talents.SilentResolve),

			BaseDamage: core.BaseDamageConfigMagic(549, 616, 0.7143),
			OutcomeApplier: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect, damage *float64) {
				if priest.SurgeOfLightProcAura.IsActive() {
					surgeOfLightOutcome(sim, spell, spellEffect, damage)
				} else {
					normalOutcome(sim, spell, spellEffect, damage)
				}
			},
		}),
	})
}
