package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Right now, add the additional GCD + mana cost for shifting back to Moonkin form as a hack
// Consider adding moonkin shapeshift spell / form tracking to balance rotation instead
// Then we can properly incur Rebirth cost through additional Moonkin form spell cast
func (druid *Druid) NewRebirth(sim *core.Simulation) *core.SimpleCast {
	var manaCost float64 = 1611 + (521.4 * (1 - (float64(druid.Talents.NaturalShapeshifter) * 0.1)))

	rb := &core.SimpleCast{
		Cast: core.Cast{
			ActionID:  core.ActionID{SpellID: 26994},
			Character: druid.GetCharacter(),
			BaseCost: core.ResourceCost{
				Type:  stats.Mana,
				Value: manaCost,
			},
			Cost: core.ResourceCost{
				Type:  stats.Mana,
				Value: manaCost,
			},
			CastTime: time.Second*3 + time.Millisecond*500,
			GCD:      core.GCDDefault,
		},
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) { druid.RebirthUsed = true },
	}

	rb.Init(sim)

	return rb
}

func (druid *Druid) TryRebirth(sim *core.Simulation) bool {
	if druid.RebirthUsed {
		return false
	}

	cast := druid.NewRebirth(sim)
	if success := cast.StartCast(sim); !success {
		druid.WaitForMana(sim, cast.GetManaCost())
	}
	return true
}
