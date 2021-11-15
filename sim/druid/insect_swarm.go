package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDIS int32 = 26988

func (druid *Druid) newInsectSwarmTemplate(sim *core.Simulation) core.DamageOverTimeSpellTemplate {
	baseCast := core.Cast{
		Name:           "Insect Swarm",
		CritMultiplier: 1.5,
		SpellSchool:    stats.NatureSpellPower,
		Character:      &druid.Character,
		BaseManaCost:   175,
		ManaCost:       175,
		CastTime:       0,
		ActionID: core.ActionID{
			SpellID: SpellIDIS,
		},
	}

	effect := core.DamageOverTimeSpellEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier: 1,
		},
		DotInput: core.DotDamageInput{
			Name:             "Insect Swarm DoT",
			NumberTicks:      6,
			TickLength:       time.Second * 2,
			BaseDamage:       792,
			SpellCoefficient: 0.127,
		},
	}
	return core.NewDamageOverTimeSpellTemplate(core.DamageOverTimeSpell{
		SpellCast: core.SpellCast{
			Cast: baseCast,
		},
		DamageOverTimeSpellEffect: effect,
	})
}

// TODO: This might behave weird if we have a moonfire still ticking when we cast one.
//   We could do a check and if the spell is still ticking allocate a new DamageOverTimeSpell?
func (druid *Druid) NewInsectSwarm(sim *core.Simulation, target *core.Target) *core.DamageOverTimeSpell {
	// Initialize cast from precomputed template.
	sf := &druid.InsectSwarmSpell
	druid.insectSwarmCastTemplate.Apply(sf)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	sf.Target = target
	sf.Init(sim)

	return sf
}
