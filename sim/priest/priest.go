package priest

import (
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core"
)

type Priest struct {
	core.Character
}

func (priest *Priest) GetCharacter() *core.Character {
	return &priest.Character
}

func (priest *Priest) AddRaidBuffs(buffs *proto.Buffs) {
	buffs.Misery = true
	buffs.DivineSpirit = proto.TristateEffect_TristateEffectRegular
}
func (priest *Priest) AddPartyBuffs(buffs *proto.Buffs) {
	buffs.ShadowPriestDps += 0
}

func (priest *Priest) ChooseAction(sim *core.Simulation) core.AgentAction {
	return core.NewWaitAction(sim, priest, core.NeverExpires) // makes the bot wait forever and do nothing.
}

func (priest *Priest) OnActionAccepted(*core.Simulation, core.AgentAction) {

}

func (priest *Priest) BuffUp(sim *core.Simulation) {
}

func (priest *Priest) Reset(sim *core.Simulation) {}
