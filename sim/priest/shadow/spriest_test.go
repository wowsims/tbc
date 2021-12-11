package shadow

import (
	"testing"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func init() {
	RegisterShadowPriest()
}

func TestSimulateP1Lazy(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "phase1-lazy",
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

		ExpectedDpsShort: 1199.2,
		ExpectedDpsLong:  1228.9,
	})
}

func TestSimulateP1Clipping(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "phase1-clipping",
		T:     t,

		Inputs: core.IndividualSimInputs{
			Player: &proto.Player{
				Race:      proto.Race_RaceUndead,
				Class:     proto.Class_ClassPriest,
				Equipment: P1Gear,
				Consumes:  FullConsumes,
				Spec:      PlayerOptionsClipping,
			},

			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Target: FullDebuffTarget,
		},

		ExpectedDpsShort: 1218.8,
		ExpectedDpsLong:  1245.1,
	})
}

func TestSimulateP1Ideal(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "phase1-ideal",
		T:     t,

		Inputs: core.IndividualSimInputs{
			Player: &proto.Player{
				Race:      proto.Race_RaceUndead,
				Class:     proto.Class_ClassPriest,
				Equipment: P1Gear,
				Consumes:  FullConsumes,
				Spec:      PlayerOptionsIdeal,
			},

			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Target: FullDebuffTarget,
		},

		ExpectedDpsShort: 1220.5,
		ExpectedDpsLong:  1266.2,
	})
}

func TestAverageDPS(t *testing.T) {
	isr := core.NewIndividualSimRequest(core.IndividualSimInputs{
		Player: &proto.Player{
			Race:      proto.Race_RaceUndead,
			Class:     proto.Class_ClassPriest,
			Equipment: P1Gear,
			Consumes:  FullConsumes,
			Spec:      PlayerOptionsIdeal,
		},

		RaidBuffs:       FullRaidBuffs,
		PartyBuffs:      FullPartyBuffs,
		IndividualBuffs: FullIndividualBuffs,

		Target: FullDebuffTarget,
	})

	core.IndividualSimAverageTest("P1Average", t, isr, 1271.8)
}
