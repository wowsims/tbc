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
	baseCost := 175.0

	druid.InsectSwarm = druid.RegisterSpell(core.SpellConfig{
		ActionID:    InsectSwarmActionID,
		SpellSchool: core.SpellSchoolNature,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost,
				GCD:  core.GCDDefault,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			OutcomeApplier:   druid.OutcomeFuncMagicHit(),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					druid.InsectSwarmDot.Apply(sim)
				}
			},
		}),
	})

	target := sim.GetPrimaryTarget()
	druid.InsectSwarmDot = core.NewDot(core.Dot{
		Spell: druid.InsectSwarm,
		Aura: target.RegisterAura(core.Aura{
			Label:    "InsectSwarm-" + strconv.Itoa(int(druid.Index)),
			ActionID: InsectSwarmActionID,
		}),
		NumberOfTicks: 6,
		TickLength:    time.Second * 2,
		TickEffects: core.TickFuncSnapshot(target, core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			IsPeriodic:       true,
			BaseDamage:       core.BaseDamageConfigMagicNoRoll(792/6, 0.127),
			OutcomeApplier:   druid.OutcomeFuncTick(),
		}),
	})
}

func (druid *Druid) ShouldCastInsectSwarm(sim *core.Simulation, target *core.Target, rotation proto.BalanceDruid_Rotation) bool {
	return rotation.InsectSwarm && !druid.InsectSwarmDot.IsActive()
}
