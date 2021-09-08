package core

import (
	"testing"
)

// Use same seed to get same result on every run.
var RSeed = int64(1)

var shortEncounter = Encounter{
	Duration:     60,
	NumClTargets: 1,
}
var longEncounter = Encounter{
	Duration:     300,
	NumClTargets: 1,
}

var basicOptions = Options{
	RSeed:        RSeed,
	NumBloodlust: 1,
	NumDrums:     0,
	Buffs: Buffs{
		ArcaneInt:                false,
		GiftOftheWild:            false,
		BlessingOfKings:          false,
		ImprovedBlessingOfWisdom: false,
		JudgementOfWisdom:        false,
		Moonkin:                  false,
		SpriestDPS:               0,
		WaterShield:              true,
		Race:                     RaceBonusTroll10,
	},
	Talents: Talents{
		LightningOverload:  5,
		ElementalPrecision: 3,
		NaturesGuidance:    3,
		TidalMastery:       5,
		ElementalMastery:   true,
		UnrelentingStorm:   3,
		CallOfThunder:      5,
		Concussion:         5,
		Convection:         5,
	},
}

var fullOptions = Options{
	RSeed:        RSeed,
	NumBloodlust: 1,
	NumDrums:     1,
	Buffs: Buffs{
		ArcaneInt:                true,
		GiftOftheWild:            true,
		BlessingOfKings:          true,
		ImprovedBlessingOfWisdom: true,
		JudgementOfWisdom:        true,
		Moonkin:                  true,
		SpriestDPS:               500,
		WaterShield:              true,
		Race:                     RaceBonusOrc,
	},
	Consumes: Consumes{
		FlaskOfBlindingLight: true,
		BrilliantWizardOil:   true,
		MajorMageblood:       false,
		BlackendBasilisk:     true,
		DestructionPotion:    true,
		SuperManaPotion:      true,
		DarkRune:             true,
	},
	Talents: Talents{
		LightningOverload:  5,
		ElementalPrecision: 3,
		NaturesGuidance:    3,
		TidalMastery:       5,
		ElementalMastery:   true,
		UnrelentingStorm:   3,
		CallOfThunder:      5,
		Concussion:         5,
		Convection:         5,
	},
	Totems: Totems{
		TotemOfWrath: 1,
		WrathOfAir:   true,
		ManaStream:   true,
	},
}

var preRaidGear = EquipmentSpec{
	{Name: "Tidefury Helm"},
	{Name: "Brooch of Heightened Potential"},
	{Name: "Tidefury Shoulderguards"},
	{Name: "Cloak of the Black Void"},
	{Name: "Tidefury Chestpiece"},
	{Name: "Shattrath Wraps"},
	{Name: "Tidefury Gauntlets"},
	{Name: "Moonrage Girdle"},
	{Name: "Tidefury Kilt"},
	{Name: "Earthbreaker's Greaves"},
	{Name: "Seal of the Exorcist"},
	{Name: "Spectral Band of Innervation"},
	{Name: "Xi'ri's Gift"},
	{Name: "Quagmirran's Eye"},
	{Name: "Totem of the Void"},
	{Name: "Sky Breaker"},
	{Name: "Silvermoon Crest Shield"},
}

var p1Gear = EquipmentSpec{
	{Name: "Cyclone Faceguard (Tier 4)"},
	{Name: "Adornment of Stolen Souls"},
	{Name: "Cyclone Shoulderguards (Tier 4)"},
	{Name: "Ruby Drape of the Mysticant"},
	{Name: "Netherstrike Breastplate"},
	{Name: "Netherstrike Bracers"},
	{Name: "Soul-Eater's Handwraps"},
	{Name: "Netherstrike Belt"},
	{Name: "Stormsong Kilt"},
	{Name: "Windshear Boots"},
	{Name: "Ring of Unrelenting Storms"},
	{Name: "Ring of Recurrence"},
	{Name: "The Lightning Capacitor"},
	{Name: "Icon of the Silver Crescent"},
	{Name: "Totem of the Void"},
	{Name: "Nathrezim Mindblade"},
	{Name: "Mazthoril Honor Shield"},
}

func TestSimulatePreRaidNoBuffs(t *testing.T) {
	simAllEncountersTest(AllEncountersTestOptions{
		label: "preRaidNoBuffs",
		t:     t,

		Options:   basicOptions,
		Gear:      preRaidGear,
		AgentType: AGENT_TYPE_ADAPTIVE,

		ExpectedDpsShort: 867,
		ExpectedDpsLong:  277,
	})
}

func TestSimulatePreRaid(t *testing.T) {
	simAllEncountersTest(AllEncountersTestOptions{
		label: "preRaid",
		t:     t,

		Options:   fullOptions,
		Gear:      preRaidGear,
		AgentType: AGENT_TYPE_ADAPTIVE,

		ExpectedDpsShort: 1406,
		ExpectedDpsLong:  1017,
	})
}

func TestSimulateP1(t *testing.T) {
	simAllEncountersTest(AllEncountersTestOptions{
		label: "p1",
		t:     t,

		Options:   fullOptions,
		Gear:      p1Gear,
		AgentType: AGENT_TYPE_ADAPTIVE,

		ExpectedDpsShort: 1539.5,
		ExpectedDpsLong:  1359,
	})
}

func TestMultiTarget(t *testing.T) {
	doSimulateTest(
		"multiTarget",
		t,
		makeOptions(
			fullOptions,
			Encounter{
				Duration:     300,
				NumClTargets: 3,
			},
			AGENT_TYPE_ADAPTIVE),
		p1Gear,
		1678.5)
}

func TestLBOnlyAgent(t *testing.T) {
	simAllEncountersTest(AllEncountersTestOptions{
		label: "lbOnly",
		t:     t,

		Options:   fullOptions,
		Gear:      p1Gear,
		AgentType: AGENT_TYPE_FIXED_LB_ONLY,

		ExpectedDpsShort: 1581.1,
		ExpectedDpsLong:  1227.6,
	})
}

func TestFixedAgent(t *testing.T) {
	simAllEncountersTest(AllEncountersTestOptions{
		label: "fixedAgent",
		t:     t,

		Options:   fullOptions,
		Gear:      p1Gear,
		AgentType: AGENT_TYPE_FIXED_4LB_1CL,

		ExpectedDpsShort: 1489.3,
		ExpectedDpsLong:  1284.2,
	})
}

func TestClearcastAgent(t *testing.T) {
	simAllEncountersTest(AllEncountersTestOptions{
		label: "clOnClearcast",
		t:     t,

		Options:   fullOptions,
		Gear:      p1Gear,
		AgentType: AGENT_TYPE_CL_ON_CLEARCAST,

		ExpectedDpsShort: 1667,
		ExpectedDpsLong:  1359.1,
	})
}

func BenchmarkSimulate(b *testing.B) {

	for i := 0; i < b.N; i++ {
		RunSimulation(SimRequest{
			Options:     fullOptions,
			Gear:        p1Gear,
			Iterations:  1000,
			IncludeLogs: false,
		})
	}
}

type AllEncountersTestOptions struct {
	label string
	t     *testing.T

	Options   Options
	Gear      EquipmentSpec
	AgentType AgentType

	ExpectedDpsShort float64
	ExpectedDpsLong  float64
}

func simAllEncountersTest(testOpts AllEncountersTestOptions) {
	doSimulateTest(
		testOpts.label+"-short",
		testOpts.t,
		makeOptions(testOpts.Options, shortEncounter, testOpts.AgentType),
		testOpts.Gear,
		testOpts.ExpectedDpsShort)

	doSimulateTest(
		testOpts.label+"-long",
		testOpts.t,
		makeOptions(testOpts.Options, longEncounter, testOpts.AgentType),
		testOpts.Gear,
		testOpts.ExpectedDpsLong)
}

// Performs a basic end-to-end test of the simulator.
//   This is where we can add more sophisticated checks if we would like.
//   Any changes to the damage output of an item set
func doSimulateTest(label string, t *testing.T, options Options, gear EquipmentSpec, expectedDps float64) {
	result := RunSimulation(SimRequest{
		Options:     options,
		Gear:        gear,
		Iterations:  1,
		IncludeLogs: false,
	})

	tolerance := 0.5
	if result.DpsAvg < expectedDps-tolerance || result.DpsAvg > expectedDps+tolerance {
		t.Fatalf("%s failed: expected %0f dps from sim but was %0f", label, expectedDps, result.DpsAvg)
	}
}

func makeOptions(baseOptions Options, encounter Encounter, agentType AgentType) Options {
	baseOptions.Encounter = encounter
	baseOptions.AgentType = agentType
	return baseOptions
}
