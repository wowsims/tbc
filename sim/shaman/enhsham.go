package shaman

import "github.com/wowsims/tbc/sim/core"

func NewEnhancement(character *core.Character, agent int, options map[string]string) *Enhancement {
	return &Enhancement{Character: character}
}

type Enhancement struct {
	core.Agent
	*core.Character
}

// BuffUp lets you buff up all players in sim.
func (e *Enhancement) BuffUp(sim *core.Simulation) {

}
