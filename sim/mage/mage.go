package mage

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

type Mage struct {
	core.Character
}

func (mage *Mage) GetCharacter() *core.Character {
	return &mage.Character
}

func (mage *Mage) AddRaidBuffs(buffs *proto.Buffs) {
	buffs.ArcaneBrilliance = true
}
func (mage *Mage) AddPartyBuffs(buffs *proto.Buffs) {
}

func (mage *Mage) BuffUp(sim *core.Simulation) {
}

func (mage *Mage) ChooseAction(sim *core.Simulation) core.AgentAction {
	return core.NewWaitAction(sim, mage, core.NeverExpires) // makes the bot wait forever and do nothing.
}
func (mage *Mage) OnActionAccepted(sim *core.Simulation, action core.AgentAction) {
}
func (mage *Mage) Reset(newsim *core.Simulation) {
}
