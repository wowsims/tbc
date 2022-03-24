package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var DwarfGunSpecializationAuraID = NewAuraID()
var HumanWeaponSpecializationAuraID = NewAuraID()

var OrcBloodFuryAuraID = NewAuraID()
var OrcBloodFuryCooldownID = NewCooldownID()
var OrcCommandAuraID = NewAuraID()
var OrcWeaponSpecializationAuraID = NewAuraID()

var TrollBowSpecializationAuraID = NewAuraID()
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
		// Gun specialization (+1% ranged crit when using a gun).
		matches := false
		if weapon := character.Equip[proto.ItemSlot_ItemSlotRanged]; weapon.ID != 0 {
			if weapon.RangedWeaponType == proto.RangedWeaponType_RangedWeaponTypeGun {
				matches = true
			}
		}

		if matches && character.Class == proto.Class_ClassHunter {
			character.AddPermanentAura(func(sim *Simulation) Aura {
				return Aura{
					ID: DwarfGunSpecializationAuraID,
					OnBeforeSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellHitEffect) {
						if spellCast.OutcomeRollCategory.Matches(OutcomeRollCategoryRanged) {
							spellEffect.BonusCritRating += 1 * MeleeCritRatingPerCritChance
						}
					},
				}
			})
		}
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
		procMask := GetMeleeProcMaskForHands(mhMatches, ohMatches)

		if procMask != ProcMaskEmpty {
			character.AddPermanentAura(func(sim *Simulation) Aura {
				return Aura{
					ID: HumanWeaponSpecializationAuraID,
					OnBeforeSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellHitEffect) {
						if !spellEffect.ProcMask.Matches(procMask) {
							return
						}
						spellEffect.BonusExpertiseRating += expertiseBonus
					},
				}
			})
		}
	case proto.Race_RaceNightElf:
	case proto.Race_RaceOrc:
		// Command (Pet damage +5%)
		if len(character.Pets) > 0 {
			const multiplier = 1.05
			for _, petAgent := range character.Pets {
				pet := petAgent.GetPet()
				pet.AddPermanentAura(func(sim *Simulation) Aura {
					return Aura{
						ID: OrcCommandAuraID,
						OnBeforeSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellHitEffect) {
							spellEffect.DamageMultiplier *= multiplier
						},
						OnBeforePeriodicDamage: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect, tickDamage *float64) {
							*tickDamage *= multiplier
						},
					}
				})
			}
		}

		// Blood Fury
		const cd = time.Minute * 2
		const dur = time.Second * 15
		const apBonus = float64(CharacterLevel)*4 + 2
		const spBonus = float64(CharacterLevel)*2 + 3
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
				applyStatAura := character.NewTemporaryStatsAuraApplier(OrcBloodFuryAuraID, actionID, stats.Stats{stats.AttackPower: apBonus, stats.RangedAttackPower: apBonus, stats.SpellPower: spBonus}, dur)
				return func(sim *Simulation, character *Character) {
					applyStatAura(sim)
					character.SetCD(OrcBloodFuryCooldownID, sim.CurrentTime+cd)
					character.Metrics.AddInstantCast(actionID)
				}
			},
		})

		// Axe specialization
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
		procMask := GetMeleeProcMaskForHands(mhMatches, ohMatches)

		if procMask != ProcMaskEmpty {
			character.AddPermanentAura(func(sim *Simulation) Aura {
				return Aura{
					ID: OrcWeaponSpecializationAuraID,
					OnBeforeSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellHitEffect) {
						if !spellEffect.ProcMask.Matches(procMask) {
							return
						}
						spellEffect.BonusExpertiseRating += expertiseBonus
					},
				}
			})
		}
	case proto.Race_RaceTauren:
		// TODO: Health +5%
	case proto.Race_RaceTroll10, proto.Race_RaceTroll30:
		// Bow specialization (+1% ranged crit when using a bow).
		matches := false
		if weapon := character.Equip[proto.ItemSlot_ItemSlotRanged]; weapon.ID != 0 {
			if weapon.RangedWeaponType == proto.RangedWeaponType_RangedWeaponTypeBow {
				matches = true
			}
		}

		if matches && character.Class == proto.Class_ClassHunter {
			character.AddPermanentAura(func(sim *Simulation) Aura {
				return Aura{
					ID: TrollBowSpecializationAuraID,
					OnBeforeSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellHitEffect) {
						if spellCast.OutcomeRollCategory.Matches(OutcomeRollCategoryRanged) {
							spellEffect.BonusCritRating += 1 * MeleeCritRatingPerCritChance
						}
					},
				}
			})
		}

		// Beast Slaying (+5% damage to beasts)
		character.AddPermanentAura(func(sim *Simulation) Aura {
			return Aura{
				ID: TrollBeastSlayingAuraID,
				OnBeforeSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellHitEffect) {
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

		var cost ResourceCost
		var actionID ActionID
		if character.Class == proto.Class_ClassRogue {
			actionID = ActionID{SpellID: 26297, CooldownID: TrollBerserkingCooldownID}
		} else if character.Class == proto.Class_ClassWarrior {
			actionID = ActionID{SpellID: 26296, CooldownID: TrollBerserkingCooldownID}
		} else {
			actionID = ActionID{SpellID: 20554, CooldownID: TrollBerserkingCooldownID}
		}

		character.AddMajorCooldown(MajorCooldown{
			ActionID:   actionID,
			CooldownID: TrollBerserkingCooldownID,
			Cooldown:   cd,
			Type:       CooldownTypeDPS,
			CanActivate: func(sim *Simulation, character *Character) bool {
				if character.Class == proto.Class_ClassRogue {
					return character.CurrentEnergy() >= cost.Value
				} else if character.Class == proto.Class_ClassWarrior {
					return character.CurrentRage() >= cost.Value
				} else {
					return character.CurrentMana() >= cost.Value
				}
			},
			ShouldActivate: func(sim *Simulation, character *Character) bool {
				return true
			},
			ActivationFactory: func(sim *Simulation) CooldownActivation {
				if character.Class == proto.Class_ClassRogue {
					cost = ResourceCost{Type: stats.Energy, Value: 10}
				} else if character.Class == proto.Class_ClassWarrior {
					cost = ResourceCost{Type: stats.Rage, Value: 5}
				} else {
					cost = ResourceCost{Type: stats.Mana, Value: character.BaseMana() * 0.06}
				}

				castTemplate := SimpleCast{
					Cast: Cast{
						ActionID:  actionID,
						Character: character,
						BaseCost:  cost,
						Cost:      cost,
						Cooldown:  cd,
						OnCastComplete: func(sim *Simulation, cast *Cast) {
							character.AddAura(sim, Aura{
								ID:       TrollBerserkingAuraID,
								ActionID: actionID,
								Duration: dur,
								OnGain: func(sim *Simulation) {
									character.PseudoStats.CastSpeedMultiplier *= hasteBonus
									character.MultiplyAttackSpeed(sim, hasteBonus)
								},
								OnExpire: func(sim *Simulation) {
									character.PseudoStats.CastSpeedMultiplier /= hasteBonus
									character.MultiplyAttackSpeed(sim, inverseBonus)
								},
							})
						},
					},
				}

				return func(sim *Simulation, character *Character) {
					cast := castTemplate
					cast.Init(sim)
					cast.StartCast(sim)
				}
			},
		})
	case proto.Race_RaceUndead:
	}
}
