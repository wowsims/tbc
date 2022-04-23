package retribution

// THIS IS A WORK IN PROGRESS IT IS NOT EVEN CLOSE TO COMPLETE
// I beg don't judge I'm a moron
// Orchrist

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

func RegisterProtectionPaladin() {
	core.RegisterAgentFactory(
		proto.Player_ProtectionPaladin{},
		proto.Spec_SpecProtectionPaladin,
		func(character core.Character, options proto.Player) core.Agent {
			return NewProtectionPaladin(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_ProtectionPaladin) // I don't really understand this line
			if !ok {
				panic("Invalid spec value for Protection Paladin!")
			}
			player.Spec = playerSpec
		},
	)
}

func NewProtectionPaladin(character core.Character, options proto.Player) *ProtectionPaladin {
	protOptions := options.GetProtectionPaladin()

	ret := &ProtectionPaladin{
		Paladin:             paladin.NewPaladin(character, *protOptions.Talents),
		Rotation:            *protOptions.Rotation,
		PrimaryJudgement:     protOptions.Options.PrimaryJudgement,
		BuffJudgement:     protOptions.Options.BuffJudgement,
	}

	// Convert DTPS option to bonus MP5
	spAtt := protOptions.Options.DamageTakenPerSecond * 5.0 / 10.0
	prot.AddStat(stats.MP5, spAtt)

	prot.EnableAutoAttacks(ret, core.AutoAttackOptions{
		MainHand:       prot.WeaponFromMainHand(prot.DefaultMeleeCritMultiplier()),
		AutoSwingMelee: true,
	})

	return prot
}

type ProtectionPaladin struct {
	*paladin.Paladin

	openerCompleted bool

	primaryjudgement proto.ProtectionPaladin_Options_PrimaryJudgement
	buffjudgement proto.ProtectionPaladin_Options_BuffJudgement

	Rotation proto.ProtectionPaladin_Rotation
}

func (prot *ProtectionPaladin) GetPaladin() *paladin.Paladin {
	return prot.Paladin
}

func (prot *ProtectionPaladin) Init(sim *core.Simulation) {
	prot.Paladin.Init(sim)
	prot.DelayDPSCooldownsForArmorDebuffs(sim)
}

func (prot *ProtectionPaladin) Reset(sim *core.Simulation) {
	prot.Paladin.Reset(sim)

	switch prot.buffjudgement {
	case proto.ProtectionPaladin_Options_Wisdom:
		prot.UpdateSeal(sim, prot.SealOfWisdomAura)
	case proto.ProtectionPaladin_Options_Crusader:
		prot.UpdateSeal(sim, prot.SealOfTheCrusaderAura)
	case proto.ProtectionPaladin_Options_None:
		prot.UpdateSeal(sim, prot.SealOfCommandAura)
	}

	prot.AutoAttacks.CancelAutoSwing(sim)
	prot.openerCompleted = false
}

func (prot *ProtectionPaladin) OnGCDReady(sim *core.Simulation) {
	prot.tryUseGCD(sim)
}

func (prot *ProtectionPaladin) OnManaTick(sim *core.Simulation) {
	if prot.FinishedWaitingForManaAndGCDReady(sim) {
		prot.tryUseGCD(sim)
	}
}

func (prot *ProtectionPaladin) tryUseGCD(sim *core.Simulation) {
	if !prot.openerCompleted {
		prot.openingRotation(sim)
		return
	}
	prot.ActRotation(sim)
}

func (prot *ProtectionPaladin) openingRotation(sim *core.Simulation) {
	target := sim.GetPrimaryTarget()

	// Cast selected judgement to keep on the boss
	if prot.JudgementOfWisdom.IsReady(sim) &&
		prot.judgement != proto.ProtectionPaladin_Options_None {
		var judge *core.Spell
		switch prot.judgement {
		case proto.ProtectionPaladin_Options_Wisdom:
			judge = prot.JudgementOfWisdom
		case proto.ProtectionPaladin_Options_Crusader:
			judge = prot.JudgementOfTheCrusader
		}
		if judge != nil {
			if success := judge.Cast(sim, target); !success {
				prot.WaitForMana(sim, judge.CurCast.Cost)
			}
		}
	}

	// Cast Seal of Command
	if !prot.SealOfCommandAura.IsActive() {
		if success := prot.SealOfCommand.Cast(sim, nil); !success {
			prot.WaitForMana(sim, prot.SealOfCommand.CurCast.Cost)
		}
		return
	}

	// Cast Seal of Blood and enable attacks to twist
	if !prot.SealOfBloodAura.IsActive() {
		if success := prot.SealOfBlood.Cast(sim, nil); !success {
			prot.WaitForMana(sim, prot.SealOfBlood.CurCast.Cost)
		}
		prot.AutoAttacks.EnableAutoSwing(sim)
		prot.openerCompleted = true
	}
}

func (prot *ProtectionPaladin) ActRotation(sim *core.Simulation) {
	// Setup
	target := sim.GetPrimaryTarget()

	gcdCD := prot.GCD.TimeToReady(sim)
	crusaderStrikeCD := prot.CrusaderStrike.TimeToReady(sim)
	nextCrusaderStrikeCD := prot.CrusaderStrike.CD.ReadyAt()
	judgementCD := prot.JudgementOfWisdom.TimeToReady(sim)

	sobActive := prot.SealOfBloodAura.IsActive()
	socActive := prot.SealOfCommandAura.IsActive()

	nextSwingAt := prot.AutoAttacks.NextAttackAt()
	timeTilNextSwing := nextSwingAt - sim.CurrentTime
	//weaponSpeed := prot.AutoAttacks.MainhandSwingSpeed()

	spellGCD := prot.SpellGCD()

	inTwistWindow := (sim.CurrentTime >= nextSwingAt-twistWindow) && (sim.CurrentTime < prot.AutoAttacks.NextAttackAt())
	latestTwistStart := nextSwingAt - spellGCD
	possibleTwist := timeTilNextSwing > spellGCD+gcdCD
	willTwist := possibleTwist && (nextSwingAt+spellGCD <= nextCrusaderStrikeCD+prot.crusaderStrikeDelay)

	// Use Judgement if we will prep Seal of Command
	// Or if we can squeeze it in on a Crusader Strike Swing
	if judgementCD == 0 && sobActive && willTwist {
		prot.JudgementOfBlood.Cast(sim, target)
		sobActive = false
	}

	// Judgement can affect active seals and CDs
	nextJudgementCD := prot.JudgementOfWisdom.CD.ReadyAt()

	if gcdCD == 0 {
		if socActive && inTwistWindow {
			// If Seal of Command is Active, complete the twist
			prot.SealOfBlood.Cast(sim, nil)
		} else if crusaderStrikeCD == 0 && !willTwist &&
			(sobActive || spellGCD < timeTilNextSwing) {
			// Cast Crusader Strike if we won't swing naked and we aren't twisting
			prot.CrusaderStrike.Cast(sim, target)
		} else if willTwist && !socActive && (nextJudgementCD > latestTwistStart) {
			// Prep seal of command
			prot.SealOfCommand.Cast(sim, nil)
		} else if !sobActive && !socActive && !willTwist {
			// If no seal is active, cast Seal of Blood
			prot.SealOfBlood.Cast(sim, nil)
		}
	}

	// Determine when next action is available
	// Throw everything into an array then filter and sort compared to doing individual comparisons
	events := []time.Duration{
		nextSwingAt,
		nextSwingAt - twistWindow,
		prot.GCD.ReadyAt(),
		prot.JudgementOfWisdom.CD.ReadyAt(),
		prot.CrusaderStrike.CD.ReadyAt(),
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
	if filteredEvents[0] == prot.GCD.ReadyAt() {
		return
	}

	// Otherwise add a pending action for the next time
	pa := &core.PendingAction{
		Priority:     core.ActionPriorityLow,
		OnAction:     prot.ActRotation,
		NextActionAt: filteredEvents[0],
	}

	sim.AddPendingAction(pa)
}

func (prot *ProtectionPaladin) useFillers(sim *core.Simulation, target *core.Target, sobActive bool) {

}

// Once filler moves are implemented, experiment with various mana settings
// See if its needed to use 2007 rotation or a variation at low mana
func (prot *ProtectionPaladin) _2007Rotation(sim *core.Simulation) {
	target := sim.GetPrimaryTarget()

	// judge blood whenever we can
	if prot.CanJudgementOfBlood(sim) {
		success := prot.JudgementOfBlood.Cast(sim, target)
		if !success {
			prot.WaitForMana(sim, prot.JudgementOfBlood.CurCast.Cost)
		}
	}

	// roll seal of blood
	if !prot.SealOfBloodAura.IsActive() {
		if success := prot.SealOfBlood.Cast(sim, nil); !success {
			prot.WaitForMana(sim, prot.SealOfBlood.CurCast.Cost)
		}
		return
	}

	// Crusader strike if we can
	if prot.CrusaderStrike.IsReady(sim) {
		success := prot.CrusaderStrike.Cast(sim, target)
		if !success {
			prot.WaitForMana(sim, prot.CrusaderStrike.CurCast.Cost)
		}
		return
	}

	// Proceed until SoB expires, CrusaderStrike comes off GCD, or Judgement comes off GCD
	nextEventAt := prot.CrusaderStrike.CD.ReadyAt()
	sobExpiration := sim.CurrentTime + prot.SealOfBloodAura.RemainingDuration(sim)
	nextEventAt = core.MinDuration(nextEventAt, sobExpiration)
	// Waiting for judgement CD causes a bug that infinite loops for some reason
	// nextEventAt = core.MinDuration(nextEventAt, prot.CDReadyAt(paladin.JudgementCD))
	prot.WaitUntil(sim, nextEventAt)
}
