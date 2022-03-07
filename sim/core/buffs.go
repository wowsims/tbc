package core

import (
	"math"
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
	character.AddStats(stats.Stats{
		stats.MeleeCrit: GetTristateValueFloat(partyBuffs.LeaderOfThePack, 5*MeleeCritRatingPerCritChance, 5*MeleeCritRatingPerCritChance+20),
	})

	if partyBuffs.TrueshotAura {
		character.AddStats(stats.Stats{
			stats.AttackPower:       125,
			stats.RangedAttackPower: 125,
		})
	}

	if partyBuffs.FerociousInspiration > 0 {
		character.AddPermanentAura(func(sim *Simulation) Aura {
			return FerociousInspirationAura(partyBuffs.FerociousInspiration)
		})
	}

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

	if individualBuffs.ShadowPriestDps > 0 {
		character.AddStats(stats.Stats{
			stats.MP5: float64(individualBuffs.ShadowPriestDps) * 0.25,
		})
	}

	character.AddStats(stats.Stats{
		stats.MP5: GetTristateValueFloat(individualBuffs.BlessingOfWisdom, 42.0, 50.0),
	})

	character.AddStats(stats.Stats{
		stats.AttackPower:       GetTristateValueFloat(individualBuffs.BlessingOfMight, 220, 264),
		stats.RangedAttackPower: GetTristateValueFloat(individualBuffs.BlessingOfMight, 220, 264),
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

	if individualBuffs.BlessingOfSalvation {
		character.PseudoStats.ThreatMultiplier *= 0.7
	}

	if partyBuffs.SanctityAura == proto.TristateEffect_TristateEffectImproved {
		character.AddPermanentAura(func(sim *Simulation) Aura {
			return ImprovedSanctityAura(sim, 2)
		})
	}

	if partyBuffs.BattleShout != proto.TristateEffect_TristateEffectMissing {
		character.AddStats(stats.Stats{
			stats.AttackPower: GetTristateValueFloat(partyBuffs.BattleShout, 306, 382.5),
		})
		if partyBuffs.BsSolarianSapphire {
			partyBuffs.SnapshotBsSolarianSapphire = false
			character.AddStats(stats.Stats{
				stats.AttackPower: 70,
			})
		}

		snapshotSapphire := partyBuffs.SnapshotBsSolarianSapphire
		snapshotT2 := partyBuffs.SnapshotBsT2
		if snapshotSapphire || snapshotT2 {
			character.AddPermanentAuraWithOptions(PermanentAura{
				AuraFactory:       SnapshotBattleShoutAura(character, snapshotSapphire, snapshotT2),
				RespectExpiration: true,
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
	if partyBuffs.WrathOfAirTotem == proto.TristateEffect_TristateEffectRegular && partyBuffs.SnapshotImprovedWrathOfAirTotem {
		character.AddPermanentAuraWithOptions(PermanentAura{
			AuraFactory:       SnapshotImprovedWrathOfAirTotemAura(character),
			RespectExpiration: true,
		})
	}
	character.AddStats(stats.Stats{
		stats.Agility: GetTristateValueFloat(partyBuffs.GraceOfAirTotem, 77.0, 88.55),
	})
	switch partyBuffs.StrengthOfEarthTotem {
	case proto.StrengthOfEarthType_Basic:
		character.AddStat(stats.Strength, 86)
	case proto.StrengthOfEarthType_CycloneBonus:
		character.AddStat(stats.Strength, 98)
	case proto.StrengthOfEarthType_EnhancingTotems:
		character.AddStat(stats.Strength, 98.9)
	case proto.StrengthOfEarthType_EnhancingAndCyclone:
		character.AddStat(stats.Strength, 110.9)
	}
	if (partyBuffs.StrengthOfEarthTotem == proto.StrengthOfEarthType_Basic || partyBuffs.StrengthOfEarthTotem == proto.StrengthOfEarthType_EnhancingTotems) && partyBuffs.SnapshotImprovedStrengthOfEarthTotem {
		character.AddPermanentAuraWithOptions(PermanentAura{
			AuraFactory:       SnapshotImprovedStrengthOfEarthTotemAura(character),
			RespectExpiration: true,
		})
	}
	character.AddStats(stats.Stats{
		stats.MP5: GetTristateValueFloat(partyBuffs.ManaSpringTotem, 50, 62.5),
	})
	if partyBuffs.WindfuryTotemRank > 0 && IsEligibleForWindfuryTotem(character) {
		character.HasWFTotem = true
		character.AddPermanentAura(func(sim *Simulation) Aura {
			return WindfuryTotemAura(character, partyBuffs.WindfuryTotemRank, partyBuffs.WindfuryTotemIwt)
		})
	}
	if partyBuffs.TranquilAirTotem {
		character.PseudoStats.ThreatMultiplier *= 0.8
	}

	if individualBuffs.UnleashedRage {
		character.AddStatDependency(stats.StatDependency{
			SourceStat:   stats.AttackPower,
			ModifiedStat: stats.AttackPower,
			Modifier: func(ap float64, _ float64) float64 {
				return ap * 1.1
			},
		})
	}

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
	if partyBuffs.BattleChickens > 0 {
		character.AddPermanentAuraWithOptions(PermanentAura{
			AuraFactory:       BattleChickenAura(character, partyBuffs.BattleChickens),
			RespectExpiration: true,
		})
	}
}

// Applies buffs to pets.
func applyPetBuffEffects(petAgent PetAgent, raidBuffs proto.RaidBuffs, partyBuffs proto.PartyBuffs, individualBuffs proto.IndividualBuffs) {
	// Summoned pets, like Mage Water Elemental, aren't around to receive raid buffs.
	if !petAgent.GetPet().initialEnabled {
		return
	}

	// We need to modify the buffs a bit because some things are applied to pets by
	// the owner during combat (Bloodlust) or don't make sense for a pet.
	partyBuffs.Bloodlust = 0
	partyBuffs.Drums = proto.Drums_DrumsUnknown
	individualBuffs.Innervates = 0
	individualBuffs.PowerInfusions = 0

	// For some reason pets don't benefit from buffs that are ratings, e.g. crit rating or haste rating.
	partyBuffs.LeaderOfThePack = MinTristate(partyBuffs.LeaderOfThePack, proto.TristateEffect_TristateEffectRegular)
	partyBuffs.BraidedEterniumChain = false

	applyBuffEffects(petAgent, raidBuffs, partyBuffs, individualBuffs)
}

var SnapshotImprovedStrengthOfEarthTotemAuraID = NewAuraID()

func SnapshotImprovedStrengthOfEarthTotemAura(character *Character) AuraFactory {
	return character.NewTemporaryStatsAuraFactory(SnapshotImprovedStrengthOfEarthTotemAuraID, ActionID{SpellID: 37223}, stats.Stats{stats.Strength: 12}, time.Second*110)
}

var SnapshotImprovedWrathOfAirTotemAuraID = NewAuraID()

func SnapshotImprovedWrathOfAirTotemAura(character *Character) AuraFactory {
	return character.NewTemporaryStatsAuraFactory(SnapshotImprovedWrathOfAirTotemAuraID, ActionID{SpellID: 37212}, stats.Stats{stats.SpellPower: 20}, time.Second*110)
}

var SnapshotBattleShoutAuraID = NewAuraID()

func SnapshotBattleShoutAura(character *Character, snapshotSapphire bool, snapshotT2 bool) AuraFactory {
	amount := 0.0
	if snapshotSapphire {
		amount += 70
	}
	if snapshotT2 {
		amount += 30
	}

	return character.NewTemporaryStatsAuraFactory(SnapshotBattleShoutAuraID, ActionID{SpellID: 2048, Tag: 1}, stats.Stats{stats.AttackPower: amount}, time.Second*110)
}

var BattleChickenAuraID = NewAuraID()

func BattleChickenAura(character *Character, numChickens int32) AuraFactory {
	bonus := math.Pow(1.05, float64(numChickens))
	inverseBonus := 1 / bonus

	return func(sim *Simulation) Aura {
		character.MultiplyMeleeSpeed(sim, bonus)

		return Aura{
			ID:      BattleChickenAuraID,
			Expires: sim.CurrentTime + time.Minute*4,
			OnExpire: func(sim *Simulation) {
				character.MultiplyMeleeSpeed(sim, inverseBonus)
			},
		}
	}
}

var FerociousInspirationAuraID = NewAuraID()

func FerociousInspirationAura(numBMHunters int32) Aura {
	multiplier := math.Pow(1.03, float64(numBMHunters))
	return Aura{
		ID:       FerociousInspirationAuraID,
		ActionID: ActionID{SpellID: 34460, Tag: -1},
		OnBeforeSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellHitEffect) {
			spellEffect.DamageMultiplier *= multiplier
		},
		OnBeforePeriodicDamage: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect, tickDamage *float64) {
			*tickDamage *= multiplier
		},
	}
}

var ImprovedSanctityAuraID = NewAuraID()

func ImprovedSanctityAura(sim *Simulation, level float64) Aura {
	return Aura{
		ID:       ImprovedSanctityAuraID,
		ActionID: ActionID{SpellID: 31870},
		OnBeforeSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellHitEffect) {
			// unsure if this scaling should be additive or multiplicative
			// scale 10% for holy damange
			if spellCast.SpellSchool.Matches(SpellSchoolHoly) {
				spellEffect.DamageMultiplier *= 1.1
			}
			// scale additional 2% for all damage
			spellEffect.DamageMultiplier *= 1 + 0.01*level
		},
		OnBeforePeriodicDamage: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect, tickDamage *float64) {
			if spellCast.SpellSchool.Matches(SpellSchoolHoly) {
				*tickDamage *= 1.1
			}
			*tickDamage *= 1 + 0.01*level
		},
	}
}

var WindfuryTotemAuraID = NewAuraID()

var WindfurySpellRanks = []int32{
	8512,
	10613,
	10614,
	25585,
	25587,
}

var windfuryAPBonuses = []float64{
	122,
	229,
	315,
	375,
	445,
}

func IsEligibleForWindfuryTotem(character *Character) bool {
	return character.AutoAttacks.IsEnabled() &&
		character.HasMHWeapon() &&
		!character.HasMHWeaponImbue
}

func WindfuryTotemAura(character *Character, rank int32, iwtTalentPoints int32) Aura {
	spellID := WindfurySpellRanks[rank-1]
	actionID := ActionID{SpellID: spellID}
	apBonus := windfuryAPBonuses[rank-1]
	apBonus *= 1 + 0.15*float64(iwtTalentPoints)

	mhAttack := character.AutoAttacks.MHAuto
	mhAttack.ActionID = actionID
	mhAttack.Effect.BonusAttackPower += apBonus
	cachedAttack := SimpleSpell{}

	const procChance = 0.2

	var icd InternalCD
	const icdDur = time.Duration(1) // No ICD, but only once per frame.

	return Aura{
		ID:       WindfuryTotemAuraID,
		ActionID: actionID,
		OnSpellHit: func(sim *Simulation, spellCast *SpellCast, spellEffect *SpellEffect) {
			if !spellEffect.Landed() ||
				!spellEffect.ProcMask.Matches(ProcMaskMeleeMHAuto) ||
				spellCast.IsPhantom {
				return
			}
			if icd.IsOnCD(sim) {
				return
			}
			if sim.RandomFloat("Windfury Totem") > procChance {
				return
			}
			icd = InternalCD(sim.CurrentTime + icdDur)

			cachedAttack = mhAttack
			cachedAttack.Effect.Target = spellEffect.Target
			cachedAttack.Cast(sim)
		},
	}
}

// Used for approximating cooldowns applied by other players to you, such as
// bloodlust, innervate, power infusion, etc. This is specifically for buffs
// which can be consecutively applied multiple times to a single player.
type externalConsecutiveCDApproximation struct {
	ActionID         ActionID
	AuraID           AuraID
	CooldownID       CooldownID
	CooldownPriority float64
	Type             int32
	AuraDuration     time.Duration
	AuraCD           time.Duration

	// Callback for any special initialization.
	Init func(sim *Simulation, character *Character)

	// Callback for extra activation conditions.
	ShouldActivate CooldownActivationCondition

	// Applies the buff.
	AddAura CooldownActivation
}

// numSources is the number of other players assigned to apply the buff to this player.
// E.g. the number of other shaman in the group using bloodlust.
func registerExternalConsecutiveCDApproximation(agent Agent, config externalConsecutiveCDApproximation, numSources int32) {
	if numSources == 0 {
		return
	}

	externalCDs := make([]InternalCD, numSources)
	nextExternalIndex := 0

	agent.GetCharacter().AddMajorCooldown(MajorCooldown{
		ActionID:   config.ActionID,
		CooldownID: config.CooldownID,
		Cooldown:   config.AuraDuration, // Assumes that multiple buffs are different sources.
		Priority:   config.CooldownPriority,
		Type:       config.Type,

		CanActivate: func(sim *Simulation, character *Character) bool {
			if externalCDs[nextExternalIndex].IsOnCD(sim) {
				return false
			}

			if character.HasAura(config.AuraID) {
				return false
			}

			return true
		},
		ShouldActivate: config.ShouldActivate,

		ActivationFactory: func(sim *Simulation) CooldownActivation {
			for i := 0; i < int(numSources); i++ {
				externalCDs[i] = NewICD()
			}
			nextExternalIndex = 0

			if config.Init != nil {
				config.Init(sim, agent.GetCharacter())
			}

			return func(sim *Simulation, character *Character) {
				config.AddAura(sim, character)

				externalCDs[nextExternalIndex] = InternalCD(sim.CurrentTime + config.AuraCD)
				nextExternalIndex = (nextExternalIndex + 1) % len(externalCDs)

				if externalCDs[nextExternalIndex].IsOnCD(sim) {
					character.SetCD(config.CooldownID, sim.CurrentTime+externalCDs[nextExternalIndex].GetRemainingCD(sim))
				} else {
					character.SetCD(config.CooldownID, sim.CurrentTime+config.AuraDuration)
				}
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
			ActionID:         ActionID{SpellID: 2825, Tag: -1},
			AuraID:           BloodlustAuraID,
			CooldownID:       sharedBloodlustCooldownID,
			CooldownPriority: CooldownPriorityBloodlust,
			AuraDuration:     BloodlustDuration,
			AuraCD:           BloodlustCD,
			Type:             CooldownTypeDPS,

			ShouldActivate: func(sim *Simulation, character *Character) bool {
				// Haste portion doesn't stack with Power Infusion, so prefer to wait.
				if character.HasAura(PowerInfusionAuraID) {
					return false
				}
				return true
			},
			AddAura: func(sim *Simulation, character *Character) { AddBloodlustAura(sim, character, -1) },
		},
		numBloodlusts)
}

func AddBloodlustAura(sim *Simulation, character *Character, actionTag int32) {
	const bonus = 1.3
	const inverseBonus = 1 / bonus

	if character.HasAura(PowerInfusionAuraID) {
		character.PseudoStats.CastSpeedMultiplier /= 1.2
	}
	character.PseudoStats.CastSpeedMultiplier *= bonus
	character.MultiplyAttackSpeed(sim, bonus)

	character.AddAura(sim, Aura{
		ID:       BloodlustAuraID,
		ActionID: ActionID{SpellID: 2825, Tag: actionTag},
		Expires:  sim.CurrentTime + BloodlustDuration,
		OnExpire: func(sim *Simulation) {
			character.PseudoStats.CastSpeedMultiplier *= inverseBonus
			if character.HasAura(PowerInfusionAuraID) {
				character.PseudoStats.CastSpeedMultiplier *= 1.2
			}
			character.MultiplyAttackSpeed(sim, inverseBonus)
		},
	})

	if len(character.Pets) > 0 {
		for _, petAgent := range character.Pets {
			pet := petAgent.GetPet()
			if pet.IsEnabled() {
				AddBloodlustAura(sim, &pet.Character, actionTag)
			}
		}
	}
}

var PowerInfusionAuraID = NewAuraID()
var sharedPowerInfusionCooldownID = NewCooldownID() // Different from priest PI CD.
const PowerInfusionDuration = time.Second * 15
const PowerInfusionCD = time.Minute * 3

func registerPowerInfusionCD(agent Agent, numPowerInfusions int32) {
	registerExternalConsecutiveCDApproximation(
		agent,
		externalConsecutiveCDApproximation{
			ActionID:         ActionID{SpellID: 10060, Tag: -1},
			AuraID:           PowerInfusionAuraID,
			CooldownID:       sharedPowerInfusionCooldownID,
			CooldownPriority: CooldownPriorityDefault,
			AuraDuration:     PowerInfusionDuration,
			AuraCD:           PowerInfusionCD,
			Type:             CooldownTypeDPS,

			ShouldActivate: func(sim *Simulation, character *Character) bool {
				// Haste portion doesn't stack with Bloodlust, so prefer to wait.
				if character.HasAura(BloodlustAuraID) {
					return false
				}
				return true
			},
			AddAura: func(sim *Simulation, character *Character) { AddPowerInfusionAura(sim, character, -1) },
		},
		numPowerInfusions)
}

func AddPowerInfusionAura(sim *Simulation, character *Character, actionTag int32) {
	const bonus = 1.2
	const inverseBonus = 1 / bonus

	if !character.HasAura(BloodlustAuraID) {
		character.PseudoStats.CastSpeedMultiplier *= bonus
	}

	character.AddAura(sim, Aura{
		ID:       PowerInfusionAuraID,
		ActionID: ActionID{SpellID: 10060, Tag: actionTag},
		Expires:  sim.CurrentTime + PowerInfusionDuration,
		OnCast: func(sim *Simulation, cast *Cast) {
			if cast.Cost.Type == stats.Mana {
				// TODO: Double-check this is how the calculation works.
				cast.Cost.Value = MaxFloat(0, cast.Cost.Value-cast.BaseCost.Value*0.2)
			}
		},
		OnExpire: func(sim *Simulation) {
			if !character.HasAura(BloodlustAuraID) {
				character.PseudoStats.CastSpeedMultiplier *= inverseBonus
			}
		},
	})
}

var sharedInnervateCooldownID = NewCooldownID()
var InnervateAuraID = NewAuraID()

const InnervateDuration = time.Second * 20
const InnervateCD = time.Minute * 6

func InnervateManaThreshold(character *Character) float64 {
	if character.Class == proto.Class_ClassMage {
		// Mages burn mana really fast so they need a higher threshold.
		return character.MaxMana() * 0.7
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
			ActionID:         ActionID{SpellID: 29166, Tag: -1},
			AuraID:           InnervateAuraID,
			CooldownID:       sharedInnervateCooldownID,
			CooldownPriority: CooldownPriorityDefault,
			AuraDuration:     InnervateDuration,
			AuraCD:           InnervateCD,
			Type:             CooldownTypeMana,
			Init: func(sim *Simulation, character *Character) {
				expectedManaPerInnervate = character.SpiritManaRegenPerSecond() * 5 * 20
				remainingInnervateUsages = int(1 + (MaxDuration(0, sim.Duration))/InnervateCD)
				character.ExpectedBonusMana += expectedManaPerInnervate * float64(remainingInnervateUsages)
			},
			ShouldActivate: func(sim *Simulation, character *Character) bool {
				// Only cast innervate when very low on mana, to make sure all other mana CDs are prioritized.
				if character.CurrentMana() > innervateThreshold {
					return false
				}
				return true
			},
			AddAura: func(sim *Simulation, character *Character) {
				AddInnervateAura(sim, character, expectedManaPerInnervate, -1)

				newRemainingUsages := int(sim.GetRemainingDuration() / InnervateCD)
				// AddInnervateAura already accounts for 1 usage, which is why we subtract 1 less.
				character.ExpectedBonusMana -= expectedManaPerInnervate * MaxFloat(0, float64(remainingInnervateUsages-newRemainingUsages-1))
				remainingInnervateUsages = newRemainingUsages
			},
		},
		numInnervates)
}

func AddInnervateAura(sim *Simulation, character *Character, expectedBonusManaReduction float64, actionTag int32) {
	character.PseudoStats.ForceFullSpiritRegen = true
	character.PseudoStats.SpiritRegenMultiplier *= 5.0
	character.UpdateManaRegenRates()

	lastUpdateTime := sim.CurrentTime
	bonusManaSubtracted := 0.0

	character.AddAura(sim, Aura{
		ID:       InnervateAuraID,
		ActionID: ActionID{SpellID: 29166, Tag: actionTag},
		Expires:  sim.CurrentTime + InnervateDuration,
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
			character.UpdateManaRegenRates()

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
			ActionID:         ActionID{SpellID: 16190, Tag: -1},
			AuraID:           ManaTideTotemAuraID,
			CooldownID:       sharedManaTideTotemCooldownID,
			CooldownPriority: CooldownPriorityDefault,
			AuraDuration:     ManaTideTotemDuration,
			AuraCD:           ManaTideTotemCD,
			Type:             CooldownTypeMana,
			Init: func(sim *Simulation, character *Character) {
				// Use first MTT at 60s, or halfway through the fight, whichever comes first.
				initialDelay = MinDuration(sim.Duration/2, time.Second*60)

				expectedManaPerManaTideTotem = ManaTideTotemAmount(character)
				remainingManaTideTotemUsages = int(1 + MaxDuration(0, sim.Duration-initialDelay)/ManaTideTotemCD)
				character.ExpectedBonusMana += expectedManaPerManaTideTotem * float64(remainingManaTideTotemUsages)
			},
			ShouldActivate: func(sim *Simulation, character *Character) bool {
				// A normal resto shaman would wait to use MTT.
				if sim.CurrentTime < initialDelay {
					return false
				}
				return true
			},
			AddAura: func(sim *Simulation, character *Character) {
				AddManaTideTotemAura(sim, character, -1)

				newRemainingUsages := int(sim.GetRemainingDuration() / ManaTideTotemCD)
				// AddManaTideTotemAura already accounts for 1 usage, which is why we subtract 1 less.
				character.ExpectedBonusMana -= expectedManaPerManaTideTotem * MaxFloat(0, float64(remainingManaTideTotemUsages-newRemainingUsages-1))
				remainingManaTideTotemUsages = newRemainingUsages
			},
		},
		numManaTideTotems)
}

func AddManaTideTotemAura(sim *Simulation, character *Character, actionTag int32) {
	lastUpdateTime := sim.CurrentTime
	totalBonusMana := ManaTideTotemAmount(character)
	bonusManaSubtracted := 0.0
	actionID := ActionID{SpellID: 16190, Tag: actionTag}

	character.AddAura(sim, Aura{
		ID:       ManaTideTotemAuraID,
		ActionID: actionID,
		Expires:  sim.CurrentTime + ManaTideTotemDuration,
		OnCast: func(sim *Simulation, cast *Cast) {
			if !character.HasManaBar() {
				return
			}

			timeDelta := sim.CurrentTime - lastUpdateTime
			lastUpdateTime = sim.CurrentTime
			progressRatio := float64(timeDelta) / float64(ManaTideTotemDuration)
			remainder := totalBonusMana - bonusManaSubtracted
			amount := MinFloat(remainder, totalBonusMana*progressRatio)

			character.AddMana(sim, amount, actionID, true)
			character.ExpectedBonusMana -= amount
			bonusManaSubtracted += amount
		},
		OnExpire: func(sim *Simulation) {
			if !character.HasManaBar() {
				return
			}

			remainder := totalBonusMana - bonusManaSubtracted
			character.AddMana(sim, remainder, actionID, true)
			character.ExpectedBonusMana -= remainder
		},
	})
}
