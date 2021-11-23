package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDSWD int32 = 32996

var SWDCooldownID = core.NewCooldownID()

func (priest *Priest) newSWDTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	baseCast := core.Cast{
		Name:           "Shadow Word: Death",
		CritMultiplier: 1.5,
		SpellSchool:    stats.ShadowSpellPower,
		Character:      &priest.Character,
		BaseManaCost:   309,
		ManaCost:       309,
		CastTime:       0,
		Cooldown:       time.Second * 12,
		ActionID: core.ActionID{
			SpellID:    SpellIDSWD,
			CooldownID: SWDCooldownID,
		},
	}

	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
		},
		DirectInput: core.DirectDamageSpellInput{
			MinBaseDamage:    672,
			MaxBaseDamage:    664,
			SpellCoefficient: 0.429,
		},
	}

	// TODO: Manage health with SW:D or assume healers got you?

	priest.applyTalentsToShadowSpell(&baseCast, &effect)

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: baseCast,
		},
		SpellHitEffect: effect,
	})
}

func (priest *Priest) NewSWD(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	mf := &priest.swdSpell

	priest.swdCastTemplate.Apply(mf)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	mf.Target = target
	mf.Init(sim)

	return mf
}
