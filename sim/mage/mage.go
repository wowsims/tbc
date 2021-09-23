package mage

import "github.com/wowsims/tbc/sim/core"

func NewBuffBot(sim *core.Simulation, party *core.Party, arcaneInt bool) *Mage {

	if arcaneInt {
		for _, raidParty := range sim.Raid.Parties {
			for _, pl := range raidParty.Players {
				pl.Stats[core.StatIntellect] += 40
				pl.InitialStats[core.StatIntellect] += 40
			}
		}
	}

	return &Mage{}
}

type Mage struct {
	core.Agent
}

func (m *Mage) BuffUp(sim *core.Simulation, party *core.Party) {

}
