package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (mage *Mage) registerManaGemsCD() {
	if !mage.Options.UseManaEmeralds {
		return
	}

	serpentCoilBraid := mage.HasTrinketEquipped(SerpentCoilBraidID)

	manaMultiplier := 1.0
	minManaGain := 2340.0
	maxManaGain := 2460.0
	if serpentCoilBraid {
		manaMultiplier = 1.25
		minManaGain *= manaMultiplier
		maxManaGain *= manaMultiplier
	}
	manaGainRange := maxManaGain - minManaGain

	mage.AddMajorCooldown(core.MajorCooldown{
		ActionID:   core.MageManaGemMCDActionID,
		CooldownID: core.ConjuredCooldownID,
		Cooldown:   time.Minute * 2,
		Priority:   core.CooldownPriorityDefault,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			if mage.remainingManaEmeralds == 0 {
				return false
			}
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			// Only pop if we have less than the max mana provided by the gem minus 1mp5 tick.
			totalRegen := character.ManaRegenPerSecondWhileCasting() * 5
			if character.MaxMana()-(character.CurrentMana()+totalRegen) < maxManaGain {
				return false
			}
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				// Restores 2340 to 2460 mana. (2 Min Cooldown)
				manaGain := minManaGain + (sim.RandomFloat("Mana Emerald") * manaGainRange)
				character.AddMana(sim, manaGain, "Mana Emerald", true)
				character.SetCD(core.ConjuredCooldownID, time.Minute*2+sim.CurrentTime)
				character.Metrics.AddInstantCast(core.MageManaGemMCDActionID)

				if serpentCoilBraid {
					mage.activateSerpentCoilBraid(sim)
				}

				mage.remainingManaEmeralds--
				if mage.remainingManaEmeralds == 0 {
					// Disable this cooldown so other mana consumes (potions / runes) know
					// they're free to activate.
					character.DisableMajorCooldown(core.MageManaGemMCDActionID)
				}
			}
		},
	})
}

var SerpentCoilBraidAuraID = core.NewAuraID()

func (mage *Mage) activateSerpentCoilBraid(sim *core.Simulation) {
	mage.AddAuraWithTemporaryStats(sim, SerpentCoilBraidAuraID, 37447, "Serpent-Coil Braid", stats.SpellPower, 225, time.Second*15)
}
