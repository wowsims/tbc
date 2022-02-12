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
	if ret.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime) > 0 {
		return /// wtf?
	}

	target := sim.GetPrimaryTarget()

	// check if we can use crusader strike
	if !ret.IsOnCD(paladin.CrusaderStrikeCD, sim.CurrentTime) {
		cs := ret.NewCrusaderStrike(sim, target)
		if success := cs.Attack(sim); !success {
			ret.WaitForMana(sim, cs.Cost.Value)
		}
		return
	}

	tts := ret.AutoAttacks.MainhandSwingAt - sim.CurrentTime

	if ret.HasAura(paladin.SealOfCommandAuraID) {
		// maybe do a mana check first to make sure we don't twist when we don't have mana
		if tts <= paladin.TwistWindow {
			sob := ret.NewSealOfBlood(sim)

			// this is probably not the behaviour we want for not being able to twist
			if success := sob.StartCast(sim); !success {
				ret.WaitForMana(sim, sob.GetManaCost())
			}
			return
		} else {
			ret.WaitUntil(sim, ret.AutoAttacks.MainhandSwingAt-paladin.TwistWindow)
			return
		}
	} else if tts > ret.SpellGCD() {
		soc := ret.NewSealOfCommand(sim)
		if success := soc.StartCast(sim); !success {
			ret.WaitForMana(sim, soc.GetManaCost())
		}
		return
	} else if !ret.HasAura(paladin.SealOfBloodAuraID) {
		sob := ret.NewSealOfBlood(sim)
		if success := sob.StartCast(sim); !success {
			ret.WaitForMana(sim, sob.GetManaCost())
		}
		return
	}

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

	// probably should check for the min of crusader strike CD or SoB expiration
	nextEventAt := ret.CDReadyAt(paladin.CrusaderStrikeCD)
	ret.WaitUntil(sim, nextEventAt)
}
