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
	minManaEmeraldGain := 2340.0
	maxManaEmeraldGain := 2460.0
	minManaRubyGain := 1073.0
	maxManaRubyGain := 1127.0
	if serpentCoilBraid {
		manaMultiplier = 1.25
		minManaEmeraldGain *= manaMultiplier
		maxManaEmeraldGain *= manaMultiplier
		minManaRubyGain *= manaMultiplier
		maxManaRubyGain *= manaMultiplier

		scbActionID := core.ActionID{ItemID: SerpentCoilBraidID}
		mage.serpentCoilAura = core.Aura{
			ID:       SerpentCoilBraidAuraID,
			ActionID: scbActionID,
			OnExpire: func(sim *core.Simulation) {
				if sim.Log != nil {
					mage.Log(sim, "Lost 225 SpellPower from fading "+scbActionID.String())
				}
				mage.AddStat(stats.SpellPower, -225)
			},
		}
	}
	manaEmeraldGainRange := maxManaEmeraldGain - minManaEmeraldGain
	manaRubyGainRange := maxManaRubyGain - minManaRubyGain

	mage.AddMajorCooldown(core.MajorCooldown{
		ActionID:   core.MageManaGemMCDActionID,
		CooldownID: core.ConjuredCooldownID,
		Cooldown:   time.Minute * 2,
		Priority:   core.CooldownPriorityDefault,
		Type:       core.CooldownTypeMana,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			if mage.remainingManaGems == 0 {
				return false
			}
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			// Only pop if we have less than the max mana provided by the gem minus 1mp5 tick.
			totalRegen := character.ManaRegenPerSecondWhileCasting() * 5
			maxManaGain := maxManaEmeraldGain
			if mage.remainingManaGems == 1 {
				maxManaGain = maxManaRubyGain
			}
			if character.MaxMana()-(character.CurrentMana()+totalRegen) < maxManaGain {
				return false
			}

			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			return func(sim *core.Simulation, character *core.Character) {
				if mage.remainingManaGems == 1 {
					// Mana Ruby: Restores 1073 to 1127 mana. (2 Min Cooldown)
					manaGain := minManaRubyGain + (sim.RandomFloat("Mana Gem") * manaRubyGainRange)
					character.AddMana(sim, manaGain, core.MageManaGemMCDActionID, true)
					character.SetCD(core.ConjuredCooldownID, time.Minute*2+sim.CurrentTime)
					character.Metrics.AddInstantCast(core.MageManaGemMCDActionID)
				} else {
					// Mana Emerald: Restores 2340 to 2460 mana. (2 Min Cooldown)
					manaGain := minManaEmeraldGain + (sim.RandomFloat("Mana Gem") * manaEmeraldGainRange)
					character.AddMana(sim, manaGain, core.MageManaGemMCDActionID, true)
					character.SetCD(core.ConjuredCooldownID, time.Minute*2+sim.CurrentTime)
					character.Metrics.AddInstantCast(core.MageManaGemMCDActionID)
				}

				if serpentCoilBraid {
					mage.activateSerpentCoilBraid(sim)
				}

				mage.remainingManaGems--
				if mage.remainingManaGems == 0 {
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
	const dur = time.Second * 15
	mage.AddStat(stats.SpellPower, 225)
	mage.serpentCoilAura.Expires = sim.CurrentTime + dur
	mage.AddAura(sim, mage.serpentCoilAura)
}
