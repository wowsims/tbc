package sim

import (
	"testing"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"

	balanceDruid "github.com/wowsims/tbc/sim/druid/balance"
	shadowPriest "github.com/wowsims/tbc/sim/priest/shadow"
	elementalShaman "github.com/wowsims/tbc/sim/shaman/elemental"
)

func init() {
	RegisterAll()
}

var SimOptions = &proto.SimOptions{
	Iterations: 1,
	IsTest:     true,
}

var StandardTarget = &proto.Target{
	Armor:   7700,
	MobType: proto.MobType_MobTypeDemon,
}

var STEncounter = &proto.Encounter{
	Duration: 300,
	Targets: []*proto.Target{
		StandardTarget,
	},
}

var P1BalanceDruid = &proto.Player{
	Name:      "P1 Boomkin",
	Race:      proto.Race_RaceTauren,
	Class:     proto.Class_ClassDruid,
	Equipment: balanceDruid.P1Gear,
	Consumes:  balanceDruid.FullConsumes,
	Spec:      balanceDruid.PlayerOptionsAdaptive,
	Buffs:     balanceDruid.FullIndividualBuffs,
}

var P1ElementalShaman = &proto.Player{
	Name:      "P1 Ele Shaman",
	Race:      proto.Race_RaceOrc,
	Class:     proto.Class_ClassShaman,
	Equipment: elementalShaman.P1Gear,
	Consumes:  elementalShaman.FullConsumes,
	Spec:      elementalShaman.PlayerOptionsAdaptive,
	Buffs:     elementalShaman.FullIndividualBuffs,
}

var P1ShadowPriest = &proto.Player{
	Name:      "P1 Shadow Priest",
	Race:      proto.Race_RaceUndead,
	Class:     proto.Class_ClassPriest,
	Equipment: shadowPriest.P1Gear,
	Consumes:  shadowPriest.FullConsumes,
	Spec:      shadowPriest.PlayerOptionsIdeal,
	Buffs:     shadowPriest.FullIndividualBuffs,
}

var BasicRaid = &proto.Raid{
	Parties: []*proto.Party{
		&proto.Party{
			Players: []*proto.Player{
				P1BalanceDruid,
				P1ElementalShaman,
				P1ShadowPriest,
			},
		},
	},
}

// Tests that we don't crash with various combinations of empty parties / blank players.
func TestSparseRaid(t *testing.T) {
	sparseRaid := &proto.Raid{
		Parties: []*proto.Party{
			&proto.Party{},
			&proto.Party{
				Players: []*proto.Player{
					&proto.Player{},
					P1ElementalShaman,
					&proto.Player{},
				},
			},
			&proto.Party{
				Players: []*proto.Player{
					&proto.Player{},
					&proto.Player{},
				},
			},
		},
	}

	rsr := &proto.RaidSimRequest{
		Raid:       sparseRaid,
		Encounter:  STEncounter,
		SimOptions: SimOptions,
	}

	core.RaidSimTest("Sparse", t, rsr, 1238.3)
}

func TestBasicRaid(t *testing.T) {
	rsr := &proto.RaidSimRequest{
		Raid:       BasicRaid,
		Encounter:  STEncounter,
		SimOptions: SimOptions,
	}

	core.RaidSimTest("P1 ST", t, rsr, 3818.3)
}
