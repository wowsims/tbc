package shaman

import (
	"log"
	"testing"

	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	RegisterElementalShaman()
}

// TODO:
//  1. How to handle buffs that modify stats based on stats? Kings, Unrelenting Storms, etc.
//		Possible: Add a function on player like 'AddStats' and a 'onstatbuff' aura effect?

func TestSimulatePreRaidNoBuffs(t *testing.T) {
	simAllEncountersTest(AllEncountersTestOptions{
		label: "preRaid",
		t:     t,

		Options: BasicOptions,
		// no consumes
		Buffs: BasicBuffs,
		Race:  core.RaceBonusTypeTroll10,

		PlayerOptions: &PlayerOptionsAdaptive,
		Gear:          PreRaidGear,

		ExpectedDpsShort: 867,
		ExpectedDpsLong:  269,
	})
}

func TestSimulatePreRaid(t *testing.T) {
	simAllEncountersTest(AllEncountersTestOptions{
		label: "preRaid",
		t:     t,

		Options:  BasicOptions,
		Consumes: FullConsumes,
		Buffs:    FullBuffs,
		Race:     core.RaceBonusTypeOrc,

		PlayerOptions: &PlayerOptionsAdaptive,
		Gear:          PreRaidGear,

		ExpectedDpsShort: 1398.5,
		ExpectedDpsLong:  1096.3,
	})
}

func TestSimulateP1(t *testing.T) {
	simAllEncountersTest(AllEncountersTestOptions{
		label: "phase1",
		t:     t,

		Options:  BasicOptions,
		Consumes: FullConsumes,
		Buffs:    FullBuffs,
		Race:     core.RaceBonusTypeOrc,

		PlayerOptions: &PlayerOptionsAdaptive,
		Gear:          P1Gear,

		ExpectedDpsShort: 1539.5,
		ExpectedDpsLong:  1260.3,
	})
}

// func TestMultiTarget(t *testing.T) {
// 	doSimulateTest(
// 		"multiTarget",
// 		t,
// 		makeOptions(
// 			FullOptions,
// 			Encounter{
// 				Duration:     300,
// 				NumClTargets: 3,
// 			},
// 			AGENT_TYPE_ADAPTIVE),
// 		p1Gear,
//      1533.5)
// }

func TestLBOnlyAgent(t *testing.T) {
	simAllEncountersTest(AllEncountersTestOptions{
		label: "lbonly",
		t:     t,

		Options:  BasicOptions,
		Consumes: FullConsumes,
		Buffs:    FullBuffs,
		Race:     core.RaceBonusTypeOrc,

		PlayerOptions: &PlayerOptionsLBOnly,
		Gear:          P1Gear,

		ExpectedDpsShort: 1581.1,
		ExpectedDpsLong:  1271.9,
	})
}

// func TestFixedAgent(t *testing.T) {
// 	simAllEncountersTest(AllEncountersTestOptions{
// 		label: "fixedAgent",
// 		t:     t,

// 		Options:   FullOptions,
// 		Gear:      p1Gear,
// 		AgentType: AGENT_TYPE_FIXED_4LB_1CL,

// 		ExpectedDpsShort: 1489.3,
// 		ExpectedDpsLong:  1284.2,
// 	})
// }

func TestClearcastAgent(t *testing.T) {
	simAllEncountersTest(AllEncountersTestOptions{
		label: "clearcast",
		t:     t,

		Options:  BasicOptions,
		Consumes: FullConsumes,
		Buffs:    FullBuffs,
		Race:     core.RaceBonusTypeOrc,

		PlayerOptions: &PlayerOptionsCLOnClearcast,
		Gear:          P1Gear,

		ExpectedDpsShort: 1459.8,
		ExpectedDpsLong:  1221.8,
	})
}

func TestAverageDPS(t *testing.T) {
	eq := P1Gear

	options := BasicOptions
	options.Iterations = 10000
	options.Encounter = LongEncounter
	// options.Debug = true

	params := core.IndividualParams{
		Equip:         eq,
		Race:          core.RaceBonusTypeOrc,
		Consumes:      FullConsumes,
		Buffs:         FullBuffs,
		Options:       options,
		PlayerOptions: &PlayerOptionsAdaptive,
		CustomStats:   stats.Stats{},
	}

	sim := core.NewIndividualSim(params)
	result := sim.Run()

	log.Printf("result.DpsAvg: %0.1f", result.DpsAvg)
	log.Printf("LOGS:\n %s\n", result.Logs)
}

func BenchmarkSimulate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		options := BasicOptions
		options.Iterations = 1000
		options.Encounter = LongEncounter

		params := core.IndividualParams{
			Equip:    P1Gear,
			Race:     core.RaceBonusTypeOrc,
			Consumes: FullConsumes,
			Buffs:    FullBuffs,
			Options:  options,

			PlayerOptions: &PlayerOptionsAdaptive,
			CustomStats:   stats.Stats{},
		}
		sim := core.NewIndividualSim(params)
		sim.Run()
	}
}

type AllEncountersTestOptions struct {
	label string
	t     *testing.T

	Options  core.Options
	Gear     items.EquipmentSpec
	Buffs    core.Buffs
	Consumes core.Consumes
	Race     core.RaceBonusType

	PlayerOptions *proto.PlayerOptions

	ExpectedDpsShort float64
	ExpectedDpsLong  float64
}

func simAllEncountersTest(testOpts AllEncountersTestOptions) {
	params := core.IndividualParams{
		Equip:    testOpts.Gear,
		Race:     testOpts.Race,
		Consumes: testOpts.Consumes,
		Buffs:    testOpts.Buffs,
		Options:  makeOptions(testOpts.Options, ShortEncounter),

		PlayerOptions: testOpts.PlayerOptions,
		CustomStats:   stats.Stats{},
	}
	doSimulateTest(
		testOpts.label+"-short",
		testOpts.t,
		params,
		testOpts.ExpectedDpsShort)

	params.Options = makeOptions(testOpts.Options, LongEncounter)
	doSimulateTest(
		testOpts.label+"-long",
		testOpts.t,
		params,
		testOpts.ExpectedDpsLong)
}

// Performs a basic end-to-end test of the simulator.
//   This is where we can add more sophisticated checks if we would like.
//   Any changes to the damage output of an item set
func doSimulateTest(label string, t *testing.T, params core.IndividualParams, expectedDps float64) {
	// params.Options.Debug = true
	// params.Options.Iterations = 1

	sim := core.NewIndividualSim(params)
	result := sim.Run()

	log.Printf("LOGS:\n%s\n", result.Logs)
	tolerance := 0.5
	if result.DpsAvg < expectedDps-tolerance || result.DpsAvg > expectedDps+tolerance {
		t.Fatalf("%s failed: expected %0f dps from sim but was %0f", label, expectedDps, result.DpsAvg)
	}
}

func makeOptions(baseOptions core.Options, encounter core.Encounter) core.Options {
	baseOptions.Encounter = encounter
	return baseOptions
}
