package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	// Keep these in order by item ID.
	core.AddItemEffect(25893, ApplyMysticalSkyfireDiamond)
	core.AddItemEffect(25901, ApplyInsightfulEarthstormDiamond)
	core.AddItemEffect(34220, ApplyChaoticSkyfireDiamond)
	core.AddItemEffect(35503, ApplyEmberSkyfireDiamond)
}

func ApplyMysticalSkyfireDiamond(agent core.Agent) {
	agent.GetCharacter().AddPermanentAura(func(sim *core.Simulation, character *core.Character) core.Aura {
		const hasteBonus = 320.0
		const dur = time.Second * 4
		const icdDur = time.Second * 35
		icd := core.NewICD()

		return core.Aura{
			ID:      core.MagicIDMysticSkyfire,
			OnCastComplete: func(sim *core.Simulation, cast core.DirectCastAction) {
				if !icd.IsOnCD(sim) && sim.Rando.Float64("unmarked") < 0.15 {
					icd = core.InternalCD(sim.CurrentTime + icdDur)
					core.AddAuraWithTemporaryStats(sim, character, core.MagicIDMysticFocus, stats.SpellHaste, hasteBonus, dur)
				}
			},
		}
	})
}

func ApplyInsightfulEarthstormDiamond(agent core.Agent) {
	agent.GetCharacter().AddPermanentAura(func(sim *core.Simulation, character *core.Character) core.Aura {
		icd := core.NewICD()
		const dur = time.Second * 15

		return core.Aura{
			ID:      core.MagicIDInsightfulEarthstorm,
			OnCastComplete: func(sim *core.Simulation, cast core.DirectCastAction) {
				if !icd.IsOnCD(sim) && sim.Rando.Float64("unmarked") < 0.04 {
					icd = core.InternalCD(sim.CurrentTime + dur)
					if sim.Log != nil {
						sim.Log(" *Insightful Earthstorm Mana Restore - 300\n")
					}
					character.AddStat(stats.Mana, 300)
				}
			},
		}
	})
}

func ApplyChaoticSkyfireDiamond(agent core.Agent) {
	agent.GetCharacter().AddPermanentAura(func(sim *core.Simulation, character *core.Character) core.Aura {
		return core.Aura{
			ID:      core.MagicIDChaoticSkyfire,
			OnCast: func(sim *core.Simulation, cast core.DirectCastAction, input *core.DirectCastInput) {
				// For a normal spell with crit multiplier of 1.5, this will be 1.
				// For a spell with a multiplier of 2 (i.e. 100% increased critical damage) this will be 2.
				improvedCritRatio := (input.CritMultiplier - 1) / 0.5

				input.CritMultiplier += 0.045 * improvedCritRatio
			},
		}
	})
}

func ApplyEmberSkyfireDiamond(agent core.Agent) {
	agent.GetCharacter().AddStatDependency(stats.StatDependency{
		SourceStat: stats.Intellect,
		ModifiedStat: stats.Intellect,
		Modifier: func(intellect float64, _ float64) float64 {
			return intellect * 1.02
		},
	})
}
