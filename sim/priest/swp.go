package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Starfire spell IDs
const SpellIDSWP int32 = 25368

func (priest *Priest) newSWPTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	baseCast := core.Cast{
		Name:           "Shadow Word: Pain",
		CritMultiplier: 1.5,
		SpellSchool:    stats.ShadowSpellPower,
		Character:      &priest.Character,
		BaseManaCost:   575,
		ManaCost:       575,
		CastTime:       0,
		ActionID: core.ActionID{
			SpellID: SpellIDSWP,
		},
	}

	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier: 1,
		},
		DotInput: core.DotDamageInput{
			NumberOfTicks:        6,
			TickLength:           time.Second * 3,
			TickBaseDamage:       1236 / 6,
			TickSpellCoefficient: 0.183,

			// TODO:  Shadow Weaving apply on tick?
			OnDamageTick: newVTOnTick(priest.Party),
		},
	}

	effect.DotInput.NumberOfTicks += int(priest.Talents.ImprovedShadowWordPain) // extra tick per point
	priest.applyTalentsToShadowSpell(&baseCast, &effect)

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: baseCast,
		},
		SpellHitEffect: effect,
	})
}

func (priest *Priest) NewSWP(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	mf := &priest.SWPSpell

	priest.swpCastTemplate.Apply(mf)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	mf.Target = target
	mf.Init(sim)

	return mf
}
