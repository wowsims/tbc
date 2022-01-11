package enhancement

import (
	"testing"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func init() {
	RegisterEnhancementShaman()
}

func TestAllSettings(t *testing.T) {
	core.RunTestSuite(t, t.Name(), &core.SettingsCombos{
		Class: proto.Class_ClassShaman,
		Races: []core.RaceCombo{
			core.RaceCombo{Label: "Troll10", Race: proto.Race_RaceTroll10},
		},
		GearSets: []core.GearSetCombo{
			core.GearSetCombo{Label: "P2", GearSet: Phase2Gear},
		},
		SpecOptions: []core.SpecOptionsCombo{
			core.SpecOptionsCombo{Label: "Basic", SpecOptions: PlayerOptionsBasic},
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
	core.RunTestSuite(t, t.Name(), &core.SettingsCombos{
		Class: proto.Class_ClassShaman,
		Races: []core.RaceCombo{
			core.RaceCombo{Label: "Troll10", Race: proto.Race_RaceTroll10},
		},
		GearSets: []core.GearSetCombo{
			core.GearSetCombo{Label: "P2", GearSet: Phase2Gear},
		},
		SpecOptions: []core.SpecOptionsCombo{
			core.SpecOptionsCombo{Label: "Basic", SpecOptions: PlayerOptionsBasic},
		},
		Buffs: []core.BuffsCombo{
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
