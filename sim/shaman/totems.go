package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

const windfuryTotem = int32(proto.AirTotem_WindfuryTotem)
const wrathOfAirTotem = int32(proto.AirTotem_WrathOfAirTotem)
const graceOfAirTotem = int32(proto.AirTotem_GraceOfAirTotem)

const totemOfWrath = int32(proto.FireTotem_TotemOfWrath)
const searingTotem = int32(proto.FireTotem_SearingTotem)
const magmaTotem = int32(proto.FireTotem_MagmaTotem)
const novaTotem = int32(proto.FireTotem_FireNovaTotem)

const strengthOfEarthTotem = int32(proto.EarthTotem_StrengthOfEarthTotem)
const tremorTotem = int32(proto.EarthTotem_TremorTotem)

// Totems that shaman will cast.
func (shaman *Shaman) NewWrathOfAirTotem(sim *core.Simulation) *core.SimpleCast {
	baseManaCost := 320.0
	manaCost := baseManaCost * (1 - float64(shaman.Talents.TotemicFocus)*0.05)
	manaCost -= baseManaCost * float64(shaman.Talents.MentalQuickness) * 0.02

	cast := &core.SimpleCast{
		Cast: core.Cast{
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

func (shaman *Shaman) NewGraceOfAirTotem(sim *core.Simulation) *core.SimpleCast {
	baseManaCost := 310.0
	manaCost := baseManaCost * (1 - float64(shaman.Talents.TotemicFocus)*0.05)
	manaCost -= baseManaCost * float64(shaman.Talents.MentalQuickness) * 0.02

	cast := &core.SimpleCast{
		Cast: core.Cast{
			ActionID:     core.ActionID{SpellID: 25359},
			Character:    shaman.GetCharacter(),
			BaseManaCost: baseManaCost,
			ManaCost:     manaCost,
			GCDCooldown:  time.Second * 1,
		},
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
			shaman.SelfBuffs.NextTotemDrops[AirTotem] = sim.CurrentTime + time.Second*120
			shaman.tryTwistWindfury(sim)
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
			ActionID:     core.ActionID{SpellID: 25587},
			Character:    shaman.GetCharacter(),
			BaseManaCost: baseManaCost,
			ManaCost:     manaCost,
			GCDCooldown:  time.Second * 1,
		},
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
			shaman.SelfBuffs.NextTotemDrops[AirTotem] = sim.CurrentTime + time.Second*120
			shaman.tryTwistWindfury(sim)
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

func (shaman *Shaman) NewTotemOfWrath(sim *core.Simulation) *core.SimpleCast {
	baseManaCost := shaman.BaseMana() * 0.05
	manaCost := baseManaCost * (1 - float64(shaman.Talents.TotemicFocus)*0.05)
	manaCost -= baseManaCost * float64(shaman.Talents.MentalQuickness) * 0.02

	cast := &core.SimpleCast{
		Cast: core.Cast{
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

func (shaman *Shaman) NewStrengthOfEarthTotem(sim *core.Simulation) *core.SimpleCast {
	baseManaCost := 125.0
	manaCost := baseManaCost * (1 - float64(shaman.Talents.TotemicFocus)*0.05)
	manaCost -= baseManaCost * float64(shaman.Talents.MentalQuickness) * 0.02

	cast := &core.SimpleCast{
		Cast: core.Cast{
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
func (shaman *Shaman) NewTremorTotem(sim *core.Simulation) *core.SimpleCast {
	baseManaCost := 60.0
	manaCost := baseManaCost * (1 - float64(shaman.Talents.TotemicFocus)*0.05)
	manaCost -= baseManaCost * float64(shaman.Talents.MentalQuickness) * 0.02

	cast := &core.SimpleCast{
		Cast: core.Cast{
			ActionID:     core.ActionID{SpellID: 8143},
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
func (shaman *Shaman) tryTwistWindfury(sim *core.Simulation) {
	if !shaman.SelfBuffs.TwistWindfury {
		return
	}

	// Swap to WF if we didn't just cast it, otherwise drop the other air totem immediately.
	if shaman.SelfBuffs.NextTotemDropType[AirTotem] != windfuryTotem {
		shaman.SelfBuffs.NextTotemDropType[AirTotem] = windfuryTotem
		shaman.SelfBuffs.NextTotemDrops[AirTotem] = sim.CurrentTime + time.Second*10 // 10s until you need to drop WF
	} else {
		shaman.SelfBuffs.NextTotemDropType[AirTotem] = int32(shaman.SelfBuffs.AirTotem)
		shaman.SelfBuffs.NextTotemDrops[AirTotem] = sim.CurrentTime + time.Second // drop immediately
	}
}

// TryDropTotems will check to see if totems need to be re-cast.
//  If they do time.Duration will be returned will be >0.
func (shaman *Shaman) TryDropTotems(sim *core.Simulation) time.Duration {
	gcd := shaman.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime)
	if gcd > 0 {
		return sim.CurrentTime + gcd // can't drop totems in GCD
	}

	var cast *core.SimpleCast
	var attackCast *core.SimpleSpell // if using fire totems this will be an attack cast.

	// currently hardcoded to include 25% mana cost reduction from resto talents
	for totemTypeIdx, totemExpiration := range shaman.SelfBuffs.NextTotemDrops {
		if cast != nil || attackCast != nil {
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

			case EarthTotem:
				switch nextDrop {
				case strengthOfEarthTotem:
					cast = shaman.NewStrengthOfEarthTotem(sim)
				case tremorTotem:
					cast = shaman.NewTremorTotem(sim)
				}

			case FireTotem:
				switch nextDrop {
				case totemOfWrath:
					cast = shaman.NewTotemOfWrath(sim)
				case searingTotem:
					attackCast = shaman.NewSearingTotem(sim, sim.GetPrimaryTarget())
					shaman.SelfBuffs.NextTotemDrops[FireTotem] = sim.CurrentTime + time.Second*60 + 1
				case magmaTotem:
					attackCast = shaman.NewMagmaTotem(sim)
					shaman.SelfBuffs.NextTotemDrops[FireTotem] = sim.CurrentTime + time.Second*20 + 1
				case novaTotem:
					// If we drop nova while another totem is running, cancel it.
					if shaman.FireTotemSpell.IsInUse() {
						shaman.FireTotemSpell.Cancel(sim)
					}
					attackCast = shaman.NewNovaTotem(sim)
					// use attackCast.DotInput.TickLength as input in case ImprovedFireTotems reduces time to explode
					shaman.SelfBuffs.NextTotemDrops[FireTotem] = sim.CurrentTime + attackCast.Effects[0].DotInput.TickLength + 1
				}

			case WaterTotem:
				cast = shaman.NewWaterTotem(sim)
			}
		}
	}

	success := false
	cost := 0.0
	anyCast := false
	isFireAttack := false
	if cast != nil {
		anyCast = true
		success = cast.StartCast(sim)
		cost = cast.GetManaCost()
	} else if attackCast != nil {
		anyCast = true
		isFireAttack = true
		success = attackCast.Cast(sim)
		cost = attackCast.GetManaCost()
	}

	if !anyCast {
		return 0
	}

	if isFireAttack {
		if shaman.SelfBuffs.TwistFireNova {
			nextDrop := shaman.SelfBuffs.NextTotemDropType[FireTotem]
			if nextDrop != novaTotem {
				shaman.SelfBuffs.NextTotemDropType[FireTotem] = novaTotem
				// place nova as soon as off CD
				shaman.SelfBuffs.NextTotemDrops[FireTotem] = sim.CurrentTime + shaman.GetRemainingCD(CooldownIDNovaTotem, sim.CurrentTime)
			} else {
				shaman.SelfBuffs.NextTotemDropType[FireTotem] = int32(shaman.SelfBuffs.FireTotem)
			}
		}
	}

	if !success {
		regenTime := shaman.TimeUntilManaRegen(cost)
		shaman.Character.Metrics.MarkOOM(sim, &shaman.Character, regenTime)
		return sim.CurrentTime + regenTime
	}
	return sim.CurrentTime + shaman.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime)
}
