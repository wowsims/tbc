package druid

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (druid *Druid) registerInsectSwarmSpell() {
	actionID := core.ActionID{SpellID: 27013}
	baseCost := 175.0

	druid.InsectSwarm = druid.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
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
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			OutcomeApplier:   druid.OutcomeFuncMagicHit(),
			OnSpellHitDealt: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					druid.InsectSwarmDot.Apply(sim)
				}
			},
		}),
	})

	target := druid.CurrentTarget
	druid.InsectSwarmDot = core.NewDot(core.Dot{
		Spell: druid.InsectSwarm,
		Aura: target.RegisterAura(core.Aura{
			Label:    "InsectSwarm-" + strconv.Itoa(int(druid.Index)),
			ActionID: actionID,
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

func (druid *Druid) ShouldCastInsectSwarm(sim *core.Simulation, target *core.Unit, rotation proto.BalanceDruid_Rotation) bool {
	return rotation.InsectSwarm && !druid.InsectSwarmDot.IsActive()
}
