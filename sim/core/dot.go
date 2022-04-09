package core

import (
	"time"
)

func (unit *Unit) NewDotAura(auraLabel string, actionID ActionID) *Aura {
	return unit.GetOrRegisterAura(&Aura{
		Label:    auraLabel,
		ActionID: actionID,
		Duration: NeverExpires,
	})
}

type TickEffects func(*Simulation, *Spell) func()

type Dot struct {
	Spell *Spell

	// Embed Aura so we can use IsActive/Refresh/etc directly.
	*Aura

	NumberOfTicks int           // number of ticks over the whole duration
	TickLength    time.Duration // time between each tick

	// If true, tick length will be shortened based on casting speed.
	AffectedByCastSpeed bool

	TickEffects TickEffects

	tickFn     func()
	tickAction *PendingAction
	tickPeriod time.Duration

	// Number of ticks since last call to Apply().
	TickCount int

	lastTickTime time.Duration
}

func (dot *Dot) Apply(sim *Simulation) {
	if dot.Aura.IsActive() {
		dot.Aura.Deactivate(sim)
	}

	dot.TickCount = 0
	if dot.AffectedByCastSpeed {
		castSpeed := dot.Spell.Character.CastSpeed()
		dot.tickPeriod = time.Duration(float64(dot.TickLength) / castSpeed)
		dot.Aura.Duration = dot.tickPeriod * time.Duration(dot.NumberOfTicks)
	}
	dot.Aura.Activate(sim)
}

// Call this after manually changing NumberOfTicks or TickLength.
func (dot *Dot) RecomputeAuraDuration() {
	if dot.AffectedByCastSpeed {
		castSpeed := dot.Spell.Character.CastSpeed()
		dot.tickPeriod = time.Duration(float64(dot.TickLength) / castSpeed)
		dot.Aura.Duration = dot.tickPeriod * time.Duration(dot.NumberOfTicks)
	} else {
		dot.tickPeriod = dot.TickLength
		dot.Aura.Duration = dot.tickPeriod * time.Duration(dot.NumberOfTicks)
	}
}

func NewDot(config Dot) *Dot {
	dot := &Dot{}
	*dot = config

	basePeriodicOptions := PeriodicActionOptions{
		OnAction: func(sim *Simulation) {
			if dot.lastTickTime != sim.CurrentTime {
				dot.lastTickTime = sim.CurrentTime
				dot.TickCount++
				dot.tickFn()
			}
		},
		CleanUp: func(sim *Simulation) {
			// In certain cases, the last tick and the dot aura expiration can happen in
			// different orders, so we might need to apply the last tick.
			if dot.tickAction.NextActionAt == sim.CurrentTime {
				if dot.lastTickTime != sim.CurrentTime {
					dot.lastTickTime = sim.CurrentTime
					dot.TickCount++
					dot.tickFn()
				}
			}
		},
	}

	dot.tickPeriod = dot.TickLength
	dot.Aura.Duration = dot.TickLength * time.Duration(dot.NumberOfTicks)

	dot.Aura.OnGain = func(aura *Aura, sim *Simulation) {
		dot.tickFn = dot.TickEffects(sim, dot.Spell)

		periodicOptions := basePeriodicOptions
		periodicOptions.Period = dot.tickPeriod
		dot.tickAction = NewPeriodicAction(sim, periodicOptions)
		sim.AddPendingAction(dot.tickAction)
	}
	dot.Aura.OnExpire = func(aura *Aura, sim *Simulation) {
		dot.tickAction.Cancel(sim)
		dot.tickAction = nil
	}

	return dot
}

func TickFuncSnapshot(target *Target, baseEffect SpellEffect) TickEffects {
	return func(sim *Simulation, spell *Spell) func() {
		snapshotEffect := baseEffect
		snapshotEffect.Target = target
		baseDamage := snapshotEffect.calculateBaseDamage(sim, spell) * snapshotEffect.DamageMultiplier
		snapshotEffect.DamageMultiplier = 1
		snapshotEffect.BaseDamage = BaseDamageConfigFlat(baseDamage)

		effectsFunc := ApplyEffectFuncDirectDamage(snapshotEffect)
		return func() {
			effectsFunc(sim, target, spell)
		}
	}
}
func TickFuncAOESnapshot(sim *Simulation, baseEffect SpellEffect) TickEffects {
	return func(sim *Simulation, spell *Spell) func() {
		target := sim.GetPrimaryTarget()
		snapshotEffect := baseEffect
		snapshotEffect.Target = target
		baseDamage := snapshotEffect.calculateBaseDamage(sim, spell) * snapshotEffect.DamageMultiplier
		snapshotEffect.DamageMultiplier = 1
		snapshotEffect.BaseDamage = BaseDamageConfigFlat(baseDamage)

		effectsFunc := ApplyEffectFuncAOEDamage(sim, snapshotEffect)
		return func() {
			effectsFunc(sim, target, spell)
		}
	}
}

func TickFuncApplyEffects(effectsFunc ApplySpellEffects) TickEffects {
	return func(sim *Simulation, spell *Spell) func() {
		return func() {
			effectsFunc(sim, sim.GetPrimaryTarget(), spell)
		}
	}
}
