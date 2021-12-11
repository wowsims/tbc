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

	registerDrumsCD(agent, partyBuffs, consumes)
	registerPotionCD(agent, consumes)
	registerDarkRuneCD(agent, consumes)
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

	return s
}

// Adds drums as a major cooldown to the character, if it's being used.
var DrumsAuraID = NewAuraID()
var DrumsCooldownID = NewCooldownID()

func registerDrumsCD(agent Agent, partyBuffs proto.PartyBuffs, consumes proto.Consumes) {
	//character := agent.GetCharacter()
	drumsType := proto.Drums_DrumsUnknown

	// Whether this agent is the one casting the drums.
	drumsSelfCast := false

	if consumes.Drums != proto.Drums_DrumsUnknown {
		drumsType = consumes.Drums
		//drumsSelfCast = true
	} else if partyBuffs.Drums != proto.Drums_DrumsUnknown {
		drumsType = partyBuffs.Drums
	}

	// TODO: If drumsSelfCast == true, then do a cast time
	mcd := MajorCooldown{
		CooldownID: DrumsCooldownID,
		Cooldown:   time.Minute * 2,
		Priority:   CooldownPriorityDrums,
	}

	if drumsType == proto.Drums_DrumsOfBattle {
		mcd.ActivationFactory = func(sim *Simulation) CooldownActivation {
			return func(sim *Simulation, character *Character) bool {
				const hasteBonus = 80
				for _, agent := range character.Party.Players {
					agent.GetCharacter().SetCD(DrumsCooldownID, time.Minute*2+sim.CurrentTime) // tinnitus
					agent.GetCharacter().AddAuraWithTemporaryStats(sim, DrumsAuraID, 35476, "Drums of Battle", stats.SpellHaste, hasteBonus, time.Second*30)
				}
				if drumsSelfCast {
					character.Metrics.AddInstantCast(ActionID{SpellID: 35476})
				}
				return true
			}
		}
	} else if drumsType == proto.Drums_DrumsOfRestoration {
		mcd.ActivationFactory = func(sim *Simulation) CooldownActivation {
			return func(sim *Simulation, character *Character) bool {
				// 600 mana over 15 seconds == 200 mp5
				const mp5Bonus = 200
				for _, agent := range character.Party.Players {
					agent.GetCharacter().SetCD(DrumsCooldownID, time.Minute*2+sim.CurrentTime) // tinnitus
					agent.GetCharacter().AddAuraWithTemporaryStats(sim, DrumsAuraID, 35478, "Drums of Restoration", stats.MP5, mp5Bonus, time.Second*15)
				}
				if drumsSelfCast {
					character.Metrics.AddInstantCast(ActionID{SpellID: 35478})
				}
				return true
			}
		}
	}

	if mcd.ActivationFactory != nil {
		agent.GetCharacter().AddMajorCooldown(mcd)
	}
}

var PotionAuraID = NewAuraID()
var PotionCooldownID = NewCooldownID()

func registerPotionCD(agent Agent, consumes proto.Consumes) {
	character := agent.GetCharacter()
	defaultPotion := consumes.DefaultPotion
	startingPotion := consumes.StartingPotion
	numStartingPotions := int(consumes.NumStartingPotions)

	defaultPotionActivation := makePotionActivation(defaultPotion, character)
	startingPotionActivation := makePotionActivation(startingPotion, character)

	if startingPotionActivation == nil {
		numStartingPotions = 0
	}

	if defaultPotionActivation == nil && startingPotionActivation == nil {
		return
	}

	agent.GetCharacter().AddMajorCooldown(MajorCooldown{
		CooldownID: PotionCooldownID,
		Cooldown:   time.Minute * 2,
		Priority:   CooldownPriorityDefault,
		ActivationFactory: func(sim *Simulation) CooldownActivation {
			// Capture this inside ActivationFactory so it resets on Sim reset.
			numPotionsUsed := 0

			expectedManaPerUsage := float64((3000 + 1800) / 2)
			remainingUsages := int(1 + (MaxDuration(0, sim.Duration))/(time.Minute*2))

			remainingManaPotionUsages := 0
			if startingPotion == proto.Potions_SuperManaPotion {
				remainingManaPotionUsages += MinInt(numStartingPotions, remainingUsages)
			}
			if defaultPotion == proto.Potions_SuperManaPotion {
				remainingManaPotionUsages += MaxInt(0, remainingUsages-numStartingPotions)
			}

			character.ExpectedBonusMana += expectedManaPerUsage * float64(remainingManaPotionUsages)

			return func(sim *Simulation, character *Character) bool {
				usedPotion := false
				if startingPotionActivation != nil && numPotionsUsed < numStartingPotions {
					usedPotion = startingPotionActivation(sim, character)
				} else if defaultPotionActivation != nil {
					usedPotion = defaultPotionActivation(sim, character)
				}

				if usedPotion {
					numPotionsUsed++

					// Update expected bonus mana
					newRemainingUsages := int(sim.GetRemainingDuration() / (time.Minute * 2))
					newRemainingManaPotionUsages := 0
					if startingPotion == proto.Potions_SuperManaPotion {
						newRemainingManaPotionUsages += MinInt(numStartingPotions-numPotionsUsed, remainingUsages)
					}
					if defaultPotion == proto.Potions_SuperManaPotion {
						newRemainingManaPotionUsages += MaxInt(0, newRemainingUsages-MaxInt(0, numStartingPotions-numPotionsUsed))
					}

					character.ExpectedBonusMana -= expectedManaPerUsage * float64(remainingManaPotionUsages-newRemainingManaPotionUsages)
					remainingManaPotionUsages = newRemainingManaPotionUsages
				}

				return usedPotion
			}
		},
	})
}

const alchStoneItemID = 35749

func makePotionActivation(potionType proto.Potions, character *Character) CooldownActivation {
	if potionType == proto.Potions_DestructionPotion {
		return func(sim *Simulation, character *Character) bool {
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
			return true
		}
	} else if potionType == proto.Potions_SuperManaPotion {
		alchStoneEquipped := character.HasTrinketEquipped(alchStoneItemID)
		return func(sim *Simulation, character *Character) bool {
			// Only pop if we have less than the max mana provided by the potion minus 1mp5 tick.
			totalRegen := character.manaRegenPerSecondWhileCasting() * 5
			if character.MaxMana()-(character.CurrentMana()+totalRegen) < 3000 {
				return false
			}

			// Restores 1800 to 3000 mana. (2 Min Cooldown)
			manaGain := 1800 + (sim.RandomFloat("super mana") * 1200)

			if alchStoneEquipped {
				manaGain *= 1.4
			}

			character.AddStat(stats.Mana, manaGain)
			if sim.Log != nil {
				sim.Log("Used Super Mana Potion\n")
			}

			character.SetCD(PotionCooldownID, time.Minute*2+sim.CurrentTime)
			character.Metrics.AddInstantCast(ActionID{ItemID: 22832})
			return true
		}
	} else {
		return nil
	}
}

var RuneCooldownID = NewCooldownID()

func registerDarkRuneCD(agent Agent, consumes proto.Consumes) {
	if !consumes.DarkRune {
		return
	}

	character := agent.GetCharacter()
	character.AddMajorCooldown(MajorCooldown{
		CooldownID: RuneCooldownID,
		Cooldown:   time.Minute * 2,
		Priority:   CooldownPriorityDefault,
		ActivationFactory: func(sim *Simulation) CooldownActivation {
			expectedManaPerUsage := float64((900 + 600) / 2)
			remainingUsages := int(1 + (MaxDuration(0, sim.Duration))/(time.Minute*2))
			character.ExpectedBonusMana += expectedManaPerUsage * float64(remainingUsages)

			return func(sim *Simulation, character *Character) bool {
				// Only pop if we have less than the max mana provided by the potion minus 1mp5 tick.
				totalRegen := character.manaRegenPerSecondWhileCasting() * 5
				if character.MaxMana()-(character.CurrentMana()+totalRegen) < 1500 {
					return false
				}

				// Restores 900 to 1500 mana. (2 Min Cooldown)
				character.AddStat(stats.Mana, 900+(sim.RandomFloat("dark rune")*600))
				character.SetCD(RuneCooldownID, time.Minute*2+sim.CurrentTime)

				// Update expected bonus mana
				newRemainingUsages := int(sim.GetRemainingDuration() / (time.Minute * 2))
				character.ExpectedBonusMana -= expectedManaPerUsage * float64(remainingUsages-newRemainingUsages)
				remainingUsages = newRemainingUsages

				if sim.Log != nil {
					sim.Log("Used Dark Rune\n")
				}
				character.Metrics.AddInstantCast(ActionID{SpellID: 27869})

				return true
			}
		},
	})
}
