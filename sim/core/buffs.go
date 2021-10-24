package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Applies buffs that affect the sim as a whole.
func applyBuffsToSim(sim *Simulation, buffs proto.Buffs) {
	if buffs.Misery {
		sim.AddPermanentAura(func(sim *Simulation) Aura {
			return MiseryAura()
		})
	}

	if buffs.JudgementOfWisdom {
		sim.AddPermanentAura(func(sim *Simulation) Aura {
			return AuraJudgementOfWisdom()
		})
	}
}

// Applies buffs that affect individual players.
func applyBuffEffects(agent Agent, buffs proto.Buffs) {
	character := agent.GetCharacter()

	if buffs.ArcaneBrilliance {
		character.AddStats(stats.Stats{
			stats.Intellect: 40,
		})
	}

	// TODO: Double-check these numbers.
	gotwAmount := GetTristateValueFloat(buffs.GiftOfTheWild, 14.0, 14.0 * 1.35)
	// TODO: Pretty sure some of these dont stack with fort/ai/divine spirit
	character.AddStats(stats.Stats{
		stats.Stamina:   gotwAmount,
		stats.Agility:   gotwAmount,
		stats.Strength:  gotwAmount,
		stats.Intellect: gotwAmount,
		stats.Spirit:    gotwAmount,
	})

	character.AddStats(stats.Stats{
		stats.SpellCrit: GetTristateValueFloat(buffs.MoonkinAura, 5*SpellCritRatingPerCritChance, 5*SpellCritRatingPerCritChance+20),
	})

	if (buffs.DraeneiRacialMelee) {
		character.AddStats(stats.Stats{
			stats.MeleeHit: 1 * MeleeHitRatingPerHitChance,
		})
	}

	if (buffs.DraeneiRacialCaster) {
		character.AddStats(stats.Stats{
			stats.SpellHit: 1 * SpellHitRatingPerHitChance,
		})
	}

	character.AddStats(stats.Stats{
		stats.Spirit: GetTristateValueFloat(buffs.DivineSpirit, 50.0, 50.0),
	})
	if buffs.DivineSpirit == proto.TristateEffect_TristateEffectImproved {
		character.AddStatDependency(stats.StatDependency{
			SourceStat: stats.Spirit,
			ModifiedStat: stats.SpellPower,
			Modifier: func(spirit float64, spellPower float64) float64 {
				return spellPower + spirit * 0.1
			},
		})
	}

	// shadow priest buff bot just statically applies mp5
	if buffs.ShadowPriestDps > 0 {
		character.AddStats(stats.Stats{
			stats.MP5: float64(buffs.ShadowPriestDps) * 0.25,
		})
	}

	// TODO: Double-check these numbers
	character.AddStats(stats.Stats{
		stats.MP5: GetTristateValueFloat(buffs.BlessingOfWisdom, 42.0, 50.0),
	})

	if buffs.BlessingOfKings {
		bokStats := [5]stats.Stat{
			stats.Agility,
			stats.Strength,
			stats.Stamina,
			stats.Intellect,
			stats.Spirit,
		}

		for _, stat := range bokStats {
			character.AddStatDependency(stats.StatDependency{
				SourceStat: stat,
				ModifiedStat: stat,
				Modifier: func(curValue float64, _ float64) float64 {
					return curValue * 1.1
				},
			})
		}
	}

	if buffs.ImprovedSealOfTheCrusader {
		character.AddStats(stats.Stats{
			stats.SpellCrit: 3 * SpellCritRatingPerCritChance,
		})
		// FUTURE: melee crit bonus, research actual value
	}

	if buffs.TotemOfWrath > 0 {
		character.AddStats(stats.Stats{
			stats.SpellCrit: 3 * SpellCritRatingPerCritChance * float64(buffs.TotemOfWrath),
			stats.SpellHit:  3 * SpellHitRatingPerHitChance * float64(buffs.TotemOfWrath),
		})
	}
	character.AddStats(stats.Stats{
		stats.SpellPower: GetTristateValueFloat(buffs.WrathOfAirTotem, 101.0, 121.0),
	})
	character.AddStats(stats.Stats{
		stats.MP5: GetTristateValueFloat(buffs.ManaSpringTotem, 50, 62.5),
	})

	registerBloodlustCD(agent, buffs)

	character.AddStats(stats.Stats{
		stats.SpellCrit: 28 * float64(buffs.AtieshMage),
	})
	character.AddStats(stats.Stats{
		stats.SpellPower:   33 * float64(buffs.AtieshWarlock),
		stats.HealingPower: 33 * float64(buffs.AtieshWarlock),
	})

	if buffs.BraidedEterniumChain {
		character.AddStats(stats.Stats{stats.MeleeCrit: 28})
	}
	if buffs.EyeOfTheNight {
		character.AddStats(stats.Stats{stats.SpellPower: 34})
	}
	if buffs.JadePendantOfBlasting {
		character.AddStats(stats.Stats{stats.SpellPower: 15})
	}
	if buffs.ChainOfTheTwilightOwl {
		character.AddStats(stats.Stats{stats.SpellCrit: 2 * SpellCritRatingPerCritChance})
	}
}

func registerBloodlustCD(agent Agent, buffs proto.Buffs) {
	numBloodlusts := buffs.Bloodlust
	if numBloodlusts == 0 {
		return
	}

	const dur = time.Second * 40

	agent.GetCharacter().AddMajorCooldown(MajorCooldown{
		CooldownID: MagicIDBloodlust,
		Cooldown: dur, // assumes that multiple BLs are different shaman.
		Priority: CooldownPriorityBloodlust,
		ActivationFactory: func(sim *Simulation) CooldownActivation {
			// Capture this inside ActivationFactory so it resets on Sim reset.
			bloodlustsUsed := int32(0)

			return func(sim *Simulation, character *Character) bool {
				if bloodlustsUsed < numBloodlusts {
					character.SetCD(MagicIDBloodlust, sim.CurrentTime+dur)
					character.Party.AddAura(sim, Aura{
						ID:      MagicIDBloodlust,
						Expires: sim.CurrentTime + dur,
						OnCast: func(sim *Simulation, cast DirectCastAction, input *DirectCastInput) {
							// Multiply and divide lets us use integer math, which is better for perf.
							input.CastTime = (input.CastTime * 10) / 13 // 30% faster
						},
					})
					bloodlustsUsed++
					return true
				} else {
					character.SetCD(MagicIDBloodlust, sim.CurrentTime+time.Minute*10)
					return true
				}
			}
		},
	})
}

func MiseryAura() Aura {
	return Aura{
		ID:      MagicIDMisery,
		Expires: NeverExpires,
		OnSpellHit: func(sim *Simulation, cast DirectCastAction, result *DirectCastDamageResult) {
			result.Damage *= 1.05
		},
	}
}

func AuraJudgementOfWisdom() Aura {
	const mana = 74 / 2 // 50% proc
	return Aura{
		ID:      MagicIDJoW,
		Expires: NeverExpires,
		OnSpellHit: func(sim *Simulation, cast DirectCastAction, result *DirectCastDamageResult) {
			if cast.GetActionID().ItemID == ItemIDTheLightningCapacitor {
				return // TLC cant proc JoW
			}

			character := cast.GetCharacter()
			// Only apply to agents that have mana.
			if character.MaxMana() > 0 {
				character.AddStat(stats.Mana, mana)
				if sim.Log != nil {
					sim.Log("(%d) +Judgement Of Wisdom: 37 mana (74 @ 50%% proc)\n", character.ID)
				}
			}
		},
	}
}
