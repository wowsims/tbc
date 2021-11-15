package core

import (
	"fmt"
	"strings"

	"github.com/wowsims/tbc/sim/core/stats"
)

// Callback for after a spell hits the target but before damage has been calculated.
type OnBeforeSpellHit func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect)

// Callback for after a spell hits the target but before damage is calculated.
// The damage result can still be modified by changing the result fields.
type OnSpellHit func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect)

// Callback for after a spell is fully resisted on a target.
type OnSpellMiss func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect)

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
		sim.Log("(%d) %s result: %s\n", spellCast.Character.ID, spellCast.Name, spellEffect)
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
