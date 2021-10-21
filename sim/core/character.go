package core

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Character is a data structure to hold all the shared values that all
// class logic shares.
// All players have stats, equipment, auras, etc
type Character struct {
	ID       int
	Consumes proto.Consumes
	Race     RaceBonusType
	Class    proto.Class

	InitialStats stats.Stats
	Stats        stats.Stats

	Equip       items.Equipment // Current Gear
	EquipSpec   items.EquipmentSpec
	ActiveEquip []*ActiveItem // cache of gear that can activate.

	// Up references to the Party and Agent for this Character
	Party *Party

	*AuraTracker

	// mutatable state

	// Used for applying the effects of hardcast / channeled spells at a later time.
	// By definition there can be only 1 hardcast spell being cast at any moment.
	HardcastAura Aura

	potionsUsed int32 // Number of potions used
}

func (character *Character) AddInitialStats(s stats.Stats) {
	character.InitialStats = character.InitialStats.Add(s)
}

func (character *Character) AddStats(s stats.Stats) {
	character.Stats = character.Stats.Add(s)
}

func (character *Character) HasteBonus() float64 {
	return 1 + (character.Stats[stats.SpellHaste] / (HasteRatingPerHastePercent * 100))
}
func NewCharacter(equipSpec items.EquipmentSpec, race RaceBonusType, class proto.Class, consumes proto.Consumes, customStats stats.Stats) Character {
	equip := items.NewEquipmentSet(equipSpec)
	// log.Printf("Gear Stats: %s", equip.Stats().Print())
	initialStats := CalculateTotalStats(race, class, equip, consumes).Add(customStats)

	character := Character{
		Race:         race,
		Class:        class,
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
	character.potionsUsed = 0
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

// Returns the amount of time this Character would need to wait in order to reach
// the desired amount of mana, via mana regen.
//
// Assumes that desiredMana > currentMana. Calculation assumes the Character
// will not take any actions during this period that would reset the 5-second rule.
func (character *Character) TimeUntilManaRegen(desiredMana float64) time.Duration {
	// +1 at the end is to deal with floating point math rounding errors.
	return DurationFromSeconds((desiredMana-character.Stats[stats.Mana])/character.manaRegenPerSecond()) + 1
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

	if character.HardcastAura.Expires != 0 && character.HardcastAura.Expires <= newTime {
		character.HardcastAura.OnExpire(sim)
		character.HardcastAura = Aura{}
	}
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
		if item.SharedID != 0 {                                       // put all shared CDs on
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

func (character *Character) AddRaidBuffs(buffs *proto.Buffs) {
}
func (character *Character) AddPartyBuffs(buffs *proto.Buffs) {
	if character.Race == RaceBonusTypeDraenei {
		class := character.Class
		if class == proto.Class_ClassHunter ||
				class == proto.Class_ClassPaladin ||
				class == proto.Class_ClassWarrior {
			buffs.DraeneiRacialMelee = true
		} else if class == proto.Class_ClassMage ||
				class == proto.Class_ClassPriest ||
				class == proto.Class_ClassShaman {
			buffs.DraeneiRacialCaster = true
		}
	}

	if character.Consumes.Drums > 0 {
		buffs.Drums = character.Consumes.Drums
	}

	if character.Equip[items.ItemSlotMainHand].ID == ItemIDAtieshMage {
		buffs.AtieshMage += 1
	}
	if character.Equip[items.ItemSlotMainHand].ID == ItemIDAtieshWarlock {
		buffs.AtieshWarlock += 1
	}

	if character.Equip[items.ItemSlotNeck].ID == ItemIDBraidedEterniumChain {
		buffs.BraidedEterniumChain = true
	}
	if character.Equip[items.ItemSlotNeck].ID == ItemIDChainOfTheTwilightOwl {
		buffs.ChainOfTheTwilightOwl = true
	}
	if character.Equip[items.ItemSlotNeck].ID == ItemIDEyeOfTheNight {
		buffs.EyeOfTheNight = true
	}
	if character.Equip[items.ItemSlotNeck].ID == ItemIDJadePendantOfBlasting {
		buffs.JadePendantOfBlasting = true
	}
}

func (character *Character) EquippedMetaGem(gemID int32) bool {
	for _, gem := range character.Equip[items.ItemSlotHead].Gems {
		if gem.ID == gemID {
			return true
		}
	}
	return false
}

type BaseStatsKey struct {
	Race  RaceBonusType
	Class proto.Class
}

var BaseStats = map[BaseStatsKey]stats.Stats{}

// CalculateTotalStats will take a set of equipment and options and add all stats/buffs/etc together
func CalculateTotalStats(race RaceBonusType, class proto.Class, equipment items.Equipment, consumes proto.Consumes) stats.Stats {
	return BaseStats[BaseStatsKey{ Race: race, Class: class }].Add(equipment.Stats()).Add(ConsumesStats(consumes))
}
