package core

import (
	"time"
)

func DotSnapshotFuncMagic(baseDamage float64, spellCoefficient float64) BaseDamageCalculator {
	if spellCoefficient == 0 {
		return BaseDamageFuncFlat(baseDamage)
	}

	if baseDamage == 0 {
		return func(_ *Simulation, hitEffect *SpellEffect, spell *Spell) float64 {
			totalSpellPower := hitEffect.SpellPower(spell.Character, spell)
			return totalSpellPower * spellCoefficient
		}
	} else {
		return func(_ *Simulation, hitEffect *SpellEffect, spell *Spell) float64 {
			totalSpellPower := hitEffect.SpellPower(spell.Character, spell)
			return baseDamage + totalSpellPower*spellCoefficient
		}
	}
}

// DotDamageInput is the data needed to kick of the dot ticking in pendingActions.
//  For now the only way for a caster to track their dot is to keep a reference to the cast object
//  that started this and check the DotDamageInput.IsTicking()
type DotDamageInput struct {
	NumberOfTicks       int           // number of ticks over the whole duration
	TickLength          time.Duration // time between each tick
	TicksCanMissAndCrit bool          // Allows individual ticks to hit/miss, and also crit.
	TickBaseDamage      BaseDamageCalculator

	// If true, tick length will be shortened based on casting speed.
	AffectedByCastSpeed bool

	// Causes all modifications applied by callbacks to the initial damagePerTick
	// value to be ignored.
	IgnoreDamageModifiers bool

	// Whether ticks can proc spell hit effects such as Judgement of Wisdom.
	TicksProcSpellEffects bool

	OnPeriodicDamage OnPeriodicDamage // After-calculation logic for this dot.

	// If this is set, will display uptime metrics for this dot.
	Aura *Aura

	// Internal fields
	startTime     time.Duration
	endTime       time.Duration
	damagePerTick float64
	tickIndex     int
	nextTickTime  time.Duration
}

func (ddi *DotDamageInput) init(spell *SpellCast) {
	if ddi.AffectedByCastSpeed {
		ddi.TickLength = time.Duration(float64(ddi.TickLength) / spell.Character.CastSpeed())
	}
}

// DamagePerTick returns the cached damage per tick on the spell.
func (ddi DotDamageInput) DamagePerTick() float64 {
	return ddi.damagePerTick
}

func (ddi DotDamageInput) FullDuration() time.Duration {
	return ddi.TickLength * time.Duration(ddi.NumberOfTicks)
}

func (ddi DotDamageInput) TimeRemaining(sim *Simulation) time.Duration {
	return MaxDuration(0, ddi.endTime-sim.CurrentTime)
}

// Returns the remaining number of times this dot is expected to tick, assuming
// it lasts for its full duration.
func (ddi DotDamageInput) RemainingTicks() int {
	return ddi.NumberOfTicks - ddi.tickIndex
}

// Returns the amount of additional damage this dot is expected to do, assuming
// it lasts for its full duration.
func (ddi DotDamageInput) RemainingDamage() float64 {
	return float64(ddi.RemainingTicks()) * ddi.DamagePerTick()
}

func (ddi DotDamageInput) IsTicking(sim *Simulation) bool {
	// It is possible that both cast and tick are to happen at the same time.
	//  In this case the dot "time remaining" will be 0 but there will be ticks left.
	//  If a DOT misses then it will have NumberOfTicks set but never have been started.
	//  So the case of 'has a final tick time but its now, but it has ticks remaining' looks like this.
	return (ddi.endTime != 0 && ddi.tickIndex < ddi.NumberOfTicks)
}

func (ddi *DotDamageInput) SetTickDamage(newDamage float64) {
	ddi.damagePerTick = newDamage
}

// Restarts the dot with the same number of ticks / duration as it started with.
// Note that this does NOT change nextTickTime.
func (ddi *DotDamageInput) RefreshDot(sim *Simulation) {
	ddi.endTime = sim.CurrentTime + time.Duration(ddi.NumberOfTicks)*ddi.TickLength
	ddi.tickIndex = 0
}

func (instance *SimpleSpell) ApplyDot(sim *Simulation, spell *Spell) {
	pa := sim.pendingActionPool.Get()
	pa.Priority = ActionPriorityDOT
	multiDot := len(instance.Effects) > 0

	if multiDot {
		pa.NextActionAt = sim.CurrentTime + instance.Effects[0].DotInput.TickLength
	} else {
		pa.NextActionAt = sim.CurrentTime + instance.Effect.DotInput.TickLength
	}

	pa.OnAction = func(sim *Simulation) {
		referenceHit := &instance.Effect
		if multiDot {
			referenceHit = &instance.Effects[0]
			if sim.CurrentTime == referenceHit.DotInput.nextTickTime {
				for i := range instance.Effects {
					instance.Effects[i].calculateDotDamage(sim, spell)
				}
				instance.applyAOECap()
				for i := range instance.Effects {
					instance.Effects[i].afterDotTick(sim, spell)
				}
			}
		} else {
			if sim.CurrentTime == referenceHit.DotInput.nextTickTime {
				referenceHit.calculateDotDamage(sim, spell)
				referenceHit.afterDotTick(sim, spell)
			}
		}

		// This assumes that all the dots have the same # of ticks and tick length.
		if referenceHit.DotInput.endTime > sim.CurrentTime {
			// Refresh action.
			pa.NextActionAt = MinDuration(referenceHit.DotInput.endTime, referenceHit.DotInput.nextTickTime)
			sim.AddPendingAction(pa)
		} else {
			pa.Cancel(sim)
		}
	}
	pa.CleanUp = func(sim *Simulation) {
		if instance.currentDotAction != nil {
			instance.currentDotAction.cancelled = true
			instance.currentDotAction = nil
		}
		if multiDot {
			for i := range instance.Effects {
				instance.Effects[i].onDotComplete(sim, spell)
			}
		} else {
			instance.Effect.onDotComplete(sim, spell)
		}
		instance.objectInUse = false
	}

	instance.currentDotAction = pa
	sim.AddPendingAction(pa)
}

// Snapshots a few values at the start of a dot.
func (hitEffect *SpellEffect) takeDotSnapshot(sim *Simulation, spell *Spell) {
	// snapshot total damage per tick, including any static damage multipliers
	hitEffect.DotInput.damagePerTick = hitEffect.DotInput.TickBaseDamage(sim, hitEffect, spell) * hitEffect.DamageMultiplier

	hitEffect.DotInput.startTime = sim.CurrentTime
	hitEffect.DotInput.RefreshDot(sim)
	hitEffect.DotInput.nextTickTime = sim.CurrentTime + hitEffect.DotInput.TickLength
	hitEffect.BeyondAOECapMultiplier = 1
}

func (hitEffect *SpellEffect) calculateDotDamage(sim *Simulation, spell *Spell) {
	damage := hitEffect.DotInput.damagePerTick

	if !hitEffect.DotInput.IgnoreDamageModifiers {
		hitEffect.applyAttackerModifiers(sim, spell, !hitEffect.DotInput.TicksCanMissAndCrit, &damage)
		hitEffect.applyTargetModifiers(sim, spell, !hitEffect.DotInput.TicksCanMissAndCrit, hitEffect.BaseDamage.TargetSpellCoefficient, &damage)
	}
	hitEffect.applyResistances(sim, spell, &damage)
	hitEffect.determineOutcome(sim, spell, true)
	hitEffect.applyOutcome(sim, spell, &damage)

	hitEffect.Damage = damage
}

// This should be called on each dot tick.
func (hitEffect *SpellEffect) afterDotTick(sim *Simulation, spell *Spell) {
	hitEffect.afterCalculations(sim, spell, true)
	hitEffect.DotInput.tickIndex++
	hitEffect.DotInput.nextTickTime = sim.CurrentTime + hitEffect.DotInput.TickLength
}

// This should be called after the final tick of the dot, or when the dot is cancelled.
func (hitEffect *SpellEffect) onDotComplete(sim *Simulation, spell *Spell) {
	// Clean up the dot object.
	hitEffect.DotInput.endTime = 0

	if hitEffect.DotInput.Aura != nil {
		hitEffect.DotInput.Aura.metrics.Uptime += sim.CurrentTime - hitEffect.DotInput.startTime
	}
}

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

	lastTickTime time.Duration
}

func (dot *Dot) Apply(sim *Simulation) {
	if dot.Aura.IsActive() {
		dot.Aura.Deactivate(sim)
	}

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
				dot.tickFn()
			}
		},
		CleanUp: func(sim *Simulation) {
			// In certain cases, the last tick and the dot aura expiration can happen in
			// different orders, so we might need to apply the last tick.
			if dot.tickAction.NextActionAt == sim.CurrentTime {
				if dot.lastTickTime != sim.CurrentTime {
					dot.lastTickTime = sim.CurrentTime
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