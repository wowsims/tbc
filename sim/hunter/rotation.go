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
	}
}

func (hunter *Hunter) OnAutoAttack(sim *core.Simulation, ability *core.ActiveMeleeAbility) {
	hunter.rotation(sim, ability.Effect.IsRanged())
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

	// Swap aspects or keep up sting if needed.
	hunter.tryUsePrioGCD(sim)

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
	} else if hunter.nextActionAt != hunter.NextGCDAt() {
		hunter.WaitUntil(sim, hunter.nextActionAt)
	}
}

func (hunter *Hunter) lazyRotation(sim *core.Simulation, followsRangedAuto bool) {
	if hunter.AutoAttacks.RangedSwingAt <= sim.CurrentTime {
		hunter.nextAction = OptionShoot
		hunter.nextActionAt = sim.CurrentTime
		return
	}

	if !followsRangedAuto {
		hunter.nextAction = OptionShoot
		hunter.nextActionAt = hunter.AutoAttacks.RangedSwingAt
		return
	}

	canMulti := hunter.Rotation.UseMultiShot && !hunter.IsOnCD(MultiShotCooldownID, sim.CurrentTime)
	if canMulti {
		hunter.nextAction = OptionMulti
		hunter.nextActionAt = sim.CurrentTime
		return
	}

	canArcane := hunter.Rotation.UseArcaneShot && !hunter.IsOnCD(ArcaneShotCooldownID, sim.CurrentTime)
	if canArcane {
		hunter.nextAction = OptionArcane
		hunter.nextActionAt = sim.CurrentTime
		return
	}

	canWeave := hunter.Rotation.MeleeWeave &&
		sim.GetRemainingDurationPercent() < hunter.Rotation.PercentWeaved &&
		hunter.AutoAttacks.MeleeSwingsReady(sim)
	if canWeave {
		hunter.nextAction = OptionWeave
		hunter.nextActionAt = sim.CurrentTime
		return
	}

	canSteady := !hunter.IsOnCD(core.GCDCooldownID, sim.CurrentTime)
	if canSteady {
		hunter.nextAction = OptionSteady
		hunter.nextActionAt = sim.CurrentTime
		return
	}

	// This case probably isn't possible.
	hunter.nextAction = OptionShoot
	hunter.nextActionAt = hunter.AutoAttacks.RangedSwingAt
}

func (hunter *Hunter) adaptiveRotation(sim *core.Simulation, followsRangedAuto bool) {
	gcdAtDuration := core.MaxDuration(sim.CurrentTime, hunter.NextGCDAt())
	weaveAtDuration := core.MaxDuration(sim.CurrentTime, hunter.AutoAttacks.MeleeSwingsReadyAt())
	shootAtDuration := core.MaxDuration(sim.CurrentTime, hunter.AutoAttacks.RangedSwingAt)
	gcdAt := gcdAtDuration.Seconds()
	weaveAt := weaveAtDuration.Seconds()
	shootAt := shootAtDuration.Seconds()

	// Use the inverse (1 / x) because multiplication is faster than division.
	gcdRate := 1.0 / 1.5
	weaveRate := 1.0 / core.MaxDuration(hunter.AutoAttacks.MainhandSwingSpeed(), hunter.AutoAttacks.OffhandSwingSpeed()).Seconds()
	shootRate := 1.0 / hunter.AutoAttacks.RangedSwingSpeed().Seconds()

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
	rangedWindup := hunter.AutoAttacks.RangedSwingWindup()
	rangedWindupSeconds := rangedWindup.Seconds()
	shootGCDDelay := core.MaxFloat(0, (shootAt+rangedWindupSeconds)-gcdAt)
	shootWeaveDelay := core.MaxFloat(0, (shootAt+rangedWindupSeconds)-weaveAt)
	dmgResults[OptionShoot] = hunter.avgShootDmg -
		(hunter.avgSteadyDmg * gcdRate * shootGCDDelay) -
		(hunter.avgWeaveDmg * weaveRate * shootWeaveDelay)

	// Dmg from choosing to weave next.
	canWeave := hunter.Rotation.MeleeWeave && sim.GetRemainingDurationPercent() < hunter.Rotation.PercentWeaved
	if canWeave {
		weaveCastTime := hunter.timeToWeave.Seconds()
		weaveShootDelay := core.MaxFloat(0, (weaveAt+weaveCastTime)-shootAt)
		weaveGCDDelay := core.MaxFloat(0, (weaveAt+weaveCastTime)-gcdAt)
		dmgResults[OptionWeave] = hunter.avgWeaveDmg -
			(hunter.avgSteadyDmg * gcdRate * weaveGCDDelay) -
			(hunter.avgShootDmg * shootRate * weaveShootDelay)
	}

	// Dmg from choosing Steady Shot next.
	rangedSwingSpeed := hunter.RangedSwingSpeed()
	steadyShotCastTime := 1.5 / rangedSwingSpeed
	steadyShootDelay := core.MaxFloat(0, (gcdAt+steadyShotCastTime)-shootAt)
	steadyWeaveDelay := core.MaxFloat(0, (gcdAt+steadyShotCastTime)-weaveAt)
	dmgResults[OptionSteady] = hunter.avgSteadyDmg -
		(hunter.avgWeaveDmg * weaveRate * steadyWeaveDelay) -
		(hunter.avgShootDmg * shootRate * steadyShootDelay)

	// Dmg from choosing Multi Shot next.
	canMulti := hunter.Rotation.UseMultiShot && hunter.CDReadyAt(MultiShotCooldownID) <= hunter.NextGCDAt()
	if canMulti {
		// https://diziet559.github.io/rotationtools/#rotation-details
		// When off CD, multi always has higher DPS than SS. Sometimes we want to
		// save it for later though, if we need to take advantage of its lower cast
		// time.
		rangedGapTime := hunter.AutoAttacks.RangedSwingSpeed() - rangedWindup

		autoCycleDuration := rangedGapTime
		for autoCycleDuration < core.GCDDefault {
			autoCycleDuration += rangedGapTime + rangedWindup
		}
		leftoverGCDRatio := float64(autoCycleDuration-core.GCDDefault) / float64(rangedGapTime+rangedWindup)
		canUseFollowingAuto := leftoverGCDRatio < 0.95
		msWouldFollowAuto := followsRangedAuto && gcdAtDuration <= sim.CurrentTime

		// If ranged swing speed lines up closely with GCD without any clipping, then
		// its never worth saving MS to use for the lower cast time.
		if canUseFollowingAuto || !msWouldFollowAuto {
			multiShotCastTime := 0.5 / rangedSwingSpeed
			multiShootDelay := core.MaxFloat(0, (gcdAt+multiShotCastTime)-shootAt)
			multiWeaveDelay := core.MaxFloat(0, (gcdAt+multiShotCastTime)-weaveAt)
			dmgResults[OptionMulti] = hunter.avgMultiDmg -
				(hunter.avgWeaveDmg * weaveRate * multiWeaveDelay) -
				(hunter.avgShootDmg * shootRate * multiShootDelay)
		}
	}

	// Dmg from choosing Arcane Shot next.
	canArcane := hunter.Rotation.UseArcaneShot && hunter.CDReadyAt(ArcaneShotCooldownID) <= hunter.NextGCDAt()
	if canArcane {
		arcaneShootDelay := core.MaxFloat(0, gcdAt-shootAt)
		arcaneWeaveDelay := core.MaxFloat(0, gcdAt-weaveAt)

		// Since steady is higher damage than arcane we need to subtract that difference too.
		// The only times we'll ever use arcane shot is if another auto will follow
		// before the steady shot, so we don't subtract the full steady damage because
		// its only being partially delayed.
		oldSteadyTime := shootAt
		newSteadyTime := core.MaxFloat(gcdAt+1.5, shootAt+arcaneShootDelay)
		arcaneSteadyDelay := core.MaxFloat(0, newSteadyTime-oldSteadyTime)

		dmgResults[OptionArcane] = hunter.avgArcaneDmg -
			(hunter.avgWeaveDmg * weaveRate * arcaneWeaveDelay) -
			(hunter.avgShootDmg * shootRate * arcaneShootDelay) -
			(hunter.avgSteadyDmg * gcdRate * arcaneSteadyDelay)
	}

	actionAtResults := []time.Duration{
		shootAtDuration,
		weaveAtDuration,
		gcdAtDuration,
		gcdAtDuration,
		gcdAtDuration,
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

	hunter.nextAction = bestOption
	hunter.nextActionAt = bestOptionAt
}

// Decides whether to use an instant-cast GCD spell.
// Returns true if any of these spells was selected.
func (hunter *Hunter) tryUsePrioGCD(sim *core.Simulation) bool {
	// First prio is swapping aspect if necessary.
	currentMana := hunter.CurrentManaPercent()
	if hunter.aspectOfTheViper && currentMana > hunter.Rotation.ViperStopManaPercent {
		aspect := hunter.NewAspectOfTheHawk(sim)
		aspect.StartCast(sim)
		return true
	} else if !hunter.aspectOfTheViper && currentMana < hunter.Rotation.ViperStartManaPercent {
		aspect := hunter.NewAspectOfTheViper(sim)
		aspect.StartCast(sim)
		return true
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

func (hunter *Hunter) doOption(sim *core.Simulation, option int) {
	hunter.nextAction = OptionNone
	target := sim.GetPrimaryTarget()
	switch option {
	case OptionShoot:
		hunter.AutoAttacks.SwingRanged(sim, target)
	case OptionWeave:
		hunter.doMeleeWeave(sim)
	case OptionSteady:
		ss := hunter.NewSteadyShot(sim, target)
		if success := ss.StartCast(sim); success {
			// Can't use kill command while casting steady shot.
			hunter.killCommandBlocked = true
		} else {
			hunter.WaitForMana(sim, ss.GetManaCost())
		}
	case OptionMulti:
		ms := hunter.NewMultiShot(sim)
		if success := ms.StartCast(sim); success {
		} else {
			hunter.WaitForMana(sim, ms.GetManaCost())
		}
	case OptionArcane:
		as := hunter.NewArcaneShot(sim, target)
		if success := as.Attack(sim); success {
			// Arcane is instant, so we can try another action immediately.
			hunter.rotation(sim, false)
		} else {
			hunter.WaitForMana(sim, as.Cost.Value)
		}
	}
}

func (hunter *Hunter) doMeleeWeave(sim *core.Simulation) {
	hunter.AutoAttacks.SwingMelee(sim, sim.GetPrimaryTarget())

	// Delay ranged autos until the weaving is done.
	hunter.AutoAttacks.DelayRangedUntil(sim, sim.CurrentTime+hunter.timeToWeave)
}

func (hunter *Hunter) GetPresimOptions() *core.PresimOptions {
	// If not adaptive, don't need to run a presim.
	if hunter.Rotation.LazyRotation {
		return nil
	}

	return &core.PresimOptions{
		SetPresimPlayerOptions: func(player *proto.Player) {
			rotation := hunter.Rotation
			rotation.LazyRotation = true
			*player.Spec.(*proto.Player_Hunter).Hunter.Rotation = rotation
		},

		OnPresimResult: func(presimResult proto.PlayerMetrics, iterations int32, duration time.Duration) bool {
			hunter.avgShootDmg = core.GetActionAvgCast(presimResult, core.ActionID{OtherID: proto.OtherAction_OtherActionShoot})
			hunter.avgWeaveDmg = core.GetActionAvgCast(presimResult, RaptorStrikeActionID) +
				core.GetActionAvgCast(presimResult, core.ActionID{OtherID: proto.OtherAction_OtherActionAttack, Tag: 1}) +
				core.GetActionAvgCast(presimResult, core.ActionID{OtherID: proto.OtherAction_OtherActionAttack, Tag: 2})
			hunter.avgSteadyDmg = core.GetActionAvgCast(presimResult, SteadyShotActionID)
			hunter.avgMultiDmg = core.GetActionAvgCast(presimResult, MultiShotActionID)
			hunter.avgArcaneDmg = core.GetActionAvgCast(presimResult, ArcaneShotActionID)
			return true
		},
	}
}
