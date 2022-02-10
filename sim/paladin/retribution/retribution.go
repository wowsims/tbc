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
		OffHand:        ret.WeaponFromOffHand(ret.DefaultMeleeCritMultiplier()),
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
	target := sim.GetPrimaryTarget()

	if !ret.HasAura(paladin.SealOfBloodAuraID) {
		cast := ret.NewSealOfBlood(sim)
		if success := cast.StartCast(sim); !success {
			ret.WaitForMana(sim, cast.GetManaCost())
		}
		return
	}
	
	// check if we can use crusader strike
	if !ret.IsOnCD(paladin.CrusaderStrikeCD, sim.CurrentTime) {
		cs := ret.NewCrusaderStrike(sim, target)
		if success := cs.Attack(sim); !success {
			ret.WaitForMana(sim, cs.Cost.Value)
		}
		return
	}

	// probably should check for the min of crusader strike CD or SoB expiration
	nextEventAt := ret.CDReadyAt(paladin.CrusaderStrikeCD)
	ret.WaitUntil(sim, nextEventAt)
}
