package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (hunter *Hunter) applyTalents() {
	if hunter.pet != nil {
		hunter.applyFocusedFire()
		hunter.applyFrenzy()
		hunter.applyFerociousInspiration()
		hunter.registerBestialWrathCD()

		hunter.pet.damageMultiplier *= 1 + 0.04*float64(hunter.Talents.UnleashedFury)
		hunter.pet.AddStat(stats.MeleeCrit, core.MeleeCritRatingPerCritChance*2*float64(hunter.Talents.Ferocity))
		hunter.pet.AddStat(stats.SpellCrit, core.SpellCritRatingPerCritChance*2*float64(hunter.Talents.Ferocity))
		hunter.pet.AddStat(stats.MeleeHit, core.MeleeHitRatingPerHitChance*2*float64(hunter.Talents.AnimalHandler))
		hunter.pet.AddStat(stats.SpellHit, core.SpellHitRatingPerHitChance*2*float64(hunter.Talents.AnimalHandler))
		hunter.pet.PseudoStats.MeleeSpeedMultiplier *= 1 + 0.04*float64(hunter.Talents.SerpentsSwiftness)
	}

	hunter.applyRangedEffects()
	hunter.applyGoForTheThroat()
	hunter.applySlaying()
	hunter.applyThrillOfTheHunt()
	hunter.applyExposeWeakness()
	hunter.applyMasterTactician()
	hunter.registerReadinessCD()

	hunter.PseudoStats.RangedSpeedMultiplier *= 1 + 0.04*float64(hunter.Talents.SerpentsSwiftness)
	hunter.AddStat(stats.MeleeHit, core.MeleeHitRatingPerHitChance*1*float64(hunter.Talents.Surefooted))
	hunter.AddStat(stats.MeleeCrit, core.MeleeCritRatingPerCritChance*1*float64(hunter.Talents.KillerInstinct))

	if hunter.Talents.CombatExperience > 0 {
		agiBonus := 1 + 0.01*float64(hunter.Talents.CombatExperience)
		hunter.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Agility,
			ModifiedStat: stats.Agility,
			Modifier: func(agility float64, _ float64) float64 {
				return agility * agiBonus
			},
		})
		intBonus := 1 + 0.03*float64(hunter.Talents.CombatExperience)
		hunter.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.Intellect,
			Modifier: func(intellect float64, _ float64) float64 {
				return intellect * intBonus
			},
		})
	}
	if hunter.Talents.CarefulAim > 0 {
		bonus := 0.15 * float64(hunter.Talents.CarefulAim)
		hunter.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.RangedAttackPower,
			Modifier: func(intellect float64, rap float64) float64 {
				return rap + intellect*bonus
			},
		})
	}
	if hunter.Talents.MasterMarksman > 0 {
		bonus := 1 + 0.02*float64(hunter.Talents.MasterMarksman)
		hunter.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.RangedAttackPower,
			ModifiedStat: stats.RangedAttackPower,
			Modifier: func(rap float64, _ float64) float64 {
				return rap * bonus
			},
		})
	}
	if hunter.Talents.SurvivalInstincts > 0 {
		apBonus := 1 + 0.02*float64(hunter.Talents.SurvivalInstincts)
		hunter.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.AttackPower,
			ModifiedStat: stats.AttackPower,
			Modifier: func(ap float64, _ float64) float64 {
				return ap * apBonus
			},
		})
		hunter.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.RangedAttackPower,
			ModifiedStat: stats.RangedAttackPower,
			Modifier: func(rap float64, _ float64) float64 {
				return rap * apBonus
			},
		})
	}
	if hunter.Talents.LightningReflexes > 0 {
		agiBonus := 1 + 0.03*float64(hunter.Talents.LightningReflexes)
		hunter.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Agility,
			ModifiedStat: stats.Agility,
			Modifier: func(agility float64, _ float64) float64 {
				return agility * agiBonus
			},
		})
	}
}

func (hunter *Hunter) critMultiplier(isRanged bool, target *core.Target) float64 {
	primaryModifier := 1.0
	secondaryModifier := 0.0

	monsterMultiplier := 1.0 + 0.01*float64(hunter.Talents.MonsterSlaying)
	humanoidMultiplier := 1.0 + 0.01*float64(hunter.Talents.HumanoidSlaying)
	if target.MobType == proto.MobType_MobTypeBeast || target.MobType == proto.MobType_MobTypeGiant || target.MobType == proto.MobType_MobTypeDragonkin {
		primaryModifier *= monsterMultiplier
	} else if target.MobType == proto.MobType_MobTypeHumanoid {
		primaryModifier *= humanoidMultiplier
	}

	if isRanged {
		secondaryModifier += 0.06 * float64(hunter.Talents.MortalShots)
	}

	return hunter.MeleeCritMultiplier(primaryModifier, secondaryModifier)
}

var FocusedFireAuraID = core.NewAuraID()

func (hunter *Hunter) applyFocusedFire() {
	if hunter.Talents.FocusedFire == 0 || hunter.pet == nil {
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
			OnExpire: func(sim *core.Simulation) {
				hunter.pet.PseudoStats.MeleeSpeedMultiplier /= 1.3
			},
		}

		tryProcAura := func() {
			if procChance == 1 || sim.RandomFloat("Frenzy") < procChance {
				hunter.pet.PseudoStats.MeleeSpeedMultiplier *= 1.3
				aura := procAura
				aura.Expires = sim.CurrentTime + time.Second*8
				hunter.pet.AddAura(sim, aura)
			}
		}

		return core.Aura{
			ID: FrenzyAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if hitEffect.Outcome.Matches(core.OutcomeCrit) {
					tryProcAura()
				}
			},
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellEffect.Outcome.Matches(core.OutcomeCrit) {
					tryProcAura()
				}
			},
		}
	})
}

// One ID for each index in the party (0-4) so auras from multiple hunters
// don't collide.
var FerociousInspirationAuraIDs = []core.AuraID{
	core.NewAuraID(),
	core.NewAuraID(),
	core.NewAuraID(),
	core.NewAuraID(),
	core.NewAuraID(),
}
var FerociousInspirationAuraID = core.NewAuraID()

func (hunter *Hunter) applyFerociousInspiration() {
	if hunter.pet == nil || hunter.Talents.FerociousInspiration == 0 {
		return
	}

	multiplier := 1.0 + 0.01*float64(hunter.Talents.FerociousInspiration)

	procAura := core.Aura{
		ID:       FerociousInspirationAuraIDs[hunter.PartyIndex],
		ActionID: core.ActionID{SpellID: 34460, Tag: int32(hunter.RaidIndex)},
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

	hunter.pet.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		applyAura := func() {
			for _, playerOrPet := range hunter.Party.PlayersAndPets {
				char := playerOrPet.GetCharacter()
				aura := procAura
				aura.Expires = sim.CurrentTime + time.Second*10
				char.AddAura(sim, aura)
			}
		}

		return core.Aura{
			ID: FerociousInspirationAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if hitEffect.Outcome.Matches(core.OutcomeCrit) {
					applyAura()
				}
			},
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellEffect.Outcome.Matches(core.OutcomeCrit) {
					applyAura()
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
			cast.Cost.Value -= cast.BaseCost.Value * 0.2
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
			ActionID:  actionID,
			Character: hunter.GetCharacter(),
			Cooldown:  cooldown,
			BaseCost: core.ResourceCost{
				Type:  stats.Mana,
				Value: manaCost,
			},
			Cost: core.ResourceCost{
				Type:  stats.Mana,
				Value: manaCost,
			},
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
			if hunter.CurrentMana() < manaCost {
				return false
			}
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
	damageBonus := 1 + 0.01*float64(hunter.Talents.RangedWeaponSpecialization)

	hunter.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: RangedEffectsAuraID,
			OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if ability.OutcomeRollCategory.Matches(core.OutcomeRollCategoryRanged) {
					hitEffect.BonusCritRating += critBonus
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
	if hunter.pet == nil {
		return
	}

	amount := 25.0 * float64(hunter.Talents.GoForTheThroat)

	hunter.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: GoForTheThroatAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !ability.OutcomeRollCategory.Matches(core.OutcomeRollCategoryRanged) || !hitEffect.Outcome.Matches(core.OutcomeCrit) {
					return
				}
				if !hunter.pet.IsEnabled() {
					return
				}
				hunter.pet.AddFocus(sim, amount, core.ActionID{SpellID: 34954})
			},
		}
	})
}

var SlayingAuraID = core.NewAuraID()

func (hunter *Hunter) applySlaying() {
	if hunter.Talents.MonsterSlaying == 0 && hunter.Talents.HumanoidSlaying == 0 {
		return
	}

	monsterMultiplier := 1.0 + 0.01*float64(hunter.Talents.MonsterSlaying)
	humanoidMultiplier := 1.0 + 0.01*float64(hunter.Talents.HumanoidSlaying)

	hunter.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: SlayingAuraID,
			OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if hitEffect.Target.MobType == proto.MobType_MobTypeBeast || hitEffect.Target.MobType == proto.MobType_MobTypeGiant || hitEffect.Target.MobType == proto.MobType_MobTypeDragonkin {
					hitEffect.DamageMultiplier *= monsterMultiplier
				} else if hitEffect.Target.MobType == proto.MobType_MobTypeHumanoid {
					hitEffect.DamageMultiplier *= humanoidMultiplier
				}
			},
			OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellEffect.Target.MobType == proto.MobType_MobTypeBeast || spellEffect.Target.MobType == proto.MobType_MobTypeGiant || spellEffect.Target.MobType == proto.MobType_MobTypeDragonkin {
					spellEffect.DamageMultiplier *= monsterMultiplier
				} else if spellEffect.Target.MobType == proto.MobType_MobTypeHumanoid {
					spellEffect.DamageMultiplier *= humanoidMultiplier
				}
			},
			OnBeforePeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage *float64) {
				if spellEffect.Target.MobType == proto.MobType_MobTypeBeast || spellEffect.Target.MobType == proto.MobType_MobTypeGiant || spellEffect.Target.MobType == proto.MobType_MobTypeDragonkin {
					*tickDamage *= monsterMultiplier
				} else if spellEffect.Target.MobType == proto.MobType_MobTypeHumanoid {
					*tickDamage *= humanoidMultiplier
				}
			},
		}
	})
}

var ThrillOfTheHuntAuraID = core.NewAuraID()

func (hunter *Hunter) applyThrillOfTheHunt() {
	if hunter.Talents.ThrillOfTheHunt == 0 {
		return
	}

	procChance := float64(hunter.Talents.ThrillOfTheHunt) / 3
	actionID := core.ActionID{SpellID: 34499}

	hunter.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: ThrillOfTheHuntAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				// mask 256
				if !hitEffect.ProcMask.Matches(core.ProcMaskRangedSpecial) {
					return
				}

				if !hitEffect.Outcome.Matches(core.OutcomeCrit) {
					return
				}

				if procChance == 1 || sim.RandomFloat("ThrillOfTheHunt") < procChance {
					hunter.AddMana(sim, ability.Cost.Value*0.4, actionID, false)
				}
			},
		}
	})
}

var ExposeWeaknessAuraID = core.NewAuraID()

func (hunter *Hunter) applyExposeWeakness() {
	if hunter.Talents.ExposeWeakness == 0 {
		return
	}

	procChance := float64(hunter.Talents.ExposeWeakness) / 3

	hunter.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: ExposeWeaknessAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !ability.OutcomeRollCategory.Matches(core.OutcomeRollCategoryRanged) {
					return
				}

				if !hitEffect.Outcome.Matches(core.OutcomeCrit) {
					return
				}

				if hitEffect.Target.RemainingAuraDuration(sim, core.ExposeWeaknessDebuffID) == core.NeverExpires {
					// Don't overwrite permanent version
					return
				}

				if procChance == 1 || sim.RandomFloat("ExposeWeakness") < procChance {
					hitEffect.Target.AddAura(sim, core.ExposeWeaknessAura(sim.CurrentTime, hunter.GetStat(stats.Agility), 1.0))
				}
			},
		}
	})
}

var MasterTacticianAuraID = core.NewAuraID()
var MasterTacticianProcAuraID = core.NewAuraID()

func (hunter *Hunter) applyMasterTactician() {
	if hunter.Talents.MasterTactician == 0 {
		return
	}

	procChance := 0.06
	critBonus := 2 * core.MeleeCritRatingPerCritChance * float64(hunter.Talents.MasterTactician)

	hunter.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		procAura := core.Aura{
			ID:       MasterTacticianProcAuraID,
			ActionID: core.ActionID{SpellID: 34839},
			OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				hitEffect.BonusCritRating += critBonus
			},
		}

		return core.Aura{
			ID: MasterTacticianAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !ability.OutcomeRollCategory.Matches(core.OutcomeRollCategoryRanged) || !hitEffect.Landed() {
					return
				}

				if sim.RandomFloat("Master Tactician") > procChance {
					return
				}

				aura := procAura
				aura.Expires = sim.CurrentTime + time.Second*8
				hunter.AddAura(sim, aura)
			},
		}
	})
}

var ReadinessCooldownID = core.NewCooldownID()

func (hunter *Hunter) registerReadinessCD() {
	if !hunter.Talents.Readiness {
		return
	}

	actionID := core.ActionID{SpellID: 23989, CooldownID: ReadinessCooldownID}
	cooldown := time.Minute * 5

	template := core.SimpleCast{
		Cast: core.Cast{
			ActionID:  actionID,
			Character: hunter.GetCharacter(),
			Cooldown:  cooldown,
			//GCD:         time.Second * 1, TODO: GCD causes panic
			//IgnoreHaste: true, // Hunter GCD is locked
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				hunter.SetCD(RapidFireCooldownID, 0)
				hunter.SetCD(MultiShotCooldownID, 0)
				hunter.SetCD(ArcaneShotCooldownID, 0)
				hunter.SetCD(AimedShotCooldownID, 0)
				hunter.SetCD(KillCommandCooldownID, 0)
				hunter.SetCD(RaptorStrikeCooldownID, 0)
			},
		},
	}

	hunter.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: ReadinessCooldownID,
		Cooldown:   cooldown,
		//UsesGCD:    true,
		Type: core.CooldownTypeDPS,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			// Don't use if there are no cooldowns to reset.
			if !character.IsOnCD(RapidFireCooldownID, sim.CurrentTime) {
				return false
			}
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
