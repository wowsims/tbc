package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SerpentCoilBraidID = 30720

var MindQuickeningGemAuraID = core.NewAuraID()
var MindQuickeningGemCooldownID = core.NewCooldownID()

func init() {
	core.AddItemEffect(19339, core.MakeTemporaryStatsOnUseCDRegistration(
		MindQuickeningGemAuraID,
		stats.Stats{stats.SpellHaste: 330},
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 19339},
			CooldownID:       MindQuickeningGemCooldownID,
			Cooldown:         time.Minute * 5,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

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

var Tirisfal4PcAuraID = core.NewAuraID()
var Tirisfal4PcProcAuraID = core.NewAuraID()

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
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				applyStatAura := character.NewTemporaryStatsAuraApplier(Tirisfal4PcProcAuraID, core.ActionID{SpellID: 37443}, stats.Stats{stats.SpellPower: 70}, time.Second*6)
				return core.Aura{
					ID: Tirisfal4PcAuraID,
					OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
						if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
							return
						}
						if spellEffect.Outcome.Matches(core.OutcomeCrit) {
							applyStatAura(sim)
						}
					},
				}
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
		applyStatAura := char.NewTemporaryStatsAuraApplier(AshtongueTalismanOfInsightProcAuraID, core.ActionID{SpellID: 32488}, stats.Stats{stats.SpellHaste: hasteBonus}, dur)

		return core.Aura{
			ID: AshtongueTalismanOfInsightAuraID,
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					return
				}
				if !spellEffect.Outcome.Matches(core.OutcomeCrit) {
					return
				}

				if sim.RandomFloat("Ashtongue Talisman of Insight") > 0.5 {
					return
				}

				applyStatAura(sim)
			},
		}
	})
}
