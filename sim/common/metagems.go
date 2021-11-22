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

var MysticalSkyfireDiamondAuraID = core.NewAuraID()
var MysticFocusAuraID = core.NewAuraID()
func ApplyMysticalSkyfireDiamond(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 320.0
		const dur = time.Second * 4
		const icdDur = time.Second * 35
		icd := core.NewICD()

		return core.Aura{
			ID:      MysticalSkyfireDiamondAuraID,
			Name:    "Mystical Skyfire Diamond",
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				if !icd.IsOnCD(sim) && sim.RandomFloat("Mystical Skyfire Diamond") < 0.15 {
					icd = core.InternalCD(sim.CurrentTime + icdDur)
					character.AddAuraWithTemporaryStats(sim, MysticFocusAuraID, 0, "Mystic Focus", stats.SpellHaste, hasteBonus, dur)
				}
			},
		}
	})
}

var InsightfulEarthstormDiamondAuraID = core.NewAuraID()
func ApplyInsightfulEarthstormDiamond(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		icd := core.NewICD()
		const dur = time.Second * 15

		return core.Aura{
			ID:      InsightfulEarthstormDiamondAuraID,
			Name:    "Insightful Earthstorm Diamond",
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				if !icd.IsOnCD(sim) && sim.RandomFloat("Insightful Earthstorm Diamond") < 0.04 {
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

var ChaoticSkyfireDiamondAuraID = core.NewAuraID()
func ApplyChaoticSkyfireDiamond(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID:      ChaoticSkyfireDiamondAuraID,
			Name:    "Chaotic Skyfire Diamond",
			OnCast: func(sim *core.Simulation, cast *core.Cast) {
				// For a normal spell with crit multiplier of 1.5, this will be 1.
				// For a spell with a multiplier of 2 (i.e. 100% increased critical damage) this will be 2.
				improvedCritRatio := (cast.CritMultiplier - 1) / 0.5

				cast.CritMultiplier += 0.045 * improvedCritRatio
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
