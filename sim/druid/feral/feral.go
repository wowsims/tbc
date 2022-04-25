package feral

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
	"github.com/wowsims/tbc/sim/druid"
)

func RegisterFeralDruid() {
	core.RegisterAgentFactory(
		proto.Player_FeralDruid{},
		proto.Spec_SpecFeralDruid,
		func(character core.Character, options proto.Player) core.Agent {
			return NewFeralDruid(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_FeralDruid)
			if !ok {
				panic("Invalid spec value for Feral Druid!")
			}
			player.Spec = playerSpec
		},
	)
}

func NewFeralDruid(character core.Character, options proto.Player) *FeralDruid {
	feralOptions := options.GetFeralDruid()

	selfBuffs := druid.SelfBuffs{}
	if feralOptions.Options.InnervateTarget != nil {
		selfBuffs.InnervateTarget = *feralOptions.Options.InnervateTarget
	} else {
		selfBuffs.InnervateTarget.TargetIndex = -1
	}

	druid := druid.New(character, selfBuffs, *feralOptions.Talents)
	druid.CatForm = true
	cat := &FeralDruid{
		Druid:    druid,
		Rotation: *feralOptions.Rotation,
	}

	// Passive Cat Form threat reduction
	cat.PseudoStats.ThreatMultiplier *= 0.71

	cat.EnableEnergyBar(100.0, func(sim *core.Simulation) {
		cat.TryUseCooldowns(sim)
		if cat.GCD.IsReady(sim) {
			cat.doRotation(sim)
		}
	})

	// Set up base paw weapon. Assume that Predatory Instincts is a primary rather than secondary modifier for now, but this needs to confirmed!
	primaryModifier := 1 + 0.02*float64(cat.Talents.PredatoryInstincts)
	critMultiplier := cat.MeleeCritMultiplier(primaryModifier, 0)
	basePaw := core.Weapon{
		BaseDamageMin:        43.5,
		BaseDamageMax:        66.5,
		SwingSpeed:           1.0,
		NormalizedSwingSpeed: 1.0,
		SwingDuration:        time.Duration(1.0 * float64(time.Second)),
		CritMultiplier:       critMultiplier,
	}
	cat.EnableAutoAttacks(cat, core.AutoAttackOptions{
		MainHand:       basePaw,
		AutoSwingMelee: true,
	})

	// Cat Form adds (2 x Level) AP + 1 AP per Agi
	cat.AddStat(stats.AttackPower, 140)
	cat.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Agility,
		ModifiedStat: stats.AttackPower,
		Modifier: func(agility float64, attackPower float64) float64 {
			return attackPower + agility*1
		},
	})

	cat.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.FeralAttackPower,
		ModifiedStat: stats.AttackPower,
		Modifier: func(feralAttackPower float64, attackPower float64) float64 {
			return attackPower + feralAttackPower*1
		},
	})

	return cat
}

type FeralDruid struct {
	*druid.Druid

	Rotation proto.FeralDruid_Rotation
}

// GetDruid is to implement druid.Agent (supports nordrassil set bonus)
func (cat *FeralDruid) GetDruid() *druid.Druid {
	return cat.Druid
}
