package core

import (
	"time"
)

// A direct spell is one that does a single instance of damage once casting is
// complete, i.e. shadowbolt or fire blast.
// Note that some spell casts can have more than 1 DirectSpellEffect, e.g.
// Chain Lightning.
//
// This struct holds additional inputs beyond what a SpellEffect already contains,
// which are necessary for a direct spell damage calculation.
type DirectDamageSpellInput struct {
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
}

func (spell *MultiTargetDirectDamageSpell) Act(sim *Simulation) bool {
	return spell.startCasting(sim, func(sim *Simulation, cast *Cast) {
		for effectIdx := range spell.Effects {
			effect := &spell.Effects[effectIdx]
			effect.apply(sim, &spell.SpellCast)
		}

		sim.MetricsAggregator.AddSpellCast(&spell.SpellCast)
		spell.objectInUse = false
	})
}

type MultiTargetDirectDamageSpellTemplate struct {
	template MultiTargetDirectDamageSpell
	effects  []SpellHitEffect
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
}

func (spell *SimpleSpell) Init(sim *Simulation) {
	spell.SpellCast.init(sim)
}

func (spell *SimpleSpell) Act(sim *Simulation) bool {
	return spell.startCasting(sim, func(sim *Simulation, cast *Cast) {
		spell.apply(sim, &spell.SpellCast)
	})
}

type SpellHitEffect struct {
	SpellEffect
	DotInput    DotDamageInput
	DirectInput DirectDamageSpellInput
}

func (hitEffect *SpellHitEffect) apply(sim *Simulation, spellCast *SpellCast) {
	hitEffect.SpellEffect.beforeCalculations(sim, spellCast)

	applyNow := !hitEffect.Hit // a miss will immediately apply

	if hitEffect.Hit {
		// Only apply direct damage if it has damage. Otherwise this is a dot without direct damage.
		if hitEffect.DirectInput.MaxBaseDamage != 0 {
			hitEffect.SpellEffect.calculateDirectDamage(sim, spellCast, &hitEffect.DirectInput)
		}

		if hitEffect.DotInput.NumberOfTicks != 0 {
			hitEffect.SpellEffect.applyDot(sim, spellCast, &hitEffect.DotInput)
		} else {
			applyNow = true // no dot means we can apply results now.
		}
	}

	// Only applyNow
	if applyNow {
		hitEffect.SpellEffect.applyResultsToCast(spellCast)
		sim.MetricsAggregator.AddSpellCast(spellCast)
		spellCast.objectInUse = false
	}
	hitEffect.SpellEffect.afterCalculations(sim, spellCast)
}

type OnDamageTick func(*Simulation)

// DotDamageInput is the data needed to kick of the dot ticking in pendingActions.
//  For now the only way for a caster to track their dot is to keep a reference to the cast object
//  that started this and check the DotDamageInput.IsTicking()
type DotDamageInput struct {
	NumberOfTicks        int           // number of ticks over the whole duration
	TickLength           time.Duration // time between each tick
	TickBaseDamage       float64
	TickSpellCoefficient float64

	OnDamageTick OnDamageTick // TODO: Do we need an OnExpire?

	// Internal fields
	damagePerTick float64
	finalTickTime time.Duration
	tickIndex     int
}

func (ddi DotDamageInput) TimeRemaining(sim *Simulation) time.Duration {
	return MaxDuration(0, ddi.finalTickTime-sim.CurrentTime)
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
		panic("Damage over time spell already in use")
	}
	*newAction = template.template
}

// Takes in a cast template and returns a template, so you don't need to keep track of which things to allocate yourself.
func NewSimpleSpellTemplate(spellTemplate SimpleSpell) SimpleSpellTemplate {
	return SimpleSpellTemplate{
		template: spellTemplate,
	}
}
