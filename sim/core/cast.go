package core

import (
	"fmt"
	"time"

	"github.com/wowsims/tbc/sim/core/stats"
)

// Callback for when a cast begins, i.e. when the in-game castbar starts filling up.
type OnCast func(sim *Simulation, cast *Cast)

// Callback for when a cast is finished, i.e. when the in-game castbar reaches full.
type OnCastComplete func(sim *Simulation, cast *Cast)

// A basic cast that costs mana and performs a callback when complete.
// Manages cooldowns and the GCD.
type Cast struct {
	// ID for the action.
	ActionID ActionID

	// The name of the cast action, e.g. 'Shadowbolt'.
	Name string

	// Subcategory for this action, for metrics.
	Tag int32

	// The character performing this action.
	Character *Character

	// If set, this action will start a cooldown using its cooldown ID.
	// Note that the GCD CD will be activated even if this is not set.
	Cooldown time.Duration

	// If set, CD for this action and GCD CD will be ignored, and this action
	// will not set new values for those CDs either.
	IgnoreCooldowns bool

	// E.g. for nature spells, set to stats.NatureSpellPower.
	SpellSchool stats.Stat

	// If set, this spell will have its mana cost ignored.
	IgnoreManaCost bool

	// Base mana cost. Many effects in the game which 'reduce mana cost by X%'
	// are calculated using the base mana cost. Any effects which reduce the base
	// mana cost should be applied before setting this value, and OnCast()
	// callbacks should not modify it.
	BaseManaCost float64

	// Actual mana cost of the spell.
	ManaCost float64

	CastTime time.Duration

	// How much to multiply damage by, if this cast crits.
	CritMultiplier float64

	// If true, will force the cast to crit (if it doesnt miss).
	GuaranteedCrit bool
}

// Should be called exactly once after creation.
func (cast *Cast) init(sim *Simulation) {
	cast.CastTime = time.Duration(float64(cast.CastTime) / cast.Character.HasteBonus())

	// Apply on-cast effects.
	cast.Character.OnCast(sim, cast)

	// By panicking if spell is on CD, we force each sim to properly check for their own CDs.
	if !cast.IgnoreCooldowns {
		if cast.Character.IsOnCD(GCDCooldownID, sim.CurrentTime) {
			panic(fmt.Sprintf("Trying to cast %s but GCD on cooldown for %s", cast.Name, cast.Character.GetRemainingCD(GCDCooldownID, sim.CurrentTime)))
		}

		cooldownID := cast.ActionID.CooldownID
		if cast.Character.IsOnCD(cooldownID, sim.CurrentTime) {
			panic(fmt.Sprintf("Trying to cast %s but is still on cooldown for %s", cast.Name, cast.Character.GetRemainingCD(cooldownID, sim.CurrentTime)))
		}
	}
}

// Start casting the spell. Return value indicates whether the spell successfully
// started casting.
func (cast *Cast) startCasting(sim *Simulation, onCastComplete OnCastComplete) bool {
	if !cast.IgnoreManaCost && cast.ManaCost > 0 {
		if cast.Character.CurrentMana() < cast.ManaCost {
			if sim.Log != nil {
				sim.Log("(%d) Failed casting %s, not enough mana. (Current Mana = %0.0f, Mana Cost = %0.0f)\n",
						cast.Character.ID, cast.Name, cast.Character.CurrentMana(), cast.ManaCost)
			}
			sim.MetricsAggregator.MarkOOM(cast.Character, sim.CurrentTime)

			return false
		}
	}

	if sim.Log != nil {
		sim.Log("(%d) Casting %s (Current Mana = %0.0f, Mana Cost = %0.0f, Cast Time = %s)\n",
				cast.Character.ID, cast.Name, cast.Character.CurrentMana(), cast.ManaCost, cast.CastTime)
	}

	// For instant-cast spells we can skip creating an aura.
	if cast.CastTime == 0 {
		cast.internalOnComplete(sim, onCastComplete)
	} else {
		cast.Character.HardcastAura = Aura{
			Expires: sim.CurrentTime + cast.CastTime,
			OnExpire: func(sim *Simulation) {
				cast.internalOnComplete(sim, onCastComplete)
			},
		}
	}

	if !cast.IgnoreCooldowns {
		// Prevent any actions on the GCD until the cast AND the GCD are done.
		gcdCD := MaxDuration(GCDMin, cast.CastTime)
		cast.Character.SetCD(GCDCooldownID, sim.CurrentTime+gcdCD)

		// TODO: Hardcasts seem to also reset swing timers, so we should set those CDs as well.
	}

	return true
}

// Cast has finished, activate the effects of the cast.
func (cast *Cast) internalOnComplete(sim *Simulation, onCastComplete OnCastComplete) {
	if !cast.IgnoreManaCost && cast.ManaCost > 0 {
		cast.Character.AddStat(stats.Mana, -cast.ManaCost)
	}

	if !cast.IgnoreCooldowns {
		if cast.Cooldown > 0 {
			cast.Character.SetCD(cast.ActionID.CooldownID, sim.CurrentTime+cast.Cooldown)
		}
	}

	onCastComplete(sim, cast)
	cast.Character.OnCastComplete(sim, cast)
}
