package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	core.AddItemEffect(30664, ApplyLivingRootoftheWildheart)
	core.AddItemEffect(33510, ApplyIdoloftheUnseenMoon)
	core.AddItemEffect(32486, ApplyAshtongueTalisman)

	core.AddItemSet(&ItemSetMalorne)
	core.AddItemSet(&ItemSetNordrassil)
	core.AddItemSet(&ItemSetThunderheart)
}

var Malorne2PcAuraID = core.NewAuraID()

var ItemSetMalorne = core.ItemSet{
	Name: "Malorne Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				return core.Aura{
					ID: Malorne2PcAuraID,
					OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
						if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
							return
						}
						if !spellEffect.Landed() {
							return
						}
						if sim.RandomFloat("malorne 2p") > 0.05 {
							return
						}
						spellCast.Character.AddMana(sim, 120, core.ActionID{SpellID: 37295}, false)
					},
				}
			})
		},
		4: func(agent core.Agent) {
			// Currently this is handled in druid.go (reducing CD of innervate)
		},
	},
}

var ItemSetNordrassil = core.ItemSet{
	Name: "Nordrassil Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		4: func(agent core.Agent) {
			// Implemented in starfire.go.
		},
	},
}

var ItemSetThunderheart = core.ItemSet{
	Name: "Thunderheart Regalia",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// handled in moonfire.go in template construction
		},
		4: func(agent core.Agent) {
			// handled in starfire.go in template construction
		},
	},
}

var LivingRootoftheWildheartAuraID = core.NewAuraID()
var LunarBlessingAuraID = core.NewAuraID()

func ApplyLivingRootoftheWildheart(agent core.Agent) {
	const spellBonus = 209
	const dur = time.Second * 15

	druidAgent := agent.(Agent)
	druid := druidAgent.GetDruid()
	druid.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		applyStatAura := agent.GetCharacter().NewTemporaryStatsAuraApplier(LunarBlessingAuraID, core.ActionID{ItemID: 30664}, stats.Stats{stats.SpellPower: spellBonus}, dur)
		return core.Aura{
			ID: LivingRootoftheWildheartAuraID,
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				// technically only works while in moonkin form... but i think we can assume thats always true.
				if druid.Talents.MoonkinForm {
					if sim.RandomFloat("Living Root of the Wildheart") > 0.03 {
						return
					}
					applyStatAura(sim)
				}
			},
		}
	})
}

var IdoloftheUnseenMoonAuraID = core.NewAuraID()
var LunarGraceAuraID = core.NewAuraID()

func ApplyIdoloftheUnseenMoon(agent core.Agent) {
	const spellBonus = 140
	const dur = time.Second * 10
	actionID := core.ActionID{ItemID: 33510}

	druidAgent := agent.(Agent)
	druid := druidAgent.GetDruid()
	druid.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		applyStatAura := agent.GetCharacter().NewTemporaryStatsAuraApplier(LunarGraceAuraID, actionID, stats.Stats{stats.SpellPower: spellBonus}, dur)
		return core.Aura{
			ID: IdoloftheUnseenMoonAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellCast.ActionID.SpellID == SpellIDMoonfire {
					if sim.RandomFloat("Idol of the Unseen Moon") > 0.5 {
						return
					}
					applyStatAura(sim)
				}
			},
		}
	})
}

var AshtongueTalismanItemAuraID = core.NewAuraID()
var AshtongueTalismanAuraID = core.NewAuraID()

func ApplyAshtongueTalisman(agent core.Agent) {
	// Not in the game yet so cant test; this logic assumes that:
	// - does not affect the starfire which procs it
	// - can proc off of any completed cast, not just hits
	const spellBonus = 150
	const dur = time.Second * 8
	actionID := core.ActionID{ItemID: 32486}

	char := agent.GetCharacter()
	char.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		applyStatAura := agent.GetCharacter().NewTemporaryStatsAuraApplier(AshtongueTalismanAuraID, actionID, stats.Stats{stats.SpellPower: spellBonus}, dur)
		return core.Aura{
			ID: AshtongueTalismanItemAuraID,
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellCast.ActionID.SpellID == SpellIDSF8 || spellCast.ActionID.SpellID == SpellIDSF6 {
					if sim.RandomFloat("Ashtongue Talisman") > 0.25 {
						return
					}
					applyStatAura(sim)
				}
			},
		}
	})
}
