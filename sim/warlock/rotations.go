package warlock

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func (warlock *Warlock) OnGCDReady(sim *core.Simulation) {
	warlock.tryUseGCD(sim)
}

func (warlock *Warlock) OnManaTick(sim *core.Simulation) {
	if warlock.FinishedWaitingForManaAndGCDReady(sim) {
		warlock.tryUseGCD(sim)
	}
}

func (warlock *Warlock) tryUseGCD(sim *core.Simulation) {

	var spell *core.Spell

	// Apply curses first
	// TODO: should this be part of setup instead of during main rotation?
	switch warlock.Rotation.Curse {
	case proto.Warlock_Rotation_Elements:
		if !warlock.CurseOfElementsAura.IsActive() {
			if !warlock.CurseOfElements.Cast(sim, sim.GetPrimaryTarget()) {
				warlock.LifeTap.Cast(sim, sim.GetPrimaryTarget())
			}
			return
		}
	}

	// main spells
	if warlock.Rotation.Immolate && !warlock.ImmolateDot.IsActive() {
		spell = warlock.Immolate
	} else {
		switch warlock.Rotation.PrimarySpell {
		case proto.Warlock_Rotation_Shadowbolt:
			spell = warlock.Shadowbolt
		case proto.Warlock_Rotation_Incinerate:
			// spell = warlock.Incinerate
		default:
			panic("no primary spell set")
		}
	}

	if success := spell.Cast(sim, sim.GetPrimaryTarget()); success {
		return
	}

	warlock.LifeTap.Cast(sim, sim.GetPrimaryTarget())
}
