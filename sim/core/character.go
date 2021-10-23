package core

import (
	"sort"
	"time"

	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Need to be a function that returns an Aura rather than an Aura, so captured
// local variables can be reset on Sim reset.
type PermanentAura func(*Simulation, *Character) Aura

// Character is a data structure to hold all the shared values that all
// class logic shares.
// All players have stats, equipment, auras, etc
type Character struct {
	ID    int
	Race  proto.Race
	Class proto.Class

	Equip       items.Equipment // Current Gear
	EquipSpec   items.EquipmentSpec

	// Consumables this Character will be using.
	consumes proto.Consumes

	// Stats this Character will have at the very start of each Sim iteration.
	// Includes all equipment / buffs / permanent effects but not temporary
	// effects from items / abilities.
	initialStats stats.Stats

	// Auras that never expire and should always be active.
	// These are automatically applied on each Sim reset.
	permanentAuras []PermanentAura

	// Cached list of major cooldowns sorted by priority, for resetting quickly.
	initialMajorCooldowns []MajorCooldown

	// Embeddded stat dependency manager.
	stats.StatDependencyManager

	// Up reference to this Character's Party.
	Party *Party

	// Whether Finalize() has been called yet for this Character.
	// All fields above this may not be altered once finalized is set.
	finalized bool

	// Embeddded aura tracker.
	*AuraTracker

	// Current stats, including temporary effects.
	stats stats.Stats

	// Major cooldowns, ordered by next available. This should always contain
	// the same cooldows as initialMajorCooldowns, but the order will change over
	// the course of the sim.
	majorCooldowns []MajorCooldown

	// Used for applying the effects of hardcast / channeled spells at a later time.
	// By definition there can be only 1 hardcast spell being cast at any moment.
	HardcastAura Aura

	potionsUsed int32 // Number of potions used
	bloodlustsUsed int32 // Number of bloodlusts used
}

func NewCharacter(equipSpec items.EquipmentSpec, race proto.Race, class proto.Class, consumes proto.Consumes, customStats stats.Stats) Character {
	equip := items.NewEquipmentSet(equipSpec)

	character := Character{
		Race:         race,
		Class:        class,
		Equip:        equip,
		EquipSpec:    equipSpec,
		consumes:     consumes,

		permanentAuras: []PermanentAura{},

		initialMajorCooldowns: []MajorCooldown{},

		AuraTracker:  NewAuraTracker(),
	}

	character.AddStats(BaseStats[BaseStatsKey{ Race: race, Class: class }])
	character.AddStats(equip.Stats())
	character.AddStats(customStats)

	// Universal stat dependencies
	character.AddStatDependency(stats.StatDependency{
		SourceStat: stats.Agility,
		ModifiedStat: stats.Armor,
		Modifier: func(agility float64, armor float64) float64 {
			return armor + agility * 2
		},
	})
	character.AddStatDependency(stats.StatDependency{
		SourceStat: stats.Intellect,
		ModifiedStat: stats.Mana,
		Modifier: func(intellect float64, mana float64) float64 {
			return mana + intellect * 15
		},
	})

	return character
}

func (character *Character) ApplyAllEffects(agent Agent, buffs proto.Buffs) {
	ApplyRaceEffects(agent)
	character.applyItemEffects(agent)
	character.applyItemSetBonusEffects(agent)
	ApplyConsumeEffects(agent)
	ApplyBuffEffects(agent, buffs)
}

// Apply effects from all equipped items.
func (character *Character) applyItemEffects(agent Agent) {
	for _, eq := range character.Equip {
		applyItemEffect, ok := itemEffects[eq.ID]
		if ok {
			applyItemEffect(agent)
		}

		for _, g := range eq.Gems {
			applyGemEffect, ok := itemEffects[g.ID]
			if ok {
				applyGemEffect(agent)
			}
		}
	}
}

// Registers a permanent aura to this Character which will be re-applied on
// every Sim reset.
func (character *Character) AddPermanentAura(permAura PermanentAura) {
	if character.finalized {
		panic("Permanent auras may not be added once finalized!")
	}

	character.permanentAuras = append(character.permanentAuras, permAura)
}

// Registers a major cooldown to the Character, which will be automatically
// used when available.
func (character *Character) AddMajorCooldown(mcd MajorCooldown) {
	if character.finalized {
		panic("Major cooldowns may not be added once finalized!")
	}

	character.initialMajorCooldowns = append(character.initialMajorCooldowns, mcd)
}

func (character *Character) AddStats(stat stats.Stats) {
	character.stats = character.stats.Add(stat)
}
func (character *Character) AddStat(stat stats.Stat, amount float64) {
	character.stats[stat] += amount
}
func (character *Character) GetInitialStat(stat stats.Stat) float64 {
	return character.initialStats[stat]
}
func (character *Character) GetStats() stats.Stats {
	return character.stats
}
func (character *Character) GetStat(stat stats.Stat) float64 {
	return character.stats[stat]
}
func (character *Character) MaxMana() float64 {
	return character.GetInitialStat(stats.Mana)
}
func (character *Character) CurrentMana() float64 {
	return character.GetStat(stats.Mana)
}

func (character *Character) HasteBonus() float64 {
	return 1 + (character.stats[stats.SpellHaste] / (HasteRatingPerHastePercent * 100))
}

func (character *Character) Finalize() {
	if character.finalized {
		return
	}
	character.finalized = true

	// Make sure we dont accidentally set initial stats instead of stats.
	if !character.initialStats.Equals(stats.Stats{}) {
		panic("Initial stats may not be set before finalized!")
	}
	character.StatDependencyManager.Finalize()
	character.stats = character.ApplyStatDependencies(character.stats)

	// All stats added up to this point are part of the 'initial' stats.
	character.initialStats = character.stats

	// Sort major cooldowns by descending priority so they get used in the correct order.
	sort.SliceStable(character.initialMajorCooldowns, func(i, j int) bool {
		return character.initialMajorCooldowns[i].Priority > character.initialMajorCooldowns[j].Priority
	})
}

func (character *Character) Reset(sim *Simulation) {
	character.potionsUsed = 0
	character.bloodlustsUsed = 0
	character.stats = character.initialStats

	character.majorCooldowns = make([]MajorCooldown, len(character.initialMajorCooldowns))
	copy(character.majorCooldowns, character.initialMajorCooldowns)

	character.AuraTracker.ResetAuras()
	for _, permAura := range character.permanentAuras {
		aura := permAura(sim, character)
		aura.Expires = NeverExpires
		character.AddAura(sim, aura)
	}
}

func (character *Character) TryUseCooldowns(sim *Simulation) {
	anyCooldownsUsed := false
	for curIdx := 0; curIdx < len(character.majorCooldowns) && !character.majorCooldowns[curIdx].IsOnCD(sim, character); curIdx++ {
		success := character.majorCooldowns[curIdx].TryActivate(sim, character)
		anyCooldownsUsed = anyCooldownsUsed || success
	}

	if anyCooldownsUsed {
		// Re-sort by availability. 
		// TODO: Probably a much faster way to do this, especially since we know which cooldowns need to be re-ordered.
		sort.Slice(character.majorCooldowns, func(i, j int) bool {
			return character.majorCooldowns[i].GetRemainingCD(sim, character) < character.majorCooldowns[j].GetRemainingCD(sim, character)
		})
	}
}

// This function should be called if the CD for a major cooldown changes outside
// of its TryActivate() call.
func (character *Character) UpdateMajorCooldowns(sim *Simulation) {
	sort.Slice(character.majorCooldowns, func(i, j int) bool {
		return character.majorCooldowns[i].GetRemainingCD(sim, character) < character.majorCooldowns[j].GetRemainingCD(sim, character)
	})
}

// Returns rate of mana regen, as mana / second
func (character *Character) manaRegenPerSecond() float64 {
	return character.stats[stats.MP5] / 5.0
}

// Returns the amount of time this Character would need to wait in order to reach
// the desired amount of mana, via mana regen.
//
// Assumes that desiredMana > currentMana. Calculation assumes the Character
// will not take any actions during this period that would reset the 5-second rule.
func (character *Character) TimeUntilManaRegen(desiredMana float64) time.Duration {
	// +1 at the end is to deal with floating point math rounding errors.
	return DurationFromSeconds((desiredMana-character.CurrentMana())/character.manaRegenPerSecond()) + 1
}

// Advance moves time forward counting down auras, CDs, mana regen, etc
func (character *Character) Advance(sim *Simulation, elapsedTime time.Duration, newTime time.Duration) {
	// MP5 regen
	regen := character.manaRegenPerSecond() * elapsedTime.Seconds()
	character.stats[stats.Mana] += regen
	if character.CurrentMana() > character.MaxMana() {
		character.stats[stats.Mana] = character.MaxMana()
	}
	if sim.Log != nil && regen != 0 {
		sim.Log("-> [%0.1f] Regenerated: %0.1f mana. Total: %0.1f\n", newTime.Seconds(), regen, character.CurrentMana())
	}

	// Advance CDs and Auras
	character.AuraTracker.Advance(sim, newTime)

	if character.HardcastAura.Expires != 0 && character.HardcastAura.Expires <= newTime {
		character.HardcastAura.OnExpire(sim)
		character.HardcastAura = Aura{}
	}
}

func (character *Character) AddRaidBuffs(buffs *proto.Buffs) {
}
func (character *Character) AddPartyBuffs(buffs *proto.Buffs) {
	if character.Race == proto.Race_RaceDraenei {
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

	if character.consumes.Drums > 0 {
		buffs.Drums = character.consumes.Drums
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

func (character *Character) HasTrinketEquipped(itemID int32) bool {
	return character.Equip[items.ItemSlotTrinket1].ID == itemID ||
			character.Equip[items.ItemSlotTrinket2].ID == itemID
}

func (character *Character) HasMetaGemEquipped(gemID int32) bool {
	for _, gem := range character.Equip[items.ItemSlotHead].Gems {
		if gem.ID == gemID {
			return true
		}
	}
	return false
}

type BaseStatsKey struct {
	Race  proto.Race
	Class proto.Class
}

var BaseStats = map[BaseStatsKey]stats.Stats{}
