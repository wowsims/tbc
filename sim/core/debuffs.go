package core

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func applyDebuffEffects(target *Target, debuffs proto.Debuffs) {
	if debuffs.Misery {
		target.AddPermanentAura(func(*Simulation) *Aura { return MiseryAura(target, 5) })
	}

	if debuffs.JudgementOfWisdom {
		target.AddPermanentAura(func(*Simulation) *Aura { return JudgementOfWisdomAura(target) })
	}

	if debuffs.ImprovedSealOfTheCrusader {
		target.AddPermanentAura(func(*Simulation) *Aura { return JudgementOfTheCrusaderAura(target, 3) })
	}

	if debuffs.CurseOfElements != proto.TristateEffect_TristateEffectMissing {
		target.AddPermanentAura(func(*Simulation) *Aura {
			return CurseOfElementsAura(target, GetTristateValueInt32(debuffs.CurseOfElements, 0, 3))
		})
	}

	if debuffs.IsbUptime > 0.0 {
		uptime := MinFloat(1.0, debuffs.IsbUptime)
		target.AddPermanentAuraWithOptions(PermanentAura{
			AuraFactory:      func(*Simulation) *Aura { return ImprovedShadowBoltAura(target, uptime) },
			UptimeMultiplier: uptime,
		})
	}

	if debuffs.ImprovedScorch {
		target.AddPermanentAura(func(*Simulation) *Aura { return ImprovedScorchAura(target, 5) })
	}

	if debuffs.WintersChill {
		target.AddPermanentAura(func(*Simulation) *Aura { return WintersChillAura(target, 5) })
	}

	if debuffs.BloodFrenzy {
		target.AddPermanentAura(func(*Simulation) *Aura { return BloodFrenzyAura(target) })
	}

	if debuffs.Mangle {
		target.AddPermanentAura(func(*Simulation) *Aura { return MangleAura(target) })
	}

	if debuffs.ExposeArmor != proto.TristateEffect_TristateEffectMissing {
		target.AddPermanentAura(func(*Simulation) *Aura {
			return ExposeArmorAura(target, GetTristateValueInt32(debuffs.ExposeArmor, 0, 2))
		})
	}

	if debuffs.SunderArmor {
		target.AddPermanentAura(func(*Simulation) *Aura { return SunderArmorAura(target, 5) })
	}

	if debuffs.FaerieFire != proto.TristateEffect_TristateEffectMissing {
		target.AddPermanentAura(func(*Simulation) *Aura {
			return FaerieFireAura(target, GetTristateValueInt32(debuffs.FaerieFire, 0, 3))
		})
	}

	if debuffs.CurseOfRecklessness {
		target.AddPermanentAura(func(*Simulation) *Aura { return CurseOfRecklessnessAura(target) })
	}

	if debuffs.ExposeWeaknessUptime > 0 && debuffs.ExposeWeaknessHunterAgility > 0 {
		uptime := MinFloat(1.0, debuffs.ExposeWeaknessUptime)
		target.AddPermanentAuraWithOptions(PermanentAura{
			AuraFactory: func(*Simulation) *Aura {
				return ExposeWeaknessAura(target, debuffs.ExposeWeaknessHunterAgility, uptime)
			},
			UptimeMultiplier: uptime,
		})
	}

	if debuffs.HuntersMark != proto.TristateEffect_TristateEffectMissing {
		if debuffs.HuntersMark == proto.TristateEffect_TristateEffectImproved {
			target.AddPermanentAura(func(*Simulation) *Aura { return HuntersMarkAura(target, 5, true) })
		} else {
			target.AddPermanentAura(func(*Simulation) *Aura { return HuntersMarkAura(target, 0, true) })
		}
	}
}

func MiseryAura(target *Target, numPoints int32) *Aura {
	multiplier := 1.0 + 0.01*float64(numPoints)

	return target.GetOrRegisterAura(&Aura{
		Label:    "Misery-" + strconv.Itoa(int(numPoints)),
		Tag:      "Misery",
		ActionID: ActionID{SpellID: 33195},
		Duration: time.Second * 24,
		Priority: numPoints,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.ArcaneDamageTakenMultiplier *= multiplier
			aura.Unit.PseudoStats.FireDamageTakenMultiplier *= multiplier
			aura.Unit.PseudoStats.FrostDamageTakenMultiplier *= multiplier
			aura.Unit.PseudoStats.HolyDamageTakenMultiplier *= multiplier
			aura.Unit.PseudoStats.NatureDamageTakenMultiplier *= multiplier
			aura.Unit.PseudoStats.ShadowDamageTakenMultiplier *= multiplier
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.ArcaneDamageTakenMultiplier /= multiplier
			aura.Unit.PseudoStats.FireDamageTakenMultiplier /= multiplier
			aura.Unit.PseudoStats.FrostDamageTakenMultiplier /= multiplier
			aura.Unit.PseudoStats.HolyDamageTakenMultiplier /= multiplier
			aura.Unit.PseudoStats.NatureDamageTakenMultiplier /= multiplier
			aura.Unit.PseudoStats.ShadowDamageTakenMultiplier /= multiplier
		},
	})
}

func ShadowWeavingAura(target *Target, startingStacks int32) *Aura {
	return target.GetOrRegisterAura(&Aura{
		Label:     "Shadow Weaving",
		ActionID:  ActionID{SpellID: 15334},
		Duration:  time.Second * 15,
		MaxStacks: 5,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.SetStacks(sim, startingStacks)
		},
		OnStacksChange: func(aura *Aura, sim *Simulation, oldStacks int32, newStacks int32) {
			aura.Unit.PseudoStats.ShadowDamageTakenMultiplier /= 1.0 + 0.02*float64(oldStacks)
			aura.Unit.PseudoStats.ShadowDamageTakenMultiplier *= 1.0 + 0.02*float64(newStacks)
		},
	})
}

func JudgementOfWisdomAura(target *Target) *Aura {
	const mana = 74 / 2 // 50% proc
	actionID := ActionID{SpellID: 27164}

	return target.GetOrRegisterAura(&Aura{
		Label:    "Judgement of Wisdom",
		ActionID: actionID,
		Duration: time.Second * 20,
		OnSpellHit: func(aura *Aura, sim *Simulation, spell *Spell, spellEffect *SpellEffect) {
			// TODO: This check is purely to maintain behavior during refactoring. Should be removed when possible.
			if !spellEffect.ProcMask.Matches(ProcMaskMeleeOrRanged) && !spellEffect.Landed() {
				return
			}
			if spellEffect.IsPhantom {
				return // Phantom spells (Romulo's, Lightning Capacitor, etc) don't proc JoW.
			}

			character := spell.Character
			if character.HasManaBar() {
				character.AddMana(sim, mana, actionID, false)
			}

			if spell.ActionID.SpellID == 35395 {
				aura.Refresh(sim)
			}
		},
	})
}

func JudgementOfTheCrusaderAura(target *Target, level int32) *Aura {
	bonusCrit := float64(level) * SpellCritRatingPerCritChance

	return target.GetOrRegisterAura(&Aura{
		Label:    "Judgement of the Crusader-" + strconv.Itoa(int(level)),
		Tag:      "Judgement of the Crusader",
		ActionID: ActionID{SpellID: 27159},
		Duration: time.Second * 20,
		Priority: level,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.BonusHolyDamageTaken += 219
			aura.Unit.PseudoStats.BonusCritRating += bonusCrit
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.BonusHolyDamageTaken -= 219
			aura.Unit.PseudoStats.BonusCritRating -= bonusCrit
		},
		OnSpellHit: func(aura *Aura, sim *Simulation, spell *Spell, spellEffect *SpellEffect) {
			if spell.ActionID.SpellID == 35395 {
				aura.Refresh(sim)
			}
		},
	})
}

func CurseOfElementsAura(target *Target, points int32) *Aura {
	multiplier := 1.1 + 0.01*float64(points)

	return target.GetOrRegisterAura(&Aura{
		Label:    "Curse of Elements-" + strconv.Itoa(int(points)),
		Tag:      "Curse of Elements",
		ActionID: ActionID{SpellID: 27228},
		Priority: points,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.ArcaneDamageTakenMultiplier *= multiplier
			aura.Unit.PseudoStats.FireDamageTakenMultiplier *= multiplier
			aura.Unit.PseudoStats.FrostDamageTakenMultiplier *= multiplier
			aura.Unit.PseudoStats.ShadowDamageTakenMultiplier *= multiplier
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.ArcaneDamageTakenMultiplier /= multiplier
			aura.Unit.PseudoStats.FireDamageTakenMultiplier /= multiplier
			aura.Unit.PseudoStats.FrostDamageTakenMultiplier /= multiplier
			aura.Unit.PseudoStats.ShadowDamageTakenMultiplier /= multiplier
		},
	})
}

func ImprovedShadowBoltAura(target *Target, uptime float64) *Aura {
	multiplier := (1 + uptime*0.2)

	return target.GetOrRegisterAura(&Aura{
		Label:     "Improved Shadow Bolt",
		ActionID:  ActionID{SpellID: 17803},
		MaxStacks: 4,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.ShadowDamageTakenMultiplier *= multiplier
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.ShadowDamageTakenMultiplier /= multiplier
		},
	})
}

func BloodFrenzyAura(target *Target) *Aura {
	return target.GetOrRegisterAura(&Aura{
		Label:    "Blood Frenzy",
		ActionID: ActionID{SpellID: 29859},
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.PhysicalDamageTakenMultiplier *= 1.04
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.PhysicalDamageTakenMultiplier /= 1.04
		},
	})
}

func MangleAura(target *Target) *Aura {
	return target.GetOrRegisterAura(&Aura{
		Label:    "Mangle",
		ActionID: ActionID{SpellID: 33876},
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.PeriodicPhysicalDamageTakenMultiplier *= 1.3
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.PeriodicPhysicalDamageTakenMultiplier /= 1.3
		},
	})
}

var ImprovedScorchAuraLabel = "Improved Scorch"

func ImprovedScorchAura(target *Target, startingStacks int32) *Aura {
	return target.GetOrRegisterAura(&Aura{
		Label:     ImprovedScorchAuraLabel,
		ActionID:  ActionID{SpellID: 12873},
		Duration:  time.Second * 30,
		MaxStacks: 5,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.SetStacks(sim, startingStacks)
		},
		OnStacksChange: func(aura *Aura, sim *Simulation, oldStacks int32, newStacks int32) {
			aura.Unit.PseudoStats.FireDamageTakenMultiplier /= 1.0 + 0.03*float64(oldStacks)
			aura.Unit.PseudoStats.FireDamageTakenMultiplier *= 1.0 + 0.03*float64(newStacks)
		},
	})
}

var WintersChillAuraLabel = "Winter's Chill"

func WintersChillAura(target *Target, startingStacks int32) *Aura {
	return target.GetOrRegisterAura(&Aura{
		Label:     WintersChillAuraLabel,
		ActionID:  ActionID{SpellID: 28595},
		Duration:  time.Second * 15,
		MaxStacks: 5,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.SetStacks(sim, startingStacks)
		},
		OnStacksChange: func(aura *Aura, sim *Simulation, oldStacks int32, newStacks int32) {
			aura.Unit.PseudoStats.BonusFrostCritRating += 2 * SpellCritRatingPerCritChance * float64(newStacks-oldStacks)
		},
	})
}

func FaerieFireAura(target *Target, level int32) *Aura {
	const armorReduction = 610

	return target.GetOrRegisterAura(&Aura{
		Label:    "Faerie Fire-" + strconv.Itoa(int(level)),
		Tag:      "Faerie Fire",
		ActionID: ActionID{SpellID: 26993},
		Duration: time.Second * 40,
		Priority: level,
		OnGain: func(aura *Aura, sim *Simulation) {
			target.AddStat(stats.Armor, -armorReduction)
			target.PseudoStats.BonusMeleeHitRating += float64(level) * MeleeHitRatingPerHitChance
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			target.AddStat(stats.Armor, armorReduction)
			target.PseudoStats.BonusMeleeHitRating -= float64(level) * MeleeHitRatingPerHitChance
		},
	})
}

var SunderArmorAuraLabel = "Sunder Armor"
var SunderExposeAuraTag = "SunderExpose"

func SunderArmorAura(target *Target, startingStacks int32) *Aura {
	armorReductionPerStack := 520.0

	return target.GetOrRegisterAura(&Aura{
		Label:     SunderArmorAuraLabel,
		Tag:       SunderExposeAuraTag,
		ActionID:  ActionID{SpellID: 25225},
		Duration:  time.Second * 30,
		MaxStacks: 5,
		Priority:  int32(armorReductionPerStack * 5),
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.SetStacks(sim, startingStacks)
		},
		OnStacksChange: func(aura *Aura, sim *Simulation, oldStacks int32, newStacks int32) {
			target.AddStat(stats.Armor, float64(newStacks-oldStacks)*armorReductionPerStack)
		},
	})
}

func ExposeArmorAura(target *Target, talentPoints int32) *Aura {
	armorReduction := 2050.0 * (1.0 + 0.25*float64(talentPoints))

	return target.GetOrRegisterAura(&Aura{
		Label:    "ExposeArmor-" + strconv.Itoa(int(talentPoints)),
		Tag:      SunderExposeAuraTag,
		ActionID: ActionID{SpellID: 26866},
		Duration: time.Second * 30,
		Priority: int32(armorReduction),
		OnGain: func(aura *Aura, sim *Simulation) {
			target.AddStat(stats.Armor, -armorReduction)
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			target.AddStat(stats.Armor, armorReduction)
		},
	})
}

func CurseOfRecklessnessAura(target *Target) *Aura {
	armorReduction := 800.0

	return target.GetOrRegisterAura(&Aura{
		Label:    "Curse of Recklessness",
		ActionID: ActionID{SpellID: 27226},
		Duration: time.Minute * 2,
		OnGain: func(aura *Aura, sim *Simulation) {
			target.AddStat(stats.Armor, -armorReduction)
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			target.AddStat(stats.Armor, armorReduction)
		},
	})
}

// Multiplier is for accomodating uptime %. For a real hunter, always pass 1.0
func ExposeWeaknessAura(target *Target, hunterAgility float64, multiplier float64) *Aura {
	apBonus := hunterAgility * 0.25 * multiplier

	return target.GetOrRegisterAura(&Aura{
		Label:    "ExposeWeakness-" + strconv.Itoa(int(hunterAgility)),
		Tag:      "ExposeWeakness",
		ActionID: ActionID{SpellID: 34503},
		Duration: time.Second * 7,
		Priority: int32(hunterAgility),
		OnGain: func(aura *Aura, sim *Simulation) {
			target.PseudoStats.BonusMeleeAttackPower += apBonus
			target.PseudoStats.BonusRangedAttackPower += apBonus
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			target.PseudoStats.BonusMeleeAttackPower -= apBonus
			target.PseudoStats.BonusRangedAttackPower -= apBonus
		},
	})
}

func HuntersMarkAura(target *Target, points int32, fullyStacked bool) *Aura {
	const baseRangedBonus = 110.0
	const bonusPerStack = 11.0
	const maxStacks = 30
	meleeBonus := baseRangedBonus * 0.2 * float64(points)

	startingStacks := int32(0)
	if fullyStacked {
		startingStacks = maxStacks
	}

	return target.GetOrRegisterAura(&Aura{
		Label:     "HuntersMark-" + strconv.Itoa(int(points)),
		Tag:       "HuntersMark",
		ActionID:  ActionID{SpellID: 14325},
		Duration:  NeverExpires,
		MaxStacks: 30,
		Priority:  points,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.BonusMeleeAttackPower += meleeBonus
			aura.Unit.PseudoStats.BonusRangedAttackPower += baseRangedBonus
			aura.SetStacks(sim, startingStacks)
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.BonusMeleeAttackPower -= meleeBonus
			aura.Unit.PseudoStats.BonusRangedAttackPower -= baseRangedBonus
		},
		OnStacksChange: func(aura *Aura, sim *Simulation, oldStacks int32, newStacks int32) {
			aura.Unit.PseudoStats.BonusRangedAttackPower += bonusPerStack * float64(newStacks-oldStacks)
		},
		OnSpellHit: func(aura *Aura, sim *Simulation, spell *Spell, spellEffect *SpellEffect) {
			if spellEffect.OutcomeRollCategory.Matches(OutcomeRollCategoryRanged) && spellEffect.Landed() {
				aura.AddStack(sim)
			}
		},
	})
}

func DemoralizingShoutAura(target *Target, boomingVoicePts int32, impDemoShoutPts int32) *Aura {
	duration := time.Duration(float64(time.Second*30) * (1 + 0.1*float64(boomingVoicePts)))

	return target.GetOrRegisterAura(&Aura{
		Label:    "DemoralizingShout-" + strconv.Itoa(int(impDemoShoutPts)),
		Tag:      "DemoralizingShout",
		ActionID: ActionID{SpellID: 25203},
		Duration: duration,
		Priority: impDemoShoutPts,
		OnGain: func(aura *Aura, sim *Simulation) {
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
		},
	})
}

func ThunderClapAura(target *Target, impThunderClapPts int32) *Aura {
	return target.GetOrRegisterAura(&Aura{
		Label:    "ThunderClap-" + strconv.Itoa(int(impThunderClapPts)),
		Tag:      "ThunderClap",
		ActionID: ActionID{SpellID: 25264},
		Duration: time.Second * 30,
		Priority: impThunderClapPts,
		OnGain: func(aura *Aura, sim *Simulation) {
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
		},
	})
}
