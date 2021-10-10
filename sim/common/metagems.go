package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	core.AddActiveItem(34220, core.ActiveItem{Activate: ActivateCSD, ActivateCD: core.NeverExpires})
	core.AddActiveItem(35503, core.ActiveItem{Activate: ActivateESD, ActivateCD: core.NeverExpires})
	core.AddActiveItem(25901, core.ActiveItem{Activate: ActivateIED, ActivateCD: core.NeverExpires})
	core.AddActiveItem(25893, core.ActiveItem{Activate: ActivateMSD, ActivateCD: core.NeverExpires})
}

func ActivateCSD(sim *core.Simulation, agent core.Agent) core.Aura {
	return core.Aura{
		ID:      core.MagicIDChaoticSkyfire,
		Expires: core.NeverExpires,
		OnCast: func(sim *core.Simulation, cast core.DirectCastAction, input *core.DirectCastInput) {
			// For a normal spell with crit multiplier of 1.5, this will be 1.
			// For a spell with a multiplier of 2 (i.e. 100% increased critical damage) this will be 2.
			improvedCritRatio := (input.CritMultiplier - 1) / 0.5

			input.CritMultiplier += 0.045 * improvedCritRatio
		},
	}
}

func ActivateESD(sim *core.Simulation, agent core.Agent) core.Aura {
	// FUTURE: this technically should be modified by blessing of kings?
	agent.GetCharacter().Stats[stats.Intellect] += agent.GetCharacter().Stats[stats.Intellect] * 0.02
	return core.Aura{
		ID:      core.MagicIDEmberSkyfire,
		Expires: core.NeverExpires,
	}
}

func ActivateIED(sim *core.Simulation, agent core.Agent) core.Aura {
	icd := core.NewICD()
	const dur = time.Second * 15
	return core.Aura{
		ID:      core.MagicIDInsightfulEarthstorm,
		Expires: core.NeverExpires,
		OnCastComplete: func(sim *core.Simulation, cast core.DirectCastAction) {
			if !icd.IsOnCD(sim) && sim.Rando.Float64("unmarked") < 0.04 {
				icd = core.InternalCD(sim.CurrentTime + dur)
				if sim.Log != nil {
					sim.Log(" *Insightful Earthstorm Mana Restore - 300\n")
				}
				agent.GetCharacter().Stats[stats.Mana] += 300
			}
		},
	}
}

func ActivateMSD(sim *core.Simulation, agent core.Agent) core.Aura {
	const hasteBonus = 320.0
	const dur = time.Second * 4
	const icdDur = time.Second * 35
	icd := core.NewICD()
	return core.Aura{
		ID:      core.MagicIDMysticSkyfire,
		Expires: core.NeverExpires,
		OnCastComplete: func(sim *core.Simulation, cast core.DirectCastAction) {
			if !icd.IsOnCD(sim) && sim.Rando.Float64("unmarked") < 0.15 {
				icd = core.InternalCD(sim.CurrentTime + icdDur)
				core.AddAuraWithTemporaryStats(sim, agent, core.MagicIDMysticFocus, stats.SpellHaste, hasteBonus, dur)
			}
		},
	}
}
