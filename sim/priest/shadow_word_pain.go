package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDShadowWordPain int32 = 25368

var ShadowWordPainActionID = core.ActionID{SpellID: SpellIDShadowWordPain}

var ShadowWordPainDebuffID = core.NewDebuffID()

func (priest *Priest) newShadowWordPainTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	cost := core.ResourceCost{Type: stats.Mana, Value: 575}
	baseCast := core.Cast{
		ActionID:            ShadowWordPainActionID,
		Character:           &priest.Character,
		OutcomeRollCategory: core.OutcomeRollCategoryMagic,
		SpellSchool:         core.SpellSchoolShadow,
		BaseCost:            cost,
		Cost:                cost,
		CastTime:            0,
		GCD:                 core.GCDDefault,
	}

	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
		},
		DotInput: core.DotDamageInput{
			NumberOfTicks:  6,
			TickLength:     time.Second * 3,
			TickBaseDamage: core.DotSnapshotFuncMagic(1236/6, 0.183),
			DebuffID:       ShadowWordPainDebuffID,
		},
	}

	effect.DotInput.NumberOfTicks += int(priest.Talents.ImprovedShadowWordPain) // extra tick per point

	if ItemSetAbsolution.CharacterHasSetBonus(&priest.Character, 2) { // Absolution 2p adds 1 extra tick to swp
		effect.DotInput.NumberOfTicks += 1
	}

	priest.applyTalentsToShadowSpell(&baseCast, &effect)

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: baseCast,
		},
		Effect: effect,
	})
}

func (priest *Priest) NewShadowWordPain(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	mf := &priest.SWPSpell

	priest.swpCastTemplate.Apply(mf)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	mf.Effect.Target = target
	mf.Init(sim)

	return mf
}
