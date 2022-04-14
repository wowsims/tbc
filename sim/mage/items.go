package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/common"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SerpentCoilBraidID = 30720

func init() {
	common.AddSimpleStatItemActiveEffect(19339, stats.Stats{stats.SpellHaste: 330}, time.Second*20, time.Minute*5, core.OffensiveTrinketSharedCooldownID) // MQG

	core.AddItemEffect(32488, ApplyAshtongueTalismanOfInsight)

	core.AddItemSet(&ItemSetAldorRegalia)
	core.AddItemSet(&ItemSetTirisfalRegalia)
	core.AddItemSet(&ItemSetTempestRegalia)

	// Even though these item effects are handled elsewhere, add them so they are
	// detected for automatic testing.
	core.AddItemEffect(SerpentCoilBraidID, func(core.Agent) {})
}

var ItemSetAldorRegalia = core.ItemSet{
	Name: "Aldor Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Interruption avoidance.
		},
		4: func(agent core.Agent) {
			// Reduces the cooldown on PoM/Blast Wave/Ice Block.
		},
	},
}

var ItemSetTirisfalRegalia = core.ItemSet{
	Name: "Tirisfal Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Increases the damage and mana cost of Arcane Blast by 20%.
			// Implemented in arcane_blast.go.
		},
		4: func(agent core.Agent) {
			// Your spell critical strikes grant you up to 70 spell damage for 6 sec.
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
				procAura := character.NewTemporaryStatsAura("Tirisfal 4pc Proc", core.ActionID{SpellID: 37443}, stats.Stats{stats.SpellPower: 70}, time.Second*6)
				return character.GetOrRegisterAura(core.Aura{
					Label: "Tirisfal 4pc",
					OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
						if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
							return
						}
						if spellEffect.Outcome.Matches(core.OutcomeCrit) {
							procAura.Activate(sim)
						}
					},
				})
			})
		},
	},
}

var ItemSetTempestRegalia = core.ItemSet{
	Name: "Tempest Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Increases the duratoin of your Evocation ability by 2 sec.
			// Implemented in mage.go.
		},
		4: func(agent core.Agent) {
			// Increases the damage of your Fireball, Frostbolt, and Arcane Missles abilities by 5%.
			// Implemented in the files for those spells.
		},
	},
}

func ApplyAshtongueTalismanOfInsight(agent core.Agent) {
	// Not in the game yet so cant test; this logic assumes that:
	// - No ICD.
	// - 50% proc rate.
	char := agent.GetCharacter()
	char.AddPermanentAura(func(sim *core.Simulation) *core.Aura {
		procAura := char.NewTemporaryStatsAura("Asghtongue Talisman Proc", core.ActionID{SpellID: 32488}, stats.Stats{stats.SpellHaste: 150}, time.Second*5)

		return char.GetOrRegisterAura(core.Aura{
			Label: "Ashtongue Talisman",
			OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					return
				}
				if !spellEffect.Outcome.Matches(core.OutcomeCrit) {
					return
				}

				if sim.RandomFloat("Ashtongue Talisman of Insight") > 0.5 {
					return
				}

				procAura.Activate(sim)
			},
		})
	})
}
