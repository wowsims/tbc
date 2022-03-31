package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var ArcaneShotCooldownID = core.NewCooldownID()
var ArcaneShotActionID = core.ActionID{SpellID: 27019, CooldownID: ArcaneShotCooldownID}

func (hunter *Hunter) registerArcaneShotSpell(sim *core.Simulation) {
	cost := core.ResourceCost{Type: stats.Mana, Value: 230}
	ama := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    ArcaneShotActionID,
				Character:   &hunter.Character,
				SpellSchool: core.SpellSchoolArcane,
				GCD:         core.GCDDefault + hunter.latency,
				IgnoreHaste: true,
				Cooldown:    time.Second * 6,
				Cost:        cost,
				BaseCost:    cost,
			},
			OutcomeRollCategory: core.OutcomeRollCategoryRanged,
			CritRollCategory:    core.CritRollCategoryPhysical,
			CritMultiplier:      hunter.critMultiplier(true, sim.GetPrimaryTarget()),
		},
		Effect: core.SpellEffect{
			ProcMask:         core.ProcMaskRangedSpecial,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BaseDamage: hunter.talonOfAlarDamageMod(core.BaseDamageConfig{
				Calculator: func(sim *core.Simulation, hitEffect *core.SpellEffect, spellCast *core.SpellCast) float64 {
					return (hitEffect.RangedAttackPower(spellCast)+hitEffect.RangedAttackPowerOnTarget())*0.15 + 273
				},
				TargetSpellCoefficient: 1,
			}),
		},
	}

	ama.Cost.Value *= 1 - 0.02*float64(hunter.Talents.Efficiency)
	ama.Cooldown -= time.Millisecond * 200 * time.Duration(hunter.Talents.ImprovedArcaneShot)

	hunter.ArcaneShot = hunter.RegisterSpell(core.SpellConfig{
		Template:   ama,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}
