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

		ExpectedDpsShort: 1184.8,
		ExpectedDpsLong:  1155.4,
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

		ExpectedDpsShort: 1160.6,
		ExpectedDpsLong:  1002.8,
	})
}
