package runner

import (
	"fmt"
	"strings"

	"github.com/wowsims/tbc/items"
	"github.com/wowsims/tbc/sim/api"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
	"github.com/wowsims/tbc/sim/druid"
	"github.com/wowsims/tbc/sim/mage"
	"github.com/wowsims/tbc/sim/paladin"
	"github.com/wowsims/tbc/sim/priest"
	"github.com/wowsims/tbc/sim/shaman"
)

// TODO: Should we move the 'buff bots' to a subpackage so we dont have to import the full class unless we are actually wanting the whole class?

type IndividualParams struct {
	Equip    items.EquipmentSpec
	Race     core.RaceBonusType
	Consumes core.Consumes
	Buffs    core.Buffs
	Options  core.Options

	PlayerOptions *api.PlayerOptions

	CustomStats stats.Stats
}

func SetupIndividualSim(params IndividualParams) *core.Simulation {
	character := core.NewCharacter(params.Equip, params.Race, params.Consumes, params.CustomStats)

	var agent core.Agent

	switch v := params.PlayerOptions.Spec.(type) {
	case *api.PlayerOptions_ElementalShaman:
		agent = shaman.NewElementalShaman(character, params.PlayerOptions, params.Buffs)
		break;
	default:
		panic(fmt.Sprintf("Class not supported: %v\n", v))
	}

	party := &core.Party{
		Players: []core.Agent{},
	}
	raid := &core.Raid{Parties: []*core.Party{party}}

	party.AddPlayer(agent)


	buffs := params.Buffs
	options := params.Options

	sim := core.NewSim(raid, options)

	// These buffs are a one-time apply... no need to add the bots to the raid group.
	//  the constructors apply their buffs to the party.
	party.AddPlayer(druid.NewBuffBot(sim, party, buffs.GiftOfTheWild, buffs.Moonkin, buffs.MoonkinRavenGoddess))
	party.AddPlayer(mage.NewBuffBot(sim, party, buffs.ArcaneInt))

	// These apply auras on every sim reset
	party.AddPlayer(priest.NewBuffBot(sim, party, buffs.Misery, float64(buffs.SpriestDPS)))
	party.AddPlayer(paladin.NewBuffBot(sim, party, buffs.BlessingOfKings, buffs.ImprovedBlessingOfWisdom, buffs.ImprovedSealOfTheCrusader, buffs.JudgementOfWisdom))

	// TODO: Buffs as a whole needs refactoring
	if buffs.TwilightOwl {
		const bonus = 2 * 22.08 // 2% crit
		party.AddInitialStats(stats.Stats{stats.SpellCrit: bonus})
	}
	if buffs.EyeOfNight {
		const bonus = 34
		party.AddInitialStats(stats.Stats{stats.SpellPower: bonus})
	}
	if buffs.ImprovedDivineSpirit {
		raid.AddInitialStats(stats.Stats{stats.Spirit: 50})
	}

	sim.Reset()

	// Now apply all the 'final' stat improvements.
	// TODO: Figure out how to handle buffs that buff based on other buffs...
	//   for now this hardcoded buffing works...
	for _, raidParty := range sim.Raid.Parties {
		for _, player := range raidParty.Players {
			if buffs.BlessingOfKings {
				player.GetCharacter().InitialStats[stats.Stamina] *= 1.1
				player.GetCharacter().InitialStats[stats.Agility] *= 1.1
				player.GetCharacter().InitialStats[stats.Strength] *= 1.1
				player.GetCharacter().InitialStats[stats.Intellect] *= 1.1
				player.GetCharacter().InitialStats[stats.Spirit] *= 1.1
			}
			if buffs.ImprovedDivineSpirit {
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

// RunIndividualSim
func RunIndividualSim(sim *core.Simulation) SimResult {
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
		metrics := sim.Run()
		aggregator.addMetrics(sim.Options, metrics)
		sim.ReturnCasts(metrics.Casts)
	}

	result := aggregator.getResult()
	result.Logs = logsBuffer.String()
	return result
}
