package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

var InnervateCooldownID = core.NewCooldownID()

// Returns the time to wait before the next action, or 0 if innervate is on CD
// or disabled.
func (druid *Druid) TryInnervate(sim *core.Simulation) time.Duration {
	if druid.innervateTarget == nil || druid.GetRemainingCD(InnervateCooldownID, sim.CurrentTime) != 0 {
		return 0
	}

	// If target already has another innervate, don't cast.
	if druid.innervateTarget.HasAura(core.InnervateAuraID) {
		return 0
	}

	// Innervate needs to be activated as late as possible to maximize DPS. The issue is that
	// innervate gives so much mana that it can cause Super Mana Potion or Dark Rune usages
	// to be delayed, if they come off CD soon after innervate. This delay is minimized by
	// activating innervate from the smallest amount of mana possible.
	if druid.innervateTarget.CurrentMana() > druid.innervateManaThreshold {
		return 0
	}

	baseManaCost := druid.BaseMana() * 0.04

	// Update expected bonus mana
	newRemainingUsages := int(sim.GetRemainingDuration() / druid.innervateCD)
	expectedBonusManaReduction := druid.expectedManaPerInnervate * float64(druid.remainingInnervateUsages-newRemainingUsages)
	druid.remainingInnervateUsages = newRemainingUsages

	cast := &core.SimpleCast{
		Cast: core.Cast{
			Name:         "Innervate",
			Character:    druid.GetCharacter(),
			BaseManaCost: baseManaCost,
			ManaCost:     baseManaCost,
			Cooldown:     druid.innervateCD,

			ActionID: core.ActionID{
				SpellID:    29166,
				CooldownID: InnervateCooldownID,
			},
		},
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
			core.AddInnervateAura(sim, druid.innervateTarget, expectedBonusManaReduction)
		},
	}
	cast.Init(sim)

	success := cast.StartCast(sim)
	if !success {
		regenTime := druid.TimeUntilManaRegen(cast.GetManaCost())
		druid.Character.Metrics.MarkOOM(sim, &druid.Character, regenTime)
		return sim.CurrentTime + regenTime
	}
	return sim.CurrentTime + druid.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime)
}
