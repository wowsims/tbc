package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var ExorcismCD = core.NewCooldownID()
var ExorcismActionID = core.ActionID{SpellID: 10314, CooldownID: ExorcismCD}

func (paladin *Paladin) newExorcismTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	exo := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            ExorcismActionID,
				Character:           &paladin.Character,
				CritRollCategory:    core.CritRollCategoryMagical,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				SpellSchool:         core.SpellSchoolHoly,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 295,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 295,
				},
				Cooldown:       time.Second * 15,
				CritMultiplier: paladin.SpellCritMultiplier(1, 0.25), // look up crit multiplier in the future
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier: 1,
				ThreatMultiplier: 1,
			},
			BaseDamage: core.BaseDamageConfigMagic(521, 579, 1),
		},
	}

	return core.NewSimpleSpellTemplate(exo)
}

func (paladin *Paladin) NewExorcism(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	if target.MobType != proto.MobType_MobTypeUndead && target.MobType != proto.MobType_MobTypeDemon {
		return nil
	}

	exo := &paladin.exorcismSpell
	paladin.exorcismTemplate.Apply(exo)

	exo.Effect.Target = target

	exo.Init(sim)

	return exo
}
