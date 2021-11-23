package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDMF int32 = 25387

const TagMF2 = 2
const TagMF3 = 3

func (priest *Priest) newMindflayTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	baseCast := core.Cast{
		Name:           "Mind Flay",
		CritMultiplier: 1.5,
		SpellSchool:    stats.ShadowSpellPower,
		Character:      &priest.Character,
		BaseManaCost:   230,
		ManaCost:       230,
		CastTime:       0,
		Binary:         true,
		ActionID: core.ActionID{
			SpellID: SpellIDMF,
		},
	}

	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier: 1,
		},
		DotInput: core.DotDamageInput{
			NumberOfTicks:        3,
			TickLength:           time.Second,
			TickBaseDamage:       528 / 3,
			TickSpellCoefficient: 0.19,
		},
	}

	priest.applyTalentsToShadowSpell(&baseCast, &effect)

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: baseCast,
		},
		SpellHitEffect: effect,
	})
}

func (priest *Priest) NewMindFlay(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	mf := &priest.MindFlaySpell

	priest.mindflayCastTemplate.Apply(mf)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	mf.Target = target
	mf.Init(sim)

	// Channels accelerate tick speed by haste.
	mf.DotInput.TickLength = time.Duration(float64(mf.DotInput.TickLength) / priest.CastSpeed())

	return mf
}
