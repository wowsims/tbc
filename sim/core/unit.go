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

	// Environment in which this Unit exists. This will be nil until after the
	// construction phase.
	Env *Environment

	// Stats this Unit will have at the very start of each Sim iteration.
	// Includes all equipment / buffs / permanent effects but not temporary
	// effects from items / abilities.
	initialStats stats.Stats

	initialPseudoStats stats.PseudoStats

	// Cast speed without any temporary effects.
	initialCastSpeed float64

	// Melee swing speed without any temporary effects.
	initialMeleeSwingSpeed float64

	// Ranged swing speed without any temporary effects.
	initialRangedSwingSpeed float64

	// Provides aura tracking behavior.
	auraTracker

	// Current stats, including temporary effects.
	stats stats.Stats

	PseudoStats stats.PseudoStats

	rageBar
	energyBar

	// All spells that can be cast by this unit.
	Spellbook []*Spell

	// AutoAttacks is the manager for auto attack swings.
	// Must be enabled to use, with "EnableAutoAttacks()".
	AutoAttacks AutoAttacks

	// Statistics describing the results of the sim.
	Metrics CharacterMetrics

	cdTimers []*Timer

	GCD *Timer

	// Used for applying the effects of hardcast / channeled spells at a later time.
	// By definition there can be only 1 hardcast spell being cast at any moment.
	Hardcast Hardcast

	// GCD-related PendingActions.
	gcdAction      *PendingAction
	hardcastAction *PendingAction

	// Fields related to waiting for certain events to happen.
	waitingForMana float64
	waitStartTime  time.Duration

	// Cached mana return values per tick.
	manaTickWhileCasting    float64
	manaTickWhileNotCasting float64
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

// Returns whether the indicates stat is currently modified by a temporary bonus.
func (unit *Unit) HasTemporaryBonusForStat(stat stats.Stat) bool {
	return unit.initialStats[stat] != unit.stats[stat]
}

// Returns if spell casting has any temporary increases active.
func (unit *Unit) HasTemporarySpellCastSpeedIncrease() bool {
	return unit.HasTemporaryBonusForStat(stats.SpellHaste) ||
		unit.PseudoStats.CastSpeedMultiplier != 1
}

// Returns if melee swings have any temporary increases active.
func (unit *Unit) HasTemporaryMeleeSwingSpeedIncrease() bool {
	return unit.SwingSpeed() != unit.initialMeleeSwingSpeed
}

// Returns if ranged swings have any temporary increases active.
func (unit *Unit) HasTemporaryRangedSwingSpeedIncrease() bool {
	return unit.RangedSwingSpeed() != unit.initialRangedSwingSpeed
}

func (unit *Unit) InitialCastSpeed() float64 {
	return unit.initialCastSpeed
}

func (unit *Unit) SpellGCD() time.Duration {
	return MaxDuration(GCDMin, time.Duration(float64(GCDDefault)/unit.CastSpeed()))
}

func (unit *Unit) CastSpeed() float64 {
	return unit.PseudoStats.CastSpeedMultiplier * (1 + (unit.stats[stats.SpellHaste] / (HasteRatingPerHastePercent * 100)))
}

func (unit *Unit) ApplyCastSpeed(dur time.Duration) time.Duration {
	return time.Duration(float64(dur) / unit.CastSpeed())
}

func (unit *Unit) SwingSpeed() float64 {
	return unit.PseudoStats.MeleeSpeedMultiplier * (1 + (unit.stats[stats.MeleeHaste] / (HasteRatingPerHastePercent * 100)))
}

func (unit *Unit) RangedSwingSpeed() float64 {
	return unit.PseudoStats.RangedSpeedMultiplier * (1 + (unit.stats[stats.MeleeHaste] / (HasteRatingPerHastePercent * 100)))
}

func (unit *Unit) AddMeleeHaste(sim *Simulation, amount float64) {
	if amount > 0 {
		mod := 1 + (amount / (HasteRatingPerHastePercent * 100))
		unit.AutoAttacks.ModifySwingTime(sim, mod)
	} else {
		mod := 1 / (1 + (-amount / (HasteRatingPerHastePercent * 100)))
		unit.AutoAttacks.ModifySwingTime(sim, mod)
	}
	unit.stats[stats.MeleeHaste] += amount

	// Could add melee haste to pets too, but not aware of any pets that scale with
	// owner's melee haste.
}

// MultiplyMeleeSpeed will alter the attack speed multiplier and change swing speed of all autoattack swings in progress.
func (unit *Unit) MultiplyMeleeSpeed(sim *Simulation, amount float64) {
	unit.PseudoStats.MeleeSpeedMultiplier *= amount
	unit.AutoAttacks.ModifySwingTime(sim, amount)
}

func (unit *Unit) MultiplyRangedSpeed(sim *Simulation, amount float64) {
	unit.PseudoStats.RangedSpeedMultiplier *= amount
}

// Helper for when both MultiplyMeleeSpeed and MultiplyRangedSpeed are needed.
func (unit *Unit) MultiplyAttackSpeed(sim *Simulation, amount float64) {
	unit.PseudoStats.MeleeSpeedMultiplier *= amount
	unit.PseudoStats.RangedSpeedMultiplier *= amount
	unit.AutoAttacks.ModifySwingTime(sim, amount)
}

func (unit *Unit) finalize() {
	if unit.Env.IsFinalized() {
		panic("Unit already finalized!")
	}

	// Make sure we dont accidentally set initial stats instead of stats.
	if !unit.initialStats.Equals(stats.Stats{}) {
		panic("Initial stats may not be set before finalized: " + unit.initialStats.String())
	}

	// All stats added up to this point are part of the 'initial' stats.
	unit.initialStats = unit.stats
	unit.initialPseudoStats = unit.PseudoStats
	unit.initialCastSpeed = unit.CastSpeed()
	unit.initialMeleeSwingSpeed = unit.SwingSpeed()
	unit.initialRangedSwingSpeed = unit.RangedSwingSpeed()
}

func (unit *Unit) init(sim *Simulation) {
	unit.auraTracker.init(sim)
}

func (unit *Unit) reset(sim *Simulation, agent Agent) {
	unit.Metrics.reset()
	unit.stats = unit.initialStats
	unit.PseudoStats = unit.initialPseudoStats
	unit.auraTracker.reset(sim)
	for _, spell := range unit.Spellbook {
		spell.reset(sim)
	}

	if agent != nil {
		unit.gcdAction = unit.newGCDAction(sim, agent)
	}
}

// Advance moves time forward counting down auras, CDs, mana regen, etc
func (unit *Unit) advance(sim *Simulation, elapsedTime time.Duration) {
	unit.auraTracker.advance(sim)

	if unit.Hardcast.Expires != 0 && unit.Hardcast.Expires <= sim.CurrentTime {
		unit.Hardcast.Expires = 0
		unit.Hardcast.OnExpire(sim)
	}
}

func (unit *Unit) doneIteration(sim *Simulation) {
	unit.Hardcast = Hardcast{}
	unit.doneIterationGCD(sim.Duration)

	unit.auraTracker.doneIteration(sim)
	for _, spell := range unit.Spellbook {
		spell.doneIteration()
	}
	unit.Metrics.doneIteration(sim.Duration.Seconds())
	unit.resetCDs(sim)
}
