package core

import (
	"log"
	"math/rand"
	"testing"

	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
)

func TestAutoSwing(t *testing.T) {
	a := &FakeAgent{}
	c := &Character{
		Metrics: NewCharacterMetrics(),
		Equip: items.Equipment{
			proto.ItemSlot_ItemSlotMainHand: items.ByID[32262],
			proto.ItemSlot_ItemSlotOffHand:  items.ByID[32262],
		},
	}
	sim := &Simulation{
		rand:    rand.New(rand.NewSource(1)),
		Options: proto.SimOptions{},
		encounter: Encounter{
			Targets: []*Target{{}},
		},
		isTest:            true,
		testRands:         make(map[uint32]*rand.Rand),
		emptyAuras:        make([]Aura, numAuraIDs),
		pendingActionPool: newPAPool(),
	}

	c.EnableAutoAttacks(a, AutoAttackOptions{
		MainHand: c.WeaponFromMainHand(c.DefaultMeleeCritMultiplier()),
		OffHand:  c.WeaponFromOffHand(c.DefaultMeleeCritMultiplier()),
	})
	c.AutoAttacks.TrySwingMH(sim, sim.GetPrimaryTarget())
	c.AutoAttacks.TrySwingOH(sim, sim.GetPrimaryTarget())

	mhAuto := NewActionKey(ActionID{SpellID: 0, ItemID: 0, OtherID: 3, Tag: 1, CooldownID: 0})
	ohAuto := NewActionKey(ActionID{SpellID: 0, ItemID: 0, OtherID: 3, Tag: 2, CooldownID: 0})
	if c.Metrics.actions[mhAuto].Damage != 290.042434315968 {
		t.Fatalf("Failed... Expected: %f, Actual: %f", 290.042434315968, c.Metrics.actions[3.00390625].Damage)
	}
	if c.Metrics.actions[ohAuto].Damage != 111.03181478285845 {
		t.Fatalf("Failed... Expected: %f, Actual: %f", 290.042434315968, c.Metrics.actions[3.00390625].Damage)
	}
	log.Printf("%#v", c.Metrics.actions)
}

// The Character controlled by this Agent.
func (fa *FakeAgent) GetCharacter() *Character {
	panic("not implemented") // TODO: Implement
}

// Updates the input Buffs to include raid-wide buffs provided by this Agent.
func (fa *FakeAgent) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
	panic("not implemented") // TODO: Implement
}

// Updates the input Buffs to include party-wide buffs provided by this Agent.
func (fa *FakeAgent) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {
	panic("not implemented") // TODO: Implement
}

// Called once before the first iteration, after all Agents and Targets are finalized.
// Use this to do any precomputations that require access to Sim or Target fields.
func (fa *FakeAgent) Init(sim *Simulation) {
	panic("not implemented") // TODO: Implement
}

// Returns this Agent to its initial state. Called before each Sim iteration
// and once after the final iteration.
func (fa *FakeAgent) Reset(sim *Simulation) {
	panic("not implemented") // TODO: Implement
}

// Called whenever the GCD becomes ready for this Agent.
func (fa *FakeAgent) OnGCDReady(sim *Simulation) {
	panic("not implemented") // TODO: Implement
}

// Called after each mana tick, if this Agent uses mana.
func (fa *FakeAgent) OnManaTick(sim *Simulation) {
	panic("not implemented") // TODO: Implement
}

type FakeAgent struct {
	Character
}
