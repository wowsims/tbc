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

type Encounter struct {
	Duration   float64
	NumTargets int32
	Armor      int32
}

type IndividualParams struct {
	Equip    items.EquipmentSpec
	Race     RaceBonusType
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
	Options      Options
	Duration     time.Duration
	*AuraTracker // Global Debuffs mostly.. put on the boss/target

	// Auras which are automatically applied on sim reset.
	InitialAuras []InitialAura

	metricsAggregator *MetricsAggregator

	Rando       *wrappedRandom
	rseed       int64
	CurrentTime time.Duration // duration that has elapsed in the sim since starting

	Log  func(string, ...interface{})
	logs []string

	// Holds the params used to create this sim, so similar sims can be run if needed.
	IndividualParams IndividualParams
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
	sim := newSim(raid, params.Options, 1)
	sim.IndividualParams = params

	character := NewCharacter(params.Equip, params.Race, params.Class, params.Consumes, params.CustomStats)
	agent := NewAgent(sim, character, params.PlayerOptions)
	raid.AddPlayer(agent)
	raid.AddPlayerBuffs()

	raid.ApplyBuffs(sim)

	sim.Reset()

	// Now apply all the 'final' stat improvements.
	// TODO: Figure out how to handle buffs that buff based on other buffs...
	//   for now this hardcoded buffing works...
	for _, raidParty := range sim.Raid.Parties {
		for _, player := range raidParty.Players {
			if raid.Buffs.BlessingOfKings {
				player.GetCharacter().InitialStats[stats.Stamina] *= 1.1
				player.GetCharacter().InitialStats[stats.Agility] *= 1.1
				player.GetCharacter().InitialStats[stats.Strength] *= 1.1
				player.GetCharacter().InitialStats[stats.Intellect] *= 1.1
				player.GetCharacter().InitialStats[stats.Spirit] *= 1.1
			}
			if raid.Buffs.DivineSpirit == proto.TristateEffect_TristateEffectImproved {
				player.GetCharacter().InitialStats[stats.SpellPower] += player.GetCharacter().InitialStats[stats.Spirit] * 0.1
			}
			// Add SpellCrit from Int and Mana from Int
			player.GetCharacter().InitialStats = player.GetCharacter().InitialStats.CalculatedTotal()
		}
	}

	// Reset again to make sure updated initial stats are propagated.
	sim.Reset()

	return sim
}

// New sim contructs a simulator with the given equipment / options.
func newSim(raid *Raid, options Options, numPlayers int) *Simulation {
	if options.RSeed == 0 {
		options.RSeed = time.Now().Unix()
	}

	sim := &Simulation{
		Raid:         raid,
		Options:      options,
		Duration:     DurationFromSeconds(options.Encounter.Duration),
		InitialAuras: []InitialAura{},
		Log: nil,
		AuraTracker: NewAuraTracker(),
		metricsAggregator: NewMetricsAggregator(numPlayers, options.Encounter.Duration),
	}
	sim.Rando = &wrappedRandom{sim: sim, Rand: rand.New(rand.NewSource(options.RSeed))}

	return sim
}

// Get the metrics for an invidual Agent, for the current iteration.
func (sim *Simulation) GetIndividualMetrics(agentID int) AgentIterationMetrics {
	return sim.metricsAggregator.agentIterations[agentID]
}

func (sim *Simulation) AddInitialAura(initialAura InitialAura) {
	sim.InitialAuras = append(sim.InitialAuras, initialAura)
}

// Reset will set sim back and erase all current state.
// This is automatically called before every 'Run'
//  This includes resetting and reactivating always on trinkets, auras, set bonuses, etc
func (sim *Simulation) Reset() {
	sim.CurrentTime = 0.0
	sim.ResetAuras()
	if sim.Log != nil {
		sim.Log("SIM RESET\n")
		sim.Log("----------------------\n")
	}

	// Reset all players
	for _, party := range sim.Raid.Parties {
		for _, agent := range party.Players {
			agent.GetCharacter().Reset()
			agent.Reset(sim)
		}
	}

	// Now buff everyone back up!
	for _, party := range sim.Raid.Parties {
		for _, agent := range party.Players {
			agent.BuffUp(sim) // for now do this first to match order of adding auras as original sim.
			agent.GetCharacter().BuffUp(sim, agent)
		}
	}

	for _, initialAura := range sim.InitialAuras {
		sim.AddAura(sim, initialAura(sim))
	}
}

type pendingAction struct {
	Agent Agent
	AgentAction
	ExecuteAt time.Duration
}

func (sim *Simulation) playerConsumes(agent Agent) {
	// Consumes before any casts
	TryActivateDrums(sim, agent)
	TryActivateRacial(sim, agent)
	TryActivatePotion(sim, agent)
	TryActivateDarkRune(sim, agent)

	// Pop activatable items if we can.
	agent.GetCharacter().TryActivateEquipment(sim, agent)
}

// Run runs the simulation for the configured number of iterations, and
// collects all the metrics together.
func (sim *Simulation) Run() SimResult {
	pid := 0
	for _, raidParty := range sim.Raid.Parties {
		for _, player := range raidParty.Players {
			player.GetCharacter().ID = pid
			player.GetCharacter().AuraTracker.PID = pid
			pid++
		}
	}
	sim.AuraTracker.PID = -1 // set main sim auras to be -1 character ID.
	logsBuffer := &strings.Builder{}

	if sim.Options.Debug {
		sim.Log = func(s string, vals ...interface{}) {
			logsBuffer.WriteString(fmt.Sprintf("[%0.1f] "+s, append([]interface{}{sim.CurrentTime.Seconds()}, vals...)...))
		}
	}

	for i := 0; i < sim.Options.Iterations; i++ {
		sim.RunOnce()
		sim.metricsAggregator.doneIteration()
	}

	result := sim.metricsAggregator.getResult()
	result.Logs = logsBuffer.String()
	return result
}

// RunOnce will run the simulation for number of seconds.
// Returns metrics for what was cast and how much damage was done.
func (sim *Simulation) RunOnce() {
	sim.Reset()

	pendingActions := make([]pendingAction, 0, 25)
	// setup initial actions.
	for _, party := range sim.Raid.Parties {
		for _, agent := range party.Players {
			sim.playerConsumes(agent)

			action := agent.ChooseAction(sim)

			// sim.CurrentTime is 0, so dont need to add it
			executeAt := action.GetDuration()

			if executeAt == NeverExpires {
				continue // This means agent will not perform any actions at all
			}

			pendingActions = append(pendingActions, pendingAction{
				ExecuteAt:   executeAt,
				Agent:       agent,
				AgentAction: action,
			})
		}
	}
	// order pending by execution time.
	sort.Slice(pendingActions, func(i, j int) bool {
		return pendingActions[i].ExecuteAt < pendingActions[j].ExecuteAt
	})

simloop:
	for sim.CurrentTime < sim.Duration {
		action := pendingActions[0]
		agent := action.Agent
		agent.OnActionAccepted(sim, action.AgentAction)
		if action.ExecuteAt > sim.CurrentTime {
			sim.Advance(action.ExecuteAt - sim.CurrentTime)
		}

		action.Act(sim)

		sim.playerConsumes(agent)
		newAction := agent.ChooseAction(sim)
		actionDuration := newAction.GetDuration()

		castAction, isCastAction := newAction.(DirectCastAction)
		if isCastAction {
			// TODO: This delays the cast damage until GCD is ready, even if the cast time is less than GCD.
			// How to handle this?
			actionDuration = MaxDuration(actionDuration, GCDMin)

			manaCost := castAction.GetManaCost()
			if agent.GetCharacter().Stats[stats.Mana] < manaCost {
				// Not enough mana, wait until there is enough mana to cast the desired spell
				// TODO: Doesn't account for spirit-based mana
				regenTime := DurationFromSeconds((manaCost-agent.GetCharacter().Stats[stats.Mana])/agent.GetCharacter().manaRegenPerSecond()) + 1
				if sim.Log != nil {
					sim.Log("Not enough mana to cast... regen for %0.1f seconds before casting.\n", regenTime.Seconds())
				}
				actionDuration = regenTime
				if sim.Options.ExitOnOOM {
					break simloop // named for clarity since this is pretty deep nested.
				}

				// Cancel cast for now.
				newAction = NewWaitAction(sim, agent, actionDuration)
				sim.metricsAggregator.markOOM(agent, sim.CurrentTime)
			}
		}
		pa := pendingAction{
			ExecuteAt:   sim.CurrentTime + actionDuration,
			Agent:       agent,
			AgentAction: newAction,
		}
		if len(pendingActions) == 1 {
			// We know in a single user sim, just always make the next pending action ours.
			pendingActions[0] = pa
		} else {
			// Insert into the list the correct execution time.
			for i, v := range pendingActions {
				if v.ExecuteAt >= pa.ExecuteAt {
					copy(pendingActions, pendingActions[:i])
					pendingActions[i] = pa
				}
			}
		}
	}
}

// Advance moves time forward counting down auras, CDs, mana regen, etc
func (sim *Simulation) Advance(elapsedTime time.Duration) {
	newTime := sim.CurrentTime + elapsedTime

	for _, party := range sim.Raid.Parties {
		for _, agent := range party.Players {
			// FUTURE: Do agents need to be notified or just advance the player state?
			agent.GetCharacter().Advance(sim, elapsedTime, newTime)
		}
	}
	sim.AuraTracker.Advance(sim, elapsedTime)
	sim.CurrentTime = newTime
}
