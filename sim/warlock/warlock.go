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

func (warlock *Warlock) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
	//sim.AddAura(sim, CurseOfElementsAura(warlock.malediction))
}
func (warlock *Warlock) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {
}

func (warlock *Warlock) Reset(sim *core.Simulation) {}

func (warlock *Warlock) Act(sim *core.Simulation) time.Duration {
	return core.NeverExpires // makes the bot wait forever and do nothing.
}

var CurseOfElementsAuraID = core.NewAuraID()
func CurseOfElementsAura(malediction int) core.Aura {
	multiplier := 1.10 + 0.1*float64(malediction)
	return core.Aura{
		ID:      CurseOfElementsAuraID,
		Expires: core.NeverExpires,
		OnSpellHit: func(sim *core.Simulation, cast *core.Cast, result *core.DirectCastDamageResult) {
			result.Damage *= multiplier
		},
	}
}
