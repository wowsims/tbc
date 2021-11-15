package balance

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/druid"
)

// This is in here because for now we need access to the agent type (BalanceDruid) and so this has to be in this package.
//  Once we make a better way to access DoTs on a target we can move this back to normal druid package.

func init() {
	core.AddItemSet(ItemSetNordrassil)
}

var Nordrassil4pAuraID = core.NewAuraID()

var ItemSetNordrassil = core.ItemSet{
	Name:  "Nordrassil Regalia",
	Items: map[int32]struct{}{30231: {}, 30232: {}, 30233: {}, 30234: {}, 30235: {}},
	Bonuses: map[int32]core.ApplyEffect{
		4: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				return core.Aura{
					ID:   Nordrassil4pAuraID,
					Name: "Nordrassil 4p Bonus",
					OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
						moonkin, ok := agent.(*BalanceDruid)
						if !ok {
							panic("why is a non-moonkin using nordassil regalia")
						}
						if spellCast.ActionID.SpellID == druid.SpellIDSF8 || spellCast.ActionID.SpellID == druid.SpellIDSF6 {
							// Check if moonfire/insectswarm is ticking on the target.
							// TODO: in a raid simulator we need to be able to see which dots are ticking from other druids.
							if (moonkin.MoonfireSpell.DotInput.IsTicking(sim) && moonkin.MoonfireSpell.Target.Index == spellEffect.Target.Index) ||
								(moonkin.InsectSwarmSpell.DotInput.IsTicking(sim) && moonkin.InsectSwarmSpell.Target.Index == spellEffect.Target.Index) {
								spellEffect.DamageMultiplier *= 1.1
							}
						}
					},
				}
			})
		},
	},
}
