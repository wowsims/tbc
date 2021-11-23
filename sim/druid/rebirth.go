package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

// Right now, add the additional GCD + mana cost for shifting back to Moonkin form as a hack
// Consider adding moonkin shapeshift spell / form tracking to balance rotation instead
// Then we can properly incur Rebirth cost through additional Moonkin form spell cast
func (druid *Druid) NewRebirth(sim *core.Simulation) *core.SimpleCast {
	var manaCost float64 = 1611 + (521.4 * (1 - (float64(druid.Talents.NaturalShapeshifter) * 0.1)))

	rb := &core.SimpleCast{
		Cast: core.Cast{
			Name:         "Rebirth",
			ActionID:     core.ActionID{SpellID: 26994},
			Character:    druid.GetCharacter(),
			BaseManaCost: manaCost,
			ManaCost:     manaCost,
			CastTime:     time.Second*3 + time.Millisecond*500,
		},
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) { druid.RebirthUsed = true },
	}

	rb.Init(sim)

	return rb
}

func (druid *Druid) TryRebirth(sim *core.Simulation) time.Duration {

	if druid.RebirthUsed {
		return 0
	}

	var cast *core.SimpleCast

	if !druid.RebirthUsed {
		cast = druid.NewRebirth(sim)
	}

	success := cast.StartCast(sim)
	if !success {
		regenTime := druid.TimeUntilManaRegen(cast.GetManaCost())
		return sim.CurrentTime + regenTime
	}
	return sim.CurrentTime + druid.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime)
}
