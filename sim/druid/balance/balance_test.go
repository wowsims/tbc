package balance

import (
	"testing"

	_ "github.com/wowsims/tbc/sim/common" // imported to get caster sets included. (we use spellfire here)
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func init() {
	RegisterBalanceDruid()
}

func TestAllSettings(t *testing.T) {
	core.TestSuiteAllSettingsCombos(t, t.Name(), core.SettingsCombos{
		Class: proto.Class_ClassDruid,
		Races: []core.RaceCombo{
			core.RaceCombo{Label: "Tauren", Race: proto.Race_RaceTauren},
		},
		GearSets: []core.GearSetCombo{
			core.GearSetCombo{Label: "P1", GearSet: P1Gear},
			core.GearSetCombo{Label: "P2", GearSet: P2Gear},
		},
		SpecOptions: []core.SpecOptionsCombo{
			core.SpecOptionsCombo{Label: "Wrath", SpecOptions: PlayerOptionsWrath},
			core.SpecOptionsCombo{Label: "Starfire", SpecOptions: PlayerOptionsStarfire},
			core.SpecOptionsCombo{Label: "Adaptive", SpecOptions: PlayerOptionsAdaptive},
		},
		Buffs: []core.BuffsCombo{
			core.BuffsCombo{
				Label: "NoBuffs",
			},
			core.BuffsCombo{
				Label:    "FullBuffs",
				Raid:     FullRaidBuffs,
				Party:    FullPartyBuffs,
				Player:   FullIndividualBuffs,
				Consumes: FullConsumes,
			},
		},
		Encounters: core.MakeDefaultEncounterCombos(FullDebuffs),
		SimOptions: core.DefaultSimTestOptions,
	})
}

func TestAverageDPS(t *testing.T) {
	core.TestSuiteAllSettingsCombos(t, t.Name(), core.SettingsCombos{
		Class: proto.Class_ClassDruid,
		Races: []core.RaceCombo{
			core.RaceCombo{Label: "Tauren", Race: proto.Race_RaceTauren},
		},
		GearSets: []core.GearSetCombo{
			core.GearSetCombo{Label: "P1", GearSet: P1Gear},
			core.GearSetCombo{Label: "P2", GearSet: P2Gear},
		},
		SpecOptions: []core.SpecOptionsCombo{
			core.SpecOptionsCombo{Label: "Adaptive", SpecOptions: PlayerOptionsAdaptive},
		},
		Buffs: []core.BuffsCombo{
			core.BuffsCombo{
				Label: "NoBuffs",
			},
			core.BuffsCombo{
				Label:    "FullBuffs",
				Raid:     FullRaidBuffs,
				Party:    FullPartyBuffs,
				Player:   FullIndividualBuffs,
				Consumes: FullConsumes,
			},
		},
		Encounters: core.MakeAverageDefaultEncounterCombos(FullDebuffs),
		SimOptions: core.AverageDefaultSimTestOptions,
	})
}
