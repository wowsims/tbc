package druid

import (
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core"
)

type Druid struct {
	core.Character
}

func (druid *Druid) GetCharacter() *core.Character {
	return &druid.Character
}

func (druid *Druid) AddRaidBuffs(buffs *core.Buffs) {
	// TODO: Use talents to check for imp gotw
	buffs.GiftOfTheWild = proto.TristateEffect_TristateEffectRegular
}
func (druid *Druid) AddPartyBuffs(buffs *core.Buffs) {
	//buffs.Moonkin = proto.TristateEffect_TristateEffectRegular
	// check for idol of raven goddess equipped
}

func (druid *Druid) BuffUp(sim *core.Simulation) {
}

func (druid *Druid) ChooseAction(sim *core.Simulation) core.AgentAction {
	return core.NewWaitAction(sim, druid, core.NeverExpires) // makes the bot wait forever and do nothing.
}
func (druid *Druid) OnActionAccepted(sim *core.Simulation, action core.AgentAction) {
}
func (druid *Druid) Reset(newsim *core.Simulation) {
}
