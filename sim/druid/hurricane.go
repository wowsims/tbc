package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDHurricane int32 = 27012

var HurricaneCooldownID = core.NewCooldownID()
var HurricaneDebuffID = core.NewDebuffID()

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
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
		},
	}

	baseEffect := core.SpellEffect{
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		DotInput: core.DotDamageInput{
			NumberOfTicks:       10,
			TickLength:          time.Second * 1,
			TickBaseDamage:      core.DotSnapshotFuncMagic(206, 0.107),
			DebuffID:            HurricaneDebuffID,
			AffectedByCastSpeed: true,
		},
	}

	numHits := sim.GetNumTargets()
	effects := make([]core.SpellEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)
	}
	template.Effects = effects

	druid.Hurricane = druid.RegisterSpell(core.SpellConfig{
		Template: template,
	})
}

func (druid *Druid) ShouldCastHurricane(sim *core.Simulation, rotation proto.BalanceDruid_Rotation) bool {
	return rotation.Hurricane && !druid.IsOnCD(HurricaneCooldownID, sim.CurrentTime)
}
