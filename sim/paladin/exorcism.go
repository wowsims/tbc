package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var ExorcismActionID = core.ActionID{SpellID: 10314}

func (paladin *Paladin) registerExorcismSpell(sim *core.Simulation) {
	baseCost := 295.0

	paladin.Exorcism = paladin.RegisterSpell(core.SpellConfig{
		ActionID:    ExorcismActionID,
		SpellSchool: core.SpellSchoolHoly,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost,
			},
			CD: core.Cooldown{
				Timer:    paladin.NewTimer(),
				Duration: time.Second * 15,
			},
		},

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
