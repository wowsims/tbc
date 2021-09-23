package runner

import (
	"log"
	"testing"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/shaman"
)

// TODO:
//  1. How to handle buffs that modify stats based on stats? Kings, Unrelenting Storms, etc.
//		Possible: Add a function on player like 'AddStats' and a 'onstatbuff' aura effect?

// Use same seed to get same result on every run.
var RSeed = int64(1)

var shortEncounter = core.Encounter{
	Duration:   60,
	NumTargets: 1,
}
var longEncounter = core.Encounter{
	Duration:   300,
	NumTargets: 1,
}

var basicOptions = core.Options{
	Iterations: 1,
	RSeed:      RSeed,
	Debug:      false,
}

var basicBuffs = core.Buffs{
	ArcaneInt:                false,
	GiftOfTheWild:            false,
	BlessingOfKings:          false,
	ImprovedBlessingOfWisdom: false,
	JudgementOfWisdom:        false,
	Moonkin:                  false,
	SpriestDPS:               0,
	Bloodlust:                1,
}

var shamTalents = shaman.Talents{
	LightningMastery:   5,
	LightningOverload:  5,
	ElementalPrecision: 3,
	NaturesGuidance:    3,
	TidalMastery:       5,
	ElementalMastery:   true,
	UnrelentingStorm:   3,
	CallOfThunder:      5,
	Concussion:         5,
	Convection:         5,
}

var shamTotems = shaman.Totems{
	WrathOfAir:   true,
	TotemOfWrath: 1,
	ManaStream:   true,
}

var fullBuffs = core.Buffs{
	ArcaneInt:                true,
	GiftOfTheWild:            true,
	BlessingOfKings:          true,
	ImprovedBlessingOfWisdom: true,
	JudgementOfWisdom:        true,
	Moonkin:                  true,
	SpriestDPS:               500,
	Bloodlust:                1,
	// Misery:                   true,
}

var fullConsumes = core.Consumes{
	FlaskOfBlindingLight: true,
	BrilliantWizardOil:   true,
	MajorMageblood:       false,
	BlackendBasilisk:     true,
	DestructionPotion:    true,
	SuperManaPotion:      true,
	DarkRune:             true,
	DrumsOfBattle:        true,
}

var preRaidGear = []string{
	"Tidefury Helm",
	"Brooch of Heightened Potential",
	"Tidefury Shoulderguards",
	"Cloak of the Black Void",
	"Tidefury Chestpiece",
	"Shattrath Wraps",
	"Tidefury Gauntlets",
	"Moonrage Girdle",
	"Tidefury Kilt",
	"Earthbreaker's Greaves",
	"Seal of the Exorcist",
	"Spectral Band of Innervation",
	"Xi'ri's Gift",
	"Quagmirran's Eye",
	"Totem of the Void",
	"Sky Breaker",
	"Silvermoon Crest Shield",
}

var p1Gear = []string{
	"Cyclone Faceguard (Tier 4)",
	"Adornment of Stolen Souls",
	"Cyclone Shoulderguards (Tier 4)",
	"Ruby Drape of the Mysticant",
	"Netherstrike Breastplate",
	"Netherstrike Bracers",
	"Soul-Eater's Handwraps",
	"Netherstrike Belt",
	"Stormsong Kilt",
	"Windshear Boots",
	"Ring of Unrelenting Storms",
	"Ring of Recurrence",
	"The Lightning Capacitor",
	"Icon of the Silver Crescent",
	"Totem of the Void",
	"Nathrezim Mindblade",
	"Mazthoril Honor Shield",
}

// func TestSimulatePreRaidNoBuffs(t *testing.T) {
// 	simAllEncountersTest(AllEncountersTestOptions{
// 		label: "preRaidNoBuffs",
// 		t:     t,

// 		Options:   basicOptions,
// 		Gear:      preRaidGear,
// 		AgentType: AGENT_TYPE_ADAPTIVE,

// 		ExpectedDpsShort: 867,
// 		ExpectedDpsLong:  277,
// 	})
// }

func gearFromStrings(gears []string) core.EquipmentSpec {
	eq := core.EquipmentSpec{}
	for i, gear := range gears {
		item := core.ItemsByName[gear]
		eq[i].ID = item.ID
	}
	return eq
}

func TestSimulatePreRaid(t *testing.T) {
	simAllEncountersTest(AllEncountersTestOptions{
		label: "preRaid",
		t:     t,

		Options:         basicOptions,
		Gear:            gearFromStrings(preRaidGear),
		ShamanAgentType: shaman.AgentTypeAdaptive,
		Buffs:           fullBuffs,

		ExpectedDpsShort: 1406,
		ExpectedDpsLong:  1017,
	})
}

func TestSimulateP1(t *testing.T) {

	simAllEncountersTest(AllEncountersTestOptions{
		label: "phase1",
		t:     t,

		Options:         basicOptions,
		Gear:            gearFromStrings(p1Gear),
		ShamanAgentType: shaman.AgentTypeAdaptive,
		Buffs:           fullBuffs,

		ExpectedDpsShort: 1527,
		ExpectedDpsLong:  1226.6,
	})
}

// func TestMultiTarget(t *testing.T) {
// 	doSimulateTest(
// 		"multiTarget",
// 		t,
// 		makeOptions(
// 			fullOptions,
// 			Encounter{
// 				Duration:     300,
// 				NumClTargets: 3,
// 			},
// 			AGENT_TYPE_ADAPTIVE),
// 		p1Gear,
// 		1678.5)
// }

// func TestLBOnlyAgent(t *testing.T) {
// 	simAllEncountersTest(AllEncountersTestOptions{
// 		label: "lbOnly",
// 		t:     t,

// 		Options:   fullOptions,
// 		Gear:      p1Gear,
// 		AgentType: AGENT_TYPE_FIXED_LB_ONLY,

// 		ExpectedDpsShort: 1581.1,
// 		ExpectedDpsLong:  1227.6,
// 	})
// }

// func TestFixedAgent(t *testing.T) {
// 	simAllEncountersTest(AllEncountersTestOptions{
// 		label: "fixedAgent",
// 		t:     t,

// 		Options:   fullOptions,
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

		Options:         basicOptions,
		Gear:            gearFromStrings(p1Gear),
		ShamanAgentType: shaman.AgentTypeCLOnClearcast,
		Buffs:           fullBuffs,

		ExpectedDpsShort: 1468.4,
		ExpectedDpsLong:  1214.2,
	})
}

func TestAverageDPS(t *testing.T) {
	eq := gearFromStrings(p1Gear)
	player := core.NewPlayer(eq, core.RaceBonusTypeOrc, fullConsumes)
	party := &core.Party{Players: []core.PlayerAgent{{Player: player}}}
	raid := &core.Raid{Parties: []*core.Party{party}}
	party.Players[0].Agent = shaman.NewShaman(player, party, shamTalents, shamTotems, shaman.AgentTypeAdaptive)

	options := basicOptions
	options.Iterations = 5
	options.Encounter = shortEncounter
	buffs := fullBuffs
	options.Debug = true

	sim := SetupSim(raid, buffs, options)
	result := RunIndividualSim(sim, options)

	log.Printf("result.DpsAvg: %0.1f", result.DpsAvg)
	log.Printf("LOGS:\n %s\n", result.Logs)
}

// func BenchmarkSimulate(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		RunSimulation(SimRequest{
// 			Options:     fullOptions,
// 			Gear:        p1Gear,
// 			Iterations:  1000,
// 			IncludeLogs: false,
// 		})
// 	}
// }

type AllEncountersTestOptions struct {
	label string
	t     *testing.T

	Options core.Options
	Gear    core.EquipmentSpec
	Buffs   core.Buffs

	ShamanAgentType shaman.AgentType

	ExpectedDpsShort float64
	ExpectedDpsLong  float64
}

func simAllEncountersTest(testOpts AllEncountersTestOptions) {
	doSimulateTest(
		testOpts.label+"-short",
		testOpts.t,
		testOpts.ShamanAgentType,
		testOpts.Gear,
		makeOptions(testOpts.Options, shortEncounter),
		testOpts.Buffs,
		testOpts.ExpectedDpsShort)

	// doSimulateTest(
	// 	testOpts.label+"-long",
	// 	testOpts.t,
	// 	testOpts.Gear,
	// 	makeOptions(testOpts.Options, longEncounter),
	// 	testOpts.Buffs,
	// 	testOpts.ExpectedDpsLong)
}

// Performs a basic end-to-end test of the simulator.
//   This is where we can add more sophisticated checks if we would like.
//   Any changes to the damage output of an item set
func doSimulateTest(label string, t *testing.T, agent shaman.AgentType, eq core.EquipmentSpec, options core.Options, buffs core.Buffs, expectedDps float64) {
	player := core.NewPlayer(eq, core.RaceBonusTypeOrc, fullConsumes)
	party := &core.Party{Players: []core.PlayerAgent{{Player: player}}}
	raid := &core.Raid{Parties: []*core.Party{party}}
	party.Players[0].Agent = shaman.NewShaman(player, party, shamTalents, shamTotems, agent)

	options.Debug = true

	sim := SetupSim(raid, buffs, options)
	result := RunIndividualSim(sim, options)

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
