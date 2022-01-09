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
	actionID := core.ActionID{SpellID: 12051}

	mage.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: EvocationCooldownID,
		Cooldown:   cooldown,
		Type:       core.CooldownTypeMana,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			if character.IsOnCD(core.GCDCooldownID, sim.CurrentTime) {
				return false
			}
			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			if character.HasAura(core.InnervateAuraID) || character.HasAura(core.ManaTideTotemAuraID) {
				return false
			}

			curMana := character.CurrentMana()
			if curMana > manaThreshold {
				return false
			}

			if character.HasAura(core.BloodlustAuraID) && curMana > manaThreshold/2 {
				return false
			}

			if mage.isBlastSpamming {
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
			manaThreshold = mage.MaxMana() * 0.2

			return func(sim *core.Simulation, character *core.Character) {
				duration := time.Duration(float64(baseDuration) / character.CastSpeed())

				character.AddMana(sim, totalAmount, actionID, true)
				character.AddAura(sim, core.Aura{
					ID:       EvocationAuraID,
					ActionID: actionID,
					Expires:  sim.CurrentTime + duration,
				})
				character.Metrics.AddInstantCast(actionID)
				character.SetCD(EvocationCooldownID, sim.CurrentTime+cooldown)
				character.SetCD(core.GCDCooldownID, sim.CurrentTime+duration)
			}
		},
	})
}
