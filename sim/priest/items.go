package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	core.AddItemEffect(32490, ApplyAshtongueTalismanOfAcumen)

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
					ID: Avatar2PcAuraID,
					OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
						if sim.RandomFloat("avatar 2p") > 0.06 {
							return
						}
						// This is a cheat...
						// easier than adding another aura the subtracts 150 mana from next cast.
						character.AddMana(sim, 150, core.ActionID{SpellID: 37600}, false)
					},
				}
			})
		},
		4: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				return core.Aura{
					ID: Avatar4PcAuraID,
					OnPeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage float64) {
						if spellCast.ActionID.SpellID != SpellIDShadowWordPain {
							return
						}

						if sim.RandomFloat("avatar 4p") > 0.4 { // 60% chance of not activating.
							return
						}

						character.AddAura(sim, core.Aura{
							ID:       SadistAuraID,
							ActionID: core.ActionID{SpellID: 37604},
							Expires:  sim.CurrentTime + time.Second*15,
							OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellHitEffect) {
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

var AshtongueTalismanOfAcumenItemAuraID = core.NewAuraID()
var AshtongueTalismanOfAcumenAuraID = core.NewAuraID()

func ApplyAshtongueTalismanOfAcumen(agent core.Agent) {
	// Not in the game yet so cant test; this logic assumes that:
	// - procrate is 10%
	// - no ICD on proc
	const spellBonus = 220
	const dur = time.Second * 10
	const procrate = 0.1

	char := agent.GetCharacter()
	char.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID: AshtongueTalismanOfAcumenItemAuraID,
			OnPeriodicDamage: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect, tickDamage float64) {
				if spellCast.ActionID.SpellID != SpellIDShadowWordPain {
					return
				}

				if sim.RandomFloat("Ashtongue Talisman of Acumen") > procrate {
					return
				}

				char.AddAuraWithTemporaryStats(sim, AshtongueTalismanOfAcumenAuraID, core.ActionID{ItemID: 32490}, stats.SpellPower, spellBonus, dur)
			},
		}
	})
}
