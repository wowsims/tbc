package core

import (
	"math"
	"math/rand"
	"sync"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
	googleProto "google.golang.org/protobuf/proto"
)

type StatWeightsResult struct {
	Weights       stats.Stats
	WeightsStdev  stats.Stats
	EpValues      stats.Stats
	EpValuesStdev stats.Stats
}

func CalcStatWeight(isr proto.IndividualSimRequest, statsToWeigh []stats.Stat, referenceStat stats.Stat) StatWeightsResult {
	if isr.Player.BonusStats == nil {
		isr.Player.BonusStats = make([]float64, stats.Len)
	}

	baseStatsResult := ComputeStats(&proto.ComputeStatsRequest{
		Player:     isr.Player,
		RaidBuffs:  isr.RaidBuffs,
		PartyBuffs: isr.PartyBuffs,
	})
	baseStats := baseStatsResult.FinalStats

	baselineResult := RunIndividualSim(&isr)

	var waitGroup sync.WaitGroup
	result := StatWeightsResult{}
	dpsHists := [stats.Len]map[int32]int32{}

	doStat := func(stat stats.Stat, value float64) {
		defer waitGroup.Done()

		simRequest := googleProto.Clone(&isr).(*proto.IndividualSimRequest)
		simRequest.Player.BonusStats[stat] += value

		simResult := RunIndividualSim(simRequest)

		result.Weights[stat] = (simResult.PlayerMetrics.DpsAvg - baselineResult.PlayerMetrics.DpsAvg) / value
		dpsHists[stat] = simResult.PlayerMetrics.DpsHist
	}

	// Spell hit mod shouldn't go over hit cap.
	spellHitMod := math.Max(0, math.Min(10, 202-baseStats[stats.SpellHit]))

	statMods := stats.Stats{}
	statMods[referenceStat] = 50 // make sure reference stat is included
	for _, v := range statsToWeigh {
		statMods[v] = 50
		if v == stats.SpellHit {
			statMods[v] = spellHitMod
		}
	}
	for stat, mod := range statMods {
		if mod == 0 {
			continue
		}
		waitGroup.Add(1)
		go doStat(stats.Stat(stat), mod)
	}

	waitGroup.Wait()

	for statIdx, mod := range statMods {
		if mod == 0 {
			continue
		}
		stat := stats.Stat(statIdx)

		result.EpValues[stat] = result.Weights[stat] / result.Weights[referenceStat]
		result.WeightsStdev[stat] = computeStDevFromHists(isr.SimOptions.Iterations, mod, dpsHists[stat], baselineResult.PlayerMetrics.DpsHist, nil, statMods[referenceStat])
		result.EpValuesStdev[stat] = computeStDevFromHists(isr.SimOptions.Iterations, mod, dpsHists[stat], baselineResult.PlayerMetrics.DpsHist, dpsHists[referenceStat], statMods[referenceStat])
	}

	return result
}

func computeStDevFromHists(iters int32, modValue float64, moddedStatDpsHist map[int32]int32, baselineDpsHist map[int32]int32, referenceDpsHist map[int32]int32, referenceModValue float64) float64 {
	sum := 0.0
	sumSquared := 0.0
	n := iters * 10
	for i := int32(0); i < n; {
		denominator := 1.0
		if referenceDpsHist != nil {
			denominator = float64(sampleFromDpsHist(referenceDpsHist, iters)-sampleFromDpsHist(baselineDpsHist, iters)) / referenceModValue
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

// Picks a random value from a histogram, taking into account the bucket sizes.
func sampleFromDpsHist(hist map[int32]int32, histNumSamples int32) int32 {
	r := rand.Float64()
	sampleIdx := int32(math.Floor(float64(histNumSamples) * r))

	var curSampleIdx int32
	for roundedDps, count := range hist {
		curSampleIdx += count
		if curSampleIdx >= sampleIdx {
			return roundedDps
		}
	}

	panic("Invalid dps histogram")
}
