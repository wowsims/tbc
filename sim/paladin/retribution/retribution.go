package retribution

import (
	"sort"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
	"github.com/wowsims/tbc/sim/paladin"
)

// Do 1 less millisecond to solve for sim order of operation problems
// Buffs are removed before melee swing is processed
const twistWindow = 399 * time.Millisecond

func RegisterRetributionPaladin() {
	core.RegisterAgentFactory(
		proto.Player_RetributionPaladin{},
		proto.Spec_SpecRetributionPaladin,
		func(character core.Character, options proto.Player) core.Agent {
			return NewRetributionPaladin(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_RetributionPaladin) // I don't really understand this line
			if !ok {
				panic("Invalid spec value for Retribution Paladin!")
			}
			player.Spec = playerSpec
		},
	)
}

func NewRetributionPaladin(character core.Character, options proto.Player) *RetributionPaladin {
	retOptions := options.GetRetributionPaladin()

	ret := &RetributionPaladin{
		Paladin:             paladin.NewPaladin(character, *retOptions.Talents),
		Rotation:            *retOptions.Rotation,
		crusaderStrikeDelay: time.Duration(retOptions.Options.CrusaderStrikeDelayMs) * time.Millisecond,
		hasteLeeway:         time.Duration(retOptions.Options.HasteLeewayMs) * time.Millisecond,
		judgement:           retOptions.Options.Judgement,
	}

	// Convert DTPS option to bonus MP5
	spAtt := retOptions.Options.DamageTakenPerSecond * 5.0 / 10.0
	ret.AddStat(stats.MP5, spAtt)

	ret.EnableAutoAttacks(ret, core.AutoAttackOptions{
		MainHand:       ret.WeaponFromMainHand(ret.DefaultMeleeCritMultiplier()),
		AutoSwingMelee: true,
	})

	return ret
}

type RetributionPaladin struct {
	*paladin.Paladin

	openerCompleted bool

	hasteLeeway         time.Duration
	crusaderStrikeDelay time.Duration

	judgement proto.RetributionPaladin_Options_Judgement

	Rotation proto.RetributionPaladin_Rotation
}

func (ret *RetributionPaladin) GetPaladin() *paladin.Paladin {
	return ret.Paladin
}

func (ret *RetributionPaladin) Reset(sim *core.Simulation) {
	ret.Paladin.Reset(sim)

	switch ret.judgement {
	case proto.RetributionPaladin_Options_Wisdom:
		ret.UpdateSeal(sim, ret.SealOfWisdomAura)
	case proto.RetributionPaladin_Options_Crusader:
		ret.UpdateSeal(sim, ret.SealOfTheCrusaderAura)
	case proto.RetributionPaladin_Options_None:
		ret.UpdateSeal(sim, ret.SealOfCommandAura)
	}

	ret.AutoAttacks.CancelAutoSwing(sim)
	ret.openerCompleted = false
}

func (ret *RetributionPaladin) OnGCDReady(sim *core.Simulation) {
	ret.tryUseGCD(sim)
}

func (ret *RetributionPaladin) OnManaTick(sim *core.Simulation) {
	if ret.FinishedWaitingForManaAndGCDReady(sim) {
		ret.tryUseGCD(sim)
	}
}

func (ret *RetributionPaladin) tryUseGCD(sim *core.Simulation) {
	if !ret.openerCompleted {
		ret.openingRotation(sim)
		return
	}
	ret.ActRotation(sim)
}

func (ret *RetributionPaladin) openingRotation(sim *core.Simulation) {
	target := sim.GetPrimaryTarget()

	// Cast selected judgement to keep on the boss
	if !ret.IsOnCD(paladin.JudgementCD, sim.CurrentTime) &&
		ret.judgement != proto.RetributionPaladin_Options_None {
		var judge *core.Spell
		switch ret.judgement {
		case proto.RetributionPaladin_Options_Wisdom:
			judge = ret.JudgementOfWisdom
		case proto.RetributionPaladin_Options_Crusader:
			judge = ret.JudgementOfTheCrusader
		}
		if judge != nil {
			if success := judge.Cast(sim, target); !success {
				ret.WaitForMana(sim, judge.Instance.GetManaCost())
			}
		}
	}

	// Cast Seal of Command
	if !ret.HasAura(paladin.SealOfCommandAuraID) {
		soc := ret.NewSealOfCommand(sim)
		if success := soc.StartCast(sim); !success {
			ret.WaitForMana(sim, soc.GetManaCost())
		}
		return
	}

	// Cast Seal of Blood and enable attacks to twist
	if !ret.HasAura(paladin.SealOfBloodAuraID) {
		sob := ret.NewSealOfBlood(sim)
		if success := sob.StartCast(sim); !success {
			ret.WaitForMana(sim, sob.GetManaCost())
		}
		ret.AutoAttacks.EnableAutoSwing(sim)
		ret.openerCompleted = true
	}
}

func (ret *RetributionPaladin) ActRotation(sim *core.Simulation) {
	// Setup
	target := sim.GetPrimaryTarget()

	gcdCD := ret.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime)
	crusaderStrikeCD := ret.GetRemainingCD(paladin.CrusaderStrikeCD, sim.CurrentTime)
	nextCrusaderStrikeCD := ret.CDReadyAt(paladin.CrusaderStrikeCD)
	judgementCD := ret.GetRemainingCD(paladin.JudgementCD, sim.CurrentTime)

	nextSwingAt := ret.AutoAttacks.NextAttackAt()
	timeTilNextSwing := nextSwingAt - sim.CurrentTime

	spellGCD := ret.SpellGCD()

	possibleTwist := timeTilNextSwing > spellGCD+gcdCD
	willTwist := possibleTwist && (nextSwingAt+spellGCD <= nextCrusaderStrikeCD+ret.crusaderStrikeDelay)
	inTwistWindow := (sim.CurrentTime >= nextSwingAt-twistWindow) && (sim.CurrentTime < ret.AutoAttacks.NextAttackAt())
	latestTwistStart := nextSwingAt - spellGCD

	sobActive := ret.RemainingAuraDuration(sim, paladin.SealOfBloodAuraID) > 0
	socActive := ret.RemainingAuraDuration(sim, paladin.SealOfCommandAuraID) > 0

	// Use Judgement if we will twist
	if judgementCD == 0 && willTwist && sobActive {
		ret.JudgementOfBlood.Cast(sim, target)
	}

	// Judgement can affect active seals and CDs
	nextJudgementCD := ret.CDReadyAt(paladin.JudgementCD)

	if gcdCD == 0 {
		if socActive && inTwistWindow {
			// If Seal of Command is Active, complete the twist
			ret.NewSealOfBlood(sim).StartCast(sim)
		} else if crusaderStrikeCD == 0 && !willTwist &&
			(sobActive || spellGCD < timeTilNextSwing) {
			// Cast Crusader Strike if we won't swing naked and we aren't twisting
			ret.CrusaderStrike.Cast(sim, target)
		} else if willTwist && !socActive && (nextJudgementCD > latestTwistStart) {
			// Prep seal of command
			ret.NewSealOfCommand(sim).StartCast(sim)
		} else if !sobActive && !socActive && !willTwist {
			// If no seal is active, cast Seal of Blood
			ret.NewSealOfBlood(sim).StartCast(sim)
		}
	}

	// Determine when next action is available
	// Throw everything into an array then filter and sort compared to doing individual comparisons
	events := []time.Duration{
		nextSwingAt,
		nextSwingAt - twistWindow,
		ret.CDReadyAt(core.GCDCooldownID),
		ret.CDReadyAt(paladin.JudgementCD),
		ret.CDReadyAt(paladin.CrusaderStrikeCD),
	}

	// Time has to move forward... so exclude any events that are at current time
	n := 0
	for _, elem := range events {
		if elem > sim.CurrentTime {
			events[n] = elem
			n++
		}
	}

	filteredEvents := events[:n]

	// Sort it to get minimum element
	sort.Slice(filteredEvents, func(i, j int) bool { return events[i] < events[j] })

	// If the next action is  the GCD, just return
	if filteredEvents[0] == ret.CDReadyAt(core.GCDCooldownID) {
		return
	}

	// Otherwise add a pending action for the next time
	pa := &core.PendingAction{
		Priority:     core.ActionPriorityLow,
		OnAction:     ret.ActRotation,
		NextActionAt: filteredEvents[0],
	}

	sim.AddPendingAction(pa)
}

func (ret *RetributionPaladin) useFillers(sim *core.Simulation) {
	return
}

// Once filler moves are implemented, experiment with various mana settings
// See if its needed to use 2007 rotation or a variation at low mana
func (ret *RetributionPaladin) _2007Rotation(sim *core.Simulation) {
	target := sim.GetPrimaryTarget()

	// judge blood whenever we can
	if ret.CanJudgementOfBlood(sim) {
		success := ret.JudgementOfBlood.Cast(sim, target)
		if !success {
			ret.WaitForMana(sim, ret.JudgementOfBlood.Instance.Cost.Value)
		}
	}

	// roll seal of blood
	if !ret.HasAura(paladin.SealOfBloodAuraID) {
		sob := ret.NewSealOfBlood(sim)
		if success := sob.StartCast(sim); !success {
			ret.WaitForMana(sim, sob.GetManaCost())
		}
		return
	}

	// Crusader strike if we can
	if !ret.IsOnCD(paladin.CrusaderStrikeCD, sim.CurrentTime) {
		success := ret.CrusaderStrike.Cast(sim, target)
		if !success {
			ret.WaitForMana(sim, ret.CrusaderStrike.Instance.Cost.Value)
		}
		return
	}

	// Proceed until SoB expires, CrusaderStrike comes off GCD, or Judgement comes off GCD
	nextEventAt := ret.CDReadyAt(paladin.CrusaderStrikeCD)
	sobExpiration := sim.CurrentTime + ret.RemainingAuraDuration(sim, paladin.SealOfBloodAuraID)
	nextEventAt = core.MinDuration(nextEventAt, sobExpiration)
	// Waiting for judgement CD causes a bug that infinite loops for some reason
	// nextEventAt = core.MinDuration(nextEventAt, ret.CDReadyAt(paladin.JudgementCD))
	ret.WaitUntil(sim, nextEventAt)
}
