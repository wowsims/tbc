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
			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Consumes: FullConsumes,
			Target:   FullDebuffTarget,
			Race:     proto.Race_RaceTauren,
			Class:    proto.Class_ClassDruid,

			PlayerOptions: PlayerOptionsStarfire,
			Gear:          P2Gear,
		},

		ExpectedDpsShort: 1538.2,
		ExpectedDpsLong:  1467.9,
	})
}

func TestSimulateP1Starfire(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "phase1-starfire",
		T:     t,

		Inputs: core.IndividualSimInputs{
			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Consumes: FullConsumes,
			Target:   FullDebuffTarget,
			Race:     proto.Race_RaceTauren,
			Class:    proto.Class_ClassDruid,

			PlayerOptions: PlayerOptionsStarfire,
			Gear:          P1Gear,
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
			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Consumes: FullConsumes,
			Target:   FullDebuffTarget,
			Race:     proto.Race_RaceTauren,
			Class:    proto.Class_ClassDruid,

			PlayerOptions: PlayerOptionsWrath,
			Gear:          P1Gear,
		},

		ExpectedDpsShort: 1258.0,
		ExpectedDpsLong:  1141.1,
	})
}

func TestAdaptive(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "phase2-adaptive",
		T:     t,

		Inputs: core.IndividualSimInputs{
			RaidBuffs:       BasicRaidBuffs,
			PartyBuffs:      BasicPartyBuffs,
			IndividualBuffs: BasicIndividualBuffs,

			Consumes: BasicConsumes,
			Target:   FullDebuffTarget,
			Race:     proto.Race_RaceTauren,
			Class:    proto.Class_ClassDruid,

			PlayerOptions: PlayerOptionsAdaptive,
			Gear:          P2Gear,
		},

		ExpectedDpsShort: 1569.7,
		ExpectedDpsLong:  1083.0,
	})
}

func TestAverageDPS(t *testing.T) {
	isr := core.NewIndividualSimRequest(core.IndividualSimInputs{
		Gear:     P1Gear,
		Race:     proto.Race_RaceNightElf,
		Class:    proto.Class_ClassDruid,
		Consumes: FullConsumes,

		RaidBuffs:       FullRaidBuffs,
		PartyBuffs:      FullPartyBuffs,
		IndividualBuffs: FullIndividualBuffs,

		Target:        FullDebuffTarget,
		PlayerOptions: PlayerOptionsStarfire,
	})

	core.IndividualSimAverageTest("P1Average", t, isr, 1258.1)
}
