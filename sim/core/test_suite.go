package core

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
	"google.golang.org/protobuf/encoding/prototext"
)

type IndividualTestSuite struct {
	Name string

	// Names of all the tests, in the order they are tested.
	testNames []string

	testResults proto.TestSuiteResult
}

func NewIndividualTestSuite(suiteName string) *IndividualTestSuite {
	return &IndividualTestSuite{
		Name:        suiteName,
		testResults: newTestSuiteResult(),
	}
}

func (testSuite *IndividualTestSuite) TestCharacterStats(testName string, csr *proto.ComputeStatsRequest) {
	fullTestName := testSuite.Name + "-" + testName
	testSuite.testNames = append(testSuite.testNames, fullTestName)

	result := ComputeStats(csr)
	finalStats := stats.FromFloatArray(result.RaidStats.Parties[0].Players[0].FinalStats)

	testSuite.testResults.CharacterStatsResults[fullTestName] = &proto.CharacterStatsTestResult{
		FinalStats: finalStats[:],
	}
}

func (testSuite *IndividualTestSuite) TestStatWeights(testName string, swr *proto.StatWeightsRequest) {
	fullTestName := testSuite.Name + "-" + testName
	testSuite.testNames = append(testSuite.testNames, fullTestName)

	result := StatWeights(swr)
	weights := stats.FromFloatArray(result.Weights)

	testSuite.testResults.StatWeightsResults[fullTestName] = &proto.StatWeightsTestResult{
		Weights: weights[:],
	}
}

func (testSuite *IndividualTestSuite) TestDPS(testName string, rsr *proto.RaidSimRequest) {
	fullTestName := testSuite.Name + "-" + testName
	testSuite.testNames = append(testSuite.testNames, fullTestName)

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

const tolerance = 0.00001

func (testSuite *IndividualTestSuite) evaluateResults(t *testing.T) {
	expectedResults := testSuite.readExpectedResults()

	// Display results in order of testNames, to keep the same order as the tests.
	for _, testName := range testSuite.testNames {
		if actualCharacterStats, ok := testSuite.testResults.CharacterStatsResults[testName]; ok {
			actualStats := stats.FromFloatArray(actualCharacterStats.FinalStats)
			if expectedCharacterStats, ok := expectedResults.CharacterStatsResults[testName]; ok {
				expectedStats := stats.FromFloatArray(expectedCharacterStats.FinalStats)
				if !actualStats.EqualsWithTolerance(expectedStats, tolerance) {
					t.Errorf("%s failed: expected %v but was %v", testName, expectedStats, actualStats)
				}
			} else {
				t.Errorf("Unexpected test %s with stats: %v", testName, actualStats)
			}
		} else if actualStatWeights, ok := testSuite.testResults.StatWeightsResults[testName]; ok {
			actualWeights := stats.FromFloatArray(actualStatWeights.Weights)
			if expectedStatWeights, ok := expectedResults.StatWeightsResults[testName]; ok {
				expectedWeights := stats.FromFloatArray(expectedStatWeights.Weights)
				if !actualWeights.EqualsWithTolerance(expectedWeights, tolerance) {
					t.Errorf("%s failed: expected %v but was %v", testName, expectedWeights, actualWeights)
				}
			} else {
				t.Errorf("Unexpected test %s with stat weights: %v", testName, actualWeights)
			}
		} else if actualDpsResult, ok := testSuite.testResults.DpsResults[testName]; ok {
			if expectedDpsResult, ok := expectedResults.DpsResults[testName]; ok {
				if actualDpsResult.Dps < expectedDpsResult.Dps-tolerance || actualDpsResult.Dps > expectedDpsResult.Dps+tolerance {
					t.Errorf("%s failed: expected %0.03f but was %0.03f!.", testName, expectedDpsResult.Dps, actualDpsResult.Dps)
				}
			} else {
				t.Errorf("Unexpected test %s with %0.03f DPS!", testName, actualDpsResult.Dps)
			}
		}
	}

	for testName, expectedCharacterStats := range expectedResults.CharacterStatsResults {
		expectedStats := stats.FromFloatArray(expectedCharacterStats.FinalStats)
		if _, ok := testSuite.testResults.CharacterStatsResults[testName]; !ok {
			t.Errorf("%s missing (expected stats %v)!", testName, expectedStats)
		}
	}

	for testName, expectedStatWeights := range expectedResults.StatWeightsResults {
		expectedWeights := stats.FromFloatArray(expectedStatWeights.Weights)
		if _, ok := testSuite.testResults.StatWeightsResults[testName]; !ok {
			t.Errorf("%s missing (expected weights %v)!", testName, expectedWeights)
		}
	}

	for testName, expectedDpsResult := range expectedResults.DpsResults {
		if _, ok := testSuite.testResults.DpsResults[testName]; !ok {
			t.Errorf("%s missing (expected %0.03f DPS)!", testName, expectedDpsResult.Dps)
		}
	}
}

func (testSuite *IndividualTestSuite) writeToFile() {
	str := prototext.Format(&testSuite.testResults)
	// For some reason the formatter sometimes outputs 2 spaces instead of one.
	// Replace so we get consistent output.
	str = strings.ReplaceAll(str, "  ", " ")
	data := []byte(str)

	err := os.WriteFile(testSuite.Name+".results.tmp", data, 0644)
	if err != nil {
		panic(err)
	}
}

func (testSuite *IndividualTestSuite) readExpectedResults() proto.TestSuiteResult {
	data, err := os.ReadFile(testSuite.Name + ".results")
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
		CharacterStatsResults: make(map[string]*proto.CharacterStatsTestResult),
		StatWeightsResults:    make(map[string]*proto.StatWeightsTestResult),
		DpsResults:            make(map[string]*proto.DpsTestResult),
	}
}

type TestGenerator interface {
	// The total number of tests that this generator can generate.
	NumTests() int

	// The name and API request for the test with the given index.
	GetTest(testIdx int) (string, *proto.ComputeStatsRequest, *proto.StatWeightsRequest, *proto.RaidSimRequest)
}

func RunTestSuite(t *testing.T, suiteName string, generator TestGenerator) {
	testSuite := NewIndividualTestSuite(suiteName)

	numTests := generator.NumTests()
	for i := 0; i < numTests; i++ {
		testName, csr, swr, rsr := generator.GetTest(i)
		if csr != nil {
			testSuite.TestCharacterStats(testName, csr)
		} else if swr != nil {
			testSuite.TestStatWeights(testName, swr)
		} else if rsr != nil {
			testSuite.TestDPS(testName, rsr)
		} else {
			panic("No test request provided")
		}
	}

	testSuite.Done(t)

	if t.Failed() {
		t.Log("One or more tests failed! If the changes are intentional, update the expected results with 'make update-tests'. Otherwise go fix your bugs!")
	}
}
