package druid

import (
	"github.com/wowsims/tbc/items"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func NewBuffBot(sim *core.Simulation, party *core.Party, gotw, moonkin, ravenIdol bool) *Druid {
	druid := &Druid{
		Character: core.NewCharacter(items.EquipmentSpec{}, core.RaceBonusTypeNone, core.Consumes{}, stats.Stats{}),
		gotw: gotw,
		moonkin: moonkin,
		ravenIdol: ravenIdol,
	}
	druid.Character.Agent = druid
	return druid
}

type Druid struct {
	*core.Character

	gotw      bool
	moonkin   bool
	ravenIdol bool
}

func (druid *Druid) GetCharacter() *core.Character {
	return druid.Character
}

func (druid *Druid) BuffUp(sim *core.Simulation) {
	if druid.gotw {
		amount := 18.0

		// TODO: Pretty sure some of these dont stack with fort/ai/divine spirit
		sim.Raid.AddStats(stats.Stats{
			stats.Stamina: amount,
			stats.Agility: amount,
			stats.Strength: amount,
			stats.Intellect: amount,
			stats.Spirit: amount,
		})
	}

	if druid.moonkin {
		druid.GetCharacter().Party.AddStats(stats.Stats{
			stats.SpellCrit: 110.4,
		})
		if druid.ravenIdol {
			druid.GetCharacter().Party.AddStats(stats.Stats{
				stats.SpellCrit: 20,
			})
		}
	}
}

func (druid *Druid) OnSpellHit(sim *core.Simulation, cast *core.Cast) {
}
func (druid *Druid) ChooseAction(sim *core.Simulation) core.AgentAction {
	return core.AgentAction{Wait: core.NeverExpires} // makes the bot wait forever and do nothing.
}
func (druid *Druid) OnActionAccepted(sim *core.Simulation, action core.AgentAction) {
}
func (druid *Druid) Reset(newsim *core.Simulation) {
}
