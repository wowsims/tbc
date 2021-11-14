package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Keep these (and their functions) in alphabetical order.
func init() {
	core.AddItemSet(ItemSetManaEtched)
	core.AddItemSet(ItemSetNetherstrike)
	core.AddItemSet(ItemSetSpellstrike)
	core.AddItemSet(ItemSetTheTwinStars)
	core.AddItemSet(ItemSetWindhawk)
	core.AddItemSet(ItemSetSpellfire)
}

var ManaEtchedAuraID = core.NewAuraID()
var ManaEtchedInsightAuraID = core.NewAuraID()
var ItemSetManaEtched = core.ItemSet{
	Name:  "Mana-Etched Regalia",
	Items: map[int32]struct{}{28193: {}, 27465: {}, 27907: {}, 27796: {}, 28191: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.SpellHit, 35)
		},
		4: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				const spellBonus = 110.0
				const duration = time.Second * 15

				return core.Aura{
					ID:   ManaEtchedAuraID,
					Name: "Mana-Etched Set",
					OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
						if sim.RandomFloat("unmarked") < 0.02 {
							character.AddAuraWithTemporaryStats(sim, ManaEtchedInsightAuraID, "Mana-Etched Insight", stats.SpellPower, spellBonus, duration)
						}
					},
				}
			})
		},
	},
}

var ItemSetNetherstrike = core.ItemSet{
	Name:  "Netherstrike Armor",
	Items: map[int32]struct{}{29519: {}, 29521: {}, 29520: {}},
	Bonuses: map[int32]core.ApplyEffect{
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.SpellPower, 23)
		},
	},
}

var SpellstrikeAuraID = core.NewAuraID()
var SpellstrikeInfusionAuraID = core.NewAuraID()
var ItemSetSpellstrike = core.ItemSet{
	Name:  "Spellstrike Infusion",
	Items: map[int32]struct{}{24266: {}, 24262: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				const spellBonus = 92.0
				const duration = time.Second * 10

				return core.Aura{
					ID:   SpellstrikeAuraID,
					Name: "Spellstrike Set",
					OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
						if sim.RandomFloat("spellstrike") < 0.05 {
							character.AddAuraWithTemporaryStats(sim, SpellstrikeInfusionAuraID, "Spellstrike Infusion", stats.SpellPower, spellBonus, duration)
						}
					},
				}
			})
		},
	},
}

var ItemSetTheTwinStars = core.ItemSet{
	Name:  "The Twin Stars",
	Items: map[int32]struct{}{31338: {}, 31339: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.SpellPower, 15)
		},
	},
}

var ItemSetWindhawk = core.ItemSet{
	Name:  "Windhawk Armor",
	Items: map[int32]struct{}{29524: {}, 29523: {}, 29522: {}},
	Bonuses: map[int32]core.ApplyEffect{
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.MP5, 8)
		},
	},
}

var ItemSetSpellfire = core.ItemSet{
	Name:  "Spellfire Set",
	Items: map[int32]struct{}{21848: {}, 21847: {}, 21846: {}},
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
