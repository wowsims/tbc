package core

import (
	"fmt"
	"strings"
	"time"

	"github.com/wowsims/tbc/sim/core/stats"
)

// Input needed to start casting a spell.
type DirectCastInput struct {
	// If set, CD for this action and GCD CD will be ignored, and this action
	// will not set new values for those CDs either.
	IgnoreCooldowns bool

	// If set, this spell will have its mana cost ignored.
	IgnoreManaCost bool

	ManaCost float64

	CastTime time.Duration

	// How much to multiply damage by, if this cast crits.
	CritMultiplier float64

	// If true, will force the cast to crit (if it doesnt miss).
	GuaranteedCrit bool
}

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

// Interface for direct cast spells to implement.
type DirectCastImpl interface {
	// Pass-through AgentAction methods
	GetActionID() ActionID
	GetName() string
	GetTag() int32
	GetCharacter() *Character

	// This is needed because a lot of effects that 'reduce mana cost by X%' are
	// calculated from the base mana cost.
	GetBaseManaCost() float64

	// I.e. for nature spells, return stats.NatureSpellPower
	GetSpellSchool() stats.Stat

	GetCooldown() time.Duration

	GetCastInput(sim *Simulation, cast DirectCastAction) DirectCastInput
	GetHitInputs(sim *Simulation, cast DirectCastAction) []DirectCastDamageInput

	// Lifecycle callbacks for additional custom effects.
	OnCastComplete(sim *Simulation, cast DirectCastAction)
	OnSpellHit(sim *Simulation, cast DirectCastAction, result *DirectCastDamageResult)
	OnSpellMiss(sim *Simulation, cast DirectCastAction)
}

type DirectCastAction struct {
	DirectCastImpl

	// The inputs to the cast. Auras are given a reference to these to modify them
	// before the spell begins casting.
	castInput DirectCastInput
}

func (action DirectCastAction) GetManaCost() float64 {
	return action.castInput.ManaCost
}

func (action DirectCastAction) GetDuration() time.Duration {
	return action.castInput.CastTime
}

func (action DirectCastAction) Act(sim *Simulation) bool {
	character := action.GetCharacter()

	if !action.castInput.IgnoreManaCost && action.castInput.ManaCost > 0 {
		if character.CurrentMana() < action.castInput.ManaCost {
			if sim.Log != nil {
				sim.Log("(%d) Failed casting %s, not enough mana. (Current Mana = %0.0f, Mana Cost = %0.0f)\n",
						character.ID, action.GetName(), character.CurrentMana(), action.castInput.ManaCost)
			}
			sim.MetricsAggregator.MarkOOM(character, sim.CurrentTime)

			return false
		}
	}

	if sim.Log != nil {
		sim.Log("(%d) Casting %s (Current Mana = %0.0f, Mana Cost = %0.0f, Cast Time = %s)\n",
				character.ID, action.GetName(), character.CurrentMana(), action.castInput.ManaCost, action.castInput.CastTime)
	}

	// For instant-cast spells we can skip creating an aura.
	if action.castInput.CastTime == 0 {
		action.internalOnCastComplete(sim)
	} else {
		character.HardcastAura = Aura{
			Expires: sim.CurrentTime + action.castInput.CastTime,
			OnExpire: func(sim *Simulation) {
				action.internalOnCastComplete(sim)
			},
		}
	}

	if !action.castInput.IgnoreCooldowns {
		// Prevent any actions on the GCD until the cast AND the GCD are done.
		gcdCD := MaxDuration(GCDMin, action.castInput.CastTime)
		character.SetCD(GCDCooldownID, sim.CurrentTime+gcdCD)

		// TODO: Hardcasts seem to also reset swing timers, so we should set those CDs as well.
	}

	return true
}

func (action DirectCastAction) internalOnCastComplete(sim *Simulation) {
	character := action.GetCharacter()

	if !action.castInput.IgnoreManaCost && action.castInput.ManaCost > 0 {
		character.AddStat(stats.Mana, -action.castInput.ManaCost)
	}

	action.OnCastComplete(sim, action)
	character.OnCastComplete(sim, action)

	hitInputs := action.GetHitInputs(sim, action)

	results := make([]DirectCastDamageResult, 0, len(hitInputs))
	for _, hitInput := range hitInputs {
		character.OnBeforeSpellHit(sim, &hitInput)
		hitInput.Target.OnBeforeSpellHit(sim, &hitInput)
		result := action.calculateDirectCastDamage(sim, hitInput)

		if result.Hit {
			// Apply any on spell hit effects.
			action.OnSpellHit(sim, action, &result)
			character.OnSpellHit(sim, action, &result)
			hitInput.Target.OnSpellHit(sim, action, &result)
		} else {
			action.OnSpellMiss(sim, action)
			character.OnSpellMiss(sim, action)
			hitInput.Target.OnSpellMiss(sim, action)
		}
		if sim.Log != nil {
			sim.Log("(%d) %s result: %s\n", character.ID, action.GetName(), result)
		}

		results = append(results, result)
	}


	if !action.castInput.IgnoreCooldowns {
		cooldown := action.GetCooldown()
		if cooldown > 0 {
			character.SetCD(action.GetActionID().CooldownID, sim.CurrentTime+cooldown)
		}
	}

	sim.MetricsAggregator.AddCastAction(action, results)
}

func (action DirectCastAction) calculateDirectCastDamage(sim *Simulation, damageInput DirectCastDamageInput) DirectCastDamageResult {
	result := DirectCastDamageResult{
		Target: damageInput.Target,
	}

	character := action.GetCharacter()

	hit := 0.83 + character.GetStat(stats.SpellHit)/(SpellHitRatingPerHitChance*100) + damageInput.BonusHit
	hit = MinFloat(hit, 0.99) // can't get away from the 1% miss

	if sim.RandomFloat("action hit") >= hit { // Miss
		return result
	}
	result.Hit = true

	baseDamage := damageInput.MinBaseDamage + sim.RandomFloat("action dmg")*(damageInput.MaxBaseDamage - damageInput.MinBaseDamage)
	totalSpellPower := character.GetStat(stats.SpellPower) + character.GetStat(action.GetSpellSchool()) + damageInput.BonusSpellPower
	damageFromSpellPower := (totalSpellPower * damageInput.SpellCoefficient)
	damage := baseDamage + damageFromSpellPower

	damage *= damageInput.DamageMultiplier

	crit := (character.GetStat(stats.SpellCrit) / (SpellCritRatingPerCritChance * 100)) + damageInput.BonusCrit
	if action.castInput.GuaranteedCrit || sim.RandomFloat("action crit") < crit {
		result.Crit = true
		damage *= action.castInput.CritMultiplier
	}

	// Average Resistance (AR) = (Target's Resistance / (Caster's Level * 5)) * 0.75
	// P(x) = 50% - 250%*|x - AR| <- where X is %resisted
	// Using these stats:
	//    13.6% chance of
	//  FUTURE: handle boss resists for fights/classes that are actually impacted by that.
	resVal := sim.RandomFloat("action resist")
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

func NewDirectCastAction(sim *Simulation, impl DirectCastImpl) DirectCastAction {
	action := DirectCastAction{
		DirectCastImpl: impl,
	}
	character := action.GetCharacter()

	cooldownID := action.GetActionID().CooldownID

	castInput := impl.GetCastInput(sim, action)
	castInput.CastTime = time.Duration(float64(castInput.CastTime) / character.HasteBonus())

	// Apply on-cast effects.
	character.OnCast(sim, action, &castInput)

	// By panicking if spell is on CD, we force each sim to properly check for their own CDs.
	if !castInput.IgnoreCooldowns {
		if character.IsOnCD(GCDCooldownID, sim.CurrentTime) {
			panic(fmt.Sprintf("Trying to cast %s but GCD on cooldown for %s", action.GetName(), character.GetRemainingCD(GCDCooldownID, sim.CurrentTime)))
		}
		if character.IsOnCD(cooldownID, sim.CurrentTime) {
			panic(fmt.Sprintf("Trying to cast %s but is still on cooldown for %s", action.GetName(), character.GetRemainingCD(cooldownID, sim.CurrentTime)))
		}
	}

	action.castInput = castInput

	return action
}
