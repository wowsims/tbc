package core

import (
	"log"
	"testing"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

const ShortDuration = 60
const LongDuration = 300

type IndividualSimInputs struct {
	SimOptions      *proto.SimOptions
	Gear            *proto.EquipmentSpec
	RaidBuffs       *proto.RaidBuffs
	PartyBuffs      *proto.PartyBuffs
	IndividualBuffs *proto.IndividualBuffs
	Consumes        *proto.Consumes
	Race            proto.Race
	Class           proto.Class

	Duration int

	// Convenience field if only 1 target is desired
	Target  *proto.Target
	Targets []*proto.Target

	PlayerOptions *proto.PlayerOptions
}

func NewIndividualSimRequest(inputs IndividualSimInputs) *proto.IndividualSimRequest {
	isr := &proto.IndividualSimRequest{
		Player: &proto.Player{
			Equipment: inputs.Gear,
			Options:   inputs.PlayerOptions,
		},

		RaidBuffs:       inputs.RaidBuffs,
		PartyBuffs:      inputs.PartyBuffs,
		IndividualBuffs: inputs.IndividualBuffs,

		Encounter:  &proto.Encounter{},
		SimOptions: inputs.SimOptions,
	}

	if isr.Player.Options == nil {
		isr.Player.Options = &proto.PlayerOptions{}
	}
	isr.Player.Options.Race = inputs.Race
	isr.Player.Options.Class = inputs.Class
	isr.Player.Options.Consumes = inputs.Consumes

	isr.Encounter.Duration = float64(inputs.Duration)
	if inputs.Target != nil {
		isr.Encounter.Targets = []*proto.Target{inputs.Target}
	} else {
		isr.Encounter.Targets = inputs.Targets
	}

	if isr.SimOptions == nil {
		isr.SimOptions = &proto.SimOptions{}
	}
	isr.SimOptions.Iterations = 1
	isr.SimOptions.IsTest = true
	isr.SimOptions.Debug = false

	return isr
}

func CharacterStatsTest(label string, t *testing.T, isr *proto.IndividualSimRequest, expectedStats stats.Stats) {
	csr := &proto.ComputeStatsRequest{
		Player:          isr.Player,
		RaidBuffs:       isr.RaidBuffs,
		PartyBuffs:      isr.PartyBuffs,
		IndividualBuffs: isr.IndividualBuffs,
	}

	result := ComputeStats(csr)
	finalStats := stats.FromFloatArray(result.FinalStats)

	const tolerance = 0.5
	if !finalStats.EqualsWithTolerance(expectedStats, tolerance) {
		t.Fatalf("%s failed: CharacterStats() = %v, expected %v", label, finalStats, expectedStats)
	}
}

func StatWeightsTest(label string, t *testing.T, isr *proto.IndividualSimRequest, statsToTest []proto.Stat, referenceStat proto.Stat, expectedStatWeights stats.Stats) {
	isr.Encounter.Duration = LongDuration
	isr.SimOptions.Iterations = 5000

	swr := &proto.StatWeightsRequest{
		Options:         isr,
		StatsToWeigh:    statsToTest,
		EpReferenceStat: referenceStat,
	}

	result := StatWeights(swr)
	resultWeights := stats.FromFloatArray(result.Weights)

	const tolerance = 0.05
	if !resultWeights.EqualsWithTolerance(expectedStatWeights, tolerance) {
		t.Fatalf("%s failed: CalcStatWeight() = %v, expected %v", label, resultWeights, expectedStatWeights)
	}
}

// Performs a basic end-to-end test of the simulator.
//   This is where we can add more sophisticated checks if we would like.
//   Any changes to the damage output of an item set
func IndividualSimTest(label string, t *testing.T, isr *proto.IndividualSimRequest, expectedDps float64) {
	result := RunIndividualSim(isr)

	if isr.SimOptions.Debug {
		log.Printf("LOGS:\n%s\n", result.Logs)
	}

	tolerance := 0.5
	if result.PlayerMetrics.DpsAvg < expectedDps-tolerance || result.PlayerMetrics.DpsAvg > expectedDps+tolerance {
		t.Fatalf("%s failed: expected %0f dps from sim but was %0f", label, expectedDps, result.PlayerMetrics.DpsAvg)
	}
}

func IndividualSimAverageTest(label string, t *testing.T, isr *proto.IndividualSimRequest, expectedDps float64) {
	isr.Encounter.Duration = LongDuration
	isr.SimOptions.Iterations = 10000

	result := RunIndividualSim(isr)

	if isr.SimOptions.Debug {
		log.Printf("LOGS:\n%s\n", result.Logs)
	}

	tolerance := 0.5
	if result.PlayerMetrics.DpsAvg < expectedDps-tolerance || result.PlayerMetrics.DpsAvg > expectedDps+tolerance {
		t.Fatalf("%s failed: expected %0f dps from sim but was %0f", label, expectedDps, result.PlayerMetrics.DpsAvg)
	}
}

type AllEncountersTestOptions struct {
	Label string
	T     *testing.T

	Inputs IndividualSimInputs

	ExpectedDpsShort float64
	ExpectedDpsLong  float64
}

func IndividualSimAllEncountersTest(testOpts AllEncountersTestOptions) {
	isr := NewIndividualSimRequest(testOpts.Inputs)
	// isr.SimOptions.Debug = true

	isr.Encounter.Duration = ShortDuration
	IndividualSimTest(
		testOpts.Label+"-short",
		testOpts.T,
		isr,
		testOpts.ExpectedDpsShort)

	isr.Encounter.Duration = LongDuration
	IndividualSimTest(
		testOpts.Label+"-long",
		testOpts.T,
		isr,
		testOpts.ExpectedDpsLong)
}

func IndividualBenchmark(b *testing.B, isr *proto.IndividualSimRequest) {
	isr.Encounter.Duration = LongDuration
	isr.SimOptions.Iterations = 1000

	// Set to false because IsTest adds a lot of computation.
	isr.SimOptions.IsTest = false

	for i := 0; i < b.N; i++ {
		sim := NewIndividualSim(*isr)
		sim.Run()
	}
}
