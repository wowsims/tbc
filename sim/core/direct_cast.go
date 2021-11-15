package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/stats"
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
}

func (spellEffect *SpellEffect) calculateDirectDamage(sim *Simulation, spellCast *SpellCast, ddInput *DirectDamageSpellInput) {
	baseDamage := ddInput.MinBaseDamage + sim.RandomFloat("DirectSpell Base Damage")*(ddInput.MaxBaseDamage-ddInput.MinBaseDamage)

	totalSpellPower := spellCast.Character.GetStat(stats.SpellPower) + spellCast.Character.GetStat(spellCast.SpellSchool) + spellEffect.BonusSpellPower
	damageFromSpellPower := (totalSpellPower * ddInput.SpellCoefficient)

	damage := baseDamage + damageFromSpellPower

	damage *= spellEffect.DamageMultiplier

	crit := (spellCast.Character.GetStat(stats.SpellCrit) + spellEffect.BonusSpellCritRating) / (SpellCritRatingPerCritChance * 100)
	if spellCast.GuaranteedCrit || sim.RandomFloat("DirectSpell Crit") < crit {
		spellEffect.Crit = true
		damage *= spellCast.CritMultiplier
	}

	damage = calculateResists(sim, damage, spellEffect)

	spellEffect.Damage = damage
}

type DirectDamageSpellEffect struct {
	SpellEffect
	DirectDamageSpellInput
}

func (ddEffect *DirectDamageSpellEffect) apply(sim *Simulation, spellCast *SpellCast) {
	ddEffect.SpellEffect.beforeCalculations(sim, spellCast)

	if ddEffect.Hit {
		ddEffect.SpellEffect.calculateDirectDamage(sim, spellCast, &ddEffect.DirectDamageSpellInput)
	}

	// Apply results to the cast before invoking callbacks, to prevent callbacks from changing results.
	ddEffect.SpellEffect.applyResultsToCast(spellCast)
	ddEffect.SpellEffect.afterCalculations(sim, spellCast)
}

type SingleTargetDirectDamageSpell struct {
	// Embedded spell cast.
	SpellCast

	// Individual direct damage effect of this spell.
	Effect DirectDamageSpellEffect
}

func (spell *SingleTargetDirectDamageSpell) Init(sim *Simulation) {
	spell.SpellCast.init(sim)
}

func (spell *SingleTargetDirectDamageSpell) Act(sim *Simulation) bool {
	return spell.startCasting(sim, func(sim *Simulation, cast *Cast) {
		spell.Effect.apply(sim, &spell.SpellCast)
		sim.MetricsAggregator.AddSpellCast(&spell.SpellCast)
	})
}

type SingleTargetDirectDamageSpellTemplate struct {
	template SingleTargetDirectDamageSpell
}

func (template *SingleTargetDirectDamageSpellTemplate) Apply(newAction *SingleTargetDirectDamageSpell) {
	*newAction = template.template
}

// Takes in a cast template and returns a template, so you don't need to keep track of which things to allocate yourself.
func NewSingleTargetDirectDamageSpellTemplate(spellTemplate SingleTargetDirectDamageSpell) SingleTargetDirectDamageSpellTemplate {
	return SingleTargetDirectDamageSpellTemplate{
		template: spellTemplate,
	}
}

type MultiTargetDirectDamageSpell struct {
	// Embedded spell cast.
	SpellCast

	// Individual direct damage effects of this spell.
	// For most spells this will only have 1 element, but for multi-damage spells
	// like Arcane Explosion of Chain Lightning this will have multiple elements.
	Effects []DirectDamageSpellEffect
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
	})
}

type MultiTargetDirectDamageSpellTemplate struct {
	template MultiTargetDirectDamageSpell
	effects  []DirectDamageSpellEffect
}

func (template *MultiTargetDirectDamageSpellTemplate) Apply(newAction *MultiTargetDirectDamageSpell) {
	*newAction = template.template
	newAction.Effects = template.effects
	copy(newAction.Effects, template.template.Effects)
}

// Takes in a cast template and returns a template, so you don't need to keep track of which things to allocate yourself.
func NewMultiTargetDirectDamageSpellTemplate(spellTemplate MultiTargetDirectDamageSpell) MultiTargetDirectDamageSpellTemplate {
	return MultiTargetDirectDamageSpellTemplate{
		template: spellTemplate,
		effects:  make([]DirectDamageSpellEffect, len(spellTemplate.Effects)),
	}
}

type DamageOverTimeSpell struct {
	// Embedded spell cast.
	SpellCast

	// Individual direct damage effect of this spell.
	DamageOverTimeSpellEffect
}

func (spell *DamageOverTimeSpell) Init(sim *Simulation) {
	spell.SpellCast.init(sim)
}

func (spell *DamageOverTimeSpell) Act(sim *Simulation) bool {
	return spell.startCasting(sim, func(sim *Simulation, cast *Cast) {
		spell.apply(sim, &spell.SpellCast)
	})
}

type DamageOverTimeSpellEffect struct {
	SpellEffect
	DotInput    DotDamageInput
	DirectInput DirectDamageSpellInput
}

func (dotEffect *DamageOverTimeSpellEffect) apply(sim *Simulation, spellCast *SpellCast) {
	dotEffect.SpellEffect.beforeCalculations(sim, spellCast)

	if dotEffect.Hit {
		// Only apply direct damage if it has damage. Otherwise this is a dot without direct damage.
		if dotEffect.DirectInput.MaxBaseDamage != 0 {
			dotEffect.SpellEffect.calculateDirectDamage(sim, spellCast, &dotEffect.DirectInput)
		}
		dotEffect.SpellEffect.applyDot(sim, spellCast, &dotEffect.DotInput)
	} else {
		// Handle a missed cast here.
		dotEffect.SpellEffect.applyResultsToCast(spellCast)
		sim.MetricsAggregator.AddSpellCast(spellCast)
	}

	dotEffect.SpellEffect.afterCalculations(sim, spellCast)
}

type OnDamageTick func(*Simulation)

// DotDamageInput is the data needed to kick of the dot ticking in pendingActions.
//  For now the only way for a caster to track their dot is to keep a reference to the cast object
//  that started this and check the DotDamageInput.IsTicking()
type DotDamageInput struct {
	Name             string
	BaseDamage       float64
	NumberTicks      int           // total time to tick for
	TickLength       time.Duration // how often to fire OnDamageTick
	SpellCoefficient float64

	OnDamageTick OnDamageTick // TODO: Do we need an OnExpire?

	DamagePerTick float64
	FinalTickTime time.Duration
	TickIndex     int
}

func (ddi DotDamageInput) TimeRemaining(sim *Simulation) time.Duration {
	return MaxDuration(0, ddi.FinalTickTime-sim.CurrentTime)
}

func (ddi DotDamageInput) IsTicking(sim *Simulation) bool {
	return ddi.TimeRemaining(sim) != 0
}

func (spellEffect *SpellEffect) applyDot(sim *Simulation, spellCast *SpellCast, ddInput *DotDamageInput) {
	totalSpellPower := spellCast.Character.GetStat(stats.SpellPower) + spellCast.Character.GetStat(spellCast.SpellSchool) + spellEffect.BonusSpellPower
	// snapshot total damage per tick
	ddInput.DamagePerTick = ddInput.BaseDamage/float64(ddInput.NumberTicks) + totalSpellPower*ddInput.SpellCoefficient
	ddInput.FinalTickTime = sim.CurrentTime + time.Duration(ddInput.NumberTicks)*ddInput.TickLength

	pa := &PendingAction{
		NextActionAt: sim.CurrentTime + ddInput.TickLength,
	}

	pa.OnAction = func(sim *Simulation) {
		damage := ddInput.DamagePerTick
		damage = calculateResists(sim, damage, spellEffect)

		if sim.Log != nil {
			sim.Log(" %s Ticked for %0.1f\n", ddInput.Name, damage)
		}

		spellEffect.Damage += damage

		if ddInput.OnDamageTick != nil {
			ddInput.OnDamageTick(sim)
		}

		ddInput.TickIndex++
		if ddInput.TickIndex < ddInput.NumberTicks {
			// add more pending
			pa.NextActionAt = sim.CurrentTime + ddInput.TickLength
		} else {
			// TODO: This needs to be called at end of sim even if it isn't done ticking.
			//  maybe need an effect/aggregator that can be started and appended to or something.
			//  instead of calling add all at end.

			// Complete metrics and adding results etc
			spellEffect.applyResultsToCast(spellCast)
			sim.MetricsAggregator.AddSpellCast(spellCast)

			// Kills the pending action from the main run loop.
			pa.NextActionAt = NeverExpires
		}
	}

	sim.AddPendingAction(pa)
}

type DamageOverTimeSpellTemplate struct {
	template DamageOverTimeSpell
}

func (template *DamageOverTimeSpellTemplate) Apply(newAction *DamageOverTimeSpell) {
	*newAction = template.template
}

// Takes in a cast template and returns a template, so you don't need to keep track of which things to allocate yourself.
func NewDamageOverTimeSpellTemplate(spellTemplate DamageOverTimeSpell) DamageOverTimeSpellTemplate {
	return DamageOverTimeSpellTemplate{
		template: spellTemplate,
	}
}

func calculateResists(sim *Simulation, damage float64, spellEffect *SpellEffect) float64 {
	// Average Resistance (AR) = (Target's Resistance / (Caster's Level * 5)) * 0.75
	// P(x) = 50% - 250%*|x - AR| <- where X is %resisted
	// Using these stats:
	//    13.6% chance of
	//  FUTURE: handle boss resists for fights/classes that are actually impacted by that.

	resVal := sim.RandomFloat("DirectSpell Resist")
	if resVal < 0.18 { // 13% chance for 25% resist, 4% for 50%, 1% for 75%
		if resVal < 0.01 {
			spellEffect.PartialResist_3_4 = true
			return damage * 0.25
		} else if resVal < 0.05 {
			spellEffect.PartialResist_2_4 = true
			return damage * 0.5
		} else {
			spellEffect.PartialResist_1_4 = true
			return damage * 0.75
		}
	}

	return damage
}
