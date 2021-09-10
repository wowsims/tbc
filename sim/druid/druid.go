package druid

import "github.com/wowsims/tbc/sim/core"

func NewBuffBot() *Druid {

	// TODO: Add these buffs to the raid/party.

	// if b.GiftOfTheWild {
	// 	s[StatIntellect] += 18 // assumes improved gotw, rounded down to nearest int... not sure if that is accurate.
	// }
	// if b.Moonkin {
	// 	s[StatSpellCrit] += 110.4
	// 	if b.MoonkinRavenGoddess {
	// 		s[StatSpellCrit] += 20
	// 	}
	// }

	return &Druid{}
}

type Druid struct {
	core.PlayerAgent
}

func (m *Druid) BuffUp(sim *core.Simulation) {

}
