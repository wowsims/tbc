package core

import (
	"fmt"
	"github.com/wowsims/tbc/sim/core/stats"
	"time"
)

// A direct spell is one that does a single instance of damage once casting is
// complete, i.e. shadowbolt or fire blast.
// Note that some spell casts can have more than 1 DirectSpellEffect, e.g.
// Chain Lightning.
//
// This struct holds additional inputs beyond what a SpellEffect already contains,
// which are necessary for a direct spell damage calculation.
type DirectDamageInput struct {
	MinBaseDamage float64
	MaxBaseDamage float64

	// Increase in damage per point of spell power (or attack power, if a physical spell).
	SpellCoefficient float64

	// Adds a fixed amount of damage to the spell, before multipliers.
	FlatDamageBonus float64
}

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

	OnBeforePeriodicDamage OnBeforePeriodicDamage // Before-calculation logic for this dot.
	OnPeriodicDamage       OnPeriodicDamage       // After-calculation logic for this dot.

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

type SpellHitEffect struct {
	SpellEffect
	DotInput    DotDamageInput
	DirectInput DirectDamageInput
	WeaponInput WeaponDamageInput
}

// SimpleSpell has a single cast and could have a dot or direct effect (or no effect)
//  A SimpleSpell without a target or effect will simply be cast and nothing else happens.
type SimpleSpell struct {
	// Embedded spell cast.
	SpellCast

	// Individual direct damage effect of this spell. Use this when there is only 1
	// effect for the spell.
	// Only one of this or Effects should be filled, not both.
	Effect SpellHitEffect

	// Individual direct damage effects of this spell. Use this for spells with
	// multiple effects, like Arcane Explosion, Chain Lightning, or Consecrate.
	Effects []SpellHitEffect

	// Maximum amount of pre-crit damage this spell is allowed to do.
	AOECap float64

	// The action currently used for the dot effects of this spell, or nil if not ticking.
	currentDotAction *PendingAction
}

// Init will call any 'OnCast' effects associated with the caster and then apply spell haste to the cast.
//  Init will panic if the spell or the GCD is still on CD.
func (spell *SimpleSpell) Init(sim *Simulation) {
	spell.SpellCast.init(sim)

	if len(spell.Effects) == 0 {
		if spell.Effect.DotInput.NumberOfTicks > 0 {
			spell.Effect.DotInput.init(&spell.SpellCast)
		}
	} else {
		for i, _ := range spell.Effects {
			if spell.Effects[i].DotInput.NumberOfTicks > 0 {
				spell.Effects[i].DotInput.init(&spell.SpellCast)
			}
		}
	}

	if spell.SpellExtras.Matches(SpellExtrasChanneled) {
		spell.AfterCastDelay += spell.GetChannelDuration()
	}
}

func (spell *SimpleSpell) GetChannelDuration() time.Duration {
	if len(spell.Effects) == 0 {
		return spell.Effect.DotInput.FullDuration()
	} else {
		return spell.Effects[0].DotInput.FullDuration()
	}
}

func (spell *SimpleSpell) GetDuration() time.Duration {
	if spell.SpellExtras.Matches(SpellExtrasChanneled) {
		return spell.CastTime + spell.GetChannelDuration()
	} else {
		return spell.CastTime
	}
}

func (spell *SimpleSpell) CastAuto(sim *Simulation) {
	cast := &spell.SpellCast
	effect := &spell.Effect
	character := spell.Character

	character.OnBeforeSpellHit(sim, cast, effect)
	effect.Target.OnBeforeSpellHit(sim, cast, effect)

	effect.Outcome = effect.WhiteHitTableResult(sim, spell)

	if effect.Landed() {
		bonusWeaponDamage := character.PseudoStats.BonusDamage + effect.BonusWeaponDamage

		dmg := 0.0
		if !effect.WeaponInput.Offhand {
			attackPower := character.stats[stats.AttackPower] + effect.BonusAttackPower + effect.BonusAttackPowerOnTarget
			dmg += character.AutoAttacks.MH.calculateWeaponDamage(sim, attackPower) + bonusWeaponDamage
		} else {
			attackPower := character.stats[stats.AttackPower] + effect.BonusAttackPower + 2*effect.BonusAttackPowerOnTarget
			dmg += character.AutoAttacks.OH.calculateWeaponDamage(sim, attackPower)*0.5 + bonusWeaponDamage
		}

		dmg += effect.WeaponInput.FlatDamageBonus
		dmg *= effect.WeaponInput.DamageMultiplier

		dmg += effect.DirectInput.FlatDamageBonus

		if effect.Outcome == OutcomeCrit {
			dmg *= spell.CritMultiplier
		} else if effect.Outcome == OutcomeGlance {
			// TODO glancing blow damage reduction is actually a range ([65%, 85%] vs. 73)
			dmg *= 0.75
		}

		// Apply damage reduction.
		dmg *= 1 - effect.Target.ArmorDamageReduction(character.stats[stats.ArmorPenetration]+effect.BonusArmorPenetration)

		// Apply all other effect multipliers.
		dmg *= effect.DamageMultiplier * effect.StaticDamageMultiplier

		effect.Damage += dmg
	}

	// since casts are currently temporary objects, this is wasteful
	effect.applyResultsToCast(cast)
	effect.afterCalculations(sim, spell)

	character.Metrics.AddSpellCast(cast)
	spell.objectInUse = false
}

func (spell *SimpleSpell) CastRefresh(sim *Simulation) {
	cast := &spell.SpellCast
	effect := &spell.Effect
	character := spell.Character

	character.OnBeforeSpellHit(sim, cast, effect)
	effect.Target.OnBeforeSpellHit(sim, cast, effect)

	if effect.hitCheck(sim, &spell.SpellCast) {
		effect.Outcome = OutcomeHit
	} else {
		effect.Outcome = OutcomeMiss
	}

	effect.afterCalculations(sim, spell)
}

func (spell *SimpleSpell) Cast(sim *Simulation) bool {
	return spell.startCasting(sim, func(sim *Simulation, cast *Cast) {
		if len(spell.Effects) == 0 {
			hitEffect := &spell.Effect
			hitEffect.beforeCalculations(sim, spell)

			if hitEffect.Landed() {
				// Weapon Damage Effects
				if hitEffect.WeaponInput.HasWeaponDamage() {
					hitEffect.calculateWeaponDamage(sim, spell)
				}

				// Direct Damage Effects
				if hitEffect.DirectInput.MaxBaseDamage != 0 {
					hitEffect.calculateDirectDamage(sim, &spell.SpellCast)
				}

				// Dot Damage Effects
				if hitEffect.DotInput.NumberOfTicks != 0 {
					hitEffect.takeDotSnapshot(sim, &spell.SpellCast)
					spell.ApplyDot(sim)
				}
			}

			hitEffect.applyResultsToCast(&spell.SpellCast)
			hitEffect.afterCalculations(sim, spell)
		} else {
			// Use a separate loop for the beforeCalculations() calls so that they all
			// come before the first afterCalculations() call. This prevents proc effects
			// on the first hit from benefitting other hits of the same spell.
			for effectIdx := range spell.Effects {
				hitEffect := &spell.Effects[effectIdx]
				hitEffect.beforeCalculations(sim, spell)
			}
			for effectIdx := range spell.Effects {
				hitEffect := &spell.Effects[effectIdx]
				if hitEffect.Landed() {
					// Weapon Damage Effects
					if hitEffect.WeaponInput.HasWeaponDamage() {
						hitEffect.calculateWeaponDamage(sim, spell)
					}
					// Direct Damage Effects
					if hitEffect.DirectInput.MaxBaseDamage != 0 {
						hitEffect.calculateDirectDamage(sim, &spell.SpellCast)
					}
					if hitEffect.DotInput.NumberOfTicks != 0 {
						hitEffect.takeDotSnapshot(sim, &spell.SpellCast)
					}
				}
			}
			spell.applyAOECap()
			// Use a separate loop for the afterCalculations() calls so all effect damage
			// is fully calculated before invoking proc callbacks.
			for effectIdx := range spell.Effects {
				hitEffect := &spell.Effects[effectIdx]
				hitEffect.applyResultsToCast(&spell.SpellCast)
				hitEffect.afterCalculations(sim, spell)
			}

			// This assumes that the effects either all have dots, or none of them do.
			if spell.Effects[0].DotInput.NumberOfTicks != 0 {
				spell.ApplyDot(sim)
			}
		}

		if spell.currentDotAction == nil {
			spell.Character.Metrics.AddSpellCast(&spell.SpellCast)
			spell.objectInUse = false
		}
	})
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

func (spell *SimpleSpell) applyAOECap() {
	if spell.AOECap == 0 {
		return
	}

	// Increased damage from crits doesn't count towards the cap, so need to
	// tally pre-crit damage.
	totalTowardsCap := 0.0
	for i, _ := range spell.Effects {
		effect := &spell.Effects[i]
		totalTowardsCap += effect.Damage / effect.BeyondAOECapMultiplier
	}

	if totalTowardsCap <= spell.AOECap {
		return
	}

	maxDamagePerHit := spell.AOECap / float64(len(spell.Effects))
	for i, _ := range spell.Effects {
		effect := &spell.Effects[i]
		damageTowardsCap := effect.Damage / effect.BeyondAOECapMultiplier
		if damageTowardsCap > maxDamagePerHit {
			effect.Damage -= damageTowardsCap - maxDamagePerHit
		}
	}
}

func (spell *SimpleSpell) Cancel(sim *Simulation) {
	spell.SpellCast.Cancel()
	if spell.currentDotAction != nil {
		spell.currentDotAction.Cancel(sim)
		spell.currentDotAction = nil
	}
}

type SimpleSpellTemplate struct {
	template SimpleSpell
	effects  []SpellHitEffect
}

func (template *SimpleSpellTemplate) Apply(newAction *SimpleSpell) {
	if newAction.objectInUse {
		panic(fmt.Sprintf("Spell (%s) already in use", newAction.ActionID))
	}
	*newAction = template.template
	newAction.Effects = template.effects
	copy(newAction.Effects, template.template.Effects)
}

// Takes in a cast template and returns a template, so you don't need to keep track of which things to allocate yourself.
func NewSimpleSpellTemplate(spellTemplate SimpleSpell) SimpleSpellTemplate {
	if len(spellTemplate.Effects) > 0 && spellTemplate.Effect.StaticDamageMultiplier != 0 {
		panic("Cannot use both Effect and Effects, pick one!")
	}

	return SimpleSpellTemplate{
		template: spellTemplate,
		effects:  make([]SpellHitEffect, len(spellTemplate.Effects)),
	}
}
