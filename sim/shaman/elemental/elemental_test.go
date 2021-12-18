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
		Player: &proto.Player{
			Race:      proto.Race_RaceTroll10,
			Class:     proto.Class_ClassShaman,
			Equipment: P1Gear,
			Consumes:  FullConsumes,
			Spec:      PlayerOptionsAdaptive,
		},

		RaidBuffs:       FullRaidBuffs,
		PartyBuffs:      FullPartyBuffs,
		IndividualBuffs: FullIndividualBuffs,
	})

	core.CharacterStatsTest("p1Full", t, isr, stats.Stats{
		stats.Strength:  140.7,
		stats.Agility:   100.0,
		stats.Stamina:   522.4,
		stats.Intellect: 531.2,
		stats.Spirit:    182.5,

		stats.SpellPower:       1109,
		stats.HealingPower:     705,
		stats.ArcaneSpellPower: 80,
		stats.HolySpellPower:   80,
		stats.NatureSpellPower: 123,

		stats.MP5:       337.9,
		stats.SpellHit:  125.600,
		stats.SpellCrit: 695.705,

		stats.AttackPower: 401.4,
		stats.MeleeCrit:   118.1,

		stats.Mana:  10646,
		stats.Armor: 9370.0,
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
	//rsr := &proto.RaidSimRequest{
	//	Raid: SinglePlayerRaidProto(
	//			FullPartyBuffs,
	//			FullRaidBuffs),
	//	Encounter:  request.Encounter,
	//}

	swr := &proto.StatWeightsRequest{
		Player: &proto.Player{
			Race:      proto.Race_RaceTroll10,
			Class:     proto.Class_ClassShaman,
			Equipment: P1Gear,
			Consumes:  FullConsumes,
			Spec:      PlayerOptionsAdaptive,
			Buffs:     FullIndividualBuffs,
		},
		RaidBuffs:  FullRaidBuffs,
		PartyBuffs: FullPartyBuffs,
		Encounter: &proto.Encounter{
			Targets: []*proto.Target{
				FullDebuffTarget,
			},
		},
		StatsToWeigh:    StatsToTest,
		EpReferenceStat: ReferenceStat,
		SimOptions:      core.DefaultSimTestOptions,
	}

	core.StatWeightsTest("p1Full", t, swr, stats.Stats{
		stats.Intellect:  0.182,
		stats.SpellPower: 0.699,
		stats.SpellHit:   0.156,
		stats.SpellCrit:  0.580,
	})
}

func TestSimulatePreRaidNoBuffs(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "preRaid",
		T:     t,

		Inputs: core.IndividualSimInputs{
			Player: &proto.Player{
				Race:      proto.Race_RaceTroll10,
				Class:     proto.Class_ClassShaman,
				Equipment: PreRaidGear,
				// no consumes
				Spec: PlayerOptionsAdaptiveNoBuffs,
			},

			RaidBuffs:       BasicRaidBuffs,
			PartyBuffs:      BasicPartyBuffs,
			IndividualBuffs: BasicIndividualBuffs,

			Target: NoDebuffTarget,
		},

		ExpectedDpsShort: 990.2,
		ExpectedDpsLong:  418.5,
	})
}

func TestSimulatePreRaid(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "preRaid",
		T:     t,

		Inputs: core.IndividualSimInputs{
			Player: &proto.Player{
				Race:      proto.Race_RaceOrc,
				Class:     proto.Class_ClassShaman,
				Equipment: PreRaidGear,
				Consumes:  FullConsumes,
				Spec:      PlayerOptionsAdaptive,
			},

			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Target: FullDebuffTarget,
		},

		ExpectedDpsShort: 1557.8,
		ExpectedDpsLong:  1181.8,
	})
}

func TestSimulateP1(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "phase1",
		T:     t,

		Inputs: core.IndividualSimInputs{
			Player: &proto.Player{
				Race:      proto.Race_RaceOrc,
				Class:     proto.Class_ClassShaman,
				Equipment: P1Gear,
				Consumes:  FullConsumes,
				Spec:      PlayerOptionsAdaptive,
			},

			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Target: FullDebuffTarget,
		},

		ExpectedDpsShort: 2166.0,
		ExpectedDpsLong:  1603.1,
	})
}
func TestMultiTarget(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "multiTarget",
		T:     t,

		Inputs: core.IndividualSimInputs{
			Player: &proto.Player{
				Race:      proto.Race_RaceOrc,
				Class:     proto.Class_ClassShaman,
				Equipment: P1Gear,
				Consumes:  FullConsumes,
				Spec:      PlayerOptionsAdaptive,
			},

			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Targets: []*proto.Target{
				FullDebuffTarget,
				NoDebuffTarget,
				NoDebuffTarget,
				NoDebuffTarget,
			},
		},

		ExpectedDpsShort: 2684.5,
		ExpectedDpsLong:  2053.3,
	})

	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "multiTarget-tidefury",
		T:     t,

		Inputs: core.IndividualSimInputs{
			Player: &proto.Player{
				Race:      proto.Race_RaceOrc,
				Class:     proto.Class_ClassShaman,
				Equipment: P1Tidefury,
				Consumes:  FullConsumes,
				Spec:      PlayerOptionsAdaptive,
			},

			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Targets: []*proto.Target{
				FullDebuffTarget,
				NoDebuffTarget,
				NoDebuffTarget,
				NoDebuffTarget,
			},
		},

		ExpectedDpsShort: 2731.0,
		ExpectedDpsLong:  2093.3,
	})
}

func TestLBOnlyAgent(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "lbonly",
		T:     t,

		Inputs: core.IndividualSimInputs{
			Player: &proto.Player{
				Race:      proto.Race_RaceOrc,
				Class:     proto.Class_ClassShaman,
				Equipment: P1Gear,
				Consumes:  FullConsumes,
				Spec:      PlayerOptionsLBOnly,
			},

			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Target: FullDebuffTarget,
		},

		ExpectedDpsShort: 2041.9,
		ExpectedDpsLong:  1536.8,
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
			Player: &proto.Player{
				Race:      proto.Race_RaceOrc,
				Class:     proto.Class_ClassShaman,
				Equipment: P1Gear,
				Consumes:  FullConsumes,
				Spec:      PlayerOptionsCLOnClearcast,
			},

			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Target: FullDebuffTarget,
		},

		ExpectedDpsShort: 2135.3,
		ExpectedDpsLong:  1578.9,
	})
}

func TestAverageDPS(t *testing.T) {
	isr := core.NewIndividualSimRequest(core.IndividualSimInputs{
		Player: &proto.Player{
			Race:      proto.Race_RaceOrc,
			Class:     proto.Class_ClassShaman,
			Equipment: P1Gear,
			Consumes:  FullConsumes,
			Spec:      PlayerOptionsAdaptive,
		},

		RaidBuffs:       FullRaidBuffs,
		PartyBuffs:      FullPartyBuffs,
		IndividualBuffs: FullIndividualBuffs,

		Target: FullDebuffTarget,
	})

	core.IndividualSimAverageTest("P1Average", t, isr, 1593.9)
}

func BenchmarkSimulate(b *testing.B) {
	isr := core.NewIndividualSimRequest(core.IndividualSimInputs{
		Player: &proto.Player{
			Race:      proto.Race_RaceOrc,
			Class:     proto.Class_ClassShaman,
			Equipment: P1Gear,
			Consumes:  FullConsumes,
			Spec:      PlayerOptionsAdaptive,
		},

		RaidBuffs:       FullRaidBuffs,
		PartyBuffs:      FullPartyBuffs,
		IndividualBuffs: FullIndividualBuffs,

		Target: FullDebuffTarget,
	})

	core.IndividualBenchmark(b, isr)
}
