package feral

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/druid"
)

func (cat *FeralDruid) OnGCDReady(sim *core.Simulation) {
	cat.doRotation(sim)
}

// Ported from https://github.com/NerdEgghead/TBC_cat_sim

func (cat *FeralDruid) shift(sim *core.Simulation) bool {
	cat.waitingForTick = false

	// If we have just now decided to shift, then we do not execute the
	// shift immediately, but instead trigger an input delay for realism.
	if !cat.readyToShift {
		cat.readyToShift = true
		return false
	}

	cat.readyToShift = false
	return cat.PowerShiftCat(sim)
}

func (cat *FeralDruid) doRotation(sim *core.Simulation) bool {
	// On gcd do nothing
	if !cat.GCD.IsReady(sim) {
		return false
	}

	// If we're out of form always shift back in
	if !cat.InForm(druid.Cat) {
		return cat.CatForm.Cast(sim, nil)
	}

	// If we previously decided to shift, then execute the shift now once
	// the input delay is over.
	if cat.readyToShift {
		return cat.shift(sim)
	}

	//strategy = {

	min_combos_for_rip := cat.Rotation.RipMinComboPoints
	use_mangle_trick := cat.Rotation.MangleTrick

	bite_trick_cp := int32(2)
	bite_trick_max := 39.0
	use_bite := (cat.Rotation.Biteweave && cat.Rotation.FinishingMove == proto.FeralDruid_Rotation_Rip) || cat.Rotation.FinishingMove == proto.FeralDruid_Rotation_Bite
	bite_time := time.Second * 0.0
	min_combos_for_bite := cat.Rotation.BiteMinComboPoints

	use_rip_trick := cat.Rotation.Ripweave
	use_bite_trick := cat.Rotation.RakeTrick
	use_rake_trick := cat.Rotation.RakeTrick

	rip_trick_cp := int32(4)
	rip_trick_min := 52.0
	max_wait_time := time.Second * 1.0

	bite_over_rip := use_bite && cat.Rotation.FinishingMove != proto.FeralDruid_Rotation_Rip
	no_finisher := false

	// }

	energy := cat.CurrentEnergy()
	cp := cat.ComboPoints()
	rip_cp := min_combos_for_rip
	bite_cp := min_combos_for_bite
	cur_time := sim.CurrentTime
	rip_debuff := cat.RipDot.IsActive()
	rip_end := cat.RipDot.ExpiresAt()
	mangle_debuff := cat.MangleAura.IsActive()
	mangle_end := cat.MangleAura.ExpiresAt()
	rake_debuff := cat.RakeDot.IsActive()
	next_tick := cat.NextEnergyTickAt()
	fight_length := sim.Duration
	wolfshead := cat.Equip[items.ItemSlotHead].ID == 8345
	shift_cost := cat.CatForm.DefaultCast.Cost
	omen_proc := cat.PseudoStats.NoCost
	latency := cat.latency

	// 10/6/21 - Added logic to not cast Rip if we're near the end of the
	// fight.
	end_thresh := time.Second * 10
	rip_now := cp >= rip_cp && !rip_debuff
	ripweave_now := (use_rip_trick &&
		cp >= rip_trick_cp &&
		!rip_debuff &&
		energy >= rip_trick_min &&
		!cat.PseudoStats.NoCost)

	rip_now = (rip_now || ripweave_now) && (fight_length-cur_time >= end_thresh)

	bite_at_end := (cp >= bite_cp &&
		!no_finisher && ((fight_length-cur_time < end_thresh) ||
		(rip_debuff && (fight_length-rip_end < end_thresh))))

	mangle_now := !rip_now && !mangle_debuff
	mangle_cost := cat.Mangle.DefaultCast.Cost

	bite_before_rip := (rip_debuff && use_bite &&
		(rip_end-cur_time >= bite_time))

	bite_now := ((bite_before_rip || bite_over_rip) &&
		cp >= bite_cp)

	rip_next := ((rip_now || ((cp >= rip_cp) && (rip_end <= next_tick))) &&
		(fight_length-next_tick >= end_thresh))

	mangle_next := (!rip_next && (mangle_now || mangle_end <= next_tick))

	// 12/2/21 - Added wait_to_mangle parameter that tells us whether we
	// should wait for the next Energy tick and cast Mangle, assuming we
	// are less than a tick's worth of Energy from being able to cast it. In
	// a standard Wolfshead rotation, wait_for_mangle is identical to
	// mangle_next, i.e. we only wait for the tick if Mangle will have
	// fallen off before the next tick. In a no-Wolfshead rotation, however,
	// it is preferable to Mangle rather than Shred as the second special in
	// a standard cycle, provided a bonus like 2pT6 is present to bring the
	// Mangle Energy cost down to 38 or below so that it can be fit in
	// alongside a Shred.
	wait_to_mangle := (mangle_next || ((!wolfshead) && (mangle_cost <= 38)))

	bite_before_rip_next := (bite_before_rip &&
		(rip_end-next_tick >= bite_time))

	prio_bite_over_mangle := (bite_over_rip || (!mangle_now))

	time_to_next_tick := next_tick - cur_time
	cat.waitingForTick = true

	if cat.CurrentMana() < shift_cost {
		// If this is the first time we're oom, log it
		//if self.time_to_oom is None:
		//    self.time_to_oom = time //TODO

		// No-shift rotation
		if rip_now && ((energy >= 30) || omen_proc) {
			cat.Rip.Cast(sim, cat.CurrentTarget)
			cat.waitingForTick = false
		} else if mangle_now &&
			((energy >= mangle_cost) || omen_proc) {
			return cat.Mangle.Cast(sim, cat.CurrentTarget)
		} else if bite_now && ((energy >= 35) || omen_proc) {
			return cat.FerociousBite.Cast(sim, cat.CurrentTarget)
		} else if (energy >= 42) || omen_proc {
			return cat.Shred.Cast(sim, cat.CurrentTarget)
		}
	} else if energy < 10 {
		cat.shift(sim)
	} else if rip_now {
		if (energy >= 30) || omen_proc {
			cat.Rip.Cast(sim, cat.CurrentTarget)
			cat.waitingForTick = false
		} else if time_to_next_tick > max_wait_time {
			cat.shift(sim)
		}
	} else if (bite_now || bite_at_end) && prio_bite_over_mangle {
		// Decision tree for Bite usage is more complicated, so there is
		// some duplicated logic with the main tree.

		// Shred versus Bite decision is the same as vanilla criteria.

		// Bite immediately if we'd have to wait for the following cast.
		cutoff_mod := 20.0
		if time_to_next_tick <= time.Second {
			cutoff_mod = 0.0
		}
		if (energy >= 57.0+cutoff_mod) ||
			((energy >= 15+cutoff_mod) && omen_proc) {
			return cat.Shred.Cast(sim, cat.CurrentTarget)
		}
		if energy >= 35 {
			return cat.FerociousBite.Cast(sim, cat.CurrentTarget)
		}
		// If we are doing the Rip rotation with Bite filler, then there is
		// a case where we would Bite now if we had enough energy, but once
		// we gain enough energy to do so, it's too late to Bite relative to
		// Rip falling off. In this case, we wait for the tick only if we
		// can Shred or Mangle afterward, and otherwise shift and won't Bite
		// at all this cycle. Returning 0.0 is the same thing as waiting for
		// the next tick, so this logic could be written differently if
		// desired to match the rest of the rotation code, where waiting for
		// tick is handled implicitly instead.
		wait := false
		if (energy >= 22) && bite_before_rip &&
			(!bite_before_rip_next) {
			wait = true
		} else if (energy >= 15) &&
			((!bite_before_rip) ||
				bite_before_rip_next || bite_at_end) {
			wait = true
		} else if (!rip_next) && ((energy < 20) || (!mangle_next)) {
			wait = false
			cat.shift(sim)
		} else {
			wait = true
		}
		if wait && (time_to_next_tick > max_wait_time) {
			cat.shift(sim)
		}
	} else if energy >= 35 && energy <= bite_trick_max &&
		use_bite_trick &&
		(time_to_next_tick > latency) &&
		!omen_proc &&
		cp >= bite_trick_cp {
		return cat.FerociousBite.Cast(sim, cat.CurrentTarget)
	} else if energy >= 35 && energy < mangle_cost &&
		use_rake_trick &&
		(time_to_next_tick > 1*time.Second+latency) &&
		!rake_debuff &&
		!omen_proc {
		return cat.Rake.Cast(sim, cat.CurrentTarget)
	} else if mangle_now {
		if (energy < mangle_cost-20) && (!rip_next) {
			cat.shift(sim)
		} else if (energy >= mangle_cost) || omen_proc {
			return cat.Mangle.Cast(sim, cat.CurrentTarget)
		} else if time_to_next_tick > max_wait_time {
			cat.shift(sim)
		}
	} else if energy >= 22 {
		if omen_proc {
			return cat.Shred.Cast(sim, cat.CurrentTarget)
		}
		// If our energy value is between 50-56 with 2pT6, or 60-61 without,
		// and we are within 1 second of an Energy tick, then Shredding now
		// forces us to shift afterwards, whereas we can instead cast two
		// Mangles instead for higher cpm. This scenario is most relevant
		// when using a no-Wolfshead rotation with 2pT6, and it will
		// occur whenever the initial Shred on a cycle misses.
		if (energy >= 2*mangle_cost-20) && (energy < 22+mangle_cost) &&
			(time_to_next_tick <= 1.0*time.Second) &&
			use_mangle_trick &&
			((!use_rake_trick && !use_bite_trick) || mangle_cost == 35) {
			return cat.Mangle.Cast(sim, cat.CurrentTarget)
		}
		if energy >= 42 {
			return cat.Shred.Cast(sim, cat.CurrentTarget)
		}
		if (energy >= mangle_cost) &&
			(time_to_next_tick > time.Second+latency) {
			return cat.Mangle.Cast(sim, cat.CurrentTarget)
		}
		if time_to_next_tick > max_wait_time {
			cat.shift(sim)
		}
	} else if (!rip_next) && ((energy < mangle_cost-20) || (!wait_to_mangle)) {
		cat.shift(sim)
	} else if time_to_next_tick > max_wait_time {
		cat.shift(sim)
	}
	// Model two types of input latency: (1) When waiting for an energy tick
	// to execute the next special ability, the special will in practice be
	// slightly delayed after the tick arrives. (2) When executing a
	// powershift without clipping the GCD, the shift will in practice be
	// slightly delayed after the GCD ends.

	if cat.readyToShift {
		cat.SetGCDTimer(sim, sim.CurrentTime+latency)
	} else if cat.waitingForTick {
		cat.SetGCDTimer(sim, sim.CurrentTime+time_to_next_tick+latency)
	}

	return false
}
