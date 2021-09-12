package runner

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/druid"
	"github.com/wowsims/tbc/sim/mage"
	"github.com/wowsims/tbc/sim/paladin"
	"github.com/wowsims/tbc/sim/priest"
)

func SetupSim(raid *core.Raid, buffs core.Buffs, options core.Options) *core.Simulation {
	sim := core.NewSim(raid, options)

	// These buffs are a one-time apply... no need to add the bots to the raid group.
	druid.NewBuffBot(sim, raid.Parties[0], buffs.GiftOfTheWild, buffs.Moonkin, buffs.MoonkinRavenGoddess)
	mage.NewBuffBot(sim, raid.Parties[0], buffs.ArcaneInt)

	// These apply auras on every sim reset
	priestBot := priest.NewBuffBot(sim, raid.Parties[0], buffs.Misery, float64(buffs.SpriestDPS))
	paladinBot := paladin.NewBuffBot(sim, raid.Parties[0], buffs.BlessingOfKings, buffs.ImprovedBlessingOfWisdom, buffs.ImprovedSealOfTheCrusader, buffs.JudgementOfWisdom)

	// Create a fake player and add the agent to do the buffing.
	if buffs.Misery {
		// Misery bot re-applies misery on every sim reset.
		sim.Raid.Parties[0].Players = append(sim.Raid.Parties[0].Players, core.PlayerAgent{
			Player: core.NewPlayer(core.EquipmentSpec{}, core.RaceBonusTypeNone, core.Consumes{}),
			Agent:  priestBot,
		})
	}
	if buffs.JudgementOfWisdom {
		// Judgement of wisdom is an aura that has to be reapplied on every reset.
		// create a bot that acts like a player and rebuffs us.
		sim.Raid.Parties[0].Players = append(sim.Raid.Parties[0].Players, core.PlayerAgent{
			Player: core.NewPlayer(core.EquipmentSpec{}, core.RaceBonusTypeNone, core.Consumes{}),
			Agent:  paladinBot,
		})
	}

	if len(sim.Raid.Parties[0].Players) == 1 && (buffs.TwilightOwl) {
		// Add a new player.
		sim.Raid.Parties[0].Players = append(sim.Raid.Parties[0].Players, core.PlayerAgent{
			Player: core.NewPlayer(core.EquipmentSpec{}, core.RaceBonusTypeNone, core.Consumes{}),
			Agent:  &nullAgent{}, // this player exists to pop items, no agent needed.
		})
	}

	if buffs.TwilightOwl {
		// Add neck to first bot player
		for i, item := range sim.Raid.Parties[0].Players[1].Equip {
			if item.ID == 0 { // no item in this slot.
				sim.Raid.Parties[0].Players[1].Equip[i] = core.ItemsByID[24121]
				sim.Raid.Parties[0].Players[1].ActiveEquip = append(sim.Raid.Parties[0].Players[1].ActiveEquip, &sim.Raid.Parties[0].Players[1].Equip[i])
				break
			}
		}
	}
	if buffs.EyeOfNight {
		// Add neck to first bot player
		for i, item := range sim.Raid.Parties[0].Players[1].Equip {
			if item.ID == 0 { // no item in this slot.
				sim.Raid.Parties[0].Players[1].Equip[i] = core.ItemsByID[24116]
				sim.Raid.Parties[0].Players[1].ActiveEquip = append(sim.Raid.Parties[0].Players[1].ActiveEquip, &sim.Raid.Parties[0].Players[1].Equip[i])
				break
			}
		}
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
			pl.Player.BuffUp(sim, party)
			pl.Agent.BuffUp(sim, party)
		}
	}

	for _, raidParty := range sim.Raid.Parties {
		for _, pl := range raidParty.Players {
			pl.Player.InitialStats = pl.Player.InitialStats.CalculatedTotal()

			// TODO: Figure out how to handle buffs that buff based on other buffs...
			//   for now this hardcoded buffing works...
			if buffs.ImprovedDivineSpirit {
				pl.Player.InitialStats[core.StatSpirit] += 50
			}
			if buffs.BlessingOfKings {
				pl.Player.InitialStats[core.StatIntellect] *= 1.1
				pl.Player.InitialStats[core.StatSpirit] *= 1.1
			}
			if buffs.ImprovedDivineSpirit {
				pl.Player.InitialStats[core.StatSpellPower] += pl.Player.InitialStats[core.StatSpirit] * 0.1
			}

			pl.Player.Stats = pl.Player.InitialStats
		}
	}

	return sim
}

// RunIndividualSim
//  TODO: Should this accept a 'PlayerSettings' instead of a constructed raid and do that work in here?
func RunIndividualSim(sim *core.Simulation, options core.Options) SimResult {
	pid := 0
	for _, raidParty := range sim.Raid.Parties {
		for _, pl := range raidParty.Players {
			pl.ID = pid
			pl.AuraTracker.PID = pid
			pid++
		}
	}
	sim.AuraTracker.PID = -1
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

	result := aggregator.getResult()
	result.Logs = logsBuffer.String()
	return result
}

type MetricsAggregator struct {
	startTime time.Time
	numSims   int

	dpsSum        float64
	dpsSumSquared float64
	dpsMax        float64
	dpsHist       map[int32]int32 // rounded DPS to count

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
	DpsHist  map[int32]int32 // rounded DPS to count

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
		dpsHist:   make(map[int32]int32),
		casts:     make(map[int32]CastMetric),
	}
}

func (aggregator *MetricsAggregator) addMetrics(options core.Options, metrics core.SimMetrics) {
	aggregator.numSims++

	dps := metrics.TotalDamage / options.Encounter.Duration

	aggregator.dpsSum += dps
	aggregator.dpsSumSquared += dps * dps
	aggregator.dpsMax = math.Max(aggregator.dpsMax, dps)

	dpsRounded := int32(math.Round(dps/10) * 10)
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
		if len(cm.Counts) <= idx {
			newArr := make([]int32, idx+1)
			copy(newArr, cm.Counts)
			cm.Counts = newArr

			newDmgs := make([]float64, idx+1)
			copy(newDmgs, cm.Dmgs)
			cm.Dmgs = newDmgs
		}
		cm.Counts[idx]++
		cm.Dmgs[idx] += cast.DidDmg

		aggregator.casts[id] = cm
	}
}

// TODO: Make sure SimResult matches the proto output
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

type nullAgent struct {
}

func (a *nullAgent) ChooseAction(_ *core.Simulation, party *core.Party) core.AgentAction {
	return core.AgentAction{Wait: core.NeverExpires} // makes the bot wait forever and do nothing.
}
func (a *nullAgent) OnActionAccepted(*core.Simulation, core.AgentAction) {
}
func (a *nullAgent) BuffUp(sim *core.Simulation, party *core.Party) {
}
func (a *nullAgent) Reset(sim *core.Simulation) {
}
func (a *nullAgent) OnSpellHit(*core.Simulation, *core.Player, *core.Cast) {}
