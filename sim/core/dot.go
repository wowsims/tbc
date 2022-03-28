package core

import (
	"time"
)

// DotDamageInput is the data needed to kick of the dot ticking in pendingActions.
//  For now the only way for a caster to track their dot is to keep a reference to the cast object
//  that started this and check the DotDamageInput.IsTicking()
type DotDamageInput struct {
	NumberOfTicks        int           // number of ticks over the whole duration
	TickLength           time.Duration // time between each tick
	TickBaseDamage       float64
	TickSpellCoefficient float64
	TicksCanMissAndCrit  bool // Allows individual ticks to hit/miss, and also crit.

	// If true, tick length will be shortened based on casting speed.
	AffectedByCastSpeed bool

	// Causes all modifications applied by callbacks to the initial damagePerTick
	// value to be ignored.
	IgnoreDamageModifiers bool

	// Whether ticks can proc spell hit effects such as Judgement of Wisdom.
	TicksProcSpellHitEffects bool

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

func (ddi *DotDamageInput) init(spellCast *SpellCast) {
	if ddi.AffectedByCastSpeed {
		ddi.TickLength = time.Duration(float64(ddi.TickLength) / spellCast.Character.CastSpeed())
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

func (spell *SimpleSpell) ApplyDot(sim *Simulation) {
	pa := sim.pendingActionPool.Get()
	pa.Priority = ActionPriorityDOT
	multiDot := len(spell.Effects) > 0

	if multiDot {
		pa.NextActionAt = sim.CurrentTime + spell.Effects[0].DotInput.TickLength
	} else {
		pa.NextActionAt = sim.CurrentTime + spell.Effect.DotInput.TickLength
	}

	pa.OnAction = func(sim *Simulation) {
		referenceHit := &spell.Effect
		if multiDot {
			referenceHit = &spell.Effects[0]
			if sim.CurrentTime == referenceHit.DotInput.nextTickTime {
				for i := range spell.Effects {
					spell.Effects[i].calculateDotDamage(sim, &spell.SpellCast)
				}
				spell.applyAOECap()
				for i := range spell.Effects {
					spell.Effects[i].afterDotTick(sim, spell)
				}
			}
		} else {
			if sim.CurrentTime == referenceHit.DotInput.nextTickTime {
				referenceHit.calculateDotDamage(sim, &spell.SpellCast)
				referenceHit.afterDotTick(sim, spell)
			}
		}

		// This assumes that all the dots have the same # of ticks and tick length.
		if referenceHit.DotInput.endTime > sim.CurrentTime {
			// Refresh action.
			pa.NextActionAt = MinDuration(referenceHit.DotInput.endTime, referenceHit.DotInput.nextTickTime)
			sim.AddPendingAction(pa)
		} else {
			pa.CleanUp(sim)
		}
	}
	pa.CleanUp = func(sim *Simulation) {
		if pa.cancelled {
			return
		}
		pa.cancelled = true
		if spell.currentDotAction != nil {
			spell.currentDotAction.cancelled = true
			spell.currentDotAction = nil
		}
		if multiDot {
			for i := range spell.Effects {
				spell.Effects[i].onDotComplete(sim, &spell.SpellCast)
			}
		} else {
			spell.Effect.onDotComplete(sim, &spell.SpellCast)
		}
		spell.Character.Metrics.AddSpellCast(&spell.SpellCast)
		spell.objectInUse = false
	}

	spell.currentDotAction = pa
	sim.AddPendingAction(pa)
}

// Snapshots a few values at the start of a dot.
func (hitEffect *SpellHitEffect) takeDotSnapshot(sim *Simulation, spellCast *SpellCast) {
	totalSpellPower := hitEffect.spellPower(spellCast.Character, spellCast)

	// snapshot total damage per tick, including any static damage multipliers
	hitEffect.DotInput.damagePerTick = (hitEffect.DotInput.TickBaseDamage + totalSpellPower*hitEffect.DotInput.TickSpellCoefficient) * hitEffect.StaticDamageMultiplier
	hitEffect.DotInput.startTime = sim.CurrentTime
	hitEffect.DotInput.RefreshDot(sim)
	hitEffect.DotInput.nextTickTime = sim.CurrentTime + hitEffect.DotInput.TickLength
	hitEffect.SpellEffect.BeyondAOECapMultiplier = 1
}

func (hitEffect *SpellHitEffect) calculateDotDamage(sim *Simulation, spellCast *SpellCast) {
	damage := hitEffect.DotInput.damagePerTick

	hitEffect.Outcome = OutcomeEmpty
	if hitEffect.DotInput.TicksCanMissAndCrit {
		if hitEffect.hitCheck(sim, spellCast) {
			hitEffect.Outcome = OutcomeHit
			if hitEffect.critCheck(sim, spellCast) {
				// TODO: Remove |=
				hitEffect.Outcome |= OutcomeCrit
			}
		} else {
			hitEffect.Outcome = OutcomeMiss
		}
	} else {
		hitEffect.Outcome = OutcomeHit
	}

	if !hitEffect.DotInput.IgnoreDamageModifiers {
		hitEffect.applyAttackerMultipliers(sim, spellCast, !hitEffect.DotInput.TicksCanMissAndCrit, &damage)
		hitEffect.applyTargetMultipliers(sim, spellCast, !hitEffect.DotInput.TicksCanMissAndCrit, &damage)
	}
	hitEffect.applyResistances(sim, spellCast, &damage)
	hitEffect.applyOutcome(sim, spellCast, &damage)

	hitEffect.SpellEffect.Damage = damage
}

// This should be called on each dot tick.
func (hitEffect *SpellHitEffect) afterDotTick(sim *Simulation, spell *SimpleSpell) {
	if sim.Log != nil {
		spell.Character.Log(sim, "%s %s. (Threat: %0.3f)", spell.ActionID, hitEffect.SpellEffect.DotResultString(), hitEffect.SpellEffect.calcThreat(&spell.SpellCast))
	}

	hitEffect.applyDotTickResultsToCast(&spell.SpellCast)

	if hitEffect.DotInput.TicksProcSpellHitEffects {
		hitEffect.SpellEffect.triggerSpellProcs(sim, spell)
	}

	spell.Character.OnPeriodicDamage(sim, &spell.SpellCast, &hitEffect.SpellEffect, hitEffect.Damage)
	hitEffect.Target.OnPeriodicDamage(sim, &spell.SpellCast, &hitEffect.SpellEffect, hitEffect.Damage)
	if hitEffect.DotInput.OnPeriodicDamage != nil {
		hitEffect.DotInput.OnPeriodicDamage(sim, &spell.SpellCast, &hitEffect.SpellEffect, hitEffect.Damage)
	}

	hitEffect.DotInput.tickIndex++
	hitEffect.DotInput.nextTickTime = sim.CurrentTime + hitEffect.DotInput.TickLength
}

// This should be called after the final tick of the dot, or when the dot is cancelled.
func (hitEffect *SpellHitEffect) onDotComplete(sim *Simulation, spellCast *SpellCast) {
	// Clean up the dot object.
	hitEffect.DotInput.endTime = 0

	if hitEffect.DotInput.DebuffID != 0 {
		hitEffect.Target.AddAuraUptime(hitEffect.DotInput.DebuffID, spellCast.ActionID, sim.CurrentTime-hitEffect.DotInput.startTime)
	}
}

func (spellEffect *SpellEffect) DotResultString() string {
	return "tick " + spellEffect.String()
}

// Only applies the results from the ticks, not the initial dot application.
func (hitEffect *SpellHitEffect) applyDotTickResultsToCast(spellCast *SpellCast) {
	if hitEffect.DotInput.TicksCanMissAndCrit {
		if hitEffect.Landed() {
			spellCast.Hits++
			if hitEffect.Outcome.Matches(OutcomeCrit) {
				spellCast.Crits++
			}

			if hitEffect.Outcome.Matches(OutcomePartial1_4) {
				spellCast.PartialResists_1_4++
			} else if hitEffect.Outcome.Matches(OutcomePartial2_4) {
				spellCast.PartialResists_2_4++
			} else if hitEffect.Outcome.Matches(OutcomePartial3_4) {
				spellCast.PartialResists_3_4++
			}
		} else {
			spellCast.Misses++
		}
	}

	spellCast.TotalDamage += hitEffect.Damage
	spellCast.TotalThreat += hitEffect.Damage * hitEffect.TotalThreatMultiplier(spellCast)
}
