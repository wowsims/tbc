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
	agent := NewAgent(*request.Player, proto.IndividualSimRequest{Player: request.Player, Buffs: request.Buffs})
	raid := NewRaid(*request.Buffs)
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
func RunSimulation(request *proto.IndividualSimRequest) *proto.IndividualSimResult {
	sim := NewIndividualSim(*request)
	result := sim.Run()

	actionMetrics := []*proto.ActionMetric{}
	// TODO: Actually return results for all agents
	for _, v := range result.Agents[0].Actions {
		metric := &proto.ActionMetric{
			Tag:    v.Tag,
			Casts:  v.Casts,
			Hits:   v.Hits,
			Crits:  v.Crits,
			Misses: v.Misses,
			Damage: v.Damage,
		}
		if v.ActionID.SpellID != 0 {
			metric.ActionId = &proto.ActionMetric_SpellId{SpellId: v.ActionID.SpellID}
		}
		if v.ActionID.ItemID != 0 {
			metric.ActionId = &proto.ActionMetric_ItemId{ItemId: v.ActionID.ItemID}
		}
		actionMetrics = append(actionMetrics, metric)
	}
	isr := &proto.IndividualSimResult{
		DpsAvg:              result.Agents[0].DpsAvg,
		DpsStdev:            result.Agents[0].DpsStDev,
		DpsHist:             result.Agents[0].DpsHist,
		Logs:                result.Logs,
		DpsMax:              result.Agents[0].DpsMax,
		ExecutionDurationMs: result.ExecutionDurationMs,
		NumOom:              int32(result.Agents[0].NumOom),
		OomAtAvg:            result.Agents[0].OomAtAvg,
		DpsAtOomAvg:         result.Agents[0].DpsAtOomAvg,
		ActionMetrics:       actionMetrics,
	}
	return isr
}
