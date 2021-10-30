package core

import (
	"fmt"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

type Party struct {
	Players []Agent

	// Party-wide buffs for this party + raid-wide buffs
	buffs proto.Buffs
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

func (p *Party) AddStats(s stats.Stats) {
	for _, agent := range p.Players {
		agent.GetCharacter().AddStats(s)
	}
}

type Raid struct {
	Parties []*Party

	// Raid-wide buffs
	buffs proto.Buffs
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
		buffs: baseBuffs,
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
func (raid *Raid) addPlayerBuffs() {
	// Add raid-wide buffs first.
	for _, party := range raid.Parties {
		for _, player := range party.Players {
			player.AddRaidBuffs(&raid.buffs)
			player.GetCharacter().AddRaidBuffs(&raid.buffs)
		}
	}

	// Add party-wide buffs for each party.
	for _, party := range raid.Parties {
		party.buffs = raid.buffs
		for _, player := range party.Players {
			player.AddPartyBuffs(&party.buffs)
			player.GetCharacter().AddPartyBuffs(&party.buffs)
		}
	}
}

// Applies buffs to the sim and all the players.
func (raid *Raid) applyAllEffects(sim *Simulation) {
	applyBuffsToSim(sim, raid.buffs)

	for _, party := range raid.Parties {
		for _, player := range party.Players {
			player.GetCharacter().applyAllEffects(player, party.buffs)
		}
	}
}

// Finalize the raid.
func (raid *Raid) Finalize(sim *Simulation) {
	raid.addPlayerBuffs()
	raid.applyAllEffects(sim)

	for _, party := range raid.Parties {
		for _, player := range party.Players {
			player.GetCharacter().Finalize()
		}
	}
}

func (raid Raid) AddStats(s stats.Stats) {
	for _, party := range raid.Parties {
		party.AddStats(s)
	}
}
