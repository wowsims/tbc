package priest

import (
	"github.com/wowsims/tbc/sim/api"
	"github.com/wowsims/tbc/sim/core"
)

type Priest struct {
	*core.Character
}

func (priest *Priest) GetCharacter() *core.Character {
	return priest.Character
}

func (priest *Priest) AddRaidBuffs(buffs *core.Buffs) {
	buffs.Misery = true
	buffs.DivineSpirit = api.TristateEffect_TristateEffectRegular
}
func (priest *Priest) AddPartyBuffs(buffs *core.Buffs) {
	buffs.ShadowPriestDPS += 0
}

func (p *Priest) ChooseAction(_ *core.Simulation) core.AgentAction {
	return core.AgentAction{Wait: core.NeverExpires} // makes the bot wait forever and do nothing.
}

func (p *Priest) OnActionAccepted(*core.Simulation, core.AgentAction) {

}

func (p *Priest) BuffUp(sim *core.Simulation) {
}

func (p *Priest) Reset(sim *core.Simulation)                                {}
func (p *Priest) OnSpellHit(*core.Simulation, *core.Cast) {}
