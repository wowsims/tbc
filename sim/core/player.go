package core

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/items"
	"github.com/wowsims/tbc/sim/core/stats"
)

type Party struct {
	Players []PlayerAgent
}

type Raid struct {
	Parties []*Party
}

type PlayerAgent struct {
	Agent
	*Player
}

// Advance moves time forward counting down auras, CDs, mana regen, etc
func (pa PlayerAgent) Advance(sim *Simulation, elapsedTime time.Duration, newTime time.Duration) {
	// MP5 regen
	regen := pa.manaRegenPerSecond() * elapsedTime.Seconds()
	pa.Stats[stats.Mana] += regen
	if pa.Stats[stats.Mana] > pa.InitialStats[stats.Mana] {
		pa.Stats[stats.Mana] = pa.InitialStats[stats.Mana]
	}
	if sim.Debug != nil && regen != 0 {
		sim.Debug("-> [%0.1f] Regenerated: %0.1f mana. Total: %0.1f\n", newTime.Seconds(), regen, pa.Stats[stats.Mana])
	}

	// Advance CDs and Auras
	pa.AuraTracker.Advance(sim, pa, newTime)
}

// Player is a data structure to hold all the shared values that all
// class logic shares.
//  All players have stats, equipment, auras, etc
type Player struct {
	ID       int
	Consumes Consumes // pretty sure most classes have consumes to care about.
	Race     RaceBonusType

	InitialStats stats.Stats
	Stats        stats.Stats

	Equip       items.Equipment // Current Gear
	EquipSpec   items.EquipmentSpec
	ActiveEquip []*ActiveItem // cache of gear that can activate.

	*AuraTracker

	// mutatable state
	destructionPotionUsed bool // set to true after using first destruction potion.
}

func (p *Player) HasteBonus() float64 {
	return 1 + (p.Stats[stats.SpellHaste] / 1576)
}

func (p *Player) GetPlayer() *Player {
	return p
}

func NewPlayer(equipSpec items.EquipmentSpec, race RaceBonusType, consumes Consumes) *Player {
	equip := items.NewEquipmentSet(equipSpec)
	// log.Printf("Gear Stats: %s", equip.Stats().Print())
	initialStats := CalculateTotalStats(race, equip, consumes)

	if race == RaceBonusTypeDraenei {
		initialStats[stats.SpellHit] += 12.60 // 1% hit
	}

	player := &Player{
		Race:         race,
		Consumes:     consumes,
		InitialStats: initialStats,
		Stats:        initialStats,
		Equip:        equip,
		EquipSpec:    equipSpec,
		ActiveEquip:  []*ActiveItem{},
		AuraTracker:  NewAuraTracker(),
	}

	// Cache the active abilities for all equipped items.
	for _, eq := range equip {
		act, ok := ActiveItemByID[eq.ID]
		if ok {
			player.ActiveEquip = append(player.ActiveEquip, &act)
		}
		for _, g := range eq.Gems {
			gemAct, ok := ActiveItemByID[g.ID]
			if !ok {
				continue
			}
			player.ActiveEquip = append(player.ActiveEquip, &gemAct)
		}
	}

	return player
}

func (p *Player) Reset() {
	p.destructionPotionUsed = false
	p.Stats = p.InitialStats
	p.AuraTracker.ResetAuras()
}

func (p *Player) BuffUp(sim *Simulation, party *Party) {
	// Activate all permanent item effects.
	for _, actItem := range p.ActiveEquip {
		if actItem.ActivateCD != NeverExpires {
			continue
		}
		p.AddAura(sim, actItem.Activate(sim, party, PlayerAgent{Player: p}))
	}

	p.ActivateSets(sim, party)
	p.TryActivateEquipment(sim, party)
}

// AddAura on player is a simple wrapper around AuraTracker so the
// consumer doesn't need to pass player back into itself.
func (p *Player) AddAura(sim *Simulation, a Aura) {
	p.AuraTracker.AddAura(sim, PlayerAgent{Player: p}, a)
}

// Returns rate of mana regen, as mana / second
func (p *Player) manaRegenPerSecond() float64 {
	return p.Stats[stats.MP5] / 5.0
}

// Pops any on-use trinkets / gear
func (p *Player) TryActivateEquipment(sim *Simulation, party *Party) {
	const sharedCD = time.Second * 20

	for _, item := range p.ActiveEquip {
		if item.Activate == nil || item.ActivateCD == NeverExpires { // ignore non-activatable, and always active items.
			continue
		}
		if p.IsOnCD(item.CoolID, sim.CurrentTime) || (item.SharedID != 0 && p.IsOnCD(item.SharedID, sim.CurrentTime)) {
			continue
		}
		p.AddAura(sim, item.Activate(sim, party, PlayerAgent{Player: p}))
		p.SetCD(item.CoolID, item.ActivateCD+sim.CurrentTime) // put item on CD
		if item.SharedID != 0 {                               // put all shared CDs on
			p.SetCD(item.SharedID, sharedCD+sim.CurrentTime)
		}
	}
}

// Activates set bonuses, returning the list of active bonuses.
func (p *Player) ActivateSets(sim *Simulation, party *Party) []string {
	active := []string{}
	// Activate Set Bonuses
	setItemCount := map[string]int{}

	for _, i := range p.Equip {
		set := itemSetLookup[i.ID]
		if set != nil {
			setItemCount[set.Name]++
			if bonus, ok := set.Bonuses[setItemCount[set.Name]]; ok {
				active = append(active, set.Name+" ("+strconv.Itoa(setItemCount[set.Name])+"pc)")
				p.AddAura(sim, bonus(sim, party, PlayerAgent{Player: p}))
			}
		}
	}

	return active
}

// TODO: This probably should be moved into each class because they all have different base stats.
func BaseStats(race RaceBonusType) stats.Stats {
	stats := stats.Stats{
		stats.Intellect: 104,    // Base int for troll,
		stats.Mana:      2678,   // level 70 shaman
		stats.Spirit:    135,    // lvl 70 shaman
		stats.SpellCrit: 48.576, // base crit for 70 sham
	}
	// TODO: Find race differences.
	switch race {
	case RaceBonusTypeOrc:
	}
	return stats
}

// CalculateTotalStats will take a set of equipment and options and add all stats/buffs/etc together
func CalculateTotalStats(race RaceBonusType, e items.Equipment, c Consumes) stats.Stats {
	stats := BaseStats(race)
	gearStats := e.Stats()
	for i := range stats {
		stats[i] += gearStats[i]
	}
	stats = c.AddStats(stats)
	return stats
}
