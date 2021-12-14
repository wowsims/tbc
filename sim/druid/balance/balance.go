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
		// if targetting myself for individual sim
		// TODO: what is my player idx for raid?
		selfBuffs.Innervate = balanceOptions.Options.InnervateTarget.TargetIndex == 0
	}

	druid := druid.New(character, selfBuffs, *balanceOptions.Talents)
	moonkin := &BalanceDruid{
		Druid:           druid,
		primaryRotation: *balanceOptions.Rotation,
	}

	moonkin.useBattleRes = balanceOptions.Options.BattleRes

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

		OnPresimResult: func(presimResult proto.PlayerMetrics, iterations int32) bool {
			if float64(presimResult.NumOom) < float64(iterations)*0.15 {
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
		projectedManaCost := moonkin.manaTracker.ProjectedManaCost(sim, moonkin.GetCharacter())
		remainingManaPool := moonkin.ExpectedRemainingManaPool(sim)

		// If we have enough mana to burn, use the surplus rotation.
		if projectedManaCost < remainingManaPool {
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

	if rotation.FaerieFire {
		ffWait := moonkin.TryFaerieFire(sim, target)
		if ffWait != 0 {
			return ffWait
		}
	}

	if rotation.InsectSwarm && !moonkin.InsectSwarmSpell.DotInput.IsTicking(sim) {
		swarm := moonkin.NewInsectSwarm(sim, target)
		success := swarm.Cast(sim)
		if !success {
			regenTime := moonkin.TimeUntilManaRegen(swarm.GetManaCost())
			return sim.CurrentTime + regenTime
		}
		return sim.CurrentTime + moonkin.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime)
	} else if rotation.Moonfire && !moonkin.MoonfireSpell.DotInput.IsTicking(sim) {
		moonfire := moonkin.NewMoonfire(sim, target)
		success := moonfire.Cast(sim)
		if !success {
			regenTime := moonkin.TimeUntilManaRegen(moonfire.GetManaCost())
			return sim.CurrentTime + regenTime
		}
		return sim.CurrentTime + moonkin.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime)
	}

	var spell *core.SimpleSpell
	switch rotation.PrimarySpell {
	case proto.BalanceDruid_Rotation_Starfire:
		spell = moonkin.NewStarfire(sim, target, 8)
	case proto.BalanceDruid_Rotation_Starfire6:
		spell = moonkin.NewStarfire(sim, target, 6)
	case proto.BalanceDruid_Rotation_Wrath:
		spell = moonkin.NewWrath(sim, target)
	}

	actionSuccessful := spell.Cast(sim)

	if !actionSuccessful {
		regenTime := moonkin.TimeUntilManaRegen(spell.GetManaCost())
		if sim.Log != nil {
			moonkin.Log(sim, "Not enough mana, regenerating for %s.", regenTime)
		}
		return sim.CurrentTime + regenTime
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
//  - SF6
//  - SF6 + IS
// Order of the last two (SF6 / SF6+IS) is swapped if 4p T5 is worn.
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

		currentRotation.InsectSwarm = false
		rotations = append(rotations, currentRotation)
	} else {
		currentRotation.Moonfire = false
		rotations = append(rotations, currentRotation)

		currentRotation.InsectSwarm = true
		rotations = append(rotations, currentRotation)
	}

	return rotations
}
