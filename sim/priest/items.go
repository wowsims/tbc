package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	core.AddItemEffect(32490, ApplyAshtongueTalismanOfAcumen)

	core.AddItemSet(&ItemSetIncarnate)
	core.AddItemSet(&ItemSetAvatar)
	core.AddItemSet(&ItemSetAbsolution)
}

var ItemSetIncarnate = core.ItemSet{
	Name: "Incarnate Raiment",
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

var ItemSetAvatar = core.ItemSet{
	Name: "Avatar Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
				return character.GetOrRegisterAura(&core.Aura{
					Label: "Avatar Regalia 2pc",
					OnCastComplete: func(aura *core.Aura, sim *core.Simulation, cast *core.Cast) {
						if sim.RandomFloat("avatar 2p") > 0.06 {
							return
						}
						// This is a cheat...
						// easier than adding another aura the subtracts 150 mana from next cast.
						character.AddMana(sim, 150, core.ActionID{SpellID: 37600}, false)
					},
				})
			})
		},
		4: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
				procAura := character.NewTemporaryStatsAura("Avatar Regalia 4pc Proc", core.ActionID{SpellID: 37604}, stats.Stats{stats.SpellPower: 100}, time.Second*15)
				procAura.OnSpellHit = func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
					aura.Deactivate(sim)
				}

				return character.GetOrRegisterAura(&core.Aura{
					Label: "Avatar Regalia 4pc",
					OnPeriodicDamage: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect, tickDamage float64) {
						if spell.ActionID.SpellID != SpellIDShadowWordPain {
							return
						}

						if sim.RandomFloat("avatar 4p") > 0.4 { // 60% chance of not activating.
							return
						}

						procAura.Activate(sim)
					},
				})
			})
		},
	},
}

var ItemSetAbsolution = core.ItemSet{
	Name: "Absolution Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// this is implemented in swp.go
		},
		4: func(agent core.Agent) {
			// this is implemented in mindblast.go
		},
	},
}

func ApplyAshtongueTalismanOfAcumen(agent core.Agent) {
	// Not in the game yet so cant test; this logic assumes that:
	// - procrate is 10%
	// - no ICD on proc
	const procrate = 0.1

	char := agent.GetCharacter()
	char.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		procAura := char.NewTemporaryStatsAura("Ashtongue Talisman Proc", core.ActionID{ItemID: 32490}, stats.Stats{stats.SpellPower: 220}, time.Second*10)
		return char.GetOrRegisterAura(&core.Aura{
			Label: "Ashtongue Talisman",
			OnPeriodicDamage: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect, tickDamage float64) {
				if spell.ActionID.SpellID != SpellIDShadowWordPain {
					return
				}

				if sim.RandomFloat("Ashtongue Talisman of Acumen") > procrate {
					return
				}

				procAura.Activate(sim)
			},
		})
	})
}
