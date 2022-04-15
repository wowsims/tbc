package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDShadowWordDeath int32 = 32996

var SWDCooldownID = core.NewCooldownID()
var ShadowWordDeathActionID = core.ActionID{SpellID: SpellIDShadowWordDeath, CooldownID: SWDCooldownID}

func (priest *Priest) registerShadowWordDeathSpell(sim *core.Simulation) {
	baseCost := 309.0

	priest.ShadowWordDeath = priest.RegisterSpell(core.SpellConfig{
		ActionID:    ShadowWordDeathActionID,
		SpellSchool: core.SpellSchoolShadow,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost * (1 - 0.02*float64(priest.Talents.MentalAgility)),
				GCD:  core.GCDDefault,
			},
			Cooldown: time.Second * 12,
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{

			BonusSpellHitRating: float64(priest.Talents.ShadowFocus) * 2 * core.SpellHitRatingPerHitChance,

			BonusSpellCritRating: float64(priest.Talents.ShadowPower) * 3 * core.SpellCritRatingPerCritChance,

			DamageMultiplier: 1 *
				(1 + float64(priest.Talents.Darkness)*0.02) *
				core.TernaryFloat64(priest.Talents.Shadowform, 1.15, 1),

			ThreatMultiplier: 1 - 0.08*float64(priest.Talents.ShadowAffinity),

			BaseDamage:     core.BaseDamageConfigMagic(572, 664, 0.429),
			OutcomeApplier: core.OutcomeFuncMagicHitAndCrit(priest.DefaultSpellCritMultiplier()),
		}),
	})
}
