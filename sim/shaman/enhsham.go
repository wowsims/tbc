package shaman

func NewEnhancement(agent int, options map[string]string) *Enhancement {
	return &Enhancement{}
}

type Enhancement struct {
	Agent
}

// BuffUp lets you buff up all players in sim.
func (e *Enhancement) BuffUp(sim *core.Simulation) {

}
