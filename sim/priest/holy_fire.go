package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDHolyFire int32 = 25384

func (priest *Priest) registerHolyFireSpell(sim *core.Simulation) {
	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    core.ActionID{SpellID: SpellIDHolyFire},
				Character:   &priest.Character,
				SpellSchool: core.SpellSchoolHoly,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 290,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 290,
				},
				CastTime: time.Millisecond * 3500,
				GCD:      core.GCDDefault,
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
			CritMultiplier:      priest.DefaultSpellCritMultiplier(),
			DamageMultiplier:    1,
			ThreatMultiplier:    1,
			BaseDamage:          core.BaseDamageConfigMagic(426, 537, 0.8571),
			DotInput: core.DotDamageInput{
				NumberOfTicks:  5,
				TickLength:     time.Second * 2,
				TickBaseDamage: core.DotSnapshotFuncMagic(33, 0.17),
				Aura:           priest.NewDotAura("Holy Fire", core.ActionID{SpellID: SpellIDHolyFire}),
			},
		},
	}

	priest.applyTalentsToHolySpell(&template.SpellCast.Cast, &template.Effect)
	template.CastTime -= time.Millisecond * 100 * time.Duration(priest.Talents.DivineFury)
	template.Effect.DamageMultiplier *= (1 + (0.05 * float64(priest.Talents.SearingLight)))

	priest.HolyFire = priest.RegisterSpell(core.SpellConfig{
		Template:   template,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}
