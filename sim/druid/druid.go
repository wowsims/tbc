package druid

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func NewBuffBot(sim *core.Simulation, party *core.Party, gotw, moonkin, ravenIdol bool) *Druid {

	if gotw {
		for _, raidParty := range sim.Raid.Parties {
			for _, pl := range raidParty.Players {
				// assumes improved gotw, rounded down to nearest int... not sure if that is accurate.
				pl.Stats[stats.Intellect] += 18
				pl.InitialStats[stats.Intellect] += 18
				// FUTURE: Add melee stats here.
			}
		}
	}

	if moonkin {
		s := stats.Stats{stats.SpellCrit: 110.4}
		if ravenIdol {
			s[stats.SpellCrit] += 20
		}
		party.AddInitialStats(s)
	}

	return &Druid{}
}

type Druid struct {
	core.Agent
}

func (m *Druid) BuffUp(sim *core.Simulation, party *core.Party) {

}
