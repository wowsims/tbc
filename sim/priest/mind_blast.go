package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDMindBlast int32 = 25375

var MBCooldownID = core.NewCooldownID()
var MindBlastActionID = core.ActionID{SpellID: SpellIDMindBlast, CooldownID: MBCooldownID}

func (priest *Priest) registerMindBlastSpell(sim *core.Simulation) {
	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    MindBlastActionID,
				Character:   &priest.Character,
				SpellSchool: core.SpellSchoolShadow,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 450,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 450,
				},
				CastTime: time.Millisecond * 1500,
				GCD:      core.GCDDefault,
				Cooldown: time.Second * 8,
			},
		},
	}
	template.Cooldown -= time.Millisecond * 500 * time.Duration(priest.Talents.ImprovedMindBlast)
	template.Cost.Value -= template.BaseCost.Value * float64(priest.Talents.FocusedMind) * 0.05

	priest.MindBlast = priest.RegisterSpell(core.SpellConfig{
		Template: template,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{

			BonusSpellHitRating: 0 +
				float64(priest.Talents.ShadowFocus)*2*core.SpellHitRatingPerHitChance +
				float64(priest.Talents.FocusedPower)*2*core.SpellHitRatingPerHitChance,

			BonusSpellCritRating: float64(priest.Talents.ShadowPower) * 3 * core.SpellCritRatingPerCritChance,

			DamageMultiplier: 1 *
				(1 + float64(priest.Talents.Darkness)*0.02) *
				core.TernaryFloat64(priest.Talents.Shadowform, 1.15, 1) *
				core.TernaryFloat64(ItemSetAbsolution.CharacterHasSetBonus(&priest.Character, 4), 1.1, 1),

			ThreatMultiplier: 1 - 0.08*float64(priest.Talents.ShadowAffinity),

			BaseDamage:     core.BaseDamageConfigMagic(711, 752, 0.429),
			OutcomeApplier: core.OutcomeFuncMagicHitAndCrit(priest.DefaultSpellCritMultiplier()),
		}),
	})
}
