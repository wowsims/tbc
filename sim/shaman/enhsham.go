package shaman

import "github.com/wowsims/tbc/sim/core"

func NewEnhancement(p *core.Player, agent int, options map[string]string) *Enhancement {
	return &Enhancement{Player: p}
}

type Enhancement struct {
	shamanAgent
	*core.Player
}

// BuffUp lets you buff up all players in sim.
func (e *Enhancement) BuffUp(sim *core.Simulation) {

}
