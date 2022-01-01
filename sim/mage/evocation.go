package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

var EvocationAuraID = core.NewAuraID()
var EvocationCooldownID = core.NewCooldownID()

func (mage *Mage) registerEvocationCD() {
	cooldown := time.Minute * 8
	manaThreshold := 0.0

	mage.AddMajorCooldown(core.MajorCooldown{
		ActionID:   core.ActionID{SpellID: 12051},
		CooldownID: EvocationCooldownID,
		Cooldown:   cooldown,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			if character.IsOnCD(core.GCDCooldownID, sim.CurrentTime) {
				return false
			}
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			if character.CurrentMana() > manaThreshold {
				return false
			}
			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			maxTicks := int32(4)
			if ItemSetTempestRegalia.CharacterHasSetBonus(&mage.Character, 2) {
				maxTicks++
			}

			numTicks := core.MaxInt32(0, core.MinInt32(maxTicks, mage.Options.EvocationTicks))
			if numTicks == 0 {
				numTicks = maxTicks
			}

			baseDuration := time.Duration(numTicks) * time.Second * 2
			totalAmount := float64(numTicks) * mage.MaxMana() * 0.15
			manaThreshold = float64(numTicks) * mage.MaxMana() * 0.15

			return func(sim *core.Simulation, character *core.Character) {
				duration := time.Duration(float64(baseDuration) / character.CastSpeed())

				character.AddMana(sim, totalAmount, "Evocation", true)
				character.AddAura(sim, core.Aura{
					ID:      EvocationAuraID,
					SpellID: 12051,
					Name:    "Evocation",
					Expires: sim.CurrentTime + duration,
				})
				character.Metrics.AddInstantCast(core.ActionID{SpellID: 12051})
				character.SetCD(EvocationCooldownID, sim.CurrentTime+cooldown)
				character.SetCD(core.GCDCooldownID, sim.CurrentTime+duration)
			}
		},
	})
}
