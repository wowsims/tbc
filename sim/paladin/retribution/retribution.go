package retribution

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/paladin"
)

func RegisterRetributionPaladin() {
	core.RegisterAgentFactory(
		proto.Player_RetributionPaladin{},
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
		Paladin:  paladin.NewPaladin(character, *retOptions.Talents),
		Rotation: *retOptions.Rotation,
	}

	ret.EnableAutoAttacks(ret, core.AutoAttackOptions{
		MainHand:       ret.WeaponFromMainHand(ret.DefaultMeleeCritMultiplier()),
		AutoSwingMelee: true,
	})

	return ret
}

type RetributionPaladin struct {
	*paladin.Paladin

	Rotation proto.RetributionPaladin_Rotation
}

func (ret *RetributionPaladin) GetPaladin() *paladin.Paladin {
	return ret.Paladin
}

func (ret *RetributionPaladin) Reset(sim *core.Simulation) {
	ret.Paladin.Reset(sim)
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
	ret._2007Rotation(sim)
}

func (ret *RetributionPaladin) _2007Rotation(sim *core.Simulation) {
	target := sim.GetPrimaryTarget()

	// judge blood whenever we can
	if !ret.IsOnCD(paladin.JudgementCD, sim.CurrentTime) {
		judge := ret.NewJudgement(sim, target)
		if judge != nil {
			if success := judge.Cast(sim); !success {
				ret.WaitForMana(sim, judge.GetManaCost())
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
		if success := cs.Attack(sim); !success {
			ret.WaitForMana(sim, cs.Cost.Value)
		}
		return
	}

	// Proceed until SoB expires, CrusaderStrike comes off GCD, or Judgement comes off GCD
	nextEventAt := ret.CDReadyAt(paladin.CrusaderStrikeCD)
	sobExpiration := sim.CurrentTime + ret.RemainingAuraDuration(sim, paladin.SealOfBloodAuraID)
	nextEventAt = core.MinDuration(nextEventAt, sobExpiration)
	nextEventAt = core.MinDuration(nextEventAt, ret.CDReadyAt(paladin.JudgementCD))
	ret.WaitUntil(sim, nextEventAt)
}

func (ret *RetributionPaladin) testingMechanics(sim *core.Simulation) {
	// target := sim.GetPrimaryTarget()

	if ret.Rotation.Consecration && (len(ret.ConsecrationSpell.Effects) == 0 || !ret.ConsecrationSpell.Effects[0].DotInput.IsTicking(sim)) {
		consc := ret.NewConsecration(sim)

		// if we dont have the mana.. do we just wait for regen?
		// Probably should have logic earlier on to decide if we can even cast this.
		//  Maybe we should have some pre-logic like elemental shaman to decide how much mana we have to spend ahead of time.
		if success := consc.Cast(sim); !success {
			ret.WaitForMana(sim, consc.GetManaCost())
		}

		return
	}
}
