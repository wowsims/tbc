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
			return JudgementOfWisdomAura()
		})
	}

	if debuffs.ImprovedSealOfTheCrusader {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return ImprovedSealOfTheCrusaderAura()
		})
	}

	if debuffs.CurseOfElements != proto.TristateEffect_TristateEffectMissing {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return CurseOfElementsAura(debuffs.CurseOfElements)
		})
	}

	if debuffs.IsbUptime > 0.0 {
		target.AddPermanentAura(func(sim *Simulation) Aura {
			return ImprovedShadowBoltAura(debuffs.IsbUptime)
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
		target.armor -= 2050 // 5 points: 2050 armor / imp 5 points: 3075 armor
		if debuffs.ExposeArmor == proto.TristateEffect_TristateEffectImproved {
			target.armor -= 1025
		}
	}

	if debuffs.FaerieFire != proto.TristateEffect_TristateEffectMissing {
		target.armor -= 610
		if debuffs.FaerieFire == proto.TristateEffect_TristateEffectImproved {
			target.AddPermanentAura(func(sim *Simulation) Aura {
				return FaerieFireAura(0, target)
			})
		}
	}

	if debuffs.SunderArmor {
		target.armor -= 5 * 520.0 // assume 5 stacks
	}

	if debuffs.CurseOfRecklessness {
		target.armor -= 800
	}
}

var MiseryDebuffID = NewDebuffID()

func MiseryAura(sim *Simulation, numPoints int32) Aura {
	multiplier := 1.0 + 0.01*float64(numPoints)

	return Aura{
		ID:      MiseryDebuffID,
		Name:    "Misery",
		SpellID: 33195,
		Expires: sim.CurrentTime + time.Second*24,
		Stacks:  numPoints,
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
		ID:      ShadowWeavingDebuffID,
		Name:    "Shadow Weaving",
		SpellID: 15334,
		Expires: sim.CurrentTime + time.Second*15,
		Stacks:  numStacks,
		OnBeforeSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect) {
			if spellCast.SpellSchool == stats.ShadowSpellPower {
				spellEffect.DamageMultiplier *= multiplier
			}
		},
		OnBeforePeriodicDamage: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect, tickDamage *float64) {
			if spellCast.SpellSchool == stats.ShadowSpellPower {
				*tickDamage *= multiplier
			}
		},
	}
}

var JudgementOfWisdomDebuffID = NewDebuffID()

func JudgementOfWisdomAura() Aura {
	const mana = 74 / 2 // 50% proc
	return Aura{
		ID:      JudgementOfWisdomDebuffID,
		Name:    "Judgement of Wisdom",
		SpellID: 27164,
		OnSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect) {
			if spellCast.ActionID.ItemID == ItemIDTheLightningCapacitor {
				return // TLC cant proc JoW
			}

			character := spellCast.Character
			// Only apply to agents that have mana.
			if character.MaxMana() > 0 {
				character.AddMana(sim, mana, "Judgement of Wisdom", false)
			}
		},
	}
}

var ImprovedSealOfTheCrusaderDebuffID = NewDebuffID()

func ImprovedSealOfTheCrusaderAura() Aura {
	return Aura{
		ID:      ImprovedSealOfTheCrusaderDebuffID,
		Name:    "Improved Seal of the Crusader",
		SpellID: 20337,
		OnBeforeSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect) {
			spellEffect.BonusSpellCritRating += 3 * SpellCritRatingPerCritChance
			// FUTURE: melee crit bonus, research actual value
		},
	}
}

var CurseOfElementsDebuffID = NewDebuffID()

func CurseOfElementsAura(coe proto.TristateEffect) Aura {
	mult := 1.1
	if coe == proto.TristateEffect_TristateEffectImproved {
		mult = 1.13
	}
	return Aura{
		ID:      CurseOfElementsDebuffID,
		Name:    "Curse of the Elements",
		SpellID: 27228,
		OnBeforeSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect) {
			if spellCast.SpellSchool == stats.NatureSpellPower ||
				spellCast.SpellSchool == stats.HolySpellPower ||
				spellCast.SpellSchool == stats.AttackPower {
				return // does not apply to these schools
			}
			spellEffect.DamageMultiplier *= mult
		},
		OnBeforePeriodicDamage: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect, tickDamage *float64) {
			if spellCast.SpellSchool == stats.NatureSpellPower ||
				spellCast.SpellSchool == stats.HolySpellPower ||
				spellCast.SpellSchool == stats.AttackPower {
				return // does not apply to these schools
			}
			*tickDamage *= mult
		},
	}
}

var ImprovedShadowBoltID = NewDebuffID()

func ImprovedShadowBoltAura(uptime float64) Aura {
	mult := (1 + uptime*0.2)
	return Aura{
		ID:   ImprovedShadowBoltID,
		Name: "Improved Shadow Bolt",
		OnBeforeSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect) {
			if spellCast.SpellSchool != stats.ShadowSpellPower {
				return // does not apply to these schools
			}
			spellEffect.DamageMultiplier *= mult
		},
		OnBeforePeriodicDamage: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect, tickDamage *float64) {
			if spellCast.SpellSchool != stats.ShadowSpellPower {
				return // does not apply to these schools
			}
			*tickDamage *= mult
		},
	}
}

var BloodFrenzyDebuffID = NewDebuffID()

func BloodFrenzyAura() Aura {
	return Aura{
		ID:   BloodFrenzyDebuffID,
		Name: "Blood Frenzy",
		OnBeforeMelee: func(sim *Simulation, ability *ActiveMeleeAbility, isOH bool) {
			ability.DamageMultiplier *= 1.04
		},
	}
}

var ImprovedScorchDebuffID = NewDebuffID()

func ImprovedScorchAura(sim *Simulation, numStacks int32) Aura {
	multiplier := 1.0 + 0.03*float64(numStacks)

	return Aura{
		ID:      ImprovedScorchDebuffID,
		Name:    "Improved Scorch",
		SpellID: 12873,
		Expires: sim.CurrentTime + time.Second*30,
		Stacks:  numStacks,
		OnBeforeSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect) {
			if spellCast.SpellSchool == stats.FireSpellPower {
				spellEffect.DamageMultiplier *= multiplier
			}
		},
		OnBeforePeriodicDamage: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect, tickDamage *float64) {
			if spellCast.SpellSchool == stats.FireSpellPower {
				*tickDamage *= multiplier
			}
		},
	}
}

var WintersChillDebuffID = NewDebuffID()

func WintersChillAura(sim *Simulation, numStacks int32) Aura {
	bonusCrit := 2 * float64(numStacks) * SpellCritRatingPerCritChance

	return Aura{
		ID:      WintersChillDebuffID,
		Name:    "Winter's Chill",
		SpellID: 28595,
		Expires: sim.CurrentTime + time.Second*15,
		Stacks:  numStacks,
		OnBeforeSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect) {
			if spellCast.SpellSchool == stats.FrostSpellPower {
				spellEffect.BonusSpellCritRating += bonusCrit
			}
		},
	}
}

var FaerieFireDebuffID = NewDebuffID()

func FaerieFireAura(currentTime time.Duration, target *Target) Aura {
	const hitBonus = 3 * MeleeHitRatingPerHitChance
	target.AddArmor(-610)
	return Aura{
		ID:      FaerieFireDebuffID,
		SpellID: 26993,
		Name:    "Faerie Fire",
		Expires: currentTime + time.Second*40,
		OnExpire: func(sim *Simulation) {
			target.AddArmor(-610)
		},
		OnBeforeMelee: func(sim *Simulation, ability *ActiveMeleeAbility, isOH bool) {
			ability.BonusHitRating += hitBonus
		},
	}

}

var SunderArmorDebuffID = NewDebuffID()
var CurseOfRecklessnessDebuffID = NewDebuffID()
