package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDConsecration int32 = 27173

func (paladin *Paladin) newConsecrationTemplate(sim *core.Simulation) core.SimpleSpellTemplate {

	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID: core.ActionID{
					SpellID: SpellIDConsecration,
				},
				Character:           &paladin.Character,
				CritRollCategory:    core.CritRollCategoryMagical,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				SpellSchool:         core.SpellSchoolHoly,
				SpellExtras:         core.SpellExtrasAlwaysHits,
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

	return core.NewSimpleSpellTemplate(spell)
}

func (paladin *Paladin) NewConsecration(sim *core.Simulation) *core.SimpleSpell {
	paladin.ConsecrationSpell.Cancel(sim)

	consecration := &paladin.ConsecrationSpell
	paladin.consecrationTemplate.Apply(consecration)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	consecration.Init(sim)

	return consecration
}
