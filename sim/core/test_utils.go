package core

import (
	"log"
	"testing"

	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Use same seed to get same result on every run.
const RSeed = int64(1)

const ShortDuration = 60
const LongDuration = 300

func MakeSingleTargetOptions(debuffs proto.Debuffs) Options {
	debuffsCopy := debuffs
	return Options{
		Encounter: NewEncounter(proto.Encounter{
			Targets: []*proto.Target{
				&proto.Target{
					Debuffs: &debuffsCopy,
				},
			},
		}),
	}
}

func CharacterStatsTest(label string, t *testing.T, params IndividualParams, expectedStats stats.Stats) {
	agent := NewAgent(params)
	raid := NewRaid(params.Buffs)
	raid.AddPlayer(agent)
	raid.Finalize()

	finalStats := agent.GetCharacter().GetStats()

	const tolerance = 0.5
	if !finalStats.EqualsWithTolerance(expectedStats, tolerance) {
		t.Fatalf("%s failed: CharacterStats() = %v, expected %v", label, finalStats, expectedStats)
	}
}

func StatWeightsTest(label string, t *testing.T, params IndividualParams, statsToTest []stats.Stat, referenceStat stats.Stat, expectedStatWeights stats.Stats) {
	params.Options.Encounter.Duration = LongDuration
	params.Options.Iterations = 5000
	params.Options.RSeed = RSeed
	params.Options.Debug = false

	results := CalcStatWeight(params, statsToTest, referenceStat)

	const tolerance = 0.05
	if !results.Weights.EqualsWithTolerance(expectedStatWeights, tolerance) {
		t.Fatalf("%s failed: CalcStatWeight() = %v, expected %v", label, results.Weights, expectedStatWeights)
	}
}

// Performs a basic end-to-end test of the simulator.
//   This is where we can add more sophisticated checks if we would like.
//   Any changes to the damage output of an item set
func IndividualSimTest(label string, t *testing.T, params IndividualParams, expectedDps float64) {
	sim := NewIndividualSim(params)
	result := sim.Run()

	if params.Options.Debug {
		log.Printf("LOGS:\n%s\n", result.Logs)
	}

	tolerance := 0.5
	if result.Agents[0].DpsAvg < expectedDps-tolerance || result.Agents[0].DpsAvg > expectedDps+tolerance {
		t.Fatalf("%s failed: expected %0f dps from sim but was %0f", label, expectedDps, result.Agents[0].DpsAvg)
	}
}

func IndividualSimAverageTest(label string, t *testing.T, params IndividualParams, expectedDps float64) {
	params.Options.Encounter.Duration = LongDuration
	params.Options.Iterations = 10000
	params.Options.RSeed = RSeed
	params.Options.Debug = false

	sim := NewIndividualSim(params)
	result := sim.Run()

	if params.Options.Debug {
		log.Printf("LOGS:\n%s\n", result.Logs)
	}

	tolerance := 0.5
	if result.Agents[0].DpsAvg < expectedDps-tolerance || result.Agents[0].DpsAvg > expectedDps+tolerance {
		t.Fatalf("%s failed: expected %0f dps from sim but was %0f", label, expectedDps, result.Agents[0].DpsAvg)
	}
}

type AllEncountersTestOptions struct {
	Label string
	T     *testing.T

	Options  Options
	Gear     items.EquipmentSpec
	Buffs    proto.Buffs
	Consumes proto.Consumes
	Race     proto.Race
	Class    proto.Class

	PlayerOptions *proto.PlayerOptions

	ExpectedDpsShort float64
	ExpectedDpsLong  float64
}

func IndividualSimAllEncountersTest(testOpts AllEncountersTestOptions) {
	params := IndividualParams{
		Equip:    testOpts.Gear,
		Race:     testOpts.Race,
		Class:    testOpts.Class,
		Consumes: testOpts.Consumes,
		Buffs:    testOpts.Buffs,
		Options:  testOpts.Options,

		PlayerOptions: testOpts.PlayerOptions,
		CustomStats:   stats.Stats{},
	}

	params.Options.Iterations = 1
	params.Options.RSeed = RSeed
	params.Options.Debug = false

	params.Options.Encounter.Duration = ShortDuration
	IndividualSimTest(
		testOpts.Label+"-short",
		testOpts.T,
		params,
		testOpts.ExpectedDpsShort)

	params.Options.Encounter.Duration = LongDuration
	IndividualSimTest(
		testOpts.Label+"-long",
		testOpts.T,
		params,
		testOpts.ExpectedDpsLong)
}

func IndividualBenchmark(b *testing.B, params IndividualParams) {
	params.Options.Encounter.Duration = LongDuration
	params.Options.Iterations = 1000
	params.Options.RSeed = RSeed
	params.Options.Debug = false

	for i := 0; i < b.N; i++ {
		sim := NewIndividualSim(params)
		sim.Run()
	}
}
