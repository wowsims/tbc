package warrior

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (warrior *Warrior) makeShoutSpellHelper(actionID core.ActionID) *core.Spell {
	cost := 10.0
	if ItemSetBoldArmor.CharacterHasSetBonus(&warrior.Character, 2) {
		cost -= 2
	}

	return warrior.RegisterSpell(core.SpellConfig{
		ActionID: actionID,

		ResourceType: stats.Rage,
		BaseCost:     cost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: cost,
				GCD:  core.GCDDefault,
			},
			IgnoreHaste:      true,
			DisableCallbacks: true,
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Target, _ *core.Spell) {
			// Actual shout effects are handled in core/buffs.go
			warrior.shoutExpiresAt = sim.CurrentTime + warrior.shoutDuration
		},
	})
}

func (warrior *Warrior) makeShoutSpell() *core.Spell {
	if warrior.ShoutType == proto.WarriorShout_WarriorShoutBattle {
		return warrior.makeShoutSpellHelper(core.ActionID{SpellID: 2048})
	} else if warrior.ShoutType == proto.WarriorShout_WarriorShoutCommanding {
		return warrior.makeShoutSpellHelper(core.ActionID{SpellID: 469})
	} else {
		return nil
	}
}

func (warrior *Warrior) ShouldShout(sim *core.Simulation) bool {
	return warrior.Shout != nil && warrior.CurrentRage() >= warrior.Shout.DefaultCast.Cost && sim.CurrentTime+time.Second*3 > warrior.shoutExpiresAt
}
