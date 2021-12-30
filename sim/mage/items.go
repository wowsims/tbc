package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	var MindQuickeningGemCooldownID = core.NewCooldownID()
	core.AddItemEffect(19339, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		23723,
		"Mind Quickening Gem",
		stats.SpellHaste,
		330,
		time.Second*20,
		core.MajorCooldown{
			CooldownID:       MindQuickeningGemCooldownID,
			Cooldown:         time.Minute * 5,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	core.AddItemEffect(32488, ApplyAshtongueTalismanOfInsight)

	core.AddItemSet(ItemSetAldorRegalia)
	core.AddItemSet(ItemSetTirisfalRegalia)
	core.AddItemSet(ItemSetTempestRegalia)
}

var ItemSetAldorRegalia = core.ItemSet{
	Name:  "Aldor Regalia",
	Items: map[int32]struct{}{29076: {}, 29077: {}, 29078: {}, 29079: {}, 29080: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Interruption avoidance.
		},
		4: func(agent core.Agent) {
			// Reduces the cooldown on PoM/Blast Wave/Ice Block.
		},
	},
}

var Tirisfal4PcAuraID = core.NewAuraID()
var Tirisfal4PcProcAuraID = core.NewAuraID()

var ItemSetTirisfalRegalia = core.ItemSet{
	Name:  "Tirisfal Regalia",
	Items: map[int32]struct{}{30196: {}, 30205: {}, 30206: {}, 30207: {}, 30210: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Increases the damage and mana cost of Arcane Blast by 20%.
			// Implemented in arcane_blast.go.
		},
		4: func(agent core.Agent) {
			// Your spell critical strikes grant you up to 70 spell damage for 6 sec.
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				return core.Aura{
					ID: Tirisfal4PcAuraID,
					OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
						if spellEffect.Crit {
							character.AddAuraWithTemporaryStats(sim, Tirisfal4PcProcAuraID, 37743, "Tirisfal Crit Bonus Damage", stats.SpellPower, 70, time.Second*6)
						}
					},
				}
			})
		},
	},
}

var ItemSetTempestRegalia = core.ItemSet{
	Name:  "Tempest Regalia",
	Items: map[int32]struct{}{31055: {}, 31056: {}, 31057: {}, 31058: {}, 31059: {}, 34447: {}, 34557: {}, 34574: {}},
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

var AshtongueTalismanOfInsightAuraID = core.NewAuraID()
var AshtongueTalismanOfInsightProcAuraID = core.NewAuraID()

func ApplyAshtongueTalismanOfInsight(agent core.Agent) {
	// Not in the game yet so cant test; this logic assumes that:
	// - No ICD.
	// - 50% proc rate.
	const hasteBonus = 150
	const dur = time.Second * 5

	char := agent.GetCharacter()
	char.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID:   AshtongueTalismanOfInsightAuraID,
			Name: "Ashtongue Talisman of Insight",
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if !spellEffect.Crit {
					return
				}

				if sim.RandomFloat("Ashtongue Talisman of Insight") > 0.5 {
					return
				}

				char.AddAuraWithTemporaryStats(sim, AshtongueTalismanOfInsightProcAuraID, 40482, "Ashtongue Talisman of Insight", stats.SpellHaste, hasteBonus, dur)
			},
		}
	})
}
