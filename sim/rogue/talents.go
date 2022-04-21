package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (rogue *Rogue) ApplyTalents() {
	// TODO: Last few talents in the sub tree.

	rogue.applyMurder()
	rogue.applySealFate()
	rogue.applyWeaponSpecializations()
	rogue.applyCombatPotency()

	rogue.AddStat(stats.Dodge, core.DodgeRatingPerDodgeChance*1*float64(rogue.Talents.LightningReflexes))
	rogue.AddStat(stats.Parry, core.ParryRatingPerParryChance*1*float64(rogue.Talents.Deflection))
	rogue.AddStat(stats.MeleeCrit, core.MeleeCritRatingPerCritChance*1*float64(rogue.Talents.Malice))
	rogue.AddStat(stats.MeleeHit, core.MeleeHitRatingPerHitChance*1*float64(rogue.Talents.Precision))
	rogue.AddStat(stats.Expertise, core.ExpertisePerQuarterPercentReduction*5*float64(rogue.Talents.WeaponExpertise))
	rogue.AddStat(stats.ArmorPenetration, 186*float64(rogue.Talents.SerratedBlades))

	if rogue.Talents.DualWieldSpecialization > 0 {
		rogue.AutoAttacks.OHEffect.BaseDamage.Calculator = core.BaseDamageFuncMeleeWeapon(core.OffHand, false, 0, 1+0.1*float64(rogue.Talents.DualWieldSpecialization), true)
	}

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

func (rogue *Rogue) makeFinishingMoveEffectApplier(_ *core.Simulation) func(sim *core.Simulation, numPoints int32) {
	ruthlessnessChance := 0.2 * float64(rogue.Talents.Ruthlessness)
	relentlessStrikes := rogue.Talents.RelentlessStrikes

	var fwAura *core.Aura
	findWeaknessMultiplier := 1.0 + 0.02*float64(rogue.Talents.FindWeakness)
	if findWeaknessMultiplier != 1 {
		fwAura = rogue.GetOrRegisterAura(core.Aura{
			Label:    "Find Weakness",
			ActionID: core.ActionID{SpellID: 31242},
			Duration: time.Second * 10,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.PseudoStats.AgentReserved1DamageDealtMultiplier *= findWeaknessMultiplier
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.PseudoStats.AgentReserved1DamageDealtMultiplier /= findWeaknessMultiplier
			},
		})
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
		if fwAura != nil {
			fwAura.Activate(sim)
		}
	}
}

func (rogue *Rogue) applyMurder() {
	if rogue.Talents.Murder == 0 {
		return
	}

	damageMultiplier := 1.0 + 0.01*float64(rogue.Talents.Murder)

	rogue.RegisterResetEffect(func(sim *core.Simulation) {
		switch sim.GetPrimaryTarget().MobType {
		case proto.MobType_MobTypeHumanoid, proto.MobType_MobTypeBeast, proto.MobType_MobTypeGiant, proto.MobType_MobTypeDragonkin:
			rogue.PseudoStats.DamageDealtMultiplier *= damageMultiplier
		}
	})
}

func (rogue *Rogue) registerColdBloodCD() {
	if !rogue.Talents.ColdBlood {
		return
	}

	actionID := core.ActionID{SpellID: 14177}

	coldBloodAura := rogue.RegisterAura(core.Aura{
		Label:    "Cold Blood",
		ActionID: actionID,
		Duration: core.NeverExpires,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.PseudoStats.BonusCritRatingAgentReserved1 += 100 * core.MeleeCritRatingPerCritChance
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.PseudoStats.BonusCritRatingAgentReserved1 -= 100 * core.MeleeCritRatingPerCritChance
		},
		OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			aura.Deactivate(sim)
		},
	})

	coldBloodSpell := rogue.RegisterSpell(core.SpellConfig{
		ActionID: actionID,

		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    rogue.NewTimer(),
				Duration: time.Minute * 3,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Target, spell *core.Spell) {
			coldBloodAura.Activate(sim)
		},
	})

	rogue.AddMajorCooldown(core.MajorCooldown{
		Spell: coldBloodSpell,
		Type:  core.CooldownTypeDPS,
	})
}

func (rogue *Rogue) applySealFate() {
	if rogue.Talents.SealFate == 0 {
		return
	}

	procChance := 0.2 * float64(rogue.Talents.SealFate)

	rogue.RegisterAura(core.Aura{
		Label:    "Seal Fate",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spell.SpellExtras.Matches(SpellFlagBuilder) {
				return
			}

			if !spellEffect.Outcome.Matches(core.OutcomeCrit) {
				return
			}

			if procChance == 1 || sim.RandomFloat("Seal Fate") < procChance {
				rogue.AddComboPoints(sim, 1, core.ActionID{SpellID: 14195})
			}
		},
	})
}

func (rogue *Rogue) applyWeaponSpecializations() {
	if weapon := rogue.Equip[proto.ItemSlot_ItemSlotMainHand]; weapon.ID != 0 {
		if weapon.WeaponType == proto.WeaponType_WeaponTypeFist {
			rogue.PseudoStats.BonusMHCritRating += 1 * core.MeleeCritRatingPerCritChance * float64(rogue.Talents.FistWeaponSpecialization)
		} else if weapon.WeaponType == proto.WeaponType_WeaponTypeDagger {
			rogue.PseudoStats.BonusMHCritRating += 1 * core.MeleeCritRatingPerCritChance * float64(rogue.Talents.DaggerSpecialization)
		}
	}
	if weapon := rogue.Equip[proto.ItemSlot_ItemSlotOffHand]; weapon.ID != 0 {
		if weapon.WeaponType == proto.WeaponType_WeaponTypeFist {
			rogue.PseudoStats.BonusOHCritRating += 1 * core.MeleeCritRatingPerCritChance * float64(rogue.Talents.FistWeaponSpecialization)
		} else if weapon.WeaponType == proto.WeaponType_WeaponTypeDagger {
			rogue.PseudoStats.BonusOHCritRating += 1 * core.MeleeCritRatingPerCritChance * float64(rogue.Talents.DaggerSpecialization)
		}
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
		var swordSpecializationSpell *core.Spell
		icd := core.Cooldown{
			Timer:    rogue.NewTimer(),
			Duration: time.Millisecond * 500,
		}
		procChance := 0.01 * float64(rogue.Talents.SwordSpecialization)

		rogue.RegisterAura(core.Aura{
			Label:    "Sword Specialization",
			Duration: core.NeverExpires,
			OnInit: func(aura *core.Aura, sim *core.Simulation) {
				swordSpecializationSpell = rogue.GetOrRegisterSpell(core.SpellConfig{
					ActionID:    core.ActionID{SpellID: 13964},
					SpellSchool: core.SpellSchoolPhysical,
					SpellExtras: core.SpellExtrasMeleeMetrics,

					ApplyEffects: core.ApplyEffectFuncDirectDamage(rogue.AutoAttacks.MHEffect),
				})
			},
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					return
				}

				if !spellEffect.ProcMask.Matches(swordSpecMask) {
					return
				}

				if !icd.IsReady(sim) {
					return
				}

				if sim.RandomFloat("Sword Specialization") > procChance {
					return
				}
				icd.Use(sim)

				swordSpecializationSpell.Cast(sim, spellEffect.Target)
			},
		})
	}
}

func (rogue *Rogue) applyCombatPotency() {
	if rogue.Talents.CombatPotency == 0 {
		return
	}

	const procChance = 0.2
	energyBonus := 3.0 * float64(rogue.Talents.CombatPotency)

	rogue.RegisterAura(core.Aura{
		Label:    "Combat Potency",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
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
	})
}

func (rogue *Rogue) registerBladeFlurryCD() {
	if !rogue.Talents.BladeFlurry {
		return
	}

	actionID := core.ActionID{SpellID: 13877}
	const hasteBonus = 1.2
	const inverseHasteBonus = 1 / 1.2
	const energyCost = 25.0

	dur := time.Second * 15

	rogue.BladeFlurryAura = rogue.RegisterAura(core.Aura{
		Label:    "Blade Flurry",
		ActionID: actionID,
		Duration: dur,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			rogue.MultiplyMeleeSpeed(sim, hasteBonus)
			if sim.GetNumTargets() > 1 {
				rogue.PseudoStats.DamageDealtMultiplier *= 2
			}
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			rogue.MultiplyMeleeSpeed(sim, inverseHasteBonus)
			if sim.GetNumTargets() > 1 {
				rogue.PseudoStats.DamageDealtMultiplier /= 2
			}
		},
	})

	cooldownDur := time.Minute * 2
	bladeFlurrySpell := rogue.RegisterSpell(core.SpellConfig{
		ActionID: actionID,

		ResourceType: stats.Energy,
		BaseCost:     energyCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: energyCost,
				GCD:  time.Second,
			},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    rogue.NewTimer(),
				Duration: cooldownDur,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Target, spell *core.Spell) {
			rogue.BladeFlurryAura.Activate(sim)
		},
	})

	rogue.AddMajorCooldown(core.MajorCooldown{
		Spell:    bladeFlurrySpell,
		Type:     core.CooldownTypeDPS,
		Priority: core.CooldownPriorityLow,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			return rogue.CurrentEnergy() >= energyCost
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			if sim.GetRemainingDuration() > cooldownDur+dur {
				// We'll have enough time to cast another BF, so use it immediately to make sure we get the 2nd one.
				return true
			}

			// Since this is our last BF, wait until we have SND / procs up.
			sndTimeRemaining := rogue.SliceAndDiceAura.RemainingDuration(sim)
			if sndTimeRemaining >= time.Second {
				return true
			}

			// TODO: Wait for dst/mongoose procs

			return false
		},
	})
}

func (rogue *Rogue) registerAdrenalineRushCD() {
	if !rogue.Talents.AdrenalineRush {
		return
	}

	actionID := core.ActionID{SpellID: 13750}

	rogue.AdrenalineRushAura = rogue.RegisterAura(core.Aura{
		Label:    "Adrenaline Rush",
		ActionID: actionID,
		Duration: time.Second * 15,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			rogue.ResetEnergyTick(sim)
			rogue.EnergyTickMultiplier = 2
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			rogue.ResetEnergyTick(sim)
			rogue.EnergyTickMultiplier = 1
		},
	})

	adrenalineRushSpell := rogue.RegisterSpell(core.SpellConfig{
		ActionID: actionID,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: time.Second,
			},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    rogue.NewTimer(),
				Duration: time.Minute * 5,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Target, spell *core.Spell) {
			rogue.AdrenalineRushAura.Activate(sim)
		},
	})

	rogue.AddMajorCooldown(core.MajorCooldown{
		Spell:    adrenalineRushSpell,
		Type:     core.CooldownTypeDPS,
		Priority: core.CooldownPriorityBloodlust,
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
	})
}
