package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

// Returns the time to wait before the next action, or 0 if innervate is on CD
// or disabled.
func (druid *Druid) TryInnervate(sim *core.Simulation) time.Duration {
	if !druid.SelfBuffs.Innervate || druid.GetRemainingCD(core.InnervateCooldownID, sim.CurrentTime) != 0 {
		return 0
	}

	// Innervate needs to be activated as late as possible to maximize DPS. The issue is that
	// innervate gives so much mana that it can cause Super Mana Potion or Dark Rune usages
	// to be delayed, if they come off CD soon after innervate. This delay is minimized by
	// activating innervate from the smallest amount of mana possible.
	//
	// 500 mana is enough to cast our most expensive spell (moonfire).
	if druid.CurrentMana() > 500 {
		return 0
	}

	cd := time.Minute * 6
	if druid.malorne4p {
		cd -= time.Second * 48
	}

	baseManaCost := druid.BaseMana() * 0.04

	// Update expected bonus mana
	newRemainingUsages := int(sim.GetRemainingDuration() / cd)
	expectedBonusManaReduction := druid.ExpectedManaPerInnervate * float64(druid.RemainingInnervateUsages-newRemainingUsages)
	druid.RemainingInnervateUsages = newRemainingUsages

	cast := &core.SimpleCast{
		Cast: core.Cast{
			Name:         "Innervate",
			Character:    druid.GetCharacter(),
			BaseManaCost: baseManaCost,
			ManaCost:     baseManaCost,
			Cooldown:     cd,

			ActionID: core.ActionID{
				SpellID:    29166,
				CooldownID: core.InnervateCooldownID,
			},
		},
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
			core.AddInnervateAura(sim, druid.GetCharacter(), expectedBonusManaReduction)
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
