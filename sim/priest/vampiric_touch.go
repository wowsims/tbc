package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var VampiricTouchActionID = core.ActionID{SpellID: 34917}

var VampiricTouchDebuffID = core.NewDebuffID()

func (priest *Priest) newVampiricTouchTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	cost := core.ResourceCost{Type: stats.Mana, Value: 425}
	baseCast := core.Cast{
		ActionID:            VampiricTouchActionID,
		Character:           &priest.Character,
		CritRollCategory:    core.CritRollCategoryMagical,
		OutcomeRollCategory: core.OutcomeRollCategoryMagic,
		SpellSchool:         core.SpellSchoolShadow,
		BaseCost:            cost,
		Cost:                cost,
		CastTime:            time.Millisecond * 1500,
		GCD:                 core.GCDDefault,
	}

	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
			ThreatMultiplier:       1,
		},
		DotInput: core.DotDamageInput{
			NumberOfTicks:        5,
			TickLength:           time.Second * 3,
			TickBaseDamage:       650 / 5,
			TickSpellCoefficient: 0.2,
			DebuffID:             VampiricTouchDebuffID,
		},
	}

	priest.applyTalentsToShadowSpell(&baseCast, &effect)

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: baseCast,
		},
		Effect: effect,
	})
}

func (priest *Priest) NewVampiricTouch(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	mf := priest.VTSpellCasting

	priest.vtCastTemplate.Apply(mf)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	mf.Effect.Target = target
	mf.Init(sim)

	return mf
}
