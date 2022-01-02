package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Registers all consume-related effects to the Agent.
func applyConsumeEffects(agent Agent, partyBuffs proto.PartyBuffs) {
	consumes := agent.GetCharacter().consumes

	agent.GetCharacter().AddStats(consumesStats(consumes))

	// TODO: demon slaying elixir needs an aura to only apply to demons...
	//  The other option is to include target type in this function.

	registerDrumsCD(agent, partyBuffs, consumes)
	registerPotionCD(agent, consumes)
	registerConjuredCD(agent, consumes)
}

func consumesStats(c proto.Consumes) stats.Stats {
	s := stats.Stats{}

	if c.BrilliantWizardOil {
		s[stats.SpellCrit] += 14
		s[stats.SpellPower] += 36
		s[stats.HealingPower] += 36
	}
	if c.SuperiorWizardOil {
		s[stats.SpellPower] += 42
		s[stats.HealingPower] += 42
	}

	if c.ElixirOfMajorMageblood {
		s[stats.MP5] += 16.0
	}
	if c.AdeptsElixir {
		s[stats.SpellCrit] += 24
		s[stats.SpellPower] += 24
		s[stats.HealingPower] += 24
	}
	if c.ElixirOfMajorFirePower {
		s[stats.FireSpellPower] += 55
	}
	if c.ElixirOfMajorFrostPower {
		s[stats.FrostSpellPower] += 55
	}
	if c.ElixirOfMajorShadowPower {
		s[stats.ShadowSpellPower] += 55
	}
	if c.ElixirOfDraenicWisdom {
		s[stats.Intellect] += 30
		s[stats.Spirit] += 30
	}
	if c.ElixirOfMajorAgility {
		s[stats.Agility] += 35
		s[stats.MeleeCrit] += 20
	}

	if c.FlaskOfSupremePower {
		s[stats.SpellPower] += 70
	}
	if c.FlaskOfBlindingLight {
		s[stats.NatureSpellPower] += 80
		s[stats.ArcaneSpellPower] += 80
		s[stats.HolySpellPower] += 80
	}
	if c.FlaskOfPureDeath {
		s[stats.FireSpellPower] += 80
		s[stats.FrostSpellPower] += 80
		s[stats.ShadowSpellPower] += 80
	}
	if c.FlaskOfMightyRestoration {
		s[stats.MP5] += 25
	}
	if c.FlaskOfRelentlessAssault {
		s[stats.AttackPower] += 120
	}

	if c.BlackenedBasilisk {
		s[stats.SpellPower] += 23
		s[stats.HealingPower] += 23
		s[stats.Spirit] += 20
	}
	if c.SkullfishSoup {
		s[stats.SpellCrit] += 20
		s[stats.Spirit] += 20
	}
	if c.KreegsStoutBeatdown {
		s[stats.Intellect] -= 5
		s[stats.Spirit] += 25
	}
	if c.RoastedClefthoof {
		s[stats.Strength] += 20
	}
	if c.ScrollOfAgilityV {
		s[stats.Agility] += 20
	}
	if c.ScrollOfStrengthV {
		s[stats.Strength] += 20
	}

	return s
}

var DrumsAuraID = NewAuraID()
var DrumsCooldownID = NewCooldownID()

const DrumsCD = time.Minute * 2 // Tinnitus

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
				partyMember.GetCharacter().consumes.Drums = proto.Drums_DrumsUnknown
			}
		}
	} else if partyBuffs.Drums != proto.Drums_DrumsUnknown {
		drumsType = partyBuffs.Drums
	}

	// If we aren't casting drums, and there is another real party member doing so, then we're done.
	if !drumsSelfCast {
		for _, partyMember := range agent.GetCharacter().Party.Players {
			if partyMember != agent && partyMember.GetCharacter().consumes.Drums != proto.Drums_DrumsUnknown {
				return
			}
		}
	}

	var applyDrums func(sim *Simulation, character *Character)
	spellID := int32(0)
	if drumsType == proto.Drums_DrumsOfBattle {
		applyDrums = applyDrumsOfBattle
		spellID = 35476
	} else if drumsType == proto.Drums_DrumsOfRestoration {
		applyDrums = applyDrumsOfRestoration
		spellID = 35478
	}

	if applyDrums == nil {
		return
	}

	mcd := MajorCooldown{
		ActionID:   ActionID{SpellID: spellID},
		CooldownID: DrumsCooldownID,
		Cooldown:   DrumsCD,
		Priority:   CooldownPriorityDrums,

		CanActivate:    func(sim *Simulation, character *Character) bool { return true },
		ShouldActivate: func(sim *Simulation, character *Character) bool { return true },
	}

	if drumsSelfCast {
		// When a real player is using drums, their cast applies to the whole party.
		mcd.ActivationFactory = func(sim *Simulation) CooldownActivation {
			return func(sim *Simulation, character *Character) {
				for _, agent := range character.Party.Players {
					applyDrums(sim, agent.GetCharacter())
				}
				// TODO: Do a cast time
				character.Metrics.AddInstantCast(ActionID{SpellID: spellID})
			}
		}
	} else {
		// When there is no real player using drums, each player gets a fake CD that
		// gives just themself the buff, with no cast time.
		mcd.ActivationFactory = func(sim *Simulation) CooldownActivation {
			return applyDrums
		}
	}

	agent.GetCharacter().AddMajorCooldown(mcd)
}

func applyDrumsOfBattle(sim *Simulation, character *Character) {
	const hasteBonus = 80
	character.SetCD(DrumsCooldownID, sim.CurrentTime+DrumsCD)
	character.AddAuraWithTemporaryStats(sim, DrumsAuraID, 35476, "Drums of Battle", stats.SpellHaste, hasteBonus, time.Second*30)
	// TODO: Add melee haste to drums
}

func applyDrumsOfRestoration(sim *Simulation, character *Character) {
	// 600 mana over 15 seconds == 200 mp5
	const mp5Bonus = 200
	character.SetCD(DrumsCooldownID, sim.CurrentTime+DrumsCD)
	character.AddAuraWithTemporaryStats(sim, DrumsAuraID, 35478, "Drums of Restoration", stats.MP5, mp5Bonus, time.Second*15)
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
		return MajorCooldown{
				ActionID: ActionID{ItemID: 22839},
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
					ID:      PotionAuraID,
					SpellID: 28508,
					Name:    "Destruction Potion",
					Expires: sim.CurrentTime + dur,
					OnExpire: func(sim *Simulation) {
						character.AddStat(stats.SpellPower, -spBonus)
						character.AddStat(stats.SpellCrit, -critBonus)
					},
				})

				character.SetCD(PotionCooldownID, time.Minute*2+sim.CurrentTime)
				character.Metrics.AddInstantCast(ActionID{ItemID: 22839})
			}
	} else if potionType == proto.Potions_SuperManaPotion {
		alchStoneEquipped := character.HasTrinketEquipped(AlchStoneItemID)
		return MajorCooldown{
				ActionID: ActionID{ItemID: 22832},
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

				character.AddMana(sim, manaGain, "Super Mana Potion", true)
				character.SetCD(PotionCooldownID, time.Minute*2+sim.CurrentTime)
				character.Metrics.AddInstantCast(ActionID{ItemID: 22832})
			}
	} else if potionType == proto.Potions_HastePotion {
		return MajorCooldown{
				ActionID: ActionID{ItemID: 22838},
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
					ID:      PotionAuraID,
					SpellID: 28507,
					Name:    "Haste Potion",
					Expires: sim.CurrentTime + dur,
					OnExpire: func(sim *Simulation) {
						character.AddMeleeHaste(sim, -hasteBonus)
					},
				})

				character.SetCD(PotionCooldownID, time.Minute*2+sim.CurrentTime)
				character.Metrics.AddInstantCast(ActionID{ItemID: 22838})
			}
	} else {
		return MajorCooldown{}, nil
	}
}

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
		return MajorCooldown{
				ActionID: ActionID{ItemID: 20520},
				Cooldown: time.Minute * 2,
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
				character.AddMana(sim, manaGain, "Dark Rune", true)
				character.SetCD(ConjuredCooldownID, time.Minute*2+sim.CurrentTime)
				character.Metrics.AddInstantCast(ActionID{ItemID: 20520})
			}
	} else if conjuredType == proto.Conjured_ConjuredFlameCap {
		return MajorCooldown{
				ActionID: ActionID{ItemID: 22788},
				Cooldown: time.Minute * 3,
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

				character.AddAuraWithTemporaryStats(sim, ConjuredAuraID, 28714, "Flame Cap", stats.FireSpellPower, fireBonus, dur)
				// TODO: Add separate aura for damage proc on melee/ranged swings

				character.SetCD(ConjuredCooldownID, time.Minute*3+sim.CurrentTime)
				character.Metrics.AddInstantCast(ActionID{ItemID: 22788})
			}
	} else {
		return MajorCooldown{}, nil
	}
}

// Mana-related consumes need to know this so they can wait for the Mage to use
// gems first.
var MageManaGemMCDActionID = ActionID{ItemID: 22044}
