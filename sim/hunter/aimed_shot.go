package hunter

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var AimedShotCooldownID = core.NewCooldownID()
var AimedShotActionID = core.ActionID{SpellID: 27065, CooldownID: AimedShotCooldownID}

func (hunter *Hunter) newAimedShotTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	cost := core.ResourceCost{Type: stats.Mana, Value: 370}
	ama := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            AimedShotActionID,
				Character:           &hunter.Character,
				OutcomeRollCategory: core.OutcomeRollCategoryRanged,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				// Actual aimed shot has a 2.5s cast time, but we only use it as an instant precast.
				//CastTime:       time.Millisecond * 2500,
				//Cooldown:       time.Second * 6,
				//GCD:            core.GCDDefault,
				Cost:           cost,
				BaseCost:       cost,
				CritMultiplier: hunter.critMultiplier(true, sim.GetPrimaryTarget()),
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ProcMask:               core.ProcMaskRangedSpecial,
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
			},
			BaseDamage: hunter.talonOfAlarDamageMod(func(sim *core.Simulation, hitEffect *core.SpellHitEffect, spellCast *core.SpellCast) float64 {
				return (hitEffect.RangedAttackPower(spellCast)+hitEffect.RangedAttackPowerOnTarget())*0.2 +
					hunter.AutoAttacks.Ranged.BaseDamage(sim) +
					hunter.AmmoDamageBonus +
					hitEffect.PlusWeaponDamage(spellCast) +
					870
			}),
		},
	}

	ama.Cost.Value *= 1 - 0.02*float64(hunter.Talents.Efficiency)

	return core.NewSimpleSpellTemplate(ama)
}

func (hunter *Hunter) NewAimedShot(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	as := &hunter.aimedShot
	hunter.aimedShotTemplate.Apply(as)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	as.Effect.Target = target
	as.Init(sim)
	return as
}
