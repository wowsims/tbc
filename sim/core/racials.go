package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var HumanWeaponSpecializationAuraID = NewAuraID()

var OrcBloodFuryAuraID = NewAuraID()
var OrcBloodFuryCooldownID = NewCooldownID()
var OrcWeaponSpecializationAuraID = NewAuraID()

var TrollBeastSlayingAuraID = NewAuraID()

var TrollBerserkingAuraID = NewAuraID()
var TrollBerserkingCooldownID = NewCooldownID()

func applyRaceEffects(agent Agent) {
	character := agent.GetCharacter()

	switch character.Race {
	case proto.Race_RaceBloodElf:
		// TODO: Add major cooldown: arcane torrent
	case proto.Race_RaceDraenei:
	case proto.Race_RaceDwarf:
		// TODO: If gun equipped, +1% ranged crit
	case proto.Race_RaceGnome:
		character.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Intellect,
			ModifiedStat: stats.Intellect,
			Modifier: func(intellect float64, _ float64) float64 {
				return intellect * 1.05
			},
		})
	case proto.Race_RaceHuman:
		character.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Spirit,
			ModifiedStat: stats.Spirit,
			Modifier: func(spirit float64, _ float64) float64 {
				return spirit * 1.1
			},
		})

		const expertiseBonus = 5 * ExpertisePerQuarterPercentReduction
		mhMatches := false
		ohMatches := false
		if weapon := character.Equip[proto.ItemSlot_ItemSlotMainHand]; weapon.ID != 0 {
			if weapon.WeaponType == proto.WeaponType_WeaponTypeSword || weapon.WeaponType == proto.WeaponType_WeaponTypeMace {
				mhMatches = true
			}
		}
		if weapon := character.Equip[proto.ItemSlot_ItemSlotOffHand]; weapon.ID != 0 {
			if weapon.WeaponType == proto.WeaponType_WeaponTypeSword || weapon.WeaponType == proto.WeaponType_WeaponTypeMace {
				ohMatches = true
			}
		}

		if mhMatches || ohMatches {
			character.AddPermanentAura(func(sim *Simulation) Aura {
				return Aura{
					ID: HumanWeaponSpecializationAuraID,
					OnBeforeMeleeHit: func(sim *Simulation, ability *ActiveMeleeAbility, hitEffect *AbilityHitEffect) {
						if hitEffect.IsMH() {
							if !mhMatches {
								return
							}
						} else if !ohMatches {
							return
						}
						hitEffect.BonusExpertiseRating += expertiseBonus
					},
				}
			})
		}
	case proto.Race_RaceNightElf:
	case proto.Race_RaceOrc:
		// TODO: Pet melee damage +5%
		// TODO: +5 expertise with axes
		const cd = time.Minute * 2
		const dur = time.Second * 15
		const spBonus = 143
		actionID := ActionID{SpellID: 33697}

		character.AddMajorCooldown(MajorCooldown{
			ActionID:   actionID,
			CooldownID: OrcBloodFuryCooldownID,
			Cooldown:   cd,
			Type:       CooldownTypeDPS,
			CanActivate: func(sim *Simulation, character *Character) bool {
				return true
			},
			ShouldActivate: func(sim *Simulation, character *Character) bool {
				return true
			},
			ActivationFactory: func(sim *Simulation) CooldownActivation {
				return func(sim *Simulation, character *Character) {
					character.SetCD(OrcBloodFuryCooldownID, cd+sim.CurrentTime)
					character.AddAuraWithTemporaryStats(sim, OrcBloodFuryAuraID, actionID, stats.SpellPower, spBonus, dur)
					character.Metrics.AddInstantCast(actionID)
				}
			},
		})

		const expertiseBonus = 5 * ExpertisePerQuarterPercentReduction
		mhMatches := false
		ohMatches := false
		if weapon := character.Equip[proto.ItemSlot_ItemSlotMainHand]; weapon.ID != 0 {
			if weapon.WeaponType == proto.WeaponType_WeaponTypeAxe {
				mhMatches = true
			}
		}
		if weapon := character.Equip[proto.ItemSlot_ItemSlotOffHand]; weapon.ID != 0 {
			if weapon.WeaponType == proto.WeaponType_WeaponTypeAxe {
				ohMatches = true
			}
		}

		if mhMatches || ohMatches {
			character.AddPermanentAura(func(sim *Simulation) Aura {
				return Aura{
					ID: OrcWeaponSpecializationAuraID,
					OnBeforeMeleeHit: func(sim *Simulation, ability *ActiveMeleeAbility, hitEffect *AbilityHitEffect) {
						if hitEffect.IsMH() {
							if !mhMatches {
								return
							}
						} else if !ohMatches {
							return
						}
						hitEffect.BonusExpertiseRating += expertiseBonus
					},
				}
			})
		}
	case proto.Race_RaceTauren:
		// TODO: Health +5%
	case proto.Race_RaceTroll10, proto.Race_RaceTroll30:
		// TODO: +1% ranged crit when using a bow

		// Beast Slaying (+5% damage to beasts)
		character.AddPermanentAura(func(sim *Simulation) Aura {
			return Aura{
				ID: TrollBeastSlayingAuraID,
				OnBeforeMeleeHit: func(sim *Simulation, ability *ActiveMeleeAbility, hitEffect *AbilityHitEffect) {
					if hitEffect.Target.MobType == proto.MobType_MobTypeBeast {
						hitEffect.DamageMultiplier *= 1.05
					}
				},
				OnBeforeSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect) {
					if spellEffect.Target.MobType == proto.MobType_MobTypeBeast {
						spellEffect.DamageMultiplier *= 1.05
					}
				},
				OnBeforePeriodicDamage: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect, tickDamage *float64) {
					if spellEffect.Target.MobType == proto.MobType_MobTypeBeast {
						*tickDamage *= 1.05
					}
				},
			}
		})

		// Berserking
		hasteBonus := 1.1
		if character.Race == proto.Race_RaceTroll30 {
			hasteBonus = 1.3
		}
		inverseBonus := 1 / hasteBonus
		const dur = time.Second * 10
		const cd = time.Minute * 3
		manaCost := 0.0
		actionID := ActionID{SpellID: 20554}

		character.AddMajorCooldown(MajorCooldown{
			ActionID:   actionID,
			CooldownID: TrollBerserkingCooldownID,
			Cooldown:   cd,
			Type:       CooldownTypeDPS,
			CanActivate: func(sim *Simulation, character *Character) bool {
				if character.CurrentMana() < manaCost {
					return false
				}
				return true
			},
			ShouldActivate: func(sim *Simulation, character *Character) bool {
				return true
			},
			ActivationFactory: func(sim *Simulation) CooldownActivation {
				manaCost = character.BaseMana() * 0.06
				return func(sim *Simulation, character *Character) {
					character.SetCD(TrollBerserkingCooldownID, cd+sim.CurrentTime)
					// Increase cast speed multiplier
					character.PseudoStats.CastSpeedMultiplier *= hasteBonus
					character.MultiplyMeleeSpeed(sim, hasteBonus)
					character.AddAura(sim, Aura{
						ID:       TrollBerserkingAuraID,
						ActionID: actionID,
						Expires:  sim.CurrentTime + dur,
						OnExpire: func(sim *Simulation) {
							character.PseudoStats.CastSpeedMultiplier /= hasteBonus
							character.MultiplyMeleeSpeed(sim, inverseBonus)
						},
					})
					character.SpendMana(sim, manaCost, actionID)
					character.Metrics.AddInstantCast(actionID)
				}
			},
		})
	case proto.Race_RaceUndead:
	}
}
