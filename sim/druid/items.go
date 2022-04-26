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

var ItemSetMalorne = core.ItemSet{
	Name: "Malorne Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.RegisterAura(core.Aura{
				Label:    "Malorne Raiment 2pc",
				Duration: core.NeverExpires,
				OnReset: func(aura *core.Aura, sim *core.Simulation) {
					aura.Activate(sim)
				},
				OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
					if spellEffect.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
						return
					}
					if !spellEffect.Landed() {
						return
					}
					if sim.RandomFloat("malorne 2p") > 0.05 {
						return
					}
					spell.Unit.AddMana(sim, 120, core.ActionID{SpellID: 37295}, false)
				},
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

func ApplyLivingRootoftheWildheart(agent core.Agent) {
	druidAgent := agent.(Agent)
	druid := druidAgent.GetDruid()

	procAura := druid.NewTemporaryStatsAura("Living Root Proc", core.ActionID{ItemID: 30664}, stats.Stats{stats.SpellPower: 209}, time.Second*15)

	druid.RegisterAura(core.Aura{
		Label:    "Living Root of the Wildheart",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
			// technically only works while in moonkin form... but i think we can assume thats always true.
			if druid.Talents.MoonkinForm {
				if sim.RandomFloat("Living Root of the Wildheart") > 0.03 {
					return
				}
				procAura.Activate(sim)
			}
		},
	})
}

func ApplyIdoloftheUnseenMoon(agent core.Agent) {
	druidAgent := agent.(Agent)
	druid := druidAgent.GetDruid()

	actionID := core.ActionID{ItemID: 33510}
	procAura := druid.NewTemporaryStatsAura("Idol of the Unseen Moon Proc", actionID, stats.Stats{stats.SpellPower: 140}, time.Second*10)

	druid.RegisterAura(core.Aura{
		Label:    "Idol of the Unseen Moon",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spell == druid.Moonfire {
				if sim.RandomFloat("Idol of the Unseen Moon") > 0.5 {
					return
				}
				procAura.Activate(sim)
			}
		},
	})
}

func ApplyAshtongueTalisman(agent core.Agent) {
	druidAgent := agent.(Agent)
	druid := druidAgent.GetDruid()

	// Not in the game yet so cant test; this logic assumes that:
	// - does not affect the starfire which procs it
	// - can proc off of any completed cast, not just hits
	actionID := core.ActionID{ItemID: 32486}
	procAura := druid.NewTemporaryStatsAura("Ashtongue Talisman Proc", actionID, stats.Stats{stats.SpellPower: 150}, time.Second*8)

	druid.RegisterAura(core.Aura{
		Label:    "Ashtongue Talisman",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spell == druid.Starfire8 || spell == druid.Starfire6 {
				if sim.RandomFloat("Ashtongue Talisman") > 0.25 {
					return
				}
				procAura.Activate(sim)
			}
		},
	})
}
