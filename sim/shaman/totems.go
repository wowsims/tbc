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

func (shaman *Shaman) newTotemCastTemplate(sim *core.Simulation, baseManaCost float64, spellID int32) core.SimpleCast {
	manaCost := baseManaCost * (1 - float64(shaman.Talents.TotemicFocus)*0.05)
	manaCost -= baseManaCost * float64(shaman.Talents.MentalQuickness) * 0.02

	template := core.SimpleCast{
		Cast: core.Cast{
			ActionID:     core.ActionID{SpellID: spellID},
			Character:    shaman.GetCharacter(),
			BaseManaCost: baseManaCost,
			ManaCost:     manaCost,
			GCDCooldown:  time.Second * 1,
		},
	}

	return template
}

func (shaman *Shaman) newWrathOfAirTotemTemplate(sim *core.Simulation) core.SimpleCast {
	cast := shaman.newTotemCastTemplate(sim, 320.0, 3738)
	cast.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		shaman.SelfBuffs.NextTotemDrops[AirTotem] = sim.CurrentTime + time.Second*120
		shaman.tryTwistWindfury(sim)
	}
	return cast
}

// Totems that shaman will cast.
func (shaman *Shaman) NewWrathOfAirTotem(sim *core.Simulation) *core.SimpleCast {
	shaman.totemSpell = shaman.wrathOfAirTotemTemplate
	shaman.totemSpell.Init(sim)
	return &shaman.totemSpell
}

func (shaman *Shaman) newGraceOfAirTotemTemplate(sim *core.Simulation) core.SimpleCast {
	cast := shaman.newTotemCastTemplate(sim, 310.0, 25359)
	cast.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		shaman.SelfBuffs.NextTotemDrops[AirTotem] = sim.CurrentTime + time.Second*120
		shaman.tryTwistWindfury(sim)
	}
	return cast
}

func (shaman *Shaman) NewGraceOfAirTotem(sim *core.Simulation) *core.SimpleCast {
	shaman.totemSpell = shaman.graceOfAirTotemTemplate
	shaman.totemSpell.Init(sim)
	return &shaman.totemSpell
}

var windfuryBaseManaCosts = []float64{
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

	baseManaCost := windfuryBaseManaCosts[rank-1]
	spellID := core.WindfurySpellRanks[rank-1]
	cast := shaman.newTotemCastTemplate(sim, baseManaCost, spellID)
	cast.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		shaman.SelfBuffs.NextTotemDrops[AirTotem] = sim.CurrentTime + time.Second*120
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

func (shaman *Shaman) newManaSpringTotemTemplate(sim *core.Simulation) core.SimpleCast {
	cast := shaman.newTotemCastTemplate(sim, 120, 25570)
	cast.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		shaman.SelfBuffs.NextTotemDrops[WaterTotem] = sim.CurrentTime + time.Second*120
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
		shaman.SelfBuffs.NextTotemDrops[FireTotem] = sim.CurrentTime + time.Second*120
	}
	return cast
}

func (shaman *Shaman) NewTotemOfWrath(sim *core.Simulation) *core.SimpleCast {
	shaman.totemSpell = shaman.totemOfWrathTemplate
	shaman.totemSpell.Init(sim)
	return &shaman.totemSpell
}

func (shaman *Shaman) newStrengthOfEarthTotemTemplate(sim *core.Simulation) core.SimpleCast {
	cast := shaman.newTotemCastTemplate(sim, 125, 8161)
	cast.OnCastComplete = func(sim *core.Simulation, cast *core.Cast) {
		shaman.SelfBuffs.NextTotemDrops[EarthTotem] = sim.CurrentTime + time.Second*120
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
		shaman.SelfBuffs.NextTotemDrops[EarthTotem] = sim.CurrentTime + time.Second*120
	}
	return cast
}

func (shaman *Shaman) NewTremorTotem(sim *core.Simulation) *core.SimpleCast {
	shaman.totemSpell = shaman.tremorTotemTemplate
	shaman.totemSpell.Init(sim)
	return &shaman.totemSpell
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
				cast = shaman.NewManaSpringTotem(sim)
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
