package core

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type Party struct {
	Players []PlayerAgent
}

type Raid struct {
	Parties []*Party
}

// Player is a data structure to hold all the shared values that all
// class logic shares.
//  All players have stats, equipment, auras, etc
type Player struct {
	ID       int
	Consumes Consumes // pretty sure most classes have consumes to care about.
	Race     RaceBonusType

	InitialStats Stats
	Stats        Stats

	Equip       Equipment // Current Gear
	EquipSpec   EquipmentSpec
	ActiveEquip []*Item // cache of gear that can activate.

	*AuraTracker

	// mutatable state
	destructionPotionUsed bool // set to true after using first destruction potion.
}

func (p *Player) HasteBonus() float64 {
	return 1 + (p.Stats[StatSpellHaste] / 1576)
}

func (p *Player) GetPlayer() *Player {
	return p
}

func NewPlayer(equipSpec EquipmentSpec, race RaceBonusType, consumes Consumes) *Player {
	equip := NewEquipmentSet(equipSpec)
	fmt.Println(equip.Stats().Print())
	initialStats := CalculateTotalStats(race, equip, consumes)

	if race == RaceBonusTypeDraenei {
		initialStats[StatSpellHit] += 12.60 // 1% hit
	}

	player := &Player{
		Race:         race,
		Consumes:     consumes,
		InitialStats: initialStats,
		Stats:        initialStats,
		Equip:        equip,
		EquipSpec:    equipSpec,
		ActiveEquip:  []*Item{},
		AuraTracker:  NewAuraTracker(),
	}

	for i, eq := range equip {
		if eq.Activate != nil {
			log.Printf("Adding %s to active equipment...", eq.Name)
			player.ActiveEquip = append(player.ActiveEquip, &equip[i])
		}
		for _, g := range eq.Gems {
			if g.Activate != nil {
				player.ActiveEquip = append(player.ActiveEquip, &equip[i])
			}
		}
	}

	return player
}

// Advance moves time forward counting down auras, CDs, mana regen, etc
func (p *Player) Advance(sim *Simulation, elapsedTime time.Duration, newTime time.Duration) {
	// MP5 regen
	p.Stats[StatMana] += p.manaRegenPerSecond() * elapsedTime.Seconds()
	if p.Stats[StatMana] > p.InitialStats[StatMana] {
		p.Stats[StatMana] = p.InitialStats[StatMana]
	}

	// Advance CDs and Auras
	p.AuraTracker.Advance(sim, p, newTime)
}

func (p *Player) Reset() {
	p.Stats = p.InitialStats
	p.AuraTracker.ResetAuras()
}

func (p *Player) BuffUp(sim *Simulation, party *Party) {
	// Activate all permanent item effects.
	for _, item := range p.ActiveEquip {
		if item.Activate != nil && item.ActivateCD == NeverExpires {
			p.AddAura(sim, item.Activate(sim, party, p))
		}
		for _, g := range item.Gems {
			if g.Activate != nil {
				p.AddAura(sim, g.Activate(sim, party, p))
			}
		}
	}

	p.ActivateSets(sim, party)
	p.TryActivateEquipment(sim, party)
}

// AddAura on player is a simple wrapper around AuraTracker so the
// consumer doesn't need to pass player back into itself.
func (p *Player) AddAura(sim *Simulation, a Aura) {
	p.AuraTracker.AddAura(sim, p, a)
}

// Returns rate of mana regen, as mana / second
func (p *Player) manaRegenPerSecond() float64 {
	return p.Stats[StatMP5] / 5.0
}

// Pops any on-use trinkets / gear
func (p *Player) TryActivateEquipment(sim *Simulation, party *Party) {
	for _, item := range p.ActiveEquip {
		if item.Activate == nil || item.ActivateCD == NeverExpires { // ignore non-activatable, and always active items.
			continue
		}
		if p.IsOnCD(item.CoolID, sim.CurrentTime) {
			continue
		}
		if item.ItemType == ItemTypeTrinket && p.IsOnCD(MagicIDAllTrinket, sim.CurrentTime) {
			continue
		}
		p.AddAura(sim, item.Activate(sim, party, p))
		p.SetCD(item.CoolID, item.ActivateCD+sim.CurrentTime)
		if item.ItemType == ItemTypeTrinket {
			p.SetCD(MagicIDAllTrinket, time.Second*30+sim.CurrentTime)
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
				p.AddAura(sim, bonus(sim, party, p))
			}
		}
	}

	return active
}
