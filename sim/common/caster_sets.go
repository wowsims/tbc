package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Keep these (and their functions) in alphabetical order.
func init() {
	core.AddItemSet(&ItemSetManaEtched)
	core.AddItemSet(&ItemSetNetherstrike)
	core.AddItemSet(&ItemSetSpellstrike)
	core.AddItemSet(&ItemSetTheTwinStars)
	core.AddItemSet(&ItemSetWindhawk)
	core.AddItemSet(&ItemSetSpellfire)
}

var ManaEtchedAuraID = core.NewAuraID()
var ManaEtchedInsightAuraID = core.NewAuraID()
var ItemSetManaEtched = core.ItemSet{
	Name: "Mana-Etched Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.SpellHit, 35)
		},
		4: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				const spellBonus = 110.0
				const duration = time.Second * 15
				applyStatAura := character.NewTemporaryStatsAuraApplier(ManaEtchedInsightAuraID, core.ActionID{SpellID: 37619}, stats.Stats{stats.SpellPower: spellBonus}, duration)
				return core.Aura{
					ID: ManaEtchedAuraID,
					OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
						if sim.RandomFloat("Mana-Etched Insight") > 0.02 {
							return
						}
						applyStatAura(sim)
					},
				}
			})
		},
	},
}

var ItemSetNetherstrike = core.ItemSet{
	Name: "Netherstrike Armor",
	Bonuses: map[int32]core.ApplyEffect{
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.SpellPower, 23)
		},
	},
}

var SpellstrikeAuraID = core.NewAuraID()
var SpellstrikeInfusionAuraID = core.NewAuraID()
var ItemSetSpellstrike = core.ItemSet{
	Name: "Spellstrike Infusion",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				const spellBonus = 92.0
				const duration = time.Second * 10
				applyStatAura := character.NewTemporaryStatsAuraApplier(SpellstrikeInfusionAuraID, core.ActionID{SpellID: 32106}, stats.Stats{stats.SpellPower: spellBonus}, duration)

				return core.Aura{
					ID: SpellstrikeAuraID,
					OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
						if sim.RandomFloat("spellstrike") > 0.05 {
							return
						}
						applyStatAura(sim)
					},
				}
			})
		},
	},
}

var ItemSetTheTwinStars = core.ItemSet{
	Name: "The Twin Stars",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.SpellPower, 15)
		},
	},
}

var ItemSetWindhawk = core.ItemSet{
	Name: "Windhawk Armor",
	Bonuses: map[int32]core.ApplyEffect{
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.MP5, 8)
		},
	},
}

var ItemSetSpellfire = core.ItemSet{
	Name: "Wrath of Spellfire",
	Bonuses: map[int32]core.ApplyEffect{
		3: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStatDependency(stats.StatDependency{
				SourceStat:   stats.Intellect,
				ModifiedStat: stats.SpellPower,
				Modifier: func(intellect float64, spellPower float64) float64 {
					return spellPower + intellect*0.07 // 7% bonus to sp from int
				},
			})
		},
	},
}
