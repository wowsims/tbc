package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Keep these (and their functions) in alphabetical order.
func init() {
	core.AddItemSet(ItemSetDesolationBattlegear)
	core.AddItemSet(ItemSetEbonNetherscale)
	core.AddItemSet(ItemSetFelstalker)
	core.AddItemSet(ItemSetPrimalstrike)
	core.AddItemSet(ItemSetWastewalkerArmor)
}

var DesolationBattlegearAuraID = core.NewAuraID()
var DesolationBattlegearProcAuraID = core.NewAuraID()
var ItemSetDesolationBattlegear = core.ItemSet{
	Name:  "Desolation Battlegear",
	Items: map[int32]struct{}{28192: {}, 27713: {}, 28401: {}, 27936: {}, 27528: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.MeleeHit, 35)
		},
		4: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				const apBonus = 160.0
				const duration = time.Second * 15
				const procChance = 0.01

				icd := core.NewICD()
				const icdDur = time.Second * 20

				return core.Aura{
					ID: DesolationBattlegearAuraID,
					OnMeleeAttack: func(sim *core.Simulation, target *core.Target, result core.MeleeHitType, ability *core.ActiveMeleeAbility, isOH bool) {
						if result == core.MeleeHitTypeMiss || result == core.MeleeHitTypeDodge || result == core.MeleeHitTypeParry {
							return
						}
						if icd.IsOnCD(sim) {
							return
						}
						if sim.RandomFloat("Desolation Battlegear") > procChance {
							return
						}
						icd = core.InternalCD(sim.CurrentTime + icdDur)
						character.AddAuraWithTemporaryStats(sim, DesolationBattlegearProcAuraID, core.ActionID{ItemID: 28192}, stats.AttackPower, apBonus, duration)
					},
				}
			})
		},
	},
}

var ItemSetEbonNetherscale = core.ItemSet{
	Name:  "Ebon Netherscale",
	Items: map[int32]struct{}{29515: {}, 29516: {}, 29517: {}},
	Bonuses: map[int32]core.ApplyEffect{
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.MeleeHit, 20)
		},
	},
}

var ItemSetFelstalker = core.ItemSet{
	Name:  "Felstalker",
	Items: map[int32]struct{}{25696: {}, 25695: {}, 25697: {}},
	Bonuses: map[int32]core.ApplyEffect{
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.MeleeHit, 20)
		},
	},
}

var ItemSetPrimalstrike = core.ItemSet{
	Name:  "Primalstrike",
	Items: map[int32]struct{}{29525: {}, 29526: {}, 29527: {}},
	Bonuses: map[int32]core.ApplyEffect{
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.AttackPower, 40)
		},
	},
}

var WastewalkerArmorAuraID = core.NewAuraID()
var WastewalkerArmorProcAuraID = core.NewAuraID()
var ItemSetWastewalkerArmor = core.ItemSet{
	Name:  "Wastewalker Armor",
	Items: map[int32]struct{}{28224: {}, 27797: {}, 28264: {}, 27837: {}, 27531: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.MeleeHit, 35)
		},
		4: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				const apBonus = 160.0
				const duration = time.Second * 15
				const procChance = 0.02

				icd := core.NewICD()
				const icdDur = time.Second * 20

				return core.Aura{
					ID: WastewalkerArmorAuraID,
					OnMeleeAttack: func(sim *core.Simulation, target *core.Target, result core.MeleeHitType, ability *core.ActiveMeleeAbility, isOH bool) {
						if result == core.MeleeHitTypeMiss || result == core.MeleeHitTypeDodge || result == core.MeleeHitTypeParry {
							return
						}
						if icd.IsOnCD(sim) {
							return
						}
						if sim.RandomFloat("Wastewalker Armor") > procChance {
							return
						}
						icd = core.InternalCD(sim.CurrentTime + icdDur)
						character.AddAuraWithTemporaryStats(sim, WastewalkerArmorProcAuraID, core.ActionID{ItemID: 28192}, stats.AttackPower, apBonus, duration)
					},
				}
			})
		},
	},
}
