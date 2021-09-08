package mage

import "github.com/wowsims/tbc/sim/core"

func NewBuffBot() *Mage {
	return &Mage{}
}

type Mage struct {
	core.PlayerAgent
}

func (m *Mage) BuffUp(sim *core.Simulation) {

}
