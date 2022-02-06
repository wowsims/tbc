package hunter

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (hunter *Hunter) OnAutoAttack(sim *core.Simulation, ability *core.ActiveMeleeAbility) {
	hitEffect := &ability.Effect
	if !hitEffect.IsRanged() || hunter.IsOnCD(core.GCDCooldownID, sim.CurrentTime) {
		return
	}
	hunter.tryUseGCD(sim)
}

func (hunter *Hunter) tryUseGCD(sim *core.Simulation) {
	if sim.CurrentTime == 0 && hunter.Rotation.PrecastAimedShot && hunter.Talents.AimedShot {
		hunter.NewAimedShot(sim, sim.GetPrimaryTarget()).Attack(sim)
		return
	}

	if hunter.IsWaitingForMana() {
		return
	}

	target := sim.GetPrimaryTarget()
	hasted := hunter.HasTemporaryRangedSwingSpeedIncrease()

	if hunter.Rotation.UseFrenchRotation && !hasted {
		// French rotation, i.e. special GCDs are used after a steady shot.
		cast := hunter.NewSteadyShot(sim, target)
		if success := cast.StartCast(sim); !success {
			hunter.WaitForMana(sim, cast.GetManaCost())
		}
	} else {
		// Regular rotation, i.e. special GCDs take the place of steady shot.
		if !hunter.tryUseSpecialGCD(sim) {
			cast := hunter.NewSteadyShot(sim, target)
			if success := cast.StartCast(sim); !success {
				hunter.WaitForMana(sim, cast.GetManaCost())
			}
		}
	}
}

// Decides whether to use a GCD spell other than Steady Shot.
// Returns true if any of these spells was selected.
func (hunter *Hunter) tryUseSpecialGCD(sim *core.Simulation) bool {
	target := sim.GetPrimaryTarget()
	currentMana := hunter.CurrentManaPercent()

	if hunter.aspectOfTheViper && currentMana > hunter.Rotation.ViperStopManaPercent {
		aspect := hunter.NewAspectOfTheHawk(sim)
		aspect.StartCast(sim)
		return true
	} else if !hunter.aspectOfTheViper && currentMana < hunter.Rotation.ViperStartManaPercent {
		aspect := hunter.NewAspectOfTheViper(sim)
		aspect.StartCast(sim)
		return true
	} else if hunter.Rotation.Sting == proto.Hunter_Rotation_ScorpidSting && !target.HasAura(ScorpidStingDebuffID) {
		ss := hunter.NewScorpidSting(sim, target)
		if success := ss.Attack(sim); !success {
			hunter.WaitForMana(sim, ss.Cost.Value)
		}
		return true
	} else if hunter.Rotation.Sting == proto.Hunter_Rotation_SerpentSting && !hunter.serpentStingDot.IsInUse() {
		ss := hunter.NewSerpentSting(sim, target)
		if success := ss.Attack(sim); !success {
			hunter.WaitForMana(sim, ss.Cost.Value)
		}
		return true
	} else if hunter.Rotation.UseMultiShot && !hunter.IsOnCD(MultiShotCooldownID, sim.CurrentTime) {
		ms := hunter.NewMultiShot(sim)
		if success := ms.StartCast(sim); !success {
			hunter.WaitForMana(sim, ms.GetManaCost())
		}
		return true
	} else if hunter.Rotation.UseArcaneShot && !hunter.IsOnCD(ArcaneShotCooldownID, sim.CurrentTime) {
		as := hunter.NewArcaneShot(sim, target)
		if success := as.Attack(sim); !success {
			hunter.WaitForMana(sim, as.Cost.Value)
		}
		return true
	}
	return false
}

func (hunter *Hunter) OnGCDReady(sim *core.Simulation) {
	// Hunters do most things between auto shots, so GCD usage is handled as an aura (see above).
	// Only use this for follow-up actions after an auto+GCD, i.e. melee weave or French rotation.
	if sim.Log != nil {
		sim.Log("hunter GCD")
	}

	if sim.CurrentTime == 0 {
		// Dont do anything fancy on the first GCD.
		return
	}

	hasted := hunter.HasTemporaryRangedSwingSpeedIncrease()
	if hunter.Rotation.UseFrenchRotation && !hasted {
		// 2nd GCD cast in the French rotation.
		hunter.tryUseSpecialGCD(sim)
	} else if hunter.Rotation.MeleeWeave && sim.GetRemainingDurationPercent() < hunter.Rotation.PercentWeaved && hunter.AutoAttacks.MeleeSwingsReady(sim) {
		// Melee weave.
		hunter.AutoAttacks.SwingMelee(sim, sim.GetPrimaryTarget())

		// Delay ranged autos until the weaving is done.
		hunter.AutoAttacks.DelayRangedUntil(sim, sim.CurrentTime+hunter.timeToWeave)
	}
}

func (hunter *Hunter) OnManaTick(sim *core.Simulation) {
	if hunter.aspectOfTheViper {
		// https://wowpedia.fandom.com/wiki/Aspect_of_the_Viper?oldid=1458832
		percentMana := core.MaxFloat(0.2, core.MinFloat(0.9, hunter.CurrentManaPercent()))
		scaling := 22.0/35.0*(0.9-percentMana) + 0.11
		if hunter.hasGronnstalker2Pc {
			scaling += 0.05
		}

		bonusPer5Seconds := hunter.GetStat(stats.Intellect)*scaling + 0.35*70
		manaGain := bonusPer5Seconds * 2 / 5
		hunter.AddMana(sim, manaGain, AspectOfTheViperActionID, false)
	}

	if hunter.IsWaitingForMana() && hunter.DoneWaitingForMana(sim) {
		hunter.TryKillCommand(sim, sim.GetPrimaryTarget())
	}
}
