package core

import (
	"errors"
	"os"
	"testing"

	"github.com/wowsims/tbc/sim/core/proto"
	"google.golang.org/protobuf/encoding/prototext"
)

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
		DpsResults: make(map[string]*proto.DpsTestResult),
	}
}

type TestGenerator interface {
	// The total number of tests that this generator can generate.
	NumTests() int

	// The name and API request for the test with the given index.
	GetTest(testIdx int) (string, *proto.RaidSimRequest)
}

func RunTestSuite(t *testing.T, suiteName string, generator TestGenerator) {
	testSuite := NewIndividualTestSuite(suiteName)

	numTests := generator.NumTests()
	for i := 0; i < numTests; i++ {
		testName, rsr := generator.GetTest(i)
		testSuite.TestDPS(testName, rsr)
	}

	testSuite.Done(t)
}
