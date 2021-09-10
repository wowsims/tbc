package paladin

import "github.com/wowsims/tbc/sim/core"

func NewBuffBot() *Paladin {

	// TODO: apply these
	// if b.ImprovedBlessingOfWisdom {
	// 	s[StatMP5] += 42
	// }
	// if b.ImpSealofCrusader {
	// 	s[StatSpellCrit] += 66.24 // 3% crit
	// }

	return &Paladin{}
}

type Paladin struct {
	core.PlayerAgent
}

func (m *Paladin) BuffUp(sim *core.Simulation) {

}
