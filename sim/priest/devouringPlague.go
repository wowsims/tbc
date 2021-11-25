package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDDevouringPlague int32 = 25467

var DPCooldownID = core.NewCooldownID()

func (priest *Priest) newDevouringPlagueTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	baseCast := core.Cast{
		Name:           "Devouring Plague",
		CritMultiplier: 1.5,
		SpellSchool:    stats.ShadowSpellPower,
		Character:      &priest.Character,
		BaseManaCost:   1145,
		ManaCost:       1145,
		CastTime:       0,
		Cooldown:       time.Minute * 3,
		ActionID: core.ActionID{
			SpellID:    SpellIDDevouringPlague,
			CooldownID: DPCooldownID,
		},
	}

	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
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
		SpellHitEffect: effect,
	})
}

func (priest *Priest) NewDevouringPlague(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	mf := &priest.DevouringPlagueSpell

	priest.devouringPlagueTemplate.Apply(mf)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	mf.Target = target
	mf.Init(sim)

	return mf
}
