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
	StatSpellCrit
	StatSpellHit
	StatHealing
	StatSpellPower
	StatSpellHaste
	StatMP5
	StatSpellpen
	StatFireSpellPower
	StatNatureSpellPower
	StatFrostSpellPower
	StatShadowSpellPower
	StatHolySpellPower
	StatArcaneSpellPower
	StatAttackPower
	StatMeleeHit
	StatMeleeHaste
	StatMeleeCrit
	StatExpertise
	StatArmorPenetration
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
	case StatHealing:
		return "Healing"
	case StatSpellPower:
		return "SpellPower"
	case StatSpellHaste:
		return "SpellHaste"
	case StatMP5:
		return "MP5"
	case StatSpellpen:
		return "Spellpen"
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
func (s Stats) CalculatedTotal() Stats {
	stats := s

	// Add crit/mana from int
	stats[StatSpellCrit] += (stats[StatInt] / 80) * 22.08
	stats[StatMana] += stats[StatInt] * 15
	return stats
}

// CalculateTotalStats will take a set of equipment and options and add all stats/buffs/etc together
func CalculateTotalStats(o Options, e Equipment) Stats {
	gearStats := e.Stats()
	stats := BaseStats(o.Buffs.Race)
	for i := range stats {
		stats[i] += gearStats[i]
	}

	stats = o.Talents.AddStats(o.Buffs.AddStats(o.Consumes.AddStats(o.Totems.AddStats(stats))))

	if o.Buffs.BlessingOfKings {
		stats[StatInt] *= 1.1 // blessing of kings
	}
	if o.Buffs.ImprovedDivineSpirit {
		stats[StatSpellDmg] += stats[StatSpirit] * 0.1
	}

	stats = stats.CalculatedTotal()

	// Add stat increases from talents
	stats[StatMP5] += stats[StatInt] * (0.02 * float64(o.Talents.UnrelentingStormP)

	return stats
}

func BaseStats(race RaceBonusType) Stats {
	stats := Stats{
		StatInt:       104,    // Base int for troll,
		StatMana:      2678,   // level 70 shaman
		StatSpirit:    135,    // lvl 70 shaman
		StatSpellCrit: 48.576, // base crit for 70 sham
	}
	// TODO: Find race int differences.
	switch race {
	case RaceBonusOrc:

	}
	return stats
}
