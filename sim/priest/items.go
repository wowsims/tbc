package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

func init() {
	// core.AddItemEffect(30664, ApplyLivingRootoftheWildheart)

	core.AddItemSet(ItemSetIncarnate)
	core.AddItemSet(ItemSetAvatar)
	core.AddItemSet(ItemSetAbsolution)
}

var ItemSetIncarnate = core.ItemSet{
	Name:  "Incarnate Raiment",
	Items: map[int32]struct{}{29056: {}, 29057: {}, 29058: {}, 29059: {}, 29060: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Your shadowfiend now has 75 more stamina and lasts 3 sec. longer.
			// Implemented in shadowfiend.go.
		},
		4: func(agent core.Agent) {
			// Your Mind Flay and Smite spells deal 5% more damage.
			// Implemented in mind_flay.go.
		},
	},
}

var Avatar2PcAuraID = core.NewAuraID()
var Avatar4PcAuraID = core.NewAuraID()
var SadistAuraID = core.NewAuraID()

var ItemSetAvatar = core.ItemSet{
	Name:  "Avatar Regalia",
	Items: map[int32]struct{}{30160: {}, 30161: {}, 30162: {}, 30159: {}, 30163: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				return core.Aura{
					ID:   Avatar2PcAuraID,
					Name: "Avatar 2pc Bonus",
					OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
						if sim.RandomFloat("avatar 2p") > 0.06 {
							return
						}
						// This is a cheat...
						// easier than adding another aura the subtracts 150 mana from next cast.
						character.AddMana(sim, 150, "Avatar 2p Bonus", false)
					},
				}
			})
		},
		4: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				return core.Aura{
					ID:   Avatar4PcAuraID,
					Name: "Avatar 4pc Bonus",
					OnPeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage float64) {
						if spellCast.ActionID.SpellID != SpellIDShadowWordPain {
							return
						}

						if sim.RandomFloat("avatar 4p") > 0.4 { // 60% chance of not activating.
							return
						}

						character.AddAura(sim, core.Aura{
							ID:      SadistAuraID,
							Name:    "Sadist",
							Expires: sim.CurrentTime + time.Second*15,
							OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
								spellEffect.BonusSpellPower += 100
								character.RemoveAura(sim, SadistAuraID)
							},
						})
					},
				}
			})
		},
	},
}

var Absolution2PcAuraID = core.NewAuraID()
var Absolution4PcAuraID = core.NewAuraID()

var ItemSetAbsolution = core.ItemSet{
	Name:  "Absolution Regalia",
	Items: map[int32]struct{}{31061: {}, 31064: {}, 31067: {}, 31070: {}, 31065: {}, 34434: {}, 34528: {}, 34563: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// this is implemented in swp.go
		},
		4: func(agent core.Agent) {
			// this is implemented in mindblast.go
		},
	},
}
