package core

import (
	"fmt"
	"time"
)

// Function for calculating the base damage of a spell.
type BaseDamageCalculator func(*Simulation, *SpellHitEffect, *SpellCast) float64

// Creates a BaseDamageCalculator function which returns a flat value.
func BaseDamageFuncFlat(damage float64) BaseDamageCalculator {
	return func(_ *Simulation, _ *SpellHitEffect, _ *SpellCast) float64 {
		return damage
	}
}

// Creates a BaseDamageCalculator function with a single damage roll.
func BaseDamageFuncRoll(minFlatDamage float64, maxFlatDamage float64) BaseDamageCalculator {
	if minFlatDamage == maxFlatDamage {
		return BaseDamageFuncFlat(minFlatDamage)
	} else {
		deltaDamage := maxFlatDamage - minFlatDamage
		return func(sim *Simulation, _ *SpellHitEffect, _ *SpellCast) float64 {
			return damageRollOptimized(sim, minFlatDamage, deltaDamage)
		}
	}
}

func BaseDamageFuncMagic(minFlatDamage float64, maxFlatDamage float64, spellCoefficient float64) BaseDamageCalculator {
	if spellCoefficient == 0 {
		return BaseDamageFuncRoll(minFlatDamage, maxFlatDamage)
	}

	if minFlatDamage == 0 && maxFlatDamage == 0 {
		return func(_ *Simulation, hitEffect *SpellHitEffect, spellCast *SpellCast) float64 {
			return hitEffect.SpellPower(spellCast.Character, spellCast) * spellCoefficient
		}
	} else if minFlatDamage == maxFlatDamage {
		return func(sim *Simulation, hitEffect *SpellHitEffect, spellCast *SpellCast) float64 {
			damage := hitEffect.SpellPower(spellCast.Character, spellCast) * spellCoefficient
			return damage + minFlatDamage
		}
	} else {
		deltaDamage := maxFlatDamage - minFlatDamage
		return func(sim *Simulation, hitEffect *SpellHitEffect, spellCast *SpellCast) float64 {
			damage := hitEffect.SpellPower(spellCast.Character, spellCast) * spellCoefficient
			damage += damageRollOptimized(sim, minFlatDamage, deltaDamage)
			return damage
		}
	}
}

type Hand bool

const MainHand Hand = true
const OffHand Hand = false

func BaseDamageFuncMeleeWeapon(hand Hand, normalized bool, flatBonus float64, weaponMultiplier float64, includeBonusWeaponDamage bool) BaseDamageCalculator {
	// Bonus weapon damage applies after OH penalty: https://www.youtube.com/watch?v=bwCIU87hqTs
	// TODO not all weapon damage based attacks "scale" with +bonusWeaponDamage (e.g. Devastate, Shiv, Mutilate don't)
	// ... but for other's, BonusAttackPowerOnTarget only applies to weapon damage based attacks
	if normalized {
		if hand == MainHand {
			return func(sim *Simulation, hitEffect *SpellHitEffect, spellCast *SpellCast) float64 {
				damage := spellCast.Character.AutoAttacks.MH.calculateNormalizedWeaponDamage(
					sim, hitEffect.MeleeAttackPower(spellCast)+hitEffect.MeleeAttackPowerOnTarget())
				damage += flatBonus
				if includeBonusWeaponDamage {
					damage += hitEffect.PlusWeaponDamage(spellCast)
				}
				return damage * weaponMultiplier
			}
		} else {
			return func(sim *Simulation, hitEffect *SpellHitEffect, spellCast *SpellCast) float64 {
				damage := spellCast.Character.AutoAttacks.OH.calculateNormalizedWeaponDamage(
					sim, hitEffect.MeleeAttackPower(spellCast)+2*hitEffect.MeleeAttackPowerOnTarget())
				damage = damage*0.5 + flatBonus
				if includeBonusWeaponDamage {
					damage += hitEffect.PlusWeaponDamage(spellCast)
				}
				return damage * weaponMultiplier
			}
		}
	} else {
		if hand == MainHand {
			return func(sim *Simulation, hitEffect *SpellHitEffect, spellCast *SpellCast) float64 {
				damage := spellCast.Character.AutoAttacks.MH.calculateWeaponDamage(
					sim, hitEffect.MeleeAttackPower(spellCast)+hitEffect.MeleeAttackPowerOnTarget())
				damage += flatBonus
				if includeBonusWeaponDamage {
					damage += hitEffect.PlusWeaponDamage(spellCast)
				}
				return damage * weaponMultiplier
			}
		} else {
			return func(sim *Simulation, hitEffect *SpellHitEffect, spellCast *SpellCast) float64 {
				damage := spellCast.Character.AutoAttacks.OH.calculateWeaponDamage(
					sim, hitEffect.MeleeAttackPower(spellCast)+2*hitEffect.MeleeAttackPowerOnTarget())
				damage = damage*0.5 + flatBonus
				if includeBonusWeaponDamage {
					damage += hitEffect.PlusWeaponDamage(spellCast)
				}
				return damage * weaponMultiplier
			}
		}
	}
}

func BaseDamageFuncRangedWeapon(flatBonus float64) BaseDamageCalculator {
	return func(sim *Simulation, hitEffect *SpellHitEffect, spellCast *SpellCast) float64 {
		return spellCast.Character.AutoAttacks.Ranged.calculateWeaponDamage(sim, hitEffect.RangedAttackPower(spellCast)+hitEffect.RangedAttackPowerOnTarget()) +
			flatBonus +
			hitEffect.PlusWeaponDamage(spellCast)
	}
}

// Performs an actual damage roll. Keep this internal because the 2nd parameter
// is the delta rather than maxDamage, which is error-prone.
func damageRollOptimized(sim *Simulation, minDamage float64, deltaDamage float64) float64 {
	return minDamage + deltaDamage*sim.RandomFloat("Damage Roll")
}

// For convenience, but try to use damageRollOptimized in most cases.
func DamageRoll(sim *Simulation, minDamage float64, maxDamage float64) float64 {
	return damageRollOptimized(sim, minDamage, maxDamage-minDamage)
}

func DamageRollFunc(minDamage float64, maxDamage float64) func(*Simulation) float64 {
	deltaDamage := maxDamage - minDamage
	return func(sim *Simulation) float64 {
		return damageRollOptimized(sim, minDamage, deltaDamage)
	}
}

type SpellHitEffect struct {
	SpellEffect

	BaseDamage BaseDamageCalculator

	DotInput DotDamageInput
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
	if len(spellTemplate.Effects) > 0 && spellTemplate.Effect.DamageMultiplier != 0 {
		panic("Cannot use both Effect and Effects, pick one!")
	}

	return SimpleSpellTemplate{
		template: spellTemplate,
		effects:  make([]SpellHitEffect, len(spellTemplate.Effects)),
	}
}
