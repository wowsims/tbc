package shadow

import (
	"testing"

	_ "github.com/wowsims/tbc/sim/common" // imported to get caster sets included.
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func init() {
	RegisterShadowPriest()
}

func TestAllSettings(t *testing.T) {
	core.RunTestSuite(t, t.Name(), &core.SettingsCombos{
		Class: proto.Class_ClassPriest,
		Races: []core.RaceCombo{
			core.RaceCombo{Label: "Undead", Race: proto.Race_RaceUndead},
		},
		GearSets: []core.GearSetCombo{
			core.GearSetCombo{Label: "P1", GearSet: P1Gear},
			core.GearSetCombo{Label: "P3", GearSet: P3Gear},
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

func TestAllItemEffects(t *testing.T) {
	core.RunTestSuite(t, t.Name(), &core.ItemsTestGenerator{
		Player: &proto.Player{
			Race:      proto.Race_RaceDwarf,
			Class:     proto.Class_ClassPriest,
			Spec:      PlayerOptionsClipping,
			Equipment: P3Gear,
			Consumes:  FullConsumes,
			Buffs:     FullIndividualBuffs,
		},
		RaidBuffs:  FullRaidBuffs,
		PartyBuffs: FullPartyBuffs,
		Encounter:  core.MakeSingleTargetFullDebuffEncounter(FullDebuffs),
		SimOptions: core.DefaultSimTestOptions,

		ItemFilter: core.ItemFilter{
			Categories: []proto.ItemCategory{
				proto.ItemCategory_ItemCategoryCaster,
			},
			ArmorTypes: []proto.ArmorType{
				proto.ArmorType_ArmorTypeUnknown,
				proto.ArmorType_ArmorTypeCloth,
			},
			RangedWeaponTypes: []proto.RangedWeaponType{
				proto.RangedWeaponType_RangedWeaponTypeWand,
			},
		},
	})
}

func TestAverageDPS(t *testing.T) {
	core.RunTestSuite(t, t.Name(), &core.SettingsCombos{
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
