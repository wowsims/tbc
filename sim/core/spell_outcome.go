package core

import (
	"github.com/wowsims/tbc/sim/core/stats"
)

// This function should do 3 things:
//  1. Set the Outcome of the hit effect.
//  2. Update spell outcome metrics.
//  3. Modify the damage if necessary.
type OutcomeApplier func(sim *Simulation, spell *Spell, spellEffect *SpellEffect, damage *float64)

func (unit *Unit) OutcomeFuncAlwaysHit() OutcomeApplier {
	return func(_ *Simulation, spell *Spell, spellEffect *SpellEffect, _ *float64) {
		spellEffect.Outcome = OutcomeHit
		spell.SpellMetrics[spellEffect.Target.Index].Hits++
	}
}

// A tick always hits, but we don't count them as hits in the metrics.
func (unit *Unit) OutcomeFuncTick() OutcomeApplier {
	return func(_ *Simulation, _ *Spell, spellEffect *SpellEffect, _ *float64) {
		spellEffect.Outcome = OutcomeHit
	}
}

func (unit *Unit) OutcomeFuncMagicHitAndCrit(critMultiplier float64) OutcomeApplier {
	return func(sim *Simulation, spell *Spell, spellEffect *SpellEffect, damage *float64) {
		if spellEffect.magicHitCheck(sim, spell) {
			if spellEffect.magicCritCheck(sim, spell) {
				spellEffect.Outcome = OutcomeCrit
				spell.SpellMetrics[spellEffect.Target.Index].Crits++
				*damage *= critMultiplier
			} else {
				spellEffect.Outcome = OutcomeHit
				spell.SpellMetrics[spellEffect.Target.Index].Hits++
			}
		} else {
			spellEffect.Outcome = OutcomeMiss
			spell.SpellMetrics[spellEffect.Target.Index].Misses++
			*damage = 0
		}
	}
}

func (unit *Unit) OutcomeFuncMagicHit() OutcomeApplier {
	return func(sim *Simulation, spell *Spell, spellEffect *SpellEffect, damage *float64) {
		if spellEffect.magicHitCheck(sim, spell) {
			spellEffect.Outcome = OutcomeHit
			spell.SpellMetrics[spellEffect.Target.Index].Hits++
		} else {
			spellEffect.Outcome = OutcomeMiss
			spell.SpellMetrics[spellEffect.Target.Index].Misses++
			*damage = 0
		}
	}
}

func (unit *Unit) OutcomeFuncMeleeWhite(critMultiplier float64) OutcomeApplier {
	if unit.PseudoStats.InFrontOfTarget {
		return func(sim *Simulation, spell *Spell, spellEffect *SpellEffect, damage *float64) {
			unit := spell.Unit
			roll := sim.RandomFloat("White Hit Table")
			chance := 0.0

			if !spellEffect.applyAttackTableMiss(spell, unit, roll, &chance, damage) &&
				!spellEffect.applyAttackTableDodge(spell, unit, roll, &chance, damage) &&
				!spellEffect.applyAttackTableParry(spell, unit, roll, &chance, damage) &&
				!spellEffect.applyAttackTableGlance(spell, unit, roll, &chance, damage) &&
				!spellEffect.applyAttackTableBlock(spell, unit, roll, &chance, damage) &&
				!spellEffect.applyAttackTableCrit(spell, unit, roll, critMultiplier, &chance, damage) {
				spellEffect.applyAttackTableHit(spell)
			}
		}
	} else {
		return func(sim *Simulation, spell *Spell, spellEffect *SpellEffect, damage *float64) {
			unit := spell.Unit
			roll := sim.RandomFloat("White Hit Table")
			chance := 0.0

			if !spellEffect.applyAttackTableMiss(spell, unit, roll, &chance, damage) &&
				!spellEffect.applyAttackTableDodge(spell, unit, roll, &chance, damage) &&
				!spellEffect.applyAttackTableGlance(spell, unit, roll, &chance, damage) &&
				!spellEffect.applyAttackTableCrit(spell, unit, roll, critMultiplier, &chance, damage) {
				spellEffect.applyAttackTableHit(spell)
			}
		}
	}
}

func (unit *Unit) OutcomeFuncMeleeSpecialHit() OutcomeApplier {
	if unit.PseudoStats.InFrontOfTarget {
		return func(sim *Simulation, spell *Spell, spellEffect *SpellEffect, damage *float64) {
			unit := spell.Unit
			roll := sim.RandomFloat("White Hit Table")
			chance := 0.0

			if !spellEffect.applyAttackTableMissNoDWPenalty(spell, unit, roll, &chance, damage) &&
				(spell.SpellExtras.Matches(SpellExtrasCannotBeDodged) || !spellEffect.applyAttackTableDodge(spell, unit, roll, &chance, damage)) &&
				!spellEffect.applyAttackTableParry(spell, unit, roll, &chance, damage) {
				spellEffect.applyAttackTableHit(spell)
			}
		}
	} else {
		return func(sim *Simulation, spell *Spell, spellEffect *SpellEffect, damage *float64) {
			unit := spell.Unit
			roll := sim.RandomFloat("White Hit Table")
			chance := 0.0

			if !spellEffect.applyAttackTableMissNoDWPenalty(spell, unit, roll, &chance, damage) &&
				(spell.SpellExtras.Matches(SpellExtrasCannotBeDodged) || !spellEffect.applyAttackTableDodge(spell, unit, roll, &chance, damage)) {
				spellEffect.applyAttackTableHit(spell)
			}
		}
	}
}

func (unit *Unit) OutcomeFuncMeleeSpecialHitAndCrit(critMultiplier float64) OutcomeApplier {
	if unit.PseudoStats.InFrontOfTarget {
		return func(sim *Simulation, spell *Spell, spellEffect *SpellEffect, damage *float64) {
			unit := spell.Unit
			roll := sim.RandomFloat("White Hit Table")
			chance := 0.0

			if !spellEffect.applyAttackTableMissNoDWPenalty(spell, unit, roll, &chance, damage) &&
				(spell.SpellExtras.Matches(SpellExtrasCannotBeDodged) || !spellEffect.applyAttackTableDodge(spell, unit, roll, &chance, damage)) &&
				!spellEffect.applyAttackTableParry(spell, unit, roll, &chance, damage) {
				if !spellEffect.applyAttackTableCritSeparateRoll(sim, spell, critMultiplier, damage) {
					if !spellEffect.applyAttackTableBlock(spell, unit, roll, &chance, damage) {
						spellEffect.applyAttackTableHit(spell)
					}
				} else {
					spellEffect.applyAttackTableBlock(spell, unit, roll, &chance, damage)
				}
			}
		}
	} else {
		return func(sim *Simulation, spell *Spell, spellEffect *SpellEffect, damage *float64) {
			unit := spell.Unit
			roll := sim.RandomFloat("White Hit Table")
			chance := 0.0

			if !spellEffect.applyAttackTableMissNoDWPenalty(spell, unit, roll, &chance, damage) &&
				(spell.SpellExtras.Matches(SpellExtrasCannotBeDodged) || !spellEffect.applyAttackTableDodge(spell, unit, roll, &chance, damage)) &&
				!spellEffect.applyAttackTableCritSeparateRoll(sim, spell, critMultiplier, damage) {
				spellEffect.applyAttackTableHit(spell)
			}
		}
	}
}

func (unit *Unit) OutcomeFuncMeleeSpecialNoBlockDodgeParry(critMultiplier float64) OutcomeApplier {
	return func(sim *Simulation, spell *Spell, spellEffect *SpellEffect, damage *float64) {
		unit := spell.Unit
		roll := sim.RandomFloat("White Hit Table")
		chance := 0.0

		if !spellEffect.applyAttackTableMissNoDWPenalty(spell, unit, roll, &chance, damage) &&
			!spellEffect.applyAttackTableCritSeparateRoll(sim, spell, critMultiplier, damage) {
			spellEffect.applyAttackTableHit(spell)
		}
	}
}

func (unit *Unit) OutcomeFuncMeleeSpecialCritOnly(critMultiplier float64) OutcomeApplier {
	return func(sim *Simulation, spell *Spell, spellEffect *SpellEffect, damage *float64) {
		if !spellEffect.applyAttackTableCritSeparateRoll(sim, spell, critMultiplier, damage) {
			spellEffect.applyAttackTableHit(spell)
		}
	}
}

func (unit *Unit) OutcomeFuncRangedHit() OutcomeApplier {
	return func(sim *Simulation, spell *Spell, spellEffect *SpellEffect, damage *float64) {
		unit := spell.Unit
		roll := sim.RandomFloat("White Hit Table")
		chance := 0.0

		if !spellEffect.applyAttackTableMissNoDWPenalty(spell, unit, roll, &chance, damage) {
			spellEffect.applyAttackTableHit(spell)
		}
	}
}

func (unit *Unit) OutcomeFuncRangedHitAndCrit(critMultiplier float64) OutcomeApplier {
	if unit.PseudoStats.InFrontOfTarget {
		return func(sim *Simulation, spell *Spell, spellEffect *SpellEffect, damage *float64) {
			unit := spell.Unit
			roll := sim.RandomFloat("White Hit Table")
			chance := 0.0

			if !spellEffect.applyAttackTableMissNoDWPenalty(spell, unit, roll, &chance, damage) {
				if !spellEffect.applyAttackTableCritSeparateRoll(sim, spell, critMultiplier, damage) {
					if !spellEffect.applyAttackTableBlock(spell, unit, roll, &chance, damage) {
						spellEffect.applyAttackTableHit(spell)
					}
				} else {
					spellEffect.applyAttackTableBlock(spell, unit, roll, &chance, damage)
				}
			}
		}
	} else {
		return func(sim *Simulation, spell *Spell, spellEffect *SpellEffect, damage *float64) {
			unit := spell.Unit
			roll := sim.RandomFloat("White Hit Table")
			chance := 0.0

			if !spellEffect.applyAttackTableMissNoDWPenalty(spell, unit, roll, &chance, damage) &&
				!spellEffect.applyAttackTableCritSeparateRoll(sim, spell, critMultiplier, damage) {
				spellEffect.applyAttackTableHit(spell)
			}
		}
	}
}

// Calculates a hit check using the stats from this spell.
func (spellEffect *SpellEffect) magicHitCheck(sim *Simulation, spell *Spell) bool {
	hit := spellEffect.Target.BaseSpellMissChance - (spell.Unit.GetStat(stats.SpellHit)+spellEffect.BonusSpellHitRating)/(SpellHitRatingPerHitChance*100)
	hit = MaxFloat(hit, 0.01) // can't get away from the 1% miss

	return sim.RandomFloat("Magical Hit Roll") > hit
}

func (spellEffect *SpellEffect) magicCritCheck(sim *Simulation, spell *Spell) bool {
	critChance := spellEffect.SpellCritChance(spell.Unit, spell)
	return sim.RandomFloat("Magical Crit Roll") < critChance
}

func (spellEffect *SpellEffect) physicalCritRoll(sim *Simulation, spell *Spell) bool {
	return sim.RandomFloat("Physical Crit Roll") < spellEffect.PhysicalCritChance(spell.Unit, spell)
}

func (spellEffect *SpellEffect) applyAttackTableMiss(spell *Spell, unit *Unit, roll float64, chance *float64, damage *float64) bool {
	missChance := spellEffect.Target.BaseMissChance - spellEffect.PhysicalHitChance(unit)
	if unit.AutoAttacks.IsDualWielding && !unit.PseudoStats.DisableDWMissPenalty {
		missChance += 0.19
	}
	*chance = MaxFloat(0, missChance)

	if roll < *chance {
		spellEffect.Outcome = OutcomeMiss
		spell.SpellMetrics[spellEffect.Target.Index].Misses++
		*damage = 0
		return true
	}
	return false
}

func (spellEffect *SpellEffect) applyAttackTableMissNoDWPenalty(spell *Spell, unit *Unit, roll float64, chance *float64, damage *float64) bool {
	missChance := spellEffect.Target.BaseMissChance - spellEffect.PhysicalHitChance(unit)
	*chance = MaxFloat(0, missChance)

	if roll < *chance {
		spellEffect.Outcome = OutcomeMiss
		spell.SpellMetrics[spellEffect.Target.Index].Misses++
		*damage = 0
		return true
	}
	return false
}

func (spellEffect *SpellEffect) applyAttackTableBlock(spell *Spell, unit *Unit, roll float64, chance *float64, damage *float64) bool {
	*chance += spellEffect.Target.BaseBlockChance

	if roll < *chance {
		spellEffect.Outcome |= OutcomeBlock
		spell.SpellMetrics[spellEffect.Target.Index].Blocks++
		*damage -= spellEffect.Target.GetStat(stats.BlockValue)
		return true
	}
	return false
}

func (spellEffect *SpellEffect) applyAttackTableDodge(spell *Spell, unit *Unit, roll float64, chance *float64, damage *float64) bool {
	*chance += MaxFloat(0, spellEffect.Target.BaseDodgeChance-spellEffect.ExpertisePercentage(unit)-unit.PseudoStats.DodgeReduction)

	if roll < *chance {
		spellEffect.Outcome = OutcomeDodge
		spell.SpellMetrics[spellEffect.Target.Index].Dodges++
		*damage = 0
		return true
	}
	return false
}

func (spellEffect *SpellEffect) applyAttackTableParry(spell *Spell, unit *Unit, roll float64, chance *float64, damage *float64) bool {
	*chance += MaxFloat(0, spellEffect.Target.BaseParryChance-spellEffect.ExpertisePercentage(unit))

	if roll < *chance {
		spellEffect.Outcome = OutcomeParry
		spell.SpellMetrics[spellEffect.Target.Index].Parries++
		*damage = 0
		return true
	}
	return false
}

func (spellEffect *SpellEffect) applyAttackTableGlance(spell *Spell, unit *Unit, roll float64, chance *float64, damage *float64) bool {
	*chance += spellEffect.Target.BaseGlanceChance

	if roll < *chance {
		spellEffect.Outcome = OutcomeGlance
		spell.SpellMetrics[spellEffect.Target.Index].Glances++
		// TODO glancing blow damage reduction is actually a range ([65%, 85%] vs. 73)
		*damage *= spellEffect.Target.GlanceMultiplier
		return true
	}
	return false
}

func (spellEffect *SpellEffect) applyAttackTableCrit(spell *Spell, unit *Unit, roll float64, critMultiplier float64, chance *float64, damage *float64) bool {
	*chance += spellEffect.PhysicalCritChance(unit, spell)

	if roll < *chance {
		spellEffect.Outcome = OutcomeCrit
		spell.SpellMetrics[spellEffect.Target.Index].Crits++
		*damage *= critMultiplier
		return true
	}
	return false
}

func (spellEffect *SpellEffect) applyAttackTableCritSeparateRoll(sim *Simulation, spell *Spell, critMultiplier float64, damage *float64) bool {
	if spellEffect.physicalCritRoll(sim, spell) {
		spellEffect.Outcome = OutcomeCrit
		spell.SpellMetrics[spellEffect.Target.Index].Crits++
		*damage *= critMultiplier
		return true
	}
	return false
}

func (spellEffect *SpellEffect) applyAttackTableHit(spell *Spell) {
	spellEffect.Outcome = OutcomeHit
	spell.SpellMetrics[spellEffect.Target.Index].Hits++
}
