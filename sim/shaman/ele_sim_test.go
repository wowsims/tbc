package shaman

import (
	"log"
	"testing"

	"github.com/wowsims/tbc/items"
	"github.com/wowsims/tbc/sim/api"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	RegisterElementalShaman()
}

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
	Bloodlust: 1,
}

var shamTalents = api.ShamanTalents{
	ElementalFocus:     true,
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

var playerOptionsAdaptive = api.PlayerOptions{
	Spec: &api.PlayerOptions_ElementalShaman{
		ElementalShaman: &api.ElementalShaman{
			Talents: &shamTalents,
			Options: &api.ElementalShaman_Options{
				WaterShield: true,
			},
			Agent: &api.ElementalShaman_Agent{
				Type: api.ElementalShaman_Agent_Adaptive,
			},
		},
	},
}

var playerOptionsLBOnly = api.PlayerOptions{
	Spec: &api.PlayerOptions_ElementalShaman{
		ElementalShaman: &api.ElementalShaman{
			Talents: &shamTalents,
			Options: &api.ElementalShaman_Options{
				WaterShield: true,
			},
			Agent: &api.ElementalShaman_Agent{
				Type: api.ElementalShaman_Agent_FixedLBCL,
			},
		},
	},
}

var playerOptionsCLOnClearcast = api.PlayerOptions{
	Spec: &api.PlayerOptions_ElementalShaman{
		ElementalShaman: &api.ElementalShaman{
			Talents: &shamTalents,
			Options: &api.ElementalShaman_Options{
				WaterShield: true,
			},
			Agent: &api.ElementalShaman_Agent{
				Type: api.ElementalShaman_Agent_CLOnClearcast,
			},
		},
	},
}

var fullBuffs = core.Buffs{
	ArcaneBrilliance:  true,
	GiftOfTheWild:     api.TristateEffect_TristateEffectRegular,
	BlessingOfKings:   true,
	BlessingOfWisdom:  api.TristateEffect_TristateEffectRegular,
	JudgementOfWisdom: true,
	MoonkinAura:       api.TristateEffect_TristateEffectRegular,
	ShadowPriestDPS:   500,
	Bloodlust:         1,
	// Misery:                   true,

	ManaSpringTotem: api.TristateEffect_TristateEffectRegular,
	TotemOfWrath:    1,
	WrathOfAirTotem: api.TristateEffect_TristateEffectRegular,
}

var fullConsumes = core.Consumes{
	FlaskOfBlindingLight: true,
	BrilliantWizardOil:   true,
	BlackenedBasilisk:    true,
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
	"Cyclone Faceguard",
	"Adornment of Stolen Souls",
	"Cyclone Shoulderguards",
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

func gearFromStrings(gears []string) items.EquipmentSpec {
	eq := items.EquipmentSpec{}
	for i, gear := range gears {
		item := items.ByName[gear]
		if item.ID == 0 {
			log.Fatalf("Item not found: %s", gear)
		}
		eq[i].ID = item.ID
	}
	return eq
}

func TestSimulatePreRaidNoBuffs(t *testing.T) {
	simAllEncountersTest(AllEncountersTestOptions{
		label: "preRaid",
		t:     t,

		Options: basicOptions,
		// no consumes
		Buffs: basicBuffs,
		Race:  core.RaceBonusTypeTroll10,

		PlayerOptions: &playerOptionsAdaptive,
		Gear:          gearFromStrings(preRaidGear),

		ExpectedDpsShort: 867,
		ExpectedDpsLong:  269,
	})
}

func TestSimulatePreRaid(t *testing.T) {
	simAllEncountersTest(AllEncountersTestOptions{
		label: "preRaid",
		t:     t,

		Options:  basicOptions,
		Consumes: fullConsumes,
		Buffs:    fullBuffs,
		Race:     core.RaceBonusTypeOrc,

		PlayerOptions: &playerOptionsAdaptive,
		Gear:          gearFromStrings(preRaidGear),

		ExpectedDpsShort: 1398.5,
		ExpectedDpsLong:  1096.3,
	})
}

func TestSimulateP1(t *testing.T) {
	simAllEncountersTest(AllEncountersTestOptions{
		label: "phase1",
		t:     t,

		Options:  basicOptions,
		Consumes: fullConsumes,
		Buffs:    fullBuffs,
		Race:     core.RaceBonusTypeOrc,

		PlayerOptions: &playerOptionsAdaptive,
		Gear:          gearFromStrings(p1Gear),

		ExpectedDpsShort: 1539.5,
		ExpectedDpsLong:  1260.3,
	})
}

func TestMultiTarget(t *testing.T) {
	params := core.IndividualParams{
		Equip:         gearFromStrings(p1Gear),
		Race:          core.RaceBonusTypeOrc,
		Consumes:      fullConsumes,
		Buffs:         fullBuffs,
		Options:       makeOptions(basicOptions, longEncounter),
		PlayerOptions: &playerOptionsAdaptive,
	}
	params.Options.Encounter.NumTargets = 3

	doSimulateTest(
		"multiTarget",
		t,
		params,
		1533.5)
}

func TestLBOnlyAgent(t *testing.T) {
	simAllEncountersTest(AllEncountersTestOptions{
		label: "lbonly",
		t:     t,

		Options:  basicOptions,
		Consumes: fullConsumes,
		Buffs:    fullBuffs,
		Race:     core.RaceBonusTypeOrc,

		PlayerOptions: &playerOptionsLBOnly,
		Gear:          gearFromStrings(p1Gear),

		ExpectedDpsShort: 1581.1,
		ExpectedDpsLong:  1271.9,
	})
}

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

		Options:  basicOptions,
		Consumes: fullConsumes,
		Buffs:    fullBuffs,
		Race:     core.RaceBonusTypeOrc,

		PlayerOptions: &playerOptionsCLOnClearcast,
		Gear:          gearFromStrings(p1Gear),

		ExpectedDpsShort: 1459.8,
		ExpectedDpsLong:  1221.8,
	})
}

func TestAverageDPS(t *testing.T) {
	eq := gearFromStrings(p1Gear)

	options := basicOptions
	options.Iterations = 10000
	options.Encounter = longEncounter
	options.Encounter.NumTargets = 3
	// options.Debug = true

	params := core.IndividualParams{
		Equip:         eq,
		Race:          core.RaceBonusTypeOrc,
		Consumes:      fullConsumes,
		Buffs:         fullBuffs,
		Options:       options,
		PlayerOptions: &playerOptionsAdaptive,
		CustomStats:   stats.Stats{},
	}

	sim := core.NewIndividualSim(params)
	result := sim.Run()

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

	Options  core.Options
	Gear     items.EquipmentSpec
	Buffs    core.Buffs
	Consumes core.Consumes
	Race     core.RaceBonusType

	PlayerOptions *api.PlayerOptions

	ExpectedDpsShort float64
	ExpectedDpsLong  float64
}

func simAllEncountersTest(testOpts AllEncountersTestOptions) {
	params := core.IndividualParams{
		Equip:    testOpts.Gear,
		Race:     testOpts.Race,
		Consumes: testOpts.Consumes,
		Buffs:    testOpts.Buffs,
		Options:  makeOptions(testOpts.Options, shortEncounter),

		PlayerOptions: testOpts.PlayerOptions,
		CustomStats:   stats.Stats{},
	}
	doSimulateTest(
		testOpts.label+"-short",
		testOpts.t,
		params,
		testOpts.ExpectedDpsShort)

	params.Options = makeOptions(testOpts.Options, longEncounter)
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
	params.Options.Debug = true
	params.Options.Iterations = 1

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
