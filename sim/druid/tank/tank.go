package tank

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
	"github.com/wowsims/tbc/sim/druid"
)

func RegisterFeralTankDruid() {
	core.RegisterAgentFactory(
		proto.Player_FeralTankDruid{},
		proto.Spec_SpecFeralTankDruid,
		func(character core.Character, options proto.Player) core.Agent {
			return NewFeralTankDruid(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_FeralTankDruid)
			if !ok {
				panic("Invalid spec value for Feral Tank Druid!")
			}
			player.Spec = playerSpec
		},
	)
}

func NewFeralTankDruid(character core.Character, options proto.Player) *FeralTankDruid {
	tankOptions := options.GetFeralTankDruid()

	selfBuffs := druid.SelfBuffs{}
	if tankOptions.Options.InnervateTarget != nil {
		selfBuffs.InnervateTarget = *tankOptions.Options.InnervateTarget
	} else {
		selfBuffs.InnervateTarget.TargetIndex = -1
	}

	bear := &FeralTankDruid{
		Druid:    druid.New(character, druid.Bear, selfBuffs, *tankOptions.Talents),
		Rotation: *tankOptions.Rotation,
		Options:  *tankOptions.Options,
	}

	bear.EnableRageBar(bear.Options.StartingRage, 1, func(sim *core.Simulation) {
		if bear.GCD.IsReady(sim) {
			bear.TryUseCooldowns(sim)
			if bear.GCD.IsReady(sim) {
				bear.doRotation(sim)
			}
		}
	})

	bear.EnableAutoAttacks(bear, core.AutoAttackOptions{
		// Base paw weapon.
		MainHand: core.Weapon{
			BaseDamageMin:        43.5,
			BaseDamageMax:        66.5,
			SwingSpeed:           2.5,
			NormalizedSwingSpeed: 2.5,
			SwingDuration:        time.Millisecond * 2500,
			CritMultiplier:       bear.MeleeCritMultiplier(),
		},
		AutoSwingMelee: true,
	})

	// bear Form adds (2 x Level) AP + 1 AP per Agi
	bear.AddStat(stats.AttackPower, 140)
	bear.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Agility,
		ModifiedStat: stats.AttackPower,
		Modifier: func(agility float64, attackPower float64) float64 {
			return attackPower + agility*1
		},
	})

	bear.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.FeralAttackPower,
		ModifiedStat: stats.AttackPower,
		Modifier: func(feralAttackPower float64, attackPower float64) float64 {
			return attackPower + feralAttackPower*1
		},
	})

	return bear
}

type FeralTankDruid struct {
	*druid.Druid

	Rotation proto.FeralTankDruid_Rotation
	Options  proto.FeralTankDruid_Options
}

func (bear *FeralTankDruid) GetDruid() *druid.Druid {
	return bear.Druid
}

func (bear *FeralTankDruid) Initialize() {
	bear.Druid.Initialize()
	bear.RegisterBearSpells(float64(bear.Rotation.MaulRageThreshold))
}
