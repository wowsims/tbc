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

	// Increase in damage per point of spell power (or weapon damage, if a physical school).
	SpellCoefficient float64

	// Adds a fixed amount of damage to the spell, before multipliers.
	FlatDamageBonus float64
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

func (spell *SimpleSpell) Cast(sim *Simulation) bool {
	return spell.startCasting(sim, func(sim *Simulation, cast *Cast) {
		if len(spell.Effects) == 0 {
			hitEffect := &spell.Effect
			hitEffect.beforeCalculations(sim, spell)

			if hitEffect.Landed() {
				hitEffect.directCalculations(sim, spell)

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
					hitEffect.directCalculations(sim, spell)
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

func (spell *SimpleSpell) applyAOECap() {
	if spell.AOECap == 0 {
		return
	}

	totalTowardsCap := 0.0
	for i := range spell.Effects {
		totalTowardsCap += spell.Effects[i].RawDamage
	}

	if totalTowardsCap <= spell.AOECap {
		return
	}

	maxDamagePerHit := spell.AOECap / float64(len(spell.Effects))
	for i := range spell.Effects {
		effect := &spell.Effects[i]
		if effect.RawDamage > maxDamagePerHit {
			multiplier := effect.RawDamage / maxDamagePerHit
			effect.RawDamage = maxDamagePerHit
			effect.Damage *= multiplier
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
