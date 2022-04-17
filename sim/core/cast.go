package core

import (
	"fmt"
	"time"

	"github.com/wowsims/tbc/sim/core/stats"
)

// A cast corresponds to any action which causes the in-game castbar to be
// shown, and activates the GCD. Note that a cast can also be instant, i.e.
// the effects are applied immediately even though the GCD is still activated.

// Callback for when a cast is finished, i.e. when the in-game castbar reaches full.
type OnCastComplete func(aura *Aura, sim *Simulation, spell *Spell)

type Hardcast struct {
	Expires    time.Duration
	OnComplete func(*Simulation, *Target)
	Target     *Target
}

func (hc *Hardcast) OnExpire(sim *Simulation) {
	hc.OnComplete(sim, hc.Target)
}

// Input for constructing the CastSpell function for a spell.
type CastConfig struct {
	// Default cast values with all static effects applied.
	DefaultCast Cast

	// Dynamic modifications for each cast.
	ModifyCast func(*Simulation, *Spell, *Cast)

	// Ignores haste when calculating the GCD and cast time for this cast.
	IgnoreHaste bool

	// If set, this action will start a cooldown using its cooldown ID.
	// Note that the GCD CD will be activated even if this is not set.
	Cooldown time.Duration

	// Secondary cooldown ID, used for shared cooldowns.
	SharedCooldownID CooldownID

	// Duration of secondary cooldown.
	SharedCooldown time.Duration

	// Callbacks for providing additional custom behavior.
	OnCastComplete func(*Simulation, *Spell)
	AfterCast      func(*Simulation, *Spell)

	DisableCallbacks bool
}

type Cast struct {
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
						spell.wrapCastFuncSharedCooldown(config,
							spell.makeCastFuncWait(config, onCastComplete)))))))
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
	empty := Cast{}
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

func (spell *Spell) wrapCastFuncSharedCooldown(config CastConfig, onCastComplete CastFunc) CastFunc {
	if config.SharedCooldown != 0 && config.SharedCooldownID == 0 {
		panic("SharedCooldown specified but no SharedCooldownID!")
	}

	if config.SharedCooldown == 0 {
		return onCastComplete
	}

	// Store separately so the lambda doesn't capture the entire config.
	cooldownDur := config.SharedCooldown
	cooldownID := config.SharedCooldownID

	return func(sim *Simulation, target *Target) {
		// By panicking if spell is on CD, we force each sim to properly check for their own CDs.
		if spell.Character.IsOnCD(cooldownID, sim.CurrentTime) {
			panic(fmt.Sprintf("Trying to cast %s but is still on shared cooldown for %s", spell.ActionID, spell.Character.GetRemainingCD(cooldownID, sim.CurrentTime)))
		}

		spell.Character.SetCD(cooldownID, sim.CurrentTime+spell.CurCast.CastTime+cooldownDur)

		onCastComplete(sim, target)
	}
}

func (spell *Spell) makeCastFuncWait(config CastConfig, onCastComplete CastFunc) CastFunc {
	if !config.DisableCallbacks {
		configOnCastComplete := config.OnCastComplete
		configAfterCast := config.AfterCast
		oldOnCastComplete1 := onCastComplete
		onCastComplete = func(sim *Simulation, target *Target) {
			spell.Character.OnCastComplete(sim, spell)
			if configOnCastComplete != nil {
				configOnCastComplete(sim, spell)
			}
			oldOnCastComplete1(sim, target)
			if configAfterCast != nil {
				configAfterCast(sim, spell)
			}
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
				spell.Character.Hardcast.OnComplete = onCastComplete
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
