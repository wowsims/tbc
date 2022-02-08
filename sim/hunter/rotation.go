package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

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

func (hunter *Hunter) OnAutoAttack(sim *core.Simulation, ability *core.ActiveMeleeAbility) {
	hitEffect := &ability.Effect
	if !hitEffect.IsRanged() || hunter.IsOnCD(core.GCDCooldownID, sim.CurrentTime) {
		return
	}

	hunter.rotation(sim, true)
}

func (hunter *Hunter) OnGCDReady(sim *core.Simulation) {
	// Hunters do most things between auto shots, so GCD usage is handled within OnAutoAttack (see above).
	// Only use this for follow-up actions after an auto+GCD, i.e. melee weave or French rotation.
	if sim.CurrentTime == 0 {
		if hunter.Rotation.PrecastAimedShot && hunter.Talents.AimedShot {
			hunter.NewAimedShot(sim, sim.GetPrimaryTarget()).Attack(sim)
		}

		// Dont do anything fancy on the first GCD, just wait for first auto.
		return
	}

	hunter.rotation(sim, false)
}

func (hunter *Hunter) rotation(sim *core.Simulation, followsAuto bool) {
	if hunter.Rotation.LazyRotation {
		hunter.lazyRotation(sim, followsAuto)
	} else {
		hunter.adaptiveRotation(sim, followsAuto)
	}
}

func (hunter *Hunter) lazyRotation(sim *core.Simulation, followsAuto bool) {
	if !followsAuto {
		return
	}

	// First prio is swapping aspect if necessary.
	currentMana := hunter.CurrentManaPercent()
	if hunter.aspectOfTheViper && currentMana > hunter.Rotation.ViperStopManaPercent {
		aspect := hunter.NewAspectOfTheHawk(sim)
		aspect.StartCast(sim)
		return
	} else if !hunter.aspectOfTheViper && currentMana < hunter.Rotation.ViperStartManaPercent {
		aspect := hunter.NewAspectOfTheViper(sim)
		aspect.StartCast(sim)
		return
	} else if hunter.tryUseInstantCast(sim) {
		return
	}

	canMulti := hunter.Rotation.UseMultiShot && !hunter.IsOnCD(MultiShotCooldownID, sim.CurrentTime)
	if canMulti {
		// Replace SS with MS because there's no way we could take advantage of MS otherwise.
		ms := hunter.NewMultiShot(sim)
		if success := ms.StartCast(sim); !success {
			hunter.WaitForMana(sim, ms.GetManaCost())
		}
		return
	}

	target := sim.GetPrimaryTarget()
	canWeave := hunter.Rotation.MeleeWeave &&
		sim.GetRemainingDurationPercent() < hunter.Rotation.PercentWeaved &&
		hunter.AutoAttacks.MeleeSwingsReady(sim)
	cast := hunter.NewSteadyShot(sim, target, canWeave)
	if success := cast.StartCast(sim); !success {
		hunter.WaitForMana(sim, cast.GetManaCost())
	} else {
		// Can't use kill command while casting steady shot.
		hunter.killCommandBlocked = true
	}
}

func (hunter *Hunter) adaptiveRotation(sim *core.Simulation, followsAuto bool) {
	timeBeforeClip := hunter.AutoAttacks.TimeBeforeClippingRanged(sim)
	if timeBeforeClip < 0 {
		return
	}

	// First prio is swapping aspect if necessary.
	currentMana := hunter.CurrentManaPercent()
	if hunter.aspectOfTheViper && currentMana > hunter.Rotation.ViperStopManaPercent {
		aspect := hunter.NewAspectOfTheHawk(sim)
		aspect.StartCast(sim)
		return
	} else if !hunter.aspectOfTheViper && currentMana < hunter.Rotation.ViperStartManaPercent {
		aspect := hunter.NewAspectOfTheViper(sim)
		aspect.StartCast(sim)
		return
	}

	target := sim.GetPrimaryTarget()

	rangedSwingSpeed := hunter.RangedSwingSpeed()
	multiShotCastTime := time.Duration(float64(time.Millisecond*500) / rangedSwingSpeed)
	canWeave := hunter.Rotation.MeleeWeave &&
		sim.GetRemainingDurationPercent() < hunter.Rotation.PercentWeaved &&
		hunter.AutoAttacks.MeleeSwingsReady(sim)
	canWeaveNow := canWeave && hunter.timeToWeave < timeBeforeClip

	if timeBeforeClip < multiShotCastTime {
		// Not enough time for anything other than an instant-cast ability.
		hunter.tryUseInstantCast(sim)

		// Using an instant cast doesn't interfere with weaving.
		if canWeaveNow {
			hunter.doMeleeWeave(sim)
		}
		return
	}

	canMulti := hunter.Rotation.UseMultiShot && !hunter.IsOnCD(MultiShotCooldownID, sim.CurrentTime)
	steadyShotCastTime := time.Duration(float64(time.Second*1) / rangedSwingSpeed)
	if timeBeforeClip < steadyShotCastTime {
		if canMulti {
			ms := hunter.NewMultiShot(sim)
			if success := ms.StartCast(sim); !success {
				hunter.WaitForMana(sim, ms.GetManaCost())
				if canWeaveNow {
					hunter.doMeleeWeave(sim)
				}
			}
		} else {
			hunter.tryUseInstantCast(sim)
			if canWeaveNow {
				hunter.doMeleeWeave(sim)
			}
		}
		return
	}

	if canMulti && core.GCDDefault+multiShotCastTime > hunter.AutoAttacks.Ranged.SwingDuration {
		// Replace SS with MS because there's no way we could take advantage of MS otherwise.
		ms := hunter.NewMultiShot(sim)
		if success := ms.StartCast(sim); !success {
			hunter.WaitForMana(sim, ms.GetManaCost())
		}
	}

	cast := hunter.NewSteadyShot(sim, target, canWeave)
	if success := cast.StartCast(sim); !success {
		hunter.WaitForMana(sim, cast.GetManaCost())
	} else {
		// Can't use kill command while casting steady shot.
		hunter.killCommandBlocked = true
	}
}

// Decides whether to use an instant-cast GCD spell.
// Returns true if any of these spells was selected.
func (hunter *Hunter) tryUseInstantCast(sim *core.Simulation) bool {
	target := sim.GetPrimaryTarget()

	if hunter.Rotation.Sting == proto.Hunter_Rotation_ScorpidSting && !target.HasAura(ScorpidStingDebuffID) {
		ss := hunter.NewScorpidSting(sim, target)
		if success := ss.Attack(sim); success {
			// Applies clipping if necessary.
			hunter.AutoAttacks.DelayRangedUntil(sim, sim.CurrentTime)
		} else {
			hunter.WaitForMana(sim, ss.Cost.Value)
		}
		return true
	} else if hunter.Rotation.Sting == proto.Hunter_Rotation_SerpentSting && !hunter.serpentStingDot.IsInUse() {
		ss := hunter.NewSerpentSting(sim, target)
		if success := ss.Attack(sim); success {
			// Applies clipping if necessary.
			hunter.AutoAttacks.DelayRangedUntil(sim, sim.CurrentTime)
		} else {
			hunter.WaitForMana(sim, ss.Cost.Value)
		}
		return true
	} else if hunter.Rotation.UseArcaneShot && !hunter.IsOnCD(ArcaneShotCooldownID, sim.CurrentTime) {
		as := hunter.NewArcaneShot(sim, target)
		if success := as.Attack(sim); success {
			// Applies clipping if necessary.
			hunter.AutoAttacks.DelayRangedUntil(sim, sim.CurrentTime)
		} else {
			hunter.WaitForMana(sim, as.Cost.Value)
		}
		return true
	}
	return false
}

func (hunter *Hunter) doMeleeWeave(sim *core.Simulation) {
	hunter.AutoAttacks.SwingMelee(sim, sim.GetPrimaryTarget())

	// Delay ranged autos until the weaving is done.
	hunter.AutoAttacks.DelayRangedUntil(sim, sim.CurrentTime+hunter.timeToWeave)
}
