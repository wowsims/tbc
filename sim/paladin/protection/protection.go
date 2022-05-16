package protection

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
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

	prot := &ProtectionPaladin{
		Paladin:  paladin.NewPaladin(character, *protOptions.Talents),
		Rotation: *protOptions.Rotation,
		Options:  *protOptions.Options,
	}

	prot.EnableAutoAttacks(prot, core.AutoAttackOptions{
		MainHand:       prot.WeaponFromMainHand(prot.DefaultMeleeCritMultiplier()),
		AutoSwingMelee: true,
	})

	return prot
}

type ProtectionPaladin struct {
	*paladin.Paladin

	Rotation proto.ProtectionPaladin_Rotation
	Options  proto.ProtectionPaladin_Options

	openerCompleted bool
}

func (prot *ProtectionPaladin) GetPaladin() *paladin.Paladin {
	return prot.Paladin
}

func (prot *ProtectionPaladin) Initialize() {
	prot.Paladin.Initialize()
}

func (prot *ProtectionPaladin) Reset(sim *core.Simulation) {
	prot.Paladin.Reset(sim)

	switch prot.Options.BuffJudgement {
	case proto.PaladinJudgement_JudgementOfWisdom:
		prot.UpdateSeal(sim, prot.SealOfWisdomAura)
	case proto.PaladinJudgement_JudgementOfCrusader:
		prot.UpdateSeal(sim, prot.SealOfTheCrusaderAura)
	}

	prot.AutoAttacks.CancelAutoSwing(sim)
	prot.openerCompleted = false
}
