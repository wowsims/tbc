package paladin

import (
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core"
)

type Paladin struct {
	core.Character
}

func (paladin *Paladin) GetCharacter() *core.Character {
	return &paladin.Character
}

func (paladin *Paladin) AddRaidBuffs(buffs *proto.Buffs) {
	buffs.BlessingOfWisdom = proto.TristateEffect_TristateEffectImproved
	buffs.BlessingOfKings = true
	buffs.JudgementOfWisdom = true
	buffs.ImprovedSealOfTheCrusader = true
}
func (paladin *Paladin) AddPartyBuffs(buffs *proto.Buffs) {
}

func (paladin *Paladin) ChooseAction(sim *core.Simulation) core.AgentAction {
	return core.NewWaitAction(sim, paladin, core.NeverExpires) // makes the bot wait forever and do nothing.
}

func (paladin *Paladin) OnActionAccepted(*core.Simulation, core.AgentAction) {
}

func (paladin *Paladin) BuffUp(sim *core.Simulation) {
}

func (paladin *Paladin) Reset(sim *core.Simulation) {
}
