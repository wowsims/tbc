package shadow

import (
	"time"

	"github.com/wowsims/tbc/sim/common"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/priest"
)

func RegisterShadowPriest() {
	core.RegisterAgentFactory(proto.PlayerOptions_ShadowPriest{}, func(character core.Character, options proto.PlayerOptions, isr proto.IndividualSimRequest) core.Agent {
		return NewShadowPriest(character, options, isr)
	})
}

func NewShadowPriest(character core.Character, options proto.PlayerOptions, isr proto.IndividualSimRequest) *ShadowPriest {
	shadowOptions := options.GetShadowPriest()

	selfBuffs := priest.SelfBuffs{}

	priest := priest.NewPriest(character, selfBuffs, *shadowOptions.Talents)
	spriest := &ShadowPriest{
		Priest: priest,
	}
	// spriest.PickRotations(*shadowOptions.Rotation, isr)

	return spriest
}

type ShadowPriest struct {
	priest.Priest

	primaryRotation proto.ShadowPriest_Rotation
	// These are only used when primary spell is set to 'Adaptive'. When the mana
	// tracker tells us we have extra mana to spare, use surplusRotation instead of
	// primaryRotation.
	useSurplusRotation bool
	surplusRotation    proto.ShadowPriest_Rotation
	manaTracker        common.ManaSpendingRateTracker
}

func (spriest *ShadowPriest) GetPriest() *priest.Priest {
	return &spriest.Priest
}

func (spriest *ShadowPriest) Reset(sim *core.Simulation) {
	if spriest.useSurplusRotation {
		spriest.manaTracker.Reset()
	}
	spriest.Priest.Reset(sim)
}

func (spriest *ShadowPriest) Act(sim *core.Simulation) time.Duration {
	if spriest.useSurplusRotation {
		spriest.manaTracker.Update(sim, spriest.GetCharacter())
		projectedManaCost := spriest.manaTracker.ProjectedManaCost(sim, spriest.GetCharacter())

		// If we have enough mana to burn, use the surplus rotation.
		if projectedManaCost < spriest.CurrentMana() {
			return spriest.actRotation(sim, spriest.surplusRotation)
		} else {
			return spriest.actRotation(sim, spriest.primaryRotation)
		}
	} else {
		return spriest.actRotation(sim, spriest.primaryRotation)
	}
}

func (spriest *ShadowPriest) actRotation(sim *core.Simulation, rotation proto.ShadowPriest_Rotation) time.Duration {
	// Activate shared druid behaviors

	// target := sim.GetPrimaryTarget()

	// if rotation.FaerieFire {
	// 	ffWait := moonkin.TryFaerieFire(sim, target)
	// 	if ffWait != 0 {
	// 		return ffWait
	// 	}
	// }

	var spell *core.SimpleSpell
	// switch rotation.PrimarySpell {
	// case proto.BalanceDruid_Rotation_Starfire:
	// 	spell = moonkin.NewStarfire(sim, target, 8)
	// case proto.BalanceDruid_Rotation_Starfire6:
	// 	spell = moonkin.NewStarfire(sim, target, 6)
	// case proto.BalanceDruid_Rotation_Wrath:
	// 	spell = moonkin.NewWrath(sim, target)
	// }

	actionSuccessful := spell.Cast(sim)

	if !actionSuccessful {
		regenTime := spriest.TimeUntilManaRegen(spell.GetManaCost())
		if sim.Log != nil {
			sim.Log("Not enough mana, regenerating for %s.\n", regenTime)
		}
		return sim.CurrentTime + regenTime
	}

	return sim.CurrentTime + core.MaxDuration(
		spriest.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime),
		spell.CastTime)
}
