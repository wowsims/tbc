package druid

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	core.AddItemSet(ItemSetMalorne)
}

var Malorne2PcAuraID = core.NewAuraID()
var Malorne4PcAuraID = core.NewAuraID()

var ItemSetMalorne = core.ItemSet{
	Name:  "Malorne Rainment",
	Items: map[int32]struct{}{29093: {}, 29094: {}, 29091: {}, 29092: {}, 29095: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				return core.Aura{
					ID:   Malorne2PcAuraID,
					Name: "Malorne 2pc Bonus",
					OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
						if sim.RandomFloat("malorne 2p") < 0.05 {
							spellCast.Character.AddStat(stats.Mana, 120)
						}
					},
				}
			})
		},
		4: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				return core.Aura{
					ID:   Malorne4PcAuraID,
					Name: "Malorne 4pc Bonus",
					// Currently this is handled in druid.go (reducing CD of innervate)
				}
			})
		},
	},
}
