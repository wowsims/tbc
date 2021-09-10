package core

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func debugFunc(sim *Simulation) func(string, ...interface{}) {
	return func(s string, vals ...interface{}) {
		fmt.Printf("[%0.1f] "+s, append([]interface{}{sim.CurrentTime.Seconds()}, vals...)...)
	}
}

type Options struct {
	Encounter  Encounter
	Iterations int32
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

type Simulation struct {
	Raid         Raid
	Options      Options
	Duration     time.Duration
	*AuraTracker // Global Debuffs mostly.. put on the boss/target

	// Clears and regenerates on each Run call.
	Metrics SimMetrics

	Rando       *rand.Rand
	rseed       int64
	CurrentTime time.Duration // duration that has elapsed in the sim since starting

	Debug func(string, ...interface{})

	// caches to speed up perf and store temp state
	cache *cache
}

type SimMetrics struct {
	TotalDamage       float64
	Casts             []*Cast
	IndividualMetrics []IndividualMetric
}

type IndividualMetric struct {
	// TODO: ID the player somehow?
	TotalDamage float64
	DamageAtOOM float64
	OOMAt       float64
	ManaSpent   float64
}

// New sim contructs a simulator with the given equipment / options.
func NewSim(raid *Raid, options Options) *Simulation {
	if options.GCDMin == 0 {
		options.GCDMin = durationFromSeconds(0.75) // default to 0.75s GCD
	}
	if options.RSeed == 0 {
		options.RSeed = time.Now().Unix()
	}

	sim := &Simulation{
		Options:  options,
		Duration: durationFromSeconds(options.Encounter.Duration),
		Rando:    rand.New(rand.NewSource(options.RSeed)),
		Debug:    nil,
		cache: &cache{
			castPool: make([]*Cast, 0, 1000),
		},
		AuraTracker: NewAuraTracker(),
	}

	sim.cache.fillCasts()

	return sim
}

func (sim *Simulation) NewCast() *Cast {
	return sim.cache.NewCast()
}
func (sim *Simulation) ReturnCasts(casts []*Cast) {
	sim.cache.ReturnCasts(casts)
}

// reset will set sim back and erase all current state.
// This is automatically called before every 'Run'
//  This includes resetting and reactivating always on trinkets, auras, set bonuses, etc
func (sim *Simulation) reset() {
	// sim.rseed++
	// sim.Rando.Seed(sim.rseed)

	sim.CurrentTime = 0.0
	sim.ResetAuras()
	sim.Metrics = SimMetrics{
		Casts: make([]*Cast, 0, 1000),
	}
	if sim.Debug != nil {
		sim.Debug("SIM RESET\n")
		sim.Debug("----------------------\n")
	}

	// Reset all players
	for _, party := range sim.Raid.Parties {
		for _, pl := range party.Players {
			pl.Player.Reset()
			pl.Agent.Reset(sim)
		}
	}

	// Now buff everyone back up!
	for _, party := range sim.Raid.Parties {
		for _, pl := range party.Players {
			pl.ActivateSets(sim, party)
			pl.TryActivateEquipment(sim, party)
			pl.BuffUp(sim, party)
		}
	}

	// TODO: probably need to special case the kings buff here.
}

type pendingAction struct {
	Agent PlayerAgent
	AgentAction
	ExecuteAt time.Duration
}

// Run will run the simulation for number of seconds.
// Returns metrics for what was cast and how much damage was done.
func (sim *Simulation) Run() SimMetrics {
	sim.reset()

	pendingActions := make([]pendingAction, 0, 25)
	for _, party := range sim.Raid.Parties {
		for _, pl := range party.Players {
			action := pl.ChooseAction(sim)
			pendingActions = append(pendingActions, pendingAction{
				ExecuteAt:   action.Wait,
				Agent:       pl,
				AgentAction: action,
			})
		}
	}
	// order pending by execution time.
	sort.Slice(pendingActions, func(i, j int) bool {
		return pendingActions[i].ExecuteAt < pendingActions[j].ExecuteAt
	})

	for sim.CurrentTime < sim.Duration {
		action := pendingActions[0]
		agent := action.Agent
		agent.OnActionAccepted(sim, action.AgentAction)
		if action.ExecuteAt > sim.CurrentTime {
			sim.Advance(action.ExecuteAt - sim.CurrentTime)
		}
		if action.Cast != nil {
			sim.Cast(action.Agent.Player, action.Cast)
		} else {
			// FUTURE: Swing timers could be handled in this if block.
			panic("Agent returned nil action")
		}

		newAction := agent.ChooseAction(sim)
		wait := newAction.Wait
		if newAction.Cast != nil {
			wait = newAction.Cast.CastTime
			if agent.Stats[StatMana] < newAction.Cast.ManaCost {
				// Not enough mana, wait until there is enough mana to cast the desired spell
				regenTime := durationFromSeconds((newAction.Cast.ManaCost-agent.Stats[StatMana])/agent.manaRegenPerSecond()) + 1
				if regenTime > wait {
					wait = regenTime
				}

				if sim.Options.ExitOnOOM {
					// TODO: implement this... first player to OOM ends sim (probably only makes sense in an individual sim)
				}
			}
			if newAction.Cast.CastTime < sim.Options.GCDMin {
				newAction.Cast.CastTime = sim.Options.GCDMin
			}
			if sim.Debug != nil {
				sim.Debug("Start Casting %s Cast Time: %0.1fs\n", newAction.Cast.Spell.Name, newAction.Cast.CastTime.Seconds())
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

// Cast will actually cast and treat all casts as having no 'flight time'.
// This will activate any auras around casting, calculate hit/crit and add to sim metrics.
func (sim *Simulation) Cast(p *Player, cast *Cast) {
	if sim.Debug != nil {
		sim.Debug("Current Mana %0.0f, Cast Cost: %0.0f\n", p.Stats[StatMana], cast.ManaCost)
	}
	p.Stats[StatMana] -= cast.ManaCost
	// sim.Metrics.ManaSpent += cast.ManaCost

	for _, id := range p.ActiveAuraIDs {
		if p.Auras[id].OnCastComplete != nil {
			p.Auras[id].OnCastComplete(sim, p, cast)
		}
	}
	hit := 0.83 + p.Stats[StatSpellHit]/1260.0 + cast.Hit // 12.6 hit == 1% hit
	hit = math.Min(hit, 0.99)                             // can't get away from the 1% miss

	dbgCast := cast.Spell.Name
	if sim.Debug != nil {
		sim.Debug("Completed Cast (%s)\n", dbgCast)
	}
	if sim.Rando.Float64() < hit {
		dmg := (sim.Rando.Float64() * cast.Spell.DmgDiff) + cast.Spell.MinDmg + (p.Stats[StatSpellPower] * cast.Spell.Coeff)
		if cast.DidDmg != 0 { // use the pre-set dmg
			dmg = cast.DidDmg
		}
		cast.DidHit = true

		crit := (p.Stats[StatSpellCrit] / 2208.0) + cast.Crit // 22.08 crit == 1% crit
		if sim.Rando.Float64() < crit {
			cast.DidCrit = true
			dmg *= cast.CritBonus
			if sim.Debug != nil {
				dbgCast += " crit"
			}
		} else if sim.Debug != nil {
			dbgCast += " hit"
		}

		// Average Resistance (AR) = (Target's Resistance / (Caster's Level * 5)) * 0.75
		// P(x) = 50% - 250%*|x - AR| <- where X is %resisted
		// Using these stats:
		//    13.6% chance of
		//  FUTURE: handle boss resists for fights/classes that are actually impacted by that.
		resVal := sim.Rando.Float64()
		if resVal < 0.18 { // 13% chance for 25% resist, 4% for 50%, 1% for 75%
			if sim.Debug != nil {
				dbgCast += " (partial resist: "
			}
			if resVal < 0.01 {
				dmg *= .25
				if sim.Debug != nil {
					dbgCast += "75%)"
				}
			} else if resVal < 0.05 {
				dmg *= .5
				if sim.Debug != nil {
					dbgCast += "50%)"
				}
			} else {
				dmg *= .75
				if sim.Debug != nil {
					dbgCast += "25%)"
				}
			}
		}
		cast.DidDmg = dmg
		// Apply any effects specific to this cast.
		if cast.Effect != nil {
			cast.Effect(sim, p, cast)
		}
		// Apply any on spell hit effects.
		for _, id := range p.ActiveAuraIDs {
			if p.Auras[id].OnSpellHit != nil {
				p.Auras[id].OnSpellHit(sim, p, cast)
			}
		}
	} else {
		if sim.Debug != nil {
			dbgCast += " miss"
		}
		cast.DidDmg = 0
		cast.DidCrit = false
		cast.DidHit = false
		for _, id := range p.ActiveAuraIDs {
			if p.Auras[id].OnSpellMiss != nil {
				p.Auras[id].OnSpellMiss(sim, p, cast)
			}
		}
	}

	if cast.Spell.Cooldown > 0 {
		p.SetCD(cast.Spell.ID, cast.Spell.Cooldown)
	}

	if sim.Debug != nil {
		sim.Debug("%s: %0.0f\n", dbgCast, cast.DidDmg)
	}

	sim.Metrics.Casts = append(sim.Metrics.Casts, cast)
	sim.Metrics.TotalDamage += cast.DidDmg
}

// Advance moves time forward counting down auras, CDs, mana regen, etc
func (sim *Simulation) Advance(elapsedTime time.Duration) {
	newTime := sim.CurrentTime + elapsedTime

	for _, party := range sim.Raid.Parties {
		for _, pl := range party.Players {
			// FUTURE: Do agents need to be notified or just advance the player state?
			pl.Advance(sim, elapsedTime, newTime)
		}
	}
	// Go in reverse order so we can safely delete while looping
	for i := len(sim.ActiveAuraIDs) - 1; i >= 0; i-- {
		id := sim.ActiveAuraIDs[i]
		if sim.Auras[id].Expires != 0 && sim.Auras[id].Expires <= newTime {
			sim.RemoveAura(sim, nil, id) // auras on the sim have no player attached.
		}
	}
	sim.CurrentTime = newTime
}

func durationFromSeconds(numSeconds float64) time.Duration {
	return time.Duration(float64(time.Second) * numSeconds)
}
