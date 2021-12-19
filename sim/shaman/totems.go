package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

// Totems that shaman will cast.
// TODO: add logic inside these to select each totem based on options on the shaman?
// TODO: Include mental quickness mana cost reduction when we figure out what it is.

func (shaman *Shaman) NewAirTotem(sim *core.Simulation) *core.SimpleCast {
	baseManaCost := 320.0
	manaCost := baseManaCost * (1 - float64(shaman.Talents.TotemicFocus)*0.05)

	cast := &core.SimpleCast{
		Cast: core.Cast{
			Name:         "Wrath of Air",
			ActionID:     core.ActionID{SpellID: 3738},
			Character:    shaman.GetCharacter(),
			BaseManaCost: baseManaCost,
			ManaCost:     manaCost,
			GCDCooldown:  time.Second * 1,
		},
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
			shaman.SelfBuffs.NextTotemDrops[AirTotem] = sim.CurrentTime + time.Second*120
		},
	}
	cast.Init(sim)
	return cast
}

func (shaman *Shaman) NewWaterTotem(sim *core.Simulation) *core.SimpleCast {
	baseManaCost := 120.0
	manaCost := baseManaCost * (1 - float64(shaman.Talents.TotemicFocus)*0.05)

	cast := &core.SimpleCast{
		Cast: core.Cast{
			Name:         "Mana Stream",
			ActionID:     core.ActionID{SpellID: 25570},
			Character:    shaman.GetCharacter(),
			BaseManaCost: baseManaCost,
			ManaCost:     manaCost,
			GCDCooldown:  time.Second * 1,
		},
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
			shaman.SelfBuffs.NextTotemDrops[WaterTotem] = sim.CurrentTime + time.Second*120
		},
	}
	cast.Init(sim)
	return cast
}

func (shaman *Shaman) NewFireTotem(sim *core.Simulation) *core.SimpleCast {
	baseManaCost := shaman.BaseMana() * 0.05
	manaCost := baseManaCost * (1 - float64(shaman.Talents.TotemicFocus)*0.05)

	cast := &core.SimpleCast{
		Cast: core.Cast{
			Name:         "Totem of Wrath",
			ActionID:     core.ActionID{SpellID: 30706},
			Character:    shaman.GetCharacter(),
			BaseManaCost: baseManaCost,
			ManaCost:     manaCost,
			GCDCooldown:  time.Second * 1,
		},
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
			shaman.SelfBuffs.NextTotemDrops[FireTotem] = sim.CurrentTime + time.Second*120
		},
	}
	cast.Init(sim)
	return cast
}

// TryDropTotems will check to see if totems need to be re-cast.
//  If they do time.Duration will be returned will be >0.
func (shaman *Shaman) TryDropTotems(sim *core.Simulation) time.Duration {

	var cast *core.SimpleCast

	// currently hardcoded to include 25% mana cost reduction from resto talents
	for totemTypeIdx, totemExpiration := range shaman.SelfBuffs.NextTotemDrops {
		if cast != nil {
			break
		}
		if sim.CurrentTime > totemExpiration {
			switch totemTypeIdx {
			case AirTotem:
				cast = shaman.NewAirTotem(sim)
			case EarthTotem:
				// no earth totem right now
			case FireTotem:
				cast = shaman.NewFireTotem(sim)
			case WaterTotem:
				cast = shaman.NewWaterTotem(sim)
			}
		}
	}

	if cast == nil {
		return 0 // no totem to cast
	}

	success := cast.StartCast(sim)
	if !success {
		regenTime := shaman.TimeUntilManaRegen(cast.GetManaCost())
		shaman.Character.Metrics.MarkOOM(sim, &shaman.Character, regenTime)
		return sim.CurrentTime + regenTime
	}
	return sim.CurrentTime + shaman.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime)
}
