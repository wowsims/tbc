package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

// Totems that shaman will cast.
// TODO: add logic inside these to select each totem based on options on the shaman?
// TODO: Include mental quickness mana cost reduction when we figure out what it is.

func (shaman *Shaman) NewAirTotem(sim *core.Simulation) *core.SimpleCast {
	manaCost := 360 * (1 - float64(shaman.Talents.TotemicFocus)*0.05)

	cast := &core.SimpleCast{
		Cast: core.Cast{
			Name:            "Wrath of Air",
			ActionID:        core.ActionID{SpellID: 3738},
			Character:       shaman.GetCharacter(),
			BaseManaCost:    manaCost,
			ManaCost:        manaCost,
			CastTime:        time.Second * 1,
			IgnoreCooldowns: true, // lets us override the GCD
		},
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
			shaman.SelfBuffs.NextTotemDrops[AirTotem] = sim.CurrentTime + time.Second*120
		},
	}
	cast.Init(sim)
	return cast
}

func (shaman *Shaman) NewWaterTotem(sim *core.Simulation) *core.SimpleCast {
	manaCost := 120 * (1 - float64(shaman.Talents.TotemicFocus)*0.05)

	cast := &core.SimpleCast{
		Cast: core.Cast{
			Name:            "Mana Stream",
			ActionID:        core.ActionID{SpellID: 25570},
			Character:       shaman.GetCharacter(),
			BaseManaCost:    manaCost,
			ManaCost:        manaCost,
			CastTime:        time.Second * 1,
			IgnoreCooldowns: true, // lets us override the GCD
		},
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
			shaman.SelfBuffs.NextTotemDrops[WaterTotem] = sim.CurrentTime + time.Second*120
		},
	}
	cast.Init(sim)
	return cast
}

func (shaman *Shaman) NewFireTotem(sim *core.Simulation) *core.SimpleCast {
	manaCost := 360 * (1 - float64(shaman.Talents.TotemicFocus)*0.05)

	cast := &core.SimpleCast{
		Cast: core.Cast{
			Name:            "Totem of Wrath",
			ActionID:        core.ActionID{SpellID: 30706},
			Character:       shaman.GetCharacter(),
			BaseManaCost:    manaCost,
			ManaCost:        manaCost,
			CastTime:        time.Second * 1,
			IgnoreCooldowns: true, // lets us override the GCD
		},
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
			shaman.SelfBuffs.NextTotemDrops[FireTotem] = sim.CurrentTime + time.Second*120
		},
	}
	cast.Init(sim)
	return cast
}
