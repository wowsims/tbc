package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

const (
	OptionShoot = iota
	OptionWeave
	OptionSteady
	OptionMulti
	OptionArcane
	OptionNone
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
		if hunter.nextAction == OptionNone && hunter.Hardcast.Expires <= sim.CurrentTime {
			hunter.rotation(sim, false)
		}
	}
}

func (hunter *Hunter) OnAutoAttack(sim *core.Simulation, ability *core.SimpleSpell) {
	hunter.TryKillCommand(sim, sim.GetPrimaryTarget())
	if !ability.OutcomeRollCategory.Matches(core.OutcomeRollCategoryRanged) {
		return
	}
	hunter.rotation(sim, true)
}

func (hunter *Hunter) OnGCDReady(sim *core.Simulation) {
	if sim.CurrentTime == 0 {
		if hunter.Rotation.PrecastAimedShot && hunter.Talents.AimedShot {
			hunter.NewAimedShot(sim, sim.GetPrimaryTarget()).Attack(sim)
		}
		hunter.AutoAttacks.SwingRanged(sim, sim.GetPrimaryTarget())
		return
	}

	if hunter.AutoAttacks.RangedSwingInProgress {
		return
	}

	hunter.TryKillCommand(sim, sim.GetPrimaryTarget())

	hunter.rotation(sim, false)
}

func (hunter *Hunter) rotation(sim *core.Simulation, followsRangedAuto bool) {
	if hunter.nextAction == OptionNone {
		if hunter.Rotation.LazyRotation {
			hunter.lazyRotation(sim, followsRangedAuto)
		} else {
			hunter.adaptiveRotation(sim, followsRangedAuto)
		}
	}

	if hunter.nextActionAt <= sim.CurrentTime {
		hunter.doOption(sim, hunter.nextAction)
		if hunter.IsWaitingForMana() && hunter.nextAction != OptionShoot && hunter.nextAction != OptionWeave {
			if hunter.Hardcast.Expires <= sim.CurrentTime {
				hunter.nextAction = OptionShoot
				hunter.nextActionAt = hunter.AutoAttacks.RangedSwingAt
				hunter.HardcastWaitUntil(sim, hunter.nextActionAt, &hunter.fakeHardcast)
			}
		}
	} else if hunter.nextActionAt != hunter.NextGCDAt() {
		if hunter.Hardcast.Expires <= sim.CurrentTime {
			hunter.HardcastWaitUntil(sim, hunter.nextActionAt, &hunter.fakeHardcast)
		}
	}
}

func (hunter *Hunter) lazyRotation(sim *core.Simulation, followsRangedAuto bool) {
	shootAt := hunter.AutoAttacks.RangedSwingAt
	shootReady := shootAt <= sim.CurrentTime
	gcdAt := hunter.NextGCDAt()
	gcdReady := gcdAt <= sim.CurrentTime

	waitingForMana := hunter.IsWaitingForMana()
	canWeave := hunter.Rotation.Weave != proto.Hunter_Rotation_WeaveNone &&
		(hunter.Rotation.Weave != proto.Hunter_Rotation_WeaveRaptorOnly || !hunter.IsOnCD(RaptorStrikeCooldownID, sim.CurrentTime)) &&
		sim.CurrentTime >= hunter.weaveStartTime &&
		hunter.AutoAttacks.MainhandSwingAt <= sim.CurrentTime
	if canWeave && !shootReady && (!gcdReady || (waitingForMana && hunter.Rotation.Weave != proto.Hunter_Rotation_WeaveRaptorOnly)) {
		hunter.nextAction = OptionWeave
		hunter.nextActionAt = sim.CurrentTime
		return
	}

	if shootAt <= gcdAt || waitingForMana {
		hunter.nextAction = OptionShoot
		hunter.nextActionAt = shootAt
		return
	}

	canMulti := hunter.Rotation.UseMultiShot && !hunter.IsOnCD(MultiShotCooldownID, sim.CurrentTime)
	if canMulti {
		hunter.nextAction = OptionMulti
		hunter.nextActionAt = gcdAt
		return
	}

	steadyShotCastTime := time.Duration(float64(time.Millisecond*1500) / hunter.RangedSwingSpeed())
	ssWouldClip := gcdAt+steadyShotCastTime > shootAt

	canArcane := hunter.Rotation.UseArcaneShot && !hunter.IsOnCD(ArcaneShotCooldownID, sim.CurrentTime)
	if canArcane && ssWouldClip {
		hunter.nextAction = OptionArcane
		hunter.nextActionAt = gcdAt + hunter.latency
		return
	}

	hunter.nextAction = OptionSteady
	hunter.nextActionAt = gcdAt
}

func (hunter *Hunter) adaptiveRotation(sim *core.Simulation, followsRangedAuto bool) {
	gcdAtDuration := core.MaxDuration(sim.CurrentTime, hunter.NextGCDAt())
	shootAtDuration := core.MaxDuration(sim.CurrentTime, hunter.AutoAttacks.RangedSwingAt)
	weaveAtDuration := core.MaxDuration(sim.CurrentTime, hunter.AutoAttacks.MainhandSwingAt)
	if hunter.Rotation.Weave == proto.Hunter_Rotation_WeaveRaptorOnly {
		weaveAtDuration = core.MaxDuration(weaveAtDuration, hunter.CDReadyAt(RaptorStrikeCooldownID))
	}

	gcdAt := gcdAtDuration.Seconds()
	shootAt := shootAtDuration.Seconds()
	weaveAt := weaveAtDuration.Seconds()

	rangedSwingSpeed := hunter.RangedSwingSpeed()
	if rangedSwingSpeed != hunter.rangedSwingSpeed {
		// A lot of the calculations only need to be done when ranged speed changes.
		hunter.rangedSwingSpeed = rangedSwingSpeed
		rangedWindupDuration := hunter.AutoAttacks.RangedSwingWindup()
		hunter.rangedWindup = rangedWindupDuration.Seconds()

		// Use the inverse (1 / x) because multiplication is faster than division.
		gcdRate := 1.0 / 1.5
		weaveRate := 1.0 / hunter.AutoAttacks.MainhandSwingSpeed().Seconds()
		shootRate := 1.0 / hunter.AutoAttacks.RangedSwingSpeed().Seconds()

		hunter.shootDPS = hunter.avgShootDmg * shootRate
		hunter.weaveDPS = hunter.avgWeaveDmg * weaveRate
		hunter.steadyDPS = hunter.avgSteadyDmg * gcdRate

		hunter.steadyShotCastTime = hunter.SteadyShotCastTime().Seconds()
		hunter.multiShotCastTime = hunter.MultiShotCastTime().Seconds()
		hunter.arcaneShotCastTime = hunter.latency.Seconds()

		// https://diziet559.github.io/rotationtools/#rotation-details
		// When off CD, multi always has higher DPS than SS. Sometimes we want to
		// save it for later though, if we need to take advantage of its lower cast
		// time.
		rangedGapTime := hunter.AutoAttacks.RangedSwingSpeed() - rangedWindupDuration
		autoCycleDuration := rangedGapTime
		for autoCycleDuration < core.GCDDefault {
			autoCycleDuration += rangedGapTime + rangedWindupDuration
		}
		leftoverGCDRatio := float64(autoCycleDuration-core.GCDDefault) / float64(rangedGapTime+rangedWindupDuration)
		hunter.useMultiForCatchup = leftoverGCDRatio < 0.95
	}

	// For each ability option, we calculate the expected damage as the avg damage
	// of that ability with damage lost from delaying other abilities subtracted.
	// Damage lost is calculated as (DPS * delay).
	dmgResults := []float64{
		-10000.0,
		-10000.0,
		-10000.0,
		-10000.0,
		-10000.0,
	}

	// DPS from choosing to auto next.
	shootDoneAt := shootAt + hunter.rangedWindup
	shootGCDDelay := core.MaxFloat(0, shootDoneAt-gcdAt)
	dmgResults[OptionShoot] = hunter.avgShootDmg - (hunter.steadyDPS * shootGCDDelay)

	waitingForMana := hunter.IsWaitingForMana()
	if !waitingForMana {
		// Dmg from choosing Steady Shot next.
		steadyShootDelay := core.MaxFloat(0, (gcdAt+hunter.steadyShotCastTime)-shootAt)
		dmgResults[OptionSteady] = hunter.avgSteadyDmg - (hunter.shootDPS * steadyShootDelay)

		// Dmg from choosing Multi Shot next.
		canMulti := hunter.Rotation.UseMultiShot && hunter.CDReadyAt(MultiShotCooldownID) <= hunter.NextGCDAt()
		if canMulti {
			multiShootDelay := core.MaxFloat(0, (gcdAt+hunter.multiShotCastTime)-shootAt)

			// If ranged swing speed lines up closely with GCD without any clipping, then
			// its never worth saving MS to use for the lower cast time.
			if !hunter.useMultiForCatchup || multiShootDelay < steadyShootDelay {
				dmgResults[OptionMulti] = hunter.avgMultiDmg - (hunter.shootDPS * multiShootDelay)
			}
		}

		// Dmg from choosing Arcane Shot next.
		canArcane := hunter.Rotation.UseArcaneShot && hunter.CDReadyAt(ArcaneShotCooldownID) <= hunter.NextGCDAt()
		if canArcane {
			arcaneShootDelay := core.MaxFloat(0, (gcdAt+hunter.arcaneShotCastTime)-shootAt)
			dmgResults[OptionArcane] = hunter.avgArcaneDmg - (hunter.shootDPS * arcaneShootDelay)
		}
	}

	// Only allow weaving if autos and GCD will both be on CD. Otherwise it will
	// get used even when it would cause delays to them.
	canWeave := hunter.Rotation.Weave != proto.Hunter_Rotation_WeaveNone &&
		sim.CurrentTime >= hunter.weaveStartTime &&
		weaveAt < shootAt &&
		(weaveAt < gcdAt || (waitingForMana && hunter.Rotation.Weave != proto.Hunter_Rotation_WeaveRaptorOnly))
	if canWeave {
		// Dmg from choosing to weave next.
		weaveCastTime := hunter.timeToWeave.Seconds()
		weaveShootDelay := core.MaxFloat(0, (weaveAt+weaveCastTime)-shootAt)
		weaveGCDDelay := core.MaxFloat(0, (weaveAt+weaveCastTime)-gcdAt)
		dmgResults[OptionWeave] = hunter.avgWeaveDmg -
			(hunter.steadyDPS * weaveGCDDelay) -
			(hunter.shootDPS * weaveShootDelay)

		shootWeaveDelay := core.MaxFloat(0, shootDoneAt-weaveAt)
		dmgResults[OptionShoot] -= hunter.weaveDPS * shootWeaveDelay

		steadyWeaveDelay := core.MaxFloat(0, (gcdAt+hunter.steadyShotCastTime)-weaveAt)
		dmgResults[OptionSteady] -= hunter.weaveDPS * steadyWeaveDelay

		multiWeaveDelay := core.MaxFloat(0, (gcdAt+hunter.multiShotCastTime)-weaveAt)
		dmgResults[OptionMulti] -= hunter.weaveDPS * multiWeaveDelay

		arcaneWeaveDelay := core.MaxFloat(0, gcdAt-weaveAt)
		dmgResults[OptionArcane] -= hunter.weaveDPS * arcaneWeaveDelay
	}

	actionAtResults := []time.Duration{
		shootAtDuration,
		weaveAtDuration,
		gcdAtDuration,
		gcdAtDuration,
		gcdAtDuration + hunter.latency,
	}

	bestOption := 0
	bestDmg := dmgResults[OptionShoot]
	bestOptionAt := actionAtResults[OptionShoot]
	for i := range dmgResults {
		if dmgResults[i] > bestDmg {
			bestOption = i
			bestDmg = dmgResults[i]
			bestOptionAt = actionAtResults[i]
		}
	}

	//if sim.Log != nil {
	//	hunter.Log(sim, "Choosing option: %d, %s, shoot: %0.01f, weave: %0.01f, ss: %0.01f, ms: %0.01f, as: %0.01f", bestOption, bestOptionAt, dmgResults[0], dmgResults[1], dmgResults[2], dmgResults[3], dmgResults[4])
	//}

	hunter.nextAction = bestOption
	hunter.nextActionAt = bestOptionAt
}

func (hunter *Hunter) doOption(sim *core.Simulation, option int) {
	hunter.nextAction = OptionNone
	target := sim.GetPrimaryTarget()
	switch option {
	case OptionShoot:
		hunter.AutoAttacks.SwingRanged(sim, target)
	case OptionWeave:
		hunter.doMeleeWeave(sim)
	case OptionSteady:
		if !hunter.tryUsePrioGCD(sim) {
			ss := hunter.NewSteadyShot(sim, target)
			if success := ss.StartCast(sim); success {
				// Can't use kill command while casting steady shot.
				hunter.killCommandBlocked = true
			} else {
				hunter.WaitForMana(sim, ss.GetManaCost())
			}
		}
	case OptionMulti:
		if !hunter.tryUsePrioGCD(sim) {
			ms := hunter.NewMultiShot(sim)
			if success := ms.StartCast(sim); success {
			} else {
				hunter.WaitForMana(sim, ms.GetManaCost())
			}
		}
	case OptionArcane:
		if !hunter.tryUsePrioGCD(sim) {
			as := hunter.NewArcaneShot(sim, target)
			if success := as.Attack(sim); success {
				// Arcane is instant, so we can try another action immediately.
				hunter.rotation(sim, false)
			} else {
				hunter.WaitForMana(sim, as.Cost.Value)
			}
		}
	}
}

// Decides whether to use an instant-cast GCD spell.
// Returns true if any of these spells was selected.
func (hunter *Hunter) tryUsePrioGCD(sim *core.Simulation) bool {
	// First prio is swapping aspect if necessary.
	currentMana := hunter.CurrentManaPercent()
	if hunter.aspectOfTheViper {
		if !hunter.permaHawk &&
			hunter.CurrentMana() > hunter.manaSpentPerSecondAtFirstAspectSwap*sim.GetRemainingDuration().Seconds() {
			hunter.permaHawk = true
		}
		if hunter.permaHawk || currentMana > hunter.Rotation.ViperStopManaPercent {
			aspect := hunter.NewAspectOfTheHawk(sim)
			aspect.StartCast(sim)
			return true
		}
	} else if !hunter.aspectOfTheViper && !hunter.permaHawk && currentMana < hunter.Rotation.ViperStartManaPercent {
		if hunter.manaSpentPerSecondAtFirstAspectSwap == 0 {
			hunter.manaSpentPerSecondAtFirstAspectSwap = (hunter.Metrics.ManaSpent - hunter.Metrics.ManaGained) / sim.CurrentTime.Seconds()
		}
		if !hunter.permaHawk &&
			hunter.CurrentMana() > hunter.manaSpentPerSecondAtFirstAspectSwap*sim.GetRemainingDuration().Seconds() {
			hunter.permaHawk = true
		} else {
			aspect := hunter.NewAspectOfTheViper(sim)
			aspect.StartCast(sim)
			return true
		}
	}

	target := sim.GetPrimaryTarget()

	if hunter.Rotation.Sting == proto.Hunter_Rotation_ScorpidSting && !target.HasAura(ScorpidStingDebuffID) {
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
	}
	return false
}

func (hunter *Hunter) doMeleeWeave(sim *core.Simulation) {
	// Delay gcd and ranged autos until the weaving is done.
	doneWeavingAt := sim.CurrentTime + hunter.timeToWeave
	hunter.AutoAttacks.DelayRangedUntil(sim, doneWeavingAt)
	if doneWeavingAt > hunter.NextGCDAt() {
		hunter.SetGCDTimer(sim, doneWeavingAt)
	}

	hunter.AutoAttacks.TrySwingMH(sim, sim.GetPrimaryTarget())
	hunter.HardcastWaitUntil(sim, doneWeavingAt, &hunter.fakeHardcast)
}

func (hunter *Hunter) GetPresimOptions() *core.PresimOptions {
	// If not adaptive, don't need to run a presim.
	if hunter.Rotation.LazyRotation {
		return nil
	}

	return &core.PresimOptions{
		SetPresimPlayerOptions: func(player *proto.Player) {
			player.Spec.(*proto.Player_Hunter).Hunter.Rotation.LazyRotation = true
			player.Spec.(*proto.Player_Hunter).Hunter.Options.RemoveRandomness = true
		},

		OnPresimResult: func(presimResult proto.PlayerMetrics, iterations int32, duration time.Duration) bool {
			hunter.avgShootDmg = core.GetActionAvgCast(presimResult, core.ActionID{OtherID: proto.OtherAction_OtherActionShoot})
			hunter.avgWeaveDmg = core.GetActionAvgCast(presimResult, RaptorStrikeActionID) +
				core.GetActionAvgCast(presimResult, core.ActionID{OtherID: proto.OtherAction_OtherActionAttack, Tag: 1})
			hunter.avgSteadyDmg = core.GetActionAvgCast(presimResult, SteadyShotActionID)
			hunter.avgMultiDmg = core.GetActionAvgCast(presimResult, MultiShotActionID)
			hunter.avgArcaneDmg = core.GetActionAvgCast(presimResult, ArcaneShotActionID)
			return true
		},
	}
}
