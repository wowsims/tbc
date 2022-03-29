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
		Paladin:              paladin.NewPaladin(character, *retOptions.Talents),
		Rotation:             *retOptions.Rotation,
		crusaderStrikeDelay:  time.Duration(retOptions.Options.CrusaderStrikeDelayMs) * time.Millisecond,
		hasteLeeway:          time.Duration(retOptions.Options.HasteLeewayMs) * time.Millisecond,
		nextCrusaderStrikeCD: 0,
		judgement:            retOptions.Options.Judgement,
	}

	// Convert DTPS option to bonus MP5
	spAtt := retOptions.Options.DamageTakenPerSecond * 5.0 / 10.0
	ret.AddStat(stats.MP5, spAtt)

	ret.EnableAutoAttacks(ret, core.AutoAttackOptions{
		MainHand:       ret.WeaponFromMainHand(ret.DefaultMeleeCritMultiplier()),
		AutoSwingMelee: true,
	})

	ret.SetupSealOfCommand()

	return ret
}

type RetributionPaladin struct {
	*paladin.Paladin

	openerCompleted bool

	hasteLeeway          time.Duration
	crusaderStrikeDelay  time.Duration
	nextCrusaderStrikeCD time.Duration

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

	ret.SetupSealOfCommand() // Reset this to reset the internal CD back to time 0
	ret.AutoAttacks.CancelAutoSwing(sim)
	ret.nextCrusaderStrikeCD = 0
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
		var judge *core.SimpleSpell
		switch ret.judgement {
		case proto.RetributionPaladin_Options_Wisdom:
			judge = ret.NewJudgementOfWisdom(sim, target)
		case proto.RetributionPaladin_Options_Crusader:
			judge = ret.NewJudgementOfTheCrusader(sim, target)
		}
		if judge != nil {
			if success := judge.Cast(sim); !success {
				ret.WaitForMana(sim, judge.GetManaCost())
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

	gcdActive := ret.IsOnCD(core.GCDCooldownID, sim.CurrentTime)
	crusaderStrikeCD := ret.GetRemainingCD(paladin.CrusaderStrikeCD, sim.CurrentTime)
	judgmentCD := ret.GetRemainingCD(paladin.JudgementCD, sim.CurrentTime)
	// consecrateCD := ret.GetRemainingCD(paladin.ConsecrateCD, sim.CurrentTime)
	// exorcismCD := ret.GetRemainingCD(paladin.ExorcismCD, sim.CurrentTime)

	nextSwing := ret.AutoAttacks.NextAttackAt() - sim.CurrentTime
	// effWeaponSpeed := ret.AutoAttacks.MainhandSwingSpeed()
	spellGCD := ret.SpellGCD()

	possibleTwist := nextSwing > spellGCD
	willTwist := possibleTwist && (nextSwing+spellGCD <= ret.nextCrusaderStrikeCD+ret.crusaderStrikeDelay)
	inTwistWindow := sim.CurrentTime >= ret.AutoAttacks.NextAttackAt()-twistWindow && sim.CurrentTime < ret.AutoAttacks.NextAttackAt()
	latestTwistStart := nextSwing - spellGCD

	sobActive := ret.RemainingAuraDuration(sim, paladin.SealOfBloodAuraID) > 0
	socActive := ret.RemainingAuraDuration(sim, paladin.SealOfCommandAuraID) > 0

	// Use Judgement if we will twist
	if judgmentCD == 0 && willTwist && sobActive {
		judgement := ret.NewJudgementOfBlood(sim, target)
		if judgement != nil {
			judgement.Cast(sim)
		}
	}

	if !gcdActive {
		if socActive && inTwistWindow {
			// If Seal of Command is Active, complete the twist
			ret.NewSealOfBlood(sim).StartCast(sim)
		} else if crusaderStrikeCD == 0 &&
			(sobActive || spellGCD < nextSwing) {
			// Cast Crusader Strike if its up and we won't swing naked
			ret.NewCrusaderStrike(sim, target).Cast(sim)
			ret.nextCrusaderStrikeCD = sim.CurrentTime + 6*time.Second
			// } else if crusaderStrikeCD > spellGCD && !willTwist || (willTwist && nextSwing > spellGCD*2) {
			// 	// Use fillers if it won't clip Crusader Strike or a Twist
			// 	ret.useFillers(sim)
		} else if willTwist && !socActive && judgmentCD > latestTwistStart {
			// Prep seal of command
			ret.NewSealOfCommand(sim).StartCast(sim)
		} else if !sobActive && !socActive {
			// If no seal is active, cast Seal of Blood
			ret.NewSealOfBlood(sim).StartCast(sim)
		}
	}

	// Determine when next action is available
	// Throw everything into an array then filter and sort compared to doing individual comparisons
	var events [5]time.Duration
	events[0] = ret.AutoAttacks.NextAttackAt()                          // next swing
	events[1] = ret.AutoAttacks.NextAttackAt() - twistWindow            // next twist window
	events[2] = ret.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime) // next GCD
	events[3] = sim.CurrentTime + judgmentCD                            // next Judgement CD
	events[4] = sim.CurrentTime + crusaderStrikeCD                      // next Crusader Strike CD
	// events[5] = sim.CurrentTime + consecrateCD               // next Consecrate CD
	// events[6] = sim.CurrentTime + exorcismCD                 // next Exorcism CD

	// Time has to move forward... so exclude any events that are at current time
	n := 0
	for _, elem := range events {
		if elem > sim.CurrentTime {
			events[n] = elem
			n++
		}
	}

	var filteredEvents []time.Duration = events[:n]

	// Sort it to get minimum element
	sort.Slice(filteredEvents, func(i, j int) bool { return events[i] < events[j] })

	// If the next action is  the GCD, just return
	if filteredEvents[0] == ret.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime) {
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
	if !ret.IsOnCD(paladin.JudgementCD, sim.CurrentTime) {
		judge := ret.NewJudgementOfBlood(sim, target)
		if judge != nil {
			if success := judge.Cast(sim); !success {
				ret.WaitForMana(sim, judge.Cost.Value)
			}
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
		cs := ret.NewCrusaderStrike(sim, target)
		if success := cs.Cast(sim); !success {
			ret.WaitForMana(sim, cs.Cost.Value)
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
