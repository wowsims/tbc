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

// Callback for when a cast is finished, i.e. when the in-game castbar reaches full.
type OnCastComplete func(aura *Aura, sim *Simulation, cast *Cast)
type OnSpellCastComplete func(aura *Aura, sim *Simulation, spell *Spell)

// Callback for when a cast is finished and all its immediate effects have taken effect.
type AfterCast func(sim *Simulation, cast *Cast)

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

	SpellSchool SpellSchool
	SpellExtras SpellExtras

	// Base cost. Many effects in the game which 'reduce mana cost by X%'
	// are calculated using the base mana cost. Any effects which reduce the base
	// mana cost should be applied before setting this value, and OnCast()
	// callbacks should not modify it.
	BaseCost ResourceCost

	// Actual mana cost of the spell.
	Cost ResourceCost

	CastTime time.Duration

	// TODO: Figure out how to compute this automatically based on channel settings.
	ChannelTime time.Duration

	// Adds additional delay to the GCD after the cast is completed. This is usually
	// used for adding latency following the cast.
	AfterCastDelay time.Duration

	// Callbacks for providing additional custom behavior.
	OnCastComplete func(sim *Simulation, cast *Cast)

	// Callbacks for providing additional custom behavior.
	AfterCast AfterCast

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

func (cast *Cast) ApplyCostModifiers(curCost *float64) {
	if cast.Character.PseudoStats.NoCost {
		*curCost = 0
	} else {
		*curCost -= cast.BaseCost.Value * (1 - cast.Character.PseudoStats.CostMultiplier)
		*curCost -= cast.Character.PseudoStats.CostReduction
		*curCost = MaxFloat(0, *curCost)
	}
}
func (cast *Cast) ApplyCastTimeModifiers(dur *time.Duration) {
	if !cast.IgnoreHaste {
		*dur = time.Duration(float64(*dur) / cast.Character.CastSpeed())
	}
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

	cast.ApplyCastTimeModifiers(&cast.CastTime)
	cast.ApplyCastTimeModifiers(&cast.ChannelTime)
	cast.ApplyCostModifiers(&cast.Cost.Value)

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
func (cast *Cast) startCasting(sim *Simulation, onCastComplete func(*Simulation, *Cast)) bool {
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
		gcdCD := MaxDuration(cast.CalculatedGCD(cast.Character), cast.CastTime+cast.ChannelTime+cast.AfterCastDelay)
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
			cast.Character.newHardcastAction(sim)
		}

		if cast.Character.AutoAttacks.IsEnabled() {
			// Delay autoattacks until the cast is complete.
			cast.Character.AutoAttacks.DelayAllUntil(sim, cast.Character.Hardcast.Expires)
		}
	}

	return true
}

func (character *Character) ApplyCastSpeed(dur time.Duration) time.Duration {
	return time.Duration(float64(dur) / character.CastSpeed())
}

func (cast *Cast) CalculatedGCD(char *Character) time.Duration {
	// TODO: switch on melee or physical, to apply spell haste to GCD or not?
	//   Or does spell haste always decrease GCD (its just most non-casters dont have spell haste?)

	if cast.IgnoreHaste {
		return cast.GCD
	} else {
		return MaxDuration(GCDMin, time.Duration(float64(cast.GCD)/char.CastSpeed()))
	}
}

// Cast has finished, activate the effects of the cast.
func (cast *Cast) internalOnComplete(sim *Simulation, onCastComplete func(sim *Simulation, cast *Cast)) {
	if sim.Log != nil {
		// Hunter fake cast has no ID.
		if !cast.ActionID.SameAction(ActionID{}) {
			cast.Character.Log(sim, "Completed cast %s", cast.ActionID)
		}
	}

	if cast.Cost.Value > 0 && cast.Cost.Type == stats.Mana {
		cast.Character.SpendMana(sim, cast.Cost.Value, cast.ActionID)
		cast.Character.PseudoStats.FiveSecondRuleRefreshTime = sim.CurrentTime + time.Second*5
	}

	cast.Character.OnCastComplete(sim, cast)
	if cast.OnCastComplete != nil {
		cast.OnCastComplete(sim, cast)
	}
	onCastComplete(sim, cast)
	if cast.AfterCast != nil {
		cast.AfterCast(sim, cast)
	}
}

// A simple cast is just a cast with a callback, no calculations or damage.
type SimpleCast struct {
	// Embedded Cast
	Cast

	OnCastComplete func(*Simulation, *Cast)
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
		simpleCast.objectInUse = false
	})
}

type Hardcast struct {
	Cast          *Cast
	Expires       time.Duration
	OnComplete    func(*Simulation, *Cast)
	NewOnComplete func(*Simulation, *Target)
	Target        *Target
}

func (hc *Hardcast) OnExpire(sim *Simulation) {
	if hc.NewOnComplete != nil {
		fn := hc.NewOnComplete
		hc.NewOnComplete = nil
		fn(sim, hc.Target)
	} else {
		hc.Cast.internalOnComplete(sim, hc.OnComplete)
	}
}

// Input for constructing the CastSpell function for a spell.
type CastConfig struct {
	// Default cast values with all static effects applied.
	DefaultCast NewCast

	// Dynamic modifications for each cast.
	ModifyCast func(*Simulation, *Spell, *NewCast)

	// Ignores haste when calculating the GCD and cast time for this cast.
	IgnoreHaste bool

	// If set, this action will start a cooldown using its cooldown ID.
	// Note that the GCD CD will be activated even if this is not set.
	Cooldown time.Duration

	// Callbacks for providing additional custom behavior.
	OnCastComplete func(*Simulation, *Spell)
	AfterCast      func(*Simulation, *Spell)
}

type NewCast struct {
	// Amount of resource that will be consumed by this cast.
	Cost float64

	// The length of time the GCD will be on CD as a result of this cast.
	GCD time.Duration

	// The amount of time between the call to spell.Cast() and when the spell
	// effects are invoked.
	CastTime time.Duration

	// Additional GCD delay after the cast completes.
	ChannelTime time.Duration

	// Additional GCD delay after the cast ends. Never affected by cast speed.
	// This is typically used for latency.
	AfterCastDelay time.Duration
}

type CastFunc func(*Simulation, *Target)
type CastSuccessFunc func(*Simulation, *Target) bool

func (spell *Spell) makeCastFunc(config CastConfig, onCastComplete CastFunc) CastSuccessFunc {
	return spell.wrapCastFuncInit(config,
		spell.wrapCastFuncResources(config,
			spell.wrapCastFuncHaste(config,
				spell.wrapCastFuncGCD(config,
					spell.wrapCastFuncCooldown(config,
						spell.makeCastFuncWait(config, onCastComplete))))))
}

func (spell *Spell) ApplyCostModifiers(cost float64) float64 {
	if spell.Character.PseudoStats.NoCost {
		return 0
	} else {
		cost -= spell.BaseCost * (1 - spell.Character.PseudoStats.CostMultiplier)
		cost -= spell.Character.PseudoStats.CostReduction
		return MaxFloat(0, cost)
	}
}

func (spell *Spell) wrapCastFuncInit(config CastConfig, onCastComplete CastSuccessFunc) CastSuccessFunc {
	empty := NewCast{}
	if config.DefaultCast == empty {
		return onCastComplete
	}

	if config.ModifyCast == nil {
		return func(sim *Simulation, target *Target) bool {
			spell.CurCast = spell.DefaultCast
			return onCastComplete(sim, target)
		}
	} else {
		modifyCast := config.ModifyCast
		return func(sim *Simulation, target *Target) bool {
			spell.CurCast = spell.DefaultCast
			modifyCast(sim, spell, &spell.CurCast)
			return onCastComplete(sim, target)
		}
	}
}

func (spell *Spell) wrapCastFuncResources(config CastConfig, onCastComplete CastFunc) CastSuccessFunc {
	if spell.ResourceType == 0 || config.DefaultCast.Cost == 0 {
		if spell.ResourceType != 0 {
			panic("ResourceType set for spell " + spell.ActionID.String() + " but no cost")
		}
		if config.DefaultCast.Cost != 0 {
			panic("Cost set for spell " + spell.ActionID.String() + " but no ResourceType")
		}
		return func(sim *Simulation, target *Target) bool {
			onCastComplete(sim, target)
			return true
		}
	}

	switch spell.ResourceType {
	case stats.Mana:
		return func(sim *Simulation, target *Target) bool {
			spell.CurCast.Cost = spell.ApplyCostModifiers(spell.CurCast.Cost)
			if spell.Character.CurrentMana() < spell.CurCast.Cost {
				if sim.Log != nil {
					spell.Character.Log(sim, "Failed casting %s, not enough mana. (Current Mana = %0.03f, Mana Cost = %0.03f)",
						spell.ActionID, spell.Character.CurrentMana(), spell.CurCast.Cost)
				}
				return false
			}

			// Mana is subtracted at the end of the cast.
			onCastComplete(sim, target)
			return true
		}
	case stats.Rage:
		return func(sim *Simulation, target *Target) bool {
			spell.CurCast.Cost = spell.ApplyCostModifiers(spell.CurCast.Cost)
			if spell.Character.CurrentRage() < spell.CurCast.Cost {
				return false
			}
			spell.Character.SpendRage(sim, spell.CurCast.Cost, spell.ActionID)
			onCastComplete(sim, target)
			return true
		}
	case stats.Energy:
		return func(sim *Simulation, target *Target) bool {
			spell.CurCast.Cost = spell.ApplyCostModifiers(spell.CurCast.Cost)
			if spell.Character.CurrentEnergy() < spell.CurCast.Cost {
				return false
			}
			spell.Character.SpendEnergy(sim, spell.CurCast.Cost, spell.ActionID)
			onCastComplete(sim, target)
			return true
		}
	}

	panic("Invalid resource type")
}

func (spell *Spell) wrapCastFuncHaste(config CastConfig, onCastComplete CastFunc) CastFunc {
	if config.IgnoreHaste || (config.DefaultCast.GCD == 0 && config.DefaultCast.CastTime == 0 && config.DefaultCast.ChannelTime == 0) {
		return onCastComplete
	}

	return func(sim *Simulation, target *Target) {
		spell.CurCast.GCD = spell.Character.ApplyCastSpeed(spell.CurCast.GCD)
		spell.CurCast.CastTime = spell.Character.ApplyCastSpeed(spell.CurCast.CastTime)
		spell.CurCast.ChannelTime = spell.Character.ApplyCastSpeed(spell.CurCast.ChannelTime)

		onCastComplete(sim, target)
	}
}

func (spell *Spell) wrapCastFuncGCD(config CastConfig, onCastComplete CastFunc) CastFunc {
	if config.DefaultCast.GCD == 0 {
		return onCastComplete
	}

	return func(sim *Simulation, target *Target) {
		// By panicking if spell is on CD, we force each sim to properly check for their own CDs.
		if spell.CurCast.GCD != 0 && spell.Character.IsOnCD(GCDCooldownID, sim.CurrentTime) {
			panic(fmt.Sprintf("Trying to cast %s but GCD on cooldown for %s", spell.ActionID, spell.Character.GetRemainingCD(GCDCooldownID, sim.CurrentTime)))
		}

		gcd := spell.CurCast.GCD
		if spell.CurCast.GCD != 0 {
			gcd = MaxDuration(GCDMin, gcd)
		}

		fullCastTime := spell.CurCast.CastTime + spell.CurCast.ChannelTime + spell.CurCast.AfterCastDelay
		spell.Character.SetGCDTimer(sim, sim.CurrentTime+MaxDuration(gcd, fullCastTime))

		onCastComplete(sim, target)
	}
}

func (spell *Spell) wrapCastFuncCooldown(config CastConfig, onCastComplete CastFunc) CastFunc {
	if config.Cooldown != 0 && spell.ActionID.CooldownID == 0 {
		panic("Cooldown specified but no CooldownID!")
	}

	if config.Cooldown == 0 {
		return onCastComplete
	}

	// Store separately so the lambda doesn't capture the entire config.
	cooldownDur := config.Cooldown

	return func(sim *Simulation, target *Target) {
		// By panicking if spell is on CD, we force each sim to properly check for their own CDs.
		if spell.Character.IsOnCD(spell.ActionID.CooldownID, sim.CurrentTime) {
			panic(fmt.Sprintf("Trying to cast %s but is still on cooldown for %s", spell.ActionID, spell.Character.GetRemainingCD(spell.ActionID.CooldownID, sim.CurrentTime)))
		}

		spell.Character.SetCD(spell.ActionID.CooldownID, sim.CurrentTime+spell.CurCast.CastTime+cooldownDur)

		onCastComplete(sim, target)
	}
}

func (spell *Spell) makeCastFuncWait(config CastConfig, onCastComplete CastFunc) CastFunc {
	configOnCastComplete := config.OnCastComplete
	configAfterCast := config.AfterCast
	oldOnCastComplete1 := onCastComplete
	onCastComplete = func(sim *Simulation, target *Target) {
		spell.Character.OnSpellCastComplete(sim, spell)
		if configOnCastComplete != nil {
			configOnCastComplete(sim, spell)
		}
		oldOnCastComplete1(sim, target)
		if configAfterCast != nil {
			configAfterCast(sim, spell)
		}
	}

	if spell.ResourceType == stats.Mana && config.DefaultCast.Cost != 0 {
		oldOnCastComplete2 := onCastComplete
		onCastComplete = func(sim *Simulation, target *Target) {
			if spell.CurCast.Cost > 0 {
				spell.Character.SpendMana(sim, spell.CurCast.Cost, spell.ActionID)
				spell.Character.PseudoStats.FiveSecondRuleRefreshTime = sim.CurrentTime + time.Second*5
			}
			oldOnCastComplete2(sim, target)
		}
	}

	if config.DefaultCast.CastTime == 0 {
		return func(sim *Simulation, target *Target) {
			if sim.Log != nil {
				// Hunter fake cast has no ID.
				if !spell.ActionID.IsEmptyAction() {
					spell.Character.Log(sim, "Casting %s (Cost = %0.03f, Cast Time = %s)",
						spell.ActionID, MaxFloat(0, spell.CurCast.Cost), spell.CurCast.CastTime)
					spell.Character.Log(sim, "Completed cast %s", spell.ActionID)
				}
			}
			onCastComplete(sim, target)
		}
	} else {
		oldOnCastComplete3 := onCastComplete
		onCastComplete = func(sim *Simulation, target *Target) {
			if sim.Log != nil {
				// Hunter fake cast has no ID.
				if !spell.ActionID.SameAction(ActionID{}) {
					spell.Character.Log(sim, "Completed cast %s", spell.ActionID)
				}
			}
			oldOnCastComplete3(sim, target)
		}

		return func(sim *Simulation, target *Target) {
			if sim.Log != nil {
				spell.Character.Log(sim, "Casting %s (Cost = %0.03f, Cast Time = %s)",
					spell.ActionID, MaxFloat(0, spell.CurCast.Cost), spell.CurCast.CastTime)
			}

			// For instant-cast spells we can skip creating an aura.
			if spell.CurCast.CastTime == 0 {
				onCastComplete(sim, target)
			} else {
				spell.Character.Hardcast.Expires = sim.CurrentTime + spell.CurCast.CastTime
				spell.Character.Hardcast.NewOnComplete = onCastComplete
				spell.Character.Hardcast.Target = target

				// If hardcast and GCD happen at the same time then we don't need a separate action.
				if spell.Character.Hardcast.Expires != spell.Character.NextGCDAt() {
					spell.Character.newHardcastAction(sim)
				}

				if spell.Character.AutoAttacks.IsEnabled() {
					// Delay autoattacks until the cast is complete.
					spell.Character.AutoAttacks.DelayAllUntil(sim, spell.Character.Hardcast.Expires)
				}
			}
		}
	}
}
