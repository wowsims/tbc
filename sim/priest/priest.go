package priest

import (
	"github.com/wowsims/tbc/items"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func NewBuffBot(sim *core.Simulation, party *core.Party, misery bool, spriestDPS float64) *Priest {
	priest := &Priest{
		Character: core.NewCharacter(items.EquipmentSpec{}, core.RaceBonusTypeNone, core.Consumes{}, stats.Stats{}),
		misery: misery,
		spriestDPS: spriestDPS,
	}
	priest.Character.Agent = priest
	return priest
}

type Priest struct {
	*core.Character

	misery bool
	spriestDPS float64
	core.Agent
}

func (priest *Priest) GetCharacter() *core.Character {
	return priest.Character
}

func (p *Priest) ChooseAction(_ *core.Simulation) core.AgentAction {
	return core.AgentAction{Wait: core.NeverExpires} // makes the bot wait forever and do nothing.
}

func (p *Priest) OnActionAccepted(*core.Simulation, core.AgentAction) {

}

func (p *Priest) BuffUp(sim *core.Simulation) {
	// shadow priest buff bot just statically applies mp5
	if p.spriestDPS > 0 {
		p.GetCharacter().Party.AddStats(stats.Stats{stats.MP5: float64(p.spriestDPS) * 0.25})
	}

	if p.misery {
		sim.AddAura(sim, nil, MiseryAura())
	}
}

func (p *Priest) Reset(sim *core.Simulation)                                {}
func (p *Priest) OnSpellHit(*core.Simulation, *core.Cast) {}

func MiseryAura() core.Aura {
	return core.Aura{
		ID:      core.MagicIDMisery,
		Expires: core.NeverExpires,
		OnSpellHit: func(sim *core.Simulation, agent core.Agent, c *core.Cast) {
			c.DidDmg *= 1.05
		},
	}
}
