package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

type Buffs struct {
	// Raid buffs
	ArcaneBrilliance bool
	GiftOfTheWild    proto.TristateEffect
	BlessingOfKings  bool
	BlessingOfWisdom proto.TristateEffect
	DivineSpirit     proto.TristateEffect

	// Party class buffs
	MoonkinAura     proto.TristateEffect
	ShadowPriestDPS uint16 // adds Mp5 ~ 25% (dps*5%*5sec = 25%)
	Bloodlust       int32

	// Totems
	ManaSpringTotem proto.TristateEffect
	ManaTideTotem   bool
	TotemOfWrath    int32
	WrathOfAirTotem proto.TristateEffect

	// Target debuff
	JudgementOfWisdom         bool
	ImprovedSealOfTheCrusader bool
	Misery                    bool

	// Drums
	DrumsOfBattle      bool
	DrumsOfRestoration bool

	// Party item buffs
	AtieshMage            int32
	AtieshWarlock         int32
	BraidedEterniumChain  bool
	ChainOfTheTwilightOwl bool
	EyeOfTheNight         bool
	JadePendantOfBlasting bool
}

func ProtoToBuffs(inBuff *proto.Buffs) Buffs {
	return Buffs{
		ArcaneBrilliance: inBuff.ArcaneBrilliance,
		Bloodlust:        inBuff.Bloodlust,
		GiftOfTheWild:    inBuff.GiftOfTheWild,
		BlessingOfKings:  inBuff.BlessingOfKings,
		BlessingOfWisdom: inBuff.BlessingOfWisdom,
		DivineSpirit:     inBuff.DivineSpirit,
		MoonkinAura:      inBuff.MoonkinAura,
		ShadowPriestDPS:  uint16(inBuff.ShadowPriestDps),

		JudgementOfWisdom:         inBuff.JudgementOfWisdom,
		ImprovedSealOfTheCrusader: inBuff.ImprovedSealOfTheCrusader,
		Misery:                    inBuff.Misery,

		ManaSpringTotem: inBuff.ManaSpringTotem,
		ManaTideTotem:   inBuff.ManaTideTotem,
		TotemOfWrath:    inBuff.TotemOfWrath,
		WrathOfAirTotem: inBuff.WrathOfAirTotem,

		DrumsOfBattle:      inBuff.DrumsOfBattle,
		DrumsOfRestoration: inBuff.DrumsOfRestoration,

		AtieshMage:            inBuff.AtieshMage,
		AtieshWarlock:         inBuff.AtieshWarlock,
		BraidedEterniumChain:  inBuff.BraidedEterniumChain,
		ChainOfTheTwilightOwl: inBuff.ChainOfTheTwilightOwl,
		EyeOfTheNight:         inBuff.EyeOfTheNight,
		JadePendantOfBlasting: inBuff.JadePendantOfBlasting,
	}
}

func GetTristateValueFloat(effect proto.TristateEffect, regularValue float64, impValue float64) float64 {
	if effect == proto.TristateEffect_TristateEffectRegular {
		return regularValue
	} else if effect == proto.TristateEffect_TristateEffectImproved {
		return impValue
	} else {
		return 0
	}
}

// Applies buffs that affect the sim as a whole.
func (buffs Buffs) ApplyToSim(sim *Simulation) {
	if buffs.Misery {
		sim.AddInitialAura(func(sim *Simulation) Aura {
			return MiseryAura()
		})
	}

	if buffs.JudgementOfWisdom {
		sim.AddInitialAura(func(sim *Simulation) Aura {
			return AuraJudgementOfWisdom()
		})
	}
}

// Applies buffs that affect individual players.
func (buffs Buffs) ApplyToPlayer(agent Agent) {
	character := agent.GetCharacter()

	if buffs.ArcaneBrilliance {
		character.AddInitialStats(stats.Stats{
			stats.Intellect: 40,
		})
	}

	// TODO: Figure out imp value
	gotwAmount := GetTristateValueFloat(buffs.GiftOfTheWild, 18.0, 18.0)
	// TODO: Pretty sure some of these dont stack with fort/ai/divine spirit
	character.AddInitialStats(stats.Stats{
		stats.Stamina:   gotwAmount,
		stats.Agility:   gotwAmount,
		stats.Strength:  gotwAmount,
		stats.Intellect: gotwAmount,
		stats.Spirit:    gotwAmount,
	})

	character.AddInitialStats(stats.Stats{
		stats.SpellCrit: GetTristateValueFloat(buffs.MoonkinAura, 5*SpellCritRatingPerCritChance, 5*SpellCritRatingPerCritChance+20),
	})

	character.AddInitialStats(stats.Stats{
		stats.Spirit: GetTristateValueFloat(buffs.DivineSpirit, 50.0, 50.0),
	})

	// shadow priest buff bot just statically applies mp5
	if buffs.ShadowPriestDPS > 0 {
		character.AddInitialStats(stats.Stats{
			stats.MP5: float64(buffs.ShadowPriestDPS) * 0.25,
		})
	}

	// TODO: Double-check these numbers
	character.AddInitialStats(stats.Stats{
		stats.MP5: GetTristateValueFloat(buffs.BlessingOfWisdom, 42.0, 50.0),
	})

	if buffs.ImprovedSealOfTheCrusader {
		character.AddInitialStats(stats.Stats{
			stats.SpellCrit: 3 * SpellCritRatingPerCritChance,
		})
		// FUTURE: melee crit bonus, research actual value
	}

	if buffs.TotemOfWrath > 0 {
		character.AddInitialStats(stats.Stats{
			stats.SpellCrit: 3 * SpellCritRatingPerCritChance * float64(buffs.TotemOfWrath),
			stats.SpellHit:  37.8 * float64(buffs.TotemOfWrath),
		})
	}
	character.AddInitialStats(stats.Stats{
		stats.SpellPower: GetTristateValueFloat(buffs.WrathOfAirTotem, 101.0, 121.0),
	})
	character.AddInitialStats(stats.Stats{
		stats.MP5: GetTristateValueFloat(buffs.ManaSpringTotem, 50, 62.5),
	})

	character.AddInitialStats(stats.Stats{
		stats.SpellCrit: 28 * float64(buffs.AtieshMage),
	})
	character.AddInitialStats(stats.Stats{
		stats.SpellPower:   33 * float64(buffs.AtieshWarlock),
		stats.HealingPower: 33 * float64(buffs.AtieshWarlock),
	})

	if buffs.BraidedEterniumChain {
		character.AddInitialStats(stats.Stats{stats.MeleeCrit: 28})
	}
	if buffs.EyeOfTheNight {
		character.AddInitialStats(stats.Stats{stats.SpellPower: 34})
	}
	if buffs.JadePendantOfBlasting {
		character.AddInitialStats(stats.Stats{stats.SpellPower: 15})
	}
	if buffs.ChainOfTheTwilightOwl {
		character.AddInitialStats(stats.Stats{stats.SpellCrit: 2 * SpellCritRatingPerCritChance})
	}
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
			if cast.GetActionID().ItemID == ItemIDTLC {
				return // TLC cant proc JoW
			}

			character := cast.GetAgent().GetCharacter()
			// Only apply to agents that have mana.
			if character.InitialStats[stats.Mana] > 0 {
				character.Stats[stats.Mana] += mana
				if sim.Log != nil {
					sim.Log("(%d) +Judgement Of Wisdom: 37 mana (74 @ 50%% proc)\n", character.ID)
				}
			}
		},
	}
}

type RaceBonusType byte

// These values are used directly in the dropdown in index.html
const (
	RaceBonusTypeNone RaceBonusType = iota
	RaceBonusTypeBloodElf
	RaceBonusTypeDraenei
	RaceBonusTypeDwarf
	RaceBonusTypeGnome
	RaceBonusTypeHuman
	RaceBonusTypeNightElf
	RaceBonusTypeOrc
	RaceBonusTypeTauren
	RaceBonusTypeTroll10
	RaceBonusTypeTroll30
	RaceBonusTypeUndead
)

func TryActivateRacial(sim *Simulation, agent Agent) {
	switch agent.GetCharacter().Race {
	case RaceBonusTypeOrc:
		if agent.GetCharacter().IsOnCD(MagicIDOrcBloodFury, sim.CurrentTime) {
			return
		}

		const spBonus = 143
		const dur = time.Second * 15
		const cd = time.Minute * 2

		agent.GetCharacter().SetCD(MagicIDOrcBloodFury, cd+sim.CurrentTime)
		AddAuraWithTemporaryStats(sim, agent, MagicIDOrcBloodFury, stats.SpellPower, spBonus, dur)

	case RaceBonusTypeTroll10, RaceBonusTypeTroll30:
		if agent.GetCharacter().IsOnCD(MagicIDTrollBerserking, sim.CurrentTime) {
			return
		}
		hasteBonus := time.Duration(11) // 10% haste
		if agent.GetCharacter().Race == RaceBonusTypeTroll30 {
			hasteBonus = time.Duration(13) // 30% haste
		}
		const dur = time.Second * 10
		const cd = time.Minute * 3

		agent.GetCharacter().SetCD(MagicIDTrollBerserking, cd+sim.CurrentTime)
		agent.GetCharacter().AddAura(sim, Aura{
			ID:      MagicIDTrollBerserking,
			Expires: sim.CurrentTime + dur,
			OnCast: func(sim *Simulation, cast DirectCastAction, inputs *DirectCastInput) {
				inputs.CastTime = (inputs.CastTime * 10) / hasteBonus
			},
		})
	}
}
