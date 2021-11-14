package druid

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	core.AddItemSet(ItemSetMalorne)
}

var Malorne2PcAuraID = core.NewAuraID()
var ItemSetMalorne = core.ItemSet{
	Name:  "Malorne Rainment",
	Items: map[int32]struct{}{29033: {}, 29035: {}, 29034: {}, 29036: {}, 29037: {}},
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
	},
}
