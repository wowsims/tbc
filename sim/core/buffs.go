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

	if individualBuffs.BattleShout != proto.TristateEffect_TristateEffectMissing {
		character.AddStat(stats.AttackPower, 306)
		if individualBuffs.BattleShout == proto.TristateEffect_TristateEffectImproved {
			character.AddStat(stats.AttackPower, 76.5) // assumes max rank of Commanding Presence
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
	registerPowerInfusionCD(agent, individualBuffs.PowerInfusions)
	registerManaTideTotemCD(agent, partyBuffs.ManaTideTotems)
	registerInnervateCD(agent, individualBuffs.Innervates)

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

// Used for approximating cooldowns applied by other players to you, such as
// bloodlust, innervate, power infusion, etc. This is specifically for buffs
// which can be consecutively applied multiple times to a single player.
type externalConsecutiveCDApproximation struct {
	AuraID           AuraID
	CooldownID       CooldownID
	CooldownPriority float64
	AuraDuration     time.Duration
	AuraCD           time.Duration

	// Callback for any special initialization.
	Init func(sim *Simulation, character *Character)

	// Applies the buff.
	AddAura func(sim *Simulation, character *Character) bool
}

// numSources is the number of other players assigned to apply the buff to this player.
// E.g. the number of other shaman in the group using bloodlust.
func registerExternalConsecutiveCDApproximation(agent Agent, config externalConsecutiveCDApproximation, numSources int32) {
	if numSources == 0 {
		return
	}

	externalCDs := make([]InternalCD, numSources)

	agent.GetCharacter().AddMajorCooldown(MajorCooldown{
		CooldownID: config.CooldownID,
		Cooldown:   config.AuraDuration, // Assumes that multiple buffs are different sources.
		Priority:   config.CooldownPriority,
		ActivationFactory: func(sim *Simulation) CooldownActivation {
			for i := 0; i < int(numSources); i++ {
				externalCDs[i] = NewICD()
			}
			nextExternalIndex := 0

			if config.Init != nil {
				config.Init(sim, agent.GetCharacter())
			}

			return func(sim *Simulation, character *Character) bool {
				if externalCDs[nextExternalIndex].IsOnCD(sim) {
					return false
				}

				if character.HasAura(config.AuraID) {
					return false
				}

				success := config.AddAura(sim, character)
				if !success {
					return false
				}

				externalCDs[nextExternalIndex] = InternalCD(sim.CurrentTime + config.AuraCD)
				nextExternalIndex = (nextExternalIndex + 1) % len(externalCDs)

				if externalCDs[nextExternalIndex].IsOnCD(sim) {
					character.SetCD(config.CooldownID, sim.CurrentTime+externalCDs[nextExternalIndex].GetRemainingCD(sim))
				} else {
					character.SetCD(config.CooldownID, sim.CurrentTime+config.AuraDuration)
				}
				return true
			}
		},
	})
}

var BloodlustAuraID = NewAuraID()
var sharedBloodlustCooldownID = NewCooldownID() // Different from shaman bloodlust CD.
const BloodlustDuration = time.Second * 40
const BloodlustCD = time.Minute * 10

func registerBloodlustCD(agent Agent, numBloodlusts int32) {
	registerExternalConsecutiveCDApproximation(
		agent,
		externalConsecutiveCDApproximation{
			AuraID:           BloodlustAuraID,
			CooldownID:       sharedBloodlustCooldownID,
			CooldownPriority: CooldownPriorityBloodlust,
			AuraDuration:     BloodlustDuration,
			AuraCD:           BloodlustCD,
			AddAura: func(sim *Simulation, character *Character) bool {
				AddBloodlustAura(sim, character)
				return true
			},
		},
		numBloodlusts)
}

func AddBloodlustAura(sim *Simulation, character *Character) {
	const bonus = 1.3
	const inverseBonus = 1 / bonus

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

var PowerInfusionAuraID = NewAuraID()
var sharedPowerInfusionCooldownID = NewCooldownID() // Different from priest PI CD.
const PowerInfusionDuration = time.Second * 15
const PowerInfusionCD = time.Minute * 3

func registerPowerInfusionCD(agent Agent, numPowerInfusions int32) {
	registerExternalConsecutiveCDApproximation(
		agent,
		externalConsecutiveCDApproximation{
			AuraID:           PowerInfusionAuraID,
			CooldownID:       sharedPowerInfusionCooldownID,
			CooldownPriority: CooldownPriorityDefault,
			AuraDuration:     PowerInfusionDuration,
			AuraCD:           PowerInfusionCD,
			AddAura: func(sim *Simulation, character *Character) bool {
				AddPowerInfusionAura(sim, character)
				return true
			},
		},
		numPowerInfusions)
}

func AddPowerInfusionAura(sim *Simulation, character *Character) {
	const bonus = 1.2
	const inverseBonus = 1 / bonus

	character.PseudoStats.CastSpeedMultiplier *= bonus

	character.AddAura(sim, Aura{
		ID:      PowerInfusionAuraID,
		SpellID: 10060,
		Name:    "Power Infusion",
		Expires: sim.CurrentTime + PowerInfusionDuration,
		OnCast: func(sim *Simulation, cast *Cast) {
			// TODO: Double-check this is how the calculation works.
			cast.ManaCost = MaxFloat(0, cast.ManaCost-cast.BaseManaCost*0.2)
		},
		OnExpire: func(sim *Simulation) {
			character.PseudoStats.CastSpeedMultiplier *= inverseBonus
		},
	})
}

var sharedInnervateCooldownID = NewCooldownID()
var InnervateAuraID = NewAuraID()

const InnervateDuration = time.Second * 20
const InnervateCD = time.Minute * 6

func InnervateManaThreshold(character *Character) float64 {
	if character.Class == proto.Class_ClassMage {
		// Mages burn mana really fast so they probably need a higher threshold.
		return 2000
	} else {
		return 1000
	}
}

func registerInnervateCD(agent Agent, numInnervates int32) {
	innervateThreshold := InnervateManaThreshold(agent.GetCharacter())
	expectedManaPerInnervate := 0.0
	remainingInnervateUsages := 0

	registerExternalConsecutiveCDApproximation(
		agent,
		externalConsecutiveCDApproximation{
			AuraID:           InnervateAuraID,
			CooldownID:       sharedInnervateCooldownID,
			CooldownPriority: CooldownPriorityDefault,
			AuraDuration:     InnervateDuration,
			AuraCD:           InnervateCD,
			Init: func(sim *Simulation, character *Character) {
				expectedManaPerInnervate = character.SpiritManaRegenPerSecond() * 5 * 20
				remainingInnervateUsages = int(1 + (MaxDuration(0, sim.Duration))/InnervateCD)
				character.ExpectedBonusMana += expectedManaPerInnervate * float64(remainingInnervateUsages)
			},
			AddAura: func(sim *Simulation, character *Character) bool {
				// Only cast innervate when very low on mana, to make sure all other mana CDs are prioritized.
				if character.CurrentMana() > innervateThreshold {
					return false
				}

				AddInnervateAura(sim, character, expectedManaPerInnervate)

				newRemainingUsages := int(sim.GetRemainingDuration() / InnervateCD)
				// AddInnervateAura already accounts for 1 usage, which is why we subtract 1 less.
				character.ExpectedBonusMana -= expectedManaPerInnervate * MaxFloat(0, float64(remainingInnervateUsages-newRemainingUsages-1))
				remainingInnervateUsages = newRemainingUsages

				return true
			},
		},
		numInnervates)
}

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

var sharedManaTideTotemCooldownID = NewCooldownID()
var ManaTideTotemAuraID = NewAuraID()

const ManaTideTotemDuration = time.Second * 12
const ManaTideTotemCD = time.Minute * 5

func ManaTideTotemAmount(character *Character) float64 {
	// Subtract 120 mana to simulate the loss of mana spring while MTT is active.
	// This isn't correct for multi-resto shaman groups, but that isnt a common case.
	return character.MaxMana()*0.24 - 120
}

func registerManaTideTotemCD(agent Agent, numManaTideTotems int32) {
	expectedManaPerManaTideTotem := 0.0
	remainingManaTideTotemUsages := 0
	initialDelay := time.Duration(0)

	registerExternalConsecutiveCDApproximation(
		agent,
		externalConsecutiveCDApproximation{
			AuraID:           ManaTideTotemAuraID,
			CooldownID:       sharedManaTideTotemCooldownID,
			CooldownPriority: CooldownPriorityDefault,
			AuraDuration:     ManaTideTotemDuration,
			AuraCD:           ManaTideTotemCD,
			Init: func(sim *Simulation, character *Character) {
				// Use first MTT at 60s, or halfway through the fight, whichever comes first.
				initialDelay = MinDuration(sim.Duration/2, time.Second*60)

				expectedManaPerManaTideTotem = ManaTideTotemAmount(character)
				remainingManaTideTotemUsages = int(1 + MaxDuration(0, sim.Duration-initialDelay)/ManaTideTotemCD)
				character.ExpectedBonusMana += expectedManaPerManaTideTotem * float64(remainingManaTideTotemUsages)
			},
			AddAura: func(sim *Simulation, character *Character) bool {
				// A normal resto shaman would wait to use MTT.
				if sim.CurrentTime < initialDelay {
					return false
				}

				AddManaTideTotemAura(sim, character)

				newRemainingUsages := int(sim.GetRemainingDuration() / ManaTideTotemCD)
				// AddManaTideTotemAura already accounts for 1 usage, which is why we subtract 1 less.
				character.ExpectedBonusMana -= expectedManaPerManaTideTotem * MaxFloat(0, float64(remainingManaTideTotemUsages-newRemainingUsages-1))
				remainingManaTideTotemUsages = newRemainingUsages

				return true
			},
		},
		numManaTideTotems)
}

func AddManaTideTotemAura(sim *Simulation, character *Character) {
	lastUpdateTime := sim.CurrentTime
	totalBonusMana := ManaTideTotemAmount(character)
	bonusManaSubtracted := 0.0

	character.AddAura(sim, Aura{
		ID:      ManaTideTotemAuraID,
		SpellID: 16190,
		Name:    "Mana Tide Totem",
		Expires: sim.CurrentTime + ManaTideTotemDuration,
		OnCast: func(sim *Simulation, cast *Cast) {
			timeDelta := sim.CurrentTime - lastUpdateTime
			lastUpdateTime = sim.CurrentTime
			progressRatio := float64(timeDelta) / float64(ManaTideTotemDuration)
			remainder := totalBonusMana - bonusManaSubtracted
			amount := MinFloat(remainder, totalBonusMana*progressRatio)

			character.AddMana(sim, amount, "Mana Tide Totem", true)
			character.ExpectedBonusMana -= amount
			bonusManaSubtracted += amount
		},
		OnExpire: func(sim *Simulation) {
			remainder := totalBonusMana - bonusManaSubtracted
			character.AddMana(sim, remainder, "Mana Tide Totem", true)
			character.ExpectedBonusMana -= remainder
		},
	})
}
