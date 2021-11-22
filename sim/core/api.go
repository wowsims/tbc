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
func ComputeStats(request *proto.ComputeStatsRequest) *proto.ComputeStatsResult {
	agent := NewAgent(*request.Player, proto.IndividualSimRequest{
		Player: request.Player,
		RaidBuffs: request.RaidBuffs,
		PartyBuffs: request.PartyBuffs,
		IndividualBuffs: request.IndividualBuffs,
	})

	raid := NewRaid(*request.RaidBuffs, *request.PartyBuffs, *request.IndividualBuffs)
	raid.AddPlayer(agent)
	raid.Finalize()

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

	result := CalcStatWeight(*request.Options, statsToWeigh, stats.Stat(request.EpReferenceStat))

	return &proto.StatWeightsResult{
		Weights:       result.Weights[:],
		WeightsStdev:  result.WeightsStdev[:],
		EpValues:      result.EpValues[:],
		EpValuesStdev: result.EpValuesStdev[:],
	}
}

/**
 * Runs multiple iterations of the sim with a single set of options / gear.
 */
func RunIndividualSim(request *proto.IndividualSimRequest) *proto.IndividualSimResult {
	sim := NewIndividualSim(*request)
	return sim.RunIndividual()
}
