package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func applyDebuffEffects(target *Target, debuffs proto.Debuffs) {
	if debuffs.Misery {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return MiseryAura(target, 5)
		})
	}

	if debuffs.JudgementOfWisdom {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return JudgementOfWisdomAura()
		})
	}

	if debuffs.ImprovedSealOfTheCrusader {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return JudgementOfTheCrusaderAura(target, 3)
		})
	}

	if debuffs.CurseOfElements != proto.TristateEffect_TristateEffectMissing {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return CurseOfElementsAura(target, debuffs.CurseOfElements)
		})
	}

	if debuffs.IsbUptime > 0.0 {
		uptime := MinFloat(1.0, debuffs.IsbUptime)
		target.AddPermanentAuraWithOptions(PermanentAura{
			AuraFactory: func(sim *Simulation) Aura {
				return ImprovedShadowBoltAura(target, uptime)
			},
			UptimeMultiplier: uptime,
		})
	}

	if debuffs.ImprovedScorch {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return ImprovedScorchAura(target, 5)
		})
	}

	if debuffs.WintersChill {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return WintersChillAura(target, 5)
		})
	}

	if debuffs.BloodFrenzy {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return BloodFrenzyAura(target)
		})
	}

	if debuffs.Mangle {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return MangleAura(target)
		})
	}

	if debuffs.ExposeArmor != proto.TristateEffect_TristateEffectMissing {
		points := int32(0)
		if debuffs.ExposeArmor == proto.TristateEffect_TristateEffectImproved {
			points = 2
		}
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return ExposeArmorAura(sim, target, points)
		})
	} else if debuffs.SunderArmor {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return SunderArmorAura(target, 5)
		})
	}

	if debuffs.FaerieFire != proto.TristateEffect_TristateEffectMissing {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return FaerieFireAura(target, debuffs.FaerieFire == proto.TristateEffect_TristateEffectImproved)
		})
	}

	if debuffs.CurseOfRecklessness {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return CurseOfRecklessnessAura(target)
		})
	}

	if debuffs.ExposeWeaknessUptime > 0 && debuffs.ExposeWeaknessHunterAgility > 0 {
		uptime := MinFloat(1.0, debuffs.ExposeWeaknessUptime)
		target.AddPermanentAuraWithOptions(PermanentAura{
			AuraFactory: func(sim *Simulation) Aura {
				return ExposeWeaknessAura(target, debuffs.ExposeWeaknessHunterAgility, uptime)
			},
			UptimeMultiplier: uptime,
		})
	}

	if debuffs.HuntersMark != proto.TristateEffect_TristateEffectMissing {
		if debuffs.HuntersMark == proto.TristateEffect_TristateEffectImproved {
			target.AddPermanentAura(func(sim *Simulation) Aura {
				return HuntersMarkAura(target, 5, true)
			})
		} else {
			target.AddPermanentAura(func(sim *Simulation) Aura {
				return HuntersMarkAura(target, 0, true)
			})
		}
	}
}

var MiseryAuraID = NewAuraID()

func MiseryAura(target *Target, numPoints int32) Aura {
	multiplier := 1.0 + 0.01*float64(numPoints)

	return Aura{
		ID:       MiseryAuraID,
		ActionID: ActionID{SpellID: 33195},
		Duration: time.Second * 24,
		Stacks:   numPoints,
		OnGain: func(sim *Simulation) {
			target.PseudoStats.ArcaneDamageTakenMultiplier *= multiplier
			target.PseudoStats.FireDamageTakenMultiplier *= multiplier
			target.PseudoStats.FrostDamageTakenMultiplier *= multiplier
			target.PseudoStats.HolyDamageTakenMultiplier *= multiplier
			target.PseudoStats.NatureDamageTakenMultiplier *= multiplier
			target.PseudoStats.ShadowDamageTakenMultiplier *= multiplier
		},
		OnExpire: func(sim *Simulation) {
			target.PseudoStats.ArcaneDamageTakenMultiplier /= multiplier
			target.PseudoStats.FireDamageTakenMultiplier /= multiplier
			target.PseudoStats.FrostDamageTakenMultiplier /= multiplier
			target.PseudoStats.HolyDamageTakenMultiplier /= multiplier
			target.PseudoStats.NatureDamageTakenMultiplier /= multiplier
			target.PseudoStats.ShadowDamageTakenMultiplier /= multiplier
		},
	}
}

var ShadowWeavingAuraID = NewAuraID()

func ShadowWeavingAura(target *Target, numStacks int32) Aura {
	multiplier := 1.0 + 0.02*float64(numStacks)

	return Aura{
		ID:        ShadowWeavingAuraID,
		ActionID:  ActionID{SpellID: 15334},
		Duration:  time.Second * 15,
		Stacks:    numStacks,
		MaxStacks: 5,
		OnGain: func(sim *Simulation) {
			target.PseudoStats.ShadowDamageTakenMultiplier *= multiplier
		},
		OnExpire: func(sim *Simulation) {
			target.PseudoStats.ShadowDamageTakenMultiplier /= multiplier
		},
	}
}

var JudgementOfWisdomAuraID = NewAuraID()

func JudgementOfWisdomAura() Aura {
	const mana = 74 / 2 // 50% proc
	actionID := ActionID{SpellID: 27164}

	return Aura{
		ID:       JudgementOfWisdomAuraID,
		ActionID: actionID,
		Duration: time.Second * 20,
		OnSpellHit: func(sim *Simulation, spell *Spell, spellEffect *SpellEffect) {
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
				spellEffect.Target.RefreshAura(sim, JudgementOfWisdomAuraID)
			}
		},
	}
}

var ImprovedSealOfTheCrusaderAuraID = NewAuraID()

func JudgementOfTheCrusaderAura(target *Target, level float64) Aura {
	bonusCrit := level * SpellCritRatingPerCritChance

	return Aura{
		ID:       ImprovedSealOfTheCrusaderAuraID,
		ActionID: ActionID{SpellID: 27159},
		Duration: time.Second * 20,
		OnGain: func(sim *Simulation) {
			target.PseudoStats.BonusHolyDamageTaken += 219
			target.PseudoStats.BonusCritRating += bonusCrit
		},
		OnExpire: func(sim *Simulation) {
			target.PseudoStats.BonusHolyDamageTaken -= 219
			target.PseudoStats.BonusCritRating -= bonusCrit
		},
		OnSpellHit: func(sim *Simulation, spell *Spell, spellEffect *SpellEffect) {
			if spell.ActionID.SpellID == 35395 {
				spellEffect.Target.RefreshAura(sim, ImprovedSealOfTheCrusaderAuraID)
			}
		},
	}
}

var CurseOfElementsAuraID = NewAuraID()

func CurseOfElementsAura(target *Target, coe proto.TristateEffect) Aura {
	multiplier := 1.1
	level := int32(0)
	if coe == proto.TristateEffect_TristateEffectImproved {
		multiplier = 1.13
		level = 3
	}

	return Aura{
		ID:       CurseOfElementsAuraID,
		ActionID: ActionID{SpellID: 27228},
		Stacks:   level, // Use stacks to store talent level for detection by other code.
		OnGain: func(sim *Simulation) {
			target.PseudoStats.ArcaneDamageTakenMultiplier *= multiplier
			target.PseudoStats.FireDamageTakenMultiplier *= multiplier
			target.PseudoStats.FrostDamageTakenMultiplier *= multiplier
			target.PseudoStats.ShadowDamageTakenMultiplier *= multiplier
		},
		OnExpire: func(sim *Simulation) {
			target.PseudoStats.ArcaneDamageTakenMultiplier /= multiplier
			target.PseudoStats.FireDamageTakenMultiplier /= multiplier
			target.PseudoStats.FrostDamageTakenMultiplier /= multiplier
			target.PseudoStats.ShadowDamageTakenMultiplier /= multiplier
		},
	}
}

var ImprovedShadowBoltID = NewAuraID()

func ImprovedShadowBoltAura(target *Target, uptime float64) Aura {
	multiplier := (1 + uptime*0.2)

	return Aura{
		ID:        ImprovedShadowBoltID,
		ActionID:  ActionID{SpellID: 17803},
		MaxStacks: 4,
		OnGain: func(sim *Simulation) {
			target.PseudoStats.ShadowDamageTakenMultiplier *= multiplier
		},
		OnExpire: func(sim *Simulation) {
			target.PseudoStats.ShadowDamageTakenMultiplier /= multiplier
		},
	}
}

var BloodFrenzyAuraID = NewAuraID()

func BloodFrenzyAura(target *Target) Aura {
	return Aura{
		ID:       BloodFrenzyAuraID,
		ActionID: ActionID{SpellID: 29859},
		OnGain: func(sim *Simulation) {
			target.PseudoStats.PhysicalDamageTakenMultiplier *= 1.04
		},
		OnExpire: func(sim *Simulation) {
			target.PseudoStats.PhysicalDamageTakenMultiplier /= 1.04
		},
	}
}

var MangleAuraID = NewAuraID()

func MangleAura(target *Target) Aura {
	return Aura{
		ID:       MangleAuraID,
		ActionID: ActionID{SpellID: 33876},
		OnGain: func(sim *Simulation) {
			target.PseudoStats.PeriodicPhysicalDamageTakenMultiplier *= 1.3
		},
		OnExpire: func(sim *Simulation) {
			target.PseudoStats.PeriodicPhysicalDamageTakenMultiplier /= 1.3
		},
	}
}

var ImprovedScorchAuraID = NewAuraID()

func ImprovedScorchAura(target *Target, numStacks int32) Aura {
	multiplier := 1.0 + 0.03*float64(numStacks)

	return Aura{
		ID:        ImprovedScorchAuraID,
		ActionID:  ActionID{SpellID: 12873},
		Duration:  time.Second * 30,
		Stacks:    numStacks,
		MaxStacks: 5,
		OnGain: func(sim *Simulation) {
			target.PseudoStats.FireDamageTakenMultiplier *= multiplier
		},
		OnExpire: func(sim *Simulation) {
			target.PseudoStats.FireDamageTakenMultiplier /= multiplier
		},
	}
}

var WintersChillAuraID = NewAuraID()

func WintersChillAura(target *Target, numStacks int32) Aura {
	bonusCrit := 2 * float64(numStacks) * SpellCritRatingPerCritChance

	return Aura{
		ID:        WintersChillAuraID,
		ActionID:  ActionID{SpellID: 28595},
		Duration:  time.Second * 15,
		Stacks:    numStacks,
		MaxStacks: 5,
		OnGain: func(sim *Simulation) {
			target.PseudoStats.BonusFrostCritRating += bonusCrit
		},
		OnExpire: func(sim *Simulation) {
			target.PseudoStats.BonusFrostCritRating -= bonusCrit
		},
	}
}

var FaerieFireAuraID = NewAuraID()

func FaerieFireAura(target *Target, improved bool) Aura {
	const hitBonus = 3 * MeleeHitRatingPerHitChance
	const armorReduction = 610

	return Aura{
		ID:       FaerieFireAuraID,
		ActionID: ActionID{SpellID: 26993},
		Duration: time.Second * 40,
		OnGain: func(sim *Simulation) {
			target.AddStat(stats.Armor, -armorReduction)
			if improved {
				target.PseudoStats.BonusMeleeHitRating += hitBonus
			}
		},
		OnExpire: func(sim *Simulation) {
			target.AddStat(stats.Armor, armorReduction)
			if improved {
				target.PseudoStats.BonusMeleeHitRating -= hitBonus
			}
		},
	}
}

var SunderArmorAuraID = NewAuraID()

func SunderArmorAura(target *Target, stacks int32) Aura {
	armorReduction := 520.0 * float64(stacks)

	return Aura{
		ID:        SunderArmorAuraID,
		ActionID:  ActionID{SpellID: 25225},
		Duration:  time.Second * 30,
		Stacks:    stacks,
		MaxStacks: 5,
		OnGain: func(sim *Simulation) {
			target.AddStat(stats.Armor, -armorReduction)
		},
		OnExpire: func(sim *Simulation) {
			target.AddStat(stats.Armor, armorReduction)
		},
	}
}

var ExposeArmorAuraID = NewAuraID()

func ExposeArmorAura(sim *Simulation, target *Target, talentPoints int32) Aura {
	armorReduction := 2050.0 * (1.0 + 0.25*float64(talentPoints))

	return Aura{
		ID:       ExposeArmorAuraID,
		ActionID: ActionID{SpellID: 26866},
		Duration: time.Second * 30,
		OnGain: func(sim *Simulation) {
			if target.HasAura(SunderArmorAuraID) {
				target.RemoveAura(sim, SunderArmorAuraID)
			}
			target.AddStat(stats.Armor, -armorReduction)
		},
		OnExpire: func(sim *Simulation) {
			target.AddStat(stats.Armor, armorReduction)
		},
	}
}

var CurseOfRecklessnessAuraID = NewAuraID()

func CurseOfRecklessnessAura(target *Target) Aura {
	armorReduction := 800.0

	return Aura{
		ID:       CurseOfRecklessnessAuraID,
		ActionID: ActionID{SpellID: 27226},
		Duration: time.Minute * 2,
		OnGain: func(sim *Simulation) {
			target.AddStat(stats.Armor, -armorReduction)
		},
		OnExpire: func(sim *Simulation) {
			target.AddStat(stats.Armor, armorReduction)
		},
	}
}

var ExposeWeaknessAuraID = NewAuraID()

// Multiplier is for accomodating uptime %. For a real hunter, always pass 1.0
func ExposeWeaknessAura(target *Target, hunterAgility float64, multiplier float64) Aura {
	apBonus := hunterAgility * 0.25 * multiplier

	return Aura{
		ID:       ExposeWeaknessAuraID,
		ActionID: ActionID{SpellID: 34503},
		Duration: time.Second * 7,
		OnGain: func(sim *Simulation) {
			target.PseudoStats.BonusMeleeAttackPower += apBonus
			target.PseudoStats.BonusRangedAttackPower += apBonus
		},
		OnExpire: func(sim *Simulation) {
			target.PseudoStats.BonusMeleeAttackPower -= apBonus
			target.PseudoStats.BonusRangedAttackPower -= apBonus
		},
	}
}

var HuntersMarkAuraID = NewAuraID()

func HuntersMarkAura(target *Target, points int32, fullyStacked bool) Aura {
	const baseRangedBonus = 110.0
	const bonusPerStack = 11.0
	const maxStacks = 30
	meleeBonus := baseRangedBonus * 0.2 * float64(points)

	stacks := 0
	if fullyStacked {
		stacks = maxStacks
	}

	rangedBonus := baseRangedBonus + bonusPerStack*float64(stacks)

	return Aura{
		ID:       HuntersMarkAuraID,
		ActionID: ActionID{SpellID: 14325},
		Stacks:   points, // Use this to check whether to override in hunter/hunter.go
		Duration: NeverExpires,
		OnGain: func(sim *Simulation) {
			target.PseudoStats.BonusMeleeAttackPower += meleeBonus
			target.PseudoStats.BonusRangedAttackPower += rangedBonus
		},
		OnExpire: func(sim *Simulation) {
			target.PseudoStats.BonusMeleeAttackPower -= meleeBonus
			target.PseudoStats.BonusRangedAttackPower -= rangedBonus
		},
		OnSpellHit: func(sim *Simulation, spell *Spell, spellEffect *SpellEffect) {
			if !spellEffect.OutcomeRollCategory.Matches(OutcomeRollCategoryRanged) || !spellEffect.Landed() {
				return
			}
			if stacks < maxStacks {
				stacks++
				target.PseudoStats.BonusRangedAttackPower += bonusPerStack
				rangedBonus += bonusPerStack
			}
		},
	}
}

var DemoralizingShoutAuraID = NewAuraID()

func DemoralizingShoutAura(target *Target, boomingVoicePts int32, impDemoShoutPts int32) Aura {
	duration := time.Duration(float64(time.Second*30) * (1 + 0.1*float64(boomingVoicePts)))
	return Aura{
		ID:       DemoralizingShoutAuraID,
		ActionID: ActionID{SpellID: 25203},
		Duration: duration,
		OnGain: func(sim *Simulation) {
		},
		OnExpire: func(sim *Simulation) {
		},
	}
}

var ThunderClapAuraID = NewAuraID()

func ThunderClapAura(target *Target, impThunderClapPts int32) Aura {
	return Aura{
		ID:       ThunderClapAuraID,
		ActionID: ActionID{SpellID: 25264},
		Duration: time.Second * 30,
		OnGain: func(sim *Simulation) {
		},
		OnExpire: func(sim *Simulation) {
		},
	}
}
