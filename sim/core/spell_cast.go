package core

import (
	"fmt"
	"strings"
	"time"

	"github.com/wowsims/tbc/sim/core/stats"
)

// Callback for after a spell hits the target but before damage has been calculated.
type OnBeforeSpellHit func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect)

// Callback for after a spell hits the target but before damage is calculated.
// The damage result can still be modified by changing the result fields.
type OnSpellHit func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect)

// Callback for after a spell is fully resisted on a target.
type OnSpellMiss func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect)

// OnPeriodicDamage is called when dots tick. Able to mutate tickDamage as needed
type OnPeriodicDamage func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect, tickDamage *float64)

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

	spellEffect.calculateHit(sim, spellCast)
}

func (spellEffect *SpellEffect) afterCalculations(sim *Simulation, spellCast *SpellCast) {
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

	if sim.Log != nil {
		spellCast.Character.Log(sim, "%s result: %s", spellCast.Name, spellEffect)
	}
}

// Calculates whether this spell 'hit' and updates the effect field with the result.
func (spellEffect *SpellEffect) calculateHit(sim *Simulation, spellCast *SpellCast) {
	hit := 0.83 + (spellCast.Character.GetStat(stats.SpellHit)+spellEffect.BonusSpellHitRating)/(SpellHitRatingPerHitChance*100)
	hit = MinFloat(hit, 0.99) // can't get away from the 1% miss

	spellEffect.Hit = sim.RandomFloat("SpellCast Hit") < hit
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

func (spellEffect *SpellEffect) calculateDirectDamage(sim *Simulation, spellCast *SpellCast, ddInput *DirectDamageInput) {
	baseDamage := ddInput.MinBaseDamage + sim.RandomFloat("DirectSpell Base Damage")*(ddInput.MaxBaseDamage-ddInput.MinBaseDamage)

	totalSpellPower := spellCast.Character.GetStat(stats.SpellPower) + spellCast.Character.GetStat(spellCast.SpellSchool) + spellEffect.BonusSpellPower
	damageFromSpellPower := (totalSpellPower * ddInput.SpellCoefficient)

	damage := baseDamage + damageFromSpellPower + ddInput.FlatDamageBonus

	damage *= spellEffect.DamageMultiplier * spellEffect.StaticDamageMultiplier

	crit := (spellCast.Character.GetStat(stats.SpellCrit) + spellEffect.BonusSpellCritRating) / (SpellCritRatingPerCritChance * 100)
	if spellCast.GuaranteedCrit || sim.RandomFloat("DirectSpell Crit") < crit {
		spellEffect.Crit = true
		damage *= spellCast.CritMultiplier
	}

	if !spellCast.Binary {
		damage = calculateResists(sim, damage, spellEffect)
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
		Name:         spellCast.Name,
		NextActionAt: sim.CurrentTime + ddInput.TickLength,
	}

	pa.OnAction = func(sim *Simulation) {
		// fmt.Printf("DOT (%s) Ticking, Time Remaining: %0.2f\n", spellCast.Name, ddInput.TimeRemaining(sim).Seconds())
		damage := ddInput.damagePerTick
		spellCast.Character.OnPeriodicDamage(sim, spellCast, spellEffect, &damage)
		spellEffect.Target.OnPeriodicDamage(sim, spellCast, spellEffect, &damage)

		if !spellCast.Binary {
			damage = calculateResists(sim, damage, spellEffect)
		}

		if ddInput.OnPeriodicDamage != nil {
			ddInput.OnPeriodicDamage(sim, spellCast, spellEffect, &damage)
		}

		if sim.Log != nil {
			spellCast.Character.Log(sim, "%s ticked for %0.1f.", spellCast.Name, damage)
		}

		spellEffect.Damage += damage

		ddInput.tickIndex++
		if ddInput.tickIndex < ddInput.NumberOfTicks {
			// Refresh pending time.
			pa.NextActionAt = sim.CurrentTime + ddInput.TickLength
		}
	}
	pa.CleanUp = func(sim *Simulation) {
		if pa.NextActionAt == NeverExpires {
			panic("Already cleaned up dot")
		}
		pa.NextActionAt = NeverExpires

		// Complete metrics and adding results etc
		spellEffect.applyResultsToCast(spellCast)
		spellCast.Character.Metrics.AddSpellCast(spellCast)
		spellCast.objectInUse = false

		if ddInput.DebuffID != 0 && ddInput.SpellID != 0 {
			spellEffect.Target.AddAuraUptime(ddInput.DebuffID, ddInput.SpellID, sim.CurrentTime-ddInput.startTime)
		}
	}

	sim.AddPendingAction(pa)
}

func (spellEffect *SpellEffect) String() string {
	if !spellEffect.Hit {
		return "Miss"
	}

	var sb strings.Builder

	if spellEffect.PartialResist_1_4 {
		sb.WriteString("25% Resist ")
	} else if spellEffect.PartialResist_2_4 {
		sb.WriteString("50% Resist ")
	} else if spellEffect.PartialResist_3_4 {
		sb.WriteString("75% Resist ")
	}

	if spellEffect.Crit {
		sb.WriteString("Crit")
	} else {
		sb.WriteString("Hit")
	}

	fmt.Fprintf(&sb, " for %0.2f damage", spellEffect.Damage)
	return sb.String()
}

func calculateResists(sim *Simulation, damage float64, spellEffect *SpellEffect) float64 {
	// Average Resistance (AR) = (Target's Resistance / (Caster's Level * 5)) * 0.75
	// P(x) = 50% - 250%*|x - AR| <- where X is %resisted
	// Using these stats:
	//    13.6% chance of
	//  FUTURE: handle boss resists for fights/classes that are actually impacted by that.
	resVal := sim.RandomFloat("DirectSpell Resist")
	if resVal > 0.18 { // 13% chance for 25% resist, 4% for 50%, 1% for 75%
		return damage // means didn't resist
	}

	if resVal < 0.01 {
		spellEffect.PartialResist_3_4 = true
		damage *= 0.25
	} else if resVal < 0.05 {
		spellEffect.PartialResist_2_4 = true
		damage *= 0.5
	} else {
		spellEffect.PartialResist_1_4 = true
		damage *= 0.75
	}

	return damage
}
