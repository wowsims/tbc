package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDDevouringPlague int32 = 25467

var DevouringPlagueCooldownID = core.NewCooldownID()

func (priest *Priest) newDevouringPlagueTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	baseCast := core.Cast{
		ActionID: core.ActionID{
			SpellID:    SpellIDDevouringPlague,
			CooldownID: DevouringPlagueCooldownID,
		},
		Character:    &priest.Character,
		SpellSchool:  stats.ShadowSpellPower,
		BaseManaCost: 1145,
		ManaCost:     1145,
		CastTime:     0,
		GCD:          core.GCDDefault,
		Cooldown:     time.Minute * 3,
	}

	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
			ThreatMultiplier:       1,
		},
		DotInput: core.DotDamageInput{
			NumberOfTicks:        8,
			TickLength:           time.Second * 3,
			TickBaseDamage:       1216 / 8,
			TickSpellCoefficient: 0.1,
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

func (priest *Priest) NewDevouringPlague(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	mf := &priest.DevouringPlagueSpell

	priest.devouringPlagueTemplate.Apply(mf)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	mf.Effect.Target = target
	mf.Init(sim)

	return mf
}
