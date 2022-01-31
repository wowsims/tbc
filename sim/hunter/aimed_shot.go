package hunter

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var AimedShotCooldownID = core.NewCooldownID()
var AimedShotActionID = core.ActionID{SpellID: 27065, CooldownID: AimedShotCooldownID}

func (hunter *Hunter) newAimedShotTemplate(sim *core.Simulation) core.MeleeAbilityTemplate {
	ama := core.ActiveMeleeAbility{
		MeleeAbility: core.MeleeAbility{
			ActionID:       AimedShotActionID,
			Character:      &hunter.Character,
			SpellSchool:    stats.AttackPower,
			CritMultiplier: 2.0,
			// Actual aimed shot has a 2.5s cast time, but we only use it as an instant precast.
			//CastTime:       time.Millisecond * 2500,
			//Cooldown:       time.Second * 6,
			//GCD:            core.GCDDefault,
			Cost: core.ResourceCost{
				Type:  stats.Mana,
				Value: 370,
			},
		},
		Effect: core.AbilityHitEffect{
			AbilityEffect: core.AbilityEffect{
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
			},
			WeaponInput: core.WeaponDamageInput{
				IsRanged: true,
				CalculateDamage: func(attackPower float64, bonusWeaponDamage float64) float64 {
					return attackPower*0.2 +
						hunter.AmmoDamageBonus +
						hunter.AutoAttacks.Ranged.BaseDamage(sim) +
						bonusWeaponDamage +
						870
				},
			},
		},
	}

	return core.NewMeleeAbilityTemplate(ama)
}

func (hunter *Hunter) NewAimedShot(sim *core.Simulation, target *core.Target) *core.ActiveMeleeAbility {
	as := &hunter.aimedShot
	hunter.aimedShotTemplate.Apply(as)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	as.Effect.Target = target

	return as
}
