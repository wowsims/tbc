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

		RaidBuffs:       FullRaidBuffs,
		PartyBuffs:      FullPartyBuffs,
		IndividualBuffs: FullIndividualBuffs,

		PlayerOptions: PlayerOptionsAdaptive,
	})

	core.CharacterStatsTest("p1Full", t, isr, stats.Stats{
		stats.Strength:  27.4,
		stats.Agility:   27.4,
		stats.Stamina:   395.9,
		stats.Intellect: 531.2,
		stats.Spirit:    198.0,

		stats.SpellPower:       1109,
		stats.HealingPower:     705,
		stats.ArcaneSpellPower: 80,
		stats.HolySpellPower:   80,
		stats.NatureSpellPower: 123,

		stats.MP5:       337.9,
		stats.SpellHit:  87.8,
		stats.SpellCrit: 695.705,

		stats.Mana:  10646,
		stats.Armor: 9224.8,
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
		Gear:     P1Gear,
		Race:     proto.Race_RaceTroll10,
		Class:    proto.Class_ClassShaman,
		Consumes: FullConsumes,

		RaidBuffs:       FullRaidBuffs,
		PartyBuffs:      FullPartyBuffs,
		IndividualBuffs: FullIndividualBuffs,

		Target:        FullDebuffTarget,
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
			RaidBuffs:       BasicRaidBuffs,
			PartyBuffs:      BasicPartyBuffs,
			IndividualBuffs: BasicIndividualBuffs,

			Target: NoDebuffTarget,

			Race:  proto.Race_RaceTroll10,
			Class: proto.Class_ClassShaman,

			PlayerOptions: PlayerOptionsAdaptiveNoBuffs,
			Gear:          PreRaidGear,
		},

		ExpectedDpsShort: 828.0,
		ExpectedDpsLong:  261.2,
	})
}

func TestSimulatePreRaid(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "preRaid",
		T:     t,

		Inputs: core.IndividualSimInputs{
			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Consumes: FullConsumes,
			Target:   FullDebuffTarget,
			Race:     proto.Race_RaceOrc,
			Class:    proto.Class_ClassShaman,

			PlayerOptions: PlayerOptionsAdaptive,
			Gear:          PreRaidGear,
		},

		ExpectedDpsShort: 1429.1,
		ExpectedDpsLong:  1169.2,
	})
}

func TestSimulateP1(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "phase1",
		T:     t,

		Inputs: core.IndividualSimInputs{
			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Consumes: FullConsumes,
			Target:   FullDebuffTarget,
			Race:     proto.Race_RaceOrc,
			Class:    proto.Class_ClassShaman,

			PlayerOptions: PlayerOptionsAdaptive,
			Gear:          P1Gear,
		},

		ExpectedDpsShort: 1883.1,
		ExpectedDpsLong:  1625.6,
	})
}
func TestMultiTarget(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "multiTarget",
		T:     t,

		Inputs: core.IndividualSimInputs{
			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Consumes: FullConsumes,
			Targets: []*proto.Target{
				FullDebuffTarget,
				NoDebuffTarget,
				NoDebuffTarget,
				NoDebuffTarget,
			},
			Race:  proto.Race_RaceOrc,
			Class: proto.Class_ClassShaman,

			PlayerOptions: PlayerOptionsAdaptive,
			Gear:          P1Gear,
		},

		ExpectedDpsShort: 2497.0,
		ExpectedDpsLong:  1943.5,
	})

	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "multiTarget-tidefury",
		T:     t,

		Inputs: core.IndividualSimInputs{
			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Consumes: FullConsumes,
			Targets: []*proto.Target{
				FullDebuffTarget,
				NoDebuffTarget,
				NoDebuffTarget,
				NoDebuffTarget,
			},
			Race:  proto.Race_RaceOrc,
			Class: proto.Class_ClassShaman,

			PlayerOptions: PlayerOptionsAdaptive,
			Gear:          P1Tidefury,
		},

		ExpectedDpsShort: 2504.6,
		ExpectedDpsLong:  1930.4,
	})
}

func TestLBOnlyAgent(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "lbonly",
		T:     t,

		Inputs: core.IndividualSimInputs{
			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Consumes: FullConsumes,
			Target:   FullDebuffTarget,
			Race:     proto.Race_RaceOrc,
			Class:    proto.Class_ClassShaman,

			PlayerOptions: PlayerOptionsLBOnly,
			Gear:          P1Gear,
		},

		ExpectedDpsShort: 1880.7,
		ExpectedDpsLong:  1601.9,
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
			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Consumes: FullConsumes,
			Target:   FullDebuffTarget,
			Race:     proto.Race_RaceOrc,
			Class:    proto.Class_ClassShaman,

			PlayerOptions: PlayerOptionsCLOnClearcast,
			Gear:          P1Gear,
		},

		ExpectedDpsShort: 2023.1,
		ExpectedDpsLong:  1677.5,
	})
}

func TestAverageDPS(t *testing.T) {
	isr := core.NewIndividualSimRequest(core.IndividualSimInputs{
		Gear:     P1Gear,
		Race:     proto.Race_RaceOrc,
		Class:    proto.Class_ClassShaman,
		Consumes: FullConsumes,

		RaidBuffs:       FullRaidBuffs,
		PartyBuffs:      FullPartyBuffs,
		IndividualBuffs: FullIndividualBuffs,

		Target:        FullDebuffTarget,
		PlayerOptions: PlayerOptionsAdaptive,
	})

	core.IndividualSimAverageTest("P1Average", t, isr, 1622.65)
}

func BenchmarkSimulate(b *testing.B) {
	isr := core.NewIndividualSimRequest(core.IndividualSimInputs{
		Gear:     P1Gear,
		Race:     proto.Race_RaceOrc,
		Class:    proto.Class_ClassShaman,
		Consumes: FullConsumes,

		RaidBuffs:       FullRaidBuffs,
		PartyBuffs:      FullPartyBuffs,
		IndividualBuffs: FullIndividualBuffs,

		Target: FullDebuffTarget,

		PlayerOptions: PlayerOptionsAdaptive,
	})

	core.IndividualBenchmark(b, isr)
}
