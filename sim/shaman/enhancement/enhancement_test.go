package enhancement

import (
	"testing"

	_ "github.com/wowsims/tbc/sim/common" // imported to get item effects included.
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	RegisterEnhancementShaman()
}

func TestP1FullCharacterStats(t *testing.T) {
	isr := core.NewIndividualSimRequest(core.IndividualSimInputs{
		Player: &proto.Player{
			Race:      proto.Race_RaceTroll10,
			Class:     proto.Class_ClassShaman,
			Equipment: Phase2Gear,
			Consumes:  FullConsumes,
			Spec:      PlayerOptionsBasic,
		},

		RaidBuffs:       FullRaidBuffs,
		PartyBuffs:      FullPartyBuffs,
		IndividualBuffs: FullIndividualBuffs,
	})

	core.CharacterStatsTest("p2Full", t, isr, stats.Stats{
		stats.Strength:  519.0,
		stats.Agility:   512.0,
		stats.Stamina:   595.0,
		stats.Intellect: 314.5,
		stats.Spirit:    160.5,

		stats.SpellPower: 826.0,
		stats.MP5:        100,
		stats.SpellCrit:  136.8,

		stats.AttackPower: 2753.5,
		stats.MeleeHit:    270.6,
		stats.MeleeCrit:   893.2,
		stats.Expertise:   35,

		stats.Mana:  7765.1,
		stats.Armor: 5452.0,
	})
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

func TestAllItemEffects(t *testing.T) {
	core.RunTestSuite(t, t.Name(), &core.ItemsTestGenerator{
		Player: &proto.Player{
			Race:      proto.Race_RaceOrc,
			Class:     proto.Class_ClassShaman,
			Spec:      PlayerOptionsBasic,
			Equipment: Phase2Gear,
			Consumes:  FullConsumes,
			Buffs:     FullIndividualBuffs,
		},
		RaidBuffs:  FullRaidBuffs,
		PartyBuffs: FullPartyBuffs,
		Encounter:  core.MakeSingleTargetFullDebuffEncounter(FullDebuffs),
		SimOptions: core.DefaultSimTestOptions,

		ItemFilter: core.ItemFilter{
			ArmorTypes: []proto.ArmorType{
				proto.ArmorType_ArmorTypeUnknown,
				proto.ArmorType_ArmorTypeCloth,
				proto.ArmorType_ArmorTypeLeather,
				proto.ArmorType_ArmorTypeMail,
			},
			RangedWeaponTypes: []proto.RangedWeaponType{
				proto.RangedWeaponType_RangedWeaponTypeTotem,
			},
		},
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
