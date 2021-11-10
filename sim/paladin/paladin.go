package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core"
)

type Paladin struct {
	core.Character
}

func (paladin *Paladin) GetCharacter() *core.Character {
	return &paladin.Character
}

func (paladin *Paladin) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
}
func (paladin *Paladin) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {
}

func (paladin *Paladin) Reset(sim *core.Simulation) {
}

func (paladin *Paladin) Act(sim *core.Simulation) time.Duration {
	return core.NeverExpires // makes the bot wait forever and do nothing.
}
