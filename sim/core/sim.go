package core

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func debugFunc(sim *Simulation) func(string, ...interface{}) {
	return func(s string, vals ...interface{}) {
		fmt.Printf("[%0.1f] "+s, append([]interface{}{sim.CurrentTime.Seconds()}, vals...)...)
	}
}

type Options struct {
	Encounter  Encounter
	Iterations int
	RSeed      int64
	ExitOnOOM  bool
	Debug      bool          // enables debug printing.
}

type IndividualParams struct {
	Equip    items.EquipmentSpec
	Race     proto.Race
	Class    proto.Class
	Consumes proto.Consumes
	Buffs    proto.Buffs
	Options  Options

	PlayerOptions *proto.PlayerOptions

	CustomStats stats.Stats
}

type InitialAura func(*Simulation) Aura

type Simulation struct {
	Raid         *Raid
	targets      []*Target
	Options      Options
	Duration     time.Duration

	MetricsAggregator *MetricsAggregator

	Rando       *wrappedRandom
	rseed       int64
	CurrentTime time.Duration // duration that has elapsed in the sim since starting

	Log  func(string, ...interface{})
	logs []string
}

type wrappedRandom struct {
	sim *Simulation
	*rand.Rand
}

func (wr *wrappedRandom) Float64(src string) float64 {
	// if wr.sim.Log != nil {
	// 	wr.sim.Log("FLOAT64 FROM: %s\n", src)
	// }
	return wr.Rand.Float64()
}

func NewIndividualSim(params IndividualParams) *Simulation {
	raid := NewRaid(params.Buffs)
	raid.AddPlayer(NewAgent(params))
	raid.Finalize()

	for _, target := range params.Options.Encounter.Targets {
		target.Finalize()
	}

	return newSim(raid, params.Options)
}

// New sim contructs a simulator with the given raid and target settings.
func newSim(raid *Raid, options Options) *Simulation {
	if options.RSeed == 0 {
		options.RSeed = time.Now().Unix()
	}

	if len(options.Encounter.Targets) == 0 {
		panic("Must have at least 1 target!")
	}

	sim := &Simulation{
		Raid:         raid,
		targets:      options.Encounter.Targets,
		Options:      options,
		Duration:     DurationFromSeconds(options.Encounter.Duration),
		Log: nil,
		MetricsAggregator: NewMetricsAggregator(raid.Size(), options.Encounter.Duration),
	}
	sim.Rando = &wrappedRandom{sim: sim, Rand: rand.New(rand.NewSource(options.RSeed))}

	return sim
}

// Get the metrics for an invidual Agent, for the current iteration.
func (sim *Simulation) GetIndividualMetrics(agentID int) AgentIterationMetrics {
	return sim.MetricsAggregator.agentIterations[agentID]
}

// Reset will set sim back and erase all current state.
// This is automatically called before every 'Run'.
func (sim *Simulation) reset() {
	sim.CurrentTime = 0.0

	if sim.Log != nil {
		sim.Log("SIM RESET\n")
		sim.Log("----------------------\n")
	}

	// Reset all players
	for _, party := range sim.Raid.Parties {
		for _, agent := range party.Players {
			agent.GetCharacter().Reset(sim)
			agent.Reset(sim)
		}
	}

	for _, target := range sim.targets {
		target.Reset(sim)
	}
}

type pendingAgent struct {
	Agent Agent
	NextActionAt time.Duration
}

// Run runs the simulation for the configured number of iterations, and
// collects all the metrics together.
func (sim *Simulation) Run() SimResult {
	pid := 0
	for _, raidParty := range sim.Raid.Parties {
		for _, player := range raidParty.Players {
			player.GetCharacter().ID = pid
			player.GetCharacter().auraTracker.playerID = pid
			pid++
		}
	}
	logsBuffer := &strings.Builder{}

	if sim.Options.Debug {
		sim.Log = func(s string, vals ...interface{}) {
			logsBuffer.WriteString(fmt.Sprintf("[%0.1f] "+s, append([]interface{}{sim.CurrentTime.Seconds()}, vals...)...))
		}
	}

	for i := 0; i < sim.Options.Iterations; i++ {
		sim.RunOnce()
		sim.MetricsAggregator.doneIteration()
	}

	result := sim.MetricsAggregator.getResult()
	result.Logs = logsBuffer.String()
	return result
}

// RunOnce is the main event loop. It will run the simulation for number of seconds.
func (sim *Simulation) RunOnce() {
	sim.reset()

	pendingAgents := make([]pendingAgent, 0, 25)
	// setup initial actions.
	for _, party := range sim.Raid.Parties {
		for _, agent := range party.Players {
			pendingAgents = append(pendingAgents, pendingAgent{
				NextActionAt: 0,
				Agent:        agent,
			})
		}
	}
	// order pending by execution time.
	sort.Slice(pendingAgents, func(i, j int) bool {
		return pendingAgents[i].NextActionAt < pendingAgents[j].NextActionAt
	})

	for sim.CurrentTime < sim.Duration {
		pa := pendingAgents[0]

		if pa.NextActionAt > sim.CurrentTime {
			sim.Advance(pa.NextActionAt - sim.CurrentTime)
		}

		pa.Agent.GetCharacter().TryUseCooldowns(sim)
		pa.NextActionAt = pa.Agent.Act(sim)

		if len(pendingAgents) == 1 {
			// We know in a single user sim, just always make the next pending action ours.
			pendingAgents[0] = pa
		} else {
			// Insert into the list the correct execution time.
			for i, v := range pendingAgents {
				if v.NextActionAt >= pa.NextActionAt {
					copy(pendingAgents, pendingAgents[:i])
					pendingAgents[i] = pa
				}
			}
		}
	}
}

// Advance moves time forward counting down auras, CDs, mana regen, etc
func (sim *Simulation) Advance(elapsedTime time.Duration) {
	newTime := sim.CurrentTime + elapsedTime
	sim.CurrentTime = newTime

	for _, party := range sim.Raid.Parties {
		for _, agent := range party.Players {
			// FUTURE: Do agents need to be notified or just advance the player state?
			agent.GetCharacter().Advance(sim, elapsedTime, newTime)
		}
	}

	for _, target := range sim.targets {
		target.auraTracker.advance(sim, elapsedTime)
	}
}

func (sim *Simulation) GetNumTargets() int32 {
	return int32(len(sim.targets))
}

func (sim *Simulation) GetTarget(index int32) *Target {
	return sim.targets[index]
}

func (sim *Simulation) GetPrimaryTarget() *Target {
	return sim.GetTarget(0)
}
