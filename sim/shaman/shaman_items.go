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
}

func ActivateSkycall(sim *core.Simulation, agent core.Agent) core.Aura {
	const hasteBonus = 101
	const dur = time.Second * 10
	return core.Aura{
		ID:      core.MagicIDSkycall,
		Expires: core.NeverExpires,
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
			if cast.Spell.ID == core.MagicIDLB12 && sim.Rando.Float64("skycall") < 0.15 {
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
			if cast.Spell.ID == core.MagicIDLB12 {
				// TODO: how to make sure this goes in before clearcasting?
				cast.ManaCost = math.Max(cast.ManaCost-27, 0)
			}
		},
	}
}
