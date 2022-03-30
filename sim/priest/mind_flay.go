package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDMindFlay int32 = 25387

const TagMF2 = 2
const TagMF3 = 3

func (priest *Priest) newMindflayTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	baseCast := core.Cast{
		ActionID: core.ActionID{
			SpellID: SpellIDMindFlay,
			Tag:     3, // default to 3 tick mf
		},
		Character:   &priest.Character,
		SpellSchool: core.SpellSchoolShadow,
		SpellExtras: core.SpellExtrasBinary | core.SpellExtrasChanneled,
		BaseCost: core.ResourceCost{
			Type:  stats.Mana,
			Value: 230,
		},
		Cost: core.ResourceCost{
			Type:  stats.Mana,
			Value: 230,
		},
		CastTime: 0,
		GCD:      core.GCDDefault,
	}

	effect := core.SpellEffect{
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		DotInput: core.DotDamageInput{
			NumberOfTicks:       3,
			TickLength:          time.Second,
			TickBaseDamage:      core.DotSnapshotFuncMagic(528/3, 0.19),
			AffectedByCastSpeed: true,
		},
	}

	priest.applyTalentsToShadowSpell(&baseCast, &effect)

	if ItemSetIncarnate.CharacterHasSetBonus(&priest.Character, 4) {
		effect.DamageMultiplier *= 1.05
	}

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast:                baseCast,
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
		},
		Effect: effect,
	})
}

func (priest *Priest) NewMindFlay(sim *core.Simulation, target *core.Target, numTicks int) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	mf := &priest.MindFlaySpell
	priest.mindflayCastTemplate.Apply(mf)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	mf.ActionID.Tag = int32(numTicks)
	mf.Effect.DotInput.NumberOfTicks = numTicks
	mf.Effect.Target = target

	mf.Init(sim)

	return mf
}
