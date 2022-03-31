package core

import (
	"fmt"
	"time"
)

// SimpleSpell has a single cast and could have a dot or direct effect (or no effect)
//  A SimpleSpell without a target or effect will simply be cast and nothing else happens.
type SimpleSpell struct {
	// Embedded spell cast.
	SpellCast

	// Individual direct damage effect of this spell. Use this when there is only 1
	// effect for the spell.
	// Only one of this or Effects should be filled, not both.
	Effect SpellEffect

	// Individual direct damage effects of this spell. Use this for spells with
	// multiple effects, like Arcane Explosion, Chain Lightning, or Consecrate.
	Effects []SpellEffect

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
		spellCast := &spell.SpellCast
		if len(spell.Effects) == 0 {
			hitEffect := &spell.Effect
			hitEffect.determineOutcome(sim, spellCast, spell)

			if hitEffect.Landed() {
				hitEffect.directCalculations(sim, spellCast)

				// Dot Damage Effects
				if hitEffect.DotInput.NumberOfTicks != 0 {
					hitEffect.takeDotSnapshot(sim, spellCast)
					spell.ApplyDot(sim)
				}
			}

			hitEffect.applyResultsToCast(spellCast)
			hitEffect.afterCalculations(sim, spellCast)
		} else {
			// Use a separate loop for the beforeCalculations() calls so that they all
			// come before the first afterCalculations() call. This prevents proc effects
			// on the first hit from benefitting other hits of the same spell.
			for effectIdx := range spell.Effects {
				hitEffect := &spell.Effects[effectIdx]
				hitEffect.determineOutcome(sim, spellCast, spell)
			}
			for effectIdx := range spell.Effects {
				hitEffect := &spell.Effects[effectIdx]
				if hitEffect.Landed() {
					hitEffect.directCalculations(sim, spellCast)
					if hitEffect.DotInput.NumberOfTicks != 0 {
						hitEffect.takeDotSnapshot(sim, spellCast)
					}
				}
			}

			// TODO: Reenable this when spell code is cleaned up.
			//spell.applyAOECap()

			// Use a separate loop for the afterCalculations() calls so all effect damage
			// is fully calculated before invoking proc callbacks.
			for effectIdx := range spell.Effects {
				hitEffect := &spell.Effects[effectIdx]
				hitEffect.applyResultsToCast(spellCast)
				hitEffect.afterCalculations(sim, spellCast)
			}

			// This assumes that the effects either all have dots, or none of them do.
			if spell.Effects[0].DotInput.NumberOfTicks != 0 {
				spell.ApplyDot(sim)
			}
		}

		if spell.currentDotAction == nil {
			spell.Character.Metrics.AddSpellCast(spellCast)
			spell.objectInUse = false
		}
	})
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

type SpellMetrics struct {
	// Metric totals for this spell, for the current iteration.
	Casts              int32
	Misses             int32
	Hits               int32
	Crits              int32
	Dodges             int32
	Glances            int32
	Parries            int32
	Blocks             int32
	PartialResists_1_4 int32   // 1/4 of the spell was resisted
	PartialResists_2_4 int32   // 2/4 of the spell was resisted
	PartialResists_3_4 int32   // 3/4 of the spell was resisted
	TotalDamage        float64 // Damage done by this cast.
	TotalThreat        float64 // Threat generated by this cast.
}

// TODO: Rename to 'Spell' when we're done with the refactoring.
type SimpleSpellTemplate struct {
	// ID for the action.
	ActionID

	// The character performing this action.
	Character *Character

	SpellSchool SpellSchool
	SpellExtras SpellExtras

	SpellMetrics

	modifyCast ModifySpellCast

	// Templates for creating new casts of this spell.
	Template SimpleSpell
	effects  []SpellEffect

	// Current instantiation of this spell. Can only be casting 1 instance of this spell at a time.
	Instance SimpleSpell
}

func (template *SimpleSpellTemplate) Apply(newAction *SimpleSpell) {
	if newAction.objectInUse {
		panic(fmt.Sprintf("Spell (%s) already in use", newAction.ActionID))
	}
	*newAction = template.Template
	newAction.Effects = template.effects
	copy(newAction.Effects, template.Template.Effects)
}

func (spell *SimpleSpellTemplate) reset(_ *Simulation) {
	spell.SpellMetrics = SpellMetrics{}
}

func (spell *SimpleSpellTemplate) doneIteration() {
	spell.Character.Metrics.addSpell(spell)
}

func (spell *SimpleSpellTemplate) Cast(sim *Simulation, target *Target) bool {
	// Initialize cast from precomputed template.
	instance := &spell.Instance
	spell.Apply(instance)

	if spell.modifyCast != nil {
		spell.modifyCast(sim, target, instance)
	}

	instance.Init(sim)
	return instance.Cast(sim)
}

type ModifySpellCast func(*Simulation, *Target, *SimpleSpell)

func ModifyCastAssignTarget(_ *Simulation, target *Target, instance *SimpleSpell) {
	instance.Effect.Target = target
}

type SpellConfig struct {
	Template SimpleSpell

	ModifyCast ModifySpellCast
}

// Registers a new spell to the character. Returns the newly created spell.
func (character *Character) RegisterSpell(config SpellConfig) *SimpleSpellTemplate {
	if len(config.Template.Effects) > 0 && config.Template.Effect.DamageMultiplier != 0 {
		panic("Cannot use both Effect and Effects, pick one!")
	}
	config.Template.Character = character

	spell := &SimpleSpellTemplate{
		ActionID:    config.Template.ActionID,
		Character:   character,
		SpellSchool: config.Template.SpellSchool,
		SpellExtras: config.Template.SpellExtras,

		modifyCast: config.ModifyCast,

		Template: config.Template,
		effects:  make([]SpellEffect, len(config.Template.Effects)),
	}

	character.Spellbook = append(character.Spellbook, spell)

	return spell
}
