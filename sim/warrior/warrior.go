package warrior

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

const (
	BattleStance = iota
	DefensiveStance
	BerserkerStance
)

type Warrior struct {
	core.Character

	Talents proto.WarriorTalents

	// Current state
	stance             int
	heroicStrikeQueued bool
	revengeTriggered   bool

	// Cached values
	heroicStrikeCost float64
	canShieldSlam    bool

	bloodthirstTemplate core.SimpleSpellTemplate
	bloodthirst         core.SimpleSpell

	devastateTemplate core.SimpleSpellTemplate
	devastate         core.SimpleSpell

	heroicStrikeTemplate core.SimpleSpellTemplate
	heroicStrike         core.SimpleSpell

	revengeTemplate core.SimpleSpellTemplate
	revenge         core.SimpleSpell

	shieldSlamTemplate core.SimpleSpellTemplate
	shieldSlam         core.SimpleSpell

	sunderArmorTemplate core.SimpleSpellTemplate
	sunderArmor         core.SimpleSpell

	whirlwindTemplate core.SimpleSpellTemplate
	whirlwind         core.SimpleSpell
}

func (warrior *Warrior) GetCharacter() *core.Character {
	return &warrior.Character
}

func (warrior *Warrior) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {
}

func (warrior *Warrior) Init(sim *core.Simulation) {
	warrior.bloodthirstTemplate = warrior.newBloodthirstTemplate(sim)
	warrior.devastateTemplate = warrior.newDevastateTemplate(sim)
	warrior.heroicStrikeTemplate = warrior.newHeroicStrikeTemplate(sim)
	warrior.revengeTemplate = warrior.newRevengeTemplate(sim)
	warrior.shieldSlamTemplate = warrior.newShieldSlamTemplate(sim)
	warrior.sunderArmorTemplate = warrior.newSunderArmorTemplate(sim)
	warrior.whirlwindTemplate = warrior.newWhirlwindTemplate(sim)
}

func (warrior *Warrior) Reset(newsim *core.Simulation) {
	warrior.stance = BerserkerStance
	warrior.heroicStrikeQueued = false
	warrior.revengeTriggered = false
}

func NewWarrior(character core.Character, talents proto.WarriorTalents) *Warrior {
	warrior := &Warrior{
		Character: character,
		Talents:   talents,
	}

	warrior.Character.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Agility,
		ModifiedStat: stats.MeleeCrit,
		Modifier: func(agility float64, meleecrit float64) float64 {
			return meleecrit + (agility/33)*core.MeleeCritRatingPerCritChance
		},
	})
	warrior.Character.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Agility,
		ModifiedStat: stats.Dodge,
		Modifier: func(agility float64, dodge float64) float64 {
			return dodge + (agility/30)*core.DodgeRatingPerDodgeChance
		},
	})
	warrior.Character.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Strength,
		ModifiedStat: stats.AttackPower,
		Modifier: func(strength float64, attackPower float64) float64 {
			return attackPower + strength*2
		},
	})
	warrior.Character.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Strength,
		ModifiedStat: stats.BlockValue,
		Modifier: func(strength float64, blockValue float64) float64 {
			return blockValue + strength/20
		},
	})

	return warrior
}

func (warrior *Warrior) critMultiplier(applyImpale bool) float64 {
	primaryModifier := 1.0
	secondaryModifier := 0.0

	if applyImpale {
		secondaryModifier += 0.1 * float64(warrior.Talents.Impale)
	}

	return warrior.MeleeCritMultiplier(primaryModifier, secondaryModifier)
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
