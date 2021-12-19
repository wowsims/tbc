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

func TestNordBonus(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "phase2-nordbonus",
		T:     t,

		Inputs: core.IndividualSimInputs{
			Player: &proto.Player{
				Race:      proto.Race_RaceTauren,
				Class:     proto.Class_ClassDruid,
				Equipment: P2Gear,
				Consumes:  FullConsumes,
				Spec:      PlayerOptionsStarfire,
			},

			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Target: FullDebuffTarget,
		},

		ExpectedDpsShort: 1538.2,
		ExpectedDpsLong:  1467.8,
	})
}

func TestSimulateP1Starfire(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "phase1-starfire",
		T:     t,

		Inputs: core.IndividualSimInputs{
			Player: &proto.Player{
				Race:      proto.Race_RaceTauren,
				Class:     proto.Class_ClassDruid,
				Equipment: P1Gear,
				Consumes:  FullConsumes,
				Spec:      PlayerOptionsStarfire,
			},

			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Target: FullDebuffTarget,
		},

		ExpectedDpsShort: 1413.4,
		ExpectedDpsLong:  1302.9,
	})
}

func TestSimulateP1Wrath(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "phase1-wrath",
		T:     t,

		Inputs: core.IndividualSimInputs{
			Player: &proto.Player{
				Race:      proto.Race_RaceTauren,
				Class:     proto.Class_ClassDruid,
				Equipment: P1Gear,
				Consumes:  FullConsumes,
				Spec:      PlayerOptionsWrath,
			},

			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Target: FullDebuffTarget,
		},

		ExpectedDpsShort: 1141.3,
		ExpectedDpsLong:  1033.4,
	})
}

func TestAdaptive(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "phase2-adaptive",
		T:     t,

		Inputs: core.IndividualSimInputs{
			Player: &proto.Player{
				Race:      proto.Race_RaceTauren,
				Class:     proto.Class_ClassDruid,
				Equipment: P2Gear,
				Consumes:  BasicConsumes,
				Spec:      PlayerOptionsAdaptive,
			},

			RaidBuffs:       BasicRaidBuffs,
			PartyBuffs:      BasicPartyBuffs,
			IndividualBuffs: BasicIndividualBuffs,

			Target: FullDebuffTarget,
		},

		ExpectedDpsShort: 1569.7,
		ExpectedDpsLong:  1153.9,
	})
}

func TestAdaptiveFull(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "phase2-adaptiveFull",
		T:     t,

		Inputs: core.IndividualSimInputs{
			Player: &proto.Player{
				Race:      proto.Race_RaceTauren,
				Class:     proto.Class_ClassDruid,
				Equipment: P2Gear,
				Consumes:  FullConsumes,
				Spec:      PlayerOptionsAdaptive,
			},

			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Target: FullDebuffTarget,
		},

		ExpectedDpsShort: 1538.2,
		ExpectedDpsLong:  1467.9,
	})
}

func TestAverageDPS(t *testing.T) {
	isr := core.NewIndividualSimRequest(core.IndividualSimInputs{
		Player: &proto.Player{
			Race:      proto.Race_RaceNightElf,
			Class:     proto.Class_ClassDruid,
			Equipment: P1Gear,
			Consumes:  FullConsumes,
			Spec:      PlayerOptionsStarfire,
		},

		RaidBuffs:       FullRaidBuffs,
		PartyBuffs:      FullPartyBuffs,
		IndividualBuffs: FullIndividualBuffs,

		Target: FullDebuffTarget,
	})

	core.IndividualSimAverageTest("P1Average", t, isr, 1258.1)
}
