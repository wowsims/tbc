package core

import (
	"time"
)

type Consumes struct {
	// Buffs
	BrilliantWizardOil       bool
	MajorMageblood           bool
	FlaskOfBlindingLight     bool
	FlaskOfMightyRestoration bool
	BlackendBasilisk         bool

	// Used in rotations
	DestructionPotion bool
	SuperManaPotion   bool
	DarkRune          bool
	DrumsOfBattle     bool
}

func (c Consumes) AddStats(s Stats) Stats {
	if c.BrilliantWizardOil {
		s[StatSpellCrit] += 14
		s[StatSpellPower] += 36
	}
	if c.MajorMageblood {
		s[StatMP5] += 16.0
	}
	if c.FlaskOfBlindingLight {
		s[StatSpellPower] += 80
	}
	if c.FlaskOfMightyRestoration {
		s[StatMP5] += 25
	}
	if c.BlackendBasilisk {
		s[StatSpellPower] += 23
	}
	return s
}

func TryActivateDrums(sim *Simulation, party *Party, player *Player) {
	if !player.Consumes.DrumsOfBattle || player.IsOnCD(MagicIDDrums, sim.CurrentTime) {
		return
	}

	const hasteBonus = 80
	for _, p := range party.Players {
		p.Stats[StatSpellHaste] += hasteBonus
		p.SetCD(MagicIDDrums, time.Minute*2+sim.CurrentTime) // tinnitus
		p.AddAura(sim, AuraStatRemoval(sim.CurrentTime, time.Second*30, hasteBonus, StatSpellHaste, MagicIDDrums))
	}
}

func TryActivateDestructionPotion(sim *Simulation, party *Party, player *Player) {
	if !player.Consumes.DestructionPotion || player.IsOnCD(MagicIDPotion, sim.CurrentTime) {
		return
	}

	// Only use dest potion if not using mana or if we haven't used it once.
	// If we are using mana, only use destruction potion on the pull.
	if player.destructionPotionUsed && player.Consumes.SuperManaPotion {
		return
	}

	const spBonus = 120
	const critBonus = 44.16
	const dur = time.Second * 15

	player.destructionPotionUsed = true
	player.SetCD(MagicIDPotion, time.Second*120+sim.CurrentTime)
	player.Stats[StatSpellPower] += spBonus
	player.Stats[StatSpellCrit] += critBonus

	player.AddAura(sim, Aura{
		ID:      MagicIDDestructionPotion,
		Expires: sim.CurrentTime + dur,
		OnExpire: func(sim *Simulation, player *Player, c *Cast) {
			player.Stats[StatSpellPower] -= spBonus
			player.Stats[StatSpellCrit] -= critBonus
		},
	})
}

func TryActivateDarkRune(sim *Simulation, party *Party, player *Player) {
	if !player.Consumes.DarkRune || player.IsOnCD(MagicIDRune, sim.CurrentTime) {
		return
	}

	// Only pop if we have less than the max mana provided by the potion minus 1mp5 tick.
	totalRegen := player.manaRegenPerSecond() * 5
	if player.Stats[StatMana]-(player.Stats[StatMana]+totalRegen) < 1500 {
		return
	}

	// Restores 900 to 1500 mana. (2 Min Cooldown)
	player.Stats[StatMana] += 900 + (sim.Rando.Float64() * 600)
	player.SetCD(MagicIDRune, time.Second*120+sim.CurrentTime)
	if sim.Debug != nil {
		sim.Debug("Used Dark Rune\n")
	}
	return
}

func TryActivateSuperManaPotion(sim *Simulation, party *Party, player *Player) {
	if !player.Consumes.SuperManaPotion || player.IsOnCD(MagicIDPotion, sim.CurrentTime) {
		return
	}

	// Only pop if we have less than the max mana provided by the potion minus 1mp5 tick.
	totalRegen := player.manaRegenPerSecond() * 5
	if player.Stats[StatMana]-(player.Stats[StatMana]+totalRegen) < 3000 {
		return
	}

	// Restores 1800 to 3000 mana. (2 Min Cooldown)
	manaGain := 1800 + (sim.Rando.Float64() * 1200)

	if player.HasAura(MagicIDAlchStone) {
		manaGain *= 1.4
	}

	player.Stats[StatMana] += manaGain
	player.SetCD(MagicIDPotion, time.Second*120+sim.CurrentTime)
	if sim.Debug != nil {
		sim.Debug("Used Mana Potion\n")
	}
	return
}
