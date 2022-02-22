package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

const SpellIDShadowWordDeath int32 = 32996

var SWDCooldownID = core.NewCooldownID()

func (priest *Priest) newShadowWordDeathTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	baseCast := core.Cast{
		ActionID: core.ActionID{
			SpellID:    SpellIDShadowWordDeath,
			CooldownID: SWDCooldownID,
		},
		Character:           &priest.Character,
		CritRollCategory:    core.CritRollCategoryMagical,
		OutcomeRollCategory: core.OutcomeRollCategoryMagic,
		SpellSchool:         core.SpellSchoolShadow,
		BaseManaCost:        309,
		ManaCost:            309,
		CastTime:            0,
		GCD:                 core.GCDDefault,
		Cooldown:            time.Second * 12,
		CritMultiplier:      priest.DefaultSpellCritMultiplier(),
	}

	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
			ThreatMultiplier:       1,
		},
		DirectInput: core.DirectDamageInput{
			MinBaseDamage:    572,
			MaxBaseDamage:    664,
			SpellCoefficient: 0.429,
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

func (priest *Priest) NewShadowWordDeath(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	mf := &priest.swdSpell

	priest.swdCastTemplate.Apply(mf)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	mf.Effect.Target = target
	mf.Init(sim)

	return mf
}
