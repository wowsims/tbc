package elemental

import (
	"testing"

	_ "github.com/wowsims/tbc/sim/common"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	RegisterElementalShaman()
}

func TestP1FullCharacterStats(t *testing.T) {
	isr := core.NewIndividualSimRequest(core.IndividualSimInputs{
		Player: &proto.Player{
			Race:      proto.Race_RaceTroll10,
			Class:     proto.Class_ClassShaman,
			Equipment: P1Gear,
			Consumes:  FullConsumes,
			Spec:      PlayerOptionsAdaptive,
		},

		RaidBuffs:       FullRaidBuffs,
		PartyBuffs:      FullPartyBuffs,
		IndividualBuffs: FullIndividualBuffs,
	})

	core.CharacterStatsTest("p1Full", t, isr, stats.Stats{
		stats.Strength:  140.7,
		stats.Agility:   100.0,
		stats.Stamina:   522.4,
		stats.Intellect: 531.2,
		stats.Spirit:    182.5,

		stats.SpellPower:       1109,
		stats.HealingPower:     791,
		stats.ArcaneSpellPower: 80,
		stats.HolySpellPower:   80,
		stats.NatureSpellPower: 123,

		stats.MP5:       337.9,
		stats.SpellHit:  125.600,
		stats.SpellCrit: 695.705,

		stats.AttackPower: 401.4,
		stats.MeleeHit:    47.310,
		stats.MeleeCrit:   139.541,

		stats.Mana:  10646,
		stats.Armor: 9370.0,
	})
}

var StatsToTest = []proto.Stat{
	proto.Stat_StatIntellect,
	proto.Stat_StatSpellPower,
	proto.Stat_StatSpellHit,
	proto.Stat_StatSpellCrit,
}

var ReferenceStat = proto.Stat_StatSpellPower

func TestCalcStatWeight(t *testing.T) {
	swr := &proto.StatWeightsRequest{
		Player: &proto.Player{
			Race:      proto.Race_RaceTroll10,
			Class:     proto.Class_ClassShaman,
			Equipment: P1Gear,
			Consumes:  FullConsumes,
			Spec:      PlayerOptionsAdaptive,
			Buffs:     FullIndividualBuffs,
		},
		RaidBuffs:  FullRaidBuffs,
		PartyBuffs: FullPartyBuffs,
		Encounter: &proto.Encounter{
			Targets: []*proto.Target{
				FullDebuffTarget,
			},
		},
		StatsToWeigh:    StatsToTest,
		EpReferenceStat: ReferenceStat,
		SimOptions:      core.DefaultSimTestOptions,
	}

	core.StatWeightsTest("p1Full", t, swr, stats.Stats{
		stats.Intellect:  0.190,
		stats.SpellPower: 0.701,
		stats.SpellHit:   1.445,
		stats.SpellCrit:  0.581,
	})
}

func TestAllSettings(t *testing.T) {
	core.RunTestSuite(t, t.Name(), &core.SettingsCombos{
		Class: proto.Class_ClassShaman,
		Races: []core.RaceCombo{
			core.RaceCombo{Label: "Orc", Race: proto.Race_RaceOrc},
			core.RaceCombo{Label: "Troll10", Race: proto.Race_RaceTroll10},
		},
		GearSets: []core.GearSetCombo{
			core.GearSetCombo{Label: "P1", GearSet: P1Gear},
		},
		SpecOptions: []core.SpecOptionsCombo{
			core.SpecOptionsCombo{Label: "LBOnly", SpecOptions: PlayerOptionsLBOnly},
			core.SpecOptionsCombo{Label: "Fixed3LBCL", SpecOptions: PlayerOptionsFixed3LBCL},
			core.SpecOptionsCombo{Label: "CLOnClearcast", SpecOptions: PlayerOptionsCLOnClearcast},
			core.SpecOptionsCombo{Label: "CLOnClearcastNoBuffs", SpecOptions: PlayerOptionsCLOnClearcastNoBuffs},
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

func TestAllItemEffects(t *testing.T) {
	core.RunTestSuite(t, t.Name(), &core.ItemsTestGenerator{
		Player: &proto.Player{
			Race:      proto.Race_RaceOrc,
			Class:     proto.Class_ClassShaman,
			Spec:      PlayerOptionsCLOnClearcast,
			Equipment: P1Gear,
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
			core.RaceCombo{Label: "Orc", Race: proto.Race_RaceOrc},
		},
		GearSets: []core.GearSetCombo{
			core.GearSetCombo{Label: "P1", GearSet: P1Gear},
		},
		SpecOptions: []core.SpecOptionsCombo{
			core.SpecOptionsCombo{Label: "Adaptive", SpecOptions: PlayerOptionsAdaptive},
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

func BenchmarkSimulate(b *testing.B) {
	isr := core.NewIndividualSimRequest(core.IndividualSimInputs{
		Player: &proto.Player{
			Race:      proto.Race_RaceOrc,
			Class:     proto.Class_ClassShaman,
			Equipment: P1Gear,
			Consumes:  FullConsumes,
			Spec:      PlayerOptionsAdaptive,
		},

		RaidBuffs:       FullRaidBuffs,
		PartyBuffs:      FullPartyBuffs,
		IndividualBuffs: FullIndividualBuffs,

		Target: FullDebuffTarget,
	})

	core.IndividualBenchmark(b, isr)
}
