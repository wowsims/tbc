package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var OrcBloodFuryAuraID = NewAuraID()
var OrcBloodFuryCooldownID = NewCooldownID()

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
			SourceStat: stats.Intellect,
			ModifiedStat: stats.Intellect,
			Modifier: func(intellect float64, _ float64) float64 {
				return intellect * 1.05
			},
		})
	case proto.Race_RaceHuman:
		character.AddStatDependency(stats.StatDependency{
			SourceStat: stats.Spirit,
			ModifiedStat: stats.Spirit,
			Modifier: func(spirit float64, _ float64) float64 {
				return spirit * 1.1
			},
			// TODO: +5 expertise for swords and maces
		})
	case proto.Race_RaceNightElf:
	case proto.Race_RaceOrc:
		// TODO: Pet melee damage +5%
		// TODO: +5 expertise with axes
		const cd = time.Minute * 2
		const dur = time.Second * 15
		const spBonus = 143

		character.AddMajorCooldown(MajorCooldown{
			CooldownID: OrcBloodFuryCooldownID,
			Cooldown: cd,
			ActivationFactory: func(sim *Simulation) CooldownActivation {
				return func(sim *Simulation, character *Character) bool {
					character.SetCD(OrcBloodFuryCooldownID, cd+sim.CurrentTime)
					character.AddAuraWithTemporaryStats(sim, OrcBloodFuryAuraID, "Orc Blood Fury", stats.SpellPower, spBonus, dur)
					return true
				}
			},
		})

	case proto.Race_RaceTauren:
		// TODO: Health +5%
	case proto.Race_RaceTroll10, proto.Race_RaceTroll30:
		// TODO: +1% ranged crit when using a bow

		// Beast Slaying (+5% damage to beasts)
		character.AddPermanentAura(func(sim *Simulation) Aura {
			return Aura{
				ID:      TrollBeastSlayingAuraID,
				Name:    "Beast Slaying (Troll Racial)",
				OnSpellHit: func(sim *Simulation, cast DirectCastAction, result *DirectCastDamageResult) {
					if result.Target.MobType == proto.MobType_MobTypeBeast {
						result.Damage *= 1.05
					}
				},
			}
		})

		// Berserking
		hasteBonus := time.Duration(11) // 10% haste
		if character.Race == proto.Race_RaceTroll30 {
			hasteBonus = time.Duration(13) // 30% haste
		}
		const dur = time.Second * 10
		const cd = time.Minute * 3

		character.AddMajorCooldown(MajorCooldown{
			CooldownID: TrollBerserkingCooldownID,
			Cooldown: cd,
			ActivationFactory: func(sim *Simulation) CooldownActivation {
				return func(sim *Simulation, character *Character) bool {
					character.SetCD(TrollBerserkingCooldownID, cd+sim.CurrentTime)
					character.AddAura(sim, Aura{
						ID:      TrollBerserkingAuraID,
						Name:    "Troll Berserking",
						Expires: sim.CurrentTime + dur,
						OnCast: func(sim *Simulation, cast DirectCastAction, inputs *DirectCastInput) {
							// Multiplying and then dividing lets us use integer multiplication/division which is faster.
							inputs.CastTime = (inputs.CastTime * 10) / hasteBonus
						},
					})
					return true
				}
			},
		})
	case proto.Race_RaceUndead:
	}
}
