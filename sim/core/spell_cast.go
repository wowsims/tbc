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

func (hitEffect *SpellHitEffect) calculateDirectDamage(sim *Simulation, spellCast *SpellCast) {
	baseDamage := hitEffect.DirectInput.MinBaseDamage + sim.RandomFloat("DirectSpell Base Damage")*(hitEffect.DirectInput.MaxBaseDamage-hitEffect.DirectInput.MinBaseDamage)

	totalSpellPower := spellCast.Character.GetStat(stats.SpellPower) + spellCast.Character.GetStat(spellCast.SpellSchool) + hitEffect.SpellEffect.BonusSpellPower
	damageFromSpellPower := (totalSpellPower * hitEffect.DirectInput.SpellCoefficient)

	damage := baseDamage + damageFromSpellPower + hitEffect.DirectInput.FlatDamageBonus

	damage *= hitEffect.SpellEffect.DamageMultiplier * hitEffect.SpellEffect.StaticDamageMultiplier

	if hitEffect.SpellEffect.critCheck(sim, spellCast) {
		hitEffect.SpellEffect.Crit = true
		damage *= spellCast.CritMultiplier
	}

	if !spellCast.Binary {
		damage, _ = calculateResists(sim, damage, &hitEffect.SpellEffect)
	}

	hitEffect.SpellEffect.Damage = damage
}

// Snapshots a few values at the start of a dot.
func (hitEffect *SpellHitEffect) takeDotSnapshot(sim *Simulation, spellCast *SpellCast) {
	totalSpellPower := spellCast.Character.GetStat(stats.SpellPower) + spellCast.Character.GetStat(spellCast.SpellSchool) + hitEffect.BonusSpellPower

	// snapshot total damage per tick, including any static damage multipliers
	hitEffect.DotInput.startTime = sim.CurrentTime
	hitEffect.DotInput.finalTickTime = sim.CurrentTime + time.Duration(hitEffect.DotInput.NumberOfTicks)*hitEffect.DotInput.TickLength
	hitEffect.DotInput.damagePerTick = (hitEffect.DotInput.TickBaseDamage + totalSpellPower*hitEffect.DotInput.TickSpellCoefficient) * hitEffect.StaticDamageMultiplier
}

// This should be called on each dot tick.
func (hitEffect *SpellHitEffect) onDotTick(sim *Simulation, spellCast *SpellCast) {
	// fmt.Printf("DOT (%s) Ticking, Time Remaining: %0.2f\n", spellCast.Name, hitEffect.DotInput.TimeRemaining(sim).Seconds())
	damage := hitEffect.DotInput.damagePerTick

	spellCast.Character.OnBeforePeriodicDamage(sim, spellCast, &hitEffect.SpellEffect, &damage)
	hitEffect.Target.OnBeforePeriodicDamage(sim, spellCast, &hitEffect.SpellEffect, &damage)
	if hitEffect.DotInput.OnBeforePeriodicDamage != nil {
		hitEffect.DotInput.OnBeforePeriodicDamage(sim, spellCast, &hitEffect.SpellEffect, &damage)
	}
	if hitEffect.DotInput.IgnoreDamageModifiers {
		damage = hitEffect.DotInput.damagePerTick
	}

	hit := !hitEffect.DotInput.TicksCanMissAndCrit || hitEffect.hitCheck(sim, spellCast)
	crit := false
	resistMultiplier := 1.0

	if hit {
		if hitEffect.DotInput.TicksCanMissAndCrit && hitEffect.critCheck(sim, spellCast) {
			crit = true
			damage *= spellCast.CritMultiplier
		}

		if !spellCast.Binary {
			damage, resistMultiplier = calculateResists(sim, damage, &hitEffect.SpellEffect)
		}
	} else {
		damage = 0
	}

	if sim.Log != nil {
		spellCast.Character.Log(sim, "%s %s.", spellCast.ActionID, hitEffect.SpellEffect.DotResultToString(damage, hit, crit, resistMultiplier))
	}
	hitEffect.SpellEffect.Damage += damage

	if hitEffect.DotInput.TicksCanMissAndCrit {
		hitEffect.SpellEffect.applyDotTickResultsToCast(spellCast, hit, crit, resistMultiplier, damage)
	}

	if hitEffect.DotInput.TicksProcSpellHitEffects {
		hitEffect.SpellEffect.triggerSpellProcs(sim, spellCast)
	}

	spellCast.Character.OnPeriodicDamage(sim, spellCast, &hitEffect.SpellEffect, damage)
	hitEffect.Target.OnPeriodicDamage(sim, spellCast, &hitEffect.SpellEffect, damage)
	if hitEffect.DotInput.OnPeriodicDamage != nil {
		hitEffect.DotInput.OnPeriodicDamage(sim, spellCast, &hitEffect.SpellEffect, damage)
	}

	hitEffect.DotInput.tickIndex++
}

// This should be called after the final tick of the dot, or when the dot is cancelled.
func (hitEffect *SpellHitEffect) onDotComplete(sim *Simulation, spellCast *SpellCast) {
	// Complete metrics and adding results etc
	if !hitEffect.DotInput.TicksCanMissAndCrit { // If false, results were already applied in onDotTick().
		hitEffect.SpellEffect.applyResultsToCast(spellCast)
	}

	// Clean up the dot object.
	hitEffect.DotInput.finalTickTime = 0

	if hitEffect.DotInput.DebuffID != 0 {
		hitEffect.Target.AddAuraUptime(hitEffect.DotInput.DebuffID, spellCast.ActionID, sim.CurrentTime-hitEffect.DotInput.startTime)
	}
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
