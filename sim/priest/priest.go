package priest

import "github.com/wowsims/tbc/sim/core"

func NewBuffBot(sim *core.Simulation, party *core.Party, misery bool, spriestDPS float64) *Priest {

	// shadow priest buff bot just statically applies mp5
	if spriestDPS > 0 {
		for _, pl := range party.Players {
			pl.InitialStats[core.StatMP5] += float64(spriestDPS) * 0.25
			pl.Stats = pl.InitialStats
		}
	}

	return &Priest{
		misery: misery,
	}
}

type Priest struct {
	misery bool
	core.Agent
}

func (p *Priest) ChooseAction(_ *core.Simulation, party *core.Party) core.AgentAction {
	return core.AgentAction{Wait: core.NeverExpires} // makes the bot wait forever and do nothing.
}

func (p *Priest) OnActionAccepted(*core.Simulation, core.AgentAction) {

}

func (p *Priest) BuffUp(sim *core.Simulation, party *core.Party) {
	if p.misery {
		sim.AddAura(sim, nil, MiseryAura())
	}
}

func (p *Priest) Reset(sim *core.Simulation)                            {}
func (p *Priest) OnSpellHit(*core.Simulation, *core.Player, *core.Cast) {}

func MiseryAura() core.Aura {
	return core.Aura{
		ID:      core.MagicIDMisery,
		Expires: core.NeverExpires,
		OnSpellHit: func(sim *core.Simulation, p *core.Player, c *core.Cast) {
			c.DidDmg *= 1.05
		},
	}
}
