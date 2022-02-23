package rogue

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func RegisterRogue() {
	core.RegisterAgentFactory(
		proto.Player_Rogue{},
		proto.Spec_SpecRogue,
		func(character core.Character, options proto.Player) core.Agent {
			return NewRogue(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_Rogue)
			if !ok {
				panic("Invalid spec value for Rogue!")
			}
			player.Spec = playerSpec
		},
	)
}

type Rogue struct {
	core.Character

	Talents  proto.RogueTalents
	Options  proto.Rogue_Options
	Rotation proto.Rogue_Rotation

	comboPoints int32

	deathmantle4pcProc bool

	builderEnergyCost float64
	newBuilder        func(sim *core.Simulation, target *core.Target) *core.SimpleSpell

	sinisterStrikeTemplate core.SimpleSpellTemplate
	sinisterStrike         core.SimpleSpell

	castSliceAndDice func()

	eviscerateEnergyCost  float64
	eviscerateDamageCalcs []core.MeleeDamageCalculator
	eviscerateTemplate    core.SimpleSpellTemplate
	eviscerate            core.SimpleSpell
}

func (rogue *Rogue) GetCharacter() *core.Character {
	return &rogue.Character
}

func (rogue *Rogue) GetRogue() *Rogue {
	return rogue
}

func (rogue *Rogue) AddRaidBuffs(raidBuffs *proto.RaidBuffs)    {}
func (rogue *Rogue) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {}

func (rogue *Rogue) Init(sim *core.Simulation) {
	// Precompute all the spell templates.
	rogue.sinisterStrikeTemplate = rogue.newSinisterStrikeTemplate(sim)

	rogue.initSliceAndDice(sim)
	rogue.eviscerateTemplate = rogue.newEviscerateTemplate(sim)
}

func (rogue *Rogue) Reset(sim *core.Simulation) {
	rogue.comboPoints = 0
	rogue.deathmantle4pcProc = false
}

func (rogue *Rogue) AddComboPoint(sim *core.Simulation) {
	if rogue.comboPoints == 5 {
		if sim.Log != nil {
			rogue.Log(sim, "Failed to gain 1 combo point, already full")
		}
	} else {
		if sim.Log != nil {
			rogue.Log(sim, "Gained 1 combo point (%d --> %d)", rogue.comboPoints, rogue.comboPoints+1)
		}
		rogue.comboPoints++
	}
}

func (rogue *Rogue) SpendComboPoints(sim *core.Simulation) {
	if sim.Log != nil {
		rogue.Log(sim, "Spent all combo points.")
	}
	rogue.comboPoints = 0
}

func NewRogue(character core.Character, options proto.Player) *Rogue {
	rogueOptions := options.GetRogue()

	rogue := &Rogue{
		Character: character,
		Talents:   *rogueOptions.Talents,
		Options:   *rogueOptions.Options,
		Rotation:  *rogueOptions.Rotation,
	}

	rogue.builderEnergyCost = rogue.SinisterStrikeEnergyCost()
	rogue.newBuilder = func(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
		return rogue.NewSinisterStrike(sim, target)
	}

	maxEnergy := 100.0
	if rogue.Talents.Vigor {
		maxEnergy = 110
	}
	rogue.EnableEnergyBar(maxEnergy, func(sim *core.Simulation) {
		if !rogue.IsOnCD(core.GCDCooldownID, sim.CurrentTime) {
			rogue.doRotation(sim)
		}
	})
	rogue.EnableAutoAttacks(rogue, core.AutoAttackOptions{
		MainHand:       rogue.WeaponFromMainHand(rogue.critMultiplier(true, false)),
		OffHand:        rogue.WeaponFromOffHand(rogue.critMultiplier(false, false)),
		AutoSwingMelee: true,
	})

	rogue.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Strength,
		ModifiedStat: stats.AttackPower,
		Modifier: func(strength float64, attackPower float64) float64 {
			return attackPower + strength*1
		},
	})

	rogue.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Agility,
		ModifiedStat: stats.AttackPower,
		Modifier: func(agility float64, attackPower float64) float64 {
			return attackPower + agility*1
		},
	})

	rogue.AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Agility,
		ModifiedStat: stats.MeleeCrit,
		Modifier: func(agility float64, meleeCrit float64) float64 {
			return meleeCrit + (agility/40)*core.MeleeCritRatingPerCritChance
		},
	})

	rogue.applyTalents()

	return rogue
}

func init() {
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceBloodElf, Class: proto.Class_ClassRogue}] = stats.Stats{
		stats.Strength:  92,
		stats.Agility:   160,
		stats.Stamina:   88,
		stats.Intellect: 43,
		stats.Spirit:    57,

		stats.AttackPower: 120,
		stats.MeleeCrit:   -0.3 * core.MeleeCritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceDwarf, Class: proto.Class_ClassRogue}] = stats.Stats{
		stats.Strength:  97,
		stats.Agility:   154,
		stats.Stamina:   92,
		stats.Intellect: 38,
		stats.Spirit:    57,

		stats.AttackPower: 120,
		stats.MeleeCrit:   -0.3 * core.MeleeCritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceGnome, Class: proto.Class_ClassRogue}] = stats.Stats{
		stats.Strength:  90,
		stats.Agility:   161,
		stats.Stamina:   88,
		stats.Intellect: 45,
		stats.Spirit:    58,

		stats.AttackPower: 120,
		stats.MeleeCrit:   -0.3 * core.MeleeCritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceHuman, Class: proto.Class_ClassRogue}] = stats.Stats{
		stats.Strength:  95,
		stats.Agility:   158,
		stats.Stamina:   89,
		stats.Intellect: 39,
		stats.Spirit:    58,

		stats.AttackPower: 120,
		stats.MeleeCrit:   -0.3 * core.MeleeCritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceNightElf, Class: proto.Class_ClassRogue}] = stats.Stats{
		stats.Strength:  92,
		stats.Agility:   163,
		stats.Stamina:   88,
		stats.Intellect: 39,
		stats.Spirit:    58,

		stats.AttackPower: 120,
		stats.MeleeCrit:   -0.3 * core.MeleeCritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceOrc, Class: proto.Class_ClassRogue}] = stats.Stats{
		stats.Strength:  98,
		stats.Agility:   155,
		stats.Stamina:   91,
		stats.Intellect: 36,
		stats.Spirit:    61,

		stats.AttackPower: 120,
		stats.MeleeCrit:   -0.3 * core.MeleeCritRatingPerCritChance,
	}
	trollStats := stats.Stats{
		stats.Strength:  96,
		stats.Agility:   160,
		stats.Stamina:   90,
		stats.Intellect: 35,
		stats.Spirit:    59,

		stats.AttackPower: 120,
		stats.MeleeCrit:   -0.3 * core.MeleeCritRatingPerCritChance,
	}
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll10, Class: proto.Class_ClassRogue}] = trollStats
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceTroll30, Class: proto.Class_ClassRogue}] = trollStats
	core.BaseStats[core.BaseStatsKey{Race: proto.Race_RaceUndead, Class: proto.Class_ClassRogue}] = stats.Stats{
		stats.Strength:  94,
		stats.Agility:   156,
		stats.Stamina:   90,
		stats.Intellect: 37,
		stats.Spirit:    63,

		stats.AttackPower: 120,
		stats.MeleeCrit:   -0.3 * core.MeleeCritRatingPerCritChance,
	}
}

// Agent is a generic way to access underlying rogue on any of the agents.
type RogueAgent interface {
	GetRogue() *Rogue
}
