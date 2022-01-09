package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

var InnervateCooldownID = core.NewCooldownID()

// Returns the time to wait before the next action, or 0 if innervate is on CD
// or disabled.
func (druid *Druid) registerInnervateCD() {
	actionID := core.ActionID{SpellID: 29166, Tag: int32(druid.RaidIndex)}
	baseCastTime := time.Millisecond * 1500
	baseManaCost := druid.BaseMana() * 0.04
	innervateCD := core.InnervateCD
	if ItemSetMalorne.CharacterHasSetBonus(druid.GetCharacter(), 4) {
		innervateCD -= time.Second * 48
	}

	var innervateTarget *core.Character
	expectedManaPerInnervate := 0.0
	innervateManaThreshold := 0.0
	remainingInnervateUsages := 0
	expectedBonusMana := 0.0

	druid.AddMajorCooldown(core.MajorCooldown{
		ActionID:   actionID,
		CooldownID: InnervateCooldownID,
		Cooldown:   innervateCD,
		Type:       core.CooldownTypeMana,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			if innervateTarget == nil {
				return false
			}

			if character.IsOnCD(core.GCDCooldownID, sim.CurrentTime) {
				return false
			}

			if character.CurrentMana() < baseManaCost {
				return false
			}

			// If target already has another innervate, don't cast.
			if innervateTarget.HasAura(core.InnervateAuraID) {
				return false
			}

			return true
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			// Innervate needs to be activated as late as possible to maximize DPS. The issue is that
			// innervate gives so much mana that it can cause Super Mana Potion or Dark Rune usages
			// to be delayed, if they come off CD soon after innervate. This delay is minimized by
			// activating innervate from the smallest amount of mana possible.
			if innervateTarget.CurrentMana() > innervateManaThreshold {
				return false
			}

			return true
		},
		ActivationFactory: func(sim *core.Simulation) core.CooldownActivation {
			innervateTargetAgent := sim.Raid.GetPlayerFromRaidTarget(druid.SelfBuffs.InnervateTarget)
			if innervateTargetAgent != nil {
				innervateTarget = innervateTargetAgent.GetCharacter()
				expectedManaPerInnervate = innervateTarget.SpiritManaRegenPerSecond() * 5 * 20
				if innervateTarget == druid.GetCharacter() {
					// Threshold can be lower when casting on self because its never mid-cast.
					innervateManaThreshold = 500
				} else {
					innervateManaThreshold = core.InnervateManaThreshold(innervateTarget)
				}

				remainingInnervateUsages = int(1 + (core.MaxDuration(0, sim.Duration))/innervateCD)
				expectedBonusMana += expectedManaPerInnervate * float64(remainingInnervateUsages)
			}

			return func(sim *core.Simulation, character *core.Character) {
				// Update expected bonus mana
				newRemainingUsages := int(sim.GetRemainingDuration() / innervateCD)
				expectedBonusManaReduction := expectedManaPerInnervate * float64(remainingInnervateUsages-newRemainingUsages)
				remainingInnervateUsages = newRemainingUsages

				castTime := time.Duration(float64(baseCastTime) / character.CastSpeed())

				core.AddInnervateAura(sim, innervateTarget, expectedBonusManaReduction, actionID.Tag)
				character.SpendMana(sim, baseManaCost, actionID)
				character.Metrics.AddInstantCast(actionID)
				character.SetCD(InnervateCooldownID, sim.CurrentTime+innervateCD)
				character.SetCD(core.GCDCooldownID, sim.CurrentTime+castTime)
			}
		},
	})
}
