package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var MultiShotCooldownID = core.NewCooldownID()
var MultiShotActionID = core.ActionID{SpellID: 27021, CooldownID: MultiShotCooldownID}

// ActiveMeleeAbility doesn't support cast times, so we wrap it in a SimpleCast.
func (hunter *Hunter) newMultiShotCastTemplate(sim *core.Simulation) core.SimpleCast {
	template := core.SimpleCast{
		Cast: core.Cast{
			ActionID:  MultiShotActionID,
			Character: hunter.GetCharacter(),
			BaseCost: core.ResourceCost{
				Type:  stats.Mana,
				Value: 275,
			},
			Cost: core.ResourceCost{
				Type:  stats.Mana,
				Value: 275,
			},
			// Cast time is affected by ranged attack speed so set it later.
			//CastTime:     time.Millisecond * 500,
			GCD:         core.GCDDefault,
			Cooldown:    time.Second * 10,
			IgnoreHaste: true, // Hunter GCD is locked at 1.5s
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				ms := &hunter.multiShotAbility
				hunter.multiShotAbilityTemplate.Apply(ms)
				ms.Attack(sim)
				hunter.rotation(sim, false)
			},
		},
		DisableMetrics: true,
	}

	template.Cost.Value *= 1 - 0.02*float64(hunter.Talents.Efficiency)
	if ItemSetDemonStalker.CharacterHasSetBonus(&hunter.Character, 4) {
		template.Cost.Value -= 275.0 * 0.1
	}

	return template
}

func (hunter *Hunter) newMultiShotAbilityTemplate(sim *core.Simulation) core.MeleeAbilityTemplate {
	ama := core.ActiveMeleeAbility{
		Cast: core.Cast{
			ActionID:            MultiShotActionID,
			Character:           &hunter.Character,
			OutcomeRollCategory: core.OutcomeRollCategoryRanged,
			CritRollCategory:    core.CritRollCategoryPhysical,
			SpellSchool:         core.SpellSchoolPhysical,
			// TODO: If we ever allow multiple targets to have their own type, need to
			// update this.
			CritMultiplier: hunter.critMultiplier(true, sim.GetPrimaryTarget()),
		},
	}

	baseEffect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			ProcMask:               core.ProcMaskRangedSpecial,
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
			ThreatMultiplier:       1,
		},
		WeaponInput: core.WeaponDamageInput{
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
	effects := make([]core.SpellHitEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)
	}
	ama.Effects = effects

	return core.NewMeleeAbilityTemplate(ama)
}

func (hunter *Hunter) NewMultiShot(sim *core.Simulation) core.SimpleCast {
	hunter.multiShotCast = hunter.multiShotCastTemplate

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	hunter.multiShotCast.CastTime = hunter.MultiShotCastTime()

	hunter.multiShotCast.Init(sim)
	return hunter.multiShotCast
}

func (hunter *Hunter) MultiShotCastTime() time.Duration {
	return time.Duration(float64(time.Millisecond*500)/hunter.RangedSwingSpeed()) + hunter.latency
}
