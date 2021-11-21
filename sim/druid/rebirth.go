package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

func (druid *Druid) NewRebirth(sim *core.Simulation) *core.SimpleCast {
	var manaCost float64 = 1611

	rb := &core.SimpleCast {
		Cast: core.Cast{
			Name:            "Rebirth",
			ActionID:        core.ActionID{SpellID: 26994},
			Character:       druid.GetCharacter(),
			BaseManaCost:    manaCost,
			ManaCost:        manaCost,
			CastTime:        time.Second * 2,
		},
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) { druid.Rebirth = true },
	}

	rb.Init(sim)

	return rb
}

func (druid *Druid) TryRebirth(sim *core.Simulation) time.Duration {

	var cast *core.SimpleCast

	if !druid.Rebirth {
		cast = druid.NewRebirth(sim);
	}

	if cast == nil {
		return 0 
	}

	success := cast.StartCast(sim)
	if !success {
		regenTime := druid.TimeUntilManaRegen(cast.GetManaCost())
		return sim.CurrentTime + regenTime
	}
	return sim.CurrentTime + druid.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime)
}
