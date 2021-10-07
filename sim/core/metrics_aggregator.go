package core

import (
	"math"
	"time"
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

	// Actions metrics
	// IDs can overlap so create separate maps
	// We could probably use a couple bits of a map[int64] so we only need 1 map.
	casts map[int32]ActionMetric
	items map[int32]ActionMetric
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

	Actions []ActionMetric
}

type ActionMetric struct {
	ActionID ActionID

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
		casts:     make(map[int32]ActionMetric),
		items:     make(map[int32]ActionMetric),
	}
}

func (aggregator *MetricsAggregator) addMetrics(options Options, metrics SimMetrics) {
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
		var metric ActionMetric
		if cast.Spell.ActionID.SpellID != 0 {
			metric = aggregator.casts[cast.Spell.ActionID.SpellID]
		} else if cast.Spell.ActionID.ItemID != 0 {
			metric = aggregator.items[cast.Spell.ActionID.ItemID]
		}
		idx := int(cast.Tag)

		// Construct new arrays for a tag we haven't seen before.
		if len(metric.Casts) <= idx {
			newArr := make([]int32, idx+1)
			copy(newArr, metric.Casts)
			metric.Casts = newArr

			newCritsArr := make([]int32, idx+1)
			copy(newCritsArr, metric.Crits)
			metric.Crits = newCritsArr

			newDmgs := make([]float64, idx+1)
			copy(newDmgs, metric.Dmgs)
			metric.Dmgs = newDmgs

			newMissArr := make([]int32, idx+1)
			copy(newMissArr, metric.Misses)
			metric.Misses = newMissArr
		}

		metric.Casts[idx]++
		if cast.DidCrit {
			metric.Crits[idx]++
		} else if !cast.DidHit {
			metric.Misses[idx]++
		}
		metric.Dmgs[idx] += cast.DidDmg

		if cast.Spell.ActionID.SpellID != 0 {
			aggregator.casts[cast.Spell.ActionID.SpellID] = metric
		} else if cast.Spell.ActionID.ItemID != 0 {
			aggregator.items[cast.Spell.ActionID.ItemID] = metric
		}
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

	for id, v := range aggregator.casts {
		v.ActionID.SpellID = id
		result.Actions = append(result.Actions, v)
	}

	for id, v := range aggregator.items {
		v.ActionID.ItemID = id
		result.Actions = append(result.Actions, v)
	}

	return result
}
