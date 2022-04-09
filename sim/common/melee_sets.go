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

var ItemSetDesolationBattlegear = core.ItemSet{
	Name: "Desolation Battlegear",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.MeleeHit, 35)
		},
		4: func(agent core.Agent) {
			character := agent.GetCharacter()

			character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
				procAura := character.NewTemporaryStatsAura("Desolation Battlegear Proc", core.ActionID{SpellID: 37617}, stats.Stats{stats.AttackPower: 160, stats.RangedAttackPower: 160}, time.Second*15)
				icd := core.NewICD()
				const icdDur = time.Second * 20
				const procChance = 0.01

				return character.GetOrRegisterAura(&core.Aura{
					Label: "Desolation Battlegear",
					OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
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
						procAura.Activate(sim)
					},
				})
			})
		},
	},
}

var ItemSetDoomplateBattlegear = core.ItemSet{
	Name: "Doomplate Battlegear",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.MeleeHit, 35)
		},
		4: func(agent core.Agent) {
			character := agent.GetCharacter()

			character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
				procAura := character.NewTemporaryStatsAura("Doomplate Battlegear Proc", core.ActionID{SpellID: 37611}, stats.Stats{stats.AttackPower: 160, stats.RangedAttackPower: 160}, time.Second*15)
				const procChance = 0.02
				return character.GetOrRegisterAura(&core.Aura{
					Label: "Doomplate Battlegear",
					OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
						if !spellEffect.Landed() {
							return
						}
						if !spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
							return
						}
						if sim.RandomFloat("Doomplate Battlegear") > procChance {
							return
						}
						procAura.Activate(sim)
					},
				})
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
				},
				ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
					IsPhantom:        true,
					DamageMultiplier: 1,
					ThreatMultiplier: 1,

					BaseDamage:     core.BaseDamageConfigRoll(100, 150),
					OutcomeApplier: core.OutcomeFuncMagicHitAndCrit(character.DefaultSpellCritMultiplier()),
				}),
			})

			character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
				ppmm := character.AutoAttacks.NewPPMManager(2)

				return character.GetOrRegisterAura(&core.Aura{
					Label: "Fists of Fury",
					OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
						if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
							return
						}
						if !ppmm.Proc(sim, spellEffect.IsMH(), false, "The Fists of Fury") {
							return
						}

						procSpell.Cast(sim, spellEffect.Target)
					},
				})
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

			character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
				procAura := character.NewTemporaryStatsAura("Twin Blade of Azzinoth Proc", core.ActionID{SpellID: 41435}, stats.Stats{stats.MeleeHaste: 450}, time.Second*10)
				ppmm := character.AutoAttacks.NewPPMManager(1.0)

				icd := core.NewICD()
				const icdDur = time.Second * 45

				return character.GetOrRegisterAura(&core.Aura{
					Label: "Twin Blades of Azzinoth",
					OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
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
						procAura.Activate(sim)
					},
				})
			})
		},
	},
}

var ItemSetWastewalkerArmor = core.ItemSet{
	Name: "Wastewalker Armor",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			agent.GetCharacter().AddStat(stats.MeleeHit, 35)
		},
		4: func(agent core.Agent) {
			character := agent.GetCharacter()

			character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
				procAura := character.NewTemporaryStatsAura("Wastewalker Armor Proc", core.ActionID{SpellID: 37618}, stats.Stats{stats.AttackPower: 160, stats.RangedAttackPower: 160}, time.Second*15)
				icd := core.NewICD()
				const icdDur = time.Second * 20
				const procChance = 0.02

				return character.GetOrRegisterAura(&core.Aura{
					Label: "Wastewalker Armor",
					OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
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
						procAura.Activate(sim)
					},
				})
			})
		},
	},
}
