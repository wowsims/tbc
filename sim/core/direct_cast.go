package core

import (
	"fmt"
	"strings"
	"time"

	"github.com/wowsims/tbc/sim/core/stats"
)

// A direct spell is one that does a single instance of damage once casting is
// complete, i.e. shadowbolt or fire blast.

// Input needed to calculate the damage of a direct spell.
type DirectCastDamageInput struct {
	// Target of the spell.
	Target *Target

	MinBaseDamage float64
	MaxBaseDamage float64

	// Increase in damage per point of spell power.
	SpellCoefficient float64

	// Bonus stats to be added to the damage calculation.
	BonusSpellPower      float64
	BonusSpellHitRating  float64
	BonusSpellCritRating float64

	// Additional multiplier that is always applied.
	DamageMultiplier float64
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

	// Inputs for each hit.
	HitInputs []DirectCastDamageInput

	// Results from each hit. For performance reasons, this should be pre-allocated
	// by the caller.
	HitResults []DirectCastDamageResult

	// Callbacks for providing additional custom behavior.
	OnSpellHit  OnSpellHit
	OnSpellMiss OnSpellMiss
}

func (action *DirectCastAction) GetActionID() ActionID {
	return action.Cast.ActionID
}

func (action *DirectCastAction) GetName() string {
	return action.Cast.Name
}

func (action *DirectCastAction) GetTag() int32 {
	return action.Cast.Tag
}

func (action *DirectCastAction) GetCharacter() *Character {
	return action.Cast.Character
}

func (action *DirectCastAction) GetManaCost() float64 {
	return action.Cast.ManaCost
}

func (action *DirectCastAction) GetDuration() time.Duration {
	return action.Cast.CastTime
}

func (action *DirectCastAction) Init(sim *Simulation) {
	action.Cast.init(sim)
}

func (action *DirectCastAction) Act(sim *Simulation) bool {
	return action.Cast.startCasting(sim, func(sim *Simulation, cast *Cast) {
		for hitIdx := range action.HitInputs {
			hitInput := &action.HitInputs[hitIdx]

			cast.Character.OnBeforeSpellHit(sim, cast, hitInput)
			hitInput.Target.OnBeforeSpellHit(sim, cast, hitInput)

			action.HitResults[hitIdx] = action.calculateDirectCastDamage(sim, hitInput)
		}

		// TODO: Allow hit results to be added to metrics individually, so we can merge
		// this and the surrounding loops together.
		sim.MetricsAggregator.AddCastAction(action, action.HitResults)

		for hitIdx := range action.HitInputs {
			hitInput := &action.HitInputs[hitIdx]
			hitResult := &action.HitResults[hitIdx]

			if hitResult.Hit {
				action.OnSpellHit(sim, cast, hitResult)
				cast.Character.OnSpellHit(sim, cast, hitResult)
				hitInput.Target.OnSpellHit(sim, cast, hitResult)
			} else {
				action.OnSpellMiss(sim, cast)
				cast.Character.OnSpellMiss(sim, cast)
				hitInput.Target.OnSpellMiss(sim, cast)
			}
			if sim.Log != nil {
				sim.Log("(%d) %s result: %s\n", cast.Character.ID, action.Cast.Name, hitResult)
			}
		}
	})
}

func (action *DirectCastAction) calculateDirectCastDamage(sim *Simulation, damageInput *DirectCastDamageInput) DirectCastDamageResult {
	result := DirectCastDamageResult{
		Target: damageInput.Target,
	}

	character := action.Cast.Character

	hit := 0.83 + (character.GetStat(stats.SpellHit)+damageInput.BonusSpellHitRating)/(SpellHitRatingPerHitChance*100)
	hit = MinFloat(hit, 0.99) // can't get away from the 1% miss

	if sim.RandomFloat("DirectCast Hit") >= hit { // Miss
		return result
	}
	result.Hit = true

	baseDamage := damageInput.MinBaseDamage + sim.RandomFloat("DirectCast Base Damage")*(damageInput.MaxBaseDamage-damageInput.MinBaseDamage)
	totalSpellPower := character.GetStat(stats.SpellPower) + character.GetStat(action.Cast.SpellSchool) + damageInput.BonusSpellPower
	damageFromSpellPower := (totalSpellPower * damageInput.SpellCoefficient)
	damage := baseDamage + damageFromSpellPower

	damage *= damageInput.DamageMultiplier

	crit := (character.GetStat(stats.SpellCrit) + damageInput.BonusSpellCritRating) / (SpellCritRatingPerCritChance * 100)
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

type DirectCastTemplateGenerator func() DirectCastAction

// NewDirectCastTemplateGenerator will take in a cast template and create a generator so you dont have to manually manage hit inputs.
func NewDirectCastTemplateGenerator(template DirectCastAction) DirectCastTemplateGenerator {
	hitInput := make([]DirectCastDamageInput, len(template.HitInputs))
	return func() DirectCastAction {
		newAction := template
		newAction.HitInputs = hitInput
		copy(newAction.HitInputs, template.HitInputs)
		return newAction
	}
}
