package elemental

import (
	"testing"

	_ "github.com/wowsims/tbc/sim/common"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	RegisterElementalShaman()
}

func TestP1FullCharacterStats(t *testing.T) {
	isr := core.NewIndividualSimRequest(core.IndividualSimInputs{
		Gear:     P1Gear,
		Race:     proto.Race_RaceTroll10,
		Class:    proto.Class_ClassShaman,
		Consumes: FullConsumes,
		Buffs:    FullBuffs,

		PlayerOptions: PlayerOptionsAdaptive,
	})

	core.CharacterStatsTest("p1Full", t, isr, stats.Stats{
		stats.Strength:   20.8,
		stats.Agility:    20.8,
		stats.Stamina:    395.9,
		stats.Intellect:  531.2,
		stats.Spirit:     198.0,

		stats.SpellPower:       1109,
		stats.HealingPower:     705,
		stats.ArcaneSpellPower: 80,
		stats.HolySpellPower:   80,
		stats.NatureSpellPower: 123,

		stats.MP5:       337.8,
		stats.SpellHit:  87.8,
		stats.SpellCrit: 696.4,

		stats.Mana:  10646,
		stats.Armor: 9211.6,
	})
}

var StatsToTest = []proto.Stat{
	proto.Stat_StatIntellect,
	proto.Stat_StatSpellPower,
	proto.Stat_StatSpellHit,
	proto.Stat_StatSpellCrit,
}

var ReferenceStat = proto.Stat_StatSpellPower

func TestCalcStatWeight(t *testing.T) {
	isr := core.NewIndividualSimRequest(core.IndividualSimInputs{
		Gear:        P1Gear,
		Race:        proto.Race_RaceTroll10,
		Class:       proto.Class_ClassShaman,
		Consumes:    FullConsumes,
		Buffs:       FullBuffs,
		Target:      FullDebuffTarget,
		PlayerOptions: PlayerOptionsAdaptive,
	})

	core.StatWeightsTest("p1Full", t, isr, StatsToTest, ReferenceStat, stats.Stats{
		stats.Intellect:  0.19,
		stats.SpellPower: 0.67,
		stats.SpellHit:   0.10,
		stats.SpellCrit:  0.59,
	})
}

// TODO:
//  1. How to handle buffs that modify stats based on stats? Kings, Unrelenting Storms, etc.
//		Possible: Add a function on player like 'AddStats' and a 'onstatbuff' aura effect?

func TestSimulatePreRaidNoBuffs(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "preRaid",
	  T:     t,

		Inputs: core.IndividualSimInputs{
			// no consumes
			Buffs:   BasicBuffs,
			Target:  NoDebuffTarget,

			Race:  proto.Race_RaceTroll10,
			Class: proto.Class_ClassShaman,

			PlayerOptions: PlayerOptionsAdaptiveNoBuffs,
			Gear:          PreRaidGear,
		},

		ExpectedDpsShort: 828.0,
		ExpectedDpsLong:  261.9,
	})
}

func TestSimulatePreRaid(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "preRaid",
	  T:     t,

		Inputs: core.IndividualSimInputs{
			Consumes: FullConsumes,
			Buffs:    FullBuffs,
			Target:   FullDebuffTarget,
			Race:     proto.Race_RaceOrc,
			Class:    proto.Class_ClassShaman,

			PlayerOptions: PlayerOptionsAdaptive,
			Gear:          PreRaidGear,
		},

		ExpectedDpsShort: 1297.5,
		ExpectedDpsLong:  1009.1,
	})
}

func TestSimulateP1(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "phase1",
	  T:     t,

		Inputs: core.IndividualSimInputs{
			Consumes: FullConsumes,
			Buffs:    FullBuffs,
			Target:   FullDebuffTarget,
			Race:     proto.Race_RaceOrc,
			Class:    proto.Class_ClassShaman,

			PlayerOptions: PlayerOptionsAdaptive,
			Gear:          P1Gear,
		},

		ExpectedDpsShort: 1928.0,
		ExpectedDpsLong:  1470.0,
	})
}

// func TestMultiTarget(t *testing.T) {
// 	params := core.IndividualParams{
// 		Equip:         P1Gear,
// 		Race:          proto.Race_RaceOrc,
//    Class:         proto.Class_ClassShaman,
// 		Consumes:      FullConsumes,
// 		Buffs:         FullBuffs,
//    Options:       FullDebuffOptions,
// 		Options:       makeOptions(core.BasicOptions, LongEncounter),
// 		PlayerOptions: &PlayerOptionsAdaptive,
// 	}
// 	params.Options.Encounter.NumTargets = 3

// 	doSimulateTest(
// 		"multiTarget",
// 		t,
// 		params,
// 		1533.5)
// }

func TestLBOnlyAgent(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "lbonly",
	  T:     t,

		Inputs: core.IndividualSimInputs{
			Consumes: FullConsumes,
			Buffs:    FullBuffs,
			Target:   FullDebuffTarget,
			Race:     proto.Race_RaceOrc,
			Class:    proto.Class_ClassShaman,

			PlayerOptions: PlayerOptionsLBOnly,
			Gear:          P1Gear,
		},

		ExpectedDpsShort: 1761.8,
		ExpectedDpsLong:  1483.9,
	})
}

// func TestFixedAgent(t *testing.T) {
// 	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
// 		Label: "fixedAgent",
// 	 T:     t,

// 		Options:   FullOptions,
// 		Gear:      p1Gear,
// 		AgentType: AGENT_TYPE_FIXED_4LB_1CL,

// 		ExpectedDpsShort: 1489.3,
// 		ExpectedDpsLong:  1284.2,
// 	})
// }

func TestClearcastAgent(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "clearcast",
	  T:     t,

		Inputs: core.IndividualSimInputs{
			Consumes: FullConsumes,
			Buffs:    FullBuffs,
			Target:   FullDebuffTarget,
			Race:     proto.Race_RaceOrc,
			Class:    proto.Class_ClassShaman,

			PlayerOptions: PlayerOptionsCLOnClearcast,
			Gear:          P1Gear,
		},

		ExpectedDpsShort: 1715.9,
		ExpectedDpsLong:  1526.6,
	})
}

func TestAverageDPS(t *testing.T) {
	isr := core.NewIndividualSimRequest(core.IndividualSimInputs{
		Gear:          P1Gear,
		Race:          proto.Race_RaceOrc,
    Class:         proto.Class_ClassShaman,
		Consumes:      FullConsumes,
		Buffs:         FullBuffs,
		Target:        FullDebuffTarget,
		PlayerOptions: PlayerOptionsAdaptive,
	})

	core.IndividualSimAverageTest("P1Average", t, isr, 1508.6)
}

func BenchmarkSimulate(b *testing.B) {
	isr := core.NewIndividualSimRequest(core.IndividualSimInputs{
		Gear:     P1Gear,
		Race:     proto.Race_RaceOrc,
    Class:    proto.Class_ClassShaman,
		Consumes: FullConsumes,
		Buffs:    FullBuffs,
		Target:   FullDebuffTarget,

		PlayerOptions: PlayerOptionsAdaptive,
	})

	core.IndividualBenchmark(b, isr)
}
