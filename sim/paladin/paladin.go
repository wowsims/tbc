package paladin

import (
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core"
)

type Paladin struct {
	*core.Character
}

func (paladin *Paladin) GetCharacter() *core.Character {
	return paladin.Character
}

func (paladin *Paladin) AddRaidBuffs(buffs *core.Buffs) {
	buffs.BlessingOfWisdom = proto.TristateEffect_TristateEffectImproved
	buffs.BlessingOfKings = true
	buffs.JudgementOfWisdom = true
	buffs.ImprovedSealOfTheCrusader = true
}
func (paladin *Paladin) AddPartyBuffs(buffs *core.Buffs) {
}

func (p *Paladin) ChooseAction(_ *core.Simulation) core.AgentAction {
	return core.AgentAction{Wait: core.NeverExpires} // makes the bot wait forever and do nothing.
}

func (p *Paladin) OnActionAccepted(*core.Simulation, core.AgentAction) {

}

func (p *Paladin) BuffUp(sim *core.Simulation) {
}

func (p *Paladin) Reset(sim *core.Simulation) {

}
func (p *Paladin) OnSpellHit(*core.Simulation, *core.Cast) {}
