package core

import (
	"time"

	"github.com/wowsims/tbc/sim/api"
	"github.com/wowsims/tbc/sim/core/stats"
)

type Buffs struct {
	// // Raid Buff
	// ArcaneBrilliance bool
	// BlessingOfKings  bool
	// BlessingOfWisdom TristateEffect
	// DivineSpirit     TristateEffect
	// GiftOfTheWild    TristateEffect

	// // Party Buff
	// Bloodlust       int32
	// MoonkinAura     TristateEffect
	// ShadowPriestDps int32

	// // Target Debuff
	// JudgementOfWisdom         bool
	// ImprovedSealOfTheCrusader bool
	// Misery                    bool

	// // Items
	// AtieshMage            int32
	// AtieshWarlock         int32
	// BraidedEterniumChain  bool
	// EyeOfTheNight         bool
	// ChainOfTheTwilightOwl bool
	// JadePendantOfBlasting bool

	// Totems
	ManaSpringTotem api.TristateEffect
	ManaTideTotem   bool
	TotemOfWrath    int32
	WrathOfAirTotem api.TristateEffect

	// Raid buffs
	ArcaneInt                bool
	GiftOfTheWild            bool
	BlessingOfKings          bool
	ImprovedBlessingOfWisdom bool
	ImprovedDivineSpirit     bool

	// Party class buffs
	Moonkin             bool
	MoonkinRavenGoddess bool   // adds 20 spell crit to moonkin aura
	SpriestDPS          uint16 // adds Mp5 ~ 25% (dps*5%*5sec = 25%)
	Bloodlust           int

	// TODO: Do these need to be here? Should I just use this instead of the shaman.Totems struct?
	//  Balance druids wouldnt want to import shaman.Totems probably? Or maybe those can go in a reduced package that can be shared.
	// WrathOfAir          bool
	// TotemOfWrath        bool
	// ManaStream          bool

	// Party item buffs
	EyeOfNight  bool // Eye of night bonus from party member (not you)
	TwilightOwl bool // from party member

	// Target debuff
	JudgementOfWisdom         bool
	ImprovedSealOfTheCrusader bool
	Misery                    bool
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
			OnCast: func(sim *Simulation, agent Agent, c *Cast) {
				c.CastTime = (c.CastTime * 10) / hasteBonus
			},
		})
	}
}
