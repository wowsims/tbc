package druid

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDInsectSwarm int32 = 27013

var InsectSwarmActionID = core.ActionID{SpellID: SpellIDInsectSwarm}

func (druid *Druid) registerInsectSwarmSpell(sim *core.Simulation) {
	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    InsectSwarmActionID,
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
	}

	baseEffect := core.SpellEffect{
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		OutcomeApplier:   core.OutcomeFuncMagicHit(),
	}

	target := sim.GetPrimaryTarget()
	debuffAura := target.RegisterAura(&core.Aura{
		Label:    "InsectSwarm-" + strconv.Itoa(int(druid.Index)),
		ActionID: InsectSwarmActionID,
	})

	dotEffect := baseEffect
	dotEffect.IsPeriodic = true
	dotEffect.BaseDamage = core.BaseDamageConfigMagicNoRoll(792/6, 0.127)
	dotEffect.OutcomeApplier = core.OutcomeFuncTick()

	druid.InsectSwarmDot = core.NewDot(core.Dot{
		Aura:          debuffAura,
		NumberOfTicks: 6,
		TickLength:    time.Second * 2,
		TickEffects:   core.TickFuncSnapshot(dotEffect, target),
	})

	druid.InsectSwarm = druid.RegisterSpell(core.SpellConfig{
		Template:     template,
		ModifyCast:   core.ModifyCastAssignTarget,
		ApplyEffects: core.ApplyEffectFuncDot(baseEffect, druid.InsectSwarmDot),
	})
	druid.InsectSwarmDot.Spell = druid.InsectSwarm
}

func (druid *Druid) ShouldCastInsectSwarm(sim *core.Simulation, target *core.Target, rotation proto.BalanceDruid_Rotation) bool {
	return rotation.InsectSwarm && !druid.InsectSwarmDot.IsActive()
}
