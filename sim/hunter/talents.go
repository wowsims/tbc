package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (hunter *Hunter) applyTalents() {
	if hunter.pet != nil {
		hunter.applyFocusedFire()
		hunter.applyFrenzy()
		hunter.registerBestialWrathCD()

		hunter.pet.damageMultiplier *= 1 + 0.04*float64(hunter.Talents.UnleashedFury)
		hunter.pet.AddStat(stats.MeleeCrit, core.MeleeCritRatingPerCritChance*2*float64(hunter.Talents.Ferocity))
		hunter.pet.AddStat(stats.MeleeHit, core.MeleeHitRatingPerHitChance*2*float64(hunter.Talents.AnimalHandler))
		hunter.pet.PseudoStats.MeleeSpeedMultiplier *= 1 + 0.04*float64(hunter.Talents.SerpentsSwiftness)
	}

	hunter.applyRangedEffects()
	hunter.applyGoForTheThroat()

	hunter.PseudoStats.RangedSpeedMultiplier *= 1 + 0.04*float64(hunter.Talents.SerpentsSwiftness)

	if hunter.Talents.CombatExperience > 0 {
		agiBonus := 1 + 0.01*float64(hunter.Talents.CombatExperience)
		hunter.Character.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Agility,
			ModifiedStat: stats.Agility,
			Modifier: func(agility float64, _ float64) float64 {
				return agility * agiBonus
			},
		})
		intBonus := 1 + 0.03*float64(hunter.Talents.CombatExperience)
		hunter.Character.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.Intellect,
			Modifier: func(intellect float64, _ float64) float64 {
				return intellect * intBonus
			},
		})
	}
	if hunter.Talents.CarefulAim > 0 {
		bonus := 0.15 * float64(hunter.Talents.CarefulAim)
		hunter.Character.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.RangedAttackPower,
			Modifier: func(intellect float64, rap float64) float64 {
				return rap + intellect*bonus
			},
		})
	}
	if hunter.Talents.MasterMarksman > 0 {
		bonus := 1 + 0.02*float64(hunter.Talents.MasterMarksman)
		hunter.Character.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.RangedAttackPower,
			ModifiedStat: stats.RangedAttackPower,
			Modifier: func(rap float64, _ float64) float64 {
				return rap * bonus
			},
		})
	}
}

var FocusedFireAuraID = core.NewAuraID()

func (hunter *Hunter) applyFocusedFire() {
	if hunter.Talents.FocusedFire == 0 {
		return
	}

	multiplier := 1.0 + 0.01*float64(hunter.Talents.FocusedFire)

	hunter.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: FocusedFireAuraID,
			OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				hitEffect.DamageMultiplier *= multiplier
			},
			OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				spellEffect.DamageMultiplier *= multiplier
			},
			OnBeforePeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage *float64) {
				*tickDamage *= multiplier
			},
		}
	})
}

var FrenzyAuraID = core.NewAuraID()
var FrenzyProcAuraID = core.NewAuraID()

func (hunter *Hunter) applyFrenzy() {
	if hunter.Talents.Frenzy == 0 {
		return
	}

	procChance := 0.2 * float64(hunter.Talents.Frenzy)

	hunter.pet.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		procAura := core.Aura{
			ID:       FrenzyProcAuraID,
			ActionID: core.ActionID{SpellID: 19625},
			Expires:  sim.CurrentTime + time.Second*8,
			OnExpire: func(sim *core.Simulation) {
				hunter.pet.PseudoStats.MeleeSpeedMultiplier /= 1.3
			},
		}
		tryProcAura := func() {
			if procChance == 1 || sim.RandomFloat("Frenzy") < procChance {
				hunter.pet.PseudoStats.MeleeSpeedMultiplier *= 1.3
				hunter.pet.AddAura(sim, procAura)
			}
		}

		return core.Aura{
			ID: FrenzyAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if hitEffect.HitType == core.MeleeHitTypeCrit {
					tryProcAura()
				}
			},
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellEffect.Crit {
					tryProcAura()
				}
			},
		}
	})
}

var BestialWrathAuraID = core.NewAuraID()
var BestialWrathPetAuraID = core.NewAuraID()
var BestialWrathCooldownID = core.NewCooldownID()

func (hunter *Hunter) registerBestialWrathCD() {
	if !hunter.Talents.BestialWrath {
		return
	}

	actionID := core.ActionID{SpellID: 19574, CooldownID: BestialWrathCooldownID}

	bestialWrathPetAura := core.Aura{
		ID:       BestialWrathPetAuraID,
		ActionID: actionID,
		OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
			hitEffect.DamageMultiplier *= 1.5
		},
		OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
			spellEffect.DamageMultiplier *= 1.5
		},
		OnBeforePeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage *float64) {
			*tickDamage *= 1.5
		},
	}

	bestialWrathAura := core.Aura{
		ID:       BestialWrathAuraID,
		ActionID: actionID,
		OnCast: func(sim *core.Simulation, cast *core.Cast) {
			cast.ManaCost -= cast.BaseManaCost * 0.2
		},
		OnBeforeMelee: func(sim *core.Simulation, ability *core.ActiveMeleeAbility) {
			ability.Cost.Value *= 0.8
		},
		OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
			hitEffect.DamageMultiplier *= 1.1
		},
		OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
			spellEffect.DamageMultiplier *= 1.1
		},
		OnBeforePeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage *float64) {
			*tickDamage *= 1.1
		},
	}

	manaCost := hunter.BaseMana() * 0.1
	cooldown := time.Minute * 2

	template := core.SimpleCast{
		Cast: core.Cast{
			ActionID:     actionID,
			Character:    hunter.GetCharacter(),
			Cooldown:     cooldown,
			BaseManaCost: manaCost,
			ManaCost:     manaCost,
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				petAura := bestialWrathPetAura
				petAura.Expires = sim.CurrentTime + time.Second*18
				hunter.pet.AddAura(sim, petAura)

				if hunter.Talents.TheBeastWithin {
					aura := bestialWrathAura
					aura.Expires = petAura.Expires
					hunter.AddAura(sim, aura)
				}
			},
		},
	}

	hunter.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: BestialWrathCooldownID,
		Cooldown:   cooldown,
		Type:       core.CooldownTypeDPS,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				cast := template
				cast.Init(sim)
				cast.StartCast(sim)
			}
		},
	})
}

var RangedEffectsAuraID = core.NewAuraID()

func (hunter *Hunter) applyRangedEffects() {
	critBonus := 1 * float64(hunter.Talents.LethalShots) * core.MeleeCritRatingPerCritChance
	critDamageBonus := 0.06 * float64(hunter.Talents.MortalShots)
	damageBonus := 1 + 0.01*float64(hunter.Talents.RangedWeaponSpecialization)

	hunter.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: RangedEffectsAuraID,
			OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if hitEffect.IsRanged() {
					hitEffect.BonusCritRating += critBonus
					ability.CritMultiplier = (ability.CritMultiplier-1)*critDamageBonus + 1
					hitEffect.DamageMultiplier *= damageBonus
				}
			},
		}
	})
}

var GoForTheThroatAuraID = core.NewAuraID()

func (hunter *Hunter) applyGoForTheThroat() {
	if hunter.Talents.GoForTheThroat == 0 {
		return
	}

	amount := 25.0 * float64(hunter.Talents.GoForTheThroat)

	hunter.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: GoForTheThroatAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.IsRanged() || hitEffect.HitType != core.MeleeHitTypeCrit {
					return
				}
				if hunter.pet == nil {
					return
				}
				hunter.pet.AddFocus(sim, amount, core.ActionID{SpellID: 34954})
			},
		}
	})
}
