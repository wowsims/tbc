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

func ActivateSkycall(sim *core.Simulation, party *core.Party, player core.PlayerAgent) core.Aura {
	const hasteBonus = 101
	const dur = time.Second * 10
	return core.Aura{
		ID:      core.MagicIDSkycall,
		Expires: core.NeverExpires,
		OnCastComplete: func(sim *core.Simulation, player core.PlayerAgent, c *core.Cast) {
			if c.Spell.ID == core.MagicIDLB12 && sim.Rando.Float64("skycall") < 0.15 {
				player.Stats[stats.SpellHaste] += hasteBonus
				player.AddAura(sim, core.Aura{
					ID:      core.MagicIDEnergized,
					Expires: sim.CurrentTime + dur,
					OnExpire: func(sim *core.Simulation, player core.PlayerAgent, c *core.Cast) {
						player.Stats[stats.SpellHaste] -= hasteBonus
					},
				})
			}
		},
	}
}

func ActivateNAC(sim *core.Simulation, party *core.Party, player core.PlayerAgent) core.Aura {
	const sp = 250
	player.Stats[stats.SpellPower] += sp
	return core.Aura{
		ID:      core.MagicIDNAC,
		Expires: sim.CurrentTime + time.Second*20,
		OnCast: func(sim *core.Simulation, p core.PlayerAgent, c *core.Cast) {
			c.ManaCost *= 1.2
		},
		OnExpire: func(sim *core.Simulation, player core.PlayerAgent, c *core.Cast) {
			player.Stats[stats.SpellPower] -= sp
		},
	}
}

func ActivateTotemOfPulsingEarth(sim *core.Simulation, party *core.Party, player core.PlayerAgent) core.Aura {
	return core.Aura{
		ID:      core.MagicIDTotemOfPulsingEarth,
		Expires: core.NeverExpires,
		OnCast: func(sim *core.Simulation, p core.PlayerAgent, c *core.Cast) {
			if c.Spell.ID == core.MagicIDLB12 {
				// TODO: how to make sure this goes in before clearcasting?
				c.ManaCost = math.Max(c.ManaCost-27, 0)
			}
		},
	}
}
