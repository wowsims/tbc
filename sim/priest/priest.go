package priest

import "github.com/wowsims/tbc/sim/core"

func NewBuffBot() *Priest {

	// shadow priest buff bot just statically applies mp5
	// s[StatMP5] += float64(b.SpriestDPS) * 0.25

	return &Priest{}
}

type Priest struct {
	core.PlayerAgent
}

func (m *Priest) BuffUp(sim *core.Simulation) {

}
