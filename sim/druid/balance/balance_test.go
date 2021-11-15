package balance

import (
	"testing"

	_ "github.com/wowsims/tbc/sim/common" // imported to get caster sets included. (we use spellfire here)
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/druid"
)

func init() {
	RegisterBalanceDruid()
}

func TestMoonfireReset(t *testing.T) {
	isr := core.NewIndividualSimRequest(core.IndividualSimInputs{
		RaidBuffs:       FullRaidBuffs,
		PartyBuffs:      FullPartyBuffs,
		IndividualBuffs: FullIndividualBuffs,

		Consumes: FullConsumes,
		Target:   FullDebuffTarget,
		Race:     proto.Race_RaceTauren,
		Class:    proto.Class_ClassDruid,

		PlayerOptions: PlayerOptionsStarfire,
		Gear:          P1Gear,
		SimOptions: &proto.SimOptions{
			Iterations: 1,
		},
	})
	isr.Encounter.Duration = 300
	sim := core.NewIndividualSim(*isr)
	result := sim.Run()

	var mfCasts int32
	for _, act := range result.Agents[0].Actions {
		if act.ActionID.SpellID == druid.SpellIDMF {
			mfCasts = act.Casts
		}
	}
	isr.SimOptions.Iterations = 10
	sim2 := core.NewIndividualSim(*isr)
	result2 := sim2.Run()

	for _, act := range result2.Agents[0].Actions {
		if act.ActionID.SpellID == druid.SpellIDMF {
			if mfCasts == act.Casts {
				// TODO: This is really a failure of the framework to cleanup pending dots.
				t.Fatalf("No new moonfire casts after first sim run. This means moonfire action is still ticking.")
			}
		}
	}
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
		ExpectedDpsLong:  1168.6,
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

		ExpectedDpsShort: 1258.2,
		ExpectedDpsLong:  1026.3,
	})
}
