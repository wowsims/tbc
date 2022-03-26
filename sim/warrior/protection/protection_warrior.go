package protection

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/warrior"
)

func RegisterProtectionWarrior() {
	core.RegisterAgentFactory(
		proto.Player_ProtectionWarrior{},
		proto.Spec_SpecProtectionWarrior,
		func(character core.Character, options proto.Player) core.Agent {
			return NewProtectionWarrior(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_ProtectionWarrior)
			if !ok {
				panic("Invalid spec value for Protection Warrior!")
			}
			player.Spec = playerSpec
		},
	)
}

func NewProtectionWarrior(character core.Character, options proto.Player) *ProtectionWarrior {
	warOptions := options.GetProtectionWarrior()

	war := &ProtectionWarrior{
		Warrior: warrior.NewWarrior(character, *warOptions.Talents),
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
		ReplaceMHSwing: func(sim *core.Simulation) *core.SimpleSpell {
			return war.TryHeroicStrike(sim)
		},
	})

	return war
}

type ProtectionWarrior struct {
	*warrior.Warrior

	Options proto.Warrior_Options
}

func (war *ProtectionWarrior) GetWarrior() *warrior.Warrior {
	return war.Warrior
}
