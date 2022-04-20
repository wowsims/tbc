package dps

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
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
		Warrior: warrior.NewWarrior(character, *warOptions.Talents, warrior.WarriorInputs{
			ShoutType:            warOptions.Options.Shout,
			PrecastShout:         warOptions.Options.PrecastShout,
			PrecastShoutSapphire: warOptions.Options.PrecastShoutSapphire,
			PrecastShoutT2:       warOptions.Options.PrecastShoutT2,
			RampageCDThreshold:   core.DurationFromSeconds(warOptions.Rotation.RampageCdThreshold),
		}),
		Rotation: *warOptions.Rotation,
	}

	war.EnableRageBar(warOptions.Options.StartingRage, core.TernaryFloat64(war.Talents.EndlessRage, 1.25, 1), func(sim *core.Simulation) {
		if war.GCD.IsReady(sim) {
			war.doRotation(sim)
		}
	})
	war.EnableAutoAttacks(war, core.AutoAttackOptions{
		MainHand:       war.WeaponFromMainHand(war.DefaultMeleeCritMultiplier()),
		OffHand:        war.WeaponFromOffHand(war.DefaultMeleeCritMultiplier()),
		AutoSwingMelee: true,
		ReplaceMHSwing: func(sim *core.Simulation) *core.Spell {
			if war.Rotation.UseCleave {
				return war.TryCleave(sim)
			} else {
				return war.TryHeroicStrike(sim)
			}
		},
	})

	// TODO: This should only be applied while berserker stance is active.
	if war.Talents.ImprovedBerserkerStance > 0 {
		bonus := 1 + 0.02*float64(war.Talents.ImprovedBerserkerStance)
		war.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.AttackPower,
			ModifiedStat: stats.AttackPower,
			Modifier: func(ap float64, _ float64) float64 {
				return ap * bonus
			},
		})
	}

	if war.Options.UseRecklessness {
		war.RegisterRecklessnessCD()
	}

	return war
}

type DpsWarrior struct {
	*warrior.Warrior

	Options  proto.Warrior_Options
	Rotation proto.Warrior_Rotation
}

func (war *DpsWarrior) GetWarrior() *warrior.Warrior {
	return war.Warrior
}

func (war *DpsWarrior) Init(sim *core.Simulation) {
	war.Warrior.Init(sim)
	war.DelayDPSCooldownsForArmorDebuffs(sim)
}

func (war *DpsWarrior) Reset(sim *core.Simulation) {
	war.Warrior.Reset(sim)
	war.BerserkerStanceAura.Activate(sim)
}
