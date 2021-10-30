package mage

import (
	"time"

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

func (mage *Mage) Reset(newsim *core.Simulation) {
}

func (mage *Mage) Act(sim *core.Simulation) time.Duration {
	return core.NeverExpires // makes the bot wait forever and do nothing.
}
