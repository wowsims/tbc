package warlock

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

type Warlock struct {
	core.Character

	malediction int // bonus level of coe
}

func (warlock *Warlock) GetCharacter() *core.Character {
	return &warlock.Character
}

func (warlock *Warlock) AddRaidBuffs(buffs *proto.Buffs) {
	//sim.AddAura(sim, CurseOfElementsAura(warlock.malediction))
}
func (warlock *Warlock) AddPartyBuffs(buffs *proto.Buffs) {
}

func (warlock *Warlock) Reset(sim *core.Simulation) {}

func (warlock *Warlock) Act(sim *core.Simulation) time.Duration {
	return core.NeverExpires // makes the bot wait forever and do nothing.
}

func CurseOfElementsAura(malediction int) core.Aura {
	multiplier := 1.10 + 0.1*float64(malediction)
	return core.Aura{
		ID:      core.MagicIDCurseOfElements,
		Expires: core.NeverExpires,
		OnSpellHit: func(sim *core.Simulation, cast core.DirectCastAction, result *core.DirectCastDamageResult) {
			result.Damage *= multiplier
		},
	}
}
