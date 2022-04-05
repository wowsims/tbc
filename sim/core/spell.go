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

func (instance *SimpleSpell) Cast(sim *Simulation, target *Target, spell *Spell) bool {
	return instance.startCasting(sim, func(sim *Simulation, cast *Cast) {
		spell.Casts++
		spell.MostRecentCost = cast.Cost.Value
		spell.MostRecentBaseCost = cast.BaseCost.Value

		if spell.ApplyEffects != nil {
			spell.ApplyEffects(sim, target, spell)
			return
		}

		if len(instance.Effects) == 0 {
			hitEffect := &instance.Effect
			hitEffect.determineOutcome(sim, spell, false)

			if hitEffect.Landed() {
				hitEffect.directCalculations(sim, spell)

				// Dot Damage Effects
				if hitEffect.DotInput.NumberOfTicks != 0 {
					hitEffect.takeDotSnapshot(sim, spell)
					instance.ApplyDot(sim, spell)
				}
			}

			hitEffect.afterCalculations(sim, spell, false)
		} else {
			// Use a separate loop for the beforeCalculations() calls so that they all
			// come before the first afterCalculations() call. This prevents proc effects
			// on the first hit from benefitting other hits of the same spell.
			for effectIdx := range instance.Effects {
				hitEffect := &instance.Effects[effectIdx]
				hitEffect.determineOutcome(sim, spell, false)
				if hitEffect.Landed() {
					hitEffect.directCalculations(sim, spell)
					if hitEffect.DotInput.NumberOfTicks != 0 {
						hitEffect.takeDotSnapshot(sim, spell)
					}
				}
			}

			// TODO: Reenable this when spell code is cleaned up.
			//spell.applyAOECap()

			// Use a separate loop for the afterCalculations() calls so all effect damage
			// is fully calculated before invoking proc callbacks.
			for effectIdx := range instance.Effects {
				instance.Effects[effectIdx].afterCalculations(sim, spell, false)
			}

			// This assumes that the effects either all have dots, or none of them do.
			if instance.Effects[0].DotInput.NumberOfTicks != 0 {
				instance.ApplyDot(sim, spell)
			}
		}

		if instance.currentDotAction == nil {
			instance.objectInUse = false
		}
	})
}

func (instance *SimpleSpell) applyAOECap() {
	if instance.AOECap == 0 {
		return
	}

	// Increased damage from crits doesn't count towards the cap, so need to
	// tally pre-crit damage.
	totalTowardsCap := 0.0
	for i, _ := range instance.Effects {
		effect := &instance.Effects[i]
		totalTowardsCap += effect.Damage / effect.BeyondAOECapMultiplier
	}

	if totalTowardsCap <= instance.AOECap {
		return
	}

	maxDamagePerHit := instance.AOECap / float64(len(instance.Effects))
	for i, _ := range instance.Effects {
		effect := &instance.Effects[i]
		damageTowardsCap := effect.Damage / effect.BeyondAOECapMultiplier
		if damageTowardsCap > maxDamagePerHit {
			effect.Damage -= damageTowardsCap - maxDamagePerHit
		}
	}
}

func (instance *SimpleSpell) Cancel(sim *Simulation) {
	instance.SpellCast.Cancel()
	if instance.currentDotAction != nil {
		instance.currentDotAction.Cancel(sim)
		instance.currentDotAction = nil
	}
}

type ModifySpellCast func(*Simulation, *Target, *SimpleSpell)
type ApplySpellEffects func(*Simulation, *Target, *Spell)

type SpellConfig struct {
	Template SimpleSpell

	ModifyCast ModifySpellCast

	//CastFn SpellCast
	ApplyEffects ApplySpellEffects
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

type Spell struct {
	// ID for the action.
	ActionID

	// The character performing this action.
	Character *Character

	SpellSchool SpellSchool
	SpellExtras SpellExtras

	SpellMetrics

	ModifyCast   ModifySpellCast
	ApplyEffects ApplySpellEffects

	//castFn SpellCast

	// The amount of resource spent by the most recent cast of this spell.
	// TODO: Find a way to remove this later, as its a bit hacky.
	MostRecentBaseCost float64
	MostRecentCost     float64

	// Templates for creating new casts of this spell.
	Template SimpleSpell
	effects  []SpellEffect

	// Current instantiation of this spell. Can only be casting 1 instance of this spell at a time.
	Instance SimpleSpell
}

func (spell *Spell) reset(_ *Simulation) {
	spell.SpellMetrics = SpellMetrics{}
}

func (spell *Spell) doneIteration() {
	spell.Character.Metrics.addSpell(spell)
}

func (spell *Spell) Cast(sim *Simulation, target *Target) bool {
	// Initialize cast from precomputed template.
	instance := &spell.Instance
	if instance.objectInUse {
		panic(fmt.Sprintf("Spell (%s) already in use", instance.ActionID))
	}
	*instance = spell.Template
	instance.Effects = spell.effects
	copy(instance.Effects, spell.Template.Effects)

	if spell.ModifyCast != nil {
		spell.ModifyCast(sim, target, instance)
	}

	instance.Init(sim)
	return instance.Cast(sim, target, spell)
}

func ModifyCastAssignTarget(_ *Simulation, target *Target, instance *SimpleSpell) {
	instance.Effect.Target = target
}

// User-provided function for performing a cast of a spell. Should return whether
// the spell was cast (e.g. not blocked by CDs or resources).
//type SpellCast func(*Simulation, *Spell, *Target) bool

// Registers a new spell to the character. Returns the newly created spell.
func (character *Character) RegisterSpell(config SpellConfig) *Spell {
	if len(character.Spellbook) > 100 {
		panic(fmt.Sprintf("Over 100 registered spells when registering %s! There is probably a spell being registered every iteration.", config.Template.ActionID))
	}
	if len(config.Template.Effects) > 0 && config.Template.Effect.DamageMultiplier != 0 {
		panic("Cannot use both Effect and Effects, pick one!")
	}
	config.Template.Character = character

	spell := &Spell{
		ActionID:    config.Template.ActionID,
		Character:   character,
		SpellSchool: config.Template.SpellSchool,
		SpellExtras: config.Template.SpellExtras,

		ModifyCast:   config.ModifyCast,
		ApplyEffects: config.ApplyEffects,

		Template: config.Template,
		effects:  make([]SpellEffect, len(config.Template.Effects)),
	}

	character.Spellbook = append(character.Spellbook, spell)

	return spell
}

// Returns the first registered spell with the given ID, or nil if there are none.
func (character *Character) GetSpell(actionID ActionID) *Spell {
	for _, spell := range character.Spellbook {
		if spell.ActionID.SameAction(actionID) {
			return spell
		}
	}
	return nil
}

// Retrieves an existing spell with the same ID as the config uses, or registers it if there is none.
func (character *Character) GetOrRegisterSpell(config SpellConfig) *Spell {
	registered := character.GetSpell(config.Template.ActionID)
	if registered == nil {
		return character.RegisterSpell(config)
	} else {
		return registered
	}
}

func ApplyEffectFuncAll(effectFuncs []ApplySpellEffects) ApplySpellEffects {
	if len(effectFuncs) == 0 {
		return nil
	} else if len(effectFuncs) == 1 {
		return effectFuncs[0]
	} else {
		return func(sim *Simulation, target *Target, spell *Spell) {
			for _, effectFunc := range effectFuncs {
				effectFunc(sim, target, spell)
			}
		}
	}
}

func ApplyEffectFuncDirectDamage(baseEffect SpellEffect, outcomeDeterminer OutcomeDeterminer) ApplySpellEffects {
	return func(sim *Simulation, target *Target, spell *Spell) {
		effect := baseEffect
		effect.Target = target
		effect.Outcome = outcomeDeterminer(sim, spell, &effect)

		if effect.Landed() {
			effect.directCalculations(sim, spell)
		}

		effect.afterCalculations(sim, spell, false)
		spell.Instance.objectInUse = false
	}
}
