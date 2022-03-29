package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDInsectSwarm int32 = 27013

var InsectSwarmDebuffID = core.NewDebuffID()

func (druid *Druid) newInsectSwarmTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	baseCast := core.Cast{
		ActionID:            core.ActionID{SpellID: SpellIDInsectSwarm},
		CritRollCategory:    core.CritRollCategoryMagical,
		OutcomeRollCategory: core.OutcomeRollCategoryMagic,
		SpellSchool:         core.SpellSchoolNature,
		Character:           &druid.Character,
		BaseCost: core.ResourceCost{
			Type:  stats.Mana,
			Value: 175,
		},
		Cost: core.ResourceCost{
			Type:  stats.Mana,
			Value: 175,
		},
		GCD: core.GCDDefault,
	}

	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
		},
		DotInput: core.DotDamageInput{
			NumberOfTicks:  6,
			TickLength:     time.Second * 2,
			TickBaseDamage: core.DotSnapshotFuncMagic(792/6, 0.127),
			DebuffID:       InsectSwarmDebuffID,
		},
	}
	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: baseCast,
		},
		Effect: effect,
	})
}

func (druid *Druid) NewInsectSwarm(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	sf := &druid.InsectSwarmSpell
	druid.insectSwarmCastTemplate.Apply(sf)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	sf.Effect.Target = target
	sf.Init(sim)

	return sf
}

func (druid *Druid) ShouldCastInsectSwarm(sim *core.Simulation, target *core.Target, rotation proto.BalanceDruid_Rotation) bool {
	return rotation.InsectSwarm && !druid.InsectSwarmSpell.Effect.DotInput.IsTicking(sim)
}
