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
			character.RegisterResetEffect(func(sim *Simulation) {
				if sim.GetPrimaryTarget().MobType == proto.MobType_MobTypeDemon {
					character.PseudoStats.MobTypeAttackPower += 265
				}
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
		case proto.BattleElixir_FelStrengthElixir:
			character.AddStats(stats.Stats{
				stats.AttackPower:       90,
				stats.RangedAttackPower: 90,
				stats.Stamina:           -10,
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
	allowMHImbue := character.HasMHWeapon() && (character.HasMHWeaponImbue || partyBuffs.WindfuryTotemRank == 0)
	if allowMHImbue {
		addImbueStats(character, consumes.MainHandImbue)
	}
	if character.HasOHWeapon() {
		addImbueStats(character, consumes.OffHandImbue)
	}
	applyAdamantiteSharpeningStoneAura(character, consumes, allowMHImbue)

	character.RegisterFinalizeEffect(func() {
		registerPotionCD(agent, consumes)
		registerConjuredCD(agent, consumes)
	})
	registerDrumsCD(agent, partyBuffs, consumes)
	registerExplosivesCD(agent, consumes)
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
		character.PseudoStats.BonusDamage += 12
		// Melee crit component handled separately because its melee-only.
	} else if imbue == proto.WeaponImbue_WeaponImbueAdamantiteWeightstone {
		character.PseudoStats.BonusDamage += 12
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

func applyAdamantiteSharpeningStoneAura(character *Character, consumes proto.Consumes, allowMHImbue bool) {
	critBonus := 0.0
	if consumes.MainHandImbue == proto.WeaponImbue_WeaponImbueAdamantiteSharpeningStone && allowMHImbue {
		critBonus += 14
	}
	if consumes.OffHandImbue == proto.WeaponImbue_WeaponImbueAdamantiteSharpeningStone {
		critBonus += 14
	}

	if critBonus == 0 {
		return
	}

	// Crit rating for sharpening stone applies to melee only.
	if character.Class != proto.Class_ClassHunter {
		// For non-hunters just give direct crit so it shows on the stats panel.
		character.AddStats(stats.Stats{
			stats.MeleeCrit: critBonus,
		})
	} else {
		character.PseudoStats.BonusMeleeCritRating += critBonus
	}
}

var DrumsAuraTag = "Drums"
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

	var actionID ActionID
	var cooldownType int32
	if drumsType == proto.Drums_DrumsOfBattle {
		actionID = DrumsOfBattleActionID
		cooldownType = CooldownTypeDPS
	} else if drumsType == proto.Drums_DrumsOfRestoration {
		actionID = DrumsOfRestorationActionID
		cooldownType = CooldownTypeMana
	} else {
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

			auras := []*Aura{}
			for _, agent := range character.Party.Players {
				auras = append(auras, makeDrumsAura(agent.GetCharacter(), drumsType))
			}

			drumsTemplate := SimpleCast{
				Cast: Cast{
					ActionID:  actionID,
					Character: character,
					GCD:       GCDDefault,
					OnCastComplete: func(sim *Simulation, cast *Cast) {
						// When a real player is using drums, their cast applies to the whole party.
						for _, aura := range auras {
							aura.Activate(sim)
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
			drumsAura := makeDrumsAura(agent.GetCharacter(), drumsType)
			return func(sim *Simulation, character *Character) {
				drumsAura.Activate(sim)
			}
		}
	}

	agent.GetCharacter().AddMajorCooldown(mcd)
}

func makeDrumsAura(character *Character, drumsType proto.Drums) *Aura {
	if drumsType == proto.Drums_DrumsOfBattle {
		const hasteBonus = 80
		return character.NewTemporaryStatsAuraWrapped("Drums of Battle", DrumsOfBattleActionID, stats.Stats{stats.MeleeHaste: hasteBonus, stats.SpellHaste: hasteBonus}, time.Second*30, func(aura *Aura) {
			oldOnGain := aura.OnGain
			aura.OnGain = func(aura *Aura, sim *Simulation) {
				oldOnGain(aura, sim)
				character.SetCD(DrumsCooldownID, sim.CurrentTime+DrumsCD)
				// Drums of battle doesn't affect pets, ask Blizzard.
			}
		})
	} else if drumsType == proto.Drums_DrumsOfRestoration {
		// 600 mana over 15 seconds == 200 mp5
		const mp5Bonus = 200

		petAuras := []*Aura{}
		for _, petAgent := range character.Party.Pets {
			pet := petAgent.GetPet()
			petAuras = append(petAuras, pet.NewTemporaryStatsAura("Drums of Restoration", DrumsOfRestorationActionID, stats.Stats{stats.MP5: mp5Bonus}, time.Second*15))
		}

		return character.NewTemporaryStatsAuraWrapped("Drums of Restoration", DrumsOfRestorationActionID, stats.Stats{stats.MP5: mp5Bonus}, time.Second*15, func(aura *Aura) {
			oldOnGain := aura.OnGain
			aura.OnGain = func(aura *Aura, sim *Simulation) {
				oldOnGain(aura, sim)
				character.SetCD(DrumsCooldownID, sim.CurrentTime+DrumsCD)

				for i, petAgent := range character.Party.Pets {
					pet := petAgent.GetPet()
					if pet.IsEnabled() {
						petAuras[i].Activate(sim)
					}
				}
			}
		})
	} else {
		return nil
	}
}

var PotionAuraTag = "Potion"
var PotionCooldownID = NewCooldownID()

func registerPotionCD(agent Agent, consumes proto.Consumes) {
	if consumes.DefaultPotion == consumes.StartingPotion {
		// Starting potion is redundant in this case.
		consumes.StartingPotion = proto.Potions_UnknownPotion
	}
	if consumes.StartingPotion == proto.Potions_UnknownPotion {
		consumes.NumStartingPotions = 0
	}
	if consumes.NumStartingPotions == 0 {
		consumes.StartingPotion = proto.Potions_UnknownPotion
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
		aura := character.NewTemporaryStatsAura("Destruction Potion", actionID, stats.Stats{stats.SpellPower: 120, stats.SpellCrit: 2 * SpellCritRatingPerCritChance}, time.Second*15)
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
				aura.Activate(sim)
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
		aura := character.NewTemporaryStatsAura("Haste Potion", actionID, stats.Stats{stats.MeleeHaste: 400}, time.Second*15)
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
				aura.Activate(sim)
				character.SetCD(PotionCooldownID, time.Minute*2+sim.CurrentTime)
				character.Metrics.AddInstantCast(actionID)
			}
	} else if potionType == proto.Potions_MightyRagePotion {
		actionID := ActionID{ItemID: 13442}
		aura := character.NewTemporaryStatsAura("Mighty Rage Potion", actionID, stats.Stats{stats.Strength: 20}, time.Second*15)
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
				aura.Activate(sim)
				if character.Class == proto.Class_ClassWarrior {
					bonusRage := 45.0 + (75.0-45.0)*sim.RandomFloat("Mighty Rage Potion")
					character.AddRage(sim, bonusRage, actionID)
				}

				character.SetCD(PotionCooldownID, time.Minute*2+sim.CurrentTime)
				character.Metrics.AddInstantCast(actionID)
			}
	} else if potionType == proto.Potions_FelManaPotion {
		actionID := ActionID{ItemID: 31677}

		// Restores 3200 mana over 24 seconds.
		manaGain := 3200.0
		alchStoneEquipped := character.HasTrinketEquipped(AlchStoneItemID)
		if alchStoneEquipped {
			manaGain *= 1.4
		}
		mp5 := manaGain / 24 * 5

		buffAura := character.NewTemporaryStatsAura("Fel Mana Potion", actionID, stats.Stats{stats.MP5: mp5}, time.Second*24)
		debuffAura := character.NewTemporaryStatsAura("Fel Mana Potion Debuff", ActionID{SpellID: 38927}, stats.Stats{stats.SpellPower: -25}, time.Minute*15)

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
				buffAura.Activate(sim)
				debuffAura.Activate(sim)
				debuffAura.Refresh(sim)

				character.SetCD(PotionCooldownID, time.Minute*2+sim.CurrentTime)
				character.Metrics.AddInstantCast(actionID)
			}
	} else {
		return MajorCooldown{}, nil
	}
}

var ConjuredAuraTag = "Conjured"
var ConjuredCooldownID = NewCooldownID()

func registerConjuredCD(agent Agent, consumes proto.Consumes) {
	if consumes.DefaultConjured == consumes.StartingConjured {
		// Starting conjured is redundant in this case.
		consumes.StartingConjured = proto.Conjured_ConjuredUnknown
	}
	if consumes.StartingConjured == proto.Conjured_ConjuredUnknown {
		consumes.NumStartingConjured = 0
	}
	if consumes.NumStartingConjured == 0 {
		consumes.StartingConjured = proto.Conjured_ConjuredUnknown
	}
	character := agent.GetCharacter()

	defaultMCD, defaultActivationFn := makeConjuredActivation(consumes.DefaultConjured, character)
	startingMCD, startingActivationFn := makeConjuredActivation(consumes.StartingConjured, character)
	numStartingConjured := int(consumes.NumStartingConjured)
	if defaultActivationFn == nil && startingActivationFn == nil {
		return
	}

	numStartingConjuredUsed := 0

	if startingActivationFn != nil {
		character.AddMajorCooldown(MajorCooldown{
			ActionID:   startingMCD.ActionID,
			CooldownID: ConjuredCooldownID,
			Cooldown:   startingMCD.Cooldown,
			Type:       startingMCD.Type,
			CanActivate: func(sim *Simulation, character *Character) bool {
				if numStartingConjuredUsed >= numStartingConjured {
					return false
				}
				return startingMCD.CanActivate(sim, character)
			},
			ShouldActivate: func(sim *Simulation, character *Character) bool {
				return startingMCD.ShouldActivate(sim, character)
			},
			ActivationFactory: func(sim *Simulation) CooldownActivation {
				numStartingConjuredUsed = 0
				return func(sim *Simulation, character *Character) {
					startingActivationFn(sim, character)
					numStartingConjuredUsed++
				}
			},
		})
	}

	if defaultActivationFn != nil {
		character.AddMajorCooldown(MajorCooldown{
			ActionID:   defaultMCD.ActionID,
			CooldownID: ConjuredCooldownID,
			Cooldown:   defaultMCD.Cooldown,
			Type:       defaultMCD.Type,
			CanActivate: func(sim *Simulation, character *Character) bool {
				if numStartingConjuredUsed < numStartingConjured {
					return false
				}
				return defaultMCD.CanActivate(sim, character)
			},
			ShouldActivate: func(sim *Simulation, character *Character) bool {
				return defaultMCD.ShouldActivate(sim, character)
			},
			ActivationFactory: func(sim *Simulation) CooldownActivation {
				expectedManaPerUsage := float64((900 + 600) / 2)

				remainingUsages := int(1 + (MaxDuration(0, sim.Duration))/(time.Minute*2))

				if consumes.DefaultConjured == proto.Conjured_ConjuredDarkRune {
					character.ExpectedBonusMana += expectedManaPerUsage * float64(remainingUsages)
				}

				return func(sim *Simulation, character *Character) {
					defaultActivationFn(sim, character)

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

		flameCapProc := character.RegisterSpell(SpellConfig{
			Template: SimpleSpell{
				SpellCast: SpellCast{
					Cast: Cast{
						ActionID:    actionID,
						Character:   character,
						SpellSchool: SpellSchoolFire,
					},
				},
				Effect: SpellEffect{
					IsPhantom:           true,
					OutcomeRollCategory: OutcomeRollCategoryMagic,
					CritRollCategory:    CritRollCategoryMagical,
					CritMultiplier:      1.5,
					DamageMultiplier:    1,
					ThreatMultiplier:    1,
					BaseDamage:          BaseDamageConfigFlat(40),
				},
			},
			ModifyCast: ModifyCastAssignTarget,
		})

		const procChance = 0.185
		flameCapAura := character.NewTemporaryStatsAura("Flame Cap", actionID, stats.Stats{stats.FireSpellPower: 80}, time.Minute)
		flameCapAura.OnSpellHit = func(aura *Aura, sim *Simulation, spell *Spell, spellEffect *SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(ProcMaskMeleeOrRanged) || spellEffect.IsPhantom {
				return
			}
			if sim.RandomFloat("Flame Cap Melee") > procChance {
				return
			}

			flameCapProc.Cast(sim, spellEffect.Target)
		}

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
				flameCapAura.Activate(sim)
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

var ExplosivesCooldownID = NewCooldownID()
var SuperSapperCooldownID = NewCooldownID()
var GoblinSapperCooldownID = NewCooldownID()
var SuperSapperActionID = ActionID{ItemID: 23827, CooldownID: SuperSapperCooldownID}
var GoblinSapperActionID = ActionID{ItemID: 10646, CooldownID: GoblinSapperCooldownID}
var FelIronBombActionID = ActionID{ItemID: 23736}
var AdamantiteGrenadeActionID = ActionID{ItemID: 23737}
var HolyWaterActionID = ActionID{ItemID: 13180}

func registerExplosivesCD(agent Agent, consumes proto.Consumes) {
	if !consumes.SuperSapper && !consumes.GoblinSapper && consumes.FillerExplosive == proto.Explosive_ExplosiveUnknown {
		return
	}
	character := agent.GetCharacter()

	var gnomishFlameTurretCaster func(sim *Simulation)
	if consumes.FillerExplosive == proto.Explosive_ExplosiveGnomishFlameTurret {
		gnomishFlameTurretCaster = character.newGnomishFlameTurretCaster()
	}

	character.AddMajorCooldown(MajorCooldown{
		ActionID:   SuperSapperActionID,
		CooldownID: ExplosivesCooldownID,
		Cooldown:   time.Minute,
		Type:       CooldownTypeDPS,
		CanActivate: func(sim *Simulation, character *Character) bool {
			return true
		},
		ShouldActivate: func(sim *Simulation, character *Character) bool {
			return true
		},
		ActivationFactory: func(sim *Simulation) CooldownActivation {
			superSapperCaster := character.newSuperSapperCaster(sim)
			goblinSapperCaster := character.newGoblinSapperCaster(sim)

			var castFiller func(sim *Simulation)
			hasFiller := false
			switch consumes.FillerExplosive {
			case proto.Explosive_ExplosiveFelIronBomb:
				hasFiller = true
				castFiller = character.newFelIronBombCaster(sim)
			case proto.Explosive_ExplosiveAdamantiteGrenade:
				hasFiller = true
				castFiller = character.newAdamantiteGrenadeCaster(sim)
			case proto.Explosive_ExplosiveGnomishFlameTurret:
				hasFiller = true
				castFiller = gnomishFlameTurretCaster
			case proto.Explosive_ExplosiveHolyWater:
				hasFiller = true
				castFiller = character.newHolyWaterCaster(sim)
			}

			cdAfterGoblinSapper := time.Minute
			if !hasFiller {
				if consumes.SuperSapper {
					cdAfterGoblinSapper = time.Minute * 4
				} else {
					cdAfterGoblinSapper = time.Minute * 5
				}
			}

			cdAfterSuperSapper := time.Minute
			if !hasFiller && !consumes.GoblinSapper {
				cdAfterSuperSapper = time.Minute * 5
			}

			return func(sim *Simulation, character *Character) {
				var castExplosive func(sim *Simulation)
				var cd time.Duration
				if consumes.SuperSapper && !character.IsOnCD(SuperSapperCooldownID, sim.CurrentTime) {
					castExplosive = superSapperCaster
					cd = cdAfterSuperSapper
				} else if consumes.GoblinSapper && !character.IsOnCD(GoblinSapperCooldownID, sim.CurrentTime) {
					castExplosive = goblinSapperCaster
					cd = cdAfterGoblinSapper
				} else {
					castExplosive = castFiller
					cd = time.Minute
				}
				castExplosive(sim)
				character.SetCD(ExplosivesCooldownID, sim.CurrentTime+cd)
			}
		},
	})
}

// Creates a spell object for the common explosive case.
func (character *Character) newBasicExplosiveSpell(sim *Simulation, actionID ActionID, minDamage float64, maxDamage float64, cooldown time.Duration) SpellConfig {
	spell := SimpleSpell{
		SpellCast: SpellCast{
			Cast: Cast{
				ActionID:    actionID,
				Character:   character,
				SpellSchool: SpellSchoolFire,
				Cooldown:    cooldown,
			},
		},
	}

	baseEffect := SpellEffect{
		DamageMultiplier:    1,
		ThreatMultiplier:    1,
		OutcomeRollCategory: OutcomeRollCategoryMagic,
		CritRollCategory:    CritRollCategoryMagical,
		CritMultiplier:      2,

		// Explosives always have 1% resist chance, so just give them hit cap.
		BonusSpellHitRating: 100 * SpellHitRatingPerHitChance,

		BaseDamage: BaseDamageConfigRoll(minDamage, maxDamage),
	}

	numHits := sim.GetNumTargets()
	effects := make([]SpellEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)
	}
	spell.Effects = effects

	return SpellConfig{
		Template: spell,
	}
}
func (character *Character) newSuperSapperCaster(sim *Simulation) func(sim *Simulation) {
	spell := character.GetOrRegisterSpell(character.newBasicExplosiveSpell(sim, SuperSapperActionID, 900, 1500, time.Minute*5))
	return func(sim *Simulation) {
		spell.Cast(sim, nil)
	}
}
func (character *Character) newGoblinSapperCaster(sim *Simulation) func(sim *Simulation) {
	spell := character.GetOrRegisterSpell(character.newBasicExplosiveSpell(sim, GoblinSapperActionID, 450, 750, time.Minute*5))
	return func(sim *Simulation) {
		spell.Cast(sim, nil)
	}
}
func (character *Character) newFelIronBombCaster(sim *Simulation) func(sim *Simulation) {
	spell := character.GetOrRegisterSpell(character.newBasicExplosiveSpell(sim, FelIronBombActionID, 330, 770, 0))
	return func(sim *Simulation) {
		spell.Cast(sim, nil)
	}
}
func (character *Character) newAdamantiteGrenadeCaster(sim *Simulation) func(sim *Simulation) {
	spell := character.GetOrRegisterSpell(character.newBasicExplosiveSpell(sim, AdamantiteGrenadeActionID, 450, 750, 0))
	return func(sim *Simulation) {
		spell.Cast(sim, nil)
	}
}
func (character *Character) newHolyWaterCaster(sim *Simulation) func(sim *Simulation) {
	config := character.newBasicExplosiveSpell(sim, HolyWaterActionID, 438, 562, 0)
	config.Template.SpellSchool = SpellSchoolHoly
	for i, _ := range config.Template.Effects {
		effect := &config.Template.Effects[i]
		if effect.Target.MobType != proto.MobType_MobTypeUndead {
			effect.DamageMultiplier = 0
		}
	}
	spell := character.GetOrRegisterSpell(config)

	return func(sim *Simulation) {
		spell.Cast(sim, nil)
	}
}
