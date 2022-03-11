package retribution

import (
	"testing"

	_ "github.com/wowsims/tbc/sim/common" // imported to get item effects included.
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func init() {
	RegisterRetributionPaladin()
}

func TestRetribution(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator(core.CharacterSuiteConfig{
		Class: proto.Class_ClassPaladin,

		Race:       proto.Race_RaceBloodElf,
		OtherRaces: []proto.Race{proto.Race_RaceHuman, proto.Race_RaceDraenei, proto.Race_RaceDwarf}, // To-do deal with tests for other races

		GearSet: core.GearSetCombo{Label: "P2", GearSet: Phase2Gear},

		SpecOptions: core.SpecOptionsCombo{Label: "Retribution Paladin", SpecOptions: PlayerOptionsBasic},

		RaidBuffs:   FullRaidBuffs,
		PartyBuffs:  FullPartyBuffs,
		PlayerBuffs: FullIndividualBuffs,
		Consumes:    FullConsumes,
		Debuffs:     FullDebuffs,

		ItemFilter: core.ItemFilter{
			WeaponTypes: []proto.WeaponType{
				proto.WeaponType_WeaponTypeAxe,
				proto.WeaponType_WeaponTypeSword,
				proto.WeaponType_WeaponTypePolearm,
			},
			ArmorType: proto.ArmorType_ArmorTypePlate,
			RangedWeaponTypes: []proto.RangedWeaponType{
				proto.RangedWeaponType_RangedWeaponTypeLibram,
			},
		},
	}))
}

func BenchmarkSimulate(b *testing.B) {
	// rsr := &proto.RaidSimRequest{
	// 	Raid: core.SinglePlayerRaidProto(
	// 		&proto.Player{
	// 			Race:      proto.Race_RaceOrc,
	// 			Class:     proto.Class_ClassShaman,
	// 			Equipment: Phase2Gear,
	// 			Consumes:  FullConsumes,
	// 			Spec:      PlayerOptionsBasic,
	// 			Buffs:     FullIndividualBuffs,
	// 		},
	// 		FullPartyBuffs,
	// 		FullRaidBuffs),
	// 	Encounter: &proto.Encounter{
	// 		Duration: 300,
	// 		Targets: []*proto.Target{
	// 			FullDebuffTarget,
	// 		},
	// 	},
	// 	SimOptions: core.AverageDefaultSimTestOptions,
	// }

	// core.RaidBenchmark(b, rsr)
}
