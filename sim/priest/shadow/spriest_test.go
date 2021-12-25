package shadow

import (
	"testing"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func init() {
	RegisterShadowPriest()
}

func TestAllSettings(t *testing.T) {
	core.TestSuiteAllSettingsCombos(t, t.Name(), core.SettingsCombos{
		Class: proto.Class_ClassPriest,
		Races: []core.RaceCombo{
			core.RaceCombo{Label: "Undead", Race: proto.Race_RaceUndead},
		},
		GearSets: []core.GearSetCombo{
			core.GearSetCombo{Label: "P1", GearSet: P1Gear},
		},
		SpecOptions: []core.SpecOptionsCombo{
			core.SpecOptionsCombo{Label: "Basic", SpecOptions: PlayerOptionsBasic},
			core.SpecOptionsCombo{Label: "Clipping", SpecOptions: PlayerOptionsClipping},
			core.SpecOptionsCombo{Label: "Ideal", SpecOptions: PlayerOptionsIdeal},
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
		Class: proto.Class_ClassPriest,
		Races: []core.RaceCombo{
			core.RaceCombo{Label: "Undead", Race: proto.Race_RaceUndead},
		},
		GearSets: []core.GearSetCombo{
			core.GearSetCombo{Label: "P1", GearSet: P1Gear},
		},
		SpecOptions: []core.SpecOptionsCombo{
			core.SpecOptionsCombo{Label: "Clipping", SpecOptions: PlayerOptionsClipping},
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
