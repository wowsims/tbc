package mage

import (
	"testing"

	_ "github.com/wowsims/tbc/sim/common"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func init() {
	RegisterMage()
}

func TestAllFireSettings(t *testing.T) {
	core.RunTestSuite(t, t.Name(), &core.SettingsCombos{
		Class: proto.Class_ClassMage,
		Races: []core.RaceCombo{
			core.RaceCombo{Label: "Troll10", Race: proto.Race_RaceTroll10},
		},
		GearSets: []core.GearSetCombo{
			core.GearSetCombo{Label: "P1Fire", GearSet: P1FireGear},
		},
		SpecOptions: []core.SpecOptionsCombo{
			core.SpecOptionsCombo{Label: "FireRotation", SpecOptions: PlayerOptionsFire},
		},
		Buffs: []core.BuffsCombo{
			core.BuffsCombo{
				Label: "NoBuffs",
			},
			core.BuffsCombo{
				Label:    "FullBuffs",
				Raid:     FullRaidBuffs,
				Party:    FullFirePartyBuffs,
				Player:   FullIndividualBuffs,
				Consumes: FullFireConsumes,
			},
		},
		Encounters: core.MakeDefaultEncounterCombos(FullDebuffs),
		SimOptions: core.DefaultSimTestOptions,
	})
}

func TestAllFrostSettings(t *testing.T) {
	core.RunTestSuite(t, t.Name(), &core.SettingsCombos{
		Class: proto.Class_ClassMage,
		Races: []core.RaceCombo{
			core.RaceCombo{Label: "Troll10", Race: proto.Race_RaceTroll10},
		},
		GearSets: []core.GearSetCombo{
			core.GearSetCombo{Label: "P1Frost", GearSet: P1FrostGear},
		},
		SpecOptions: []core.SpecOptionsCombo{
			core.SpecOptionsCombo{Label: "FrostRotation", SpecOptions: PlayerOptionsFrost},
		},
		Buffs: []core.BuffsCombo{
			core.BuffsCombo{
				Label: "NoBuffs",
			},
			core.BuffsCombo{
				Label:    "FullBuffs",
				Raid:     FullRaidBuffs,
				Party:    FullFrostPartyBuffs,
				Player:   FullIndividualBuffs,
				Consumes: FullFrostConsumes,
			},
		},
		Encounters: core.MakeDefaultEncounterCombos(FullDebuffs),
		SimOptions: core.DefaultSimTestOptions,
	})
}

func TestAllArcaneSettings(t *testing.T) {
	core.RunTestSuite(t, t.Name(), &core.SettingsCombos{
		Class: proto.Class_ClassMage,
		Races: []core.RaceCombo{
			core.RaceCombo{Label: "Troll10", Race: proto.Race_RaceTroll10},
		},
		GearSets: []core.GearSetCombo{
			core.GearSetCombo{Label: "P1Arcane", GearSet: P1ArcaneGear},
		},
		SpecOptions: []core.SpecOptionsCombo{
			core.SpecOptionsCombo{Label: "ArcaneRotation", SpecOptions: PlayerOptionsArcane},
		},
		Buffs: []core.BuffsCombo{
			core.BuffsCombo{
				Label: "NoBuffs",
			},
			core.BuffsCombo{
				Label:    "FullBuffs",
				Raid:     FullRaidBuffs,
				Party:    FullArcanePartyBuffs,
				Player:   FullArcaneIndividualBuffs,
				Consumes: FullArcaneConsumes,
			},
		},
		Encounters: core.MakeDefaultEncounterCombos(FullDebuffs),
		SimOptions: core.DefaultSimTestOptions,
	})
}

func TestAllItemEffects(t *testing.T) {
	core.RunTestSuite(t, t.Name(), &core.ItemsTestGenerator{
		Player: &proto.Player{
			Race:      proto.Race_RaceUndead,
			Class:     proto.Class_ClassMage,
			Spec:      PlayerOptionsArcane,
			Equipment: P1ArcaneGear,
			Consumes:  FullArcaneConsumes,
			Buffs:     FullArcaneIndividualBuffs,
		},
		RaidBuffs:  FullRaidBuffs,
		PartyBuffs: FullArcanePartyBuffs,
		Encounter:  core.MakeSingleTargetFullDebuffEncounter(FullDebuffs),
		SimOptions: core.DefaultSimTestOptions,

		ItemFilter: core.ItemFilter{
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
		Class: proto.Class_ClassMage,
		Races: []core.RaceCombo{
			core.RaceCombo{Label: "Troll10", Race: proto.Race_RaceTroll10},
		},
		GearSets: []core.GearSetCombo{
			core.GearSetCombo{Label: "P1Arcane", GearSet: P1ArcaneGear},
		},
		SpecOptions: []core.SpecOptionsCombo{
			core.SpecOptionsCombo{Label: "Arcane", SpecOptions: PlayerOptionsArcane},
		},
		Buffs: []core.BuffsCombo{
			core.BuffsCombo{
				Label:    "FullBuffs",
				Raid:     FullRaidBuffs,
				Party:    FullArcanePartyBuffs,
				Player:   FullArcaneIndividualBuffs,
				Consumes: FullArcaneConsumes,
			},
		},
		Encounters: core.MakeAverageDefaultEncounterCombos(FullDebuffs),
		SimOptions: core.AverageDefaultSimTestOptions,
	})
}
