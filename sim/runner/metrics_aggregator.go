package runner

import (
	"math"
	"time"

	"github.com/wowsims/tbc/sim/core"
)

type MetricsAggregator struct {
	startTime time.Time
	numSims   int

	dpsSum        float64
	dpsSumSquared float64
	dpsMax        float64
	dpsHist       map[int32]int32 // rounded DPS to count

	numOom      int
	oomAtSum    float64
	dpsAtOomSum float64

	casts map[int32]CastMetric
}

type SimResult struct {
	ExecutionDurationMs int64
	Logs                string

	DpsAvg   float64
	DpsStDev float64
	DpsMax   float64
	DpsHist  map[int32]int32 // rounded DPS to count

	NumOom      int
	OomAtAvg    float64
	DpsAtOomAvg float64

	Casts map[int32]CastMetric
}

type CastMetric struct {
	// Index 0 of each slice is the 'normal' cast data.
	// Count & Dmg of spells cast by Tag
	Casts  []int32 // Total Count of Casts
	Crits  []int32 // Count of Crits
	Misses []int32 // Count of Misses
	// Resists []int32   // Count of Resists
	Dmgs []float64 // Total Damage
}

func NewMetricsAggregator() *MetricsAggregator {
	return &MetricsAggregator{
		startTime: time.Now(),
		dpsHist:   make(map[int32]int32),
		casts:     make(map[int32]CastMetric),
	}
}

func (aggregator *MetricsAggregator) addMetrics(options core.Options, metrics core.SimMetrics) {
	aggregator.numSims++

	dps := metrics.TotalDamage / options.Encounter.Duration
	// log.Printf("total: %0.1f, dur: %0.1f, dps: %0.1f", metrics.TotalDamage, options.Encounter.Duration, dps)

	aggregator.dpsSum += dps
	aggregator.dpsSumSquared += dps * dps
	aggregator.dpsMax = math.Max(aggregator.dpsMax, dps)

	dpsRounded := int32(math.Round(dps/10) * 10)
	aggregator.dpsHist[dpsRounded]++

	// TODO: Fix me
	firstPlayer := metrics.IndividualMetrics[0]
	if firstPlayer.OOMAt > 0 {
		aggregator.numOom++
		aggregator.oomAtSum += float64(firstPlayer.OOMAt)
		aggregator.dpsAtOomSum += float64(firstPlayer.DamageAtOOM) / float64(firstPlayer.OOMAt)
	}

	for _, cast := range metrics.Casts {
		var id = cast.Spell.ID
		cm := aggregator.casts[id]
		idx := int(cast.Tag)

		if len(cm.Casts) <= idx {
			newArr := make([]int32, idx+1)
			copy(newArr, cm.Casts)
			cm.Casts = newArr

			newDmgs := make([]float64, idx+1)
			copy(newDmgs, cm.Dmgs)
			cm.Dmgs = newDmgs
		}
		cm.Casts[idx]++
		if cast.DidCrit {
			if len(cm.Crits) <= idx {
				newArr := make([]int32, idx+1)
				copy(newArr, cm.Crits)
				cm.Crits = newArr
			}
			cm.Crits[idx]++
		} else if !cast.DidHit {
			if len(cm.Misses) <= idx {
				newArr := make([]int32, idx+1)
				copy(newArr, cm.Misses)
				cm.Misses = newArr
			}
			cm.Misses[idx]++
		}
		cm.Dmgs[idx] += cast.DidDmg
		aggregator.casts[id] = cm
	}
}

func (aggregator *MetricsAggregator) getResult() SimResult {
	result := SimResult{}
	result.ExecutionDurationMs = time.Since(aggregator.startTime).Milliseconds()

	numSims := float64(aggregator.numSims)
	result.DpsAvg = aggregator.dpsSum / numSims
	result.DpsStDev = math.Sqrt((aggregator.dpsSumSquared / numSims) - (result.DpsAvg * result.DpsAvg))
	result.DpsMax = aggregator.dpsMax
	result.DpsHist = aggregator.dpsHist

	result.NumOom = aggregator.numOom
	if result.NumOom > 0 {
		result.OomAtAvg = aggregator.oomAtSum / float64(aggregator.numOom)
		result.DpsAtOomAvg = aggregator.dpsAtOomSum / float64(aggregator.numOom)
	}

	result.Casts = aggregator.casts

	return result
}
