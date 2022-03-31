package hunter

import (
	"testing"

	_ "github.com/wowsims/tbc/sim/common" // imported to get item effects included.
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func init() {
	RegisterHunter()
}

func TestBestialWrath(t *testing.T) {
	defaultRaid := core.SinglePlayerRaidProto(&proto.Player{
		Class:     proto.Class_ClassHunter,
		Spec:      PlayerOptionsBasic,
		Race:      proto.Race_RaceTroll10,
		Equipment: P1Gear,
	}, FullPartyBuffs, FullRaidBuffs)
	sim := core.NewSim(proto.RaidSimRequest{
		Raid:      defaultRaid,
		Encounter: core.MakeSingleTargetFullDebuffEncounter(FullDebuffs, 5),
		SimOptions: &proto.SimOptions{
			Iterations: 1,
			IsTest:     true,
			Debug:      false,
			RandomSeed: 101,
		},
	})
	h := sim.Raid.Parties[0].Players[0].(*Hunter)
	h.Init(sim)

	sim.Reset()
	h.TryUseCooldowns(sim)
	h.AimedShot.Cast(sim, sim.GetPrimaryTarget())
	if h.AimedShot.Instance.Cost.Value != 259 {
		t.Logf("cost is wrong, expected: %0.1f, actual: %0.1f", 259.0, h.AimedShot.Instance.Cost.Value)
	}
}

func TestHunter(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator(core.CharacterSuiteConfig{
		Class: proto.Class_ClassHunter,

		Race:       proto.Race_RaceOrc,
		OtherRaces: []proto.Race{proto.Race_RaceDwarf},

		GearSet: core.GearSetCombo{Label: "P1", GearSet: P1Gear},

		SpecOptions: core.SpecOptionsCombo{Label: "Basic", SpecOptions: PlayerOptionsBasic},
		OtherSpecOptions: []core.SpecOptionsCombo{
			core.SpecOptionsCombo{Label: "French", SpecOptions: PlayerOptionsFrench},
			core.SpecOptionsCombo{Label: "MeleeWeave", SpecOptions: PlayerOptionsMeleeWeave},
			core.SpecOptionsCombo{Label: "SV", SpecOptions: PlayerOptionsSV},
		},

		RaidBuffs:   FullRaidBuffs,
		PartyBuffs:  FullPartyBuffs,
		PlayerBuffs: FullIndividualBuffs,
		Consumes:    FullConsumes,
		Debuffs:     FullDebuffs,

		ItemFilter: core.ItemFilter{
			ArmorType: proto.ArmorType_ArmorTypeMail,
			RangedWeaponTypes: []proto.RangedWeaponType{
				proto.RangedWeaponType_RangedWeaponTypeBow,
				proto.RangedWeaponType_RangedWeaponTypeCrossbow,
				proto.RangedWeaponType_RangedWeaponTypeGun,
			},
		},
	}))
}

func BenchmarkSimulate(b *testing.B) {
	rsr := &proto.RaidSimRequest{
		Raid: core.SinglePlayerRaidProto(
			&proto.Player{
				Race:      proto.Race_RaceOrc,
				Class:     proto.Class_ClassHunter,
				Equipment: P1Gear,
				Consumes:  FullConsumes,
				Spec:      PlayerOptionsMeleeWeave,
				Buffs:     FullIndividualBuffs,
			},
			FullPartyBuffs,
			FullRaidBuffs),
		Encounter: &proto.Encounter{
			Duration: 300,
			Targets: []*proto.Target{
				FullDebuffTarget,
			},
		},
		SimOptions: core.AverageDefaultSimTestOptions,
	}

	core.RaidBenchmark(b, rsr)
}
