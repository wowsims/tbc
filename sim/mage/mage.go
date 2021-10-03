package mage

import (
	"github.com/wowsims/tbc/items"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func NewBuffBot(sim *core.Simulation, party *core.Party, arcaneInt bool) *Mage {
	mage := &Mage{
		Character: core.NewCharacter(items.EquipmentSpec{}, core.RaceBonusTypeNone, core.Consumes{}, stats.Stats{}),
		arcaneInt: arcaneInt,
	}
	mage.Character.Agent = mage
	return mage
}

type Mage struct {
	*core.Character

	arcaneInt bool
}

func (mage *Mage) GetCharacter() *core.Character {
	return mage.Character
}

func (mage *Mage) BuffUp(sim *core.Simulation) {
	if mage.arcaneInt {
		sim.Raid.AddStats(stats.Stats{
			stats.Intellect: 40,
		})
	}
}

func (mage *Mage) OnSpellHit(sim *core.Simulation, cast *core.Cast) {
}
func (mage *Mage) ChooseAction(sim *core.Simulation) core.AgentAction {
	return core.AgentAction{Wait: core.NeverExpires} // makes the bot wait forever and do nothing.
}
func (mage *Mage) OnActionAccepted(sim *core.Simulation, action core.AgentAction) {
}
func (mage *Mage) Reset(newsim *core.Simulation) {
}
