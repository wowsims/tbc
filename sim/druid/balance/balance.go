package balance

import (
	"time"

	"github.com/wowsims/tbc/sim/common"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/druid"
)

func RegisterBalanceDruid() {
	core.RegisterAgentFactory(proto.Player_BalanceDruid{}, func(character core.Character, options proto.Player) core.Agent {
		return NewBalanceDruid(character, options)
	})
}

func NewBalanceDruid(character core.Character, options proto.Player) *BalanceDruid {
	balanceOptions := options.GetBalanceDruid()

	selfBuffs := druid.SelfBuffs{}
	if balanceOptions.Options.InnervateTarget != nil {
		selfBuffs.Innervate = balanceOptions.Options.InnervateTarget.TargetIndex == int32(character.RaidIndex)
	}

	druid := druid.New(character, selfBuffs, *balanceOptions.Talents)
	moonkin := &BalanceDruid{
		Druid:           druid,
		primaryRotation: *balanceOptions.Rotation,
		useBattleRes:    balanceOptions.Options.BattleRes,
	}

	return moonkin
}

type BalanceDruid struct {
	druid.Druid

	primaryRotation proto.BalanceDruid_Rotation
	useBattleRes    bool

	// These are only used when primary spell is set to 'Adaptive'. When the mana
	// tracker tells us we have extra mana to spare, use surplusRotation instead of
	// primaryRotation.
	useSurplusRotation bool
	surplusRotation    proto.BalanceDruid_Rotation
	manaTracker        common.ManaSpendingRateTracker
}

// GetDruid is to implement druid.Agent (supports nordrassil set bonus)
func (moonkin *BalanceDruid) GetDruid() *druid.Druid {
	return &moonkin.Druid
}

func (moonkin *BalanceDruid) GetPresimOptions() *core.PresimOptions {
	// If not adaptive, just use the primary rotation directly.
	if moonkin.primaryRotation.PrimarySpell != proto.BalanceDruid_Rotation_Adaptive {
		return nil
	}

	rotations := moonkin.GetDpsRotationHierarchy(moonkin.primaryRotation)
	rotationIdx := 0

	return &core.PresimOptions{
		SetPresimPlayerOptions: func(player *proto.Player) {
			*player.Spec.(*proto.Player_BalanceDruid).BalanceDruid.Rotation = rotations[rotationIdx]
		},

		OnPresimResult: func(presimResult proto.PlayerMetrics, iterations int32, duration time.Duration) bool {
			if float64(presimResult.SecondsOomAvg) <= 0.03*duration.Seconds() {
				moonkin.primaryRotation = rotations[rotationIdx]

				// If the highest dps rotation is fine, we dont need any adaptive logic.
				if rotationIdx == 0 {
					return true
				}

				moonkin.useSurplusRotation = true
				moonkin.surplusRotation = rotations[rotationIdx-1]
				moonkin.manaTracker = common.NewManaSpendingRateTracker()

				return true
			}

			rotationIdx++
			if rotationIdx == len(rotations) {
				// If we are here than all of the rotations went oom. No adaptive logic needed, just use the lowest one.
				moonkin.primaryRotation = rotations[len(rotations)-1]
				return true
			}

			return false
		},
	}
}

func (moonkin *BalanceDruid) Reset(sim *core.Simulation) {
	if moonkin.useSurplusRotation {
		moonkin.manaTracker.Reset()
	}
	moonkin.Druid.Reset(sim)
}

func (moonkin *BalanceDruid) Act(sim *core.Simulation) time.Duration {
	if moonkin.useSurplusRotation {
		moonkin.manaTracker.Update(sim, moonkin.GetCharacter())

		// If we have enough mana to burn, use the surplus rotation.
		if moonkin.manaTracker.ProjectedManaSurplus(sim, moonkin.GetCharacter()) {
			return moonkin.actRotation(sim, moonkin.surplusRotation)
		} else {
			return moonkin.actRotation(sim, moonkin.primaryRotation)
		}
	} else {
		return moonkin.actRotation(sim, moonkin.primaryRotation)
	}
}

func (moonkin *BalanceDruid) actRotation(sim *core.Simulation, rotation proto.BalanceDruid_Rotation) time.Duration {
	// Activate shared druid behaviors
	// Use Rebirth at the beginning of the fight if flagged in rotation settings
	// Potentially allow options for "Time of cast" in future or default cast like 1 min into fight
	// Currently just casts at the beginning of encounter (with all CDs popped)
	if moonkin.useBattleRes {
		rebirthTime := moonkin.TryRebirth(sim)
		if rebirthTime > 0 {
			return rebirthTime
		}
	}
	innervateWait := moonkin.TryInnervate(sim)
	if innervateWait != 0 {
		return innervateWait
	}

	target := sim.GetPrimaryTarget()

	var spell *core.SimpleSpell

	if moonkin.ShouldCastFaerieFire(sim, target, rotation) {
		spell = moonkin.NewFaerieFire(sim, target)
	} else if moonkin.ShouldCastInsectSwarm(sim, target, rotation) {
		spell = moonkin.NewInsectSwarm(sim, target)
	} else if moonkin.ShouldCastMoonfire(sim, target, rotation) {
		spell = moonkin.NewMoonfire(sim, target)
	} else {
		switch rotation.PrimarySpell {
		case proto.BalanceDruid_Rotation_Starfire:
			spell = moonkin.NewStarfire(sim, target, 8)
		case proto.BalanceDruid_Rotation_Starfire6:
			spell = moonkin.NewStarfire(sim, target, 6)
		case proto.BalanceDruid_Rotation_Wrath:
			spell = moonkin.NewWrath(sim, target)
		}
	}

	actionSuccessful := spell.Cast(sim)

	if !actionSuccessful {
		regenTime := moonkin.TimeUntilManaRegen(spell.GetManaCost())
		waitAction := core.NewWaitAction(sim, moonkin.GetCharacter(), regenTime, core.WaitReasonOOM)
		waitAction.Cast(sim)
		return sim.CurrentTime + waitAction.GetDuration()
	}

	return sim.CurrentTime + core.MaxDuration(
		moonkin.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime),
		spell.CastTime)
}

// Returns the order of DPS rotations to try, from highest to lowest dps. The
// lower DPS rotations are more mana efficient.
//
// Rotation tiers, from highest dps to lowest:
//  - SF8 + MF
//  - SF6 + MF
//  - SF6, or SF6 + IS if 4p T5 is worn.
func (moonkin *BalanceDruid) GetDpsRotationHierarchy(baseRotation proto.BalanceDruid_Rotation) []proto.BalanceDruid_Rotation {
	rotations := []proto.BalanceDruid_Rotation{}

	currentRotation := baseRotation
	currentRotation.PrimarySpell = proto.BalanceDruid_Rotation_Starfire
	currentRotation.Moonfire = true
	rotations = append(rotations, currentRotation)

	currentRotation.PrimarySpell = proto.BalanceDruid_Rotation_Starfire6
	rotations = append(rotations, currentRotation)

	if druid.ItemSetNordrassil.CharacterHasSetBonus(&moonkin.Character, 4) {
		currentRotation.Moonfire = false
		currentRotation.InsectSwarm = true
		rotations = append(rotations, currentRotation)
	} else {
		currentRotation.Moonfire = false
		rotations = append(rotations, currentRotation)
	}

	return rotations
}
