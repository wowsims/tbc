package core

import (
	"time"
)

func DotSnapshotFuncMagic(baseDamage float64, spellCoefficient float64) BaseDamageCalculator {
	if spellCoefficient == 0 {
		return BaseDamageFuncFlat(baseDamage)
	}

	if baseDamage == 0 {
		return func(_ *Simulation, hitEffect *SpellEffect, spell *SimpleSpellTemplate) float64 {
			totalSpellPower := hitEffect.SpellPower(spell.Character, spell)
			return totalSpellPower * spellCoefficient
		}
	} else {
		return func(_ *Simulation, hitEffect *SpellEffect, spell *SimpleSpellTemplate) float64 {
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

	// If both of these are set, will display uptime metrics for this dot.
	DebuffID AuraID

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

func (instance *SimpleSpell) ApplyDot(sim *Simulation, spell *SimpleSpellTemplate) {
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
func (hitEffect *SpellEffect) takeDotSnapshot(sim *Simulation, spell *SimpleSpellTemplate) {
	// snapshot total damage per tick, including any static damage multipliers
	hitEffect.DotInput.damagePerTick = hitEffect.DotInput.TickBaseDamage(sim, hitEffect, spell) * hitEffect.DamageMultiplier

	hitEffect.DotInput.startTime = sim.CurrentTime
	hitEffect.DotInput.RefreshDot(sim)
	hitEffect.DotInput.nextTickTime = sim.CurrentTime + hitEffect.DotInput.TickLength
	hitEffect.BeyondAOECapMultiplier = 1
}

func (hitEffect *SpellEffect) calculateDotDamage(sim *Simulation, spell *SimpleSpellTemplate) {
	damage := hitEffect.DotInput.damagePerTick

	hitEffect.determineOutcome(sim, spell, true)

	if !hitEffect.DotInput.IgnoreDamageModifiers {
		hitEffect.applyAttackerModifiers(sim, spell, !hitEffect.DotInput.TicksCanMissAndCrit, &damage)
		hitEffect.applyTargetModifiers(sim, spell, !hitEffect.DotInput.TicksCanMissAndCrit, hitEffect.BaseDamage.TargetSpellCoefficient, &damage)
	}
	hitEffect.applyResistances(sim, spell, &damage)
	hitEffect.applyOutcome(sim, spell, &damage)

	hitEffect.Damage = damage
}

// This should be called on each dot tick.
func (hitEffect *SpellEffect) afterDotTick(sim *Simulation, spell *SimpleSpellTemplate) {
	hitEffect.afterCalculations(sim, spell, true)
	hitEffect.DotInput.tickIndex++
	hitEffect.DotInput.nextTickTime = sim.CurrentTime + hitEffect.DotInput.TickLength
}

// This should be called after the final tick of the dot, or when the dot is cancelled.
func (hitEffect *SpellEffect) onDotComplete(sim *Simulation, spell *SimpleSpellTemplate) {
	// Clean up the dot object.
	hitEffect.DotInput.endTime = 0

	if hitEffect.DotInput.DebuffID != 0 {
		hitEffect.Target.AddAuraUptime(hitEffect.DotInput.DebuffID, spell.ActionID, sim.CurrentTime-hitEffect.DotInput.startTime)
	}
}
