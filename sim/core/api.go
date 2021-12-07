// Proto-based function interface for the simulator
package core

import (
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

/**
 * Returns all items, enchants, and gems recognized by the sim.
 */
func GetGearList(request *proto.GearListRequest) *proto.GearListResult {
	result := &proto.GearListResult{}

	for i := range items.Items {
		item := items.Items[i]
		result.Items = append(result.Items, item.ToProto())
	}
	for i := range items.Gems {
		gem := items.Gems[i]
		result.Gems = append(result.Gems, gem.ToProto())
	}
	for i := range items.Enchants {
		enchant := items.Enchants[i]
		result.Enchants = append(result.Enchants, enchant.ToProto())
	}

	return result
}

/**
 * Returns character stats taking into account gear / buffs / consumes / etc
 */
func ComputeStats(csr *proto.ComputeStatsRequest) *proto.ComputeStatsResult {
	raid := NewRaid(proto.Raid{
		Parties: []*proto.Party{
			&proto.Party{
				Players: []*proto.Player{
					csr.Player,
				},
				Buffs: csr.PartyBuffs,
			},
		},
		Buffs: csr.RaidBuffs,
	})

	agent := raid.Parties[0].Players[0]

	gearStats := agent.GetCharacter().Equip.Stats()
	finalStats := agent.GetCharacter().GetStats()
	setBonusNames := agent.GetCharacter().GetActiveSetBonusNames()

	return &proto.ComputeStatsResult{
		GearOnly:   gearStats[:],
		FinalStats: finalStats[:],
		Sets:       setBonusNames,
	}
}

/**
 * Returns stat weights and EP values, with standard deviations, for all stats.
 */
func StatWeights(request *proto.StatWeightsRequest) *proto.StatWeightsResult {
	statsToWeigh := stats.ProtoArrayToStatsList(request.StatsToWeigh)

	result := CalcStatWeight(*request, statsToWeigh, stats.Stat(request.EpReferenceStat))

	return &proto.StatWeightsResult{
		Weights:       result.Weights[:],
		WeightsStdev:  result.WeightsStdev[:],
		EpValues:      result.EpValues[:],
		EpValuesStdev: result.EpValuesStdev[:],
	}
}

/**
 * Runs multiple iterations of the sim with just 1 player.
 */
func RunIndividualSim(request *proto.IndividualSimRequest) *proto.IndividualSimResult {
	raidResult := RunRaidSim(&proto.RaidSimRequest{
		Raid:       SinglePlayerRaidProto(request.Player, request.PartyBuffs, request.RaidBuffs),
		Encounter:  request.Encounter,
		SimOptions: request.SimOptions,
	})

	return &proto.IndividualSimResult{
		PlayerMetrics:    raidResult.RaidMetrics.Parties[0].Players[0],
		EncounterMetrics: raidResult.EncounterMetrics,
		Logs:             raidResult.Logs,
	}
}

/**
 * Runs multiple iterations of the sim with a full raid.
 */
func RunRaidSim(request *proto.RaidSimRequest) *proto.RaidSimResult {
	return RunSim(*request)
}
