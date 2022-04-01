package hunter

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var AimedShotCooldownID = core.NewCooldownID()
var AimedShotActionID = core.ActionID{SpellID: 27065, CooldownID: AimedShotCooldownID}

func (hunter *Hunter) registerAimedShotSpell(sim *core.Simulation) {
	cost := core.ResourceCost{Type: stats.Mana, Value: 370}
	ama := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    AimedShotActionID,
				Character:   &hunter.Character,
				SpellSchool: core.SpellSchoolPhysical,
				// Actual aimed shot has a 2.5s cast time, but we only use it as an instant precast.
				//CastTime:       time.Millisecond * 2500,
				//Cooldown:       time.Second * 6,
				//GCD:            core.GCDDefault,
				Cost:     cost,
				BaseCost: cost,
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategoryRanged,
			CritRollCategory:    core.CritRollCategoryPhysical,
			CritMultiplier:      hunter.critMultiplier(true, sim.GetPrimaryTarget()),
			ProcMask:            core.ProcMaskRangedSpecial,
			DamageMultiplier:    1,
			ThreatMultiplier:    1,
			BaseDamage: hunter.talonOfAlarDamageMod(core.BaseDamageConfig{
				Calculator: func(sim *core.Simulation, hitEffect *core.SpellEffect, spell *core.SimpleSpellTemplate) float64 {
					return (hitEffect.RangedAttackPower(spell.Character)+hitEffect.RangedAttackPowerOnTarget())*0.2 +
						hunter.AutoAttacks.Ranged.BaseDamage(sim) +
						hunter.AmmoDamageBonus +
						hitEffect.BonusWeaponDamage(spell.Character) +
						870
				},
				TargetSpellCoefficient: 1,
			}),
		},
	}

	ama.Cost.Value *= 1 - 0.02*float64(hunter.Talents.Efficiency)

	hunter.AimedShot = hunter.RegisterSpell(core.SpellConfig{
		Template:   ama,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}
