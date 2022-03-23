package core

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"strings"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
)

type InitialAura func(*Simulation) Aura

type Simulation struct {
	Raid              *Raid
	encounter         Encounter
	Options           proto.SimOptions
	BaseDuration      time.Duration // base duration
	DurationVariation time.Duration // variation per duration
	Duration          time.Duration // Duration of current iteration

	rand *rand.Rand

	// Used for testing only, see RandomFloat().
	isTest    bool
	testRands map[string]*rand.Rand

	// Current Simulation State
	pendingActions    []*PendingAction
	pendingActionPool *paPool
	CurrentTime       time.Duration // duration that has elapsed in the sim since starting

	ProgressReport func(*proto.ProgressMetrics)

	Log  func(string, ...interface{})
	logs []string

	emptyAuras []Aura
}

func RunSim(rsr proto.RaidSimRequest, progress chan *proto.ProgressMetrics) *proto.RaidSimResult {
	sim := NewSim(rsr)
	sim.runPresims(rsr)
	if progress != nil {
		sim.ProgressReport = func(progMetric *proto.ProgressMetrics) {
			progress <- progMetric
		}
	}
	return sim.run()
}

func NewSim(rsr proto.RaidSimRequest) *Simulation {
	raid := NewRaid(*rsr.Raid)
	encounter := NewEncounter(*rsr.Encounter)
	simOptions := *rsr.SimOptions

	if len(encounter.Targets) == 0 {
		panic("Must have at least 1 target!")
	}

	rseed := simOptions.RandomSeed
	if rseed == 0 {
		rseed = time.Now().Unix()
	}

	return &Simulation{
		Raid:              raid,
		encounter:         encounter,
		Options:           simOptions,
		BaseDuration:      encounter.Duration,
		DurationVariation: encounter.DurationVariation,
		Log:               nil,

		rand: rand.New(rand.NewSource(rseed)),

		isTest:    simOptions.IsTest,
		testRands: make(map[string]*rand.Rand),

		emptyAuras: make([]Aura, numAuraIDs),

		pendingActionPool: newPAPool(),
	}
}

// Returns a random float.
//
// In tests, although we can set the initial seed, test results are still very
// sensitive to the exact order of RandomFloat() calls. To mitigate this, when
// testing we use a separate rand object for each RandomFloat callsite,
// distinguished by the label string.
func (sim *Simulation) RandomFloat(label string) float64 {
	if !sim.isTest {
		return sim.rand.Float64()
	}

	labelRand, isPresent := sim.testRands[label]
	if !isPresent {
		labelRand = rand.New(rand.NewSource(int64(hash(label))))
		sim.testRands[label] = labelRand
	}
	v := labelRand.Float64()
	// if sim.Log != nil {
	// 	sim.Log("FLOAT64 '%s': %0.5f", label, v)
	// }
	return v
}

func (sim *Simulation) Reset() {
	sim.reset()
}

// Reset will set sim back and erase all current state.
// This is automatically called before every 'Run'.
func (sim *Simulation) reset() {
	if sim.Log != nil {
		sim.Log("SIM RESET")
		sim.Log("----------------------")
	}
	variation := sim.DurationVariation * 2

	sim.Duration = sim.BaseDuration + time.Duration((sim.RandomFloat("sim duration") * float64(variation))) - sim.DurationVariation
	sim.CurrentTime = 0.0

	sim.pendingActions = make([]*PendingAction, 0, 64)

	// Targets need to be reset before the raid, so that players can check for
	// the presence of permanent target auras in their Reset handlers.
	for _, target := range sim.encounter.Targets {
		target.Reset(sim)
	}

	sim.Raid.reset(sim)

	sim.initManaTickAction()
}

// Run runs the simulation for the configured number of iterations, and
// collects all the metrics together.
func (sim *Simulation) run() *proto.RaidSimResult {
	t := time.Now()

	logsBuffer := &strings.Builder{}
	if sim.Options.Debug || sim.Options.DebugFirstIteration {
		sim.Log = func(message string, vals ...interface{}) {
			logsBuffer.WriteString(fmt.Sprintf("[%0.2f] "+message+"\n", append([]interface{}{sim.CurrentTime.Seconds()}, vals...)...))
		}
	}

	// Uncomment this to print logs directly to console.
	// sim.Log = func(message string, vals ...interface{}) {
	// 	fmt.Printf(fmt.Sprintf("[%0.1f] "+message+"\n", append([]interface{}{sim.CurrentTime.Seconds()}, vals...)...))
	// }

	for _, party := range sim.Raid.Parties {
		for _, player := range party.Players {
			character := player.GetCharacter()
			character.auraTracker.logFn = func(message string, vals ...interface{}) {
				character.Log(sim, message, vals...)
			}
			player.Init(sim)

			for _, petAgent := range character.Pets {
				petCharacter := petAgent.GetCharacter()
				petCharacter.auraTracker.logFn = func(message string, vals ...interface{}) {
					petCharacter.Log(sim, message, vals...)
				}
				petAgent.Init(sim)
			}
		}
	}

	for _, target := range sim.encounter.Targets {
		target.auraTracker.logFn = func(message string, vals ...interface{}) {
			target.Log(sim, message, vals...)
		}
	}

	sim.runOnce()
	firstIterationDuration := sim.Duration

	if !sim.Options.Debug {
		sim.Log = nil
	}

	st := time.Now()
	for i := int32(1); i < sim.Options.Iterations; i++ {
		// fmt.Printf("Iteration: %d\n", i)
		if sim.ProgressReport != nil && time.Since(st) > time.Millisecond*100 {
			metrics := sim.Raid.GetMetrics(i + 1)
			sim.ProgressReport(&proto.ProgressMetrics{TotalIterations: sim.Options.Iterations, CompletedIterations: i + 1, Dps: metrics.Dps.Avg})
			runtime.Gosched() // ensure that reporting threads are given time to report, mostly only important in wasm (only 1 thread)
			st = time.Now()
		}
		sim.runOnce()
	}
	result := &proto.RaidSimResult{
		RaidMetrics:      sim.Raid.GetMetrics(sim.Options.Iterations),
		EncounterMetrics: sim.encounter.GetMetricsProto(sim.Options.Iterations),

		Logs:                   logsBuffer.String(),
		FirstIterationDuration: firstIterationDuration.Seconds(),
	}

	// Final progress report
	if sim.ProgressReport != nil {
		sim.ProgressReport(&proto.ProgressMetrics{TotalIterations: sim.Options.Iterations, CompletedIterations: sim.Options.Iterations, Dps: result.RaidMetrics.Dps.Avg, FinalRaidResult: result})
	}

	if sim.Options.Iterations > 100 {
		log.Printf("running %d iterations took %d ms", sim.Options.Iterations, time.Since(t).Milliseconds())
	}

	return result
}

// RunOnce is the main event loop. It will run the simulation for number of seconds.
func (sim *Simulation) runOnce() {
	sim.reset()

	for true {
		last := len(sim.pendingActions) - 1
		pa := sim.pendingActions[last]
		sim.pendingActions = sim.pendingActions[:last]
		if pa.cancelled {
			sim.pendingActionPool.Put(pa)
			continue
		}

		if pa.NextActionAt > sim.Duration {
			if pa.CleanUp != nil {
				pa.CleanUp(sim)
			}
			break
		}

		if pa.NextActionAt > sim.CurrentTime {
			sim.advance(pa.NextActionAt - sim.CurrentTime)
		}

		pa.OnAction(sim)
	}

	for _, pa := range sim.pendingActions {
		if pa == nil {
			continue
		}
		if pa.CleanUp != nil {
			pa.CleanUp(sim)
		}
	}

	sim.Raid.doneIteration(sim.Duration)
	sim.encounter.doneIteration(sim.Duration)
}

func (sim *Simulation) AddPendingAction(pa *PendingAction) {
	oldlen := len(sim.pendingActions)

	// The logic to calculate the index to insert at can be replaced with sort.Search() which uses a binary search.
	//   However I haven't found any cases yet in our simulator that it is faster.
	var index = 0
	for _, v := range sim.pendingActions {
		if v.NextActionAt < pa.NextActionAt || (v.NextActionAt == pa.NextActionAt && v.Priority >= pa.Priority) {
			break
		}
		index++
	}

	sim.pendingActions = append(sim.pendingActions, pa)
	if index == oldlen { // if the insert was at the end, just return now.
		return
	} else if oldlen == 1 { // simple case we can just swap the two
		sim.pendingActions[0], sim.pendingActions[1] = sim.pendingActions[1], sim.pendingActions[0]
		return
	}

	copy(sim.pendingActions[index+1:], sim.pendingActions[index:])
	sim.pendingActions[index] = pa
}

// Advance moves time forward counting down auras, CDs, mana regen, etc
func (sim *Simulation) advance(elapsedTime time.Duration) {
	sim.CurrentTime += elapsedTime

	for _, party := range sim.Raid.Parties {
		for _, agent := range party.Players {
			agent.GetCharacter().advance(sim, elapsedTime)
		}
	}

	for _, target := range sim.encounter.Targets {
		target.Advance(sim, elapsedTime)
	}
}

func (sim *Simulation) IsExecutePhase() bool {
	return sim.CurrentTime > sim.encounter.executePhaseBegins
}

func (sim *Simulation) GetRemainingDuration() time.Duration {
	return sim.Duration - sim.CurrentTime
}

// Returns the percentage of time remaining in the current iteration, as a value from 0-1.
func (sim *Simulation) GetRemainingDurationPercent() float64 {
	return float64(sim.Duration-sim.CurrentTime) / float64(sim.Duration)
}

// The maximum possible duration for any iteration.
func (sim *Simulation) GetMaxDuration() time.Duration {
	return sim.BaseDuration + sim.DurationVariation
}

func (sim *Simulation) GetNumTargets() int32 {
	return int32(len(sim.encounter.Targets))
}

func (sim *Simulation) GetTarget(index int32) *Target {
	return sim.encounter.Targets[index]
}

func (sim *Simulation) GetPrimaryTarget() *Target {
	return sim.GetTarget(0)
}
