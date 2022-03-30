package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var SerpentStingDebuffID = core.NewDebuffID()
var SerpentStingActionID = core.ActionID{SpellID: 27016}

func (hunter *Hunter) newSerpentStingTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	cost := core.ResourceCost{Type: stats.Mana, Value: 275}
	ama := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    SerpentStingActionID,
				Character:   &hunter.Character,
				SpellSchool: core.SpellSchoolNature,
				GCD:         core.GCDDefault,
				Cost:        cost,
				BaseCost:    cost,
				IgnoreHaste: true, // Hunter GCD is locked at 1.5s
			},
			OutcomeRollCategory: core.OutcomeRollCategoryRanged,
			CritRollCategory:    core.CritRollCategoryNone,
		},
		Effect: core.SpellEffect{
			ProcMask:         core.ProcMaskRangedSpecial,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			DotInput: core.DotDamageInput{
				NumberOfTicks: 5,
				TickLength:    time.Second * 3,
				TickBaseDamage: func(sim *core.Simulation, hitEffect *core.SpellEffect, spellCast *core.SpellCast) float64 {
					attackPower := hitEffect.RangedAttackPower(spellCast) + hitEffect.RangedAttackPowerOnTarget()
					return 132 + attackPower*0.02
				},
				DebuffID: SerpentStingDebuffID,
			},
		},
	}

	ama.Cost.Value *= 1 - 0.02*float64(hunter.Talents.Efficiency)
	ama.Effect.DamageMultiplier *= 1 + 0.06*float64(hunter.Talents.ImprovedStings)

	return core.NewSimpleSpellTemplate(ama)
}

func (hunter *Hunter) NewSerpentSting(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	ss := &hunter.serpentSting
	hunter.serpentStingTemplate.Apply(ss)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	ss.Effect.Target = target

	ss.Init(sim)
	return ss
}
