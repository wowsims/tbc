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
			ActionID:     SteadyShotActionID,
			Character:    hunter.GetCharacter(),
			BaseManaCost: 110,
			ManaCost:     110,
			// Cast time is affected by ranged attack speed so set it later.
			//CastTime:     time.Second * 1,
			GCD:         core.GCDDefault,
			IgnoreHaste: true, // Hunter GCD is locked at 1.5s
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				hunter.killCommandBlocked = false
				hunter.TryKillCommand(sim, sim.GetPrimaryTarget())
			},
		},
		DisableMetrics: true,
	}

	template.ManaCost *= 1 - 0.02*float64(hunter.Talents.Efficiency)

	return template
}

func (hunter *Hunter) newSteadyShotAbilityTemplate(sim *core.Simulation) core.MeleeAbilityTemplate {
	ama := core.ActiveMeleeAbility{
		MeleeAbility: core.MeleeAbility{
			ActionID:       SteadyShotActionID,
			Character:      &hunter.Character,
			SpellSchool:    stats.AttackPower,
			IgnoreCost:     true,
			CritMultiplier: hunter.critMultiplier(true, sim.GetPrimaryTarget()),
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

func (hunter *Hunter) NewSteadyShot(sim *core.Simulation, target *core.Target, canWeave bool) core.SimpleCast {
	hunter.steadyShotCast = hunter.steadyShotCastTemplate

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	hunter.steadyShotCast.CastTime = time.Duration(float64(time.Second*1) / hunter.RangedSwingSpeed())

	// Might be able to fill the gap between SS complete and GCD ready with a melee weave.
	leftoverGCDTime := core.GCDDefault - hunter.steadyShotCast.CastTime
	wouldClipAuto := hunter.steadyShotCast.CastTime+hunter.timeToWeave > hunter.AutoAttacks.TimeBeforeClippingRanged(sim)
	canWeaveAfterSS := canWeave &&
		hunter.timeToWeave < leftoverGCDTime &&
		!wouldClipAuto

	hunter.steadyShotCast.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		ss := &hunter.steadyShotAbility
		hunter.steadyShotAbilityTemplate.Apply(ss)
		ss.Effect.Target = target
		ss.Attack(sim)

		if canWeaveAfterSS {
			hunter.doMeleeWeave(sim)
		}
	}

	hunter.steadyShotCast.Init(sim)
	return hunter.steadyShotCast
}
