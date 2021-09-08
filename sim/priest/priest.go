package priest

import "github.com/wowsims/tbc/sim/core"

func NewBuffBot() *Priest {
	return &Priest{}
}

type Priest struct {
	core.PlayerAgent
}

func (m *Priest) BuffUp(sim *core.Simulation) {

}
