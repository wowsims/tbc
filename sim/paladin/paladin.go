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

func (paladin *Paladin) AddRaidBuffs(buffs *proto.Buffs) {
	buffs.BlessingOfWisdom = proto.TristateEffect_TristateEffectImproved
	buffs.BlessingOfKings = true
}
func (paladin *Paladin) AddPartyBuffs(buffs *proto.Buffs) {
}

func (paladin *Paladin) Reset(sim *core.Simulation) {
}

func (paladin *Paladin) Act(sim *core.Simulation) time.Duration {
	return core.NeverExpires // makes the bot wait forever and do nothing.
}
