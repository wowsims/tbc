package shaman

import (
	"math"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	core.AddActiveItem(33506, core.ActiveItem{Activate: ActivateSkycall, ActivateCD: core.NeverExpires})
	core.AddActiveItem(29389, core.ActiveItem{Activate: ActivateTotemOfPulsingEarth, ActivateCD: core.NeverExpires})
	core.AddActiveItem(19344, core.ActiveItem{Activate: ActivateNAC, ActivateCD: time.Second * 300, CoolID: core.MagicIDNACTrink, SharedID: core.MagicIDAtkTrinket})

	// core.AddItemSet()
}

func ActivateSkycall(sim *core.Simulation, agent core.Agent) core.Aura {
	const hasteBonus = 101
	const dur = time.Second * 10
	return core.Aura{
		ID:      core.MagicIDSkycall,
		Expires: core.NeverExpires,
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
			if cast.Spell.ID == MagicIDLB12 && sim.Rando.Float64("skycall") < 0.15 {
				core.AddAuraWithTemporaryStats(sim, agent, core.MagicIDEnergized, stats.SpellHaste, hasteBonus, dur)
			}
		},
	}
}

func ActivateNAC(sim *core.Simulation, agent core.Agent) core.Aura {
	const sp = 250
	agent.GetCharacter().Stats[stats.SpellPower] += sp
	return core.Aura{
		ID:      core.MagicIDNAC,
		Expires: sim.CurrentTime + time.Second*20,
		OnCast: func(sim *core.Simulation, cast *core.Cast) {
			cast.ManaCost *= 1.2
		},
		OnExpire: func(sim *core.Simulation, cast *core.Cast) {
			agent.GetCharacter().Stats[stats.SpellPower] -= sp
		},
	}
}

func ActivateTotemOfPulsingEarth(sim *core.Simulation, agent core.Agent) core.Aura {
	return core.Aura{
		ID:      core.MagicIDTotemOfPulsingEarth,
		Expires: core.NeverExpires,
		OnCast: func(sim *core.Simulation, cast *core.Cast) {
			if cast.Spell.ID == MagicIDLB12 {
				// TODO: how to make sure this goes in before clearcasting?
				cast.ManaCost = math.Max(cast.ManaCost-27, 0)
			}
		},
	}
}

func ActivateCycloneManaReduce(sim *core.Simulation, agent core.Agent) core.Aura {
	return core.Aura{
		ID:      core.MagicIDCyclone4pc,
		Expires: core.NeverExpires,
		OnSpellHit: func(sim *core.Simulation, cast *core.Cast) {
			if cast.DidCrit && sim.Rando.Float64("unmarked") < 0.11 {
				agent.GetCharacter().AddAura(sim, core.Aura{
					ID: core.MagicIDCycloneMana,
					OnCast: func(sim *core.Simulation, cast *core.Cast) {
						// TODO: how to make sure this goes in before clearcasting?
						cast.ManaCost -= 270
					},
					OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
						agent.GetCharacter().RemoveAura(sim, core.MagicIDCycloneMana)
					},
				})
			}
		},
	}
}

func ActivateCataclysmLBDiscount(sim *core.Simulation, agent core.Agent) core.Aura {
	return core.Aura{
		ID:      core.MagicIDCataclysm4pc,
		Expires: core.NeverExpires,
		OnSpellHit: func(sim *core.Simulation, cast *core.Cast) {
			if cast.DidCrit && sim.Rando.Float64("unmarked") < 0.25 {
				agent.GetCharacter().Stats[stats.Mana] += 120
			}
		},
	}
}

func ActivateSkyshatterImpLB(sim *core.Simulation, agent core.Agent) core.Aura {
	return core.Aura{
		ID:      core.MagicIDSkyshatter4pc,
		Expires: core.NeverExpires,
		OnSpellHit: func(sim *core.Simulation, cast *core.Cast) {
			if cast.Spell.ID == MagicIDLB12 {
				cast.DidDmg *= 1.05
			}
		},
	}
}
