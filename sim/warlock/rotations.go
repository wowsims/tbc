package warlock

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
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
	var target = sim.GetPrimaryTarget()

	// Apply curses first
	// TODO: should this be part of setup instead of during main rotation?
	switch warlock.Rotation.Curse {
	case proto.Warlock_Rotation_Elements:
		if !warlock.CurseOfElementsAura.IsActive() {
			if !warlock.CurseOfElements.Cast(sim, target) {
				warlock.LifeTap.Cast(sim, target)
			}
			return
		}
	}

	bigCDs := warlock.GetMajorCooldowns()
	nextBigCD := sim.Duration
	for _, cd := range bigCDs {
		if cd == nil {
			continue // not on cooldown right now.
		}
		cdReadyAt := cd.Cooldown.ReadyAt()
		if cd.Type == core.CooldownTypeDPS && cdReadyAt < nextBigCD {
			nextBigCD = cdReadyAt
		}
	}

	// If big CD coming up and we don't have enough mana for it, lifetap
	// Also, never do a big regen in the last few seconds of the fight.
	if !warlock.DoingRegen && nextBigCD-sim.CurrentTime < time.Second*15 && sim.Duration-sim.CurrentTime > time.Second*20 {
		if warlock.GetStat(stats.SpellPower) > warlock.GetInitialStat(stats.SpellPower) || warlock.CastSpeed() > warlock.InitialCastSpeed() {
			// never start regen if you have boosted sp or boosted cast speed
		} else if warlock.CurrentManaPercent() < 0.2 {
			warlock.DoingRegen = true
			// Try to make sure at least immolate is ticking while doing regen.
			if warlock.ImmolateDot.RemainingDuration(sim) < time.Second*10 {
				if sucess := warlock.Immolate.Cast(sim, target); sucess {
					return
				}
			}
		}
	}

	if warlock.DoingRegen {
		if nextBigCD-sim.CurrentTime < time.Second*2 {
			// stop regen, start blasting
			warlock.DoingRegen = false
		} else {
			warlock.LifeTap.Cast(sim, target)
			if warlock.CurrentManaPercent() > 0.6 {
				warlock.DoingRegen = false
			}
			return
		}
	}

	// main spells
	// TODO: optimize so that cast time of DoT is included in calculation so you can cast right before falling off.
	if warlock.Talents.UnstableAffliction && !warlock.UnstableAffDot.IsActive() {
		spell = warlock.UnstableAff
	} else if warlock.Rotation.Corruption && !warlock.CorruptionDot.IsActive() {
		spell = warlock.Corruption
	} else if warlock.Talents.SiphonLife && !warlock.SiphonLifeDot.IsActive() && warlock.ImpShadowboltAura.IsActive() {
		spell = warlock.SiphonLife
	} else if warlock.Rotation.Immolate && !warlock.ImmolateDot.IsActive() {
		spell = warlock.Immolate
	} else {
		switch warlock.Rotation.PrimarySpell {
		case proto.Warlock_Rotation_Shadowbolt:
			spell = warlock.Shadowbolt
		case proto.Warlock_Rotation_Incinerate:
			spell = warlock.Incinerate
		default:
			panic("no primary spell set")
		}
	}

	if success := spell.Cast(sim, target); success {
		return
	}

	// If we were not successful at anything else, lifetap.
	warlock.LifeTap.Cast(sim, target)
}
