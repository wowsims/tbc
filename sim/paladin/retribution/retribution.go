package retribution

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
	"github.com/wowsims/tbc/sim/paladin"
)

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
		Paladin:     paladin.NewPaladin(character, *retOptions.Talents),
		Rotation:    *retOptions.Rotation,
		csDelay:     time.Duration(retOptions.Options.CrusaderStrikeDelayMs),
		hasteLeeway: time.Duration(retOptions.Options.HasteLeewayMs),
		judgement:   retOptions.Options.Judgement,
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

	hasteLeeway time.Duration
	csDelay     time.Duration

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
	ret._2007Rotation(sim)
}

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
	if !ret.SealOfBloodAura.IsActive() {
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
	sobExpiration := sim.CurrentTime + ret.SealOfBloodAura.RemainingDuration(sim)
	nextEventAt = core.MinDuration(nextEventAt, sobExpiration)
	// Waiting for judgement CD causes a bug that infinite loops for some reason
	// nextEventAt = core.MinDuration(nextEventAt, ret.CDReadyAt(paladin.JudgementCD))
	ret.WaitUntil(sim, nextEventAt)
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
	if !ret.SealOfCommandAura.IsActive() {
		soc := ret.NewSealOfCommand(sim)
		if success := soc.StartCast(sim); !success {
			ret.WaitForMana(sim, soc.GetManaCost())
		}
		return
	}

	// Cast Seal of Blood and enable attacks to twist
	if !ret.SealOfBloodAura.IsActive() {
		sob := ret.NewSealOfBlood(sim)
		if success := sob.StartCast(sim); !success {
			ret.WaitForMana(sim, sob.GetManaCost())
		}
		ret.AutoAttacks.EnableAutoSwing(sim)
		ret.openerCompleted = true
	}
}

func (ret *RetributionPaladin) testingMechanics(sim *core.Simulation) {

}
