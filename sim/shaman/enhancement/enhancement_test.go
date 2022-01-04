package enhancement

import (
	"testing"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func init() {
	RegisterEnhancementShaman()
}

func TestSimulatePhase2(t *testing.T) {
	core.IndividualSimAllEncountersTest(core.AllEncountersTestOptions{
		Label: "p2-basic",
		T:     t,

		Inputs: core.IndividualSimInputs{
			Player: &proto.Player{
				Race:      proto.Race_RaceTroll10,
				Class:     proto.Class_ClassShaman,
				Equipment: Phase2Gear,
				// no consumes
				Spec: PlayerOptionsBasic,
			},

			RaidBuffs:       BasicRaidBuffs,
			PartyBuffs:      BasicPartyBuffs,
			IndividualBuffs: BasicIndividualBuffs,

			Target: NoDebuffTarget,
		},

		// these numbers will change while we are still implementing and fixing up enh shaman and melee
		ExpectedDpsShort: 755.0,
		ExpectedDpsLong:  623.5,
	})
}

func TestAverageDPS(t *testing.T) {
	isr := core.NewIndividualSimRequest(core.IndividualSimInputs{
		Player: &proto.Player{
			Race:      proto.Race_RaceTroll10,
			Class:     proto.Class_ClassShaman,
			Equipment: Phase2Gear,
			Consumes:  FullConsumes,
			Spec:      PlayerOptionsBasic,
		},

		RaidBuffs:       FullRaidBuffs,
		PartyBuffs:      FullPartyBuffs,
		IndividualBuffs: FullIndividualBuffs,

		Target: FullDebuffTarget,
	})

	core.IndividualSimAverageTest("P2Average", t, isr, 1024.0)
}
