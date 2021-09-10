package runner

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/wowsims/tbc/sim/core"
)

func RunSim(raid *core.Raid, buffs core.Buffs, options core.Options) {
	// TODO: Create buff bot players here
	// TODO: Since buff bots will just be implemented separate from the real
	//   agent for the class... maybe just add a public "ApplyXX" function for teh buffs

	// TODO: Move these to an "item buff bot"?
	// Or just the first party member other than you? (create item buff bot if not i guess)
	// if b.TwilightOwl {
	// 	s[StatSpellCrit] += 44.16
	// }
	// if b.EyeOfNight {
	// 	s[StatSpellPower] += 34
	// }

	sim := core.NewSim(raid, options)

	// if sim.Options.Buffs.Misery {
	// 	sim.cache.elcDmgBonus += 0.05
	// 	sim.cache.dmgBonus += 0.05
	// }

	// Buffs    Buffs
	// sim := NewSim(request.Gear, request.Options)

	logsBuffer := &strings.Builder{}
	aggregator := NewMetricsAggregator()

	if options.Debug {
		sim.Debug = func(s string, vals ...interface{}) {
			logsBuffer.WriteString(fmt.Sprintf("[%0.1f] "+s, append([]interface{}{sim.CurrentTime.Seconds()}, vals...)...))
		}
	}

	for i := int32(0); i < options.Iterations; i++ {
		metrics := sim.Run()
		aggregator.addMetrics(options, metrics)
		sim.ReturnCasts(metrics.Casts)
	}

	// result := aggregator.getResult()
	// result.Logs = logsBuffer.String()
	// return result
}

type MetricsAggregator struct {
	startTime time.Time
	numSims   int

	dpsSum        float64
	dpsSumSquared float64
	dpsMax        float64
	dpsHist       map[int]int // rounded DPS to count

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
	DpsHist  map[int]int // rounded DPS to count

	NumOom      int
	OomAtAvg    float64
	DpsAtOomAvg float64

	Casts map[int32]CastMetric
}

type CastMetric struct {
	Counts []int32
	Dmgs   []float64
	Tags   []int32
}

func NewMetricsAggregator() *MetricsAggregator {
	return &MetricsAggregator{
		startTime: time.Now(),
		dpsHist:   make(map[int]int),
		casts:     make(map[int32]CastMetric),
	}
}

func (aggregator *MetricsAggregator) addMetrics(options core.Options, metrics core.SimMetrics) {
	aggregator.numSims++

	dps := metrics.TotalDamage / options.Encounter.Duration

	aggregator.dpsSum += dps
	aggregator.dpsSumSquared += dps * dps
	aggregator.dpsMax = math.Max(aggregator.dpsMax, dps)

	dpsRounded := int(math.Round(dps/10) * 10)
	aggregator.dpsHist[dpsRounded]++

	// TODO: Fix me
	// if metrics.OOMAt > 0 {
	// 	aggregator.numOom++
	// 	aggregator.oomAtSum += float64(metrics.OOMAt)
	// 	aggregator.dpsAtOomSum += float64(metrics.DamageAtOOM) / float64(metrics.OOMAt)
	// }

	for _, cast := range metrics.Casts {
		var id = cast.Spell.ID
		cm := aggregator.casts[id]
		idx := 0
		if cast.IsLO {
			idx = 2
		} else if cast.DidCrit {
			idx = 1
		}
		cm.Counts[idx]++
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
