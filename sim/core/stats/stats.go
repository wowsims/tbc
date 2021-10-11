package stats

import (
	"fmt"
	"strings"
)

type Stats [Len]float64

type Stat byte

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

// Adds two Stats together, returning the new Stats.
func (this Stats) Add(other Stats) Stats {
	newStats := Stats{}

	for i, thisStat := range this {
		newStats[i] = thisStat + other[i]
	}

	return newStats
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

// CalculatedTotal will add Mana and Crit from Int and return the new stats.
//   TODO: These numbers might change from class to class and so we might need to make this per-class.
func (s Stats) CalculatedTotal() Stats {
	stats := s
	// Add crit/mana from int
	stats[SpellCrit] += (stats[Intellect] / 78.1) * 22.08
	stats[Mana] += stats[Intellect] * 15
	return stats
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
