package druid

import "github.com/wowsims/tbc/sim/core"

func NewBuffBot(sim *core.Simulation, party *core.Party, gotw, moonkin, ravenIdol bool) *Druid {

	if gotw {
		for _, raidParty := range sim.Raid.Parties {
			for _, pl := range raidParty.Players {
				// assumes improved gotw, rounded down to nearest int... not sure if that is accurate.
				pl.Stats[core.StatIntellect] += 18
				pl.InitialStats[core.StatIntellect] += 18
				// FUTURE: Add melee stats here.
			}
		}
	}

	if moonkin {
		for _, pl := range party.Players {
			pl.Stats[core.StatSpellCrit] += 110.4
			pl.InitialStats[core.StatSpellCrit] += 110.4

			if ravenIdol {
				pl.Stats[core.StatSpellCrit] += 20
				pl.InitialStats[core.StatSpellCrit] += 20
			}
		}

	}

	return &Druid{}
}

type Druid struct {
	core.Agent
}

func (m *Druid) BuffUp(sim *core.Simulation, party *core.Party) {

}
