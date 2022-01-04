package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

const windfuryTotem = int32(proto.AirTotem_WindfuryTotem)
const wrathOfAirTotem = int32(proto.AirTotem_WrathOfAirTotem)
const graceOfAirTotem = int32(proto.AirTotem_GraceOfAirTotem)

// Totems that shaman will cast.
// TODO: add logic inside these to select each totem based on options on the shaman?
// TODO: Include mental quickness mana cost reduction when we figure out what it is.

func (shaman *Shaman) NewWrathOfAirTotem(sim *core.Simulation) *core.SimpleCast {
	baseManaCost := 320.0
	manaCost := baseManaCost * (1 - float64(shaman.Talents.TotemicFocus)*0.05)
	manaCost -= baseManaCost * float64(shaman.Talents.MentalQuickness) * 0.02

	cast := &core.SimpleCast{
		Cast: core.Cast{
			Name:         "Wrath of Air",
			ActionID:     core.ActionID{SpellID: 3738},
			Character:    shaman.GetCharacter(),
			BaseManaCost: baseManaCost,
			ManaCost:     manaCost,
			GCDCooldown:  time.Second * 1,
		},
	}
	cast.Init(sim)
	return cast
}

func (shaman *Shaman) NewGraceOfAirTotem(sim *core.Simulation) *core.SimpleCast {
	baseManaCost := 310.0
	manaCost := baseManaCost * (1 - float64(shaman.Talents.TotemicFocus)*0.05)
	manaCost -= baseManaCost * float64(shaman.Talents.MentalQuickness) * 0.02

	cast := &core.SimpleCast{
		Cast: core.Cast{
			Name:         "Grace Of Air",
			ActionID:     core.ActionID{SpellID: 25359},
			Character:    shaman.GetCharacter(),
			BaseManaCost: baseManaCost,
			ManaCost:     manaCost,
			GCDCooldown:  time.Second * 1,
		},
	}
	cast.Init(sim)
	return cast
}

func (shaman *Shaman) NewWindfuryTotem(sim *core.Simulation) *core.SimpleCast {
	baseManaCost := 325.0
	manaCost := baseManaCost * (1 - float64(shaman.Talents.TotemicFocus)*0.05)
	manaCost -= baseManaCost * float64(shaman.Talents.MentalQuickness) * 0.02

	cast := &core.SimpleCast{
		Cast: core.Cast{
			Name:         "Windfury Totem",
			ActionID:     core.ActionID{SpellID: 25587},
			Character:    shaman.GetCharacter(),
			BaseManaCost: baseManaCost,
			ManaCost:     manaCost,
			GCDCooldown:  time.Second * 1,
		},
	}
	cast.Init(sim)
	return cast
}

func (shaman *Shaman) NewWaterTotem(sim *core.Simulation) *core.SimpleCast {
	baseManaCost := 120.0
	manaCost := baseManaCost * (1 - float64(shaman.Talents.TotemicFocus)*0.05)
	manaCost -= baseManaCost * float64(shaman.Talents.MentalQuickness) * 0.02

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
	manaCost -= baseManaCost * float64(shaman.Talents.MentalQuickness) * 0.02

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

func (shaman *Shaman) NewEarthTotem(sim *core.Simulation) *core.SimpleCast {
	baseManaCost := 125.0
	manaCost := baseManaCost * (1 - float64(shaman.Talents.TotemicFocus)*0.05)
	manaCost -= baseManaCost * float64(shaman.Talents.MentalQuickness) * 0.02

	cast := &core.SimpleCast{
		Cast: core.Cast{
			Name:         "Strength of Earth",
			ActionID:     core.ActionID{SpellID: 8161},
			Character:    shaman.GetCharacter(),
			BaseManaCost: baseManaCost,
			ManaCost:     manaCost,
			GCDCooldown:  time.Second * 1,
		},
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
			shaman.SelfBuffs.NextTotemDrops[EarthTotem] = sim.CurrentTime + time.Second*120
		},
	}
	cast.Init(sim)
	return cast
}

// TryDropTotems will check to see if totems need to be re-cast.
//  If they do time.Duration will be returned will be >0.
func (shaman *Shaman) TryDropTotems(sim *core.Simulation) time.Duration {
	gcd := shaman.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime)
	if gcd > 0 {
		return sim.CurrentTime + gcd // can't drop totems in GCD
	}
	var cast *core.SimpleCast

	// currently hardcoded to include 25% mana cost reduction from resto talents
	for totemTypeIdx, totemExpiration := range shaman.SelfBuffs.NextTotemDrops {
		if cast != nil {
			break
		}
		nextDrop := shaman.SelfBuffs.NextTotemDropType[totemTypeIdx]
		if sim.CurrentTime > totemExpiration {
			switch totemTypeIdx {
			case AirTotem:
				switch nextDrop {
				case wrathOfAirTotem:
					cast = shaman.NewWrathOfAirTotem(sim)
				case windfuryTotem:
					cast = shaman.NewWindfuryTotem(sim)
				case graceOfAirTotem:
					cast = shaman.NewGraceOfAirTotem(sim)
				}
				if shaman.SelfBuffs.TwistWindfury {
					if nextDrop != windfuryTotem {
						shaman.SelfBuffs.NextTotemDropType[AirTotem] = windfuryTotem
						shaman.SelfBuffs.NextTotemDrops[AirTotem] = sim.CurrentTime + time.Second*10 // 10s until you need to drop WF
					} else {
						shaman.SelfBuffs.NextTotemDropType[AirTotem] = int32(shaman.SelfBuffs.AirTotem)
						shaman.SelfBuffs.NextTotemDrops[AirTotem] = sim.CurrentTime + time.Second // drop immediately
					}
				}
			case EarthTotem:
				cast = shaman.NewEarthTotem(sim)
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
