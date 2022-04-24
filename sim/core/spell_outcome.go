package core

import (
	"github.com/wowsims/tbc/sim/core/stats"
)

// This function should do 3 things:
//  1. Set the Outcome of the hit effect.
//  2. Update spell outcome metrics.
//  3. Modify the damage if necessary.
type OutcomeApplier func(sim *Simulation, spell *Spell, spellEffect *SpellEffect, damage *float64)

func OutcomeFuncAlwaysHit() OutcomeApplier {
	return func(_ *Simulation, spell *Spell, spellEffect *SpellEffect, _ *float64) {
		spellEffect.Outcome = OutcomeHit
		spell.SpellMetrics[spellEffect.Target.Index].Hits++
	}
}

// A tick always hits, but we don't count them as hits in the metrics.
func OutcomeFuncTick() OutcomeApplier {
	return func(_ *Simulation, _ *Spell, spellEffect *SpellEffect, _ *float64) {
		spellEffect.Outcome = OutcomeHit
	}
}

func OutcomeFuncMagicHitAndCrit(critMultiplier float64) OutcomeApplier {
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

func OutcomeFuncMagicHit() OutcomeApplier {
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

func OutcomeFuncMeleeWhite(critMultiplier float64) OutcomeApplier {
	return func(sim *Simulation, spell *Spell, spellEffect *SpellEffect, damage *float64) {
		character := spell.Character
		roll := sim.RandomFloat("White Hit Table")

		// Miss
		missChance := spellEffect.Target.MissChance - spellEffect.PhysicalHitChance(character)
		if character.AutoAttacks.IsDualWielding && !character.PseudoStats.DisableDWMissPenalty {
			missChance += 0.19
		}

		chance := MaxFloat(0, missChance)
		if roll < chance {
			spellEffect.Outcome = OutcomeMiss
			spell.SpellMetrics[spellEffect.Target.Index].Misses++
			*damage = 0
			return
		}

		// Dodge
		chance += MaxFloat(0, spellEffect.Target.Dodge-spellEffect.ExpertisePercentage(character)-character.PseudoStats.DodgeReduction)
		if roll < chance {
			spellEffect.Outcome = OutcomeDodge
			spell.SpellMetrics[spellEffect.Target.Index].Dodges++
			*damage = 0
			return
		}

		// Parry (if in front)
		// If the target is a mob and defense minus weapon skill is 11 or more:
		// ParryChance = 5% + (TargetLevel*5 - AttackerSkill) * 0.6%
		// If the target is a mob and defense minus weapon skill is 10 or less:
		// ParryChance = 5% + (TargetLevel*5 - AttackerSkill) * 0.1%

		// Block (if in front)
		// If the target is a mob:
		// BlockChance = MIN(5%, 5% + (TargetLevel*5 - AttackerSkill) * 0.1%)
		// If we actually implement blocks, ranged hits can be blocked.

		// Glance
		chance += spellEffect.Target.Glance
		if roll < chance {
			spellEffect.Outcome = OutcomeGlance
			spell.SpellMetrics[spellEffect.Target.Index].Glances++
			// TODO glancing blow damage reduction is actually a range ([65%, 85%] vs. 73)
			*damage *= 0.75
			return
		}

		// Crit
		chance += spellEffect.PhysicalCritChance(character, spell)
		if roll < chance {
			spellEffect.Outcome = OutcomeCrit
			spell.SpellMetrics[spellEffect.Target.Index].Crits++
			*damage *= critMultiplier
			return
		}

		// Hit
		spellEffect.Outcome = OutcomeHit
		spell.SpellMetrics[spellEffect.Target.Index].Hits++
	}
}

func OutcomeFuncMeleeSpecialHit() OutcomeApplier {
	return func(sim *Simulation, spell *Spell, spellEffect *SpellEffect, damage *float64) {
		character := spell.Character
		roll := sim.RandomFloat("White Hit Table")

		// Miss
		missChance := spellEffect.Target.MissChance - spellEffect.PhysicalHitChance(character)
		chance := MaxFloat(0, missChance)
		if roll < chance {
			spellEffect.Outcome = OutcomeMiss
			spell.SpellMetrics[spellEffect.Target.Index].Misses++
			*damage = 0
			return
		}

		// Dodge
		if !spell.SpellExtras.Matches(SpellExtrasCannotBeDodged) {
			chance += MaxFloat(0, spellEffect.Target.Dodge-spellEffect.ExpertisePercentage(character)-character.PseudoStats.DodgeReduction)
			if roll < chance {
				spellEffect.Outcome = OutcomeDodge
				spell.SpellMetrics[spellEffect.Target.Index].Dodges++
				*damage = 0
				return
			}
		}

		// Parry (if in front)
		// If the target is a mob and defense minus weapon skill is 11 or more:
		// ParryChance = 5% + (TargetLevel*5 - AttackerSkill) * 0.6%
		// If the target is a mob and defense minus weapon skill is 10 or less:
		// ParryChance = 5% + (TargetLevel*5 - AttackerSkill) * 0.1%

		// Hit
		spellEffect.Outcome = OutcomeHit
		spell.SpellMetrics[spellEffect.Target.Index].Hits++
	}
}

func OutcomeFuncMeleeSpecialHitAndCrit(critMultiplier float64) OutcomeApplier {
	return func(sim *Simulation, spell *Spell, spellEffect *SpellEffect, damage *float64) {
		character := spell.Character
		roll := sim.RandomFloat("White Hit Table")

		// Miss
		missChance := spellEffect.Target.MissChance - spellEffect.PhysicalHitChance(character)
		chance := MaxFloat(0, missChance)
		if roll < chance {
			spellEffect.Outcome = OutcomeMiss
			spell.SpellMetrics[spellEffect.Target.Index].Misses++
			*damage = 0
			return
		}

		// Dodge
		if !spell.SpellExtras.Matches(SpellExtrasCannotBeDodged) {
			chance += MaxFloat(0, spellEffect.Target.Dodge-spellEffect.ExpertisePercentage(character)-character.PseudoStats.DodgeReduction)
			if roll < chance {
				spellEffect.Outcome = OutcomeDodge
				spell.SpellMetrics[spellEffect.Target.Index].Dodges++
				*damage = 0
				return
			}
		}

		// Parry (if in front)
		// If the target is a mob and defense minus weapon skill is 11 or more:
		// ParryChance = 5% + (TargetLevel*5 - AttackerSkill) * 0.6%
		// If the target is a mob and defense minus weapon skill is 10 or less:
		// ParryChance = 5% + (TargetLevel*5 - AttackerSkill) * 0.1%

		// Block (if in front). Note that critical blocks are allowed for 2-roll hits.
		// If the target is a mob:
		// BlockChance = MIN(5%, 5% + (TargetLevel*5 - AttackerSkill) * 0.1%)
		// If we actually implement blocks, ranged hits can be blocked.

		// Crit (separate roll)
		if spellEffect.physicalCritRoll(sim, spell) {
			spellEffect.Outcome = OutcomeCrit
			spell.SpellMetrics[spellEffect.Target.Index].Crits++
			*damage *= critMultiplier
			return
		}

		// Hit
		spellEffect.Outcome = OutcomeHit
		spell.SpellMetrics[spellEffect.Target.Index].Hits++
	}
}

func OutcomeFuncMeleeSpecialNoBlockDodgeParry(critMultiplier float64) OutcomeApplier {
	return func(sim *Simulation, spell *Spell, spellEffect *SpellEffect, damage *float64) {
		character := spell.Character
		roll := sim.RandomFloat("White Hit Table")

		// Miss
		missChance := spellEffect.Target.MissChance - spellEffect.PhysicalHitChance(character)
		chance := MaxFloat(0, missChance)
		if roll < chance {
			spellEffect.Outcome = OutcomeMiss
			spell.SpellMetrics[spellEffect.Target.Index].Misses++
			*damage = 0
			return
		}

		// Crit (separate roll)
		if spellEffect.physicalCritRoll(sim, spell) {
			spellEffect.Outcome = OutcomeCrit
			spell.SpellMetrics[spellEffect.Target.Index].Crits++
			*damage *= critMultiplier
			return
		}

		// Hit
		spellEffect.Outcome = OutcomeHit
		spell.SpellMetrics[spellEffect.Target.Index].Hits++
	}
}

func OutcomeFuncMeleeSpecialCritOnly(critMultiplier float64) OutcomeApplier {
	return func(sim *Simulation, spell *Spell, spellEffect *SpellEffect, damage *float64) {
		if spellEffect.physicalCritRoll(sim, spell) {
			spellEffect.Outcome = OutcomeCrit
			spell.SpellMetrics[spellEffect.Target.Index].Crits++
			*damage *= critMultiplier
			return
		}

		// Hit
		spellEffect.Outcome = OutcomeHit
		spell.SpellMetrics[spellEffect.Target.Index].Hits++
	}
}

func OutcomeFuncRangedHit() OutcomeApplier {
	return func(sim *Simulation, spell *Spell, spellEffect *SpellEffect, damage *float64) {
		character := spell.Character
		roll := sim.RandomFloat("White Hit Table")

		// Miss
		missChance := spellEffect.Target.MissChance - spellEffect.PhysicalHitChance(character)
		chance := MaxFloat(0, missChance)
		if roll < chance {
			spellEffect.Outcome = OutcomeMiss
			spell.SpellMetrics[spellEffect.Target.Index].Misses++
			*damage = 0
			return
		}

		// Hit
		spellEffect.Outcome = OutcomeHit
		spell.SpellMetrics[spellEffect.Target.Index].Hits++
	}
}

func OutcomeFuncRangedHitAndCrit(critMultiplier float64) OutcomeApplier {
	return func(sim *Simulation, spell *Spell, spellEffect *SpellEffect, damage *float64) {
		character := spell.Character
		roll := sim.RandomFloat("White Hit Table")

		// Miss
		missChance := spellEffect.Target.MissChance - spellEffect.PhysicalHitChance(character)
		chance := MaxFloat(0, missChance)
		if roll < chance {
			spellEffect.Outcome = OutcomeMiss
			spell.SpellMetrics[spellEffect.Target.Index].Misses++
			*damage = 0
			return
		}

		// Block (if in front). Note that critical blocks are allowed for 2-roll hits.
		// If the target is a mob:
		// BlockChance = MIN(5%, 5% + (TargetLevel*5 - AttackerSkill) * 0.1%)
		// If we actually implement blocks, ranged hits can be blocked.

		// Crit (separate roll)
		if spellEffect.physicalCritRoll(sim, spell) {
			spellEffect.Outcome = OutcomeCrit
			spell.SpellMetrics[spellEffect.Target.Index].Crits++
			*damage *= critMultiplier
			return
		}

		// Hit
		spellEffect.Outcome = OutcomeHit
		spell.SpellMetrics[spellEffect.Target.Index].Hits++
	}
}

// Calculates a hit check using the stats from this spell.
func (spellEffect *SpellEffect) magicHitCheck(sim *Simulation, spell *Spell) bool {
	hit := 0.83 + (spell.Character.GetStat(stats.SpellHit)+spellEffect.BonusSpellHitRating)/(SpellHitRatingPerHitChance*100)
	hit = MinFloat(hit, 0.99) // can't get away from the 1% miss

	return sim.RandomFloat("Magical Hit Roll") < hit
}

func (spellEffect *SpellEffect) magicCritCheck(sim *Simulation, spell *Spell) bool {
	critChance := spellEffect.SpellCritChance(spell.Character, spell)
	return sim.RandomFloat("Magical Crit Roll") < critChance
}

func (spellEffect *SpellEffect) physicalCritRoll(sim *Simulation, spell *Spell) bool {
	return sim.RandomFloat("Physical Crit Roll") < spellEffect.PhysicalCritChance(spell.Character, spell)
}
