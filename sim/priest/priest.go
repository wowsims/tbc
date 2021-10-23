package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core"
)

type Priest struct {
	core.Character
}

func (priest *Priest) GetCharacter() *core.Character {
	return &priest.Character
}

func (priest *Priest) AddRaidBuffs(buffs *proto.Buffs) {
	buffs.Misery = true
	buffs.DivineSpirit = proto.TristateEffect_TristateEffectRegular
}
func (priest *Priest) AddPartyBuffs(buffs *proto.Buffs) {
}

func (priest *Priest) Reset(sim *core.Simulation) {}

func (priest *Priest) Act(sim *core.Simulation) time.Duration {
	return core.NeverExpires // makes the bot wait forever and do nothing.
}
