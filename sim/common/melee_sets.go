package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Keep these (and their functions) in alphabetical order.
func init() {
	core.AddItemSet(&ItemSetBurningRage)
	core.AddItemSet(&ItemSetDesolationBattlegear)
	core.AddItemSet(&ItemSetDoomplateBattlegear)
	core.AddItemSet(&ItemSetEbonNetherscale)
	core.AddItemSet(&ItemSetFaithInFelsteel)
	core.AddItemSet(&ItemSetFelstalker)
	core.AddItemSet(&ItemSetFistsOfFury)
	core.AddItemSet(&ItemSetFlameGuard)
	core.AddItemSet(&ItemSetPrimalstrike)
	core.AddItemSet(&ItemSetStrengthOfTheClefthoof)
	core.AddItemSet(&ItemSetTwinBladesOfAzzinoth)
	core.AddItemSet(&ItemSetWastewalkerArmor)
}

var ItemSetBurningRage = core.ItemSet{
	Name: "Burning Rage",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.MeleeHit, 20)
		},
	},
}

var DesolationBattlegearAuraID = core.NewAuraID()
var DesolationBattlegearProcAuraID = core.NewAuraID()
var ItemSetDesolationBattlegear = core.ItemSet{
	Name: "Desolation Battlegear",
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
				applyStatAura := character.NewTemporaryStatsAuraApplier(DesolationBattlegearProcAuraID, core.ActionID{SpellID: 37617}, stats.Stats{stats.AttackPower: apBonus, stats.RangedAttackPower: apBonus}, duration)

				icd := core.NewICD()
				const icdDur = time.Second * 20

				return core.Aura{
					ID: DesolationBattlegearAuraID,
					OnSpellHit: func(sim *core.Simulation, spell *core.SimpleSpellTemplate, spellEffect *core.SpellEffect) {
						if !spellEffect.Landed() {
							return
						}
						if !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
							return
						}
						if icd.IsOnCD(sim) {
							return
						}
						if sim.RandomFloat("Desolation Battlegear") > procChance {
							return
						}
						icd = core.InternalCD(sim.CurrentTime + icdDur)
						applyStatAura(sim)
					},
				}
			})
		},
	},
}

var DoomplateBattlegearAuraID = core.NewAuraID()
var DoomplateBattlegearProcAuraID = core.NewAuraID()
var ItemSetDoomplateBattlegear = core.ItemSet{
	Name: "Doomplate Battlegear",
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
				applyStatAura := character.NewTemporaryStatsAuraApplier(DoomplateBattlegearProcAuraID, core.ActionID{SpellID: 37611}, stats.Stats{stats.AttackPower: apBonus, stats.RangedAttackPower: apBonus}, duration)

				return core.Aura{
					ID: DoomplateBattlegearAuraID,
					OnSpellHit: func(sim *core.Simulation, spell *core.SimpleSpellTemplate, spellEffect *core.SpellEffect) {
						if !spellEffect.Landed() {
							return
						}
						if !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
							return
						}
						if sim.RandomFloat("Doomplate Battlegear") > procChance {
							return
						}
						applyStatAura(sim)
					},
				}
			})
		},
	},
}

var ItemSetEbonNetherscale = core.ItemSet{
	Name: "Netherscale Armor",
	Bonuses: map[int32]core.ApplyEffect{
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.MeleeHit, 20)
		},
	},
}

var ItemSetFaithInFelsteel = core.ItemSet{
	Name: "Faith in Felsteel",
	Bonuses: map[int32]core.ApplyEffect{
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.Strength, 25)
		},
	},
}

var ItemSetFelstalker = core.ItemSet{
	Name: "Felstalker Armor",
	Bonuses: map[int32]core.ApplyEffect{
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.MeleeHit, 20)
		},
	},
}

var FistsOfFuryAuraID = core.NewAuraID()
var ItemSetFistsOfFury = core.ItemSet{
	Name: "The Fists of Fury",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			character := agent.GetCharacter()

			procSpell := character.RegisterSpell(core.SpellConfig{
				Template: core.SimpleSpell{
					SpellCast: core.SpellCast{
						Cast: core.Cast{
							ActionID:    core.ActionID{SpellID: 41989},
							Character:   character,
							SpellSchool: core.SpellSchoolFire,
						},
					},
					Effect: core.SpellEffect{
						OutcomeRollCategory: core.OutcomeRollCategoryMagic,
						CritRollCategory:    core.CritRollCategoryMagical,
						CritMultiplier:      character.DefaultSpellCritMultiplier(),
						IsPhantom:           true,
						DamageMultiplier:    1,
						ThreatMultiplier:    1,
						BaseDamage:          core.BaseDamageConfigRoll(100, 150),
					},
				},
				ModifyCast: core.ModifyCastAssignTarget,
			})

			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				ppmm := character.AutoAttacks.NewPPMManager(2)

				return core.Aura{
					ID: FistsOfFuryAuraID,
					OnSpellHit: func(sim *core.Simulation, spell *core.SimpleSpellTemplate, spellEffect *core.SpellEffect) {
						if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
							return
						}
						if !ppmm.Proc(sim, spellEffect.IsMH(), false, "The Fists of Fury") {
							return
						}

						procSpell.Cast(sim, spellEffect.Target)
					},
				}
			})
		},
	},
}

var ItemSetFlameGuard = core.ItemSet{
	Name: "Flame Guard",
	Bonuses: map[int32]core.ApplyEffect{
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.Parry, 20)
		},
	},
}

var ItemSetPrimalstrike = core.ItemSet{
	Name: "Primal Intent",
	Bonuses: map[int32]core.ApplyEffect{
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.AttackPower, 40)
			agent.GetCharacter().AddStat(stats.RangedAttackPower, 40)
		},
	},
}

var ItemSetStrengthOfTheClefthoof = core.ItemSet{
	Name: "Strength of the Clefthoof",
	Bonuses: map[int32]core.ApplyEffect{
		3: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.Strength, 20)
		},
	},
}

var TwinBladesOfAzzinothAuraID = core.NewAuraID()
var TwinBladesOfAzzinothProcAuraID = core.NewAuraID()
var ItemSetTwinBladesOfAzzinoth = core.ItemSet{
	Name: "The Twin Blades of Azzinoth",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			character := agent.GetCharacter()

			character.RegisterResetEffect(func(sim *core.Simulation) {
				if sim.GetPrimaryTarget().MobType == proto.MobType_MobTypeDemon {
					character.PseudoStats.MobTypeAttackPower += 200
				}
			})

			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				const hasteBonus = 450.0
				const duration = time.Second * 10
				applyStatAura := character.NewTemporaryStatsAuraApplier(TwinBladesOfAzzinothProcAuraID, core.ActionID{SpellID: 41435}, stats.Stats{stats.MeleeHaste: hasteBonus}, duration)
				ppmm := character.AutoAttacks.NewPPMManager(1.0)

				icd := core.NewICD()
				const icdDur = time.Second * 45

				return core.Aura{
					ID: TwinBladesOfAzzinothAuraID,
					OnSpellHit: func(sim *core.Simulation, spell *core.SimpleSpellTemplate, spellEffect *core.SpellEffect) {
						if !spellEffect.Landed() {
							return
						}

						// https://tbc.wowhead.com/spell=41434/the-twin-blades-of-azzinoth, proc mask = 20.
						if !spellEffect.ProcMask.Matches(core.ProcMaskMelee) || spellEffect.IsPhantom {
							return
						}

						if icd.IsOnCD(sim) {
							return
						}

						if !ppmm.Proc(sim, spellEffect.IsMH(), false, "Twin Blades of Azzinoth") {
							return
						}
						icd = core.InternalCD(sim.CurrentTime + icdDur)
						applyStatAura(sim)
					},
				}
			})
		},
	},
}

var WastewalkerArmorAuraID = core.NewAuraID()
var WastewalkerArmorProcAuraID = core.NewAuraID()
var ItemSetWastewalkerArmor = core.ItemSet{
	Name: "Wastewalker Armor",
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
				applyStatAura := character.NewTemporaryStatsAuraApplier(WastewalkerArmorProcAuraID, core.ActionID{SpellID: 37618}, stats.Stats{stats.AttackPower: apBonus, stats.RangedAttackPower: apBonus}, duration)

				icd := core.NewICD()
				const icdDur = time.Second * 20

				return core.Aura{
					ID: WastewalkerArmorAuraID,
					OnSpellHit: func(sim *core.Simulation, spell *core.SimpleSpellTemplate, spellEffect *core.SpellEffect) {
						if !spellEffect.Landed() {
							return
						}
						if !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
							return
						}
						if icd.IsOnCD(sim) {
							return
						}
						if sim.RandomFloat("Wastewalker Armor") > procChance {
							return
						}
						icd = core.InternalCD(sim.CurrentTime + icdDur)
						applyStatAura(sim)
					},
				}
			})
		},
	},
}
