package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	core.AddItemSet(&ItemSetJusticarBattlegear)
	core.AddItemSet(&ItemSetJusticarArmor)
	core.AddItemSet(&ItemSetCrystalforgeBattlegear)
	core.AddItemSet(&ItemSetCrystalforgeArmor)
	core.AddItemSet(&ItemSetLightbringerBattlegear)
	core.AddItemSet(&ItemSetLightbringerArmor)

	core.AddItemEffect(27484, ApplyLibramOfAvengement)
	core.AddItemEffect(32368, ApplyTomeOfTheLightbringer)
}

var ItemSetJusticarBattlegear = core.ItemSet{
	Name: "Justicar Battlegear",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// sim/debuffs.go handles this (and paladin/judgement.go)
		},
		4: func(agent core.Agent) {
			// TODO: if we ever implemented judgement of command, add bonus from 4p
		},
	},
}

var ItemSetJusticarArmor = core.ItemSet{
	Name: "Justicar Armor",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Increases the damage dealt by your Seal of Righteousness, Seal of
			// Vengeance, and Seal of Blood by 10%.
			// Implemented in seals.go.
		},
		4: func(agent core.Agent) {
			// Increases the damage dealt by Holy Shield by 15.
			// Implemented in holy_shield.go.
		},
	},
}

var ItemSetCrystalforgeBattlegear = core.ItemSet{
	Name: "Crystalforge Battlegear",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// judgement.go
		},
		4: func(agent core.Agent) {
			// TODO: if we implement healing, this heals party.
		},
	},
}

var ItemSetCrystalforgeArmor = core.ItemSet{
	Name: "Crystalforge Armor",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Increases the damage from your Retribution Aura by 15.
			// TODO
		},
		4: func(agent core.Agent) {
			// Each time you use your Holy Shield ability, you gain 100 Block Value
			// against a single attack in the next 6 seconds.
			paladin := agent.(PaladinAgent).GetPaladin()

			procAura := paladin.RegisterAura(core.Aura{
				Label:    "Crystalforge 2pc Proc",
				ActionID: core.ActionID{SpellID: 37191},
				Duration: time.Second * 6,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					paladin.AddStatDynamic(sim, stats.BlockValue, 100)
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					paladin.AddStatDynamic(sim, stats.BlockValue, -100)
				},
				OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
					if spellEffect.Outcome.Matches(core.OutcomeBlock) {
						aura.Deactivate(sim)
					}
				},
			})

			paladin.RegisterAura(core.Aura{
				Label:    "Crystalforge 2pc",
				Duration: core.NeverExpires,
				OnReset: func(aura *core.Aura, sim *core.Simulation) {
					aura.Activate(sim)
				},
				OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
					if spell == paladin.HolyShield {
						procAura.Activate(sim)
					}
				},
			})
		},
	},
}

var ItemSetLightbringerBattlegear = core.ItemSet{
	Name: "Lightbringer Battlegear",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			paladin := agent.(PaladinAgent).GetPaladin()
			paladin.RegisterAura(core.Aura{
				Label:    "Lightbringer Battlegear 2pc",
				Duration: core.NeverExpires,
				OnReset: func(aura *core.Aura, sim *core.Simulation) {
					aura.Activate(sim)
				},
				OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
					if !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
						return
					}
					if sim.RandomFloat("lightbringer 2pc") > 0.2 {
						return
					}
					paladin.AddMana(sim, 50, core.ActionID{SpellID: 38428}, true)
				},
			})
		},
		4: func(agent core.Agent) {
			// TODO: if we implemented hammer of wrath.. this ups dmg
		},
	},
}

var ItemSetLightbringerArmor = core.ItemSet{
	Name: "Lightbringer Armor",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Increases the mana gained from your Spiritual Attunement ability by 10%.
		},
		4: func(agent core.Agent) {
			// Increases the damage dealt by Consecration by 10%.
		},
	},
}

// Librams implemented in seals.go and judgement.go

// TODO: once we have judgement of command.. https://tbc.wowhead.com/item=33503/libram-of-divine-judgement

func ApplyLibramOfAvengement(agent core.Agent) {
	paladin := agent.(PaladinAgent).GetPaladin()
	procAura := paladin.NewTemporaryStatsAura("Libram of Avengement Proc", core.ActionID{SpellID: 34260}, stats.Stats{stats.MeleeCrit: 53, stats.SpellCrit: 53}, time.Second*5)

	paladin.RegisterAura(core.Aura{
		Label:    "Libram of Avengement",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spell == paladin.JudgementOfBlood || spell == paladin.JudgementOfRighteousness {
				procAura.Activate(sim)
			}
		},
	})
}

func ApplyTomeOfTheLightbringer(agent core.Agent) {
	paladin := agent.(PaladinAgent).GetPaladin()
	procAura := paladin.NewTemporaryStatsAura("Tome of the Lightbringer Proc", core.ActionID{SpellID: 41042}, stats.Stats{stats.BlockValue: 186}, time.Second*5)

	paladin.RegisterAura(core.Aura{
		Label:    "Tome of the Lightbringer",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spell.SpellExtras.Matches(SpellFlagJudgement) {
				procAura.Activate(sim)
			}
		},
	})
}
