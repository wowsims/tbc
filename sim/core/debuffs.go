package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func applyDebuffEffects(target *Target, debuffs proto.Debuffs) {
	if debuffs.Misery {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return MiseryAura(sim, 5)
		})
	}

	if debuffs.JudgementOfWisdom {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return JudgementOfWisdomAura(sim)
		})
	}

	if debuffs.ImprovedSealOfTheCrusader {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return JudgementOfTheCrusaderAura(sim, 3)
		})
	}

	if debuffs.CurseOfElements != proto.TristateEffect_TristateEffectMissing {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return CurseOfElementsAura(debuffs.CurseOfElements)
		})
	}

	if debuffs.IsbUptime > 0.0 {
		uptime := MinFloat(1.0, debuffs.IsbUptime)
		target.AddPermanentAuraWithOptions(PermanentAura{
			AuraFactory: func(sim *Simulation) Aura {
				return ImprovedShadowBoltAura(uptime)
			},
			UptimeMultiplier: uptime,
		})
	}

	if debuffs.ImprovedScorch {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return ImprovedScorchAura(sim, 5)
		})
	}

	if debuffs.WintersChill {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return WintersChillAura(sim, 5)
		})
	}

	if debuffs.BloodFrenzy {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return BloodFrenzyAura()
		})
	}

	if debuffs.ExposeArmor != proto.TristateEffect_TristateEffectMissing {
		points := 0
		if debuffs.ExposeArmor == proto.TristateEffect_TristateEffectImproved {
			points = 2
		}
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return ExposeArmorAura(0, target, points)
		})
	} else if debuffs.SunderArmor {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return SunderArmorAura(0, target, 5)
		})
	}

	if debuffs.FaerieFire != proto.TristateEffect_TristateEffectMissing {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return FaerieFireAura(0, target, debuffs.FaerieFire == proto.TristateEffect_TristateEffectImproved)
		})
	}

	if debuffs.CurseOfRecklessness {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return CurseOfRecklessnessAura(0, target)
		})
	}

	if debuffs.ExposeWeaknessUptime > 0 && debuffs.ExposeWeaknessHunterAgility > 0 {
		multiplier := MinFloat(1.0, debuffs.ExposeWeaknessUptime)
		target.AddPermanentAuraWithOptions(PermanentAura{
			AuraFactory: func(sim *Simulation) Aura {
				return ExposeWeaknessAura(0, debuffs.ExposeWeaknessHunterAgility, multiplier)
			},
			UptimeMultiplier: multiplier,
		})
	}

	if debuffs.HuntersMark != proto.TristateEffect_TristateEffectMissing {
		if debuffs.HuntersMark == proto.TristateEffect_TristateEffectImproved {
			target.AddPermanentAura(func(sim *Simulation) Aura {
				return HuntersMarkAura(5, true)
			})
		} else {
			target.AddPermanentAura(func(sim *Simulation) Aura {
				return HuntersMarkAura(0, true)
			})
		}
	}
}

var MiseryDebuffID = NewDebuffID()

func MiseryAura(sim *Simulation, numPoints int32) Aura {
	multiplier := 1.0 + 0.01*float64(numPoints)

	return Aura{
		ID:       MiseryDebuffID,
		ActionID: ActionID{SpellID: 33195},
		Expires:  sim.CurrentTime + time.Second*24,
		Stacks:   numPoints,
		OnBeforeSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect) {
			spellEffect.DamageMultiplier *= multiplier
		},
		OnBeforePeriodicDamage: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect, tickDamage *float64) {
			*tickDamage *= multiplier
		},
	}
}

var ShadowWeavingDebuffID = NewDebuffID()

func ShadowWeavingAura(sim *Simulation, numStacks int32) Aura {
	multiplier := 1.0 + 0.02*float64(numStacks)

	return Aura{
		ID:       ShadowWeavingDebuffID,
		ActionID: ActionID{SpellID: 15334},
		Expires:  sim.CurrentTime + time.Second*15,
		Stacks:   numStacks,
		OnBeforeSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect) {
			if spellCast.SpellSchool.Matches(SpellSchoolShadow) {
				spellEffect.DamageMultiplier *= multiplier
			}
		},
		OnBeforePeriodicDamage: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect, tickDamage *float64) {
			if spellCast.SpellSchool.Matches(SpellSchoolShadow) {
				*tickDamage *= multiplier
			}
		},
	}
}

var JudgementOfWisdomDebuffID = NewDebuffID()

func JudgementOfWisdomAura(sim *Simulation) Aura {
	const mana = 74 / 2 // 50% proc
	actionID := ActionID{SpellID: 27164}
	var aura Aura
	aura = Aura{
		ID:       JudgementOfWisdomDebuffID,
		ActionID: actionID,
		Expires:  sim.CurrentTime + time.Second*20,
		OnSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect) {
			if spellCast.IsPhantom {
				return // Phantom spells (Romulo's, Lightning Capacitor, etc) don't proc JoW.
			}

			character := spellCast.Character
			if character.HasManaBar() {
				character.AddMana(sim, mana, actionID, false)
			}
		},
		OnMeleeAttack: func(sim *Simulation, ability *ActiveMeleeAbility, hitEffect *AbilityHitEffect) {
			if ability.ActionID.SpellID == 35395 {
				aura.Expires = sim.CurrentTime + time.Second*20
				hitEffect.Target.ReplaceAura(sim, aura)
			}
			if ability.IsPhantom {
				return
			}
			character := ability.Character
			if character.HasManaBar() {
				character.AddMana(sim, mana, actionID, false)
			}
		},
	}
	return aura
}

var ImprovedSealOfTheCrusaderDebuffID = NewDebuffID()

func JudgementOfTheCrusaderAura(sim *Simulation, level float64) Aura {
	bonusSPCrit := level * SpellCritRatingPerCritChance
	bonusMCrit := level * MeleeCritRatingPerCritChance
	var aura Aura
	aura = Aura{
		ID:       ImprovedSealOfTheCrusaderDebuffID,
		ActionID: ActionID{SpellID: 27159},
		Expires:  sim.CurrentTime + time.Second*20,
		OnBeforeSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect) {
			if spellCast.SpellSchool.Matches(SpellSchoolHoly) {
				spellEffect.BonusSpellPower += 219
			}
			spellEffect.BonusSpellCritRating += bonusSPCrit
		},
		OnBeforeMeleeHit: func(sim *Simulation, ability *ActiveMeleeAbility, hitEffect *AbilityHitEffect) {
			if ability.ActionID.SpellID == 35395 {
				aura.Expires = sim.CurrentTime + time.Second*20
				hitEffect.Target.ReplaceAura(sim, aura)
			}
			hitEffect.BonusCritRating += bonusMCrit
		},
	}
	return aura
}

var CurseOfElementsDebuffID = NewDebuffID()

func CurseOfElementsAura(coe proto.TristateEffect) Aura {
	mult := 1.1
	level := int32(0)
	if coe == proto.TristateEffect_TristateEffectImproved {
		mult = 1.13
		level = 3
	}

	return Aura{
		ID:       CurseOfElementsDebuffID,
		ActionID: ActionID{SpellID: 27228},
		Stacks:   level, // Use stacks to store talent level for detection by other code.
		OnBeforeSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect) {
			if spellCast.SpellSchool.Matches(SpellSchoolArcane | SpellSchoolFire | SpellSchoolFrost | SpellSchoolShadow) {
				spellEffect.DamageMultiplier *= mult
			}
		},
		OnBeforePeriodicDamage: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect, tickDamage *float64) {
			if spellCast.SpellSchool.Matches(SpellSchoolArcane | SpellSchoolFire | SpellSchoolFrost | SpellSchoolShadow) {
				*tickDamage *= mult
			}
		},
	}
}

var ImprovedShadowBoltID = NewDebuffID()

func ImprovedShadowBoltAura(uptime float64) Aura {
	mult := (1 + uptime*0.2)
	return Aura{
		ID:       ImprovedShadowBoltID,
		ActionID: ActionID{SpellID: 17803},
		OnBeforeSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect) {
			if !spellCast.SpellSchool.Matches(SpellSchoolShadow) {
				return // does not apply to these schools
			}
			spellEffect.DamageMultiplier *= mult
		},
		OnBeforePeriodicDamage: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect, tickDamage *float64) {
			if !spellCast.SpellSchool.Matches(SpellSchoolShadow) {
				return // does not apply to these schools
			}
			*tickDamage *= mult
		},
	}
}

var BloodFrenzyDebuffID = NewDebuffID()

func BloodFrenzyAura() Aura {
	return Aura{
		ID:       BloodFrenzyDebuffID,
		ActionID: ActionID{SpellID: 29859},
		OnBeforeMeleeHit: func(sim *Simulation, ability *ActiveMeleeAbility, hitEffect *AbilityHitEffect) {
			if ability.SpellSchool != stats.AttackPower {
				return
			}
			hitEffect.DamageMultiplier *= 1.04
		},
	}
}

var ImprovedScorchDebuffID = NewDebuffID()

func ImprovedScorchAura(sim *Simulation, numStacks int32) Aura {
	multiplier := 1.0 + 0.03*float64(numStacks)

	return Aura{
		ID:       ImprovedScorchDebuffID,
		ActionID: ActionID{SpellID: 12873},
		Expires:  sim.CurrentTime + time.Second*30,
		Stacks:   numStacks,
		OnBeforeSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect) {
			if spellCast.SpellSchool.Matches(SpellSchoolFire) {
				spellEffect.DamageMultiplier *= multiplier
			}
		},
		OnBeforePeriodicDamage: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect, tickDamage *float64) {
			if spellCast.SpellSchool.Matches(SpellSchoolFire) {
				*tickDamage *= multiplier
			}
		},
	}
}

var WintersChillDebuffID = NewDebuffID()

func WintersChillAura(sim *Simulation, numStacks int32) Aura {
	bonusCrit := 2 * float64(numStacks) * SpellCritRatingPerCritChance

	return Aura{
		ID:       WintersChillDebuffID,
		ActionID: ActionID{SpellID: 28595},
		Expires:  sim.CurrentTime + time.Second*15,
		Stacks:   numStacks,
		OnBeforeSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect) {
			if spellCast.SpellSchool.Matches(SpellSchoolFrost) {
				spellEffect.BonusSpellCritRating += bonusCrit
			}
		},
	}
}

var FaerieFireDebuffID = NewDebuffID()

func FaerieFireAura(currentTime time.Duration, target *Target, improved bool) Aura {
	const hitBonus = 3 * MeleeHitRatingPerHitChance
	target.AddArmor(-610)
	aura := Aura{
		ID:       FaerieFireDebuffID,
		ActionID: ActionID{SpellID: 26993},
		Expires:  currentTime + time.Second*40,
		OnExpire: func(sim *Simulation) {
			target.AddArmor(610)
		},
	}
	if improved {
		aura.OnBeforeMeleeHit = func(sim *Simulation, ability *ActiveMeleeAbility, hitEffect *AbilityHitEffect) {
			hitEffect.BonusHitRating += hitBonus
		}
	}

	return aura
}

var SunderArmorDebuffID = NewDebuffID()

func SunderArmorAura(currentTime time.Duration, target *Target, stacks int) Aura {
	armorReduction := 520.0 * float64(stacks)
	target.AddArmor(-armorReduction)

	aura := Aura{
		ID:       SunderArmorDebuffID,
		ActionID: ActionID{SpellID: 25225},
		Expires:  currentTime + time.Second*30,
		OnExpire: func(sim *Simulation) {
			target.AddArmor(armorReduction)
		},
	}

	return aura
}

var ExposeArmorDebuffID = NewDebuffID()

func ExposeArmorAura(currentTime time.Duration, target *Target, talentPoints int) Aura {
	// TODO: Make this override sunder, not add
	armorReduction := 2050.0 * (1.0 + 0.25*float64(talentPoints))
	target.AddArmor(-armorReduction)

	aura := Aura{
		ID:       ExposeArmorDebuffID,
		ActionID: ActionID{SpellID: 26866},
		Expires:  currentTime + time.Second*30,
		OnExpire: func(sim *Simulation) {
			target.AddArmor(armorReduction)
		},
	}

	return aura
}

var CurseOfRecklessnessDebuffID = NewDebuffID()

func CurseOfRecklessnessAura(currentTime time.Duration, target *Target) Aura {
	armorReduction := 800.0
	target.AddArmor(-armorReduction)

	aura := Aura{
		ID:       CurseOfRecklessnessDebuffID,
		ActionID: ActionID{SpellID: 27226},
		Expires:  currentTime + time.Minute*2,
		OnExpire: func(sim *Simulation) {
			target.AddArmor(armorReduction)
		},
	}

	return aura
}

var ExposeWeaknessDebuffID = NewDebuffID()

// Multiplier is for accomodating uptime %. For a real hunter, always pass 1.0
func ExposeWeaknessAura(currentTime time.Duration, hunterAgility float64, multiplier float64) Aura {
	apBonus := hunterAgility * 0.25 * multiplier

	return Aura{
		ID:       ExposeWeaknessDebuffID,
		ActionID: ActionID{SpellID: 34503},
		Expires:  currentTime + time.Second*7,
		OnBeforeMeleeHit: func(sim *Simulation, ability *ActiveMeleeAbility, hitEffect *AbilityHitEffect) {
			hitEffect.BonusAttackPower += apBonus
		},
	}
}

var HuntersMarkDebuffID = NewDebuffID()

func HuntersMarkAura(points int32, fullyStacked bool) Aura {
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
		ID:       HuntersMarkDebuffID,
		ActionID: ActionID{SpellID: 14325},
		Stacks:   points, // Use this to check whether to override in hunter/hunter.go
		OnBeforeMeleeHit: func(sim *Simulation, ability *ActiveMeleeAbility, hitEffect *AbilityHitEffect) {
			if hitEffect.IsMelee() {
				hitEffect.BonusAttackPower += meleeBonus
			} else {
				hitEffect.BonusAttackPower += rangedBonus
			}
		},
		OnMeleeAttack: func(sim *Simulation, ability *ActiveMeleeAbility, hitEffect *AbilityHitEffect) {
			if !hitEffect.IsRanged() || !hitEffect.Landed() {
				return
			}
			if stacks < maxStacks {
				stacks++
				rangedBonus = baseRangedBonus + bonusPerStack*float64(stacks)
			}
		},
	}
}
