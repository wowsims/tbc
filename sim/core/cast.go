package core

import (
	"fmt"
	"time"

	"github.com/wowsims/tbc/sim/core/stats"
)

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
	ActionID ActionID

	// The name of the cast action, e.g. 'Shadowbolt'.
	Name string

	// The character performing this action.
	Character *Character

	// If set, this action will start a cooldown using its cooldown ID.
	// Note that the GCD CD will be activated even if this is not set.
	Cooldown time.Duration

	// If set, this will be used as the GCD instead of the default value (1.5s).
	GCDCooldown time.Duration

	// If set, CD for this action and GCD CD will be ignored, and this action
	// will not set new values for those CDs either.
	IgnoreCooldowns bool

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

	// E.g. for nature spells, set to stats.NatureSpellPower.
	SpellSchool stats.Stat

	// How much to multiply damage by, if this cast crits.
	CritMultiplier float64

	// If true, will force the cast to crit (if it doesnt miss).
	GuaranteedCrit bool

	// Callbacks for providing additional custom behavior.
	OnCastComplete OnCastComplete

	Binary bool // if spell is binary it ignores partial resists

	// Internal field only, used to prevent cast pool objects from being used by
	// multiple casts simultaneously.
	objectInUse bool
}

// AgentAction functions for actions that embed a Cast.

func (cast *Cast) GetActionID() ActionID {
	return cast.ActionID
}

func (cast *Cast) GetName() string {
	return cast.Name
}

func (cast *Cast) GetCharacter() *Character {
	return cast.Character
}

func (cast *Cast) GetManaCost() float64 {
	return cast.ManaCost
}

func (cast *Cast) GetDuration() time.Duration {
	return cast.CastTime
}

func (cast *Cast) IsInUse() bool {
	return cast.objectInUse
}

// Should be called exactly once after creation.
func (cast *Cast) init(sim *Simulation) {
	if cast.Character == nil {
		panic("character not set on cast")
	}
	cast.objectInUse = true
	cast.CastTime = time.Duration(float64(cast.CastTime) / cast.Character.CastSpeed())

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
				cast.Character.Log(sim, "Failed casting %s, not enough mana. (Current Mana = %0.0f, Mana Cost = %0.0f)",
					cast.Name, cast.Character.CurrentMana(), cast.ManaCost)
			}
			cast.Character.Metrics.MarkOOM(sim, cast.Character)
			cast.objectInUse = false // cast failed and we aren't using it
			return false
		}
	}

	if sim.Log != nil {
		cast.Character.Log(sim, "Casting %s (Current Mana = %0.0f, Mana Cost = %0.0f, Cast Time = %s)",
			cast.Name, cast.Character.CurrentMana(), cast.ManaCost, cast.CastTime)
	}

	// For instant-cast spells we can skip creating an aura.
	if cast.CastTime == 0 {
		cast.internalOnComplete(sim, onCastComplete)
	} else {
		cast.Character.Hardcast.Expires = sim.CurrentTime + cast.CastTime
		cast.Character.Hardcast.Cast = cast
		cast.Character.Hardcast.OnComplete = onCastComplete
	}

	if !cast.IgnoreCooldowns {
		// Prevent any actions on the GCD until the cast AND the GCD are done.
		gcdCD := MaxDuration(cast.CalculatedGCD(cast.Character), cast.CastTime)
		cast.Character.SetCD(GCDCooldownID, sim.CurrentTime+gcdCD)

		// TODO: Hardcasts seem to also reset swing timers, so we should set those CDs as well.
	}

	return true
}

func (cast *Cast) CalculatedGCD(char *Character) time.Duration {
	baseGCD := GCDDefault
	if cast.GCDCooldown != 0 {
		baseGCD = cast.GCDCooldown
	}
	return MaxDuration(GCDMin, time.Duration(float64(baseGCD)/char.CastSpeed()))
}

// Cast has finished, activate the effects of the cast.
func (cast *Cast) internalOnComplete(sim *Simulation, onCastComplete OnCastComplete) {
	if !cast.IgnoreManaCost && cast.ManaCost > 0 {
		cast.Character.AddStat(stats.Mana, -cast.ManaCost)
		cast.Character.PseudoStats.FiveSecondRuleRefreshTime = sim.CurrentTime + time.Second*5
	}

	if !cast.IgnoreCooldowns && cast.Cooldown > 0 {
		cast.Character.SetCD(cast.ActionID.CooldownID, sim.CurrentTime+cast.Cooldown)
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
}

func (simpleCast *SimpleCast) Init(sim *Simulation) {
	simpleCast.Cast.init(sim)
}

// TODO: Need to rename this. Cant call it Cast() because of conflict with field of the same name.
func (simpleCast *SimpleCast) StartCast(sim *Simulation) bool {
	return simpleCast.Cast.startCasting(sim, func(sim *Simulation, cast *Cast) {
		cast.Character.Metrics.AddCast(cast)
		if simpleCast.OnCastComplete != nil {
			simpleCast.OnCastComplete(sim, cast)
		}
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
