package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	core.AddItemSet(&ItemSetJusticarBattlegear)
	core.AddItemSet(&ItemSetCrystalforgeBattlegear)

	core.AddItemEffect(27484, ApplyAvengementLibramEffect)
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

var ItemSetLightbringerBattlegear = core.ItemSet{
	Name: "Lightbringer Battlegear",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.RegisterAura(core.Aura{
				Label:    "Lightbringer Battlegear 2pc",
				Duration: core.NeverExpires,
				OnReset: func(aura *core.Aura, sim *core.Simulation) {
					aura.Activate(sim)
				},
				OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
					if !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
						return
					}
					if sim.RandomFloat("lightbringer 2pc") > 0.2 {
						return
					}
					character.AddMana(sim, 50, core.ActionID{SpellID: 38428}, true)
				},
			})
		},
		4: func(agent core.Agent) {
			// TODO: if we implemented hammer of wrath.. this ups dmg
		},
	},
}

// https://tbc.wowhead.com/item=27484/libram-of-avengement
func ApplyAvengementLibramEffect(agent core.Agent) {
	character := agent.GetCharacter()
	procAura := character.NewTemporaryStatsAura("Justice", core.ActionID{SpellID: 34260}, stats.Stats{stats.MeleeCrit: 53, stats.SpellCrit: 53}, time.Second*6)

	character.RegisterAura(core.Aura{
		Label:    "Libram Of Avengement",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
			if spell.ActionID.SpellID == JudgementOfBloodActionID.SpellID {
				procAura.Activate(sim)
			}
		},
	})
}

// Librams implemented in seals.go and judgement.go

// TODO: once we have judgement of command.. https://tbc.wowhead.com/item=33503/libram-of-divine-judgement
