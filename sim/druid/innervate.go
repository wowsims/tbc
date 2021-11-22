package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

var InnervateCooldownID = core.NewCooldownID()

// TODO: This probably needs to allow for multiple innervates later
//  would need to solve the same issue we had as dots (maybe ID per user)
var InnervateAuraID = core.NewAuraID()

// Returns the time to wait before the next action, or 0 if innervate is on CD
// or disabled.
func (druid *Druid) TryInnervate(sim *core.Simulation) time.Duration {
	if !druid.SelfBuffs.Innervate || druid.GetRemainingCD(InnervateCooldownID, sim.CurrentTime) != 0 {
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

	cast := &core.SimpleCast{
		Cast: core.Cast{
			Name:         "Innervate",
			Character:    druid.GetCharacter(),
			BaseManaCost: 94,
			ManaCost:     94,
			Cooldown:     cd,

			ActionID: core.ActionID{
				SpellID: 29166,
				CooldownID: InnervateCooldownID,
			},
		},
		OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
			druid.PseudoStats.ForceFullSpiritRegen = true
			druid.PseudoStats.SpiritRegenMultiplier *= 5.0

			druid.AddAura(sim, core.Aura{
				ID:      InnervateAuraID,
				SpellID: 29166,
				Name:    "Innervate",
				Expires: sim.CurrentTime + time.Second*20,
				OnExpire: func(sim *core.Simulation) {
					druid.PseudoStats.ForceFullSpiritRegen = false
					druid.PseudoStats.SpiritRegenMultiplier /= 5.0
				},
			})
		},
	}
	cast.Init(sim)

	success := cast.StartCast(sim)
	if !success {
		regenTime := druid.TimeUntilManaRegen(cast.GetManaCost())
		return sim.CurrentTime + regenTime
	}
	return sim.CurrentTime + druid.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime)
}
