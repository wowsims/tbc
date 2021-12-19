package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Applies buffs that affect individual players.
func applyBuffEffects(agent Agent, raidBuffs proto.RaidBuffs, partyBuffs proto.PartyBuffs, individualBuffs proto.IndividualBuffs) {
	character := agent.GetCharacter()

	if raidBuffs.ArcaneBrilliance {
		character.AddStats(stats.Stats{
			stats.Intellect: 40,
		})
	}

	gotwAmount := GetTristateValueFloat(raidBuffs.GiftOfTheWild, 14.0, 14.0*1.35)
	character.AddStats(stats.Stats{
		stats.Stamina:   gotwAmount,
		stats.Agility:   gotwAmount,
		stats.Strength:  gotwAmount,
		stats.Intellect: gotwAmount,
		stats.Spirit:    gotwAmount,
	})

	character.AddStats(stats.Stats{
		stats.SpellCrit: GetTristateValueFloat(partyBuffs.MoonkinAura, 5*SpellCritRatingPerCritChance, 5*SpellCritRatingPerCritChance+20),
	})

	if partyBuffs.DraeneiRacialMelee {
		character.AddStats(stats.Stats{
			stats.MeleeHit: 1 * MeleeHitRatingPerHitChance,
		})
	}

	if partyBuffs.DraeneiRacialCaster {
		character.AddStats(stats.Stats{
			stats.SpellHit: 1 * SpellHitRatingPerHitChance,
		})
	}
	character.AddStats(stats.Stats{
		stats.Spirit: GetTristateValueFloat(raidBuffs.DivineSpirit, 50.0, 50.0),
	})
	if raidBuffs.DivineSpirit == proto.TristateEffect_TristateEffectImproved {
		character.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.Spirit,
			ModifiedStat: stats.SpellPower,
			Modifier: func(spirit float64, spellPower float64) float64 {
				return spellPower + spirit*0.1
			},
		})
	}

	// shadow priest buff bot just statically applies mp5
	if individualBuffs.ShadowPriestDps > 0 {
		character.AddStats(stats.Stats{
			stats.MP5: float64(individualBuffs.ShadowPriestDps) * 0.25,
		})
	}

	// TODO: Double-check these numbers
	character.AddStats(stats.Stats{
		stats.MP5: GetTristateValueFloat(individualBuffs.BlessingOfWisdom, 42.0, 50.0),
	})

	if individualBuffs.BlessingOfKings {
		bokStats := [5]stats.Stat{
			stats.Agility,
			stats.Strength,
			stats.Stamina,
			stats.Intellect,
			stats.Spirit,
		}

		for _, stat := range bokStats {
			character.AddStatDependency(stats.StatDependency{
				SourceStat:   stat,
				ModifiedStat: stat,
				Modifier: func(curValue float64, _ float64) float64 {
					return curValue * 1.1
				},
			})
		}
	}

	if partyBuffs.TotemOfWrath > 0 {
		character.AddStats(stats.Stats{
			stats.SpellCrit: 3 * SpellCritRatingPerCritChance * float64(partyBuffs.TotemOfWrath),
			stats.SpellHit:  3 * SpellHitRatingPerHitChance * float64(partyBuffs.TotemOfWrath),
		})
	}
	character.AddStats(stats.Stats{
		stats.SpellPower: GetTristateValueFloat(partyBuffs.WrathOfAirTotem, 101.0, 121.0),
	})
	character.AddStats(stats.Stats{
		stats.Agility: GetTristateValueFloat(partyBuffs.GraceOfAirTotem, 77.0, 88.55),
	})
	character.AddStats(stats.Stats{
		stats.Strength: GetTristateValueFloat(partyBuffs.StrengthOfEarthTotem, 86.0, 98.9),
	})
	character.AddStats(stats.Stats{
		stats.MP5: GetTristateValueFloat(partyBuffs.ManaSpringTotem, 50, 62.5),
	})

	registerBloodlustCD(agent, partyBuffs.Bloodlust)

	character.AddStats(stats.Stats{
		stats.SpellCrit: 28 * float64(partyBuffs.AtieshMage),
	})
	character.AddStats(stats.Stats{
		stats.SpellPower:   33 * float64(partyBuffs.AtieshWarlock),
		stats.HealingPower: 33 * float64(partyBuffs.AtieshWarlock),
	})

	if partyBuffs.BraidedEterniumChain {
		character.AddStats(stats.Stats{stats.MeleeCrit: 28})
	}
	if partyBuffs.EyeOfTheNight {
		character.AddStats(stats.Stats{stats.SpellPower: 34})
	}
	if partyBuffs.JadePendantOfBlasting {
		character.AddStats(stats.Stats{stats.SpellPower: 15})
	}
	if partyBuffs.ChainOfTheTwilightOwl {
		character.AddStats(stats.Stats{stats.SpellCrit: 2 * SpellCritRatingPerCritChance})
	}
}

var BloodlustAuraID = NewAuraID()
var sharedBloodlustCooldownID = NewCooldownID() // Different from shaman bloodlust CD.
const BloodlustDuration = time.Second * 40
const BloodlustCD = time.Minute * 10

func registerBloodlustCD(agent Agent, numBloodlusts int32) {
	if numBloodlusts == 0 {
		return
	}

	bloodlustCDs := make([]InternalCD, numBloodlusts)

	agent.GetCharacter().AddMajorCooldown(MajorCooldown{
		CooldownID: sharedBloodlustCooldownID,
		Cooldown:   BloodlustDuration, // assumes that multiple BLs are different shaman.
		Priority:   CooldownPriorityBloodlust,
		ActivationFactory: func(sim *Simulation) CooldownActivation {
			for i := 0; i < int(numBloodlusts); i++ {
				bloodlustCDs[i] = NewICD()
			}
			nextBloodlustIndex := 0

			return func(sim *Simulation, character *Character) bool {
				if bloodlustCDs[nextBloodlustIndex].IsOnCD(sim) {
					return false
				}

				if character.HasAura(BloodlustAuraID) {
					return false
				}

				AddBloodlustAura(sim, character)
				bloodlustCDs[nextBloodlustIndex] = InternalCD(sim.CurrentTime + BloodlustCD)
				nextBloodlustIndex = (nextBloodlustIndex + 1) % len(bloodlustCDs)

				if bloodlustCDs[nextBloodlustIndex].IsOnCD(sim) {
					character.SetCD(sharedBloodlustCooldownID, sim.CurrentTime+bloodlustCDs[nextBloodlustIndex].GetRemainingCD(sim))
				} else {
					character.SetCD(sharedBloodlustCooldownID, sim.CurrentTime+BloodlustDuration)
				}
				return true
			}
		},
	})
}

func AddBloodlustAura(sim *Simulation, character *Character) {
	const bonus = 1.3
	const inverseBonus = 1 / 1.3

	character.PseudoStats.CastSpeedMultiplier *= bonus
	character.MultiplyMeleeSpeed(sim, bonus)

	character.AddAura(sim, Aura{
		ID:      BloodlustAuraID,
		SpellID: 2825,
		Name:    "Bloodlust",
		Expires: sim.CurrentTime + BloodlustDuration,
		OnExpire: func(sim *Simulation) {
			character.PseudoStats.CastSpeedMultiplier *= inverseBonus
			character.MultiplyMeleeSpeed(sim, inverseBonus)
		},
	})
}

func registerInnervateCD(agent Agent, numInnervates int) {
	if numInnervates == 0 {
		return
	}

	// Cooldowns for each innervate are separate, since they are cast by different players.
	innervateCDs := make([]InternalCD, numInnervates)
	const dur = time.Second * 20
	const cd = time.Minute * 6

	agent.GetCharacter().AddMajorCooldown(MajorCooldown{
		CooldownID: InnervateCooldownID,
		Cooldown:   dur, // Just put on CD for the duration because we can get other innervates after
		ActivationFactory: func(sim *Simulation) CooldownActivation {
			for i := 0; i < numInnervates; i++ {
				innervateCDs[i] = NewICD()
			}
			nextInnervateIndex := 0

			return func(sim *Simulation, character *Character) bool {
				if innervateCDs[nextInnervateIndex].IsOnCD(sim) {
					return false
				}

				if character.HasAura(InnervateAuraID) {
					return false
				}

				// Only cast innervate when very low on mana, to make sure all other mana CDs are prioritized.
				if character.CurrentMana() > 1000 {
					return false
				}

				AddInnervateAura(sim, character, 0)
				innervateCDs[nextInnervateIndex] = InternalCD(sim.CurrentTime + cd)
				nextInnervateIndex = (nextInnervateIndex + 1) % len(innervateCDs)

				if innervateCDs[nextInnervateIndex].IsOnCD(sim) {
					character.SetCD(InnervateCooldownID, sim.CurrentTime+innervateCDs[nextInnervateIndex].GetRemainingCD(sim))
				} else {
					character.SetCD(InnervateCooldownID, sim.CurrentTime+dur)
				}
				return true
			}
		},
	})
}

var InnervateCooldownID = NewCooldownID()
var InnervateAuraID = NewAuraID()

const InnervateDuration = time.Second * 20

func AddInnervateAura(sim *Simulation, character *Character, expectedBonusManaReduction float64) {
	character.PseudoStats.ForceFullSpiritRegen = true
	character.PseudoStats.SpiritRegenMultiplier *= 5.0

	lastUpdateTime := sim.CurrentTime
	bonusManaSubtracted := 0.0

	character.AddAura(sim, Aura{
		ID:      InnervateAuraID,
		SpellID: 29166,
		Name:    "Innervate",
		Expires: sim.CurrentTime + InnervateDuration,
		OnCast: func(sim *Simulation, cast *Cast) {
			timeDelta := sim.CurrentTime - lastUpdateTime
			lastUpdateTime = sim.CurrentTime
			progressRatio := float64(timeDelta) / float64(InnervateDuration)
			amount := expectedBonusManaReduction * progressRatio

			character.ExpectedBonusMana -= amount
			character.Metrics.BonusManaGained += amount
			bonusManaSubtracted += amount
		},
		OnExpire: func(sim *Simulation) {
			character.PseudoStats.ForceFullSpiritRegen = false
			character.PseudoStats.SpiritRegenMultiplier /= 5.0

			remainder := expectedBonusManaReduction - bonusManaSubtracted
			character.ExpectedBonusMana -= remainder
			character.Metrics.BonusManaGained += remainder
		},
	})
}
