package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Keep these (and their functions) in alphabetical order.
func init() {
	core.AddItemSet(ItemSetDesolationBattlegear)
	core.AddItemSet(ItemSetEbonNetherscale)
	core.AddItemSet(ItemSetFelstalker)
	core.AddItemSet(ItemSetFistsOfFury)
	core.AddItemSet(ItemSetPrimalstrike)
	core.AddItemSet(ItemSetTwinBladesOfAzzinoth)
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
					OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
						if !hitEffect.Landed() {
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

var FistsOfFuryAuraID = core.NewAuraID()
var ItemSetFistsOfFury = core.ItemSet{
	Name:  "The Fists of Fury",
	Items: map[int32]struct{}{32945: {}, 32946: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			spellObj := core.SimpleSpell{}

			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				// TODO: Get real numbers for this.
				ppmm := character.AutoAttacks.NewPPMManager(1.0)

				castTemplate := core.NewSimpleSpellTemplate(core.SimpleSpell{
					SpellCast: core.SpellCast{
						Cast: core.Cast{
							ActionID:       core.ActionID{SpellID: 41989},
							Character:      character,
							IsPhantom:      true,
							SpellSchool:    stats.FireSpellPower,
							CritMultiplier: character.DefaultSpellCritMultiplier(),
						},
					},
					Effect: core.SpellHitEffect{
						SpellEffect: core.SpellEffect{
							DamageMultiplier:       1,
							StaticDamageMultiplier: 1,
							ThreatMultiplier:       1,
						},
						DirectInput: core.DirectDamageInput{
							MinBaseDamage: 100,
							MaxBaseDamage: 150,
						},
					},
				})

				return core.Aura{
					ID: FistsOfFuryAuraID,
					OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
						if !hitEffect.Landed() || !hitEffect.ProcMask.Matches(core.ProcMaskMelee) {
							return
						}
						if !ppmm.Proc(sim, hitEffect.IsMH(), false, "The Fists of Fury") {
							return
						}

						castAction := &spellObj
						castTemplate.Apply(castAction)
						castAction.Effect.Target = hitEffect.Target
						castAction.Init(sim)
						castAction.Cast(sim)
					},
				}
			})
		},
	},
}

var ItemSetPrimalstrike = core.ItemSet{
	Name:  "Primalstrike",
	Items: map[int32]struct{}{29525: {}, 29526: {}, 29527: {}},
	Bonuses: map[int32]core.ApplyEffect{
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.AttackPower, 40)
			agent.GetCharacter().AddStat(stats.RangedAttackPower, 40)
		},
	},
}

var TwinBladesOfAzzinothAuraID = core.NewAuraID()
var TwinBladesOfAzzinothProcAuraID = core.NewAuraID()
var ItemSetTwinBladesOfAzzinoth = core.ItemSet{
	Name:  "The Twin Blades of Azzinoth",
	Items: map[int32]struct{}{32837: {}, 32838: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				const hasteBonus = 450.0
				const duration = time.Second * 10
				statApplier := character.NewTempStatAuraApplier(sim, TwinBladesOfAzzinothProcAuraID, core.ActionID{SpellID: 41435}, stats.MeleeHaste, hasteBonus, duration)
				ppmm := character.AutoAttacks.NewPPMManager(1.0)

				icd := core.NewICD()
				const icdDur = time.Second * 45

				return core.Aura{
					ID: TwinBladesOfAzzinothAuraID,
					OnBeforeMeleeHit: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
						if hitEffect.Target.MobType == proto.MobType_MobTypeDemon {
							hitEffect.BonusAttackPower += 200
						}
					},
					OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
						if !hitEffect.Landed() {
							return
						}

						// https://tbc.wowhead.com/spell=41434/the-twin-blades-of-azzinoth, proc mask = 20.
						if !hitEffect.ProcMask.Matches(core.ProcMaskMelee) || ability.IsPhantom {
							return
						}

						if icd.IsOnCD(sim) {
							return
						}

						if !ppmm.Proc(sim, hitEffect.IsMH(), false, "Twin Blades of Azzinoth") {
							return
						}
						icd = core.InternalCD(sim.CurrentTime + icdDur)
						statApplier(sim)
					},
				}
			})
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
					OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
						if !hitEffect.Landed() {
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
