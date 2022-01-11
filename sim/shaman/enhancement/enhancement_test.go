package enhancement

import (
	"testing"

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
		stats.Strength:  416.7,
		stats.Agility:   477.9,
		stats.Stamina:   575.2,
		stats.Intellect: 307.9,
		stats.Spirit:    153.9,

		stats.SpellPower: 730.5,
		stats.MP5:        100,
		stats.SpellCrit:  134.9,

		stats.AttackPower: 2434.9,
		stats.MeleeHit:    254.6,
		stats.MeleeCrit:   833.1,
		stats.Expertise:   35,

		stats.Mana:  7661.2,
		stats.Armor: 5383.8,
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
