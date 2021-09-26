package core

import (
	"strconv"
)

type Stats [StatLen]float64

type Stat byte

const (
	StatStrength Stat = iota
	StatAgility
	StatStamina
	StatIntellect
	StatSpirit
	StatSpellPower
	StatHealingPower
	StatArcaneSpellPower
	StatFireSpellPower
	StatFrostSpellPower
	StatHolySpellPower
	StatNatureSpellPower
	StatShadowSpellPower
	StatMP5
	StatSpellHit
	StatSpellCrit
	StatSpellHaste
	StatSpellPenetration
	StatAttackPower
	StatMeleeHit
	StatMeleeCrit
	StatMeleeHaste
	StatArmorPenetration
	StatExpertise
	StatMana
	StatEnergy
	StatRage
	StatArmor

	StatLen
)

func (s Stat) StatName() string {
	switch s {
	case StatStrength:
		return "Strength"
	case StatAgility:
		return "Agility"
	case StatStamina:
		return "Stamina"
	case StatIntellect:
		return "Intellect"
	case StatSpirit:
		return "Spirit"
	case StatSpellCrit:
		return "SpellCrit"
	case StatSpellHit:
		return "SpellHit"
	case StatHealingPower:
		return "HealingPower"
	case StatSpellPower:
		return "Spell Power"
	case StatSpellHaste:
		return "SpellHaste"
	case StatMP5:
		return "MP5"
	case StatSpellPenetration:
		return "StatSpellPenetration"
	case StatFireSpellPower:
		return "FireSpellPower"
	case StatNatureSpellPower:
		return "NatureSpellPower"
	case StatFrostSpellPower:
		return "FrostSpellPower"
	case StatShadowSpellPower:
		return "ShadowSpellPower"
	case StatHolySpellPower:
		return "HolySpellPower"
	case StatArcaneSpellPower:
		return "ArcaneSpellPower"
	case StatAttackPower:
		return "AttackPower"
	case StatMeleeHit:
		return "MeleeHit"
	case StatMeleeHaste:
		return "MeleeHaste"
	case StatMeleeCrit:
		return "MeleeCrit"
	case StatExpertise:
		return "Expertise"
	case StatArmorPenetration:
		return "ArmorPenetration"
	case StatMana:
		return "Mana"
	case StatEnergy:
		return "Energy"
	case StatRage:
		return "Rage"
	case StatArmor:
		return "Armor"
	}

	return "none"
}

// Print is debug print function
func (st Stats) Print() string {
	output := "{ "
	printed := false
	for k, v := range st {
		name := Stat(k).StatName()
		if name == "none" {
			continue
		}
		if v == 0 {
			continue
		}
		if printed {
			printed = false
			output += ",\n"
		}
		output += "\t"
		if v < 50 {
			printed = true
			output += "\"" + name + "\": " + strconv.FormatFloat(v, 'f', 3, 64)
		} else {
			printed = true
			output += "\"" + name + "\": " + strconv.FormatFloat(v, 'f', 0, 64)
		}
	}
	output += " }"
	return output
}

// CalculatedTotal will add Mana and Crit from Int and return the new stats.
//   TODO: These numbers might change from class to class and so we might need to make this per-class.
func (s Stats) CalculatedTotal() Stats {
	stats := s
	// Add crit/mana from int
	stats[StatSpellCrit] += (stats[StatIntellect] / 78.1) * 22.08
	stats[StatMana] += stats[StatIntellect] * 15
	return stats
}

// CalculateTotalStats will take a set of equipment and options and add all stats/buffs/etc together
func CalculateTotalStats(race RaceBonusType, e Equipment, c Consumes) Stats {
	stats := BaseStats(race)
	gearStats := e.Stats()
	for i := range stats {
		stats[i] += gearStats[i]
	}
	stats = c.AddStats(stats)
	return stats
}

// TODO: This probably should be moved into each class because they all have different base stats.
func BaseStats(race RaceBonusType) Stats {
	stats := Stats{
		StatIntellect: 104,    // Base int for troll,
		StatMana:      2678,   // level 70 shaman
		StatSpirit:    135,    // lvl 70 shaman
		StatSpellCrit: 48.576, // base crit for 70 sham
	}
	// TODO: Find race differences.
	switch race {
	case RaceBonusTypeOrc:
	}
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
