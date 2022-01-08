package core

import (
	"fmt"
	"strings"
	"time"

	"github.com/wowsims/tbc/sim/core/stats"
)

// Callback for after a spell hits the target, before damage has been calculated.
// Use it to modify the spell damage or results.
type OnBeforeSpellHit func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect)

// Callback for after a spell hits the target and after damage is calculated. Use it for proc effects
// or anything that comes from the final result of the spell.
type OnSpellHit func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect)

// Callback for after a spell is fully resisted on a target.
type OnSpellMiss func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect)

// OnBeforePeriodicDamage is called when dots tick, before damage is calculated.
// Use it to modify the spell damage or results.
type OnBeforePeriodicDamage func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect, tickDamage *float64)

// OnPeriodicDamage is called when dots tick, after damage is calculated. Use it for proc effects
// or anything that comes from the final result of a tick.
type OnPeriodicDamage func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect, tickDamage float64)

// A Spell is a type of cast that can hit/miss using spell stats, and has a spell school.
type SpellCast struct {
	// Embedded Cast
	Cast

	// Results from the spell cast. Spell casts can have multiple effects (e.g.
	// Chain Lightning, Moonfire) so these are totals from all the effects.
	Hits               int32
	Misses             int32
	Crits              int32
	PartialResists_1_4 int32   // 1/4 of the spell was resisted
	PartialResists_2_4 int32   // 2/4 of the spell was resisted
	PartialResists_3_4 int32   // 3/4 of the spell was resisted
	TotalDamage        float64 // Damage done by this cast.
}

type SpellEffect struct {
	// Target of the spell.
	Target *Target

	// Bonus stats to be added to the spell.
	BonusSpellHitRating  float64
	BonusSpellPower      float64
	BonusSpellCritRating float64

	// Additional multiplier that is always applied.
	DamageMultiplier float64

	// applies fixed % increases to damage at cast time.
	//  Only use multipliers that don't change for the lifetime of the sim.
	//  This should probably only be mutated in a template and not changed in auras.
	StaticDamageMultiplier float64

	// Skips the hit check, i.e. this effect will always hit.
	// This is generally used only for proc effects, like Mage Ignite.
	IgnoreHitCheck bool

	// Callbacks for providing additional custom behavior.
	OnSpellHit  OnSpellHit
	OnSpellMiss OnSpellMiss

	// Results

	Hit  bool // True = hit, False = resisted
	Crit bool // Whether this cast was a critical strike.

	PartialResist_1_4 bool // 1/4 of the spell was resisted
	PartialResist_2_4 bool // 2/4 of the spell was resisted
	PartialResist_3_4 bool // 3/4 of the spell was resisted

	Damage float64 // Damage done by this cast.
}

func (spellEffect *SpellEffect) beforeCalculations(sim *Simulation, spellCast *SpellCast) {
	spellCast.Character.OnBeforeSpellHit(sim, spellCast, spellEffect)
	spellEffect.Target.OnBeforeSpellHit(sim, spellCast, spellEffect)

	spellEffect.Hit = spellEffect.IgnoreHitCheck || spellEffect.hitCheck(sim, spellCast)
}

func (spellEffect *SpellEffect) triggerSpellProcs(sim *Simulation, spellCast *SpellCast) {
	if spellEffect.Hit {
		if spellEffect.OnSpellHit != nil {
			spellEffect.OnSpellHit(sim, spellCast, spellEffect)
		}
		spellCast.Character.OnSpellHit(sim, spellCast, spellEffect)
		spellEffect.Target.OnSpellHit(sim, spellCast, spellEffect)
	} else {
		if spellEffect.OnSpellMiss != nil {
			spellEffect.OnSpellMiss(sim, spellCast, spellEffect)
		}
		spellCast.Character.OnSpellMiss(sim, spellCast, spellEffect)
		spellEffect.Target.OnSpellMiss(sim, spellCast, spellEffect)
	}
}

func (spellEffect *SpellEffect) afterCalculations(sim *Simulation, spellCast *SpellCast) {
	if sim.Log != nil && !spellEffect.IgnoreHitCheck {
		spellCast.Character.Log(sim, "%s %s.", spellCast.ActionID, spellEffect)
	}

	spellEffect.triggerSpellProcs(sim, spellCast)
}

// Calculates a hit check using the stats from this spell.
func (spellEffect *SpellEffect) hitCheck(sim *Simulation, spellCast *SpellCast) bool {
	hit := 0.83 + (spellCast.Character.GetStat(stats.SpellHit)+spellEffect.BonusSpellHitRating)/(SpellHitRatingPerHitChance*100)
	hit = MinFloat(hit, 0.99) // can't get away from the 1% miss

	return sim.RandomFloat("SpellCast Hit") < hit
}

// Calculates a crit check using the stats from this spell.
func (spellEffect *SpellEffect) critCheck(sim *Simulation, spellCast *SpellCast) bool {
	critChance := (spellCast.Character.GetStat(stats.SpellCrit) + spellCast.BonusCritRating + spellEffect.BonusSpellCritRating) / (SpellCritRatingPerCritChance * 100)
	return sim.RandomFloat("DirectSpell Crit") < critChance
}

func (spellEffect *SpellEffect) applyResultsToCast(spellCast *SpellCast) {
	if spellEffect.Hit {
		spellCast.Hits++
		if spellEffect.Crit {
			spellCast.Crits++
		}
	} else {
		spellCast.Misses++
	}

	if spellEffect.PartialResist_1_4 {
		spellCast.PartialResists_1_4++
	} else if spellEffect.PartialResist_2_4 {
		spellCast.PartialResists_2_4++
	} else if spellEffect.PartialResist_3_4 {
		spellCast.PartialResists_3_4++
	}

	spellCast.TotalDamage += spellEffect.Damage
}

// Only applies the results from the ticks, not the initial dot application.
// This is used for spells like Arcane Missiles or Magma Totem, where the initial
// cast always hits and we just care about the dot itself.
func (spellEffect *SpellEffect) applyDotTickResultsToCast(spellCast *SpellCast, hit bool, crit bool, resistMultiplier float64, damage float64) {
	spellEffect.Hit = hit
	spellEffect.Crit = crit
	if hit {
		spellCast.Hits++
		if crit {
			spellCast.Crits++
		}
	} else {
		spellCast.Misses++
	}

	if resistMultiplier == 0.75 {
		spellCast.PartialResists_1_4++
	} else if resistMultiplier == 0.5 {
		spellCast.PartialResists_2_4++
	} else if resistMultiplier == 0.25 {
		spellCast.PartialResists_3_4++
	}

	spellCast.TotalDamage += damage
}

func (spellEffect *SpellEffect) calculateDirectDamage(sim *Simulation, spellCast *SpellCast, ddInput *DirectDamageInput) {
	baseDamage := ddInput.MinBaseDamage + sim.RandomFloat("DirectSpell Base Damage")*(ddInput.MaxBaseDamage-ddInput.MinBaseDamage)

	totalSpellPower := spellCast.Character.GetStat(stats.SpellPower) + spellCast.Character.GetStat(spellCast.SpellSchool) + spellEffect.BonusSpellPower
	damageFromSpellPower := (totalSpellPower * ddInput.SpellCoefficient)

	damage := baseDamage + damageFromSpellPower + ddInput.FlatDamageBonus

	damage *= spellEffect.DamageMultiplier * spellEffect.StaticDamageMultiplier

	if spellEffect.critCheck(sim, spellCast) {
		spellEffect.Crit = true
		damage *= spellCast.CritMultiplier
	}

	if !spellCast.Binary {
		damage, _ = calculateResists(sim, damage, spellEffect)
	}

	spellEffect.Damage = damage
}

func (spellEffect *SpellEffect) applyDot(sim *Simulation, spellCast *SpellCast, ddInput *DotDamageInput) {
	totalSpellPower := spellCast.Character.GetStat(stats.SpellPower) + spellCast.Character.GetStat(spellCast.SpellSchool) + spellEffect.BonusSpellPower

	// snapshot total damage per tick, including any static damage multipliers
	ddInput.startTime = sim.CurrentTime
	ddInput.finalTickTime = sim.CurrentTime + time.Duration(ddInput.NumberOfTicks)*ddInput.TickLength
	ddInput.damagePerTick = (ddInput.TickBaseDamage + totalSpellPower*ddInput.TickSpellCoefficient) * spellEffect.StaticDamageMultiplier

	pa := &PendingAction{
		Name:         spellCast.ActionID.String(),
		NextActionAt: sim.CurrentTime + ddInput.TickLength,
	}

	pa.OnAction = func(sim *Simulation) {
		// fmt.Printf("DOT (%s) Ticking, Time Remaining: %0.2f\n", spellCast.Name, ddInput.TimeRemaining(sim).Seconds())
		damage := ddInput.damagePerTick

		spellCast.Character.OnBeforePeriodicDamage(sim, spellCast, spellEffect, &damage)
		spellEffect.Target.OnBeforePeriodicDamage(sim, spellCast, spellEffect, &damage)
		if ddInput.OnBeforePeriodicDamage != nil {
			ddInput.OnBeforePeriodicDamage(sim, spellCast, spellEffect, &damage)
		}
		if ddInput.IgnoreDamageModifiers {
			damage = ddInput.damagePerTick
		}

		hit := !ddInput.TicksCanMissAndCrit || spellEffect.hitCheck(sim, spellCast)
		crit := false
		resistMultiplier := 1.0

		if hit {
			if ddInput.TicksCanMissAndCrit && spellEffect.critCheck(sim, spellCast) {
				crit = true
				damage *= spellCast.CritMultiplier
			}

			if !spellCast.Binary {
				damage, resistMultiplier = calculateResists(sim, damage, spellEffect)
			}
		} else {
			damage = 0
		}

		if sim.Log != nil {
			spellCast.Character.Log(sim, "%s %s.", spellCast.ActionID, spellEffect.DotResultToString(damage, hit, crit, resistMultiplier))
		}
		spellEffect.Damage += damage

		if ddInput.TicksCanMissAndCrit {
			spellEffect.applyDotTickResultsToCast(spellCast, hit, crit, resistMultiplier, damage)
		}

		if ddInput.TicksProcSpellHitEffects {
			spellEffect.triggerSpellProcs(sim, spellCast)
		}

		spellCast.Character.OnPeriodicDamage(sim, spellCast, spellEffect, damage)
		spellEffect.Target.OnPeriodicDamage(sim, spellCast, spellEffect, damage)
		if ddInput.OnPeriodicDamage != nil {
			ddInput.OnPeriodicDamage(sim, spellCast, spellEffect, damage)
		}

		ddInput.tickIndex++
		if ddInput.tickIndex < ddInput.NumberOfTicks {
			// Refresh action.
			pa.NextActionAt = sim.CurrentTime + ddInput.TickLength
			sim.AddPendingAction(pa)
		} else {
			pa.CleanUp(sim)
		}
	}
	pa.CleanUp = func(sim *Simulation) {
		if ddInput.currentDotAction == nil {
			return
		}
		ddInput.currentDotAction = nil

		// Complete metrics and adding results etc
		if !ddInput.TicksCanMissAndCrit { // If false, results were already applied in applyDot().
			spellEffect.applyResultsToCast(spellCast)
		}
		spellCast.Character.Metrics.AddSpellCast(spellCast)

		// Clean up the dot object.
		ddInput.finalTickTime = 0
		spellCast.objectInUse = false

		if ddInput.DebuffID != 0 {
			spellEffect.Target.AddAuraUptime(ddInput.DebuffID, spellCast.ActionID, sim.CurrentTime-ddInput.startTime)
		}
	}

	ddInput.currentDotAction = pa
	sim.AddPendingAction(pa)
}

func (spellEffect *SpellEffect) String() string {
	if !spellEffect.Hit {
		return "Miss"
	}

	var sb strings.Builder

	if spellEffect.Crit {
		sb.WriteString("Crit")
	} else {
		sb.WriteString("Hit")
	}

	fmt.Fprintf(&sb, " for %0.3f damage", spellEffect.Damage)

	if spellEffect.PartialResist_1_4 {
		sb.WriteString(" (25% Resist)")
	} else if spellEffect.PartialResist_2_4 {
		sb.WriteString(" (50% Resist)")
	} else if spellEffect.PartialResist_3_4 {
		sb.WriteString(" (75% Resist)")
	}

	return sb.String()
}

func (spellEffect *SpellEffect) DotResultToString(damage float64, hit bool, crit bool, resistMultiplier float64) string {
	if !hit {
		return "tick Missed"
	}

	var sb strings.Builder

	fmt.Fprintf(&sb, "ticked for %0.3f damage", damage)

	if resistMultiplier == 0.75 {
		sb.WriteString(" (25% Resist)")
	} else if resistMultiplier == 0.5 {
		sb.WriteString(" (50% Resist)")
	} else if resistMultiplier == 0.25 {
		sb.WriteString(" (75% Resist)")
	}

	if crit {
		sb.WriteString(" (Crit)")
	}

	return sb.String()
}

// Return value is (newDamage, resistMultiplier)
func calculateResists(sim *Simulation, damage float64, spellEffect *SpellEffect) (float64, float64) {
	// Average Resistance (AR) = (Target's Resistance / (Caster's Level * 5)) * 0.75
	// P(x) = 50% - 250%*|x - AR| <- where X is %resisted
	// Using these stats:
	//    13.6% chance of
	//  FUTURE: handle boss resists for fights/classes that are actually impacted by that.
	resVal := sim.RandomFloat("DirectSpell Resist")
	if resVal > 0.18 { // 13% chance for 25% resist, 4% for 50%, 1% for 75%
		// No partial resist.
		return damage, 0
	}

	var multiplier float64
	if resVal < 0.01 {
		spellEffect.PartialResist_3_4 = true
		multiplier = 0.25
	} else if resVal < 0.05 {
		spellEffect.PartialResist_2_4 = true
		multiplier = 0.5
	} else {
		spellEffect.PartialResist_1_4 = true
		multiplier = 0.75
	}

	return damage * multiplier, multiplier
}
