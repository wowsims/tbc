package shadow

import (
	"testing"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func init() {
	RegisterShadowPriest()
}

func TestSimulateP1Basic(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "phase1-basic",
		T:     t,

		Inputs: core.IndividualSimInputs{
			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Consumes: FullConsumes,
			Target:   FullDebuffTarget,
			Race:     proto.Race_RaceUndead,
			Class:    proto.Class_ClassPriest,

			PlayerOptions: PlayerOptionsBasic,
			Gear:          P1Gear,
		},

		ExpectedDpsShort: 1128.9,
		ExpectedDpsLong:  48.7,
	})
}
