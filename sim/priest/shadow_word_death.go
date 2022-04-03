package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDShadowWordDeath int32 = 32996

var SWDCooldownID = core.NewCooldownID()

func (priest *Priest) registerShadowWordDeathSpell(sim *core.Simulation) {
	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID: core.ActionID{
					SpellID:    SpellIDShadowWordDeath,
					CooldownID: SWDCooldownID,
				},
				Character:   &priest.Character,
				SpellSchool: core.SpellSchoolShadow,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 309,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 309,
				},
				CastTime: 0,
				GCD:      core.GCDDefault,
				Cooldown: time.Second * 12,
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
			CritMultiplier:      priest.DefaultSpellCritMultiplier(),
			DamageMultiplier:    1,
			ThreatMultiplier:    1,
			BaseDamage:          core.BaseDamageConfigMagic(572, 664, 0.429),
		},
	}

	priest.applyTalentsToShadowSpell(&template.SpellCast.Cast, &template.Effect)

	priest.ShadowWordDeath = priest.RegisterSpell(core.SpellConfig{
		Template:   template,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}
