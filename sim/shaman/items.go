package shaman

import (
	"log"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	core.AddActiveItem(33506, core.ActiveItem{BuffUp: ActivateSkycall, ActivateCD: core.NeverExpires})
	core.AddActiveItem(19344, core.ActiveItem{Activate: ActivateNAC, ActivateCD: time.Second * 300, CoolID: core.MagicIDNACTrink, SharedID: core.MagicIDAtkTrinket})
	core.AddActiveItem(30663, core.ActiveItem{BuffUp: ActivateFathomBrooch, ActivateCD: core.NeverExpires})

	core.AddItemSet(ItemSetTidefury)
	core.AddItemSet(ItemSetCycloneRegalia)
	core.AddItemSet(ItemSetCataclysmRegalia)
	core.AddItemSet(ItemSetSkyshatterRegalia)
}

var ItemSetTidefury = core.ItemSet{
	Name:  "Tidefury",
	Items: map[int32]struct{}{28231: {}, 27510: {}, 28349: {}, 27909: {}, 27802: {}},
	Bonuses: map[int]core.ItemBuffUpFunc{
		2: func(sim *core.Simulation, agent core.Agent) {
			agent.GetCharacter().AddAura(sim, core.Aura{ID: core.MagicIDTidefury, Expires: core.NeverExpires})
		},
		4: func(sim *core.Simulation, agent core.Agent) {
			shaman, ok := agent.(*Shaman)
			if !ok {
				log.Fatalf("Non-shaman attempted to activate shaman cyclone set bonus.")
			}

			if shaman.SelfBuffs.WaterShield {
				shaman.Stats[stats.MP5] += 3
			}
		},
	},
}

var ItemSetCycloneRegalia = core.ItemSet{
	Name:  "Cyclone Regalia",
	Items: map[int32]struct{}{29033: {}, 29035: {}, 29034: {}, 29036: {}, 29037: {}},
	Bonuses: map[int]core.ItemBuffUpFunc{
		2: func(sim *core.Simulation, agent core.Agent) {
			// Handled in shaman.go
		},
		4: func(sim *core.Simulation, agent core.Agent) {
			character := agent.GetCharacter()
			character.AddAura(sim, core.Aura{
				ID:      core.MagicIDCyclone4pc,
				Expires: core.NeverExpires,
				OnSpellHit: func(sim *core.Simulation, cast core.DirectCastAction, result *core.DirectCastDamageResult) {
					if result.Crit && sim.Rando.Float64("cycl4p") < 0.11 {
						character.AddAura(sim, core.Aura{
							ID: core.MagicIDCycloneMana,
							OnCast: func(sim *core.Simulation, cast core.DirectCastAction, input *core.DirectCastInput) {
								// TODO: how to make sure this goes in before clearcasting?
								input.ManaCost -= 270
							},
							OnCastComplete: func(sim *core.Simulation, cast core.DirectCastAction) {
								character.RemoveAura(sim, core.MagicIDCycloneMana)
							},
						})
					}
				},
			})
		},
	},
}

var ItemSetCataclysmRegalia = core.ItemSet{
	Name:    "Cataclysm Regalia",
	Items:   map[int32]struct{}{30169: {}, 30170: {}, 30171: {}, 30172: {}, 30173: {}},
	Bonuses: map[int]core.ItemBuffUpFunc{
		4: func(sim *core.Simulation, agent core.Agent) {
			character := agent.GetCharacter()
			character.AddAura(sim, core.Aura{
				ID:      core.MagicIDCataclysm4pc,
				Expires: core.NeverExpires,
				OnSpellHit: func(sim *core.Simulation, cast core.DirectCastAction, result *core.DirectCastDamageResult) {
					if result.Crit && sim.Rando.Float64("cata4p") < 0.25 {
						character.Stats[stats.Mana] += 120
					}
				},
			})
		},
	},
}

var ItemSetSkyshatterRegalia = core.ItemSet{
	Name:  "Skyshatter Regalia",
	Items: map[int32]struct{}{34437: {}, 31017: {}, 34542: {}, 31008: {}, 31014: {}, 31020: {}, 31023: {}, 34566: {}},
	Bonuses: map[int]core.ItemBuffUpFunc{
		2: func(sim *core.Simulation, agent core.Agent) {
			agent.GetCharacter().Stats[stats.MP5] += 15
			agent.GetCharacter().Stats[stats.SpellCrit] += 35
			agent.GetCharacter().Stats[stats.SpellPower] += 45
		},
		4: func(sim *core.Simulation, agent core.Agent) {
			character := agent.GetCharacter()
			character.AddAura(sim, core.Aura{
				ID:      core.MagicIDSkyshatter4pc,
				Expires: core.NeverExpires,
				OnSpellHit: func(sim *core.Simulation, cast core.DirectCastAction, result *core.DirectCastDamageResult) {
					if cast.GetActionID().SpellID == SpellIDLB12 {
						result.Damage *= 1.05
					}
				},
			})
		},
	},
}

func ActivateSkycall(sim *core.Simulation, agent core.Agent) {
	character := agent.GetCharacter()
	const hasteBonus = 101
	const dur = time.Second * 10

	character.AddAura(sim, core.Aura{
		ID:      core.MagicIDSkycall,
		Expires: core.NeverExpires,
		OnCastComplete: func(sim *core.Simulation, cast core.DirectCastAction) {
			if cast.GetActionID().SpellID == SpellIDLB12 && sim.Rando.Float64("skycall") < 0.15 {
				core.AddAuraWithTemporaryStats(sim, character, core.MagicIDEnergized, stats.SpellHaste, hasteBonus, dur)
			}
		},
	})
}

func ActivateNAC(sim *core.Simulation, character *core.Character) core.Aura {
	const sp = 250
	character.Stats[stats.SpellPower] += sp
	return core.Aura{
		ID:      core.MagicIDNAC,
		Expires: sim.CurrentTime + time.Second*20,
		OnCast: func(sim *core.Simulation, cast core.DirectCastAction, input *core.DirectCastInput) {
			input.ManaCost *= 1.2
		},
		OnExpire: func(sim *core.Simulation) {
			character.Stats[stats.SpellPower] -= sp
		},
	}
}

// ActivateFathomBrooch adds an aura that has a chance on cast of nature spell
//  to restore 335 mana. 40s ICD
func ActivateFathomBrooch(sim *core.Simulation, agent core.Agent) {
	character := agent.GetCharacter()
	icd := core.NewICD()
	const icdDur = time.Second * 40

	character.AddAura(sim, core.Aura{
		ID:      core.MagicIDRegainMana,
		Expires: core.NeverExpires,
		OnCastComplete: func(sim *core.Simulation, cast core.DirectCastAction) {
			if icd.IsOnCD(sim) {
				return
			}
			if cast.GetSpellSchool() != stats.NatureSpellPower {
				return
			}
			if sim.Rando.Float64("unmarked") < 0.15 {
				icd = core.InternalCD(sim.CurrentTime + icdDur)
				character.Stats[stats.Mana] += 335
			}
		},
	})
}
