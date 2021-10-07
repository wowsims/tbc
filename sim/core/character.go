package core

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Character is a data structure to hold all the shared values that all
// class logic shares.
// All players have stats, equipment, auras, etc
type Character struct {
	ID       int
	Consumes Consumes // pretty sure most classes have consumes to care about.
	Race     RaceBonusType

	InitialStats stats.Stats
	Stats        stats.Stats

	Equip       items.Equipment // Current Gear
	EquipSpec   items.EquipmentSpec
	ActiveEquip []*ActiveItem // cache of gear that can activate.

	// Up references to the Party and Agent for this Character
	Party *Party

	*AuraTracker

	// mutatable state
	destructionPotionUsed bool // set to true after using first destruction potion.
}

func (character *Character) AddInitialStats(s stats.Stats) {
	character.InitialStats = character.InitialStats.Add(s)
}

func (character *Character) AddStats(s stats.Stats) {
	character.Stats = character.Stats.Add(s)
}

func (character *Character) HasteBonus() float64 {
	return 1 + (character.Stats[stats.SpellHaste] / 1576)
}
func NewCharacter(equipSpec items.EquipmentSpec, race RaceBonusType, consumes Consumes, customStats stats.Stats) *Character {
	equip := items.NewEquipmentSet(equipSpec)
	// log.Printf("Gear Stats: %s", equip.Stats().Print())
	initialStats := CalculateTotalStats(race, equip, consumes).Add(customStats)

	character := &Character{
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
			character.ActiveEquip = append(character.ActiveEquip, &act)
		}
		for _, g := range eq.Gems {
			gemAct, ok := ActiveItemByID[g.ID]
			if !ok {
				continue
			}
			character.ActiveEquip = append(character.ActiveEquip, &gemAct)
		}
	}

	return character
}

func (character *Character) Reset() {
	character.destructionPotionUsed = false
	character.Stats = character.InitialStats
	character.AuraTracker.ResetAuras()
}

func (character *Character) BuffUp(sim *Simulation, agent Agent) {
	// Activate all permanent item effects.
	for _, actItem := range character.ActiveEquip {
		if actItem.ActivateCD != NeverExpires {
			continue
		}
		character.AddAura(sim, actItem.Activate(sim, agent))
	}

	character.ActivateSets(sim, agent)
	character.TryActivateEquipment(sim, agent)
}

// AddAura on player is a simple wrapper around AuraTracker so the
// consumer doesn't need to pass player back into itself.
func (character *Character) AddAura(sim *Simulation, aura Aura) {
	character.AuraTracker.AddAura(sim, aura)
}

// Returns rate of mana regen, as mana / second
func (character *Character) manaRegenPerSecond() float64 {
	return character.Stats[stats.MP5] / 5.0
}

// Advance moves time forward counting down auras, CDs, mana regen, etc
func (character *Character) Advance(sim *Simulation, elapsedTime time.Duration, newTime time.Duration) {
	// MP5 regen
	regen := character.manaRegenPerSecond() * elapsedTime.Seconds()
	character.Stats[stats.Mana] += regen
	if character.Stats[stats.Mana] > character.InitialStats[stats.Mana] {
		character.Stats[stats.Mana] = character.InitialStats[stats.Mana]
	}
	if sim.Log != nil && regen != 0 {
		sim.Log("-> [%0.1f] Regenerated: %0.1f mana. Total: %0.1f\n", newTime.Seconds(), regen, character.Stats[stats.Mana])
	}

	// Advance CDs and Auras
	character.AuraTracker.Advance(sim, newTime)
}

// Pops any on-use trinkets / gear
func (character *Character) TryActivateEquipment(sim *Simulation, agent Agent) {
	const sharedCD = time.Second * 20

	for _, item := range character.ActiveEquip {
		if item.Activate == nil || item.ActivateCD == NeverExpires { // ignore non-activatable, and always active items.
			continue
		}
		if character.IsOnCD(item.CoolID, sim.CurrentTime) || (item.SharedID != 0 && character.IsOnCD(item.SharedID, sim.CurrentTime)) {
			continue
		}
		character.AddAura(sim, item.Activate(sim, agent))
		character.SetCD(item.CoolID, item.ActivateCD+sim.CurrentTime) // put item on CD
		if item.SharedID != 0 {                               // put all shared CDs on
			character.SetCD(item.SharedID, sharedCD+sim.CurrentTime)
		}
	}
}

// Activates set bonuses, returning the list of active bonuses.
func (character *Character) ActivateSets(sim *Simulation, agent Agent) []string {
	active := []string{}
	// Activate Set Bonuses
	setItemCount := map[string]int{}

	for _, i := range character.Equip {
		set := itemSetLookup[i.ID]
		if set != nil {
			setItemCount[set.Name]++
			if bonus, ok := set.Bonuses[setItemCount[set.Name]]; ok {
				active = append(active, set.Name+" ("+strconv.Itoa(setItemCount[set.Name])+"pc)")
				character.AddAura(sim, bonus(sim, agent))
			}
		}
	}

	return active
}

const (
	AtieshMage            = 22589
	AtieshWarlock         = 22630
	BraidedEterniumChain  = 24114
	ChainOfTheTwilightOwl = 24121
	EyeOfTheNight         = 24116
	JadePendantOfBlasting = 20966
)

func (character *Character) AddRaidBuffs(buffs *Buffs) {
}
func (character *Character) AddPartyBuffs(buffs *Buffs) {
	// TODO: Figure out how to sync these with general settings
	//if character.Equip[items.ItemSlotMainHand].ID == AtieshMage {
	//	buffs.AtieshMage += 1
	//}
	//if character.Equip[items.ItemSlotMainHand].ID == AtieshWarlock {
	//	buffs.AtieshWarlock += 1
	//}

	if character.Equip[items.ItemSlotNeck].ID == BraidedEterniumChain {
		buffs.BraidedEterniumChain = true
	}
	if character.Equip[items.ItemSlotNeck].ID == ChainOfTheTwilightOwl {
		buffs.ChainOfTheTwilightOwl = true
	}
	if character.Equip[items.ItemSlotNeck].ID == EyeOfTheNight {
		buffs.EyeOfTheNight = true
	}
	if character.Equip[items.ItemSlotNeck].ID == JadePendantOfBlasting {
		buffs.JadePendantOfBlasting = true
	}
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
func CalculateTotalStats(race RaceBonusType, equipment items.Equipment, consumes Consumes) stats.Stats {
	totalStats := BaseStats(race).Add(equipment.Stats()).Add(consumes.Stats())

	if race == RaceBonusTypeDraenei {
		totalStats[stats.SpellHit] += 12.60 // 1% hit
	}

	return totalStats
}
