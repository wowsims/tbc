package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDSmite int32 = 25364

var SmiteCooldownID = core.NewCooldownID()

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
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
			CritMultiplier:      priest.DefaultSpellCritMultiplier(),
		},
		Effect: core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BaseDamage:       core.BaseDamageConfigMagic(549, 616, 0.7143),
		},
	}

	priest.applyTalentsToHolySpell(&template.SpellCast.Cast, &template.Effect)

	template.CastTime -= time.Millisecond * 100 * time.Duration(priest.Talents.DivineFury)
	template.Effect.DamageMultiplier *= (1 + (0.05 * float64(priest.Talents.SearingLight)))
	template.Effect.BonusSpellHitRating += float64(priest.Talents.FocusedPower) * 2 * core.SpellHitRatingPerHitChance // 2% crit per point
	template.Effect.OnSpellHit = priest.applyOnHitTalents

	priest.Smite = priest.RegisterSpell(core.SpellConfig{
		Template: template,
		ModifyCast: func(sim *core.Simulation, target *core.Target, instance *core.SimpleSpell) {
			instance.Effect.Target = target
			priest.applySurgeOfLight(&instance.SpellCast)
		},
	})
}
