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

	Log  func(string, ...interface{})
	logs []string
}

func RunSim(rsr proto.RaidSimRequest) *proto.RaidSimResult {
	sim := newSim(rsr)
	sim.runPresims(rsr)
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
	Name         string
	OnAction     func(*Simulation)
	CleanUp      func(*Simulation)
	NextActionAt time.Duration
}

// Run runs the simulation for the configured number of iterations, and
// collects all the metrics together.
func (sim *Simulation) run() *proto.RaidSimResult {
	for _, party := range sim.Raid.Parties {
		for _, player := range party.Players {
			player.Init(sim)
		}
	}

	logsBuffer := &strings.Builder{}
	if sim.Options.Debug {
		sim.Log = func(s string, vals ...interface{}) {
			logsBuffer.WriteString(fmt.Sprintf("[%0.1f] "+s, append([]interface{}{sim.CurrentTime.Seconds()}, vals...)...))
		}
	}

	for i := int32(0); i < sim.Options.Iterations; i++ {
		sim.runOnce()
	}

	result := &proto.RaidSimResult{
		RaidMetrics:      sim.Raid.GetMetrics(sim.Options.Iterations),
		EncounterMetrics: sim.encounter.GetMetricsProto(sim.Options.Iterations),

		Logs: logsBuffer.String(),
	}
	return result
}

// RunOnce is the main event loop. It will run the simulation for number of seconds.
func (sim *Simulation) runOnce() {
	sim.reset()

	sim.pendingActions = make([]*PendingAction, 0, 25)
	// setup initial actions.
	for _, party := range sim.Raid.Parties {
		for _, agent := range party.Players {
			ag := agent
			pa := &PendingAction{
				Name: "Agent",
			}
			pa.OnAction = func(sim *Simulation) {
				ag.GetCharacter().TryUseCooldowns(sim)
				dur := ag.Act(sim)
				if dur == 0 {
					panic("Agent returned a 0 time wait")
				}
				pa.NextActionAt = dur
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
			} else {
				handled := false
				for i, v := range sim.pendingActions {
					if i == 0 {
						continue
					}
					if v.NextActionAt >= pa.NextActionAt {
						handled = true
						if i == 1 {
							sim.pendingActions[0] = pa
							break // just leave it there
						}
						copy(sim.pendingActions, sim.pendingActions[1:i])
						sim.pendingActions[i-1] = pa
						break
					}
				}
				if !handled {
					copy(sim.pendingActions, sim.pendingActions[1:])
					sim.pendingActions[len(sim.pendingActions)-1] = pa
				}
			}
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
	handled := false
	for i, v := range sim.pendingActions {
		if v.NextActionAt >= pa.NextActionAt {
			handled = true
			sim.pendingActions = append(sim.pendingActions, &PendingAction{})
			copy(sim.pendingActions[i+1:], sim.pendingActions[i:])
			sim.pendingActions[i] = pa
			break
		}
	}
	if !handled {
		sim.pendingActions = append(sim.pendingActions, pa)
	}
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
