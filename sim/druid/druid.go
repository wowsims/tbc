package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core"
)

type Druid struct {
	core.Character
}

func (druid *Druid) GetCharacter() *core.Character {
	return &druid.Character
}

func (druid *Druid) AddRaidBuffs(buffs *proto.Buffs) {
	// TODO: Use talents to check for imp gotw
	buffs.GiftOfTheWild = proto.TristateEffect_TristateEffectRegular
}
func (druid *Druid) AddPartyBuffs(buffs *proto.Buffs) {
	//buffs.Moonkin = proto.TristateEffect_TristateEffectRegular
	// check for idol of raven goddess equipped
}

func (druid *Druid) BuffUp(sim *core.Simulation) {
}

func (druid *Druid) Act(sim *core.Simulation) time.Duration {
	return core.NeverExpires // makes the bot wait forever and do nothing.
}
func (druid *Druid) Start(sim *core.Simulation) time.Duration {
	return druid.Act(sim)
}
func (druid *Druid) Reset(newsim *core.Simulation) {
}
