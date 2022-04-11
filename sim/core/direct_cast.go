package core

// Function for calculating the base damage of a spell.
type BaseDamageCalculator func(*Simulation, *SpellEffect, *Spell) float64

type BaseDamageConfig struct {
	// Lambda for calculating the base damage.
	Calculator BaseDamageCalculator

	// Spell coefficient for +damage effects on the target.
	TargetSpellCoefficient float64
}

func BuildBaseDamageConfig(calculator BaseDamageCalculator, coeff float64) BaseDamageConfig {
	return BaseDamageConfig{
		Calculator:             calculator,
		TargetSpellCoefficient: coeff,
	}
}

func WrapBaseDamageConfig(config BaseDamageConfig, wrapper func(oldCalculator BaseDamageCalculator) BaseDamageCalculator) BaseDamageConfig {
	return BaseDamageConfig{
		Calculator:             wrapper(config.Calculator),
		TargetSpellCoefficient: config.TargetSpellCoefficient,
	}
}

// Creates a BaseDamageCalculator function which returns a flat value.
func BaseDamageFuncFlat(damage float64) BaseDamageCalculator {
	return func(_ *Simulation, _ *SpellEffect, _ *Spell) float64 {
		return damage
	}
}
func BaseDamageConfigFlat(damage float64) BaseDamageConfig {
	return BuildBaseDamageConfig(BaseDamageFuncFlat(damage), 0)
}

// Creates a BaseDamageCalculator function with a single damage roll.
func BaseDamageFuncRoll(minFlatDamage float64, maxFlatDamage float64) BaseDamageCalculator {
	if minFlatDamage == maxFlatDamage {
		return BaseDamageFuncFlat(minFlatDamage)
	} else {
		deltaDamage := maxFlatDamage - minFlatDamage
		return func(sim *Simulation, _ *SpellEffect, _ *Spell) float64 {
			return damageRollOptimized(sim, minFlatDamage, deltaDamage)
		}
	}
}
func BaseDamageConfigRoll(minFlatDamage float64, maxFlatDamage float64) BaseDamageConfig {
	return BuildBaseDamageConfig(BaseDamageFuncRoll(minFlatDamage, maxFlatDamage), 0)
}

func BaseDamageFuncMagic(minFlatDamage float64, maxFlatDamage float64, spellCoefficient float64) BaseDamageCalculator {
	if spellCoefficient == 0 {
		return BaseDamageFuncRoll(minFlatDamage, maxFlatDamage)
	}

	if minFlatDamage == 0 && maxFlatDamage == 0 {
		return func(_ *Simulation, hitEffect *SpellEffect, spell *Spell) float64 {
			return hitEffect.SpellPower(spell.Character, spell) * spellCoefficient
		}
	} else if minFlatDamage == maxFlatDamage {
		return func(sim *Simulation, hitEffect *SpellEffect, spell *Spell) float64 {
			damage := hitEffect.SpellPower(spell.Character, spell) * spellCoefficient
			return damage + minFlatDamage
		}
	} else {
		deltaDamage := maxFlatDamage - minFlatDamage
		return func(sim *Simulation, hitEffect *SpellEffect, spell *Spell) float64 {
			damage := hitEffect.SpellPower(spell.Character, spell) * spellCoefficient
			damage += damageRollOptimized(sim, minFlatDamage, deltaDamage)
			return damage
		}
	}
}
func BaseDamageConfigMagic(minFlatDamage float64, maxFlatDamage float64, spellCoefficient float64) BaseDamageConfig {
	return BuildBaseDamageConfig(BaseDamageFuncMagic(minFlatDamage, maxFlatDamage, spellCoefficient), spellCoefficient)
}
func BaseDamageConfigMagicNoRoll(flatDamage float64, spellCoefficient float64) BaseDamageConfig {
	return BaseDamageConfigMagic(flatDamage, flatDamage, spellCoefficient)
}

func MultiplyByStacks(config BaseDamageConfig, aura *Aura) BaseDamageConfig {
	return WrapBaseDamageConfig(config, func(oldCalculator BaseDamageCalculator) BaseDamageCalculator {
		return func(sim *Simulation, hitEffect *SpellEffect, spell *Spell) float64 {
			return oldCalculator(sim, hitEffect, spell) * float64(aura.GetStacks())
		}
	})
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
			return func(sim *Simulation, hitEffect *SpellEffect, spell *Spell) float64 {
				damage := spell.Character.AutoAttacks.MH.calculateNormalizedWeaponDamage(
					sim, hitEffect.MeleeAttackPower(spell.Character)+hitEffect.MeleeAttackPowerOnTarget())
				damage += flatBonus
				if includeBonusWeaponDamage {
					damage += hitEffect.BonusWeaponDamage(spell.Character)
				}
				return damage * weaponMultiplier
			}
		} else {
			return func(sim *Simulation, hitEffect *SpellEffect, spell *Spell) float64 {
				damage := spell.Character.AutoAttacks.OH.calculateNormalizedWeaponDamage(
					sim, hitEffect.MeleeAttackPower(spell.Character)+2*hitEffect.MeleeAttackPowerOnTarget())
				damage = damage*0.5 + flatBonus
				if includeBonusWeaponDamage {
					damage += hitEffect.BonusWeaponDamage(spell.Character)
				}
				return damage * weaponMultiplier
			}
		}
	} else {
		if hand == MainHand {
			return func(sim *Simulation, hitEffect *SpellEffect, spell *Spell) float64 {
				damage := spell.Character.AutoAttacks.MH.calculateWeaponDamage(
					sim, hitEffect.MeleeAttackPower(spell.Character)+hitEffect.MeleeAttackPowerOnTarget())
				damage += flatBonus
				if includeBonusWeaponDamage {
					damage += hitEffect.BonusWeaponDamage(spell.Character)
				}
				return damage * weaponMultiplier
			}
		} else {
			return func(sim *Simulation, hitEffect *SpellEffect, spell *Spell) float64 {
				damage := spell.Character.AutoAttacks.OH.calculateWeaponDamage(
					sim, hitEffect.MeleeAttackPower(spell.Character)+2*hitEffect.MeleeAttackPowerOnTarget())
				damage = damage*0.5 + flatBonus
				if includeBonusWeaponDamage {
					damage += hitEffect.BonusWeaponDamage(spell.Character)
				}
				return damage * weaponMultiplier
			}
		}
	}
}
func BaseDamageConfigMeleeWeapon(hand Hand, normalized bool, flatBonus float64, weaponMultiplier float64, includeBonusWeaponDamage bool) BaseDamageConfig {
	calculator := BaseDamageFuncMeleeWeapon(hand, normalized, flatBonus, weaponMultiplier, includeBonusWeaponDamage)
	if includeBonusWeaponDamage {
		return BuildBaseDamageConfig(calculator, 1)
	} else {
		return BuildBaseDamageConfig(calculator, 0)
	}
}

func BaseDamageFuncRangedWeapon(flatBonus float64) BaseDamageCalculator {
	return func(sim *Simulation, hitEffect *SpellEffect, spell *Spell) float64 {
		return spell.Character.AutoAttacks.Ranged.calculateWeaponDamage(sim, hitEffect.RangedAttackPower(spell.Character)+hitEffect.RangedAttackPowerOnTarget()) +
			flatBonus +
			hitEffect.BonusWeaponDamage(spell.Character)
	}
}
func BaseDamageConfigRangedWeapon(flatBonus float64) BaseDamageConfig {
	return BuildBaseDamageConfig(BaseDamageFuncRangedWeapon(flatBonus), 1)
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
