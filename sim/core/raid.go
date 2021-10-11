package core

import (
	"fmt"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

type Party struct {
	Players []Agent

	// Party-wide buffs for this party + raid-wide buffs
	Buffs proto.Buffs
}

func (party *Party) Size() int {
	return len(party.Players)
}

func (party *Party) IsFull() bool {
	return party.Size() >= 5
}

func (party *Party) AddPlayer(player Agent) {
	if (party.IsFull()) {
		// Just print a warning, dont need to panic
		fmt.Printf("Party is full\n")
	}

	party.Players = append(party.Players, player)
	player.GetCharacter().Party = party
}

func (party *Party) AddAura(sim *Simulation, aura Aura) {
	for _, agent := range party.Players {
		agent.GetCharacter().AddAura(sim, aura)
	}
}

func (party *Party) AddInitialStats(s stats.Stats) {
	for _, agent := range party.Players {
		agent.GetCharacter().AddInitialStats(s)
	}
}

// AddStats adds a temporary increase to each players stats.
//  This will be reset at the end of the simulation. (using player InitialStats)
func (p *Party) AddStats(s stats.Stats) {
	for _, agent := range p.Players {
		agent.GetCharacter().AddStats(s)
	}
}

type Raid struct {
	Parties []*Party

	// Raid-wide buffs
	Buffs proto.Buffs
}

// Makes a new raid. baseBuffs are extra additional buffs applied to all players in the raid.
func NewRaid(baseBuffs proto.Buffs) *Raid {
	return &Raid{
		Parties: []*Party{
			&Party{ Players: []Agent{}, },
			&Party{ Players: []Agent{}, },
			&Party{ Players: []Agent{}, },
			&Party{ Players: []Agent{}, },
			&Party{ Players: []Agent{}, },
		},
		Buffs: baseBuffs,
	}
}

func (raid *Raid) Size() int {
	totalPlayers := 0
	for _, party := range raid.Parties {
		totalPlayers += party.Size()
	}
	return totalPlayers
}

func (raid *Raid) IsFull() bool {
	return raid.Size() >= 25
}

// Adds the player to the first non-full party in the raid and returns the
// party to which it was added.
func (raid *Raid) AddPlayer(player Agent) *Party {
	for _, party := range raid.Parties {
		if !party.IsFull() {
			party.AddPlayer(player)
			return party
		}
	}

	// All parties are full. For now, just put extra players in party 1.
	party := raid.Parties[0]
	party.AddPlayer(player)
	return party
}

// Adds buffs from every player to the raid and party buffs.
func (raid *Raid) AddPlayerBuffs() {
	// Add raid-wide buffs first.
	for _, party := range raid.Parties {
		for _, player := range party.Players {
			player.AddRaidBuffs(&raid.Buffs)
			player.GetCharacter().AddRaidBuffs(&raid.Buffs)
		}
	}

	// Add party-wide buffs for each party.
	for _, party := range raid.Parties {
		party.Buffs = raid.Buffs
		for _, player := range party.Players {
			player.AddPartyBuffs(&party.Buffs)
			player.GetCharacter().AddPartyBuffs(&party.Buffs)
		}
	}
}

// Applies buffs to the sim and all the players.
func (raid *Raid) ApplyBuffs(sim *Simulation) {
	ApplyBuffsToSim(sim, raid.Buffs)

	for _, party := range raid.Parties {
		for _, player := range party.Players {
			ApplyBuffsToPlayer(player, party.Buffs)
		}
	}
}

func (raid Raid) AddInitialStats(s stats.Stats) {
	for _, party := range raid.Parties {
		party.AddInitialStats(s)
	}
}

func (raid Raid) AddStats(s stats.Stats) {
	for _, party := range raid.Parties {
		party.AddStats(s)
	}
}
