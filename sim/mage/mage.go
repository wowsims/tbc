package mage

import "github.com/wowsims/tbc/sim/core"

func NewBuffBot() *Mage {
	// TODO: apply to raid
	// if b.ArcaneInt {
	// 	s[StatIntellect] += 40
	// }

	return &Mage{}
}

type Mage struct {
	core.PlayerAgent
}

func (m *Mage) BuffUp(sim *core.Simulation) {

}
