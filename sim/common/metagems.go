package common

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	// Keep these in order by item ID.
	core.AddItemEffect(25893, ApplyMysticalSkyfireDiamond)
	core.AddItemEffect(25899, ApplyBrutalEarthstormDiamond)
	core.AddItemEffect(25901, ApplyInsightfulEarthstormDiamond)
	core.AddItemEffect(35503, ApplyEmberSkyfireDiamond)
	core.AddItemEffect(32410, ApplyThunderingSkyfireDiamond)

	// These are handled in character.go, but create empty effects so they are included in tests.
	core.AddItemEffect(34220, func(_ core.Agent) {}) // Chaotic Skyfire Diamond
	core.AddItemEffect(32409, func(_ core.Agent) {}) // Relentless Earthstorm Diamond
}

var MysticalSkyfireDiamondAuraID = core.NewAuraID()
var MysticFocusAuraID = core.NewAuraID()

func ApplyBrutalEarthstormDiamond(agent core.Agent) {
	agent.GetCharacter().PseudoStats.BonusDamage += 3
}

func ApplyMysticalSkyfireDiamond(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 320.0
		const dur = time.Second * 4
		const icdDur = time.Second * 35
		icd := core.NewICD()
		applyStatAura := character.NewTemporaryStatsAuraApplier(MysticFocusAuraID, core.ActionID{ItemID: 25893}, stats.Stats{stats.SpellHaste: hasteBonus}, dur)

		return core.Aura{
			ID: MysticalSkyfireDiamondAuraID,
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				if icd.IsOnCD(sim) || sim.RandomFloat("Mystical Skyfire Diamond") > 0.15 {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)
				applyStatAura(sim)
			},
		}
	})
}

var InsightfulEarthstormDiamondAuraID = core.NewAuraID()

func ApplyInsightfulEarthstormDiamond(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		icd := core.NewICD()
		const dur = time.Second * 15

		return core.Aura{
			ID: InsightfulEarthstormDiamondAuraID,
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				if icd.IsOnCD(sim) || sim.RandomFloat("Insightful Earthstorm Diamond") > 0.04 {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + dur)
				character.AddMana(sim, 300, core.ActionID{ItemID: 25901}, false)
			},
		}
	})
}

func ApplyEmberSkyfireDiamond(agent core.Agent) {
	agent.GetCharacter().AddStatDependency(stats.StatDependency{
		SourceStat:   stats.Intellect,
		ModifiedStat: stats.Intellect,
		Modifier: func(intellect float64, _ float64) float64 {
			return intellect * 1.02
		},
	})
}

var ThunderingSkyfireDiamondAuraID = core.NewAuraID()
var SkyfireSwiftnessAuraID = core.NewAuraID()

func ApplyThunderingSkyfireDiamond(agent core.Agent) {
	character := agent.GetCharacter()
	character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		const hasteBonus = 240.0
		const dur = time.Second * 6
		const icdDur = time.Second * 40
		icd := core.NewICD()
		applyStatAura := character.NewTemporaryStatsAuraApplier(SkyfireSwiftnessAuraID, core.ActionID{ItemID: 32410}, stats.Stats{stats.MeleeHaste: hasteBonus}, dur)

		ppmm := character.AutoAttacks.NewPPMManager(1.5)

		return core.Aura{
			ID: ThunderingSkyfireDiamondAuraID,
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				// Mask 68, melee or ranged auto attacks.
				if !spellEffect.Landed() || !spellEffect.ProcMask.Matches(core.ProcMaskWhiteHit) || spellEffect.IsPhantom {
					return
				}
				if icd.IsOnCD(sim) {
					return
				}
				if !ppmm.Proc(sim, spellEffect.IsMH(), spellEffect.OutcomeRollCategory.Matches(core.OutcomeRollCategoryRanged), "Thundering Skyfire Diamond") {
					return
				}
				icd = core.InternalCD(sim.CurrentTime + icdDur)
				applyStatAura(sim)
			},
		}
	})
}
