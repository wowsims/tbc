package core

import (
	"github.com/wowsims/tbc/sim/core/stats"
)

// Modifies damage based on Armor or Magic resistances, depending on the damage type.
func (spellEffect *SpellEffect) applyResistances(sim *Simulation, spell *Spell) {
	if spell.SpellExtras.Matches(SpellExtrasIgnoreResists) {
		return
	}

	if spell.SpellSchool.Matches(SpellSchoolPhysical) {
		// Physical resistance (armor).
		spellEffect.Damage *= spellEffect.Target.ArmorDamageReduction(spell.Unit.stats[stats.ArmorPenetration])
	} else if !spell.SpellExtras.Matches(SpellExtrasBinary) {
		// Magical resistance.
		target := spellEffect.Target

		resistanceRoll := sim.RandomFloat("Partial Resist")
		if sim.Log != nil {
			sim.Log("Resist thresholds: %0.04f, %0.04f, %0.04f",
				target.PartialResistRollThreshold00,
				target.PartialResistRollThreshold25,
				target.PartialResistRollThreshold50)
		}
		if resistanceRoll > target.PartialResistRollThreshold00 {
			// No partial resist.
		} else if resistanceRoll > target.PartialResistRollThreshold25 {
			spellEffect.Outcome |= OutcomePartial1_4
			spellEffect.Damage *= 0.75
		} else if resistanceRoll > target.PartialResistRollThreshold50 {
			spellEffect.Outcome |= OutcomePartial2_4
			spellEffect.Damage *= 0.5
		} else {
			spellEffect.Outcome |= OutcomePartial3_4
			spellEffect.Damage *= 0.25
		}
	}
}

// ArmorDamageReduction currently assumes a level 70 attacker
func (unit *Unit) ArmorDamageReduction(armorPen float64) float64 {
	// TODO: Cache this somehow so we dont have to recalculate every time.
	effectiveArmor := MaxFloat(0, unit.stats[stats.Armor]-armorPen)
	return 1 - (effectiveArmor / (effectiveArmor + 10557.5))
}

// All of the following calculations are based on this guide:
// https://royalgiraffe.github.io/resist-guide

func (unit *Unit) resistCoeff(school SpellSchool, attackerLevel int32, attackerSpellPen float64) float64 {
	resistanceCap := float64(unit.Level * 5)

	levelBasedResistance := 0.0
	if unit.Type == EnemyUnit {
		levelBasedResistance = LevelBasedNPCSpellResistancePerLevel * float64(MaxInt32(0, unit.Level-attackerLevel))
	}

	resistance := MaxFloat(0, unit.GetStat(school.ResistanceStat())-attackerSpellPen)
	if school == SpellSchoolHoly {
		resistance = 0
	}
	totalResistance := MinFloat(resistanceCap, resistance+levelBasedResistance)

	return totalResistance / resistanceCap
}

// Roll threshold for each type of partial resist.
func (unit *Unit) partialResistRollThresholds(school SpellSchool, attackerLevel int32, attackerSpellPen float64) (float64, float64, float64) {
	resistCoeff := unit.resistCoeff(school, attackerLevel, attackerSpellPen)

	// Based on the piecewise linear regression estimates at https://royalgiraffe.github.io/partial-resist-table.
	//partialResistChance00 := piecewiseLinear3(resistCoeff, 1, 0.24, 0.00, 0.00)
	partialResistChance25 := piecewiseLinear3(resistCoeff, 0, 0.55, 0.22, 0.04)
	partialResistChance50 := piecewiseLinear3(resistCoeff, 0, 0.18, 0.56, 0.16)
	partialResistChance75 := piecewiseLinear3(resistCoeff, 0, 0.03, 0.22, 0.80)

	return partialResistChance25 + partialResistChance50 + partialResistChance75,
		partialResistChance50 + partialResistChance75,
		partialResistChance75
}

// Interpolation for a 3-part piecewise linear function (which all the partial resist equations use).
func piecewiseLinear3(val float64, p0 float64, p1 float64, p2 float64, p3 float64) float64 {
	if val < 1.0/3.0 {
		return interpolate(val*3, p0, p1)
	} else if val < 2.0/3.0 {
		return interpolate((val-1.0/3.0)*3, p1, p2)
	} else {
		return interpolate((val-2.0/3.0)*3, p2, p3)
	}
}

func interpolate(val float64, p0 float64, p1 float64) float64 {
	return p0*(1-val) + p1*val
}
