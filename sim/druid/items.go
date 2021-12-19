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

	core.AddItemSet(ItemSetMalorne)
	core.AddItemSet(ItemSetNordrassil)
	core.AddItemSet(ItemSetThunderheart)
}

var Malorne2PcAuraID = core.NewAuraID()

var ItemSetMalorne = core.ItemSet{
	Name:  "Malorne Rainment",
	Items: map[int32]struct{}{29093: {}, 29094: {}, 29091: {}, 29092: {}, 29095: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				return core.Aura{
					ID:   Malorne2PcAuraID,
					Name: "Malorne 2pc Bonus",
					OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
						if sim.RandomFloat("malorne 2p") > 0.05 {
							return
						}
						spellCast.Character.AddMana(sim, 120, "Malorne 2p Bonus", false)
					},
				}
			})
		},
		4: func(agent core.Agent) {
			// Currently this is handled in druid.go (reducing CD of innervate)
		},
	},
}

var Nordrassil4pAuraID = core.NewAuraID()

var ItemSetNordrassil = core.ItemSet{
	Name:  "Nordrassil Regalia",
	Items: map[int32]struct{}{30231: {}, 30232: {}, 30233: {}, 30234: {}, 30235: {}},
	Bonuses: map[int32]core.ApplyEffect{
		4: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddPermanentAura(func(sim *core.Simulation) core.Aura {
				return core.Aura{
					ID:   Nordrassil4pAuraID,
					Name: "Nordrassil 4p Bonus",
					OnBeforeSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
						agent, ok := agent.(Agent)
						if !ok {
							panic("why is a non-druid using nordassil regalia")
						}
						druid := agent.GetDruid()
						if spellCast.ActionID.SpellID == SpellIDSF8 || spellCast.ActionID.SpellID == SpellIDSF6 {
							// Check if moonfire/insectswarm is ticking on the target.
							// TODO: in a raid simulator we need to be able to see which dots are ticking from other druids.
							if (druid.MoonfireSpell.DotInput.IsTicking(sim) && druid.MoonfireSpell.Target.Index == spellEffect.Target.Index) ||
								(druid.InsectSwarmSpell.DotInput.IsTicking(sim) && druid.InsectSwarmSpell.Target.Index == spellEffect.Target.Index) {
								spellEffect.DamageMultiplier *= 1.1
							}
						}
					},
				}
			})
		},
	},
}

var ItemSetThunderheart = core.ItemSet{
	Name:  "Thunderheart Regalia",
	Items: map[int32]struct{}{31043: {}, 31035: {}, 31040: {}, 31046: {}, 31049: {}, 34572: {}, 34446: {}, 34555: {}},
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
		return core.Aura{
			ID:   LivingRootoftheWildheartAuraID,
			Name: "Living Root of the Wildheart",
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				// technically only works while in moonkin form... but i think we can assume thats always true.
				if druid.Talents.MoonkinForm {
					if sim.RandomFloat("Living Root of the Wildheart") > 0.03 {
						return
					}
					druid.AddAuraWithTemporaryStats(sim, LunarBlessingAuraID, 0, "Lunar Blessing", stats.SpellPower, spellBonus, dur)
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

	druidAgent := agent.(Agent)
	druid := druidAgent.GetDruid()
	druid.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID:   IdoloftheUnseenMoonAuraID,
			Name: "Idol of the Unseen Moon",
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellCast.ActionID.SpellID == SpellIDMF {
					if sim.RandomFloat("Idol of the Unseen Moon") > 0.5 {
						return
					}
					druid.AddAuraWithTemporaryStats(sim, LunarGraceAuraID, 0, "Lunar Blessing", stats.SpellPower, spellBonus, dur)
				}
			},
			OnSpellMiss: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellCast.ActionID.SpellID == SpellIDMF {
					if sim.RandomFloat("Idol of the Unseen Moon") > 0.5 {
						return
					}
					druid.AddAuraWithTemporaryStats(sim, LunarGraceAuraID, 0, "Lunar Blessing", stats.SpellPower, spellBonus, dur)
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

	char := agent.GetCharacter()
	char.AddPermanentAura(func(sim *core.Simulation) core.Aura {
		return core.Aura{
			ID:   AshtongueTalismanItemAuraID,
			Name: "Ashtongue Talisman of Equilibrium",
			OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellCast.ActionID.SpellID == SpellIDSF8 || spellCast.ActionID.SpellID == SpellIDSF6 {
					if sim.RandomFloat("Ashtongue Talisman") > 0.25 {
						return
					}
					char.AddAuraWithTemporaryStats(sim, AshtongueTalismanAuraID, 0, "Ashtongue Spellpower", stats.SpellPower, spellBonus, dur)
				}
			},
			OnSpellMiss: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
				if spellCast.ActionID.SpellID == SpellIDSF8 || spellCast.ActionID.SpellID == SpellIDSF6 {
					if sim.RandomFloat("Ashtongue Talisman") > 0.25 {
						return
					}
					char.AddAuraWithTemporaryStats(sim, AshtongueTalismanAuraID, 0, "Ashtongue Spellpower", stats.SpellPower, spellBonus, dur)
				}
			},
		}
	})
}
