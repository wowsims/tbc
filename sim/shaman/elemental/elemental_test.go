package elemental

import (
	"testing"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	RegisterElementalShaman()
}

func TestP1FullCharacterStats(t *testing.T) {
	params := core.IndividualParams{
		Equip:       P1Gear,
		Race:        core.RaceBonusTypeTroll10,
		Class:       proto.Class_ClassShaman,
		Consumes:    FullConsumes,
		Buffs:       FullBuffs,
		PlayerOptions: &PlayerOptionsAdaptive,
		CustomStats: stats.Stats{},
	}

	core.CharacterStatsTest("p1Full", t, params, stats.Stats{
		stats.Strength:   20.8,
		stats.Agility:    20.8,
		stats.Stamina:    347.5,
		stats.Intellect:  511.4,
		stats.Spirit:     191.3,

		stats.SpellPower:       989,
		stats.HealingPower:     690,
		stats.ArcaneSpellPower: 80,
		stats.HolySpellPower:   80,
		stats.NatureSpellPower: 123,

		stats.MP5:       336.7,
		stats.SpellHit:  73.8,
		stats.SpellCrit: 637.8,

		stats.Mana:  10349,
		stats.Armor: 9170,
	})
}

var StatsToTest = []stats.Stat{
	stats.SpellPower,
	stats.SpellHit,
	stats.Intellect,
	stats.SpellCrit,
}

var ReferenceStat = stats.SpellPower

func TestCalcStatWeight(t *testing.T) {
	params := core.IndividualParams{
		Equip:       P1Gear,
		Race:        core.RaceBonusTypeTroll10,
		Class:       proto.Class_ClassShaman,
		Consumes:    FullConsumes,
		Buffs:       FullBuffs,
		PlayerOptions: &PlayerOptionsAdaptive,
		CustomStats: stats.Stats{},
	}

	core.StatWeightsTest("p1Full", t, params, StatsToTest, ReferenceStat, stats.Stats{
		stats.Intellect:  0.14,
		stats.SpellPower: 0.63,
		stats.SpellHit:   1.26,
		stats.SpellCrit:  0.46,
	})
}

// TODO:
//  1. How to handle buffs that modify stats based on stats? Kings, Unrelenting Storms, etc.
//		Possible: Add a function on player like 'AddStats' and a 'onstatbuff' aura effect?

func TestSimulatePreRaidNoBuffs(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "preRaid",
	  T:     t,

		// no consumes
		Buffs: BasicBuffs,
		Race:  core.RaceBonusTypeTroll10,
		Class: proto.Class_ClassShaman,

		PlayerOptions: &PlayerOptionsAdaptiveNoBuffs,
		Gear:          PreRaidGear,

		ExpectedDpsShort: 973.7,
		ExpectedDpsLong:  293.9,
	})
}

func TestSimulatePreRaid(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "preRaid",
	  T:     t,

		Consumes: FullConsumes,
		Buffs:    FullBuffs,
		Race:     core.RaceBonusTypeOrc,
		Class:    proto.Class_ClassShaman,

		PlayerOptions: &PlayerOptionsAdaptive,
		Gear:          PreRaidGear,

		ExpectedDpsShort: 1435.9,
		ExpectedDpsLong:  1078.5,
	})
}

func TestSimulateP1(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "phase1",
	  T:     t,

		Consumes: FullConsumes,
		Buffs:    FullBuffs,
		Race:     core.RaceBonusTypeOrc,
		Class:    proto.Class_ClassShaman,

		PlayerOptions: &PlayerOptionsAdaptive,
		Gear:          P1Gear,

		ExpectedDpsShort: 1385.0,
		ExpectedDpsLong:  1317.3,
	})
}

// func TestMultiTarget(t *testing.T) {
// 	params := core.IndividualParams{
// 		Equip:         P1Gear,
// 		Race:          core.RaceBonusTypeOrc,
//    Class:         proto.Class_ClassShaman,
// 		Consumes:      FullConsumes,
// 		Buffs:         FullBuffs,
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

		Consumes: FullConsumes,
		Buffs:    FullBuffs,
		Race:     core.RaceBonusTypeOrc,
		Class:    proto.Class_ClassShaman,

		PlayerOptions: &PlayerOptionsLBOnly,
		Gear:          P1Gear,

		ExpectedDpsShort: 1413.8,
		ExpectedDpsLong:  1205.8,
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

		Consumes: FullConsumes,
		Buffs:    FullBuffs,
		Race:     core.RaceBonusTypeOrc,
    Class:    proto.Class_ClassShaman,

		PlayerOptions: &PlayerOptionsCLOnClearcast,
		Gear:          P1Gear,

		ExpectedDpsShort: 1607.6,
		ExpectedDpsLong:  1315.1,
	})
}

func TestAverageDPS(t *testing.T) {
	params := core.IndividualParams{
		Equip:         P1Gear,
		Race:          core.RaceBonusTypeOrc,
    Class:         proto.Class_ClassShaman,
		Consumes:      FullConsumes,
		Buffs:         FullBuffs,
		PlayerOptions: &PlayerOptionsAdaptive,
		CustomStats:   stats.Stats{},
	}

	core.IndividualSimAverageTest("P1Average", t, params, 1248.1)
}

func BenchmarkSimulate(b *testing.B) {
	params := core.IndividualParams{
		Equip:    P1Gear,
		Race:     core.RaceBonusTypeOrc,
    Class:    proto.Class_ClassShaman,
		Consumes: FullConsumes,
		Buffs:    FullBuffs,

		PlayerOptions: &PlayerOptionsAdaptive,
		CustomStats:   stats.Stats{},
	}

	core.IndividualBenchmark(b, params)
}
