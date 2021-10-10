package core

import (
	"math"
	"time"
)

type MetricsAggregator struct {
	// Duration of each iteration, in seconds
	encounterDuration float64

	startTime time.Time
	numIterations  int

	// Metrics for each Agent, for the current iteration
	agentIterations []AgentIterationMetrics

	// Aggregate values for each agent over all iterations
	agentAggregates []AgentAggregateMetrics
}

type AgentIterationMetrics struct {
	TotalDamage float64
	ManaSpent   float64
	DamageAtOOM float64
	OOMAt       time.Duration
}

type AgentAggregateMetrics struct {
	DpsSum        float64
	DpsSumSquared float64
	DpsMax        float64
	DpsHist       map[int32]int32 // rounded DPS to count

	NumOom      int
	OomAtSum    float64
	DpsAtOomSum float64

	Actions map[int32]ActionMetric
}

type SimResult struct {
	ExecutionDurationMs int64
	Logs                string

	Agents []AgentResult
}

type AgentResult struct {
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
	Hits   []int32 // Count of Hits
	Crits  []int32 // Count of Crits
	Misses []int32 // Count of Misses
	// Resists []int32   // Count of Resists
	Dmgs []float64 // Total Damage
}

func NewMetricsAggregator(numAgents int, encounterDuration float64) *MetricsAggregator {
	aggregator := &MetricsAggregator{
		encounterDuration: encounterDuration,
		startTime: time.Now(),
	}

	for i := 0; i < numAgents; i++ {
		aggregator.agentIterations = append(aggregator.agentIterations, AgentIterationMetrics{})
		aggregator.agentAggregates = append(aggregator.agentAggregates, AgentAggregateMetrics{})

		aggregator.agentAggregates[i].Actions = make(map[int32]ActionMetric)
		aggregator.agentAggregates[i].DpsHist = make(map[int32]int32)
	}

	return aggregator
}

// Adds the results of an action to the aggregated metrics.
func (aggregator *MetricsAggregator) addCastAction(cast DirectCastAction, castResults []DirectCastDamageResult) {
	actionID := cast.GetActionID()

	// This works by making item IDs negative to avoid collisions, and assumes
	// there are no collisions with OtherID.
	// Actual key values dont matter, just need something unique and fast to compute.
	actionKey := int32(actionID.OtherID) + actionID.SpellID - actionID.ItemID

	agentID := cast.GetAgent().GetCharacter().ID

	iterationMetrics := &aggregator.agentIterations[agentID]
	iterationMetrics.ManaSpent += cast.castInput.ManaCost

	aggregateMetrics := &aggregator.agentAggregates[agentID]
	actionMetrics := aggregateMetrics.Actions[actionKey]

	tag := int(cast.GetTag())
	// Construct new arrays for a tag we haven't seen before.
	if len(actionMetrics.Casts) <= tag {
		actionMetrics.ActionID = actionID

		newArr := make([]int32, tag+1)
		copy(newArr, actionMetrics.Casts)
		actionMetrics.Casts = newArr

		newHitsArr := make([]int32, tag+1)
		copy(newHitsArr, actionMetrics.Hits)
		actionMetrics.Hits = newHitsArr

		newCritsArr := make([]int32, tag+1)
		copy(newCritsArr, actionMetrics.Crits)
		actionMetrics.Crits = newCritsArr

		newDmgs := make([]float64, tag+1)
		copy(newDmgs, actionMetrics.Dmgs)
		actionMetrics.Dmgs = newDmgs

		newMissArr := make([]int32, tag+1)
		copy(newMissArr, actionMetrics.Misses)
		actionMetrics.Misses = newMissArr
	}

	actionMetrics.Casts[tag]++
	for _, result := range castResults {
		if result.Crit {
			actionMetrics.Crits[tag]++
		} else if result.Hit {
			actionMetrics.Hits[tag]++
		} else {
			actionMetrics.Misses[tag]++
		}
		actionMetrics.Dmgs[tag] += result.Damage
		iterationMetrics.TotalDamage += result.Damage
	}

	aggregateMetrics.Actions[actionKey] = actionMetrics
}

func (aggregator *MetricsAggregator) markOOM(agent Agent, oomAtTime time.Duration) {
	agentID := agent.GetCharacter().ID

	if aggregator.agentIterations[agentID].OOMAt == 0 {
		aggregator.agentIterations[agentID].DamageAtOOM = aggregator.agentIterations[agentID].TotalDamage
		aggregator.agentIterations[agentID].OOMAt = oomAtTime
	}
}

// This should be called when a Sim iteration is complete.
func (aggregator *MetricsAggregator) doneIteration() {
	aggregator.numIterations++

	// Loop for each agent
	for i, iterationMetrics := range aggregator.agentIterations {
		aggregateMetrics := &aggregator.agentAggregates[i]

		dps := iterationMetrics.TotalDamage / aggregator.encounterDuration
		// log.Printf("total: %0.1f, dur: %0.1f, dps: %0.1f", metrics.TotalDamage, aggregator.encounterDuration, dps)

		aggregateMetrics.DpsSum += dps
		aggregateMetrics.DpsSumSquared += dps * dps
		aggregateMetrics.DpsMax = MaxFloat(aggregateMetrics.DpsMax, dps)

		dpsRounded := int32(math.Round(dps/10) * 10)
		aggregateMetrics.DpsHist[dpsRounded]++

		if iterationMetrics.OOMAt > 0 {
			aggregateMetrics.NumOom++
			aggregateMetrics.OomAtSum += float64(iterationMetrics.OOMAt)
			aggregateMetrics.DpsAtOomSum += float64(iterationMetrics.DamageAtOOM) / float64(iterationMetrics.OOMAt.Seconds())
		}

		// Clear the iteration metrics
		aggregator.agentIterations[i] = AgentIterationMetrics{}
	}
}

func (aggregator *MetricsAggregator) getResult() SimResult {
	result := SimResult{}
	result.ExecutionDurationMs = time.Since(aggregator.startTime).Milliseconds()

	numIterations := float64(aggregator.numIterations)
	numAgents := len(aggregator.agentAggregates)

	result.Agents = make([]AgentResult, numAgents)
	for i := 0; i < numAgents; i++ {
		agentAggregate := &aggregator.agentAggregates[i]
		agentResult := &result.Agents[i]

		agentResult.DpsAvg = agentAggregate.DpsSum / numIterations
		agentResult.DpsStDev = math.Sqrt((agentAggregate.DpsSumSquared / numIterations) - (agentResult.DpsAvg * agentResult.DpsAvg))
		agentResult.DpsMax = agentAggregate.DpsMax
		agentResult.DpsHist = agentAggregate.DpsHist

		agentResult.NumOom = agentAggregate.NumOom
		if agentResult.NumOom > 0 {
			agentResult.OomAtAvg = agentAggregate.OomAtSum / float64(agentAggregate.NumOom)
			agentResult.DpsAtOomAvg = agentAggregate.DpsAtOomSum / float64(agentAggregate.NumOom)
		}

		for _, action := range agentAggregate.Actions {
			agentResult.Actions = append(agentResult.Actions, action)
		}
	}

	return result
}
