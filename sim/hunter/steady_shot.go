package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var SteadyShotActionID = core.ActionID{SpellID: 34120}

// ActiveMeleeAbility doesn't support cast times, so we wrap it in a SimpleCast.
func (hunter *Hunter) newSteadyShotCastTemplate(sim *core.Simulation) core.SimpleCast {
	template := core.SimpleCast{
		Cast: core.Cast{
			ActionID:  SteadyShotActionID,
			Character: hunter.GetCharacter(),
			BaseCost: core.ResourceCost{
				Type:  stats.Mana,
				Value: 110,
			},
			Cost: core.ResourceCost{
				Type:  stats.Mana,
				Value: 110,
			},
			// Cast time is affected by ranged attack speed so set it later.
			//CastTime:     time.Millisecond * 1500,
			GCD:         core.GCDDefault,
			IgnoreHaste: true, // Hunter GCD is locked at 1.5s
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				target := sim.GetPrimaryTarget()
				ss := &hunter.steadyShotAbility
				hunter.steadyShotAbilityTemplate.Apply(ss)
				ss.Effect.Target = target
				ss.Attack(sim)

				hunter.killCommandBlocked = false
				hunter.TryKillCommand(sim, target)

				hunter.rotation(sim, false)
			},
		},
		DisableMetrics: true,
	}

	template.Cost.Value *= 1 - 0.02*float64(hunter.Talents.Efficiency)

	return template
}

func (hunter *Hunter) newSteadyShotAbilityTemplate(sim *core.Simulation) core.MeleeAbilityTemplate {
	ama := core.ActiveMeleeAbility{
		MeleeAbility: core.MeleeAbility{
			ActionID:       SteadyShotActionID,
			Character:      &hunter.Character,
			SpellSchool:    stats.AttackPower,
			CritMultiplier: hunter.critMultiplier(true, sim.GetPrimaryTarget()),
		},
		Effect: core.AbilityHitEffect{
			AbilityEffect: core.AbilityEffect{
				ProcMask:               core.ProcMaskRangedSpecial,
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
			},
			WeaponInput: core.WeaponDamageInput{
				CalculateDamage: func(attackPower float64, bonusWeaponDamage float64) float64 {
					return attackPower*0.2 +
						hunter.AutoAttacks.Ranged.BaseDamage(sim)*2.8/hunter.AutoAttacks.Ranged.SwingSpeed +
						150
				},
			},
		},
	}

	if ItemSetRiftStalker.CharacterHasSetBonus(&hunter.Character, 4) {
		ama.Effect.BonusCritRating += 5 * core.MeleeCritRatingPerCritChance
	}
	if ItemSetGronnstalker.CharacterHasSetBonus(&hunter.Character, 4) {
		ama.Effect.DamageMultiplier *= 1.1
	}

	return core.NewMeleeAbilityTemplate(ama)
}

func (hunter *Hunter) NewSteadyShot(sim *core.Simulation, target *core.Target) core.SimpleCast {
	hunter.steadyShotCast = hunter.steadyShotCastTemplate

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	hunter.steadyShotCast.CastTime = hunter.SteadyShotCastTime()

	hunter.steadyShotCast.Init(sim)
	return hunter.steadyShotCast
}

func (hunter *Hunter) SteadyShotCastTime() time.Duration {
	return time.Duration(float64(time.Millisecond*1500)/hunter.RangedSwingSpeed()) + hunter.latency
}
