package warlock

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

type Warlock struct {
	core.Character
	Talents proto.WarlockTalents

	Shadowbolt *core.Spell

	// for reference
	copymespell *core.Spell
	copymedot   *core.Dot
	copymeaura  *core.Aura
}

func (warlock *Warlock) GetCharacter() *core.Character {
	return &warlock.Character
}

func (warlock *Warlock) Init(sim *core.Simulation) {
	warlock.Shadowbolt = warlock.newShadowboltSpell(sim)
}

func (warlock *Warlock) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
}

func (warlock *Warlock) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {
}

func (warlock *Warlock) Reset(sim *core.Simulation) {}

func (warlock *Warlock) Act(sim *core.Simulation) time.Duration {
	return core.NeverExpires // makes the bot wait forever and do nothing.
}
