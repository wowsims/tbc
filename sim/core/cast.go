package core

import (
	"fmt"
	"time"

	"github.com/wowsims/tbc/sim/core/stats"
)

type ResourceCost struct {
	Type  stats.Stat // stats.Mana, stats.Energy, stats.Rage
	Value float64
}

// A cast corresponds to any action which causes the in-game castbar to be
// shown, and activates the GCD. Note that a cast can also be instant, i.e.
// the effects are applied immediately even though the GCD is still activated.

// Callback for when a cast begins, i.e. when the in-game castbar starts filling up.
type OnCast func(sim *Simulation, cast *Cast)

// Callback for when a cast is finished, i.e. when the in-game castbar reaches full.
type OnCastComplete func(sim *Simulation, cast *Cast)

// A basic cast that costs mana and performs a callback when complete.
// Manages cooldowns and the GCD.
type Cast struct {
	// ID for the action.
	ActionID

	// The character performing this action.
	Character *Character

	// If set, this action will start a cooldown using its cooldown ID.
	// Note that the GCD CD will be activated even if this is not set.
	Cooldown time.Duration

	// The amount of GCD time incurred by this cast. This is almost always 0, 1s, or 1.5s.
	GCD time.Duration

	// Whether this is a phantom cast. Phantom casts are usually casts triggered by some effect,
	// like The Lightning Capacitor or Shaman Flametongue Weapon. Many on-hit effects do not
	// proc from phantom casts, only regular casts.
	IsPhantom bool

	OutcomeRollCategory OutcomeRollCategory
	CritRollCategory    CritRollCategory
	SpellSchool         SpellSchool
	SpellExtras         SpellExtras

	// Base cost. Many effects in the game which 'reduce mana cost by X%'
	// are calculated using the base mana cost. Any effects which reduce the base
	// mana cost should be applied before setting this value, and OnCast()
	// callbacks should not modify it.
	BaseCost ResourceCost

	// Actual mana cost of the spell.
	Cost ResourceCost

	CastTime time.Duration

	// Adds additional delay to the GCD after the cast is completed. This is usually
	// used for adding latency following the cast.
	AfterCastDelay time.Duration

	// How much to multiply damage by, if this cast crits.
	CritMultiplier float64

	// Bonus crit to be applied to all effects resulting from this cast.
	BonusCritRating float64

	// Callbacks for providing additional custom behavior.
	OnCastComplete OnCastComplete

	// Ignores haste when calculating the GCD and cast time for this cast.
	IgnoreHaste bool

	// Internal field only, used to prevent cast pool objects from being used by
	// multiple casts simultaneously.
	objectInUse bool
}

// AgentAction functions for actions that embed a Cast.

func (cast *Cast) GetActionID() ActionID {
	return cast.ActionID
}

func (cast *Cast) GetCharacter() *Character {
	return cast.Character
}

func (cast *Cast) GetManaCost() float64 {
	return cast.Cost.Value
}

func (cast *Cast) GetDuration() time.Duration {
	return cast.CastTime
}

func (cast *Cast) IsInUse() bool {
	return cast.objectInUse
}

// Cancel will disable 'in use' so the cast can be reused. Useful if deciding not to cast.
func (cast *Cast) Cancel() {
	cast.objectInUse = false
}

// Should be called exactly once after creation.
func (cast *Cast) init(sim *Simulation) {
	if cast.Character == nil {
		panic("Character not set on cast")
	}
	if cast.objectInUse {
		panic("Cast object already in use")
	}
	cast.objectInUse = true

	if !cast.IgnoreHaste {
		cast.CastTime = time.Duration(float64(cast.CastTime) / cast.Character.CastSpeed())
	}

	// Apply on-cast effects.
	cast.Character.OnCast(sim, cast)

	// By panicking if spell is on CD, we force each sim to properly check for their own CDs.
	if cast.GCD != 0 && cast.Character.IsOnCD(GCDCooldownID, sim.CurrentTime) {
		panic(fmt.Sprintf("Trying to cast %s but GCD on cooldown for %s", cast.ActionID, cast.Character.GetRemainingCD(GCDCooldownID, sim.CurrentTime)))
	}

	if cast.Cooldown != 0 {
		cooldownID := cast.ActionID.CooldownID
		if cast.Character.IsOnCD(cooldownID, sim.CurrentTime) {
			panic(fmt.Sprintf("Trying to cast %s but is still on cooldown for %s", cast.ActionID, cast.Character.GetRemainingCD(cooldownID, sim.CurrentTime)))
		}
	}
}

// Start casting the spell. Return value indicates whether the spell successfully
// started casting.
func (cast *Cast) startCasting(sim *Simulation, onCastComplete OnCastComplete) bool {

	switch cast.Cost.Type {
	case stats.Mana:
		if cast.Character.CurrentMana() < cast.Cost.Value {
			if sim.Log != nil {
				cast.Character.Log(sim, "Failed casting %s, not enough mana. (Current Mana = %0.03f, Mana Cost = %0.03f)",
					cast.ActionID, cast.Character.CurrentMana(), cast.Cost.Value)
			}
			cast.objectInUse = false
			return false
		}
	case stats.Rage:
		if cast.Character.CurrentRage() < cast.Cost.Value {
			return false
		}
		cast.Character.SpendRage(sim, cast.Cost.Value, cast.ActionID)
	case stats.Energy:
		if cast.Character.CurrentEnergy() < cast.Cost.Value {
			return false
		}
		cast.Character.SpendEnergy(sim, cast.Cost.Value, cast.ActionID)
	}

	if sim.Log != nil {
		cast.Character.Log(sim, "Casting %s (Cost = %0.03f, Cast Time = %s)",
			cast.ActionID, MaxFloat(0, cast.Cost.Value), cast.CastTime)
	}

	// This needs to come before the internalOnComplete() call so that changes to
	// casting speed caused by the cast don't affect the GCD CD.
	if cast.GCD != 0 {
		// Prevent any actions on the GCD until the cast AND the GCD are done.
		gcdCD := MaxDuration(cast.CalculatedGCD(cast.Character), cast.CastTime+cast.AfterCastDelay)
		cast.Character.SetGCDTimer(sim, sim.CurrentTime+gcdCD)
	}

	if cast.Cooldown > 0 {
		cast.Character.SetCD(cast.ActionID.CooldownID, sim.CurrentTime+cast.CastTime+cast.Cooldown)
	}

	// For instant-cast spells we can skip creating an aura.
	if cast.CastTime == 0 {
		cast.internalOnComplete(sim, onCastComplete)
	} else {
		cast.Character.Hardcast.Expires = sim.CurrentTime + cast.CastTime
		cast.Character.Hardcast.Cast = cast
		cast.Character.Hardcast.OnComplete = onCastComplete

		// If hardcast and GCD happen at the same time then we don't need a separate action.
		if cast.Character.Hardcast.Expires != cast.Character.NextGCDAt() {
			cast.Character.hardcastAction.NextActionAt = cast.Character.Hardcast.Expires
			sim.AddPendingAction(cast.Character.hardcastAction)
		}

		if cast.Character.AutoAttacks.IsEnabled() {
			// Delay autoattacks until the cast is complete.
			cast.Character.AutoAttacks.DelayAllUntil(sim, cast.Character.Hardcast.Expires)
		}
	}

	return true
}

func (cast *Cast) CalculatedGCD(char *Character) time.Duration {
	if cast.IgnoreHaste {
		return cast.GCD
	} else {
		return MaxDuration(GCDMin, time.Duration(float64(cast.GCD)/char.CastSpeed()))
	}
}

// Cast has finished, activate the effects of the cast.
func (cast *Cast) internalOnComplete(sim *Simulation, onCastComplete OnCastComplete) {
	if cast.Cost.Value > 0 && cast.Cost.Type == stats.Mana {
		cast.Character.SpendMana(sim, cast.Cost.Value, cast.ActionID)
		cast.Character.PseudoStats.FiveSecondRuleRefreshTime = sim.CurrentTime + time.Second*5
	}

	cast.Character.OnCastComplete(sim, cast)
	if cast.OnCastComplete != nil {
		cast.OnCastComplete(sim, cast)
	}
	onCastComplete(sim, cast)
}

// A simple cast is just a cast with a callback, no calculations or damage.
type SimpleCast struct {
	// Embedded Cast
	Cast

	OnCastComplete OnCastComplete

	// Turns off metrics recording for this cast.
	DisableMetrics bool
}

func (simpleCast *SimpleCast) Init(sim *Simulation) {
	simpleCast.Cast.init(sim)
}

// TODO: Need to rename this. Cant call it Cast() because of conflict with field of the same name.
func (simpleCast *SimpleCast) StartCast(sim *Simulation) bool {
	return simpleCast.Cast.startCasting(sim, func(sim *Simulation, cast *Cast) {
		if !simpleCast.DisableMetrics {
			cast.Character.Metrics.AddCast(cast)
		}
		if simpleCast.OnCastComplete != nil {
			simpleCast.OnCastComplete(sim, cast)
		}
		simpleCast.objectInUse = false
	})
}

type Hardcast struct {
	Cast       *Cast
	Expires    time.Duration
	OnComplete OnCastComplete
}

func (hc Hardcast) OnExpire(sim *Simulation) {
	hc.Cast.internalOnComplete(sim, hc.OnComplete)
}
