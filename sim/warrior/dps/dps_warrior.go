package dps

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/warrior"
)

func RegisterDpsWarrior() {
	core.RegisterAgentFactory(
		proto.Player_Warrior{},
		proto.Spec_SpecWarrior,
		func(character core.Character, options proto.Player) core.Agent {
			return NewDpsWarrior(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_Warrior)
			if !ok {
				panic("Invalid spec value for Warrior!")
			}
			player.Spec = playerSpec
		},
	)
}

func NewDpsWarrior(character core.Character, options proto.Player) *DpsWarrior {
	warOptions := options.GetWarrior()

	war := &DpsWarrior{
		Warrior:      warrior.NewWarrior(character, *warOptions.Talents, warrior.WarriorInputs{}),
		RotationType: warOptions.Rotation.Type,
	}

	if war.RotationType == proto.Warrior_Rotation_ArmsSlam && warOptions.Rotation.ArmsSlam != nil {
		war.ArmsSlamRotation = *warOptions.Rotation.ArmsSlam
	} else if war.RotationType == proto.Warrior_Rotation_ArmsDW && warOptions.Rotation.ArmsDw != nil {
		war.ArmsDwRotation = *warOptions.Rotation.ArmsDw
	} else if war.RotationType == proto.Warrior_Rotation_Fury && warOptions.Rotation.Fury != nil {
		war.FuryRotation = *warOptions.Rotation.Fury
	}

	war.EnableRageBar(warOptions.Options.StartingRage, func(sim *core.Simulation) {
		if !war.IsOnCD(core.GCDCooldownID, sim.CurrentTime) {
			war.doRotation(sim)
		}
	})
	war.EnableAutoAttacks(war, core.AutoAttackOptions{
		MainHand:       war.WeaponFromMainHand(war.DefaultMeleeCritMultiplier()),
		OffHand:        war.WeaponFromOffHand(war.DefaultMeleeCritMultiplier()),
		AutoSwingMelee: true,
		ReplaceMHSwing: func(sim *core.Simulation) *core.SimpleSpellTemplate {
			return war.TryHeroicStrike(sim)
		},
	})

	return war
}

type DpsWarrior struct {
	*warrior.Warrior

	Options          proto.Warrior_Options
	RotationType     proto.Warrior_Rotation_Type
	ArmsSlamRotation proto.Warrior_Rotation_ArmsSlamRotation
	ArmsDwRotation   proto.Warrior_Rotation_ArmsDWRotation
	FuryRotation     proto.Warrior_Rotation_FuryRotation
}

func (war *DpsWarrior) GetWarrior() *warrior.Warrior {
	return war.Warrior
}

func (war *DpsWarrior) Reset(sim *core.Simulation) {
	war.Warrior.Reset(sim)
	war.AddAura(sim, war.BerserkerStanceAura())
}
