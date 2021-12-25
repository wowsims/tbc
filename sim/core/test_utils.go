package core

import (
	"errors"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
	"google.golang.org/protobuf/encoding/prototext"
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
		MobType: proto.MobType_MobTypeDemon,
		Debuffs: &proto.Debuffs{},
	}

	var FullDebuffTarget = &proto.Target{
		Level:   73,
		Armor:   7700,
		MobType: proto.MobType_MobTypeBeast,
		Debuffs: debuffs,
	}

	return []EncounterCombo{
		EncounterCombo{
			Label: "ShortSingleTargetNoDebuffs",
			Encounter: &proto.Encounter{
				Duration: ShortDuration,
				Targets: []*proto.Target{
					NoDebuffTarget,
				},
			},
		},
		EncounterCombo{
			Label: "LongSingleTargetNoDebuffs",
			Encounter: &proto.Encounter{
				Duration: LongDuration,
				Targets: []*proto.Target{
					NoDebuffTarget,
				},
			},
		},
		EncounterCombo{
			Label: "ShortSingleTargetFullDebuffs",
			Encounter: &proto.Encounter{
				Duration: ShortDuration,
				Targets: []*proto.Target{
					FullDebuffTarget,
				},
			},
		},
		EncounterCombo{
			Label: "LongSingleTargetFullDebuffs",
			Encounter: &proto.Encounter{
				Duration: LongDuration,
				Targets: []*proto.Target{
					FullDebuffTarget,
				},
			},
		},
		EncounterCombo{
			Label: "ShortMultiTarget",
			Encounter: &proto.Encounter{
				Duration: ShortDuration,
				Targets: []*proto.Target{
					FullDebuffTarget,
					FullDebuffTarget,
					FullDebuffTarget,
				},
			},
		},
		EncounterCombo{
			Label: "LongMultiTarget",
			Encounter: &proto.Encounter{
				Duration: LongDuration,
				Targets: []*proto.Target{
					FullDebuffTarget,
					FullDebuffTarget,
					FullDebuffTarget,
				},
			},
		},
	}
}

// Returns default encounter combos, for testing average DPS.
// When doing average DPS tests we use a lot more iterations, so to save time
// we test fewer encounters.
func MakeAverageDefaultEncounterCombos(debuffs *proto.Debuffs) []EncounterCombo {
	var FullDebuffTarget = &proto.Target{
		Level:   73,
		Armor:   7700,
		MobType: proto.MobType_MobTypeBeast,
		Debuffs: debuffs,
	}

	return []EncounterCombo{
		EncounterCombo{
			Label: "LongSingleTarget",
			Encounter: &proto.Encounter{
				Duration: LongDuration,
				Targets: []*proto.Target{
					FullDebuffTarget,
				},
			},
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

type IndividualTestSuite struct {
	Name string

	testResults proto.TestSuiteResult
}

func NewIndividualTestSuite(suiteName string) *IndividualTestSuite {
	return &IndividualTestSuite{
		Name:        suiteName,
		testResults: newTestSuiteResult(),
	}
}

func (testSuite *IndividualTestSuite) TestDPS(testName string, rsr *proto.RaidSimRequest) {
	fullTestName := testSuite.Name + "-" + testName

	result := RunRaidSim(rsr)
	dps := result.RaidMetrics.Dps.Avg

	testSuite.testResults.DpsResults[fullTestName] = &proto.DpsTestResult{
		Dps: dps,
	}
}

func (testSuite *IndividualTestSuite) Done(t *testing.T) {
	testSuite.writeToFile()
	testSuite.evaluateResults(t)
}

const tolerance = 0.5

func (testSuite *IndividualTestSuite) evaluateResults(t *testing.T) {
	expectedResults := testSuite.readExpectedResults()

	for testName, expectedDpsResult := range expectedResults.DpsResults {
		if actualDpsResult, ok := testSuite.testResults.DpsResults[testName]; ok {
			if actualDpsResult.Dps < expectedDpsResult.Dps-tolerance || actualDpsResult.Dps > expectedDpsResult.Dps+tolerance {
				t.Errorf("%s failed: expected %0.03f but was %0.03f!.", testName, expectedDpsResult.Dps, actualDpsResult.Dps)
			}
		} else {
			t.Errorf("%s missing (expected %0.03f DPS)!", testName, expectedDpsResult.Dps)
		}
	}

	for testName, actualDpsResult := range testSuite.testResults.DpsResults {
		if _, ok := expectedResults.DpsResults[testName]; !ok {
			t.Errorf("Unexpected test %s with %0.03f DPS!", testName, actualDpsResult.Dps)
		}
	}
}

func (testSuite *IndividualTestSuite) writeToFile() {
	str := prototext.Format(&testSuite.testResults)
	data := []byte(str)

	err := os.WriteFile(testSuite.Name+".testresults.tmp", data, 0644)
	if err != nil {
		panic(err)
	}
}

func (testSuite *IndividualTestSuite) readExpectedResults() proto.TestSuiteResult {
	data, err := os.ReadFile(testSuite.Name + ".testresults")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return newTestSuiteResult()
		}

		panic(err)
	}

	results := &proto.TestSuiteResult{}
	if err = prototext.Unmarshal(data, results); err != nil {
		panic(err)
	}
	return *results
}

func newTestSuiteResult() proto.TestSuiteResult {
	return proto.TestSuiteResult{
		DpsResults: make(map[string]*proto.DpsTestResult),
	}
}

type RaceCombo struct {
	Label string
	Race  proto.Race
}
type GearSetCombo struct {
	Label   string
	GearSet *proto.EquipmentSpec
}
type SpecOptionsCombo struct {
	Label       string
	SpecOptions interface{}
}
type BuffsCombo struct {
	Label    string
	Raid     *proto.RaidBuffs
	Party    *proto.PartyBuffs
	Player   *proto.IndividualBuffs
	Consumes *proto.Consumes
}
type EncounterCombo struct {
	Label     string
	Encounter *proto.Encounter
}
type SettingsCombos struct {
	Class       proto.Class
	Races       []RaceCombo
	GearSets    []GearSetCombo
	SpecOptions []SpecOptionsCombo
	Buffs       []BuffsCombo
	Encounters  []EncounterCombo
	SimOptions  *proto.SimOptions
}

func (combos *SettingsCombos) NumCombos() int {
	return len(combos.Races) * len(combos.GearSets) * len(combos.SpecOptions) * len(combos.Buffs) * len(combos.Encounters)
}

func (combos *SettingsCombos) GetCombo(comboIdx int) (string, *proto.RaidSimRequest) {
	testNameParts := []string{}

	raceIdx := comboIdx % len(combos.Races)
	comboIdx /= len(combos.Races)
	raceCombo := combos.Races[raceIdx]
	testNameParts = append(testNameParts, raceCombo.Label)

	gearSetIdx := comboIdx % len(combos.GearSets)
	comboIdx /= len(combos.GearSets)
	gearSetCombo := combos.GearSets[gearSetIdx]
	testNameParts = append(testNameParts, gearSetCombo.Label)

	specOptionsIdx := comboIdx % len(combos.SpecOptions)
	comboIdx /= len(combos.SpecOptions)
	specOptionsCombo := combos.SpecOptions[specOptionsIdx]
	testNameParts = append(testNameParts, specOptionsCombo.Label)

	buffsIdx := comboIdx % len(combos.Buffs)
	comboIdx /= len(combos.Buffs)
	buffsCombo := combos.Buffs[buffsIdx]
	testNameParts = append(testNameParts, buffsCombo.Label)

	encounterIdx := comboIdx % len(combos.Encounters)
	comboIdx /= len(combos.Encounters)
	encounterCombo := combos.Encounters[encounterIdx]
	testNameParts = append(testNameParts, encounterCombo.Label)

	rsr := &proto.RaidSimRequest{
		Raid: SinglePlayerRaidProto(
			WithSpec(&proto.Player{
				Race:      raceCombo.Race,
				Class:     combos.Class,
				Equipment: gearSetCombo.GearSet,
				Consumes:  buffsCombo.Consumes,
				Buffs:     buffsCombo.Player,
			}, specOptionsCombo.SpecOptions),
			buffsCombo.Party,
			buffsCombo.Raid),
		Encounter:  encounterCombo.Encounter,
		SimOptions: combos.SimOptions,
	}

	return strings.Join(testNameParts, "-"), rsr
}

func TestSuiteAllSettingsCombos(t *testing.T, suiteName string, combos SettingsCombos) {
	testSuite := NewIndividualTestSuite(suiteName)

	numCombos := combos.NumCombos()
	for i := 0; i < numCombos; i++ {
		testName, rsr := combos.GetCombo(i)
		testSuite.TestDPS(testName, rsr)
	}

	testSuite.Done(t)
}
