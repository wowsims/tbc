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
			Player: &proto.Player{
				Race:      proto.Race_RaceUndead,
				Class:     proto.Class_ClassPriest,
				Equipment: P1Gear,
				Consumes:  FullConsumes,
				Spec:      PlayerOptionsBasic,
			},

			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Target: FullDebuffTarget,
		},

		ExpectedDpsShort: 1184.4,
		ExpectedDpsLong:  1235.9,
	})
}

func TestAverageDPS(t *testing.T) {
	isr := core.NewIndividualSimRequest(core.IndividualSimInputs{
		Player: &proto.Player{
			Race:      proto.Race_RaceUndead,
			Class:     proto.Class_ClassPriest,
			Equipment: P1Gear,
			Consumes:  FullConsumes,
			Spec:      PlayerOptionsBasic,
		},

		RaidBuffs:       FullRaidBuffs,
		PartyBuffs:      FullPartyBuffs,
		IndividualBuffs: FullIndividualBuffs,

		Target: FullDebuffTarget,
	})

	core.IndividualSimAverageTest("P1Average", t, isr, 1189.5)
}
