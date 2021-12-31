package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

var EvocationAuraID = core.NewAuraID()
var EvocationCooldownID = core.NewCooldownID()

func (mage *Mage) registerEvocationCD() {
	cooldown := time.Minute * 8

	mage.AddMajorCooldown(core.MajorCooldown{
		CooldownID: EvocationCooldownID,
		Cooldown:   cooldown,
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			numTicks := 4
			if ItemSetTempestRegalia.CharacterHasSetBonus(&mage.Character, 2) {
				numTicks++
			}

			baseDuration := time.Duration(numTicks) * time.Second * 2
			totalAmount := float64(numTicks) * mage.MaxMana() * 0.15
			threshold := float64(numTicks) * mage.MaxMana() * 0.15

			return func(sim *core.Simulation, character *core.Character) bool {
				if character.CurrentMana() > threshold {
					return false
				}

				if character.IsOnCD(core.GCDCooldownID, sim.CurrentTime) {
					return false
				}

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
				return true
			}
		},
	})
}
