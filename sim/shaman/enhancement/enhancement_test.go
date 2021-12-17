package enhancement

import (
	"testing"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func init() {
	RegisterEnhancementShaman()
}

func TestSimulatePreRaidNoBuffs(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "preRaid-basic",
		T:     t,

		Inputs: core.IndividualSimInputs{
			Player: &proto.Player{
				Race:      proto.Race_RaceTroll10,
				Class:     proto.Class_ClassShaman,
				Equipment: PreRaidGear,
				// no consumes
				Spec: PlayerOptionsBasic,
			},

			RaidBuffs:       BasicRaidBuffs,
			PartyBuffs:      BasicPartyBuffs,
			IndividualBuffs: BasicIndividualBuffs,

			Target: NoDebuffTarget,
		},

		ExpectedDpsShort: 994.9,
		ExpectedDpsLong:  364.5,
	})
}
