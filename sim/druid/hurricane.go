package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDHurricane int32 = 27012

var HurricaneCooldownID = core.NewCooldownID()

func (druid *Druid) registerHurricaneSpell(sim *core.Simulation) {
	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID: core.ActionID{
					SpellID:    SpellIDHurricane,
					CooldownID: HurricaneCooldownID,
				},
				Character:   &druid.Character,
				SpellSchool: core.SpellSchoolNature,
				SpellExtras: core.SpellExtrasChanneled | core.SpellExtrasAlwaysHits,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 1905,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 1905,
				},
				GCD:      core.GCDDefault,
				Cooldown: time.Second * 60,
			},
		},
	}

	hurricaneDot := core.NewDot(core.Dot{
		Aura: druid.RegisterAura(&core.Aura{
			Label: "Hurricane",
		}),
		NumberOfTicks:       10,
		TickLength:          time.Second * 1,
		AffectedByCastSpeed: true,
		TickEffects: core.TickFuncAOESnapshot(sim, core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BaseDamage:       core.BaseDamageConfigMagicNoRoll(206, 0.107),
			OutcomeApplier:   core.OutcomeFuncTick(),
			IsPeriodic:       true,
		}),
	})

	druid.Hurricane = druid.RegisterSpell(core.SpellConfig{
		Template:     template,
		ApplyEffects: core.ApplyEffectFuncDot(hurricaneDot),
	})
	hurricaneDot.Spell = druid.Hurricane
}

func (druid *Druid) ShouldCastHurricane(sim *core.Simulation, rotation proto.BalanceDruid_Rotation) bool {
	return rotation.Hurricane && !druid.IsOnCD(HurricaneCooldownID, sim.CurrentTime)
}
