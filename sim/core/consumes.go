package core

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

type Consumes struct {
	// Buffs
	BrilliantWizardOil       bool
	SuperiorWizardOil        bool
	ElixirOfMajorMageblood   bool
	FlaskOfBlindingLight     bool
	FlaskOfMightyRestoration bool
	FlaskOfPureDeath         bool
	FlaskOfSupremePower      bool
	BlackenedBasilisk        bool
	SkullfishSoup            bool
	AdeptsElixir             bool
	ElixirOfMajorFirePower   bool
	ElixirOfMajorFrostPower  bool
	ElixirOfMajorShadowPower bool
	ElixirOfDraenicWisdom    bool

	// Used in rotations
	DefaultPotion      proto.Potions
	StartingPotion     proto.Potions
	NumStartingPotions int32
	DarkRune           bool
	Drums              proto.Drums
}

func ProtoToConsumes(c *proto.Consumes) Consumes {
	return Consumes{
		FlaskOfBlindingLight:     c.FlaskOfBlindingLight,
		FlaskOfMightyRestoration: c.FlaskOfMightyRestoration,
		FlaskOfPureDeath:         c.FlaskOfPureDeath,
		FlaskOfSupremePower:      c.FlaskOfSupremePower,
		AdeptsElixir:             c.AdeptsElixir,
		ElixirOfMajorFirePower:   c.ElixirOfMajorFirePower,
		ElixirOfMajorFrostPower:  c.ElixirOfMajorFrostPower,
		ElixirOfMajorShadowPower: c.ElixirOfMajorShadowPower,
		ElixirOfDraenicWisdom:    c.ElixirOfDraenicWisdom,
		ElixirOfMajorMageblood:   c.ElixirOfMajorMageblood,
		BrilliantWizardOil:       c.BrilliantWizardOil,
		SuperiorWizardOil:        c.SuperiorWizardOil,
		BlackenedBasilisk:        c.BlackenedBasilisk,
		SkullfishSoup:            c.SkullfishSoup,
		DefaultPotion:            c.DefaultPotion,
		StartingPotion:           c.StartingPotion,
		NumStartingPotions:       c.NumStartingPotions,
		DarkRune:                 c.DarkRune,
		Drums:                    c.Drums,
	}
}

func (c Consumes) Stats() stats.Stats {
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

func TryActivateDrums(sim *Simulation, agent Agent) {
	character := agent.GetCharacter()
	if character.IsOnCD(MagicIDDrums, sim.CurrentTime) {
		return
	}

	partyCast := character.Party.Buffs.Drums
	if partyCast == proto.Drums_DrumsUnknown {
		return
	}

	// TODO: If this character has the drums set too, then do a cast time
	//selfCast := character.Consumes.Drums

	if partyCast == proto.Drums_DrumsOfBattle {
		const hasteBonus = 80
		for _, agent := range character.Party.Players {
			agent.GetCharacter().SetCD(MagicIDDrums, time.Minute*2+sim.CurrentTime) // tinnitus
			AddAuraWithTemporaryStats(sim, agent, MagicIDDrums, stats.SpellHaste, hasteBonus, time.Second*30)
		}
	} else if partyCast == proto.Drums_DrumsOfRestoration {
		// 600 mana over 15 seconds == 200 mp5
		const mp5Bonus = 200
		for _, agent := range character.Party.Players {
			agent.GetCharacter().SetCD(MagicIDDrums, time.Minute*2+sim.CurrentTime) // tinnitus
			AddAuraWithTemporaryStats(sim, agent, MagicIDDrums, stats.MP5, mp5Bonus, time.Second*15)
		}
	}
}

func TryActivatePotion(sim *Simulation, agent Agent) {
	character := agent.GetCharacter()
	if character.IsOnCD(MagicIDPotion, sim.CurrentTime) {
		return
	}

	potionToUse := character.Consumes.DefaultPotion
	if character.Consumes.StartingPotion != proto.Potions_UnknownPotion && character.potionsUsed < character.Consumes.NumStartingPotions {
		potionToUse = character.Consumes.StartingPotion
	}

	if potionToUse == proto.Potions_UnknownPotion {
		return
	}

	if potionToUse == proto.Potions_DestructionPotion {
		const spBonus = 120
		const critBonus = 44.16
		const dur = time.Second * 15

		character.SetCD(MagicIDPotion, time.Second*120+sim.CurrentTime)
		character.Stats[stats.SpellPower] += spBonus
		character.Stats[stats.SpellCrit] += critBonus

		character.AddAura(sim, Aura{
			ID:      MagicIDDestructionPotion,
			Expires: sim.CurrentTime + dur,
			OnExpire: func(sim *Simulation) {
				character.Stats[stats.SpellPower] -= spBonus
				character.Stats[stats.SpellCrit] -= critBonus
			},
		})
	} else if potionToUse == proto.Potions_SuperManaPotion {
		// Only pop if we have less than the max mana provided by the potion minus 1mp5 tick.
		totalRegen := character.manaRegenPerSecond() * 5
		if character.InitialStats[stats.Mana]-(character.Stats[stats.Mana]+totalRegen) < 3000 {
			return
		}

		// Restores 1800 to 3000 mana. (2 Min Cooldown)
		manaGain := 1800 + (sim.Rando.Float64("super mana") * 1200)

		if character.HasAura(MagicIDAlchStone) {
			manaGain *= 1.4
		}

		character.Stats[stats.Mana] += manaGain
		character.SetCD(MagicIDPotion, time.Second*120+sim.CurrentTime)
		if sim.Log != nil {
			sim.Log("Used Mana Potion\n")
		}
	}

	character.potionsUsed++
}

func TryActivateDarkRune(sim *Simulation, agent Agent) {
	character := agent.GetCharacter()
	if !character.Consumes.DarkRune || character.IsOnCD(MagicIDRune, sim.CurrentTime) {
		return
	}

	// Only pop if we have less than the max mana provided by the potion minus 1mp5 tick.
	totalRegen := character.manaRegenPerSecond() * 5
	if character.InitialStats[stats.Mana]-(character.Stats[stats.Mana]+totalRegen) < 1500 {
		return
	}

	// Restores 900 to 1500 mana. (2 Min Cooldown)
	character.Stats[stats.Mana] += 900 + (sim.Rando.Float64("dark rune") * 600)
	character.SetCD(MagicIDRune, time.Second*120+sim.CurrentTime)
	if sim.Log != nil {
		sim.Log("Used Dark Rune\n")
	}
}
