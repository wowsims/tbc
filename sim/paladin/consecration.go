package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDConsecration int32 = 27173

func (paladin *Paladin) registerConsecrationSpell(sim *core.Simulation) {

	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID: core.ActionID{
					SpellID: SpellIDConsecration,
				},
				Character:   &paladin.Character,
				SpellSchool: core.SpellSchoolHoly,
				SpellExtras: core.SpellExtrasAlwaysHits,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 660,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 660,
				},
				GCD: core.GCDDefault,
			},
			CritRollCategory:    core.CritRollCategoryMagical,
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
		},
	}

	effect := core.SpellEffect{
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		DotInput: core.DotDamageInput{
			NumberOfTicks:  8,
			TickLength:     time.Second,
			TickBaseDamage: core.DotSnapshotFuncMagic(64, 0.119),
		},
	}

	// TODO: consecration talents here

	numHits := sim.GetNumTargets()
	effects := make([]core.SpellEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, effect)
		effects[i].Target = sim.GetTarget(i)
	}
	spell.Effects = effects

	paladin.Consecration = paladin.RegisterSpell(core.SpellConfig{
		Template: spell,
	})
}
