package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

var EvocationCooldownID = core.NewCooldownID()

func (mage *Mage) registerEvocationCD() {
	cooldown := time.Minute * 8
	manaThreshold := 0.0
	actionID := core.ActionID{SpellID: 12051, CooldownID: EvocationCooldownID}

	maxTicks := int32(4)
	if ItemSetTempestRegalia.CharacterHasSetBonus(&mage.Character, 2) {
		maxTicks++
	}

	numTicks := core.MaxInt32(0, core.MinInt32(maxTicks, mage.Options.EvocationTicks))
	if numTicks == 0 {
		numTicks = maxTicks
	}

	castTime := time.Duration(numTicks) * time.Second * 2
	manaGain := 0.0

	template := core.SimpleCast{
		Cast: core.Cast{
			ActionID:  actionID,
			Character: mage.GetCharacter(),
			Cooldown:  time.Minute * 8,
			CastTime:  castTime,
			OnCastComplete: func(sim *core.Simulation, cast *core.Cast) {
				mage.AddMana(sim, manaGain, actionID, true)

				// All MCDs that use the GCD and have a non-zero cast time must call this.
				mage.UpdateMajorCooldowns()
			},
		},
	}

	mage.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: EvocationCooldownID,
		Cooldown:   cooldown,
		UsesGCD:    true,
		Type:       core.CooldownTypeMana,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
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
			manaGain = float64(numTicks) * mage.MaxMana() * 0.15
			manaThreshold = mage.MaxMana() * 0.2

			return func(sim *core.Simulation, character *core.Character) {
				cast := template
				cast.Init(sim)
				cast.StartCast(sim)
			}
		},
	})
}
