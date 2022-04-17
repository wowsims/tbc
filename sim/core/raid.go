package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

type Party struct {
	Raid  *Raid
	Index int

	Players []Agent
	Pets    []PetAgent // Cached list of all the pets in the party.

	PlayersAndPets []Agent // Cached list of players + pets, concatenated.

	dpsMetrics DistributionMetrics
}

func NewParty(raid *Raid, index int, partyConfig proto.Party) *Party {
	party := &Party{
		Raid:       raid,
		Index:      index,
		dpsMetrics: NewDistributionMetrics(),
	}

	for playerIndex, playerConfig := range partyConfig.Players {
		if playerConfig != nil && playerConfig.Class != proto.Class_ClassUnknown {
			party.Players = append(party.Players, NewAgent(party, playerIndex, *playerConfig))
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

func (party *Party) AddStats(newStats stats.Stats) {
	for _, agent := range party.Players {
		agent.GetCharacter().AddStats(newStats)
	}
}

func (party *Party) AddStat(stat stats.Stat, amount float64) {
	for _, agent := range party.Players {
		agent.GetCharacter().AddStat(stat, amount)
	}
}

func (party *Party) reset(sim *Simulation) {
	for _, agent := range party.Players {
		agent.GetCharacter().reset(sim, agent)

		agent.GetCharacter().SetGCDTimer(sim, 0)
		for _, petAgent := range agent.GetCharacter().Pets {
			if petAgent.GetPet().initialEnabled {
				petAgent.GetPet().Enable(sim, petAgent)
			}
		}
	}

	party.dpsMetrics.reset()
}

func (party *Party) doneIteration(sim *Simulation) {
	for _, agent := range party.Players {
		agent.GetCharacter().doneIteration(sim)
		party.dpsMetrics.Total += agent.GetCharacter().Metrics.dps.Total
	}

	party.dpsMetrics.doneIteration(sim.Duration.Seconds())
}

func (party *Party) GetMetrics(numIterations int32) *proto.PartyMetrics {
	metrics := &proto.PartyMetrics{
		Dps: party.dpsMetrics.ToProto(numIterations),
	}

	playerIdx := 0
	i := 0
	for playerIdx < len(party.Players) {
		player := party.Players[playerIdx]
		if player.GetCharacter().PartyIndex == i {
			metrics.Players = append(metrics.Players, player.GetCharacter().GetMetricsProto(numIterations))
			playerIdx++
		} else {
			metrics.Players = append(metrics.Players, &proto.PlayerMetrics{})
		}
		i++
	}

	return metrics
}
func (party *Party) GetStats() *proto.PartyStats {
	partyStats := &proto.PartyStats{}

	playerIdx := 0
	i := 0
	for playerIdx < len(party.Players) {
		player := party.Players[playerIdx]
		if player.GetCharacter().PartyIndex == i {
			partyStats.Players = append(partyStats.Players, player.GetCharacter().GetStatsProto())
			playerIdx++
		} else {
			partyStats.Players = append(partyStats.Players, &proto.PlayerStats{})
		}
		i++
	}

	return partyStats
}

type Raid struct {
	Parties []*Party

	dpsMetrics DistributionMetrics
}

// Makes a new raid.
func NewRaid(raidConfig proto.Raid) *Raid {
	raid := &Raid{
		dpsMetrics: NewDistributionMetrics(),
	}

	if raidConfig.StaggerStormstrikes {
		enhanceShaman := RaidPlayersWithSpec(raidConfig, proto.Spec_SpecEnhancementShaman)
		if len(enhanceShaman) > 1 {
			stagger := time.Duration(float64(time.Second*10) / float64(len(enhanceShaman)))
			for i, shaman := range enhanceShaman {
				delay := stagger * time.Duration(i)
				shaman.Spec.(*proto.Player_EnhancementShaman).EnhancementShaman.Rotation.FirstStormstrikeDelay = delay.Seconds()
			}
		}
	}

	for partyIndex, partyConfig := range raidConfig.Parties {
		if partyConfig != nil {
			raid.Parties = append(raid.Parties, NewParty(raid, partyIndex, *partyConfig))
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

// Finalize the raid.
func (raid *Raid) finalize(raidConfig proto.Raid) {
	// Precompute the playersAndPets array for each party.
	for _, party := range raid.Parties {
		party.Pets = []PetAgent{}
		for _, player := range party.Players {
			for _, petAgent := range player.GetCharacter().Pets {
				party.Pets = append(party.Pets, petAgent)
			}
		}
		party.PlayersAndPets = make([]Agent, len(party.Players)+len(party.Pets))
		for i, player := range party.Players {
			party.PlayersAndPets[i] = player
		}
		for i, pet := range party.Pets {
			party.PlayersAndPets[len(party.Players)+i] = pet
		}
	}

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

			player.GetCharacter().applyAllEffects(player, raidBuffs, partyBuffs, individualBuffs)
		}
	}

	for _, party := range raid.Parties {
		for _, player := range party.Players {
			player.Finalize(raid)
			player.GetCharacter().Finalize(raid)
		}
	}
}

func (raid Raid) AddStats(s stats.Stats) {
	for _, party := range raid.Parties {
		party.AddStats(s)
	}
}

func (raid Raid) GetPlayerFromRaidTarget(raidTarget proto.RaidTarget) Agent {
	raidIndex := raidTarget.TargetIndex

	partyIndex := int(raidIndex / 5)
	playerIndex := int(raidIndex % 5)

	if partyIndex < 0 || partyIndex >= len(raid.Parties) {
		return nil
	}

	party := raid.Parties[partyIndex]

	if playerIndex < 0 || playerIndex >= len(party.Players) {
		return nil
	}

	return party.Players[playerIndex]
}

func (raid *Raid) reset(sim *Simulation) {
	for _, party := range raid.Parties {
		party.reset(sim)
	}
	raid.dpsMetrics.reset()
}

func (raid *Raid) doneIteration(sim *Simulation) {
	for _, party := range raid.Parties {
		party.doneIteration(sim)
		raid.dpsMetrics.Total += party.dpsMetrics.Total
	}

	raid.dpsMetrics.doneIteration(sim.Duration.Seconds())
}

func (raid *Raid) GetMetrics(numIterations int32) *proto.RaidMetrics {
	metrics := &proto.RaidMetrics{
		Dps: raid.dpsMetrics.ToProto(numIterations),
	}
	for _, party := range raid.Parties {
		metrics.Parties = append(metrics.Parties, party.GetMetrics(numIterations))
	}
	return metrics
}

func (raid *Raid) GetStats() *proto.RaidStats {
	raidStats := &proto.RaidStats{}
	for _, party := range raid.Parties {
		raidStats.Parties = append(raidStats.Parties, party.GetStats())
	}
	return raidStats
}

func SinglePlayerRaidProto(player *proto.Player, partyBuffs *proto.PartyBuffs, raidBuffs *proto.RaidBuffs) *proto.Raid {
	return &proto.Raid{
		Parties: []*proto.Party{
			&proto.Party{
				Players: []*proto.Player{
					player,
				},
				Buffs: partyBuffs,
			},
		},
		Buffs: raidBuffs,
	}
}

func RaidPlayersWithSpec(raid proto.Raid, spec proto.Spec) []*proto.Player {
	var specPlayers []*proto.Player
	for _, party := range raid.Parties {
		for _, player := range party.Players {
			if player != nil && player.GetSpec() != nil && PlayerProtoToSpec(*player) == spec {
				specPlayers = append(specPlayers, player)
			}
		}
	}
	return specPlayers
}
