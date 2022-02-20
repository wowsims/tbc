package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Registers all consume-related effects to the Agent.
func applyConsumeEffects(agent Agent, raidBuffs proto.RaidBuffs, partyBuffs proto.PartyBuffs) {
	character := agent.GetCharacter()
	consumes := character.Consumes

	if consumes.Flask != proto.Flask_FlaskUnknown {
		switch consumes.Flask {
		case proto.Flask_FlaskOfBlindingLight:
			character.AddStats(stats.Stats{
				stats.NatureSpellPower: 80,
				stats.ArcaneSpellPower: 80,
				stats.HolySpellPower:   80,
			})
		case proto.Flask_FlaskOfMightyRestoration:
			character.AddStats(stats.Stats{
				stats.MP5: 25,
			})
		case proto.Flask_FlaskOfPureDeath:
			character.AddStats(stats.Stats{
				stats.FireSpellPower:   80,
				stats.FrostSpellPower:  80,
				stats.ShadowSpellPower: 80,
			})
		case proto.Flask_FlaskOfRelentlessAssault:
			character.AddStats(stats.Stats{
				stats.AttackPower:       120,
				stats.RangedAttackPower: 120,
			})
		case proto.Flask_FlaskOfSupremePower:
			character.AddStats(stats.Stats{
				stats.SpellPower: 70,
			})
		}
	} else {
		switch consumes.BattleElixir {
		case proto.BattleElixir_AdeptsElixir:
			character.AddStats(stats.Stats{
				stats.SpellCrit:    24,
				stats.SpellPower:   24,
				stats.HealingPower: 24,
			})
		case proto.BattleElixir_ElixirOfDemonslaying:
			character.AddPermanentAura(func(sim *Simulation) Aura {
				return ElixirOfDemonslayingAura()
			})
		case proto.BattleElixir_ElixirOfMajorAgility:
			character.AddStats(stats.Stats{
				stats.Agility:   35,
				stats.MeleeCrit: 20,
			})
		case proto.BattleElixir_ElixirOfMajorFirePower:
			character.AddStats(stats.Stats{
				stats.FireSpellPower: 55,
			})
		case proto.BattleElixir_ElixirOfMajorFrostPower:
			character.AddStats(stats.Stats{
				stats.FrostSpellPower: 55,
			})
		case proto.BattleElixir_ElixirOfMajorShadowPower:
			character.AddStats(stats.Stats{
				stats.ShadowSpellPower: 55,
			})
		case proto.BattleElixir_ElixirOfMajorStrength:
			character.AddStats(stats.Stats{
				stats.Strength: 35,
			})
		case proto.BattleElixir_ElixirOfTheMongoose:
			character.AddStats(stats.Stats{
				stats.Agility:   25,
				stats.MeleeCrit: 28,
			})
		}

		switch consumes.GuardianElixir {
		case proto.GuardianElixir_ElixirOfDraenicWisdom:
			character.AddStats(stats.Stats{
				stats.Intellect: 30,
				stats.Spirit:    30,
			})
		case proto.GuardianElixir_ElixirOfMajorMageblood:
			character.AddStats(stats.Stats{
				stats.MP5: 16.0,
			})
		}
	}

	switch consumes.Food {
	case proto.Food_FoodBlackenedBasilisk:
		character.AddStats(stats.Stats{
			stats.SpellPower:   23,
			stats.HealingPower: 23,
			stats.Spirit:       20,
		})
	case proto.Food_FoodGrilledMudfish:
		character.AddStats(stats.Stats{
			stats.Agility: 20,
			stats.Spirit:  20,
		})
	case proto.Food_FoodRavagerDog:
		character.AddStats(stats.Stats{
			stats.AttackPower:       40,
			stats.RangedAttackPower: 40,
			stats.Spirit:            20,
		})
	case proto.Food_FoodRoastedClefthoof:
		character.AddStats(stats.Stats{
			stats.Strength: 20,
			stats.Spirit:   20,
		})
	case proto.Food_FoodSkullfishSoup:
		character.AddStats(stats.Stats{
			stats.SpellCrit: 20,
			stats.Spirit:    20,
		})
	case proto.Food_FoodSpicyHotTalbuk:
		character.AddStats(stats.Stats{
			stats.MeleeHit: 20,
			stats.Spirit:   20,
		})
	}

	switch consumes.Alchohol {
	case proto.Alchohol_AlchoholKreegsStoutBeatdown:
		character.AddStats(stats.Stats{
			stats.Intellect: -5,
			stats.Spirit:    25,
		})
	}

	// Scrolls
	character.AddStat(stats.Agility, []float64{0, 5, 9, 13, 17, 20}[consumes.ScrollOfAgility])
	character.AddStat(stats.Strength, []float64{0, 5, 9, 13, 17, 20}[consumes.ScrollOfStrength])
	// Doesn't stack with DS
	if raidBuffs.DivineSpirit == proto.TristateEffect_TristateEffectMissing {
		character.AddStat(stats.Spirit, []float64{0, 3, 7, 11, 15, 30}[consumes.ScrollOfSpirit])
	}

	// Weapon Imbues
	if character.HasMHWeapon() && (character.HasMHWeaponImbue || partyBuffs.WindfuryTotemRank == 0) {
		addImbueStats(character, consumes.MainHandImbue)
	}
	if character.HasOHWeapon() {
		addImbueStats(character, consumes.OffHandImbue)
	}
	applyAdamantiteSharpeningStoneAura(character, consumes)

	registerDrumsCD(agent, partyBuffs, consumes)
	registerPotionCD(agent, consumes)
	registerConjuredCD(agent, consumes)
}

func ApplyPetConsumeEffects(pet *Character, ownerConsumes proto.Consumes) {
	switch ownerConsumes.PetFood {
	case proto.PetFood_PetFoodKiblersBits:
		pet.AddStats(stats.Stats{
			stats.Strength: 20,
			stats.Spirit:   20,
		})
	}

	pet.AddStat(stats.Agility, []float64{0, 5, 9, 13, 17, 20}[ownerConsumes.PetScrollOfAgility])
	pet.AddStat(stats.Strength, []float64{0, 5, 9, 13, 17, 20}[ownerConsumes.PetScrollOfStrength])
}

func addImbueStats(character *Character, imbue proto.WeaponImbue) {
	if imbue == proto.WeaponImbue_WeaponImbueUnknown {
		return
	}

	if imbue == proto.WeaponImbue_WeaponImbueAdamantiteSharpeningStone {
		character.PseudoStats.BonusMeleeDamage += 12
		character.PseudoStats.BonusRangedDamage += 12
		// Melee crit component handled separately because its melee-only.
	} else if imbue == proto.WeaponImbue_WeaponImbueAdamantiteWeightstone {
		character.PseudoStats.BonusMeleeDamage += 12
		character.PseudoStats.BonusRangedDamage += 12
		character.AddStats(stats.Stats{
			stats.MeleeCrit: 14,
		})
	} else if imbue == proto.WeaponImbue_WeaponImbueElementalSharpeningStone {
		character.AddStats(stats.Stats{
			stats.MeleeCrit: 28,
		})
	} else if imbue == proto.WeaponImbue_WeaponImbueBrilliantWizardOil {
		character.AddStats(stats.Stats{
			stats.SpellCrit:    14,
			stats.SpellPower:   36,
			stats.HealingPower: 36,
		})
	} else if imbue == proto.WeaponImbue_WeaponImbueSuperiorWizardOil {
		character.AddStats(stats.Stats{
			stats.SpellPower:   42,
			stats.HealingPower: 42,
		})
	}
}

var AdamantiteSharpeningStoneMeleeCritAuraID = NewAuraID()

func applyAdamantiteSharpeningStoneAura(character *Character, consumes proto.Consumes) {
	critBonus := 0.0
	if consumes.MainHandImbue == proto.WeaponImbue_WeaponImbueAdamantiteSharpeningStone {
		critBonus += 14
	}
	if consumes.OffHandImbue == proto.WeaponImbue_WeaponImbueAdamantiteSharpeningStone {
		critBonus += 14
	}

	// Crit rating for sharpening stone applies to melee only.
	if character.Class != proto.Class_ClassHunter {
		character.AddStats(stats.Stats{
			stats.MeleeCrit: critBonus,
		})
		return
	}

	character.AddPermanentAura(func(sim *Simulation) Aura {
		return Aura{
			ID: AdamantiteSharpeningStoneMeleeCritAuraID,
			OnBeforeMeleeHit: func(sim *Simulation, ability *ActiveMeleeAbility, hitEffect *AbilityHitEffect) {
				if !hitEffect.IsRanged() {
					hitEffect.BonusCritRating += critBonus
				}
			},
		}
	})
}

var ElixirOfDemonslayingAuraID = NewAuraID()

func ElixirOfDemonslayingAura() Aura {
	return Aura{
		ID:       ElixirOfDemonslayingAuraID,
		ActionID: ActionID{ItemID: 9224},
		OnBeforeMeleeHit: func(sim *Simulation, ability *ActiveMeleeAbility, hitEffect *AbilityHitEffect) {
			if hitEffect.Target.MobType == proto.MobType_MobTypeDemon {
				hitEffect.BonusAttackPower += 265
			}
		},
	}
}

var DrumsAuraID = NewAuraID()
var DrumsCooldownID = NewCooldownID()

const DrumsCD = time.Minute * 2 // Tinnitus
var DrumsOfBattleActionID = ActionID{SpellID: 35476}
var DrumsOfRestorationActionID = ActionID{SpellID: 35478}

// Adds drums as a major cooldown to the character, if it's being used.
func registerDrumsCD(agent Agent, partyBuffs proto.PartyBuffs, consumes proto.Consumes) {
	//character := agent.GetCharacter()
	drumsType := proto.Drums_DrumsUnknown

	// Whether this agent is the one casting the drums.
	drumsSelfCast := false

	if consumes.Drums != proto.Drums_DrumsUnknown {
		drumsType = consumes.Drums
		drumsSelfCast = true

		// Disable self-drums on other party members, so there is only 1 drummer.
		for _, partyMember := range agent.GetCharacter().Party.Players {
			if partyMember != agent {
				partyMember.GetCharacter().Consumes.Drums = proto.Drums_DrumsUnknown
			}
		}
	} else if partyBuffs.Drums != proto.Drums_DrumsUnknown {
		drumsType = partyBuffs.Drums
	}

	// If we aren't casting drums, and there is another real party member doing so, then we're done.
	if !drumsSelfCast {
		for _, partyMember := range agent.GetCharacter().Party.Players {
			if partyMember != agent && partyMember.GetCharacter().Consumes.Drums != proto.Drums_DrumsUnknown {
				return
			}
		}
	}

	var applyDrums func(sim *Simulation, character *Character)
	var actionID ActionID
	var cooldownType int32
	if drumsType == proto.Drums_DrumsOfBattle {
		applyDrums = applyDrumsOfBattle
		actionID = DrumsOfBattleActionID
		cooldownType = CooldownTypeDPS
	} else if drumsType == proto.Drums_DrumsOfRestoration {
		applyDrums = applyDrumsOfRestoration
		actionID = DrumsOfRestorationActionID
		cooldownType = CooldownTypeMana
	}

	if applyDrums == nil {
		return
	}

	mcd := MajorCooldown{
		ActionID:   actionID,
		CooldownID: DrumsCooldownID,
		Cooldown:   DrumsCD,
		Priority:   CooldownPriorityDrums,
		Type:       cooldownType,

		CanActivate:    func(sim *Simulation, character *Character) bool { return true },
		ShouldActivate: func(sim *Simulation, character *Character) bool { return true },
	}

	if drumsSelfCast {
		mcd.UsesGCD = true
		mcd.ActivationFactory = func(sim *Simulation) CooldownActivation {
			character := agent.GetCharacter()
			drumsTemplate := SimpleCast{
				Cast: Cast{
					ActionID:       actionID,
					Character:      character,
					IgnoreManaCost: true,
					CastTime:       time.Second * 1,
					GCD:            GCDDefault,
					OnCastComplete: func(sim *Simulation, cast *Cast) {
						// When a real player is using drums, their cast applies to the whole party.
						for _, agent := range character.Party.Players {
							applyDrums(sim, agent.GetCharacter())
						}

						// Drums of battle doesn't affect pets, ask Blizzard
						if drumsType != proto.Drums_DrumsOfBattle {
							for _, petAgent := range character.Party.Pets {
								pet := petAgent.GetPet()
								if pet.IsEnabled() {
									applyDrums(sim, &pet.Character)
								}
							}
						}

						// All MCDs that use the GCD and have a non-zero cast time must call this.
						character.UpdateMajorCooldowns()
					},
				},
			}

			return func(sim *Simulation, character *Character) {
				cast := drumsTemplate
				cast.Init(sim)
				cast.StartCast(sim)
			}
		}
	} else {
		// When there is no real player using drums, each player gets a fake CD that
		// gives just themself the buff, with no cast time.
		mcd.ActivationFactory = func(sim *Simulation) CooldownActivation {
			return applyDrums
			return func(sim *Simulation, character *Character) {
				applyDrums(sim, character)

				// Drums of battle doesn't affect pets, ask Blizzard
				if drumsType != proto.Drums_DrumsOfBattle {
					for _, petAgent := range character.Pets {
						pet := petAgent.GetPet()
						if pet.IsEnabled() {
							applyDrums(sim, &pet.Character)
						}
					}
				}
			}
		}
	}

	agent.GetCharacter().AddMajorCooldown(mcd)
}

func applyDrumsOfBattle(sim *Simulation, character *Character) {
	const hasteBonus = 80

	character.AddMeleeHaste(sim, hasteBonus)
	character.AddStat(stats.SpellHaste, hasteBonus)
	character.SetCD(DrumsCooldownID, sim.CurrentTime+DrumsCD)

	character.AddAura(sim, Aura{
		ID:       DrumsAuraID,
		ActionID: DrumsOfBattleActionID,
		Expires:  sim.CurrentTime + time.Second*30,
		OnExpire: func(sim *Simulation) {
			character.AddMeleeHaste(sim, -hasteBonus)
			character.AddStat(stats.SpellHaste, -hasteBonus)
		},
	})
}

func applyDrumsOfRestoration(sim *Simulation, character *Character) {
	// 600 mana over 15 seconds == 200 mp5
	const mp5Bonus = 200
	character.SetCD(DrumsCooldownID, sim.CurrentTime+DrumsCD)
	character.AddAuraWithTemporaryStats(sim, DrumsAuraID, DrumsOfRestorationActionID, stats.MP5, mp5Bonus, time.Second*15)
}

var PotionAuraID = NewAuraID()
var PotionCooldownID = NewCooldownID()

func registerPotionCD(agent Agent, consumes proto.Consumes) {
	if consumes.DefaultPotion == consumes.StartingPotion {
		// Starting potion is redundant in this case.
		consumes.StartingPotion = proto.Potions_UnknownPotion
	}
	if consumes.StartingPotion == proto.Potions_UnknownPotion {
		consumes.NumStartingPotions = 0
	}

	character := agent.GetCharacter()
	defaultPotion := consumes.DefaultPotion
	startingPotion := consumes.StartingPotion
	numStartingPotions := int(consumes.NumStartingPotions)

	defaultMCD, defaultPotionActivation := makePotionActivation(defaultPotion, character)
	startingMCD, startingPotionActivation := makePotionActivation(startingPotion, character)

	if defaultPotionActivation == nil && startingPotionActivation == nil {
		return
	}

	numStartingPotionsUsed := 0
	var mageManaGemMCD *MajorCooldown

	// Setup for expected bonus mana management.
	expectedManaPerUsage := float64((3000 + 1800) / 2)
	remainingUsages := 0
	remainingManaPotionUsages := 0
	eitherPotionIsManaPotion := startingPotion == proto.Potions_SuperManaPotion || defaultPotion == proto.Potions_SuperManaPotion
	initExpectedBonusMana := func(sim *Simulation, potionType proto.Potions, isStartingPotion bool) {
		if potionType != proto.Potions_SuperManaPotion {
			return
		}
		mageManaGemMCD = character.GetMajorCooldown(MageManaGemMCDActionID)

		remainingUsages = int(1 + (MaxDuration(0, sim.Duration))/(time.Minute*2))
		if isStartingPotion {
			remainingManaPotionUsages = MinInt(numStartingPotions, remainingUsages)
		} else {
			remainingManaPotionUsages = MaxInt(0, remainingUsages-numStartingPotions)
		}

		character.ExpectedBonusMana += expectedManaPerUsage * float64(remainingManaPotionUsages)
	}

	updateExpectedBonusMana := func(sim *Simulation, character *Character, potionType proto.Potions, isStartingPotion bool) {
		if !eitherPotionIsManaPotion {
			return
		}
		thisPotionIsManaPotion := potionType == proto.Potions_SuperManaPotion

		newRemainingUsages := int(sim.GetRemainingDuration() / (time.Minute * 2))
		numRemainingStartingPotionUsages := MinInt(numStartingPotions-numStartingPotionsUsed, remainingUsages)
		numRemainingDefaultPotionUsages := MaxInt(0, newRemainingUsages-MaxInt(0, numStartingPotions-numStartingPotionsUsed))

		newRemainingManaPotionUsages := 0
		if thisPotionIsManaPotion == isStartingPotion {
			newRemainingManaPotionUsages = numRemainingStartingPotionUsages
		} else {
			newRemainingManaPotionUsages = numRemainingDefaultPotionUsages
		}

		character.ExpectedBonusMana -= expectedManaPerUsage * float64(remainingManaPotionUsages-newRemainingManaPotionUsages)
		remainingManaPotionUsages = newRemainingManaPotionUsages
	}

	if startingPotionActivation != nil {
		character.AddMajorCooldown(MajorCooldown{
			ActionID: startingMCD.ActionID,
			CanActivate: func(sim *Simulation, character *Character) bool {
				if numStartingPotionsUsed >= numStartingPotions {
					return false
				}
				return startingMCD.CanActivate(sim, character)
			},
			ShouldActivate: func(sim *Simulation, character *Character) bool {
				if startingPotion == proto.Potions_SuperManaPotion &&
					mageManaGemMCD != nil &&
					mageManaGemMCD.IsEnabled() &&
					!mageManaGemMCD.IsOnCD(sim, character) {
					return false
				}
				return startingMCD.ShouldActivate(sim, character)
			},

			CooldownID: PotionCooldownID,
			Cooldown:   time.Minute * 2,
			Type:       startingMCD.Type,

			ActivationFactory: func(sim *Simulation) CooldownActivation {
				numStartingPotionsUsed = 0
				initExpectedBonusMana(sim, startingPotion, true)

				return func(sim *Simulation, character *Character) {
					startingPotionActivation(sim, character)
					numStartingPotionsUsed++
					updateExpectedBonusMana(sim, character, startingPotion, true)
				}
			},
		})
	}

	if defaultPotionActivation != nil {
		character.AddMajorCooldown(MajorCooldown{
			ActionID: defaultMCD.ActionID,
			CanActivate: func(sim *Simulation, character *Character) bool {
				if numStartingPotionsUsed < numStartingPotions {
					return false
				}
				return defaultMCD.CanActivate(sim, character)
			},
			ShouldActivate: func(sim *Simulation, character *Character) bool {
				if defaultPotion == proto.Potions_SuperManaPotion &&
					mageManaGemMCD != nil &&
					mageManaGemMCD.IsEnabled() &&
					!mageManaGemMCD.IsOnCD(sim, character) {
					return false
				}
				return defaultMCD.ShouldActivate(sim, character)
			},

			CooldownID: PotionCooldownID,
			Cooldown:   time.Minute * 2,
			Type:       defaultMCD.Type,

			ActivationFactory: func(sim *Simulation) CooldownActivation {
				initExpectedBonusMana(sim, defaultPotion, true)

				return func(sim *Simulation, character *Character) {
					defaultPotionActivation(sim, character)
					updateExpectedBonusMana(sim, character, defaultPotion, true)
				}
			},
		})
	}
}

const AlchStoneItemID = 35749

func makePotionActivation(potionType proto.Potions, character *Character) (MajorCooldown, CooldownActivation) {
	if potionType == proto.Potions_DestructionPotion {
		actionID := ActionID{ItemID: 22839}
		return MajorCooldown{
				ActionID: actionID,
				Type:     CooldownTypeDPS,
				CanActivate: func(sim *Simulation, character *Character) bool {
					return true
				},
				ShouldActivate: func(sim *Simulation, character *Character) bool {
					return true
				},
			},
			func(sim *Simulation, character *Character) {
				const spBonus = 120
				const critBonus = 2 * SpellCritRatingPerCritChance
				const dur = time.Second * 15

				character.AddStat(stats.SpellPower, spBonus)
				character.AddStat(stats.SpellCrit, critBonus)

				character.AddAura(sim, Aura{
					ID:       PotionAuraID,
					ActionID: actionID,
					Expires:  sim.CurrentTime + dur,
					OnExpire: func(sim *Simulation) {
						character.AddStat(stats.SpellPower, -spBonus)
						character.AddStat(stats.SpellCrit, -critBonus)
					},
				})

				character.SetCD(PotionCooldownID, time.Minute*2+sim.CurrentTime)
				character.Metrics.AddInstantCast(actionID)
			}
	} else if potionType == proto.Potions_SuperManaPotion {
		alchStoneEquipped := character.HasTrinketEquipped(AlchStoneItemID)
		actionID := ActionID{ItemID: 22832}
		return MajorCooldown{
				ActionID: actionID,
				Type:     CooldownTypeMana,
				CanActivate: func(sim *Simulation, character *Character) bool {
					return true
				},
				ShouldActivate: func(sim *Simulation, character *Character) bool {
					// Only pop if we have less than the max mana provided by the potion minus 1mp5 tick.
					totalRegen := character.ManaRegenPerSecondWhileCasting() * 5
					if character.MaxMana()-(character.CurrentMana()+totalRegen) < 3000 {
						return false
					}

					return true
				},
			},
			func(sim *Simulation, character *Character) {
				// Restores 1800 to 3000 mana. (2 Min Cooldown)
				manaGain := 1800 + (sim.RandomFloat("super mana") * 1200)

				if alchStoneEquipped {
					manaGain *= 1.4
				}

				character.AddMana(sim, manaGain, actionID, true)
				character.SetCD(PotionCooldownID, time.Minute*2+sim.CurrentTime)
				character.Metrics.AddInstantCast(actionID)
			}
	} else if potionType == proto.Potions_HastePotion {
		actionID := ActionID{ItemID: 22838}
		return MajorCooldown{
				ActionID: actionID,
				Type:     CooldownTypeDPS,
				CanActivate: func(sim *Simulation, character *Character) bool {
					return true
				},
				ShouldActivate: func(sim *Simulation, character *Character) bool {
					return true
				},
			},
			func(sim *Simulation, character *Character) {
				const hasteBonus = 400
				const dur = time.Second * 15

				character.AddMeleeHaste(sim, hasteBonus)

				character.AddAura(sim, Aura{
					ID:       PotionAuraID,
					ActionID: actionID,
					Expires:  sim.CurrentTime + dur,
					OnExpire: func(sim *Simulation) {
						character.AddMeleeHaste(sim, -hasteBonus)
					},
				})

				character.SetCD(PotionCooldownID, time.Minute*2+sim.CurrentTime)
				character.Metrics.AddInstantCast(actionID)
			}
	} else if potionType == proto.Potions_MightyRagePotion {
		actionID := ActionID{ItemID: 13442}
		return MajorCooldown{
				ActionID: actionID,
				Type:     CooldownTypeDPS,
				CanActivate: func(sim *Simulation, character *Character) bool {
					return true
				},
				ShouldActivate: func(sim *Simulation, character *Character) bool {
					if character.Class == proto.Class_ClassWarrior {
						return character.CurrentRage() < 25
					}
					return true
				},
			},
			func(sim *Simulation, character *Character) {
				const strBonus = 20.0
				const dur = time.Second * 15

				character.AddAuraWithTemporaryStats(sim, PotionAuraID, actionID, stats.Strength, strBonus, dur)
				if character.Class == proto.Class_ClassWarrior {
					bonusRage := 45.0 + (75.0-45.0)*sim.RandomFloat("Mighty Rage Potion")
					character.AddRage(sim, bonusRage, actionID)
				}

				character.SetCD(PotionCooldownID, time.Minute*2+sim.CurrentTime)
				character.Metrics.AddInstantCast(actionID)
			}
	} else if potionType == proto.Potions_FelManaPotion {
		alchStoneEquipped := character.HasTrinketEquipped(AlchStoneItemID)
		actionID := ActionID{ItemID: 31677}
		return MajorCooldown{
				ActionID: actionID,
				Type:     CooldownTypeMana,
				CanActivate: func(sim *Simulation, character *Character) bool {
					return true
				},
				ShouldActivate: func(sim *Simulation, character *Character) bool {
					// Only pop if we have low enough mana. The potion takes effect over 24
					// seconds so we can pop it a little earlier than the full value.
					if character.MaxMana()-character.CurrentMana() < 2000 {
						return false
					}

					return true
				},
			},
			func(sim *Simulation, character *Character) {
				// Restores 3200 mana over 24 seconds.
				manaGain := 3200.0
				if alchStoneEquipped {
					manaGain *= 1.4
				}
				mp5 := manaGain / 24 * 5
				character.AddAuraWithTemporaryStats(sim, FelManaPotionAuraID, actionID, stats.MP5, mp5, time.Second*24)

				if !character.HasAura(FelManaPotionDebuffAuraID) {
					character.AddStat(stats.SpellPower, -25)
					character.AddAura(sim, Aura{
						ID:       FelManaPotionDebuffAuraID,
						ActionID: ActionID{SpellID: 38927},
						Expires:  NeverExpires,
					})
				}

				character.SetCD(PotionCooldownID, time.Minute*2+sim.CurrentTime)
				character.Metrics.AddInstantCast(actionID)
			}
	} else {
		return MajorCooldown{}, nil
	}
}

var FelManaPotionAuraID = NewAuraID()
var FelManaPotionDebuffAuraID = NewAuraID()

var ConjuredAuraID = NewAuraID()
var ConjuredCooldownID = NewCooldownID()

func registerConjuredCD(agent Agent, consumes proto.Consumes) {
	character := agent.GetCharacter()

	mcd, activationFn := makeConjuredActivation(consumes.DefaultConjured, character)
	if activationFn == nil {
		return
	}

	var mageManaGemMCD *MajorCooldown

	character.AddMajorCooldown(MajorCooldown{
		ActionID:    mcd.ActionID,
		CooldownID:  ConjuredCooldownID,
		Cooldown:    mcd.Cooldown,
		CanActivate: mcd.CanActivate,
		Type:        mcd.Type,
		ShouldActivate: func(sim *Simulation, character *Character) bool {
			if consumes.DefaultConjured == proto.Conjured_ConjuredDarkRune &&
				mageManaGemMCD != nil &&
				mageManaGemMCD.IsEnabled() &&
				!mageManaGemMCD.IsOnCD(sim, character) {
				return false
			}
			return mcd.ShouldActivate(sim, character)
		},
		ActivationFactory: func(sim *Simulation) CooldownActivation {
			mageManaGemMCD = character.GetMajorCooldown(MageManaGemMCDActionID)
			expectedManaPerUsage := float64((900 + 600) / 2)

			remainingUsages := int(1 + (MaxDuration(0, sim.Duration))/(time.Minute*2))

			if consumes.DefaultConjured == proto.Conjured_ConjuredDarkRune {
				character.ExpectedBonusMana += expectedManaPerUsage * float64(remainingUsages)
			}

			return func(sim *Simulation, character *Character) {
				activationFn(sim, character)

				if consumes.DefaultConjured == proto.Conjured_ConjuredDarkRune {
					// Update expected bonus mana
					newRemainingUsages := int(sim.GetRemainingDuration() / (time.Minute * 2))
					character.ExpectedBonusMana -= expectedManaPerUsage * float64(remainingUsages-newRemainingUsages)
					remainingUsages = newRemainingUsages
				}
			}
		},
	})
}

func makeConjuredActivation(conjuredType proto.Conjured, character *Character) (MajorCooldown, CooldownActivation) {
	if conjuredType == proto.Conjured_ConjuredDarkRune {
		actionID := ActionID{ItemID: 20520}
		return MajorCooldown{
				ActionID: actionID,
				Cooldown: time.Minute * 2,
				Type:     CooldownTypeMana,
				CanActivate: func(sim *Simulation, character *Character) bool {
					return true
				},
				ShouldActivate: func(sim *Simulation, character *Character) bool {
					// Only pop if we have less than the max mana provided by the potion minus 1mp5 tick.
					totalRegen := character.ManaRegenPerSecondWhileCasting() * 5
					if character.MaxMana()-(character.CurrentMana()+totalRegen) < 1500 {
						return false
					}
					return true
				},
			},
			func(sim *Simulation, character *Character) {
				// Restores 900 to 1500 mana. (2 Min Cooldown)
				manaGain := 900 + (sim.RandomFloat("dark rune") * 600)
				character.AddMana(sim, manaGain, actionID, true)
				character.SetCD(ConjuredCooldownID, time.Minute*2+sim.CurrentTime)
				character.Metrics.AddInstantCast(actionID)
			}
	} else if conjuredType == proto.Conjured_ConjuredFlameCap {
		actionID := ActionID{ItemID: 22788}

		castTemplate := NewSimpleSpellTemplate(SimpleSpell{
			SpellCast: SpellCast{
				Cast: Cast{
					ActionID:       actionID,
					Character:      character,
					IgnoreManaCost: true,
					IsPhantom:      true,
					SpellSchool:    stats.FireSpellPower,
					CritMultiplier: 1.5,
				},
			},
			Effect: SpellHitEffect{
				SpellEffect: SpellEffect{
					DamageMultiplier:       1,
					StaticDamageMultiplier: 1,
				},
				DirectInput: DirectDamageInput{
					MinBaseDamage: 40,
					MaxBaseDamage: 40,
				},
			},
		})
		spellObj := SimpleSpell{}

		return MajorCooldown{
				ActionID: actionID,
				Cooldown: time.Minute * 3,
				Type:     CooldownTypeDPS,
				CanActivate: func(sim *Simulation, character *Character) bool {
					return true
				},
				ShouldActivate: func(sim *Simulation, character *Character) bool {
					return true
				},
			},
			func(sim *Simulation, character *Character) {
				const fireBonus = 80
				const dur = time.Minute * 1
				const procChance = 0.185

				aura := character.NewAuraWithTemporaryStats(sim, ConjuredAuraID, actionID, stats.FireSpellPower, fireBonus, dur)
				aura.OnMeleeAttack = func(sim *Simulation, ability *ActiveMeleeAbility, hitEffect *AbilityHitEffect) {
					if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(ProcMaskMeleeOrRanged) || ability.IsPhantom {
						return
					}
					if sim.RandomFloat("Flame Cap Melee") > procChance {
						return
					}

					castAction := &spellObj
					castTemplate.Apply(castAction)
					castAction.Effect.Target = hitEffect.Target
					castAction.Init(sim)
					castAction.Cast(sim)
				}
				character.AddAura(sim, aura)

				character.SetCD(ConjuredCooldownID, time.Minute*3+sim.CurrentTime)
				character.Metrics.AddInstantCast(actionID)
			}
	} else {
		return MajorCooldown{}, nil
	}
}

// Mana-related consumes need to know this so they can wait for the Mage to use
// gems first.
var MageManaGemMCDActionID = ActionID{ItemID: 22044}
