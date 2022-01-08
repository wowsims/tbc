package core

import (
	"fmt"
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

	// Increase in damage per point of spell power.
	SpellCoefficient float64

	// Adds a fixed amount of damage to the spell, before multipliers.
	FlatDamageBonus float64
}

type MultiTargetDirectDamageSpell struct {
	// Embedded spell cast.
	SpellCast

	// Individual direct damage effects of this spell.
	// For most spells this will only have 1 element, but for multi-damage spells
	// like Arcane Explosion of Chain Lightning this will have multiple elements.
	Effects []SpellHitEffect
}

func (spell *MultiTargetDirectDamageSpell) Init(sim *Simulation) {
	spell.SpellCast.init(sim)
	for effectIdx := range spell.Effects {
		if spell.Effects[effectIdx].DotInput.NumberOfTicks > 0 {
			spell.Effects[effectIdx].DotInput.init(&spell.SpellCast)
		}
	}
}

// TODO: If there are multiple Effects.DotEffect then each one will apply to the metrics (creating too high of a resulting DPS)
//  To handle this we would need to add a "OnDotComplete" callback to aggregate all the dots together into a single metrics.
//  Note: This might only apply to consecrate.
func (spell *MultiTargetDirectDamageSpell) Cast(sim *Simulation) bool {
	return spell.startCasting(sim, func(sim *Simulation, cast *Cast) {
		for effectIdx := range spell.Effects {
			effect := &spell.Effects[effectIdx]
			effect.apply(sim, &spell.SpellCast, false)
		}
		// Manually apply all effects at once at the end of all the apply
		cast.Character.Metrics.AddSpellCast(&spell.SpellCast)
		spell.objectInUse = false
	})
}

type MultiTargetDirectDamageSpellTemplate struct {
	template MultiTargetDirectDamageSpell
	effects  []SpellHitEffect // cached effects to use on the actual cast so we don't mutate the template
}

func (template *MultiTargetDirectDamageSpellTemplate) Apply(newAction *MultiTargetDirectDamageSpell) {
	if newAction.objectInUse {
		panic("Multi target spell already in use")
	}
	*newAction = template.template
	newAction.Effects = template.effects
	copy(newAction.Effects, template.template.Effects)
}

// Takes in a cast template and returns a template, so you don't need to keep track of which things to allocate yourself.
func NewMultiTargetDirectDamageSpellTemplate(spellTemplate MultiTargetDirectDamageSpell) MultiTargetDirectDamageSpellTemplate {
	return MultiTargetDirectDamageSpellTemplate{
		template: spellTemplate,
		effects:  make([]SpellHitEffect, len(spellTemplate.Effects)),
	}
}

// SimpleSpell has a single cast and could have a dot or direct effect (or no effect)
//  A SimpleSpell without a target or effect will simply be cast and nothing else happens.
type SimpleSpell struct {
	// Embedded spell cast.
	SpellCast

	// Individual direct damage effect of this spell.
	SpellHitEffect

	IsChannel bool
}

// Init will call any 'OnCast' effects associated with the caster and then apply spell haste to the cast.
//  Init will panic if the spell or the GCD is still on CD.
func (spell *SimpleSpell) Init(sim *Simulation) {
	spell.SpellCast.init(sim)
	if spell.SpellHitEffect.DotInput.NumberOfTicks > 0 {
		spell.SpellHitEffect.DotInput.init(&spell.SpellCast)
	}
}

func (spell *SimpleSpell) GetDuration() time.Duration {
	if spell.IsChannel {
		return spell.SpellHitEffect.DotInput.FullDuration()
	} else {
		return spell.CastTime
	}
}

func (spell *SimpleSpell) Cast(sim *Simulation) bool {
	return spell.startCasting(sim, func(sim *Simulation, cast *Cast) {
		spell.apply(sim, &spell.SpellCast, true)
	})
}

func (spell *SimpleSpell) Cancel(sim *Simulation) {
	spell.SpellCast.Cancel()
	spell.SpellHitEffect.cancel(sim)
}

type SpellHitEffect struct {
	SpellEffect
	DotInput    DotDamageInput
	DirectInput DirectDamageInput
}

// applies the hit/miss/dmg effects to the spellCast
//  If applyMetrics is false it will not apply to the sim.MetricsAggregator.. This is to support collecting multiple SpellHitEffects (like in Multi)
//  If there is a dot effect the damage will be applied to the SpellCast on each tick and on expire added to sim.MetricsAggregator.
func (hitEffect *SpellHitEffect) apply(sim *Simulation, spellCast *SpellCast, applyMetrics bool) {
	hitEffect.beforeCalculations(sim, spellCast)

	applyNow := !hitEffect.Hit // a miss will immediately apply

	if hitEffect.Hit {
		// Only apply direct damage if it has damage. Otherwise this is a dot without direct damage.
		if hitEffect.DirectInput.MaxBaseDamage != 0 {
			hitEffect.calculateDirectDamage(sim, spellCast, &hitEffect.DirectInput)
		}

		if hitEffect.DotInput.NumberOfTicks != 0 {
			hitEffect.applyDot(sim, spellCast, &hitEffect.DotInput)
		} else {
			applyNow = true // no dot means we can apply results now.
		}
	}

	// Only applyNow if there is no dot ticking
	if applyNow {
		hitEffect.applyResultsToCast(spellCast)
		if applyMetrics {
			spellCast.Character.Metrics.AddSpellCast(spellCast)
		}
		spellCast.objectInUse = false
	}
	hitEffect.afterCalculations(sim, spellCast)
}

func (hitEffect *SpellHitEffect) cancel(sim *Simulation) {
	if hitEffect.DotInput.currentDotAction != nil {
		hitEffect.DotInput.currentDotAction.Cancel(sim)
	}
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
	SpellID  int32

	// Internal fields
	startTime     time.Duration
	finalTickTime time.Duration
	damagePerTick float64
	tickIndex     int

	// The action currently used for this dot, or nil if not ticking.
	currentDotAction *PendingAction
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
	return MaxDuration(0, ddi.finalTickTime-sim.CurrentTime)
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
	return (ddi.finalTickTime != 0 && ddi.tickIndex < ddi.NumberOfTicks)
}

type SimpleSpellTemplate struct {
	template SimpleSpell
}

func (template *SimpleSpellTemplate) Apply(newAction *SimpleSpell) {
	if newAction.objectInUse {
		panic(fmt.Sprintf("Damage over time spell (%s) already in use", newAction.Name))
	}
	*newAction = template.template
}

// Takes in a cast template and returns a template, so you don't need to keep track of which things to allocate yourself.
func NewSimpleSpellTemplate(spellTemplate SimpleSpell) SimpleSpellTemplate {
	return SimpleSpellTemplate{
		template: spellTemplate,
	}
}
