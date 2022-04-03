package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDMindBlast int32 = 25375

var MBCooldownID = core.NewCooldownID()

func (priest *Priest) registerMindBlastSpell(sim *core.Simulation) {
	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID: core.ActionID{
					SpellID:    SpellIDMindBlast,
					CooldownID: MBCooldownID,
				},
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
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
			CritMultiplier:      priest.DefaultSpellCritMultiplier(),
			DamageMultiplier:    1,
			ThreatMultiplier:    1,
			BaseDamage:          core.BaseDamageConfigMagic(711, 752, 0.429),
		},
	}

	priest.applyTalentsToShadowSpell(&template.SpellCast.Cast, &template.Effect)
	template.Cooldown -= time.Millisecond * 500 * time.Duration(priest.Talents.ImprovedMindBlast)
	template.Effect.BonusSpellHitRating += float64(priest.Talents.FocusedPower) * 2 * core.SpellHitRatingPerHitChance // 2% crit per point

	if ItemSetAbsolution.CharacterHasSetBonus(&priest.Character, 4) { // Absolution 4p adds 10% damage
		template.Effect.DamageMultiplier *= 1.1
	}

	priest.MindBlast = priest.RegisterSpell(core.SpellConfig{
		Template:   template,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}
