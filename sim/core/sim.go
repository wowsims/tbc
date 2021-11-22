package core

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
)

func debugFunc(sim *Simulation) func(string, ...interface{}) {
	return func(s string, vals ...interface{}) {
		fmt.Printf("[%0.1f] "+s, append([]interface{}{sim.CurrentTime.Seconds()}, vals...)...)
	}
}

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
	pendingActions []*PendingAction
	CurrentTime    time.Duration // duration that has elapsed in the sim since starting

	ProgressReport func(*proto.ProgressMetrics)

	Log  func(string, ...interface{})
	logs []string
}

func NewIndividualSim(isr proto.IndividualSimRequest) *Simulation {
	raid := NewRaid(*isr.RaidBuffs, *isr.PartyBuffs, *isr.IndividualBuffs)
	raid.AddPlayer(NewAgent(*isr.Player, isr))
	raid.Finalize()

	encounter := NewEncounter(*isr.Encounter)
	encounter.Finalize()

	return newSim(raid, encounter, *isr.SimOptions)
}

// New sim contructs a simulator with the given raid and target settings.
func newSim(raid *Raid, encounter Encounter, simOptions proto.SimOptions) *Simulation {
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
		Duration:  DurationFromSeconds(encounter.Duration),
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
	// 	sim.Log("FLOAT64 FROM: %s\n", label)
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
		sim.Log("SIM RESET\n")
		sim.Log("----------------------\n")
	}

	// Reset all players
	for _, party := range sim.Raid.Parties {
		for _, agent := range party.Players {
			agent.GetCharacter().reset(sim)
			agent.Reset(sim)
		}
	}

	for _, target := range sim.encounter.Targets {
		target.Reset(sim)
	}

	sim.CurrentTime = 0.0
}

type PendingAction struct {
	OnAction     func(*Simulation)
	CleanUp      func(*Simulation)
	NextActionAt time.Duration
}

// Run runs the simulation for the configured number of iterations, and
// collects all the metrics together.
func (sim *Simulation) Run() *proto.RaidSimResult {
	pid := 0
	for _, raidParty := range sim.Raid.Parties {
		for _, player := range raidParty.Players {
			player.GetCharacter().ID = pid
			player.GetCharacter().auraTracker.playerID = pid
			pid++
			player.Init(sim)
		}
	}
	logsBuffer := &strings.Builder{}

	if sim.Options.Debug {
		sim.Log = func(s string, vals ...interface{}) {
			logsBuffer.WriteString(fmt.Sprintf("[%0.1f] "+s, append([]interface{}{sim.CurrentTime.Seconds()}, vals...)...))
		}
	}

	simDurationSeconds := sim.Duration.Seconds()
	st := time.Now()
	for i := int32(0); i < sim.Options.Iterations; i++ {
		sim.RunOnce()

		sim.Raid.doneIteration(simDurationSeconds)
		if sim.ProgressReport != nil && time.Since(st) > time.Millisecond*250 {
			sim.ProgressReport(&proto.ProgressMetrics{TotalIterations: sim.Options.Iterations, CompletedIterations: i, Dps: sim.Raid.Parties[0].Players[0].GetCharacter().Metrics.dpsSum / float64(i)})
			st = time.Now()
		}
	}

	// Reset after the last iteration, because some metrics get updated in reset().
	sim.reset()

	result := &proto.RaidSimResult{
		RaidMetrics:      sim.Raid.GetMetrics(sim.Options.Iterations),
		EncounterMetrics: sim.encounter.GetMetricsProto(),

		Logs: logsBuffer.String(),
	}
	return result
}

// Runs a full sim for an individual player.
func (sim *Simulation) RunIndividual() *proto.IndividualSimResult {
	raidResult := sim.Run()
	return &proto.IndividualSimResult{
		PlayerMetrics:    raidResult.RaidMetrics.Parties[0].Players[0],
		EncounterMetrics: raidResult.EncounterMetrics,
		Logs:             raidResult.Logs,
	}
}

// RunOnce is the main event loop. It will run the simulation for number of seconds.
func (sim *Simulation) RunOnce() {
	sim.reset()

	sim.pendingActions = make([]*PendingAction, 0, 25)
	// setup initial actions.
	for _, party := range sim.Raid.Parties {
		for _, agent := range party.Players {
			ag := agent
			pa := &PendingAction{}
			pa.OnAction = func(sim *Simulation) {
				ag.GetCharacter().TryUseCooldowns(sim)
				pa.NextActionAt = ag.Act(sim)
			}
			sim.AddPendingAction(pa)
		}
	}

	// order pending by execution time.
	sort.Slice(sim.pendingActions, func(i, j int) bool {
		return sim.pendingActions[i].NextActionAt < sim.pendingActions[j].NextActionAt
	})

	for true {
		pa := sim.pendingActions[0]
		if pa.NextActionAt > sim.Duration {
			break
		}

		if pa.NextActionAt > sim.CurrentTime {
			sim.advance(pa.NextActionAt - sim.CurrentTime)
		}

		pa.OnAction(sim)

		if len(sim.pendingActions) == 1 {
			// We know in a single user sim, just always make the next pending action ours.
			sim.pendingActions[0] = pa
		} else {
			// This path is only used when there is more than one
			//  action sitting on the list.
			// This path is not currently used by individual shaman sim.
			if pa.NextActionAt == NeverExpires {
				sim.pendingActions = sim.pendingActions[1:] // cut off front
			}
			sort.Slice(sim.pendingActions, func(i, j int) bool {
				return sim.pendingActions[i].NextActionAt < sim.pendingActions[j].NextActionAt
			})
		}
	}

	for _, pa := range sim.pendingActions {
		if pa.CleanUp != nil {
			pa.CleanUp(sim)
		}
	}
}

func (sim *Simulation) AddPendingAction(pa *PendingAction) {
	sim.pendingActions = append(sim.pendingActions, pa)
}

// TODO: remove pending actions
func (sim *Simulation) RemovePendingAction(id int32) {

}

// Advance moves time forward counting down auras, CDs, mana regen, etc
func (sim *Simulation) advance(elapsedTime time.Duration) {
	sim.CurrentTime += elapsedTime

	for _, party := range sim.Raid.Parties {
		for _, agent := range party.Players {
			agent.Advance(sim, elapsedTime)
			agent.GetCharacter().advance(sim, elapsedTime)
		}
	}

	for _, target := range sim.encounter.Targets {
		target.Advance(sim, elapsedTime)
	}
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
