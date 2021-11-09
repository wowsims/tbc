package core

import (
	"fmt"
	"strings"
	"time"

	"github.com/wowsims/tbc/sim/core/stats"
)

// Input needed to calculate the damage of a spell.
type DirectCastDamageInput struct {
	// Target of the spell.
	Target *Target

	MinBaseDamage float64
	MaxBaseDamage float64

	// Increase in damage per point of spell power.
	SpellCoefficient float64

	// Additional multiplier that is always applied.
	DamageMultiplier float64

	BonusSpellPower float64
	BonusHit        float64 // Direct % bonus... 0.1 == 10%
	BonusCrit       float64 // Direct % bonus... 0.1 == 10%
}

type DirectCastDamageResult struct {
	Target *Target

	Hit  bool // True = hit, False = resisted
	Crit bool // Whether this cast was a critical strike.

	PartialResist_1_4 bool // 1/4 of the spell was resisted
	PartialResist_2_4 bool // 2/4 of the spell was resisted
	PartialResist_3_4 bool // 3/4 of the spell was resisted

	Damage float64 // Damage done by this cast.
}

func (result DirectCastDamageResult) String() string {
	if !result.Hit {
		return "Miss"
	}

	var sb strings.Builder

	if result.PartialResist_1_4 {
		sb.WriteString("25% Resist ")
	} else if result.PartialResist_2_4 {
		sb.WriteString("50% Resist ")
	} else if result.PartialResist_3_4 {
		sb.WriteString("75% Resist ")
	}

	if result.Crit {
		sb.WriteString("Crit")
	} else {
		sb.WriteString("Hit")
	}

	fmt.Fprintf(&sb, " for %0.2f damage", result.Damage)
	return sb.String()
}

// Callback for before the damage calculation of a spell hit happens.
type OnBeforeSpellHit func(sim *Simulation, cast *Cast, hitInput *DirectCastDamageInput)

// Callback for after a spell hits the target and damage has been calculated.
// The damage result can still be modified by changing the result fields.
type OnSpellHit func(sim *Simulation, cast *Cast, result *DirectCastDamageResult)

// Callback for after a spell is fully resisted on a target.
type OnSpellMiss func(sim *Simulation, cast *Cast)

type DirectCastAction struct {
	Cast

	HitInputs []DirectCastDamageInput

	// Callbacks for providing additional custom behavior.
	OnCastComplete OnCastComplete
	OnSpellHit     OnSpellHit
	OnSpellMiss    OnSpellMiss
}

func (action DirectCastAction) GetActionID() ActionID {
	return action.Cast.ActionID
}

func (action DirectCastAction) GetName() string {
	return action.Cast.Name
}

func (action DirectCastAction) GetTag() int32 {
	return action.Cast.Tag
}

func (action DirectCastAction) GetCharacter() *Character {
	return action.Cast.Character
}

func (action DirectCastAction) GetManaCost() float64 {
	return action.Cast.ManaCost
}

func (action DirectCastAction) GetDuration() time.Duration {
	return action.Cast.CastTime
}

func (action DirectCastAction) Act(sim *Simulation) bool {
	return action.Cast.startCasting(sim, func(sim *Simulation, cast *Cast) {
		action.OnCastComplete(sim, cast)

		results := make([]DirectCastDamageResult, 0, len(action.HitInputs))
		for hitIdx := range action.HitInputs {
			hitInput := &action.HitInputs[hitIdx]
			hitInput.Target.OnBeforeSpellHit(sim, cast, hitInput)

			result := action.calculateDirectCastDamage(sim, hitInput)
			results = append(results, result)

			if result.Hit {
				// Apply any on spell hit effects.
				action.OnSpellHit(sim, cast, &result)
				cast.Character.OnSpellHit(sim, cast, &result)
				hitInput.Target.OnSpellHit(sim, cast, &result)
			} else {
				action.OnSpellMiss(sim, cast)
				cast.Character.OnSpellMiss(sim, cast)
				hitInput.Target.OnSpellMiss(sim, cast)
			}
			if sim.Log != nil {
				sim.Log("(%d) %s result: %s\n", cast.Character.ID, action.Cast.Name, result)
			}
		}

		sim.MetricsAggregator.AddCastAction(action, results)
	})
}

func (action DirectCastAction) calculateDirectCastDamage(sim *Simulation, damageInput *DirectCastDamageInput) DirectCastDamageResult {
	result := DirectCastDamageResult{
		Target: damageInput.Target,
	}

	character := action.Cast.Character

	hit := 0.83 + character.GetStat(stats.SpellHit)/(SpellHitRatingPerHitChance*100) + damageInput.BonusHit
	hit = MinFloat(hit, 0.99) // can't get away from the 1% miss

	if sim.RandomFloat("DirectCast Hit") >= hit { // Miss
		return result
	}
	result.Hit = true

	baseDamage := damageInput.MinBaseDamage + sim.RandomFloat("DirectCast Damage")*(damageInput.MaxBaseDamage - damageInput.MinBaseDamage)
	totalSpellPower := character.GetStat(stats.SpellPower) + character.GetStat(action.Cast.SpellSchool) + damageInput.BonusSpellPower
	damageFromSpellPower := (totalSpellPower * damageInput.SpellCoefficient)
	damage := baseDamage + damageFromSpellPower

	damage *= damageInput.DamageMultiplier

	crit := (character.GetStat(stats.SpellCrit) / (SpellCritRatingPerCritChance * 100)) + damageInput.BonusCrit
	if action.Cast.GuaranteedCrit || sim.RandomFloat("DirectCast Crit") < crit {
		result.Crit = true
		damage *= action.Cast.CritMultiplier
	}

	// Average Resistance (AR) = (Target's Resistance / (Caster's Level * 5)) * 0.75
	// P(x) = 50% - 250%*|x - AR| <- where X is %resisted
	// Using these stats:
	//    13.6% chance of
	//  FUTURE: handle boss resists for fights/classes that are actually impacted by that.
	resVal := sim.RandomFloat("DirectCast Resist")
	if resVal < 0.18 { // 13% chance for 25% resist, 4% for 50%, 1% for 75%
		if resVal < 0.01 {
			result.PartialResist_3_4 = true
			damage *= .25
		} else if resVal < 0.05 {
			result.PartialResist_2_4 = true
			damage *= .5
		} else {
			result.PartialResist_1_4 = true
			damage *= .75
		}
	}

	result.Damage = damage

	return result
}

func NewDirectCastAction(sim *Simulation, cast Cast, hitInputs []DirectCastDamageInput, onCastComplete OnCastComplete, onSpellHit OnSpellHit, onSpellMiss OnSpellMiss) DirectCastAction {
	action := DirectCastAction{
		Cast: cast,
		HitInputs: hitInputs,
		OnCastComplete: onCastComplete,
		OnSpellHit: onSpellHit,
		OnSpellMiss: onSpellMiss,
	}

	action.Cast.init(sim)

	return action
}
