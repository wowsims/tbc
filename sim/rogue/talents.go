package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (rogue *Rogue) ApplyTalents() {
	// TODO: Puncturing Wounds, IEA, poisons, mutilate, blade flurry, adrenaline rush
	// Everything in the sub tree

	rogue.applyMurder()
	rogue.applySealFate()
	rogue.applyWeaponSpecializations()
	rogue.applyCombatPotency()

	rogue.AddStat(stats.Dodge, core.DodgeRatingPerDodgeChance*1*float64(rogue.Talents.LightningReflexes))
	rogue.AddStat(stats.Parry, core.ParryRatingPerParryChance*1*float64(rogue.Talents.Deflection))
	rogue.AddStat(stats.MeleeCrit, core.MeleeCritRatingPerCritChance*1*float64(rogue.Talents.Malice))
	rogue.AddStat(stats.MeleeHit, core.MeleeHitRatingPerHitChance*1*float64(rogue.Talents.Precision))
	rogue.AddStat(stats.Expertise, core.ExpertisePerQuarterPercentReduction*5*float64(rogue.Talents.WeaponExpertise))
	rogue.AutoAttacks.OHAuto.Effect.WeaponInput.DamageMultiplier *= 1.0 + 0.1*float64(rogue.Talents.DualWieldSpecialization)
	rogue.AddStat(stats.ArmorPenetration, 186*float64(rogue.Talents.SerratedBlades))

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

	if rogue.Talents.Deadliness > 0 {
		apBonus := 1 + 0.02*float64(rogue.Talents.Deadliness)
		rogue.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.AttackPower,
			ModifiedStat: stats.AttackPower,
			Modifier: func(ap float64, _ float64) float64 {
				return ap * apBonus
			},
		})
	}

	if rogue.Talents.SinisterCalling > 0 {
		agiBonus := 1 + 0.03*float64(rogue.Talents.SinisterCalling)
		rogue.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Agility,
			ModifiedStat: stats.Agility,
			Modifier: func(agi float64, _ float64) float64 {
				return agi * agiBonus
			},
		})
	}

	rogue.registerColdBloodCD()
	rogue.registerBladeFlurryCD()
	rogue.registerAdrenalineRushCD()
}

var FindWeaknessAuraID = core.NewAuraID()

func (rogue *Rogue) makeFinishingMoveEffectApplier(_ *core.Simulation) func(sim *core.Simulation, numPoints int32) {
	ruthlessnessChance := 0.2 * float64(rogue.Talents.Ruthlessness)
	relentlessStrikes := rogue.Talents.RelentlessStrikes

	findWeaknessMultiplier := 1.0 + 0.02*float64(rogue.Talents.FindWeakness)

	findWeaknessAura := core.Aura{
		ID:       FindWeaknessAuraID,
		ActionID: core.ActionID{SpellID: 31242},
		Duration: time.Second * 10,
		OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellHitEffect) {
			// TODO: This should be rogue abilities only, not all specials.
			if spellEffect.ProcMask.Matches(core.ProcMaskMeleeSpecial) {
				spellEffect.DamageMultiplier *= findWeaknessMultiplier
			}
		},
	}

	netherblade4pc := ItemSetNetherblade.CharacterHasSetBonus(&rogue.Character, 4)

	return func(sim *core.Simulation, numPoints int32) {
		if ruthlessnessChance > 0 && sim.RandomFloat("Ruthlessness") < ruthlessnessChance {
			rogue.AddComboPoints(sim, 1, core.ActionID{SpellID: 14161})
		}
		if netherblade4pc && sim.RandomFloat("Netherblade 4pc") < 0.15 {
			rogue.AddComboPoints(sim, 1, core.ActionID{SpellID: 37168})
		}
		if relentlessStrikes {
			if numPoints == 5 || sim.RandomFloat("RelentlessStrikes") < 0.2*float64(numPoints) {
				rogue.AddEnergy(sim, 25, core.ActionID{SpellID: 14179})
			}
		}
		if findWeaknessMultiplier != 1 {
			rogue.ReplaceAura(sim, findWeaknessAura)
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
			OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellHitEffect) {
				switch spellEffect.Target.MobType {
				case proto.MobType_MobTypeHumanoid, proto.MobType_MobTypeBeast, proto.MobType_MobTypeGiant, proto.MobType_MobTypeDragonkin:
					spellEffect.DamageMultiplier *= damageMultiplier
				}
			},
			OnBeforePeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage *float64) {
				switch spellEffect.Target.MobType {
				case proto.MobType_MobTypeHumanoid, proto.MobType_MobTypeBeast, proto.MobType_MobTypeGiant, proto.MobType_MobTypeDragonkin:
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
		Duration: core.NeverExpires,
		OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellHitEffect) {
			// TODO: This should be rogue abilities only, not all specials.
			if spellEffect.ProcMask.Matches(core.ProcMaskMeleeSpecial) {
				spellEffect.BonusCritRating += 100 * core.MeleeCritRatingPerCritChance
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
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellCast.SpellExtras.Matches(SpellFlagBuilder) {
					return
				}

				if !spellEffect.Outcome.Matches(core.OutcomeCrit) {
					return
				}

				if procChance == 1 || sim.RandomFloat("Seal Fate") < procChance {
					rogue.AddComboPoints(sim, 1, core.ActionID{SpellID: 14195})
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
				OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellHitEffect) {
					if spellEffect.ProcMask.Matches(core.ProcMaskMeleeMH) {
						spellEffect.BonusCritRating += mhCritBonus
					} else if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOH) {
						spellEffect.BonusCritRating += ohCritBonus
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
			cachedAttack := core.SimpleSpell{}

			return core.Aura{
				ID: SwordSpecializationAuraID,
				OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
					if !spellEffect.Landed() {
						return
					}

					if !spellEffect.ProcMask.Matches(swordSpecMask) {
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
					cachedAttack.Effect.Target = spellEffect.Target
					cachedAttack.Cast(sim)
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
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					return
				}

				// https://tbc.wowhead.com/spell=35553/combat-potency, proc mask = 8838608.
				if !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOH) {
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

var BladeFlurryAuraID = core.NewAuraID()
var BladeFlurryCooldownID = core.NewCooldownID()

func (rogue *Rogue) registerBladeFlurryCD() {
	if !rogue.Talents.BladeFlurry {
		return
	}

	actionID := core.ActionID{SpellID: 13877, CooldownID: BladeFlurryCooldownID}
	const hasteBonus = 1.2
	const inverseHasteBonus = 1 / 1.2
	const energyCost = 25.0

	dur := time.Second * 15
	cooldown := time.Minute * 2

	bladeFlurryAura := core.Aura{
		ID:       BladeFlurryAuraID,
		ActionID: actionID,
		Duration: dur,
		OnGain: func(sim *core.Simulation) {
			rogue.MultiplyMeleeSpeed(sim, hasteBonus)
		},
		OnExpire: func(sim *core.Simulation) {
			rogue.MultiplyMeleeSpeed(sim, inverseHasteBonus)
		},
		OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellHitEffect) {
			if sim.GetNumTargets() > 1 {
				spellEffect.DamageMultiplier *= 2
			}
		},
	}

	template := core.SimpleCast{
		Cast: core.Cast{
			ActionID:    actionID,
			Character:   rogue.GetCharacter(),
			Cooldown:    cooldown,
			GCD:         time.Second,
			IgnoreHaste: true,
			Cost: core.ResourceCost{
				Type:  stats.Energy,
				Value: energyCost,
			},
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				rogue.AddAura(sim, bladeFlurryAura)
			},
		},
	}

	rogue.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: BladeFlurryCooldownID,
		Cooldown:   cooldown,
		UsesGCD:    true,
		Type:       core.CooldownTypeDPS,
		Priority:   core.CooldownPriorityLow,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			if rogue.CurrentEnergy() < energyCost {
				return false
			}
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			if sim.GetRemainingDuration() > cooldown+dur {
				// We'll have enough time to cast another BF, so use it immediately to make sure we get the 2nd one.
				return true
			}

			// Since this is our last BF, wait until we have SND / procs up.
			sndTimeRemaining := rogue.RemainingAuraDuration(sim, SliceAndDiceAuraID)
			if sndTimeRemaining >= time.Second {
				return true
			}

			// TODO: Wait for dst/mongoose procs

			return false
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

var AdrenalineRushAuraID = core.NewAuraID()
var AdrenalineRushCooldownID = core.NewCooldownID()

func (rogue *Rogue) registerAdrenalineRushCD() {
	if !rogue.Talents.AdrenalineRush {
		return
	}

	actionID := core.ActionID{SpellID: 13750, CooldownID: AdrenalineRushCooldownID}

	adrenalineRushAura := core.Aura{
		ID:       AdrenalineRushAuraID,
		ActionID: actionID,
		Duration: time.Second * 15,
		OnGain: func(sim *core.Simulation) {
			rogue.ResetEnergyTick(sim)
			rogue.EnergyTickMultiplier = 2
		},
		OnExpire: func(sim *core.Simulation) {
			rogue.ResetEnergyTick(sim)
			rogue.EnergyTickMultiplier = 1
		},
	}

	cooldown := time.Minute * 5

	template := core.SimpleCast{
		Cast: core.Cast{
			ActionID:    actionID,
			Character:   rogue.GetCharacter(),
			Cooldown:    cooldown,
			GCD:         time.Second,
			IgnoreHaste: true,
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				rogue.AddAura(sim, adrenalineRushAura)
			},
		},
	}

	rogue.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: AdrenalineRushCooldownID,
		Cooldown:   cooldown,
		UsesGCD:    true,
		Type:       core.CooldownTypeDPS,
		Priority:   core.CooldownPriorityBloodlust,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			// Make sure we have plenty of room so the big ticks dont get wasted.
			thresh := 85.0
			if rogue.NextEnergyTickAt() < sim.CurrentTime+time.Second*1 {
				thresh = 60.0
			}
			if rogue.CurrentEnergy() > thresh {
				return false
			}
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
