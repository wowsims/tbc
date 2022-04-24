package balance

import (
	"time"

	"github.com/wowsims/tbc/sim/common"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/druid"
)

func RegisterBalanceDruid() {
	core.RegisterAgentFactory(
		proto.Player_BalanceDruid{},
		proto.Spec_SpecBalanceDruid,
		func(character core.Character, options proto.Player) core.Agent {
			return NewBalanceDruid(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_BalanceDruid)
			if !ok {
				panic("Invalid spec value for Balance Druid!")
			}
			player.Spec = playerSpec
		},
	)
}

func NewBalanceDruid(character core.Character, options proto.Player) *BalanceDruid {
	balanceOptions := options.GetBalanceDruid()

	selfBuffs := druid.SelfBuffs{}
	if balanceOptions.Options.InnervateTarget != nil {
		selfBuffs.InnervateTarget = *balanceOptions.Options.InnervateTarget
	} else {
		selfBuffs.InnervateTarget.TargetIndex = -1
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
	*druid.Druid

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
	return moonkin.Druid
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

		OnPresimResult: func(presimResult proto.UnitMetrics, iterations int32, duration time.Duration) bool {
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

func (moonkin *BalanceDruid) OnGCDReady(sim *core.Simulation) {
	moonkin.tryUseGCD(sim)
}

func (moonkin *BalanceDruid) OnManaTick(sim *core.Simulation) {
	if moonkin.FinishedWaitingForManaAndGCDReady(sim) {
		moonkin.tryUseGCD(sim)
	}
}

func (moonkin *BalanceDruid) tryUseGCD(sim *core.Simulation) {
	if moonkin.useSurplusRotation {
		moonkin.manaTracker.Update(sim, moonkin.GetCharacter())

		// If we have enough mana to burn, use the surplus rotation.
		if moonkin.manaTracker.ProjectedManaSurplus(sim, moonkin.GetCharacter()) {
			moonkin.actRotation(sim, moonkin.surplusRotation)
		} else {
			moonkin.actRotation(sim, moonkin.primaryRotation)
		}
	} else {
		moonkin.actRotation(sim, moonkin.primaryRotation)
	}
}

func (moonkin *BalanceDruid) actRotation(sim *core.Simulation, rotation proto.BalanceDruid_Rotation) {
	// Activate shared druid behaviors
	// Use Rebirth at the beginning of the fight if flagged in rotation settings
	// Potentially allow options for "Time of cast" in future or default cast like 1 min into fight
	// Currently just casts at the beginning of encounter (with all CDs popped)
	if moonkin.useBattleRes && moonkin.TryRebirth(sim) {
		return
	}

	target := sim.GetPrimaryTarget()

	var spell *core.Spell

	if moonkin.ShouldCastFaerieFire(sim, target, rotation) {
		spell = moonkin.FaerieFire
	} else if moonkin.ShouldCastHurricane(sim, rotation) {
		spell = moonkin.Hurricane
	} else if moonkin.ShouldCastInsectSwarm(sim, target, rotation) {
		spell = moonkin.InsectSwarm
	} else if moonkin.ShouldCastMoonfire(sim, target, rotation) {
		spell = moonkin.Moonfire
	} else {
		switch rotation.PrimarySpell {
		case proto.BalanceDruid_Rotation_Starfire:
			spell = moonkin.Starfire8
		case proto.BalanceDruid_Rotation_Starfire6:
			spell = moonkin.Starfire6
		case proto.BalanceDruid_Rotation_Wrath:
			spell = moonkin.Wrath
		}
	}

	if success := spell.Cast(sim, target); !success {
		moonkin.WaitForMana(sim, spell.CurCast.Cost)
	}
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
