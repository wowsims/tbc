// Top-level implementations for the go functions.
package api

import (
	"math"
	"math/rand"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/runner"
	"github.com/wowsims/tbc/sim/shaman"
)

func getGearListImpl(request *GearListRequest) *GearListResult {
	result := &GearListResult{}

	if request.Spec == Spec_elemental_shaman {
		for _, v := range shaman.ElementalItems {
			item := core.ItemsByID[v]
			result.Items = append(result.Items,
				&Item{
					Id:      item.ID,
					Slot:    int32(item.Slot),
					SubSlot: int32(item.SubSlot),
					Name:    item.Name,
					// Stats:       item.Stats,
					Phase:   int32(item.Phase),
					Quality: ItemQuality(item.Quality),
					// GemSlots:    item.GemSlots,
					// SocketBonus: item.SocketBonus,
				},
			)
		}
	}
	// Items:    items,
	// Enchants: Enchants,
	// Gems:     Gems,

	return result
}

func computeStatsImpl(request *ComputeStatsRequest) *ComputeStatsResult {
	panic("not implemented")
	// fakeSim := core.NewSim(request.Gear, request.Options)

	// sets := fakeSim.ActivateSets()
	// fakeSim.reset() // this will activate any perm-effect items as well

	// gearOnlyStats := fakeSim.Equip.Stats().CalculatedTotal()
	// finalStats := fakeSim.Stats

	// return &ComputeStatsResult{
	// 	GearOnly:   gearOnlyStats,
	// 	FinalStats: finalStats,
	// 	Sets:       sets,
	// }
}

func statWeightsImpl(request *StatWeightsRequest) *StatWeightsResult {
	panic("not implemented")

	// request.Options.AgentType = AGENT_TYPE_ADAPTIVE

	// baselineSimRequest := SimRequest{
	// 	Options:    request.Options,
	// 	Gear:       request.Gear,
	// 	Iterations: request.Iterations,
	// }
	// baselineResult := RunSimulation(baselineSimRequest)

	// var waitGroup sync.WaitGroup
	// result := StatWeightsResult{}
	// dpsHists := [StatLen]map[int]int{}

	// doStat := func(stat Stat, value float64) {
	// 	defer waitGroup.Done()

	// 	simRequest := baselineSimRequest
	// 	simRequest.Options.Buffs.Custom[stat] += value

	// 	simResult := RunSimulation(simRequest)
	// 	result.Weights[stat] = (simResult.DpsAvg - baselineResult.DpsAvg) / value
	// 	dpsHists[stat] = simResult.DpsHist
	// }

	// // Spell hit mod shouldn't go over hit cap.
	// computeStatsResult := ComputeStats(ComputeStatsRequest{
	// 	Options: request.Options,
	// 	Gear:    request.Gear,
	// })
	// spellHitMod := math.Max(0, math.Min(10, 202-computeStatsResult.FinalStats[StatSpellHit]))

	// statMods := Stats{
	// 	StatInt:       50,
	// 	StatSpellDmg:  50,
	// 	StatSpellCrit: 50,
	// 	StatSpellHit:  spellHitMod,
	// 	StatHaste:     50,
	// 	StatMP5:       50,
	// }

	// for stat, mod := range statMods {
	// 	if mod == 0 {
	// 		continue
	// 	}

	// 	waitGroup.Add(1)
	// 	go doStat(Stat(stat), mod)
	// }

	// waitGroup.Wait()

	// for stat, mod := range statMods {
	// 	if mod == 0 {
	// 		continue
	// 	}

	// 	result.EpValues[stat] = result.Weights[stat] / result.Weights[StatSpellPower]
	// 	result.WeightsStDev[stat] = computeStDevFromHists(request.Iterations, mod, dpsHists[stat], baselineResult.DpsHist, nil, statMods[StatSpellDmg])
	// 	result.EpValuesStDev[stat] = computeStDevFromHists(request.Iterations, mod, dpsHists[stat], baselineResult.DpsHist, dpsHists[StatSpellDmg], statMods[StatSpellDmg])
	// }
	// return result
}

func computeStDevFromHists(iters int, modValue float64, moddedStatDpsHist map[int]int, baselineDpsHist map[int]int, spellDmgDpsHist map[int]int, spellDmgModValue float64) float64 {
	sum := 0.0
	sumSquared := 0.0
	n := iters * 10
	for i := 0; i < n; {
		denominator := 1.0
		if spellDmgDpsHist != nil {
			denominator = float64(sampleFromDpsHist(spellDmgDpsHist, iters)-sampleFromDpsHist(baselineDpsHist, iters)) / spellDmgModValue
		}

		if denominator != 0 {
			ep := (float64(sampleFromDpsHist(moddedStatDpsHist, iters)-sampleFromDpsHist(baselineDpsHist, iters)) / modValue) / denominator
			sum += ep
			sumSquared += ep * ep
			i++
		}
	}
	epAvg := sum / float64(n)
	epStDev := math.Sqrt((sumSquared / float64(n)) - (epAvg * epAvg))
	return epStDev
}

func sampleFromDpsHist(hist map[int]int, histNumSamples int) int {
	r := rand.Float64()
	sampleIdx := int(math.Floor(float64(histNumSamples) * r))

	curSampleIdx := 0
	for roundedDps, count := range hist {
		curSampleIdx += count
		if curSampleIdx >= sampleIdx {
			return roundedDps
		}
	}

	panic("Invalid dps histogram")
}

func runSimulationImpl(request *IndividualSimRequest) *IndividualSimResult {
	// panic("not implemented")

	player := core.NewPlayer(core.EquipmentSpec{}, core.Consumes{})

	var agent core.Agent
	switch v := request.Player.Options.Class.(type) {
	case *PlayerOptions_Shaman:
		agent = shaman.NewShaman(player, int(v.Shaman.AgentType), v.Shaman.AgentOptions)
	default:
		panic("class not supported")
	}

	raid := &core.Raid{
		Parties: []*core.Party{
			{
				Players: []core.PlayerAgent{
					{Player: player, Agent: agent},
				},
			},
		},
	}

	options := core.Options{}

	isr := &IndividualSimResult{
		// ExecutionDurationMs int64                 `protobuf:"varint,1,opt,name=execution_duration_ms,json=executionDurationMs,proto3" json:"execution_duration_ms,omitempty"`
		// Logs                string                `protobuf:"bytes,2,opt,name=logs,proto3" json:"logs,omitempty"`
		// DpsAvg              float64               `protobuf:"fixed64,3,opt,name=dps_avg,json=dpsAvg,proto3" json:"dps_avg,omitempty"`
		// DpsStdev            float64               `protobuf:"fixed64,4,opt,name=dps_stdev,json=dpsStdev,proto3" json:"dps_stdev,omitempty"`
		// DpsMax              float64               `protobuf:"fixed64,5,opt,name=dps_max,json=dpsMax,proto3" json:"dps_max,omitempty"`
		// DpsHist             map[int32]int32       `protobuf:"bytes,6,rep,name=dps_hist,json=dpsHist,proto3" json:"dps_hist,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
		// NumOom              int32                 `protobuf:"varint,7,opt,name=num_oom,json=numOom,proto3" json:"num_oom,omitempty"`
		// OomAtAvg            float64               `protobuf:"fixed64,8,opt,name=oom_at_avg,json=oomAtAvg,proto3" json:"oom_at_avg,omitempty"`
		// DpsAtOomAvg         float64               `protobuf:"fixed64,9,opt,name=dps_at_oom_avg,json=dpsAtOomAvg,proto3" json:"dps_at_oom_avg,omitempty"`
		// Casts               map[int32]*CastMetric `protobuf:"bytes,10,rep,name=casts,proto3" json:"casts,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	}
	runner.RunSim(raid, options)

	return isr
}
