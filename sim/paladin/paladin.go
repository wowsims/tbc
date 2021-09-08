package paladin

import "github.com/wowsims/tbc/sim/core"

func NewBuffBot() *Paladin {
	return &Paladin{}
}

type Paladin struct {
	core.PlayerAgent
}

func (m *Paladin) BuffUp(sim *core.Simulation) {

}
