package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (rogue *Rogue) applyTalents() {
	// TODO: Puncturing Wounds, IEA, poisons, mutilate, blade flurry, adrenaline rush
	// Everything in the sub tree

	rogue.applyMurder()
	rogue.applySealFate()
	rogue.applyWeaponSpecializations()
	rogue.applyCombatPotency()

	rogue.AddStat(stats.MeleeCrit, core.MeleeCritRatingPerCritChance*1*float64(rogue.Talents.Malice))
	rogue.AddStat(stats.MeleeHit, core.MeleeHitRatingPerHitChance*1*float64(rogue.Talents.Precision))
	rogue.AddStat(stats.Expertise, core.ExpertisePerQuarterPercentReduction*5*float64(rogue.Talents.WeaponExpertise))
	rogue.AutoAttacks.OHAuto.Effect.WeaponInput.DamageMultiplier *= 1.0 + 0.1*float64(rogue.Talents.DualWieldSpecialization)

	if rogue.Talents.Vitality > 0 {
		agiBonus := 1 + 0.01*float64(rogue.Talents.Vitality)
		rogue.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Agility,
			ModifiedStat: stats.Agility,
			Modifier: func(agility float64, _ float64) float64 {
				return agility * agiBonus
			},
		})
		stamBonus := 1 + 0.02*float64(rogue.Talents.Vitality)
		rogue.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Stamina,
			ModifiedStat: stats.Stamina,
			Modifier: func(stamina float64, _ float64) float64 {
				return stamina * stamBonus
			},
		})
	}

	rogue.registerColdBloodCD()
}

func (rogue *Rogue) critMultiplier(isMH bool, applyLethality bool) float64 {
	primaryModifier := 1.0
	secondaryModifier := 0.0

	isMace := false
	if weapon := rogue.Equip[proto.ItemSlot_ItemSlotMainHand]; isMH && weapon.ID != 0 {
		if weapon.WeaponType == proto.WeaponType_WeaponTypeMace {
			isMace = true
		}
	} else if weapon := rogue.Equip[proto.ItemSlot_ItemSlotOffHand]; !isMH && weapon.ID != 0 {
		if weapon.WeaponType == proto.WeaponType_WeaponTypeMace {
			isMace = true
		}
	}
	if isMace {
		primaryModifier *= 1 + 0.01*float64(rogue.Talents.MaceSpecialization)
	}

	if applyLethality {
		secondaryModifier += 0.06 * float64(rogue.Talents.Lethality)
	}

	return rogue.MeleeCritMultiplier(primaryModifier, secondaryModifier)
}

var FindWeaknessAuraID = core.NewAuraID()

func (rogue *Rogue) makeFinishingMoveEffectApplier() func(sim *core.Simulation, numPoints int32) {
	ruthlessnessChance := 0.2 * float64(rogue.Talents.Ruthlessness)
	relentlessStrikes := rogue.Talents.RelentlessStrikes

	findWeaknessMultiplier := 1.0 + 0.02*float64(rogue.Talents.FindWeakness)

	findWeaknessAura := core.Aura{
		ID:       FindWeaknessAuraID,
		ActionID: core.ActionID{SpellID: 31242},
		OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
			// TODO: This should be rogue abilities only, not all specials.
			if hitEffect.ProcMask.Matches(core.ProcMaskMeleeSpecial) {
				hitEffect.DamageMultiplier *= findWeaknessMultiplier
			}
		},
	}

	return func(sim *core.Simulation, numPoints int32) {
		if ruthlessnessChance > 0 && sim.RandomFloat("Ruthlessness") < ruthlessnessChance {
			rogue.AddComboPoint(sim)
		}
		if relentlessStrikes {
			if numPoints == 5 || sim.RandomFloat("RelentlessStrikes") < 0.2*float64(numPoints) {
				rogue.AddEnergy(sim, 25, core.ActionID{SpellID: 14179})
			}
		}
		if findWeaknessMultiplier != 1 {
			aura := findWeaknessAura
			aura.Expires = sim.CurrentTime + time.Second*10
			rogue.AddAura(sim, aura)
		}
	}
}

var MurderAuraID = core.NewAuraID()

func (rogue *Rogue) applyMurder() {
	if rogue.Talents.Murder == 0 {
		return
	}

	damageMultiplier := 1.0 + 0.01*float64(rogue.Talents.Murder)

	rogue.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: MurderAuraID,
			OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if hitEffect.Target.MobType == proto.MobType_MobTypeHumanoid || hitEffect.Target.MobType == proto.MobType_MobTypeBeast || hitEffect.Target.MobType == proto.MobType_MobTypeGiant || hitEffect.Target.MobType == proto.MobType_MobTypeDragonkin {
					hitEffect.DamageMultiplier *= damageMultiplier
				}
			},
			OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellEffect.Target.MobType == proto.MobType_MobTypeHumanoid || spellEffect.Target.MobType == proto.MobType_MobTypeBeast || spellEffect.Target.MobType == proto.MobType_MobTypeGiant || spellEffect.Target.MobType == proto.MobType_MobTypeDragonkin {
					spellEffect.DamageMultiplier *= damageMultiplier
				}
			},
			OnBeforePeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage *float64) {
				if spellEffect.Target.MobType == proto.MobType_MobTypeHumanoid || spellEffect.Target.MobType == proto.MobType_MobTypeBeast || spellEffect.Target.MobType == proto.MobType_MobTypeGiant || spellEffect.Target.MobType == proto.MobType_MobTypeDragonkin {
					*tickDamage *= damageMultiplier
				}
			},
		}
	})
}

var ColdBloodAuraID = core.NewAuraID()
var ColdBloodCooldownID = core.NewCooldownID()

func (rogue *Rogue) registerColdBloodCD() {
	if !rogue.Talents.ColdBlood {
		return
	}

	actionID := core.ActionID{SpellID: 14177, CooldownID: ColdBloodCooldownID}

	coldBloodAura := core.Aura{
		ID:       ColdBloodAuraID,
		ActionID: actionID,
		Expires:  core.NeverExpires,
		OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
			// TODO: This should be rogue abilities only, not all specials.
			if hitEffect.ProcMask.Matches(core.ProcMaskMeleeSpecial) {
				hitEffect.BonusCritRating += 100 * core.MeleeCritRatingPerCritChance
				rogue.RemoveAura(sim, ColdBloodAuraID)
			}
		},
	}

	cooldown := time.Minute * 3

	template := core.SimpleCast{
		Cast: core.Cast{
			ActionID:  actionID,
			Character: rogue.GetCharacter(),
			Cooldown:  cooldown,
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				rogue.AddAura(sim, coldBloodAura)
			},
		},
	}

	rogue.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: ColdBloodCooldownID,
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

var SealFateAuraID = core.NewAuraID()

func (rogue *Rogue) applySealFate() {
	if rogue.Talents.SealFate == 0 {
		return
	}

	procChance := 0.2 * float64(rogue.Talents.SealFate)

	rogue.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: SealFateAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !ability.ActionID.SameAction(SinisterStrikeActionID) {
					return
				}

				if !hitEffect.Outcome.Matches(core.OutcomeCrit) {
					return
				}

				if procChance == 1 || sim.RandomFloat("Seal Fate") < procChance {
					rogue.AddComboPoint(sim)
				}
			},
		}
	})
}

var DaggerAndFistSpecializationsAuraID = core.NewAuraID()
var SwordSpecializationAuraID = core.NewAuraID()

func (rogue *Rogue) applyWeaponSpecializations() {
	mhCritBonus := 0.0
	ohCritBonus := 0.0
	if weapon := rogue.Equip[proto.ItemSlot_ItemSlotMainHand]; weapon.ID != 0 {
		if weapon.WeaponType == proto.WeaponType_WeaponTypeFist {
			mhCritBonus = 1 * core.MeleeCritRatingPerCritChance * float64(rogue.Talents.FistWeaponSpecialization)
		} else if weapon.WeaponType == proto.WeaponType_WeaponTypeDagger {
			mhCritBonus = 1 * core.MeleeCritRatingPerCritChance * float64(rogue.Talents.DaggerSpecialization)
		}
	} else if weapon := rogue.Equip[proto.ItemSlot_ItemSlotOffHand]; weapon.ID != 0 {
		if weapon.WeaponType == proto.WeaponType_WeaponTypeFist {
			ohCritBonus = 1 * core.MeleeCritRatingPerCritChance * float64(rogue.Talents.FistWeaponSpecialization)
		} else if weapon.WeaponType == proto.WeaponType_WeaponTypeDagger {
			ohCritBonus = 1 * core.MeleeCritRatingPerCritChance * float64(rogue.Talents.DaggerSpecialization)
		}
	}

	if mhCritBonus > 0 || ohCritBonus > 0 {
		rogue.AddPermanentAura(func(sim *core.Simulation) core.Aura {
			return core.Aura{
				ID: DaggerAndFistSpecializationsAuraID,
				OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
					if hitEffect.ProcMask.Matches(core.ProcMaskMeleeMH) {
						hitEffect.BonusCritRating += mhCritBonus
					} else if hitEffect.ProcMask.Matches(core.ProcMaskMeleeOH) {
						hitEffect.BonusCritRating += ohCritBonus
					}
				},
			}
		})
	}

	// https://tbc.wowhead.com/spell=13964/sword-specialization, proc mask = 20.
	swordSpecMask := core.ProcMaskEmpty
	if rogue.Equip[proto.ItemSlot_ItemSlotMainHand].WeaponType == proto.WeaponType_WeaponTypeSword {
		swordSpecMask |= core.ProcMaskMeleeMH
	}
	if rogue.Equip[proto.ItemSlot_ItemSlotOffHand].WeaponType == proto.WeaponType_WeaponTypeSword {
		swordSpecMask |= core.ProcMaskMeleeOH
	}
	if rogue.Talents.SwordSpecialization > 0 && swordSpecMask != core.ProcMaskEmpty {
		rogue.AddPermanentAura(func(sim *core.Simulation) core.Aura {
			procChance := 0.01 * float64(rogue.Talents.SwordSpecialization)
			var icd core.InternalCD
			icdDur := time.Millisecond * 500

			mhAttack := rogue.AutoAttacks.MHAuto
			mhAttack.ActionID = core.ActionID{SpellID: 13964}

			cachedAttack := core.ActiveMeleeAbility{}

			return core.Aura{
				ID: SwordSpecializationAuraID,
				OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
					if !hitEffect.Landed() {
						return
					}

					if !hitEffect.ProcMask.Matches(swordSpecMask) {
						return
					}

					if icd.IsOnCD(sim) {
						return
					}

					if sim.RandomFloat("Sword Specialization") > procChance {
						return
					}
					icd = core.InternalCD(sim.CurrentTime + icdDur)

					// Got a proc
					cachedAttack = mhAttack
					cachedAttack.Effect.Target = hitEffect.Target
					cachedAttack.Attack(sim)
				},
			}
		})
	}
}

var CombatPotencyAuraID = core.NewAuraID()

func (rogue *Rogue) applyCombatPotency() {
	if rogue.Talents.CombatPotency == 0 {
		return
	}

	const procChance = 0.2
	energyBonus := 3.0 * float64(rogue.Talents.CombatPotency)

	rogue.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: CombatPotencyAuraID,
			OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
				if !hitEffect.Landed() {
					return
				}

				// https://tbc.wowhead.com/spell=35553/combat-potency, proc mask = 8838608.
				if !hitEffect.ProcMask.Matches(core.ProcMaskMeleeOHAuto) {
					return
				}

				if sim.RandomFloat("Combat Potency") > procChance {
					return
				}

				rogue.AddEnergy(sim, energyBonus, core.ActionID{SpellID: 35553})
			},
		}
	})
}
