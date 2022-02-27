package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SerpentCoilBraidID = 30720

func init() {
	var MindQuickeningGemCooldownID = core.NewCooldownID()
	core.AddItemEffect(19339, core.MakeTemporaryStatsOnUseCDRegistration(
		core.OffensiveTrinketActiveAuraID,
		stats.SpellHaste,
		330,
		time.Second*20,
		core.MajorCooldown{
			ActionID:         core.ActionID{ItemID: 19339},
			CooldownID:       MindQuickeningGemCooldownID,
			Cooldown:         time.Minute * 5,
			SharedCooldownID: core.OffensiveTrinketSharedCooldownID,
		},
	))

	core.AddItemEffect(32488, ApplyAshtongueTalismanOfInsight)

	core.AddItemSet(ItemSetAldorRegalia)
	core.AddItemSet(ItemSetTirisfalRegalia)
	core.AddItemSet(ItemSetTempestRegalia)

	// Even though these item effects are handled elsewhere, add them so they are
	// detected for automatic testing.
	core.AddItemEffect(SerpentCoilBraidID, func(core.Agent) {})
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
				statApplier := character.NewTempStatAuraApplier(sim, Tirisfal4PcProcAuraID, core.ActionID{SpellID: 37443}, stats.SpellPower, 70, time.Second*6)
				return core.Aura{
					ID: Tirisfal4PcAuraID,
					OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
						if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
							return
						}
						if spellEffect.Outcome.Matches(core.OutcomeCrit) {
							statApplier(sim)
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
		statApplier := char.NewTempStatAuraApplier(sim, AshtongueTalismanOfInsightProcAuraID, core.ActionID{ItemID: 32488}, stats.SpellHaste, hasteBonus, dur)

		return core.Aura{
			ID: AshtongueTalismanOfInsightAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					return
				}
				if !spellEffect.Outcome.Matches(core.OutcomeCrit) {
					return
				}

				if sim.RandomFloat("Ashtongue Talisman of Insight") > 0.5 {
					return
				}

				statApplier(sim)
			},
		}
	})
}
