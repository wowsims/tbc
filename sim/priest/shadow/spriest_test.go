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
				Spec:      PlayerOptionsLazy,
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

func TestSimulateP1Sweaty(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "phase1-sweaty",
		T:     t,

		Inputs: core.IndividualSimInputs{
			Player: &proto.Player{
				Race:      proto.Race_RaceUndead,
				Class:     proto.Class_ClassPriest,
				Equipment: P1Gear,
				Consumes:  FullConsumes,
				Spec:      PlayerOptionsSweaty,
			},

			RaidBuffs:       FullRaidBuffs,
			PartyBuffs:      FullPartyBuffs,
			IndividualBuffs: FullIndividualBuffs,

			Target: FullDebuffTarget,
		},

		ExpectedDpsShort: 1191.1,
		ExpectedDpsLong:  1262.2,
	})
}

func TestSimulateP1Perfect(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "phase1-sweaty",
		T:     t,

		Inputs: core.IndividualSimInputs{
			Player: &proto.Player{
				Race:      proto.Race_RaceUndead,
				Class:     proto.Class_ClassPriest,
				Equipment: P1Gear,
				Consumes:  FullConsumes,
				Spec:      PlayerOptionsPerfect,
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
			Spec:      PlayerOptionsPerfect,
		},

		RaidBuffs:       FullRaidBuffs,
		PartyBuffs:      FullPartyBuffs,
		IndividualBuffs: FullIndividualBuffs,

		Target: FullDebuffTarget,
	})

	core.IndividualSimAverageTest("P1Average", t, isr, 1271.8)
}
