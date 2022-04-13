package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (shaman *Shaman) newTotemCastTemplate(sim *core.Simulation, baseManaCost float64, spellID int32) core.SimpleCast {
	manaCost := baseManaCost * (1 - float64(shaman.Talents.TotemicFocus)*0.05)
	manaCost -= baseManaCost * float64(shaman.Talents.MentalQuickness) * 0.02

	template := core.SimpleCast{
		Cast: core.Cast{
			ActionID:  core.ActionID{SpellID: spellID},
			Character: shaman.GetCharacter(),
			BaseCost: core.ResourceCost{
				Type:  stats.Mana,
				Value: baseManaCost,
			},
			Cost: core.ResourceCost{
				Type:  stats.Mana,
				Value: manaCost,
			},
			GCD:         time.Second,
			SpellExtras: SpellFlagTotem,
		},
	}

	return template
}

func (shaman *Shaman) newWrathOfAirTotemTemplate(sim *core.Simulation) core.SimpleCast {
	cast := shaman.newTotemCastTemplate(sim, 320.0, 3738)
	cast.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		shaman.NextTotemDrops[AirTotem] = sim.CurrentTime + time.Second*120
		shaman.tryTwistWindfury(sim)
	}
	return cast
}

func (shaman *Shaman) NewWrathOfAirTotem(sim *core.Simulation) *core.SimpleCast {
	shaman.totemSpell = shaman.wrathOfAirTotemTemplate
	shaman.totemSpell.Init(sim)
	return &shaman.totemSpell
}

func (shaman *Shaman) newGraceOfAirTotemTemplate(sim *core.Simulation) core.SimpleCast {
	cast := shaman.newTotemCastTemplate(sim, 310.0, 25359)
	cast.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		shaman.NextTotemDrops[AirTotem] = sim.CurrentTime + time.Second*120
		shaman.tryTwistWindfury(sim)
	}
	return cast
}

func (shaman *Shaman) NewGraceOfAirTotem(sim *core.Simulation) *core.SimpleCast {
	shaman.totemSpell = shaman.graceOfAirTotemTemplate
	shaman.totemSpell.Init(sim)
	return &shaman.totemSpell
}

func (shaman *Shaman) newTranquilAirTotemTemplate(sim *core.Simulation) core.SimpleCast {
	baseManaCost := shaman.BaseMana() * 0.06
	cast := shaman.newTotemCastTemplate(sim, baseManaCost, 25908)
	cast.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		shaman.NextTotemDrops[AirTotem] = sim.CurrentTime + time.Second*120
		shaman.tryTwistWindfury(sim)
	}
	return cast
}

func (shaman *Shaman) NewTranquilAirTotem(sim *core.Simulation) *core.SimpleCast {
	shaman.totemSpell = shaman.tranquilAirTotemTemplate
	shaman.totemSpell.Init(sim)
	return &shaman.totemSpell
}

var windfuryTotemBaseManaCosts = []float64{
	95,
	140,
	200,
	275,
	325,
}

func (shaman *Shaman) newWindfuryTotemTemplate(sim *core.Simulation, rank int32) core.SimpleCast {
	if rank == 0 {
		// This will happen if we're not casting windfury totem. Just return a rank 1
		// template so we don't error.
		rank = 1
	}

	baseManaCost := windfuryTotemBaseManaCosts[rank-1]
	spellID := core.WindfuryTotemSpellRanks[rank-1]
	cast := shaman.newTotemCastTemplate(sim, baseManaCost, spellID)
	cast.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		shaman.NextTotemDrops[AirTotem] = sim.CurrentTime + time.Second*120
		shaman.tryTwistWindfury(sim)
	}
	return cast
}

func (shaman *Shaman) NewWindfuryTotem(sim *core.Simulation) *core.SimpleCast {
	shaman.totemSpell = shaman.windfuryTotemTemplate
	shaman.totemSpell.Init(sim)
	return &shaman.totemSpell
}

func (shaman *Shaman) tryTwistWindfury(sim *core.Simulation) {
	if !shaman.Totems.TwistWindfury {
		return
	}

	if shaman.Metrics.WentOOM && shaman.CurrentManaPercent() < 0.2 {
		shaman.NextTotemDropType[AirTotem] = int32(shaman.Totems.Air)
		return
	}

	// Swap to WF if we didn't just cast it, otherwise drop the other air totem immediately.
	if shaman.NextTotemDropType[AirTotem] != int32(proto.AirTotem_WindfuryTotem) {
		shaman.NextTotemDropType[AirTotem] = int32(proto.AirTotem_WindfuryTotem)
		shaman.NextTotemDrops[AirTotem] = sim.CurrentTime + time.Second*10 // 10s until you need to drop WF
	} else {
		shaman.NextTotemDropType[AirTotem] = int32(shaman.Totems.Air)
		shaman.NextTotemDrops[AirTotem] = sim.CurrentTime + time.Second // drop immediately
	}
}

func (shaman *Shaman) tryTwistFireNova(sim *core.Simulation) {
	if !shaman.Totems.TwistFireNova {
		return
	}

	if shaman.Metrics.WentOOM && shaman.CurrentManaPercent() < 0.2 {
		shaman.NextTotemDropType[FireTotem] = int32(shaman.Totems.Fire)
		return
	}

	if shaman.NextTotemDropType[FireTotem] != int32(proto.FireTotem_FireNovaTotem) ||
		shaman.Totems.Fire == proto.FireTotem_NoFireTotem {
		shaman.NextTotemDropType[FireTotem] = int32(proto.FireTotem_FireNovaTotem)
		shaman.NextTotemDrops[FireTotem] = sim.CurrentTime + shaman.GetRemainingCD(CooldownIDNovaTotem, sim.CurrentTime)
	} else {
		shaman.NextTotemDropType[FireTotem] = int32(shaman.Totems.Fire)
	}
}

func (shaman *Shaman) newManaSpringTotemTemplate(sim *core.Simulation) core.SimpleCast {
	cast := shaman.newTotemCastTemplate(sim, 120, 25570)
	cast.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		shaman.NextTotemDrops[WaterTotem] = sim.CurrentTime + time.Second*120
	}
	return cast
}

func (shaman *Shaman) NewManaSpringTotem(sim *core.Simulation) *core.SimpleCast {
	shaman.totemSpell = shaman.manaSpringTotemTemplate
	shaman.totemSpell.Init(sim)
	return &shaman.totemSpell
}

func (shaman *Shaman) newTotemOfWrathTemplate(sim *core.Simulation) core.SimpleCast {
	baseManaCost := shaman.BaseMana() * 0.05
	cast := shaman.newTotemCastTemplate(sim, baseManaCost, 30706)
	cast.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		shaman.NextTotemDrops[FireTotem] = sim.CurrentTime + time.Second*120
		shaman.tryTwistFireNova(sim)
	}
	return cast
}

func (shaman *Shaman) NewTotemOfWrath(sim *core.Simulation) *core.SimpleCast {
	shaman.totemSpell = shaman.totemOfWrathTemplate
	shaman.totemSpell.Init(sim)
	return &shaman.totemSpell
}

func (shaman *Shaman) newStrengthOfEarthTotemTemplate(sim *core.Simulation) core.SimpleCast {
	cast := shaman.newTotemCastTemplate(sim, 300, 25528)
	cast.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		shaman.NextTotemDrops[EarthTotem] = sim.CurrentTime + time.Second*120
	}
	return cast
}

func (shaman *Shaman) NewStrengthOfEarthTotem(sim *core.Simulation) *core.SimpleCast {
	shaman.totemSpell = shaman.strengthOfEarthTotemTemplate
	shaman.totemSpell.Init(sim)
	return &shaman.totemSpell
}

func (shaman *Shaman) newTremorTotemTemplate(sim *core.Simulation) core.SimpleCast {
	cast := shaman.newTotemCastTemplate(sim, 60, 8143)
	cast.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		shaman.NextTotemDrops[EarthTotem] = sim.CurrentTime + time.Second*120
	}
	return cast
}

func (shaman *Shaman) NewTremorTotem(sim *core.Simulation) *core.SimpleCast {
	shaman.totemSpell = shaman.tremorTotemTemplate
	shaman.totemSpell.Init(sim)
	return &shaman.totemSpell
}

func (shaman *Shaman) NextTotemAt(sim *core.Simulation) time.Duration {
	nextTotemAt := core.MinDuration(
		shaman.NextTotemDrops[0],
		core.MinDuration(
			shaman.NextTotemDrops[1],
			core.MinDuration(
				shaman.NextTotemDrops[2],
				shaman.NextTotemDrops[3])))

	return nextTotemAt
}

// TryDropTotems will check to see if totems need to be re-cast.
//  Returns whether we tried to cast a totem, regardless of whether it succeeded.
func (shaman *Shaman) TryDropTotems(sim *core.Simulation) bool {
	var cast *core.SimpleCast
	var attackCast *core.Spell // if using fire totems this will be an attack cast.

	for totemTypeIdx, totemExpiration := range shaman.NextTotemDrops {
		if cast != nil || attackCast != nil {
			break
		}
		nextDrop := shaman.NextTotemDropType[totemTypeIdx]
		if sim.CurrentTime >= totemExpiration {
			switch totemTypeIdx {
			case AirTotem:
				switch proto.AirTotem(nextDrop) {
				case proto.AirTotem_WrathOfAirTotem:
					cast = shaman.NewWrathOfAirTotem(sim)
				case proto.AirTotem_WindfuryTotem:
					cast = shaman.NewWindfuryTotem(sim)
				case proto.AirTotem_GraceOfAirTotem:
					cast = shaman.NewGraceOfAirTotem(sim)
				case proto.AirTotem_TranquilAirTotem:
					cast = shaman.NewTranquilAirTotem(sim)
				}

			case EarthTotem:
				switch proto.EarthTotem(nextDrop) {
				case proto.EarthTotem_StrengthOfEarthTotem:
					cast = shaman.NewStrengthOfEarthTotem(sim)
				case proto.EarthTotem_TremorTotem:
					cast = shaman.NewTremorTotem(sim)
				}

			case FireTotem:
				switch proto.FireTotem(nextDrop) {
				case proto.FireTotem_TotemOfWrath:
					cast = shaman.NewTotemOfWrath(sim)
				case proto.FireTotem_SearingTotem:
					attackCast = shaman.SearingTotem
				case proto.FireTotem_MagmaTotem:
					attackCast = shaman.MagmaTotem
				case proto.FireTotem_FireNovaTotem:
					attackCast = shaman.FireNovaTotem
				}

			case WaterTotem:
				cast = shaman.NewManaSpringTotem(sim)
			}
		}
	}

	if cast != nil {
		if success := cast.StartCast(sim); !success {
			shaman.WaitForMana(sim, cast.Cost.Value)
		}
		return true
	} else if attackCast != nil {
		if success := attackCast.Cast(sim, sim.GetPrimaryTarget()); !success {
			shaman.WaitForMana(sim, attackCast.CurCast.Cost)
		}
		return true
	}
	return false
}
