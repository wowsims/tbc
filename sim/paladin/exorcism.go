package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var ExorcismCD = core.NewCooldownID()
var ExorcismActionID = core.ActionID{SpellID: 10314, CooldownID: ExorcismCD}

func (paladin *Paladin) registerExorcismSpell(sim *core.Simulation) {
	exo := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    ExorcismActionID,
				Character:   &paladin.Character,
				SpellSchool: core.SpellSchoolHoly,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 295,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 295,
				},
				Cooldown: time.Second * 15,
			},
		},
	}

	paladin.Exorcism = paladin.RegisterSpell(core.SpellConfig{
		Template: exo,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			BaseDamage: core.BaseDamageConfigMagic(521, 579, 1),
			// look up crit multiplier in the future
			// TODO: What is this 0.25?
			OutcomeApplier: core.OutcomeFuncMagicHitAndCrit(paladin.SpellCritMultiplier(1, 0.25)),
		}),
	})
}

func (paladin *Paladin) CanExorcism(target *core.Target) bool {
	return target.MobType == proto.MobType_MobTypeUndead || target.MobType == proto.MobType_MobTypeDemon
}
