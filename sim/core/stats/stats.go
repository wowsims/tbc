package stats

import (
	"fmt"
	"strings"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
)

type Stats [Len]float64

type Stat byte

// Use internal representation instead of proto.Stat so we can add functions
// and use 'byte' as the data type.
//
// This needs to stay synced with proto.Stat.
const (
	Strength Stat = iota
	Agility
	Stamina
	Intellect
	Spirit
	SpellPower
	HealingPower
	ArcaneSpellPower
	FireSpellPower
	FrostSpellPower
	HolySpellPower
	NatureSpellPower
	ShadowSpellPower
	MP5
	SpellHit
	SpellCrit
	SpellHaste
	SpellPenetration
	AttackPower
	MeleeHit
	MeleeCrit
	MeleeHaste
	ArmorPenetration
	Expertise
	Mana
	Energy
	Rage
	Armor

	Len
)

func ProtoArrayToStatsList(protoStats []proto.Stat) []Stat {
	stats := make([]Stat, len(protoStats))
	for i, v := range protoStats {
		stats[i] = Stat(v)
	}
	return stats
}

func (s Stat) StatName() string {
	switch s {
	case Strength:
		return "Strength"
	case Agility:
		return "Agility"
	case Stamina:
		return "Stamina"
	case Intellect:
		return "Intellect"
	case Spirit:
		return "Spirit"
	case SpellCrit:
		return "SpellCrit"
	case SpellHit:
		return "SpellHit"
	case HealingPower:
		return "HealingPower"
	case SpellPower:
		return "SpellPower"
	case SpellHaste:
		return "SpellHaste"
	case MP5:
		return "MP5"
	case SpellPenetration:
		return "StatSpellPenetration"
	case FireSpellPower:
		return "FireSpellPower"
	case NatureSpellPower:
		return "NatureSpellPower"
	case FrostSpellPower:
		return "FrostSpellPower"
	case ShadowSpellPower:
		return "ShadowSpellPower"
	case HolySpellPower:
		return "HolySpellPower"
	case ArcaneSpellPower:
		return "ArcaneSpellPower"
	case AttackPower:
		return "AttackPower"
	case MeleeHit:
		return "MeleeHit"
	case MeleeHaste:
		return "MeleeHaste"
	case MeleeCrit:
		return "MeleeCrit"
	case Expertise:
		return "Expertise"
	case ArmorPenetration:
		return "ArmorPenetration"
	case Mana:
		return "Mana"
	case Energy:
		return "Energy"
	case Rage:
		return "Rage"
	case Armor:
		return "Armor"
	}

	return "none"
}

func FromFloatArray(values []float64) Stats {
	stats := Stats{}
	for i, v := range values {
		stats[i] = v
	}
	return stats
}

// Adds two Stats together, returning the new Stats.
func (this Stats) Add(other Stats) Stats {
	newStats := Stats{}

	for i, thisStat := range this {
		newStats[i] = thisStat + other[i]
	}

	return newStats
}

// Subtracts another Stats from this one, returning the new Stats.
func (stats Stats) Subtract(other Stats) Stats {
	newStats := Stats{}

	for k, v := range stats {
		newStats[k] = v - other[k]
	}

	return newStats
}

// Multiplies two Stats together by multiplying the values of corresponding
// stats, like a dot product operation.
func (stats Stats) DotProduct(other Stats) Stats {
	newStats := Stats{}

	for k, v := range stats {
		newStats[k] = v * other[k]
	}

	return newStats
}

func (this Stats) Equals(other Stats) bool {
	for i := range this {
		if this[i] != other[i] {
			return false
		}
	}

	return true
}

func (this Stats) EqualsWithTolerance(other Stats, tolerance float64) bool {
	for i := range this {
		if this[i] < other[i]-tolerance || this[i] > other[i]+tolerance {
			return false
		}
	}

	return true
}

func (this Stats) String() string {
	var sb strings.Builder
	sb.WriteString("\n{\n")

	for statIdx, statValue := range this {
		name := Stat(statIdx).StatName()
		if name == "none" || statValue == 0 {
			continue
		}

		fmt.Fprintf(&sb, "\t%s: %0.3f,\n", name, statValue)
	}

	sb.WriteString("\n}")
	return sb.String()
}

// Given the current values for source and mod stats, should return the new
// value for the mod stat.
type StatModifier func(sourceValue float64, modValue float64) float64

// Represents a dependency between two stats, whereby the value of one stat
// modifies the value of the other.
//
// For example, many casters have a talent to increase their spell power by
// a percentage of their intellect.
type StatDependency struct {
	// The stat which will be used to control the amount of increase.
	SourceStat Stat

	// The stat which will be modified, depending on the value of SourceStat.
	ModifiedStat Stat

	// Applies the stat modification.
	Modifier StatModifier
}

type StatDependencyManager struct {
	// Stat dependencies for each stat.
	// First dimension is the modified stat. For each modified stat, stores a list of
	// dependencies for that stat.
	deps [Len][]StatDependency

	// Whether Finalize() has been called.
	finalized bool

	// Dependencies being managed, sorted so that their modifiers can be applied
	// in-order without any issues.
	sortedDeps []StatDependency
}

func (sdm *StatDependencyManager) AddStatDependency(dep StatDependency) {
	if sdm.finalized {
		panic("Stat dependencies may not be added once finalized!")
	}

	sdm.deps[dep.ModifiedStat] = append(sdm.deps[dep.ModifiedStat], dep)
}

// Populates sortedDeps. Panics if there are any dependency cycles.
// TODO: Figure out if we need to separate additive / multiplicative dependencies.
func (sdm *StatDependencyManager) Finalize() {
	if sdm.finalized {
		return
	}
	sdm.finalized = true

	sdm.sortedDeps = []StatDependency{}

	// Set of stats we're done processing.
	processedStats := map[Stat]struct{}{}

	for len(processedStats) < int(Len) {
		numNewlyProcessed := 0
		for i := 0; i < int(Len); i++ {
			stat := Stat(i)

			if _, alreadyProcessed := processedStats[stat]; alreadyProcessed {
				continue
			}

			// If all deps for this stat have been processed or are the same stat, we can process it.
			allDepsProcessed := true
			for _, dep := range sdm.deps[stat] {
				_, depAlreadyProcessed := processedStats[dep.SourceStat]

				if !depAlreadyProcessed && dep.SourceStat != stat {
					allDepsProcessed = false
				}
			}
			if !allDepsProcessed {
				continue
			}

			// Process this stat by adding its deps to sortedDeps.

			// Add deps from other stats first.
			for _, dep := range sdm.deps[stat] {
				if dep.SourceStat != stat {
					sdm.sortedDeps = append(sdm.sortedDeps, dep)
				}
			}

			// Now add deps from the same stat.
			for _, dep := range sdm.deps[stat] {
				if dep.SourceStat == stat {
					sdm.sortedDeps = append(sdm.sortedDeps, dep)
				}
			}

			// Mark this stat as processed.
			processedStats[stat] = struct{}{}
			numNewlyProcessed++
		}

		// If we couldn't process any new stats but there are still stats left,
		// there must be a circular dependency.
		if numNewlyProcessed == 0 {
			panic("Circular stat dependency detected")
		}
	}
}

// Applies all stat dependencies and returns the new Stats.
func (sdm *StatDependencyManager) ApplyStatDependencies(stats Stats) Stats {
	newStats := stats

	for _, dep := range sdm.sortedDeps {
		newStats[dep.ModifiedStat] = dep.Modifier(newStats[dep.SourceStat], newStats[dep.ModifiedStat])
	}

	return newStats
}

type PseudoStats struct {
	CastSpeedMultiplier   float64
	AttackSpeedMultiplier float64 // not used yet

	FiveSecondRuleRefreshTime time.Duration // last time a spell was cast
	SpiritRegenRateCasting    float64       // percentage of spirit regen allowed during casting

	// Both of these are currently only used for innervate.
	ForceFullSpiritRegen  bool    // If set, automatically uses full spirit regen regardless of FSR refresh time.
	SpiritRegenMultiplier float64 // Multiplier on spirit portion of mana regen.
}

// TODO: more stat calculations

// INT

// Warlocks receive 1% Spell Critical Strike chance for every 81.9 points of intellect.
// Druids receive 1% Spell Critical Strike chance for every 79.4 points of intellect.
// Shamans receive 1% Spell Critical Strike chance for every 78.1 points of intellect.
// Mages receive 1% Spell Critical Strike chance for every 81 points of intellect.
// Priests receive 1% Spell Critical Strike chance for every 80 points of intellect.
// Paladins receive 1% Spell Critical Strike chance for every 79.4 points of intellect.

// AGI

// Rogues, Hunters, and Warriors gain 1 ranged Attack Power per point of Agility.
// Druids in Cat Form, Hunters and Rogues gain 1 melee Attack Power per point of Agility.
// You gain 2 Armor for every point of Agility.

// You gain Critical Strike Chance at varying rates, depending on your class:
// 	Paladins, Druids, and Shamans receive 1% Critical Strike Chance for every 25 points of Agility.
// 	Rogues and Hunters receive 1% Critical Strike Chance for every 40 points of Agility.
// 	Warriors receive 1% Critical Strike Chance for every 33 points of Agility.

// STR

// Feral Druids receive 2 melee Attack Power per point of Strength.
// Protection and Retribution Paladins receive 1 melee Attack Power per point of Strength.
// Rogues receive 1 melee Attack Power per point of Strength.
// Enhancement Shaman receive 2 melee Attack Power per point of Strength.
