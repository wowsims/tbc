// Proto based function interface for the simulator
package sim

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	RegisterAll()
}

func getGearListImpl(request *proto.GearListRequest) *proto.GearListResult {
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

func computeStatsImpl(request *proto.ComputeStatsRequest) *proto.ComputeStatsResult {
	return statsFromIndSimRequest(&proto.IndividualSimRequest{Player: request.Player, Buffs: request.Buffs})
}

func statsFromIndSimRequest(isr *proto.IndividualSimRequest) *proto.ComputeStatsResult {
	sim := createSim(isr)
	gearStats := sim.Raid.Parties[0].Players[0].GetCharacter().Equip.Stats()
	return &proto.ComputeStatsResult{
		GearOnly:   gearStats[:],
		FinalStats: sim.Raid.Parties[0].Players[0].GetCharacter().Stats[:], // createSim includes a call to buff up all party members.
		Sets:       []string{},
	}
}

func statWeightsImpl(request *proto.StatWeightsRequest) *proto.StatWeightsResult {
	statsToWeight := make([]stats.Stat, len(request.StatsToWeigh))
	for i, v := range request.StatsToWeigh {
		statsToWeight[i] = stats.Stat(v)
	}
	result := core.CalcStatWeight(convertSimParams(request.Options), statsToWeight, stats.Stat(request.EpReferenceStat))
	return &proto.StatWeightsResult{
		Weights:       result.Weights[:],
		WeightsStdev:  result.WeightsStdev[:],
		EpValues:      result.EpValues[:],
		EpValuesStdev: result.EpValuesStdev[:],
	}
}

func convertSimParams(request *proto.IndividualSimRequest) core.IndividualParams {
	options := core.Options{
		Iterations: int(request.Iterations),
		RSeed:      request.RandomSeed,
		ExitOnOOM:  request.ExitOnOom,
		GCDMin:     time.Duration(request.GcdMin),
		Debug:      request.Debug,
	}
	if request.Encounter != nil {
		options.Encounter = core.Encounter{
			Duration:   request.Encounter.Duration,
			NumTargets: int32(request.Encounter.NumTargets),
			Armor:      request.Encounter.TargetArmor,
		}
	}

	params := core.IndividualParams{
		Equip:    items.ProtoToEquipmentSpec(request.Player.Equipment),
		Race:     core.RaceBonusType(request.Player.Options.Race),
		Consumes: core.ProtoToConsumes(request.Player.Options.Consumes),
		Buffs:    core.ProtoToBuffs(request.Buffs),
		Options:  options,

		PlayerOptions: request.Player.Options,
	}
	copy(params.CustomStats[:], request.Player.CustomStats[:])

	return params
}

func createSim(request *proto.IndividualSimRequest) *core.Simulation {
	params := convertSimParams(request)
	sim := core.NewIndividualSim(params)
	return sim
}

func runSimulationImpl(request *proto.IndividualSimRequest) *proto.IndividualSimResult {
	sim := createSim(request)
	result := sim.Run()

	actionMetrics := []*proto.ActionMetric{}
	for _, v := range result.Actions {
		metric := &proto.ActionMetric{
			Casts:  v.Casts,
			Crits:  v.Crits,
			Misses: v.Misses,
			Dmgs:   v.Dmgs,
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
		DpsAvg:              result.DpsAvg,
		DpsStdev:            result.DpsStDev,
		DpsHist:             result.DpsHist,
		Logs:                result.Logs,
		DpsMax:              result.DpsMax,
		ExecutionDurationMs: result.ExecutionDurationMs,
		NumOom:              int32(result.NumOom),
		OomAtAvg:            result.OomAtAvg,
		DpsAtOomAvg:         result.DpsAtOomAvg,
		ActionMetrics:       actionMetrics,
	}
	return isr
}
