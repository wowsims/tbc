package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var MultiShotCooldownID = core.NewCooldownID()
var MultiShotActionID = core.ActionID{SpellID: 27021, CooldownID: MultiShotCooldownID}

func (hunter *Hunter) newMultiShotTemplate(sim *core.Simulation) core.MeleeAbilityTemplate {
	ama := core.ActiveMeleeAbility{
		MeleeAbility: core.MeleeAbility{
			ActionID:    MultiShotActionID,
			Character:   &hunter.Character,
			SpellSchool: stats.AttackPower,
			GCD:         core.GCDDefault,
			Cooldown:    time.Second * 10,
			Cost: core.ResourceCost{
				Type:  stats.Mana,
				Value: 275,
			},
			// TODO: If we ever allow multiple targets to have their own type, need to
			// update this.
			CritMultiplier: hunter.critMultiplier(true, sim.GetPrimaryTarget()),
		},
	}

	ama.Cost.Value *= 1 - 0.02*float64(hunter.Talents.Efficiency)

	baseEffect := core.AbilityHitEffect{
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
					205
			},
		},
	}

	baseEffect.DamageMultiplier *= 1 + 0.04*float64(hunter.Talents.Barrage)
	baseEffect.BonusCritRating += float64(hunter.Talents.ImprovedBarrage) * 4 * core.MeleeCritRatingPerCritChance

	numHits := core.MinInt32(3, sim.GetNumTargets())
	effects := make([]core.AbilityHitEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)
	}
	ama.Effects = effects

	return core.NewMeleeAbilityTemplate(ama)
}

func (hunter *Hunter) NewMultiShot(sim *core.Simulation) *core.ActiveMeleeAbility {
	ms := &hunter.multiShot
	hunter.multiShotTemplate.Apply(ms)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	// Nothing

	return ms
}
