package shaman

import (
	"log"
	"math"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	core.AddActiveItem(33506, core.ActiveItem{Activate: ActivateSkycall, ActivateCD: core.NeverExpires})
	core.AddActiveItem(29389, core.ActiveItem{Activate: ActivateTotemOfPulsingEarth, ActivateCD: core.NeverExpires})
	core.AddActiveItem(19344, core.ActiveItem{Activate: ActivateNAC, ActivateCD: time.Second * 300, CoolID: core.MagicIDNACTrink, SharedID: core.MagicIDAtkTrinket})

	// core.AddItemSet()
}

var shaman_sets = []core.ItemSet{
	// Netherstrike is technically not shaman exclusive, but no other class that can wear it would want it.
	{
		Name:  "Netherstrike",
		Items: map[int32]struct{}{29519: {}, 29521: {}, 29520: {}},
		Bonuses: map[int]core.ItemActivation{3: func(sim *core.Simulation, agent core.Agent) core.Aura {
			agent.GetCharacter().Stats[stats.SpellPower] += 23
			return core.Aura{ID: core.MagicIDNetherstrike}
		}},
	},
	{
		Name:  "Tidefury",
		Items: map[int32]struct{}{28231: {}, 27510: {}, 28349: {}, 27909: {}, 27802: {}},
		Bonuses: map[int]core.ItemActivation{
			2: func(sim *core.Simulation, agent core.Agent) core.Aura {
				return core.Aura{ID: core.MagicIDTidefury, Expires: core.NeverExpires}
			},
			4: func(sim *core.Simulation, agent core.Agent) core.Aura {
				shaman, ok := agent.(*Shaman)
				if !ok {
					log.Fatalf("Non-shaman attempted to activate shaman cyclone set bonus.")
				}
				if !shaman.SelfBuffs.WaterShield {
					return core.Aura{}
				}

				agent.GetCharacter().Stats[stats.MP5] += 3
				return core.Aura{} // no aura to add
			},
		},
	},
	{
		Name:  "Cyclone Regalia",
		Items: map[int32]struct{}{29033: {}, 29035: {}, 29034: {}, 29036: {}, 29037: {}},
		Bonuses: map[int]core.ItemActivation{4: ActivateCycloneManaReduce, 2: func(sim *core.Simulation, agent core.Agent) core.Aura {
			shaman, ok := agent.(*Shaman)
			if !ok {
				log.Fatalf("Non-shaman attempted to activate shaman cyclone set bonus.")
			}
			if !shaman.SelfBuffs.WrathOfAir {
				return core.Aura{}
			}
			if shaman.HasAura(core.MagicIDCyclone2pc) {
				return core.Aura{} // only can activate 2pc bonus once
			}
			a := core.Aura{
				ID:      core.MagicIDCyclone2pc,
				Expires: core.NeverExpires,
				OnExpire: func(sim *core.Simulation, c *core.Cast) {
					agent.GetCharacter().Party.AddStats(stats.Stats{stats.SpellPower: -20})
				},
			}
			// Give the party stats, and the aura.
			agent.GetCharacter().Party.AddStats(stats.Stats{stats.SpellPower: 20})
			agent.GetCharacter().Party.AddAura(sim, a)

			// Return no aura to be directly added because we added the aura already.
			return core.Aura{}
		}},
	},
	{
		Name:    "Cataclysm Regalia",
		Items:   map[int32]struct{}{30169: {}, 30170: {}, 30171: {}, 30172: {}, 30173: {}},
		Bonuses: map[int]core.ItemActivation{4: ActivateCataclysmLBDiscount},
	},
	{
		Name:  "Skyshatter Regalia",
		Items: map[int32]struct{}{34437: {}, 31017: {}, 34542: {}, 31008: {}, 31014: {}, 31020: {}, 31023: {}, 34566: {}},
		Bonuses: map[int]core.ItemActivation{2: func(sim *core.Simulation, agent core.Agent) core.Aura {
			agent.GetCharacter().Stats[stats.MP5] += 15
			agent.GetCharacter().Stats[stats.SpellCrit] += 35
			agent.GetCharacter().Stats[stats.SpellPower] += 45
			return core.Aura{ID: core.MagicIDSkyshatter2pc}
		}, 4: ActivateSkyshatterImpLB},
	},
}

func ActivateSkycall(sim *core.Simulation, agent core.Agent) core.Aura {
	const hasteBonus = 101
	const dur = time.Second * 10
	return core.Aura{
		ID:      core.MagicIDSkycall,
		Expires: core.NeverExpires,
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
			if cast.Spell.ActionID.SpellID == SpellIDLB12 && sim.Rando.Float64("skycall") < 0.15 {
				core.AddAuraWithTemporaryStats(sim, agent, core.MagicIDEnergized, stats.SpellHaste, hasteBonus, dur)
			}
		},
	}
}

func ActivateNAC(sim *core.Simulation, agent core.Agent) core.Aura {
	const sp = 250
	agent.GetCharacter().Stats[stats.SpellPower] += sp
	return core.Aura{
		ID:      core.MagicIDNAC,
		Expires: sim.CurrentTime + time.Second*20,
		OnCast: func(sim *core.Simulation, cast *core.Cast) {
			cast.ManaCost *= 1.2
		},
		OnExpire: func(sim *core.Simulation, cast *core.Cast) {
			agent.GetCharacter().Stats[stats.SpellPower] -= sp
		},
	}
}

func ActivateTotemOfPulsingEarth(sim *core.Simulation, agent core.Agent) core.Aura {
	return core.Aura{
		ID:      core.MagicIDTotemOfPulsingEarth,
		Expires: core.NeverExpires,
		OnCast: func(sim *core.Simulation, cast *core.Cast) {
			if cast.Spell.ActionID.SpellID == SpellIDLB12 {
				// TODO: how to make sure this goes in before clearcasting?
				cast.ManaCost = math.Max(cast.ManaCost-27, 0)
			}
		},
	}
}

func ActivateCycloneManaReduce(sim *core.Simulation, agent core.Agent) core.Aura {
	return core.Aura{
		ID:      core.MagicIDCyclone4pc,
		Expires: core.NeverExpires,
		OnSpellHit: func(sim *core.Simulation, cast *core.Cast) {
			if cast.DidCrit && sim.Rando.Float64("cycl4p") < 0.11 {
				agent.GetCharacter().AddAura(sim, core.Aura{
					ID: core.MagicIDCycloneMana,
					OnCast: func(sim *core.Simulation, cast *core.Cast) {
						// TODO: how to make sure this goes in before clearcasting?
						cast.ManaCost -= 270
					},
					OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
						agent.GetCharacter().RemoveAura(sim, core.MagicIDCycloneMana)
					},
				})
			}
		},
	}
}

func ActivateCataclysmLBDiscount(sim *core.Simulation, agent core.Agent) core.Aura {
	return core.Aura{
		ID:      core.MagicIDCataclysm4pc,
		Expires: core.NeverExpires,
		OnSpellHit: func(sim *core.Simulation, cast *core.Cast) {
			if cast.DidCrit && sim.Rando.Float64("cata4p") < 0.25 {
				agent.GetCharacter().Stats[stats.Mana] += 120
			}
		},
	}
}

func ActivateSkyshatterImpLB(sim *core.Simulation, agent core.Agent) core.Aura {
	return core.Aura{
		ID:      core.MagicIDSkyshatter4pc,
		Expires: core.NeverExpires,
		OnSpellHit: func(sim *core.Simulation, cast *core.Cast) {
			if cast.Spell.ActionID.SpellID == SpellIDLB12 {
				cast.DidDmg *= 1.05
			}
		},
	}
}
