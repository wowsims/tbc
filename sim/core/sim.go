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
	GCDMin     time.Duration // sets the minimum GCD
	Debug      bool          // enables debug printing.
}

type Encounter struct {
	Duration   float64
	NumTargets int
	Armor      int32
}

type IndividualParams struct {
	Equip    items.EquipmentSpec
	Race     RaceBonusType
	Consumes Consumes
	Buffs    Buffs
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

	// Clears and regenerates on each Run call.
	Metrics SimMetrics

	Rando       *wrappedRandom
	rseed       int64
	CurrentTime time.Duration // duration that has elapsed in the sim since starting

	Log  func(string, ...interface{})
	logs []string

	// caches to speed up perf and store temp state
	cache *cache

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

type SimMetrics struct {
	TotalDamage       float64
	Casts             []*Cast
	IndividualMetrics []IndividualMetric
}

type IndividualMetric struct {
	ID          int32
	TotalDamage float64
	DamageAtOOM float64
	OOMAt       float64
	ManaSpent   float64
}

func NewIndividualSim(params IndividualParams) *Simulation {
	raid := NewRaid(params.Buffs)
	sim := newSim(raid, params.Options)
	sim.IndividualParams = params

	character := NewCharacter(params.Equip, params.Race, params.Consumes, params.CustomStats)
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
func newSim(raid *Raid, options Options) *Simulation {
	if options.GCDMin == 0 {
		options.GCDMin = durationFromSeconds(1.0) // default to 0.75s GCD
	}
	if options.RSeed == 0 {
		options.RSeed = time.Now().Unix()
	}

	sim := &Simulation{
		Raid:         raid,
		Options:      options,
		Duration:     durationFromSeconds(options.Encounter.Duration),
		InitialAuras: []InitialAura{},
		// Rando:    ,
		Log: nil,
		cache: &cache{
			castPool: make([]*Cast, 0, 1000),
		},
		AuraTracker: NewAuraTracker(),
	}
	sim.Rando = &wrappedRandom{sim: sim, Rand: rand.New(rand.NewSource(options.RSeed))}

	sim.cache.fillCasts()

	return sim
}

func (sim *Simulation) NewCast() *Cast {
	return sim.cache.NewCast()
}
func (sim *Simulation) ReturnCasts(casts []*Cast) {
	sim.cache.ReturnCasts(casts)
}

func (sim *Simulation) AddInitialAura(initialAura InitialAura) {
	sim.InitialAuras = append(sim.InitialAuras, initialAura)
}

// Reset will set sim back and erase all current state.
// This is automatically called before every 'Run'
//  This includes resetting and reactivating always on trinkets, auras, set bonuses, etc
func (sim *Simulation) Reset() {
	// sim.rseed++
	// sim.Rando.Seed(sim.rseed)

	sim.CurrentTime = 0.0
	sim.ResetAuras()
	sim.Metrics = SimMetrics{
		Casts:             make([]*Cast, 0, 1000),
		IndividualMetrics: make([]IndividualMetric, 25),
	}
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
	TryActivateDestructionPotion(sim, agent)
	TryActivateDarkRune(sim, agent)
	TryActivateSuperManaPotion(sim, agent)

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
	aggregator := NewMetricsAggregator()

	if sim.Options.Debug {
		sim.Log = func(s string, vals ...interface{}) {
			logsBuffer.WriteString(fmt.Sprintf("[%0.1f] "+s, append([]interface{}{sim.CurrentTime.Seconds()}, vals...)...))
		}
	}

	for i := 0; i < sim.Options.Iterations; i++ {
		metrics := sim.RunOnce()
		aggregator.addMetrics(sim.Options, metrics)
		sim.ReturnCasts(metrics.Casts)
	}

	result := aggregator.getResult()
	result.Logs = logsBuffer.String()
	return result
}

// RunOnce will run the simulation for number of seconds.
// Returns metrics for what was cast and how much damage was done.
func (sim *Simulation) RunOnce() SimMetrics {
	sim.Reset()

	pendingActions := make([]pendingAction, 0, 25)
	// setup initial actions.
	for _, party := range sim.Raid.Parties {
		for _, agent := range party.Players {
			sim.playerConsumes(agent)
			action := agent.ChooseAction(sim)
			if action.Wait == NeverExpires {
				continue // This means agent will not perform any actions at all
			}
			wait := action.Wait
			if action.Cast != nil {
				wait = action.Cast.CastTime
			}
			pendingActions = append(pendingActions, pendingAction{
				ExecuteAt:   wait,
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

		if action.Cast != nil {
			action.Cast.DoItNow(sim, action.Cast)
		} else if action.Wait == 0 {
			// FUTURE: Swing timers could be handled in this if block.
			panic("Agent returned nil action")
		}

		sim.playerConsumes(agent)
		newAction := agent.ChooseAction(sim)
		wait := newAction.Wait
		if newAction.Cast != nil {
			if newAction.Cast.CastTime < sim.Options.GCDMin {
				newAction.Cast.CastTime = sim.Options.GCDMin
			}
			wait = newAction.Cast.CastTime
			if agent.GetCharacter().Stats[stats.Mana] < newAction.Cast.ManaCost {
				// Not enough mana, wait until there is enough mana to cast the desired spell
				regenTime := durationFromSeconds((newAction.Cast.ManaCost-agent.GetCharacter().Stats[stats.Mana])/agent.GetCharacter().manaRegenPerSecond()) + 1
				if sim.Log != nil {
					sim.Log("Not enough mana to cast... regen for %0.1f seconds before casting.\n", regenTime.Seconds())
				}
				wait = regenTime
				if sim.Options.ExitOnOOM {
					break simloop // named for clarity since this is pretty deep nested.
				}

				// Cancel cast for now.
				newAction.Cast = nil
				newAction.Wait = wait
				sim.Metrics.IndividualMetrics[agent.GetCharacter().ID].DamageAtOOM = sim.Metrics.IndividualMetrics[agent.GetCharacter().ID].TotalDamage
				sim.Metrics.IndividualMetrics[agent.GetCharacter().ID].OOMAt = sim.CurrentTime.Seconds()
			} else {
				if sim.Log != nil {
					sim.Log("(%d) Start Casting %s Cast Time: %0.1fs\n", agent.GetCharacter().ID, newAction.Cast.Spell.Name, newAction.Cast.CastTime.Seconds())
				}
			}
		}
		pa := pendingAction{
			ExecuteAt:   sim.CurrentTime + wait,
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

	return sim.Metrics
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

func durationFromSeconds(numSeconds float64) time.Duration {
	return time.Duration(float64(time.Second) * numSeconds)
}
