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
	DestructionPotion  bool
	SuperManaPotion    bool
	DarkRune           bool
	DrumsOfBattle      bool
	DrumsOfRestoration bool
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
		DestructionPotion:        c.DestructionPotion,
		SuperManaPotion:          c.SuperManaPotion,
		DarkRune:                 c.DarkRune,
		DrumsOfBattle:            c.DrumsOfBattle,
		DrumsOfRestoration:       c.DrumsOfRestoration,
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
	if !agent.GetCharacter().Consumes.DrumsOfBattle || agent.GetCharacter().IsOnCD(MagicIDDrums, sim.CurrentTime) {
		return
	}

	const hasteBonus = 80
	for _, agent := range agent.GetCharacter().Party.Players {
		agent.GetCharacter().SetCD(MagicIDDrums, time.Minute*2+sim.CurrentTime) // tinnitus
		AddAuraWithTemporaryStats(sim, agent, MagicIDDrums, stats.SpellHaste, hasteBonus, time.Second*30)
	}
}

func TryActivateDestructionPotion(sim *Simulation, agent Agent) {
	if !agent.GetCharacter().Consumes.DestructionPotion || agent.GetCharacter().IsOnCD(MagicIDPotion, sim.CurrentTime) {
		return
	}

	// Only use dest potion if not using mana or if we haven't used it once.
	// If we are using mana, only use destruction potion on the pull.
	if agent.GetCharacter().destructionPotionUsed && agent.GetCharacter().Consumes.SuperManaPotion {
		return
	}

	const spBonus = 120
	const critBonus = 44.16
	const dur = time.Second * 15

	agent.GetCharacter().destructionPotionUsed = true
	agent.GetCharacter().SetCD(MagicIDPotion, time.Second*120+sim.CurrentTime)
	agent.GetCharacter().Stats[stats.SpellPower] += spBonus
	agent.GetCharacter().Stats[stats.SpellCrit] += critBonus

	agent.GetCharacter().AddAura(sim, Aura{
		ID:      MagicIDDestructionPotion,
		Expires: sim.CurrentTime + dur,
		OnExpire: func(sim *Simulation, cast *Cast) {
			agent.GetCharacter().Stats[stats.SpellPower] -= spBonus
			agent.GetCharacter().Stats[stats.SpellCrit] -= critBonus
		},
	})
}

func TryActivateDarkRune(sim *Simulation, agent Agent) {
	if !agent.GetCharacter().Consumes.DarkRune || agent.GetCharacter().IsOnCD(MagicIDRune, sim.CurrentTime) {
		return
	}

	// Only pop if we have less than the max mana provided by the potion minus 1mp5 tick.
	totalRegen := agent.GetCharacter().manaRegenPerSecond() * 5
	if agent.GetCharacter().InitialStats[stats.Mana]-(agent.GetCharacter().Stats[stats.Mana]+totalRegen) < 1500 {
		return
	}

	// Restores 900 to 1500 mana. (2 Min Cooldown)
	agent.GetCharacter().Stats[stats.Mana] += 900 + (sim.Rando.Float64("dark rune") * 600)
	agent.GetCharacter().SetCD(MagicIDRune, time.Second*120+sim.CurrentTime)
	if sim.Log != nil {
		sim.Log("Used Dark Rune\n")
	}
}

func TryActivateSuperManaPotion(sim *Simulation, agent Agent) {
	if !agent.GetCharacter().Consumes.SuperManaPotion || agent.GetCharacter().IsOnCD(MagicIDPotion, sim.CurrentTime) {
		return
	}

	// Only pop if we have less than the max mana provided by the potion minus 1mp5 tick.
	totalRegen := agent.GetCharacter().manaRegenPerSecond() * 5
	if agent.GetCharacter().InitialStats[stats.Mana]-(agent.GetCharacter().Stats[stats.Mana]+totalRegen) < 3000 {
		return
	}

	// Restores 1800 to 3000 mana. (2 Min Cooldown)
	manaGain := 1800 + (sim.Rando.Float64("super mana") * 1200)

	if agent.GetCharacter().HasAura(MagicIDAlchStone) {
		manaGain *= 1.4
	}

	agent.GetCharacter().Stats[stats.Mana] += manaGain
	agent.GetCharacter().SetCD(MagicIDPotion, time.Second*120+sim.CurrentTime)
	if sim.Log != nil {
		sim.Log("Used Mana Potion\n")
	}
}
