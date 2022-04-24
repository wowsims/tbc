package retribution

import (
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

	// Setup Seal of Command after autos are enabled so that the PPM works
	ret.SetupSealOfCommand()

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

func (ret *RetributionPaladin) Init(sim *core.Simulation) {
	ret.Paladin.Init(sim)
	ret.DelayDPSCooldownsForArmorDebuffs(sim)
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
	ret.rotation(sim)
}

func (ret *RetributionPaladin) openingRotation(sim *core.Simulation) {
	target := sim.GetPrimaryTarget()

	// Cast selected judgement to keep on the boss
	if ret.JudgementOfWisdom.IsReady(sim) &&
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
				ret.WaitForMana(sim, judge.CurCast.Cost)
			}
		}
	}

	// Cast Seal of Command
	if !ret.SealOfCommandAura.IsActive() {
		if success := ret.SealOfCommand.Cast(sim, nil); !success {
			ret.WaitForMana(sim, ret.SealOfCommand.CurCast.Cost)
		}
		return
	}

	// Cast Seal of Blood and enable attacks to twist
	if !ret.SealOfBloodAura.IsActive() {
		if success := ret.SealOfBlood.Cast(sim, nil); !success {
			ret.WaitForMana(sim, ret.SealOfBlood.CurCast.Cost)
		}
		ret.AutoAttacks.EnableAutoSwing(sim)
		ret.openerCompleted = true
	}
}

func (ret *RetributionPaladin) rotation(sim *core.Simulation) {
	// Need to check for SoC early
	socActive := ret.SealOfCommandAura.IsActive()

	// If mana is low, do the low mana rotation instead
	// Don't do the low mana rotation in the middle of a twist
	if ret.CurrentMana() <= 1000 && !socActive {
		ret.lowManaRotation(sim)
	}

	// Setup
	target := sim.GetPrimaryTarget()

	gcdCD := ret.GCD.TimeToReady(sim)
	crusaderStrikeCD := ret.CrusaderStrike.TimeToReady(sim)
	nextCrusaderStrikeCD := ret.CrusaderStrike.CD.ReadyAt()
	judgementCD := ret.JudgementOfWisdom.TimeToReady(sim)

	sobActive := ret.SealOfBloodAura.IsActive()

	nextSwingAt := ret.AutoAttacks.NextAttackAt()
	timeTilNextSwing := nextSwingAt - sim.CurrentTime
	//weaponSpeed := ret.AutoAttacks.MainhandSwingSpeed()

	spellGCD := ret.SpellGCD()

	inTwistWindow := (sim.CurrentTime >= nextSwingAt-twistWindow) && (sim.CurrentTime < ret.AutoAttacks.NextAttackAt())
	latestTwistStart := nextSwingAt - spellGCD
	possibleTwist := timeTilNextSwing > spellGCD+gcdCD
	willTwist := possibleTwist && (nextSwingAt+spellGCD <= nextCrusaderStrikeCD+ret.crusaderStrikeDelay)

	// Use Judgement if we will prep Seal of Command
	// Or if we can squeeze it in on a Crusader Strike Swing
	if judgementCD == 0 && sobActive && willTwist {
		ret.JudgementOfBlood.Cast(sim, target)
		sobActive = false
	}

	// Judgement can affect active seals and CDs
	nextJudgementCD := ret.JudgementOfWisdom.CD.ReadyAt()

	if gcdCD == 0 {
		if socActive && inTwistWindow {
			// If Seal of Command is Active, complete the twist
			ret.SealOfBlood.Cast(sim, nil)
		} else if crusaderStrikeCD == 0 && !willTwist &&
			(sobActive || spellGCD < timeTilNextSwing) {
			// Cast Crusader Strike if we won't swing naked and we aren't twisting
			ret.CrusaderStrike.Cast(sim, target)
		} else if willTwist && !socActive && (nextJudgementCD > latestTwistStart) {
			// Prep seal of command
			ret.SealOfCommand.Cast(sim, nil)
		} else if !sobActive && !socActive && !willTwist {
			// If no seal is active, cast Seal of Blood
			ret.SealOfBlood.Cast(sim, nil)
		}
	}

	// All possible next events
	events := []time.Duration{
		nextSwingAt,
		nextSwingAt - twistWindow,
		ret.GCD.ReadyAt(),
		ret.JudgementOfWisdom.CD.ReadyAt(),
		ret.CrusaderStrike.CD.ReadyAt(),
	}

	ret.waitUntilNextEvent(sim, events)
}

func (ret *RetributionPaladin) useFillers(sim *core.Simulation, target *core.Target, sobActive bool) {
}

// Just roll seal of blood and cast crusader strike on CD to conserve mana
func (ret *RetributionPaladin) lowManaRotation(sim *core.Simulation) {
	target := sim.GetPrimaryTarget()

	nextSwingAt := ret.AutoAttacks.NextAttackAt()
	sobExpiration := ret.SealOfBloodAura.ExpiresAt()

	spellGCD := ret.SpellGCD()

	sobAndCSCost := ret.CrusaderStrike.DefaultCast.Cost + ret.SealOfBlood.DefaultCast.Cost
	sobAndJudgementCost := ret.JudgementOfBlood.DefaultCast.Cost + ret.SealOfBlood.DefaultCast.Cost

	if !ret.GCD.IsReady(sim) {
		// Roll seal of blood
		if sim.CurrentTime+time.Second >= sobExpiration {
			if ret.CanJudgementOfBlood(sim) && ret.CurrentMana() >= sobAndJudgementCost {
				ret.JudgementOfBlood.Cast(sim, target)
			}
			ret.SealOfBlood.Cast(sim, target)
		}

		// Crusader strike unless it will cause seal of blood to drop
		// Or we won't have enough mana to reseal
		if !ret.CrusaderStrike.CD.IsReady(sim) &&
			!(spellGCD+sim.CurrentTime > nextSwingAt && sobExpiration < nextSwingAt) &&
			(ret.CurrentMana() >= sobAndCSCost) {
			ret.CrusaderStrike.Cast(sim, target)
		}
	}

	events := []time.Duration{
		ret.GCD.ReadyAt(),
		ret.CrusaderStrike.CD.ReadyAt(),
		// ret.TimeUntilManaRegen(sobAndCSCost),
		sobExpiration - time.Second,
	}

	ret.waitUntilNextEvent(sim, events)
}

// Helper function for finding the next event
func (ret *RetributionPaladin) waitUntilNextEvent(sim *core.Simulation, events []time.Duration) {
	// Find the minimum possible next event that is greater than the current time
	nextEventAt := events[0]
	for _, elem := range events {
		if elem > sim.CurrentTime && elem < nextEventAt {
			nextEventAt = elem
		}
	}
	// If the next action is  the GCD, just return
	if nextEventAt == ret.GCD.ReadyAt() {
		return
	}

	// Otherwise add a pending action for the next time
	pa := &core.PendingAction{
		Priority:     core.ActionPriorityLow,
		OnAction:     ret.rotation,
		NextActionAt: nextEventAt,
	}

	sim.AddPendingAction(pa)
}
