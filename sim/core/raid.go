package core

import (
	"fmt"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

type Party struct {
	Players []Agent
}

func NewParty(partyConfig proto.Party) *Party {
	party := &Party{}

	for _, playerConfig := range partyConfig.Players {
		if playerConfig != nil {
			party.AddPlayer(NewAgent(*playerConfig))
		}
	}

	return party
}

func (party *Party) Size() int {
	return len(party.Players)
}

func (party *Party) IsFull() bool {
	return party.Size() >= 5
}

func (party *Party) AddPlayer(player Agent) {
	if party.IsFull() {
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

func (party *Party) AddStats(newStats stats.Stats) {
	for _, agent := range party.Players {
		agent.GetCharacter().AddStats(newStats)
	}
}

func (party *Party) doneIteration(simDuration time.Duration) {
	for _, agent := range party.Players {
		agent.GetCharacter().doneIteration(simDuration)
	}
}

func (party *Party) GetMetrics(numIterations int32) *proto.PartyMetrics {
	metrics := &proto.PartyMetrics{}
	for _, agent := range party.Players {
		metrics.Players = append(metrics.Players, agent.GetCharacter().GetMetricsProto(numIterations))
	}
	return metrics
}

type Raid struct {
	Parties []*Party
}

// Makes a new raid.
func NewRaid(raidConfig proto.Raid) *Raid {
	raid := &Raid{}

	for _, partyConfig := range raidConfig.Parties {
		if partyConfig != nil {
			raid.Parties = append(raid.Parties, NewParty(*partyConfig))
		}
	}

	pid := 0
	for _, party := range raid.Parties {
		for _, player := range party.Players {
			player.GetCharacter().ID = pid
			player.GetCharacter().auraTracker.playerID = pid
			pid++
		}
	}

	raid.finalize(raidConfig)

	return raid
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

// Finalize the raid.
func (raid *Raid) finalize(raidConfig proto.Raid) {
	// Compute the full raid buffs from the raid.
	raidBuffs := proto.RaidBuffs{}
	if raidConfig.Buffs != nil {
		raidBuffs = *raidConfig.Buffs
	}
	for _, party := range raid.Parties {
		for _, player := range party.Players {
			player.AddRaidBuffs(&raidBuffs)
			player.GetCharacter().AddRaidBuffs(&raidBuffs)
		}
	}

	for partyIdx, party := range raid.Parties {
		// Compute the full party buffs for this party.
		partyConfig := *raidConfig.Parties[partyIdx]
		partyBuffs := proto.PartyBuffs{}
		if partyConfig.Buffs != nil {
			partyBuffs = *partyConfig.Buffs
		}
		for _, player := range party.Players {
			player.AddPartyBuffs(&partyBuffs)
			player.GetCharacter().AddPartyBuffs(&partyBuffs)
		}

		// Apply all buffs to the players in this party.
		for playerIdx, player := range party.Players {
			playerConfig := *partyConfig.Players[playerIdx]
			individualBuffs := proto.IndividualBuffs{}
			if playerConfig.Buffs != nil {
				individualBuffs = *playerConfig.Buffs
			}

			player.GetCharacter().applyAllEffects(player)
			applyBuffEffects(player, raidBuffs, partyBuffs, individualBuffs)
			applyConsumeEffects(player, partyBuffs)
		}
	}

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

func (raid *Raid) doneIteration(simDuration time.Duration) {
	for _, party := range raid.Parties {
		party.doneIteration(simDuration)
	}
}

func (raid *Raid) GetMetrics(numIterations int32) *proto.RaidMetrics {
	metrics := &proto.RaidMetrics{}
	for _, party := range raid.Parties {
		metrics.Parties = append(metrics.Parties, party.GetMetrics(numIterations))
	}
	return metrics
}
