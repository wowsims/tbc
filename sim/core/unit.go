package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/stats"
)

type UnitType int

const (
	PlayerUnit UnitType = iota
	EnemyUnit
	PetUnit
)

// Unit is an abstraction of a Character/Boss/Pet/etc, containing functionality
// shared by all of them.
type Unit struct {
	Type UnitType

	// Index of this unit with its group.
	//  For Players, this is the 0-indexed raid index (0-24).
	//  For Enemies, this is its enemy index.
	//  For Pets, this is the same as the owner's index.
	Index int32

	// Unique label for logging.
	Label string

	Level int32 // Level of Unit, e.g. Bosses are lvl 73.

	// Stats this Character will have at the very start of each Sim iteration.
	// Includes all equipment / buffs / permanent effects but not temporary
	// effects from items / abilities.
	initialStats stats.Stats

	initialPseudoStats stats.PseudoStats

	// Provides aura tracking behavior.
	auraTracker

	// Whether finalize() has been called yet for this Unit.
	// All fields above this may not be altered once finalized is set.
	finalized bool

	// Current stats, including temporary effects.
	stats stats.Stats

	PseudoStats stats.PseudoStats

	// All spells that can be cast by this unit.
	Spellbook []*Spell

	// Statistics describing the results of the sim.
	Metrics CharacterMetrics

	cdTimers []*Timer

	GCD *Timer
}

func (unit *Unit) Log(sim *Simulation, message string, vals ...interface{}) {
	sim.Log("[%s] "+message, append([]interface{}{unit.Label}, vals...)...)
}

func (unit *Unit) GetInitialStat(stat stats.Stat) float64 {
	return unit.initialStats[stat]
}
func (unit *Unit) GetStats() stats.Stats {
	return unit.stats
}
func (unit *Unit) GetStat(stat stats.Stat) float64 {
	return unit.stats[stat]
}

func (unit *Unit) AddStats(stat stats.Stats) {
	unit.stats = unit.stats.Add(stat)
}
func (unit *Unit) AddStat(stat stats.Stat, amount float64) {
	unit.stats[stat] += amount
}

func (unit *Unit) finalize() {
	if unit.finalized {
		return
	}
	unit.finalized = true

	// Make sure we dont accidentally set initial stats instead of stats.
	if !unit.initialStats.Equals(stats.Stats{}) {
		panic("Initial stats may not be set before finalized!")
	}

	// All stats added up to this point are part of the 'initial' stats.
	unit.initialStats = unit.stats
	unit.initialPseudoStats = unit.PseudoStats

	unit.auraTracker.finalize()
}

func (unit *Unit) init(sim *Simulation) {
	unit.auraTracker.init(sim)
}

func (unit *Unit) reset(sim *Simulation) {
	unit.Metrics.reset()
	unit.stats = unit.initialStats
	unit.PseudoStats = unit.initialPseudoStats
	unit.auraTracker.reset(sim)
	for _, spell := range unit.Spellbook {
		spell.reset(sim)
	}
}

// Advance moves time forward counting down auras, CDs, mana regen, etc
func (unit *Unit) advance(sim *Simulation, elapsedTime time.Duration) {
	unit.auraTracker.advance(sim)
}

func (unit *Unit) doneIteration(sim *Simulation) {
	unit.auraTracker.doneIteration(sim)
	for _, spell := range unit.Spellbook {
		spell.doneIteration()
	}
	unit.Metrics.doneIteration(sim.Duration.Seconds())
	unit.resetCDs(sim)
}

// ArmorDamageReduction currently assumes a level 70 attacker
func (unit *Unit) ArmorDamageReduction(armorPen float64) float64 {
	// TODO: Cache this somehow so we dont have to recalculate every time.
	effectiveArmor := MaxFloat(0, unit.stats[stats.Armor]-armorPen)
	return effectiveArmor / (effectiveArmor + 10557.5)
}
