package core

import (
	"log"
	"testing"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
	googleProto "google.golang.org/protobuf/proto"
)

var DefaultSimTestOptions = &proto.SimOptions{
	Iterations: 1,
	IsTest:     true,
	Debug:      false,
}
var AverageDefaultSimTestOptions = &proto.SimOptions{
	Iterations: 10000,
	IsTest:     true,
	Debug:      false,
}

const ShortDuration = 60
const LongDuration = 300

func MakeDefaultEncounterCombos(debuffs *proto.Debuffs) []EncounterCombo {
	var NoDebuffTarget = &proto.Target{
		Level:   73,
		Armor:   7700,
		MobType: proto.MobType_MobTypeBeast,
		Debuffs: &proto.Debuffs{},
	}

	var FullDebuffTarget = &proto.Target{
		Level:   73,
		Armor:   7700,
		MobType: proto.MobType_MobTypeDemon,
		Debuffs: debuffs,
	}

	return []EncounterCombo{
		EncounterCombo{
			Label: "LongSingleTargetNoDebuffs",
			Encounter: &proto.Encounter{
				Duration:          LongDuration,
				ExecuteProportion: 0.2,
				Targets: []*proto.Target{
					NoDebuffTarget,
				},
			},
		},
		EncounterCombo{
			Label: "ShortSingleTargetFullDebuffs",
			Encounter: &proto.Encounter{
				Duration:          ShortDuration,
				ExecuteProportion: 0.2,
				Targets: []*proto.Target{
					FullDebuffTarget,
				},
			},
		},
		EncounterCombo{
			Label: "LongSingleTargetFullDebuffs",
			Encounter: &proto.Encounter{
				Duration:          LongDuration,
				ExecuteProportion: 0.2,
				Targets: []*proto.Target{
					FullDebuffTarget,
				},
			},
		},
		EncounterCombo{
			Label: "LongMultiTarget",
			Encounter: &proto.Encounter{
				Duration:          LongDuration,
				ExecuteProportion: 0.2,
				Targets: []*proto.Target{
					FullDebuffTarget,
					FullDebuffTarget,
					FullDebuffTarget,
				},
			},
		},
	}
}

func MakeSingleTargetFullDebuffEncounter(debuffs *proto.Debuffs) *proto.Encounter {
	return &proto.Encounter{
		Duration:          LongDuration,
		ExecuteProportion: 0.2,
		Targets: []*proto.Target{
			&proto.Target{
				Level:   73,
				Armor:   7700,
				MobType: proto.MobType_MobTypeDemon,
				Debuffs: debuffs,
			},
		},
	}
}

// Returns default encounter combos, for testing average DPS.
// When doing average DPS tests we use a lot more iterations, so to save time
// we test fewer encounters.
func MakeAverageDefaultEncounterCombos(debuffs *proto.Debuffs) []EncounterCombo {
	return []EncounterCombo{
		EncounterCombo{
			Label:     "LongSingleTarget",
			Encounter: MakeSingleTargetFullDebuffEncounter(debuffs),
		},
	}
}

type IndividualSimInputs struct {
	Player          *proto.Player
	RaidBuffs       *proto.RaidBuffs
	PartyBuffs      *proto.PartyBuffs
	IndividualBuffs *proto.IndividualBuffs
	SimOptions      *proto.SimOptions

	Duration int

	// Convenience field if only 1 target is desired
	Target  *proto.Target
	Targets []*proto.Target
}

func NewIndividualSimRequest(inputs IndividualSimInputs) *proto.IndividualSimRequest {
	isr := &proto.IndividualSimRequest{
		Player:     inputs.Player,
		RaidBuffs:  inputs.RaidBuffs,
		PartyBuffs: inputs.PartyBuffs,

		Encounter:  &proto.Encounter{},
		SimOptions: inputs.SimOptions,
	}

	isr.Player.Buffs = inputs.IndividualBuffs

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
		Raid: SinglePlayerRaidProto(isr.Player, isr.PartyBuffs, isr.RaidBuffs),
	}

	result := ComputeStats(csr)
	finalStats := stats.FromFloatArray(result.RaidStats.Parties[0].Players[0].FinalStats)

	const tolerance = 0.5
	if !finalStats.EqualsWithTolerance(expectedStats, tolerance) {
		t.Fatalf("%s failed: CharacterStats() = %v, expected %v", label, finalStats, expectedStats)
	}
}

func StatWeightsTest(label string, t *testing.T, _swr *proto.StatWeightsRequest, expectedStatWeights stats.Stats) {
	// Make a copy so we can safely change fields.
	swr := googleProto.Clone(_swr).(*proto.StatWeightsRequest)

	swr.Encounter.Duration = LongDuration
	swr.SimOptions.Iterations = 5000

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

	tolerance := 0.5
	if result.PlayerMetrics.Dps.Avg < expectedDps-tolerance || result.PlayerMetrics.Dps.Avg > expectedDps+tolerance {
		// Automatically print output if we had debugging enabled.
		if isr.SimOptions.Debug {
			log.Printf("LOGS:\n%s\n", result.Logs)
		}
		t.Fatalf("%s failed: expected %0f dps from sim but was %0f", label, expectedDps, result.PlayerMetrics.Dps.Avg)
	}
}

func RaidSimTest(label string, t *testing.T, rsr *proto.RaidSimRequest, expectedDps float64) {
	result := RunRaidSim(rsr)

	tolerance := 0.5
	if result.RaidMetrics.Dps.Avg < expectedDps-tolerance || result.RaidMetrics.Dps.Avg > expectedDps+tolerance {
		// Automatically print output if we had debugging enabled.
		if rsr.SimOptions.Debug {
			log.Printf("LOGS:\n%s\n", result.Logs)
		}
		t.Fatalf("%s failed: expected %0f dps from sim but was %0f", label, expectedDps, result.RaidMetrics.Dps.Avg)
	}
}

func IndividualSimAverageTest(label string, t *testing.T, isr *proto.IndividualSimRequest, expectedDps float64) {
	isr.Encounter.Duration = LongDuration
	isr.SimOptions.Iterations = 10000
	// isr.SimOptions.Debug = true

	result := RunIndividualSim(isr)

	if isr.SimOptions.Debug {
		log.Printf("LOGS:\n%s\n", result.Logs)
	}

	tolerance := 0.5
	if result.PlayerMetrics.Dps.Avg < expectedDps-tolerance || result.PlayerMetrics.Dps.Avg > expectedDps+tolerance {
		t.Fatalf("%s failed: expected %0f dps from sim but was %0f", label, expectedDps, result.PlayerMetrics.Dps.Avg)
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
		RunIndividualSim(isr)
	}
}
