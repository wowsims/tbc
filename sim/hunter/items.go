package hunter

import (
	"github.com/wowsims/tbc/sim/core"
)

func init() {
	//core.AddItemEffect(19344, ApplyNaturalAlignmentCrystal)

	core.AddItemSet(ItemSetBeastLord)
	core.AddItemSet(ItemSetDemonStalker)
	core.AddItemSet(ItemSetRiftStalker)
	core.AddItemSet(ItemSetGronnstalker)
}

var BeastLord4PcAuraID = core.NewAuraID()
var ItemSetBeastLord = core.ItemSet{
	Name:  "Beast Lord Armor",
	Items: map[int32]struct{}{28228: {}, 27474: {}, 28275: {}, 27874: {}, 27801: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
		},
		4: func(agent core.Agent) {
			// Handled in kill_command.go
		},
	},
}

var ItemSetDemonStalker = core.ItemSet{
	Name:  "Demon Stalker Armor",
	Items: map[int32]struct{}{29081: {}, 29082: {}, 29083: {}, 29084: {}, 29085: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
		},
		4: func(agent core.Agent) {
			// Handled in multi_shot.go
		},
	},
}

var ItemSetRiftStalker = core.ItemSet{
	Name:  "Rift Stalker Armor",
	Items: map[int32]struct{}{30139: {}, 30140: {}, 30141: {}, 30142: {}, 30143: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
		},
		4: func(agent core.Agent) {
			// Handled in steady_shot.go
		},
	},
}

var ItemSetGronnstalker = core.ItemSet{
	Name:  "Gronnstalker's Armor",
	Items: map[int32]struct{}{31001: {}, 31003: {}, 31004: {}, 31005: {}, 31006: {}, 34443: {}, 34549: {}, 34570: {}},
	Bonuses: map[int32]core.ApplyEffect{
		2: func(agent core.Agent) {
			// Handled in rotation.go
		},
		4: func(agent core.Agent) {
			// Handled in steady_shot.go
		},
	},
}
