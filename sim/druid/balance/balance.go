package balance

import (
	"time"

	"github.com/wowsims/tbc/sim/common"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/druid"
	googleProto "google.golang.org/protobuf/proto"
)

func RegisterBalanceDruid() {
	core.RegisterAgentFactory(proto.PlayerOptions_BalanceDruid{}, func(character core.Character, options proto.PlayerOptions, isr proto.IndividualSimRequest) core.Agent {
		return NewBalanceDruid(character, options, isr)
	})
}

func NewBalanceDruid(character core.Character, options proto.PlayerOptions, isr proto.IndividualSimRequest) *BalanceDruid {
	balanceOptions := options.GetBalanceDruid()

	selfBuffs := druid.SelfBuffs{}
	if balanceOptions.Options.InnervateTarget != nil {
		// if targetting myself for individual sim
		// TODO: what is my player idx for raid?
		selfBuffs.Innervate = balanceOptions.Options.InnervateTarget.TargetIndex == 0
	}

	druid := druid.NewDruid(character, selfBuffs, *balanceOptions.Talents)
	moonkin := &BalanceDruid{
		Druid: druid,
	}
	moonkin.PickRotations(*balanceOptions.Rotation, isr)

	return moonkin
}

type BalanceDruid struct {
	druid.Druid

	primaryRotation proto.BalanceDruid_Rotation

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

		// If we have enough mana to burn, use the surplus rotation.
		if projectedManaCost < moonkin.CurrentMana() {
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
			sim.Log("Not enough mana, regenerating for %s.\n", regenTime)
		}
		return sim.CurrentTime + regenTime
	}

	return sim.CurrentTime + core.MaxDuration(
		moonkin.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime),
		spell.CastTime)
}

// Sets the rotation fields on moonkin. If adaptive, uses multiple presims.
func (moonkin *BalanceDruid) PickRotations(baseRotation proto.BalanceDruid_Rotation, isr proto.IndividualSimRequest) {
	// If not adaptive, just use the base rotation directly.
	if baseRotation.PrimarySpell != proto.BalanceDruid_Rotation_Adaptive {
		moonkin.primaryRotation = baseRotation
		return
	}

	// If no encounter is set, it means we aren't going to run a sim at all.
	// So just return something valid.
	// TODO: Probably need some organized way of doing presims so we dont have
	// to check these types of things.
	if isr.Encounter == nil || len(isr.Encounter.Targets) == 0 {
		moonkin.primaryRotation = baseRotation
		return
	}

	rotations := moonkin.GetDpsRotationHierarchy(baseRotation)
	for i, rotation := range rotations {
		presimRequest := googleProto.Clone(&isr).(*proto.IndividualSimRequest)
		presimRequest.SimOptions.RandomSeed = 1
		presimRequest.SimOptions.Debug = false
		presimRequest.SimOptions.Iterations = 100
		*presimRequest.Player.Options.Spec.(*proto.PlayerOptions_BalanceDruid).BalanceDruid.Rotation = rotation

		presimResult := core.RunIndividualSim(presimRequest)

		if presimResult.PlayerMetrics.NumOom < 15 {
			moonkin.primaryRotation = rotation

			// If the highest dps rotation is fine, we dont need any adaptive logic.
			if i == 0 {
				return
			}

			moonkin.useSurplusRotation = true
			moonkin.surplusRotation = rotations[i-1]
			moonkin.manaTracker = common.NewManaSpendingRateTracker()
			return
		}
	}

	// If we are here than all of the rotations went oom. No adaptive logic needed, just use the lowest one.
	moonkin.primaryRotation = rotations[len(rotations)-1]
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
