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

func (druid *Druid) newHurricaneTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID: core.ActionID{
					SpellID:    SpellIDHurricane,
					CooldownID: HurricaneCooldownID,
				},
				Character:           &druid.Character,
				SpellSchool:         core.SpellSchoolNature,
				CritRollCategory:    core.CritRollCategoryMagical,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				SpellExtras:         core.SpellExtrasChanneled | core.SpellExtrasAlwaysHits,
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

	baseEffect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
			ThreatMultiplier:       1,
		},
		DotInput: core.DotDamageInput{
			NumberOfTicks:        10,
			TickLength:           time.Second * 1,
			TickBaseDamage:       206,
			TickSpellCoefficient: 0.107,
			DebuffID:             HurricaneDebuffID,
			AffectedByCastSpeed:  true,
		},
	}

	numHits := sim.GetNumTargets()
	effects := make([]core.SpellHitEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)
	}
	spell.Effects = effects

	return core.NewSimpleSpellTemplate(spell)
}

func (druid *Druid) NewHurricane(sim *core.Simulation) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	hurricane := &druid.HurricaneSpell
	druid.hurricaneCastTemplate.Apply(hurricane)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	hurricane.Init(sim)

	return hurricane
}

func (druid *Druid) ShouldCastHurricane(sim *core.Simulation, rotation proto.BalanceDruid_Rotation) bool {
	return rotation.Hurricane && !druid.IsOnCD(HurricaneCooldownID, sim.CurrentTime)
}
