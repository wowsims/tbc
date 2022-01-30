package warrior

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func RegisterWarrior() {
	core.RegisterAgentFactory(
		proto.Player_Warrior{},
		func(character core.Character, options proto.Player) core.Agent {
			return NewWarrior(character, options)
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

type Warrior struct {
	core.Character

	Talents          proto.WarriorTalents
	Options          proto.Warrior_Options
	RotationType     proto.Warrior_Rotation_Type
	ArmsSlamRotation proto.Warrior_Rotation_ArmsSlamRotation
	ArmsDwRotation   proto.Warrior_Rotation_ArmsDWRotation
	FuryRotation     proto.Warrior_Rotation_FuryRotation
}

func (warrior *Warrior) GetCharacter() *core.Character {
	return &warrior.Character
}

func (warrior *Warrior) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {
}

func (warrior *Warrior) Init(sim *core.Simulation) {
}

func (warrior *Warrior) Reset(newsim *core.Simulation) {
}

func NewWarrior(character core.Character, options proto.Player) *Warrior {
	warriorOptions := options.GetWarrior()

	warrior := &Warrior{
		Character:    character,
		Talents:      *warriorOptions.Talents,
		Options:      *warriorOptions.Options,
		RotationType: warriorOptions.Rotation.Type,
	}

	warrior.PseudoStats.MeleeSpeedMultiplier = 1
	warrior.EnableRageBar(warriorOptions.Options.StartingRage)
	warrior.EnableAutoAttacks(warrior, core.AutoAttackOptions{
		MainHand:       warrior.WeaponFromMainHand(),
		OffHand:        warrior.WeaponFromOffHand(),
		AutoSwingMelee: true,
	})

	if warrior.RotationType == proto.Warrior_Rotation_ArmsSlam && warriorOptions.Rotation.ArmsSlam != nil {
		warrior.ArmsSlamRotation = *warriorOptions.Rotation.ArmsSlam
	} else if warrior.RotationType == proto.Warrior_Rotation_ArmsDW && warriorOptions.Rotation.ArmsDw != nil {
		warrior.ArmsDwRotation = *warriorOptions.Rotation.ArmsDw
	} else if warrior.RotationType == proto.Warrior_Rotation_Fury && warriorOptions.Rotation.Fury != nil {
		warrior.FuryRotation = *warriorOptions.Rotation.Fury
	}

	warrior.Character.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Agility,
		ModifiedStat: stats.MeleeCrit,
		Modifier: func(agility float64, meleecrit float64) float64 {
			return meleecrit + (agility/33)*core.MeleeCritRatingPerCritChance
		},
	})

	warrior.Character.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Strength,
		ModifiedStat: stats.AttackPower,
		Modifier: func(strength float64, attackPower float64) float64 {
			return attackPower + strength*2
		},
	})

	//warrior.applyTalents()

	return warrior
}

func init() {
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDraenei, Class: proto.Class_ClassWarrior}] = stats.Stats{
		stats.Strength:  146,
		stats.Agility:   93,
		stats.Stamina:   132,
		stats.Intellect: 34,
		stats.Spirit:    53,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDwarf, Class: proto.Class_ClassWarrior}] = stats.Stats{
		stats.Strength:  147,
		stats.Agility:   92,
		stats.Stamina:   136,
		stats.Intellect: 32,
		stats.Spirit:    50,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceGnome, Class: proto.Class_ClassWarrior}] = stats.Stats{
		stats.Strength:  140,
		stats.Agility:   99,
		stats.Stamina:   132,
		stats.Intellect: 38,
		stats.Spirit:    51,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceHuman, Class: proto.Class_ClassWarrior}] = stats.Stats{
		stats.Strength:  145,
		stats.Agility:   96,
		stats.Stamina:   133,
		stats.Intellect: 33,
		stats.Spirit:    56,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceNightElf, Class: proto.Class_ClassWarrior}] = stats.Stats{
		stats.Strength:  142,
		stats.Agility:   101,
		stats.Stamina:   132,
		stats.Intellect: 33,
		stats.Spirit:    51,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceOrc, Class: proto.Class_ClassWarrior}] = stats.Stats{
		stats.Strength:  148,
		stats.Agility:   93,
		stats.Stamina:   135,
		stats.Intellect: 30,
		stats.Spirit:    54,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTauren, Class: proto.Class_ClassWarrior}] = stats.Stats{
		stats.Strength:  150,
		stats.Agility:   91,
		stats.Stamina:   135,
		stats.Intellect: 28,
		stats.Spirit:    53,
	}
	trollStats := stats.Stats{
		stats.Strength:  146,
		stats.Agility:   98,
		stats.Stamina:   134,
		stats.Intellect: 29,
		stats.Spirit:    52,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll10, Class: proto.Class_ClassWarrior}] = trollStats
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll30, Class: proto.Class_ClassWarrior}] = trollStats
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceUndead, Class: proto.Class_ClassWarrior}] = stats.Stats{
		stats.Strength:  144,
		stats.Agility:   94,
		stats.Stamina:   134,
		stats.Intellect: 31,
		stats.Spirit:    56,
	}
}

// Agent is a generic way to access underlying warrior on any of the agents.
type Agent interface {
	GetWarrior() *Warrior
}
