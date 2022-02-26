package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var ArcaneShotCooldownID = core.NewCooldownID()
var ArcaneShotActionID = core.ActionID{SpellID: 27019, CooldownID: ArcaneShotCooldownID}

func (hunter *Hunter) newArcaneShotTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	cost := core.ResourceCost{Type: stats.Mana, Value: 230}
	ama := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            ArcaneShotActionID,
				Character:           &hunter.Character,
				OutcomeRollCategory: core.OutcomeRollCategoryRanged,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolArcane,
				GCD:                 core.GCDDefault,
				IgnoreHaste:         true,
				Cooldown:            time.Second * 6,
				Cost:                cost,
				BaseCost:            cost,
				CritMultiplier:      hunter.critMultiplier(true, sim.GetPrimaryTarget()),
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ProcMask:               core.ProcMaskRangedSpecial,
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
				IgnoreArmor:            true,
			},
			WeaponInput: core.WeaponDamageInput{
				CalculateDamage: func(attackPower float64, bonusWeaponDamage float64) float64 {
					return attackPower*0.15 + 273
				},
			},
		},
	}

	ama.Cost.Value *= 1 - 0.02*float64(hunter.Talents.Efficiency)
	ama.Cooldown -= time.Millisecond * 200 * time.Duration(hunter.Talents.ImprovedArcaneShot)

	return core.NewSimpleSpellTemplate(ama)
}

func (hunter *Hunter) NewArcaneShot(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	as := &hunter.arcaneShot
	hunter.arcaneShotTemplate.Apply(as)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	as.Effect.Target = target

	as.Init(sim)

	return as
}
