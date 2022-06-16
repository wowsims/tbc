package warlock

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func init() {
	core.AddItemSet(&ItemSetOblivionRaiment)
	core.AddItemSet(&ItemSetVoidheartRaiment)
	core.AddItemSet(&ItemSetCorruptorRaiment)
	core.AddItemSet(&ItemSetMaleficRaiment)
}

var ItemSetOblivionRaiment = core.ItemSet{
	Name: "Oblivion Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// in pet.go constructor
		},
		4: func(agent core.Agent) {
			// in seed.go
		},
	},
}

var ItemSetVoidheartRaiment = core.ItemSet{
	Name: "Voidheart Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			warlock := agent.(WarlockAgent).GetWarlock()

			shadowBonus := warlock.NewTemporaryStatsAura("Shadowflame", core.ActionID{SpellID: 37377}, stats.Stats{stats.ShadowSpellPower: 135}, time.Second*15)
			fireBonus := warlock.NewTemporaryStatsAura("Shadowflame Hellfire", core.ActionID{SpellID: 39437}, stats.Stats{stats.ShadowSpellPower: 135}, time.Second*15)

			warlock.RegisterAura(core.Aura{
				Label:    "Voidheart Raiment 2pc",
				Duration: core.NeverExpires,
				OnReset: func(aura *core.Aura, sim *core.Simulation) {
					aura.Activate(sim)
				},
				OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
					if sim.RandomFloat("cycl4p") > 0.05 {
						return
					}
					if spell.SpellSchool.Matches(core.SpellSchoolShadow) {
						shadowBonus.Activate(sim)
					}
					if spell.SpellSchool.Matches(core.SpellSchoolFire) {
						fireBonus.Activate(sim)
					}
				},
			})
		},
		4: func(agent core.Agent) {
			// implemented in immolate.go
			// TODO: add to corruption.go
		},
	},
}

var ItemSetCorruptorRaiment = core.ItemSet{
	Name: "Corruptor Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// heals pet
		},
		4: func(agent core.Agent) {
			// TODO: increase corruption tick damage on target whenever shadowbolt hits.
		},
	},
}

var ItemSetMaleficRaiment = core.ItemSet{
	Name: "Malefic Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// heals... not implemented yet
		},
		4: func(agent core.Agent) {
			// Increases damage done by shadowbolt and incinerate by 6%.
			// Implemented in shadowbolt.go and incinerate.go
		},
	},
}
