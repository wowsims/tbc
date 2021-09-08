package core

import "time"

type Player struct {
	Spec PlayerAgent

	Consumes Consumes

	InitialStats Stats
	Stats        Stats
	Equip        Equipment // Current Gear
	EquipSpec    EquipmentSpec
	activeEquip  []*Item // cache of gear that can activate.

	CDs           [MagicIDLen]time.Duration // Map of MagicID to sim duration at which CD is done.
	auras         [MagicIDLen]Aura          // this is array instead of map to speed up browser perf.
	activeAuraIDs []int32                   // IDs of auras that are active, in no particular order
}

func (p *Player) HasteBonus() float64 {
	return 1 + (p.Stats[StatSpellHaste] / 1576)
}

type Party struct {
	Players []*Player
}

type Raid struct {
	Parties []Party
	Auras   [MagicIDLen]Aura // mostly just global debuffs
}

func NewPlayer(equip EquipmentSpec, consumes Consumes, spec PlayerAgent) Player {
	// TODO: configure player here.

	// equip := NewEquipmentSet(equipSpec)
	// initialStats := CalculateTotalStats(options, equip)

	return Player{}
}
