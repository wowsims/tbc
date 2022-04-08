package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDSmite int32 = 25364

var SmiteCooldownID = core.NewCooldownID()
var SmiteActionID = core.ActionID{SpellID: SpellIDSmite}

func (priest *Priest) registerSmiteSpell(sim *core.Simulation) {
	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID: core.ActionID{
					SpellID:    SpellIDSmite,
					CooldownID: SmiteCooldownID,
				},
				Character:   &priest.Character,
				SpellSchool: core.SpellSchoolHoly,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 385,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 385,
				},
				CastTime: time.Millisecond * 2500,
				GCD:      core.GCDDefault,
				Cooldown: time.Second * 0,
			},
		},
	}
	template.CastTime -= time.Millisecond * 100 * time.Duration(priest.Talents.DivineFury)

	priest.Smite = priest.RegisterSpell(core.SpellConfig{
		Template: template,
		ModifyCast: func(sim *core.Simulation, target *core.Target, instance *core.SimpleSpell) {
			priest.applySurgeOfLight(&instance.SpellCast)
		},
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{

			BonusSpellHitRating: float64(priest.Talents.FocusedPower) * 2 * core.SpellHitRatingPerHitChance,

			BonusSpellCritRating: float64(priest.Talents.HolySpecialization) * 1 * core.SpellCritRatingPerCritChance,

			DamageMultiplier: 1 + 0.05*float64(priest.Talents.SearingLight),

			ThreatMultiplier: 1 - 0.04*float64(priest.Talents.SilentResolve),

			BaseDamage:     core.BaseDamageConfigMagic(549, 616, 0.7143),
			OutcomeApplier: core.OutcomeFuncMagicHitAndCrit(priest.DefaultSpellCritMultiplier()),

			OnSpellHit: priest.applyOnHitTalents,
		}),
	})
}
