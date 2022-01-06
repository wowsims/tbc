package core

import (
	"container/heap"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
)

type InitialAura func(*Simulation) Aura

type Simulation struct {
	Raid      *Raid
	encounter Encounter
	Options   proto.SimOptions
	Duration  time.Duration

	rand *rand.Rand

	// Used for testing only, see RandomFloat().
	isTest    bool
	testRands map[uint32]*rand.Rand

	// Current Simulation State
	pendingActions ActionsQueue
	CurrentTime    time.Duration // duration that has elapsed in the sim since starting

	ProgressReport func(*proto.ProgressMetrics)

	Log  func(string, ...interface{})
	logs []string
}

func RunSim(rsr proto.RaidSimRequest, progress chan *proto.ProgressMetrics) *proto.RaidSimResult {
	sim := newSim(rsr)
	sim.runPresims(rsr)
	if progress != nil {
		sim.ProgressReport = func(progMetric *proto.ProgressMetrics) {
			progress <- progMetric
		}
	}
	return sim.run()
}

func newSim(rsr proto.RaidSimRequest) *Simulation {
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
		Raid:      raid,
		encounter: encounter,
		Options:   simOptions,
		Duration:  encounter.Duration,
		Log:       nil,

		rand: rand.New(rand.NewSource(rseed)),

		isTest:    simOptions.IsTest,
		testRands: make(map[uint32]*rand.Rand),
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

	// if sim.Log != nil {
	// 	sim.Log("FLOAT64 FROM: %s", label)
	// }

	labelHash := hash(label)
	labelRand, isPresent := sim.testRands[labelHash]
	if !isPresent {
		labelRand = rand.New(rand.NewSource(int64(labelHash)))
		sim.testRands[labelHash] = labelRand
	}
	return labelRand.Float64()
}

// Reset will set sim back and erase all current state.
// This is automatically called before every 'Run'.
func (sim *Simulation) reset() {
	if sim.Log != nil {
		sim.Log("SIM RESET")
		sim.Log("----------------------")
	}

	sim.CurrentTime = 0.0
	sim.pendingActions = make([]*PendingAction, 0, 64)

	sim.Raid.reset(sim)

	for _, target := range sim.encounter.Targets {
		target.Reset(sim)
	}
}

// Run runs the simulation for the configured number of iterations, and
// collects all the metrics together.
func (sim *Simulation) run() *proto.RaidSimResult {
	logsBuffer := &strings.Builder{}
	if sim.Options.Debug || sim.Options.DebugFirstIteration {
		sim.Log = func(message string, vals ...interface{}) {
			logsBuffer.WriteString(fmt.Sprintf("[%0.1f] "+message+"\n", append([]interface{}{sim.CurrentTime.Seconds()}, vals...)...))
		}
	}

	// Uncomment this to print logs directly to console.
	//sim.Log = func(message string, vals ...interface{}) {
	//	fmt.Printf(fmt.Sprintf("[%0.1f] "+message+"\n", append([]interface{}{sim.CurrentTime.Seconds()}, vals...)...))
	//}

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

	if !sim.Options.Debug {
		sim.Log = nil
	}

	st := time.Now()
	for i := int32(1); i < sim.Options.Iterations; i++ {
		if sim.ProgressReport != nil && time.Since(st) > time.Millisecond*250 {
			metrics := sim.Raid.GetMetrics(i + 1)
			sim.ProgressReport(&proto.ProgressMetrics{TotalIterations: sim.Options.Iterations, CompletedIterations: i + 1, Dps: metrics.Dps.Avg})
			st = time.Now()
		}
		sim.runOnce()
	}
	result := &proto.RaidSimResult{
		RaidMetrics:      sim.Raid.GetMetrics(sim.Options.Iterations),
		EncounterMetrics: sim.encounter.GetMetricsProto(sim.Options.Iterations),

		Logs: logsBuffer.String(),
	}

	// Final progress report
	if sim.ProgressReport != nil {
		sim.ProgressReport(&proto.ProgressMetrics{TotalIterations: sim.Options.Iterations, CompletedIterations: sim.Options.Iterations, Dps: result.RaidMetrics.Dps.Avg, FinalResult: result})
	}

	return result
}

// RunOnce is the main event loop. It will run the simulation for number of seconds.
func (sim *Simulation) runOnce() {
	sim.reset()

	for true {
		pa := heap.Pop(&sim.pendingActions).(*PendingAction)
		if pa.NextActionAt > sim.Duration {
			if pa.CleanUp != nil {
				pa.CleanUp(sim)
			}
			break
		}

		if pa.NextActionAt > sim.CurrentTime {
			sim.advance(pa.NextActionAt - sim.CurrentTime)
		}

		if !pa.cancelled {
			pa.OnAction(sim)
		}
	}

	for _, pa := range sim.pendingActions {
		if pa.CleanUp != nil {
			pa.CleanUp(sim)
		}
	}

	sim.Raid.doneIteration(sim.Duration)
	sim.encounter.doneIteration(sim.Duration)
}

func (sim *Simulation) AddPendingAction(pa *PendingAction) {
	heap.Push(&sim.pendingActions, pa)
}

// Creates a PendingAction which repeatedly asks an Agent for actions.
func (sim *Simulation) newDefaultAgentAction(agent Agent) *PendingAction {
	pa := &PendingAction{
		Name:     "Agent",
		Priority: -1, // Give lower priority so that dot ticks always happen before player actions.
	}
	pa.OnAction = func(sim *Simulation) {
		// If char has AA enabled (a MH weapon is set), try to swing
		if agent.GetCharacter().AutoAttacks.mh != nil {
			agent.GetCharacter().AutoAttacks.Swing(sim, sim.GetPrimaryTarget())
		}
		agent.GetCharacter().TryUseCooldowns(sim)
		dur := agent.Act(sim)
		if dur <= sim.CurrentTime {
			panic(fmt.Sprintf("Agent returned invalid time delta: %s (%s - %s)", sim.CurrentTime-dur, sim.CurrentTime, dur))
		}
		pa.NextActionAt = dur
		sim.AddPendingAction(pa)
	}
	return pa
}

// Advance moves time forward counting down auras, CDs, mana regen, etc
func (sim *Simulation) advance(elapsedTime time.Duration) {
	sim.CurrentTime += elapsedTime

	for _, party := range sim.Raid.Parties {
		for _, agent := range party.Players {
			// TODO: Switch the order here to match other things like this.
			agent.Advance(sim, elapsedTime)
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

func (sim *Simulation) GetNumTargets() int32 {
	return int32(len(sim.encounter.Targets))
}

func (sim *Simulation) GetTarget(index int32) *Target {
	return sim.encounter.Targets[index]
}

func (sim *Simulation) GetPrimaryTarget() *Target {
	return sim.GetTarget(0)
}

type PendingAction struct {
	Name         string
	Priority     int
	OnAction     func(*Simulation)
	CleanUp      func(*Simulation)
	NextActionAt time.Duration

	cancelled bool
}

func (pa *PendingAction) Cancel(sim *Simulation) {
	if pa.cancelled {
		return
	}

	if pa.CleanUp != nil {
		pa.CleanUp(sim)
		pa.CleanUp = nil
	}

	pa.cancelled = true
}

type ActionsQueue []*PendingAction

func (queue ActionsQueue) Len() int {
	return len(queue)
}
func (queue ActionsQueue) Less(i, j int) bool {
	return queue[i].NextActionAt < queue[j].NextActionAt ||
		(queue[i].NextActionAt == queue[j].NextActionAt && queue[i].Priority > queue[j].Priority)
}
func (queue ActionsQueue) Swap(i, j int) {
	queue[i], queue[j] = queue[j], queue[i]
}
func (queue *ActionsQueue) Push(newAction interface{}) {
	*queue = append(*queue, newAction.(*PendingAction))
}
func (queue *ActionsQueue) Pop() interface{} {
	old := *queue
	n := len(old)
	action := old[n-1]
	old[n-1] = nil
	*queue = old[0 : n-1]
	return action
}
