package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var ScorpidStingDebuffID = core.NewDebuffID()

func (hunter *Hunter) newScorpidStingTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	actionID := core.ActionID{SpellID: 3043}
	cost := core.ResourceCost{Type: stats.Mana, Value: hunter.BaseMana() * 0.09}
	ama := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            actionID,
				Character:           &hunter.Character,
				OutcomeRollCategory: core.OutcomeRollCategoryRanged,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolNature,
				GCD:                 core.GCDDefault,
				Cost:                cost,
				BaseCost:            cost,
				IgnoreHaste:         true, // Hunter GCD is locked at 1.5s
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ProcMask: core.ProcMaskRangedSpecial,
				OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
					if !spellEffect.Landed() {
						return
					}

					spellEffect.Target.AddAura(sim, core.Aura{
						ID:       ScorpidStingDebuffID,
						ActionID: actionID,
						Expires:  sim.CurrentTime + time.Second*20,
					})
				},
			},
		},
	}

	ama.Cost.Value *= 1 - 0.02*float64(hunter.Talents.Efficiency)

	return core.NewSimpleSpellTemplate(ama)
}

func (hunter *Hunter) NewScorpidSting(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	as := &hunter.scorpidSting
	hunter.scorpidStingTemplate.Apply(as)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	as.Effect.Target = target

	as.Init(sim)

	return as
}
