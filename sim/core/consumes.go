package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Registers all consume-related effects to the Agent.
func applyConsumeEffects(agent Agent) {
	consumes := agent.GetCharacter().consumes

	agent.GetCharacter().AddStats(consumesStats(consumes))

	registerDrumsCD(agent, consumes)
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

	return s
}

// Adds drums as a major cooldown to the character, if it's being used.
func registerDrumsCD(agent Agent, consumes proto.Consumes) {
	character := agent.GetCharacter()
	drumsType := proto.Drums_DrumsUnknown

	// Whether this agent is the one casting the drums.
	//drumsSelfCast := false

	if consumes.Drums != proto.Drums_DrumsUnknown {
		drumsType = consumes.Drums
		//drumsSelfCast = true
	} else if character.Party.buffs.Drums != proto.Drums_DrumsUnknown {
		drumsType = character.Party.buffs.Drums
	}

	// TODO: If drumsSelfCast == true, then do a cast time
	mcd := MajorCooldown{
		CooldownID: MagicIDDrums,
		Cooldown: time.Minute * 2,
		Priority: CooldownPriorityDrums,
	}

	if drumsType == proto.Drums_DrumsOfBattle {
		mcd.ActivationFactory = func(sim *Simulation) CooldownActivation {
			return func(sim *Simulation, character *Character) bool {
				const hasteBonus = 80
				for _, agent := range character.Party.Players {
					agent.GetCharacter().SetCD(MagicIDDrums, time.Minute*2+sim.CurrentTime) // tinnitus
					agent.GetCharacter().AddAuraWithTemporaryStats(sim, MagicIDDrums, "Drums of Battle", stats.SpellHaste, hasteBonus, time.Second*30)
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
					agent.GetCharacter().SetCD(MagicIDDrums, time.Minute*2+sim.CurrentTime) // tinnitus
					agent.GetCharacter().AddAuraWithTemporaryStats(sim, MagicIDDrums, "Drums of Restoration", stats.MP5, mp5Bonus, time.Second*15)
				}
				return true
			}
		}
	}

	if mcd.ActivationFactory != nil {
		agent.GetCharacter().AddMajorCooldown(mcd);
	}
}

func registerPotionCD(agent Agent, consumes proto.Consumes) {
	character := agent.GetCharacter()
	defaultPotionActivation := makePotionActivation(consumes.DefaultPotion, character)
	startingPotionActivation := makePotionActivation(consumes.StartingPotion, character)
	numStartingPotions := consumes.NumStartingPotions

	mcd := MajorCooldown{
		CooldownID: MagicIDPotion,
		Cooldown: time.Minute * 2,
		Priority: CooldownPriorityDefault,
	}

	if defaultPotionActivation != nil {
		if startingPotionActivation != nil && numStartingPotions > 0 {
			mcd.ActivationFactory = func(sim *Simulation) CooldownActivation {
				// Capture this inside ActivationFactory so it resets on Sim reset.
				numPotionsUsed := int32(0)

				return func(sim *Simulation, character *Character) bool {
					usedPotion := false
					if numPotionsUsed < numStartingPotions {
						usedPotion = startingPotionActivation(sim, character)
					} else {
						usedPotion = defaultPotionActivation(sim, character)
					}
					if usedPotion {
						numPotionsUsed++
					}
					return usedPotion
				}
			}
		} else {
			mcd.ActivationFactory = func(sim *Simulation) CooldownActivation {
				return defaultPotionActivation
			}
		}
	} else if startingPotionActivation != nil && numStartingPotions > 0 {
		mcd.ActivationFactory = func(sim *Simulation) CooldownActivation {
			// Capture this inside ActivationFactory so it resets on Sim reset.
			numPotionsUsed := int32(0)

			return func(sim *Simulation, character *Character) bool {
				usedPotion := false
				if numPotionsUsed < numStartingPotions {
					usedPotion = startingPotionActivation(sim, character)
					if usedPotion {
						numPotionsUsed++
					}
					return usedPotion
				} else {
					character.SetCD(MagicIDPotion, NeverExpires)
					return true
				}
			}
		}
	}

	if mcd.ActivationFactory != nil {
		agent.GetCharacter().AddMajorCooldown(mcd)
	}
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
				ID:      MagicIDDestructionPotion,
				Name:    "Destruction Potion",
				Expires: sim.CurrentTime + dur,
				OnExpire: func(sim *Simulation) {
					character.AddStat(stats.SpellPower, -spBonus)
					character.AddStat(stats.SpellCrit, -critBonus)
				},
			})

			character.SetCD(MagicIDPotion, time.Minute*2+sim.CurrentTime)
			return true
		}
	} else if potionType == proto.Potions_SuperManaPotion {
		alchStoneEquipped := character.HasTrinketEquipped(alchStoneItemID)
		return func(sim *Simulation, character *Character) bool {
			// Only pop if we have less than the max mana provided by the potion minus 1mp5 tick.
			totalRegen := character.manaRegenPerSecond() * 5
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
				sim.Log("Used Mana Potion\n")
			}

			character.SetCD(MagicIDPotion, time.Minute*2+sim.CurrentTime)
			return true
		}
	} else {
		return nil
	}
}

func registerDarkRuneCD(agent Agent, consumes proto.Consumes) {
	if !consumes.DarkRune {
		return
	}

	agent.GetCharacter().AddMajorCooldown(MajorCooldown{
		CooldownID: MagicIDRune,
		Cooldown: time.Minute * 2,
		Priority: CooldownPriorityDefault,
		ActivationFactory: func(sim *Simulation) CooldownActivation {
			return func(sim *Simulation, character *Character) bool {
				// Only pop if we have less than the max mana provided by the potion minus 1mp5 tick.
				totalRegen := character.manaRegenPerSecond() * 5
				if character.MaxMana()-(character.CurrentMana()+totalRegen) < 1500 {
					return false
				}

				// Restores 900 to 1500 mana. (2 Min Cooldown)
				character.AddStat(stats.Mana, 900 + (sim.RandomFloat("dark rune") * 600))
				character.SetCD(MagicIDRune, time.Minute*2+sim.CurrentTime)
				if sim.Log != nil {
					sim.Log("Used Dark Rune\n")
				}

				return true
			}
		},
	})
}
