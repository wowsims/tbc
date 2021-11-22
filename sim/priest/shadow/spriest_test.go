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

		ExpectedDpsShort: 1198.6,
		ExpectedDpsLong:  1257.7,
	})
}

func TestAverageDPS(t *testing.T) {
	isr := core.NewIndividualSimRequest(core.IndividualSimInputs{
		RaidBuffs:       FullRaidBuffs,
		PartyBuffs:      FullPartyBuffs,
		IndividualBuffs: FullIndividualBuffs,

		Consumes: FullConsumes,
		Target:   FullDebuffTarget,
		Race:     proto.Race_RaceUndead,
		Class:    proto.Class_ClassPriest,

		PlayerOptions: PlayerOptionsBasic,
		Gear:          P1Gear,
	})

	core.IndividualSimAverageTest("P1Average", t, isr, 1245.6)
}
