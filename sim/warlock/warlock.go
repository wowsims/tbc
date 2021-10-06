package warlock

import "github.com/wowsims/tbc/sim/core"

func NewBuffBot(sim *core.Simulation, party *core.Party, malediction int) *Warlock {
	return &Warlock{
		malediction: malediction,
	}
}

type Warlock struct {
	core.Agent

	malediction int // bonus level of coe
}

func (p *Warlock) ChooseAction(_ *core.Simulation) core.AgentAction {
	return core.AgentAction{Wait: core.NeverExpires} // makes the bot wait forever and do nothing.
}

func (p *Warlock) OnActionAccepted(*core.Simulation, core.AgentAction) {

}

func (p *Warlock) BuffUp(sim *core.Simulation) {
	sim.AddAura(sim, CurseOfElementsAura(p.malediction))
}

func (p *Warlock) Reset(sim *core.Simulation)                                {}
func (p *Warlock) OnSpellHit(*core.Simulation, core.Agent, *core.Cast) {}

func CurseOfElementsAura(malediction int) core.Aura {
	multiplier := 1.10 + 0.1*float64(malediction)
	return core.Aura{
		ID:      core.MagicIDCurseOfElements,
		Expires: core.NeverExpires,
		OnSpellHit: func(sim *core.Simulation, cast *core.Cast) {
			cast.DidDmg *= multiplier
		},
	}
}
