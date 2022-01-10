package core

import (
	"math"
	"math/rand"
	"sync"
	"sync/atomic"

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

func CalcStatWeight(swr proto.StatWeightsRequest, statsToWeigh []stats.Stat, referenceStat stats.Stat, progress chan *proto.ProgressMetrics) StatWeightsResult {
	if swr.Player.BonusStats == nil {
		swr.Player.BonusStats = make([]float64, stats.Len)
	}

	raidProto := SinglePlayerRaidProto(swr.Player, swr.PartyBuffs, swr.RaidBuffs)
	baseStatsResult := ComputeStats(&proto.ComputeStatsRequest{
		Raid: raidProto,
	})
	baseStats := baseStatsResult.RaidStats.Parties[0].Players[0].FinalStats

	baseSimRequest := &proto.RaidSimRequest{
		Raid:       raidProto,
		Encounter:  swr.Encounter,
		SimOptions: swr.SimOptions,
	}
	baselineResult := RunRaidSim(baseSimRequest)
	baselineDpsMetrics := baselineResult.RaidMetrics.Parties[0].Players[0].Dps

	var waitGroup sync.WaitGroup

	// Do half the iterations with a positive, and half with a negative value for better accuracy.
	resultLow := StatWeightsResult{}
	resultHigh := StatWeightsResult{}
	dpsHistsLow := [stats.Len]map[int32]int32{}
	dpsHistsHigh := [stats.Len]map[int32]int32{}

	var iterationsTotal int32
	var iterationsDone int32

	doStat := func(stat stats.Stat, value float64, isLow bool) {
		defer waitGroup.Done()

		simRequest := googleProto.Clone(baseSimRequest).(*proto.RaidSimRequest)
		simRequest.Raid.Parties[0].Players[0].BonusStats[stat] += value
		simRequest.SimOptions.Iterations /= 2 // Cut in half since we're doing above and below separately.

		reporter := make(chan *proto.ProgressMetrics, 10)
		go RunSim(*simRequest, reporter) // RunRaidSim(simRequest)

		var localIterations int32
		var simResult *proto.RaidSimResult
	statsim:
		for {
			select {
			case metrics, ok := <-reporter:
				if !ok {
					break statsim
				}
				if metrics.FinalResult != nil {
					simResult = metrics.FinalResult
					break statsim
				}
				atomic.AddInt32(&iterationsDone, (metrics.CompletedIterations - localIterations))
				localIterations = metrics.CompletedIterations
			}
		}
		dpsMetrics := simResult.RaidMetrics.Parties[0].Players[0].Dps
		dpsDiff := (dpsMetrics.Avg - baselineDpsMetrics.Avg) / value

		if isLow {
			resultLow.Weights[stat] = dpsDiff
			dpsHistsLow[stat] = dpsMetrics.Hist
		} else {
			resultHigh.Weights[stat] = dpsDiff
			dpsHistsHigh[stat] = dpsMetrics.Hist
		}
	}

	const defaultStatMod = 50.0
	statModsLow := stats.Stats{}
	statModsHigh := stats.Stats{}

	// Make sure reference stat is included.
	statModsLow[referenceStat] = defaultStatMod
	statModsHigh[referenceStat] = defaultStatMod

	for _, v := range statsToWeigh {
		statMod := defaultStatMod
		if v == stats.SpellHit || v == stats.MeleeHit {
			// For spell/melee hit, always pick the direction which is gauranteed to
			// not run into a hit cap.
			if baseStats[v] < 80 {
				statModsHigh[v] = 10
				statModsLow[v] = 10
			} else {
				statModsHigh[v] = -10
				statModsLow[v] = -10
			}
		} else {
			statModsHigh[v] = statMod
			statModsLow[v] = -statMod
		}
	}

	for stat, _ := range statModsLow {
		if statModsLow[stat] == 0 {
			continue
		}
		waitGroup.Add(2)
		iterationsTotal += swr.SimOptions.Iterations

		go doStat(stats.Stat(stat), statModsLow[stat], true)
		go doStat(stats.Stat(stat), statModsHigh[stat], false)
	}

	waitGroup.Wait()

	result := StatWeightsResult{}
	for statIdx, _ := range statModsLow {
		stat := stats.Stat(statIdx)
		if statModsLow[stat] == 0 {
			continue
		}
		result.Weights[stat] = (resultLow.Weights[stat] + resultHigh.Weights[stat]) / 2
	}

	for statIdx, _ := range statModsLow {
		stat := stats.Stat(statIdx)
		if statModsLow[stat] == 0 {
			continue
		}

		result.EpValues[stat] = result.Weights[stat] / result.Weights[referenceStat]

		weightStdevLow := computeStDevFromHists(swr.SimOptions.Iterations/2, statModsLow[stat], dpsHistsLow[stat], baselineDpsMetrics.Hist, nil, statModsLow[referenceStat])
		weightStdevHigh := computeStDevFromHists(swr.SimOptions.Iterations/2, statModsHigh[stat], dpsHistsHigh[stat], baselineDpsMetrics.Hist, nil, statModsHigh[referenceStat])
		result.WeightsStdev[stat] = (weightStdevLow + weightStdevHigh) / 2

		epStdevLow := computeStDevFromHists(swr.SimOptions.Iterations/2, statModsLow[stat], dpsHistsLow[stat], baselineDpsMetrics.Hist, dpsHistsLow[referenceStat], statModsLow[referenceStat])
		epStdevHigh := computeStDevFromHists(swr.SimOptions.Iterations/2, statModsHigh[stat], dpsHistsHigh[stat], baselineDpsMetrics.Hist, dpsHistsHigh[referenceStat], statModsHigh[referenceStat])
		result.EpValuesStdev[stat] = (epStdevLow + epStdevHigh) / 2
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
