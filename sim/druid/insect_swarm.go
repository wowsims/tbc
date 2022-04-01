package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDInsectSwarm int32 = 27013

var InsectSwarmDebuffID = core.NewDebuffID()

func (druid *Druid) registerInsectSwarmSpell(sim *core.Simulation) {
	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    core.ActionID{SpellID: SpellIDInsectSwarm},
				SpellSchool: core.SpellSchoolNature,
				Character:   &druid.Character,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 175,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 175,
				},
				GCD: core.GCDDefault,
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
			DamageMultiplier:    1,
			ThreatMultiplier:    1,
			DotInput: core.DotDamageInput{
				NumberOfTicks:  6,
				TickLength:     time.Second * 2,
				TickBaseDamage: core.DotSnapshotFuncMagic(792/6, 0.127),
				DebuffID:       InsectSwarmDebuffID,
			},
		},
	}

	druid.InsectSwarm = druid.RegisterSpell(core.SpellConfig{
		Template:   template,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}

func (druid *Druid) ShouldCastInsectSwarm(sim *core.Simulation, target *core.Target, rotation proto.BalanceDruid_Rotation) bool {
	return rotation.InsectSwarm && !druid.InsectSwarm.Instance.Effect.DotInput.IsTicking(sim)
}
