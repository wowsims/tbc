package druid

import "github.com/wowsims/tbc/sim/core"

func NewBuffBot() *Druid {
	return &Druid{}
}

type Druid struct {
	core.PlayerAgent
}

func (m *Druid) BuffUp(sim *core.Simulation) {

}
