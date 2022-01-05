package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDInsectSwarm int32 = 27013

var InsectSwarmDebuffID = core.NewDebuffID()

func (druid *Druid) newInsectSwarmTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	baseCast := core.Cast{
		Name:           "Insect Swarm",
		CritMultiplier: 1.5,
		SpellSchool:    stats.NatureSpellPower,
		Character:      &druid.Character,
		BaseManaCost:   175,
		ManaCost:       175,
		CastTime:       0,
		ActionID: core.ActionID{
			SpellID: SpellIDInsectSwarm,
		},
	}

	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
		},
		DotInput: core.DotDamageInput{
			NumberOfTicks:        6,
			TickLength:           time.Second * 2,
			TickBaseDamage:       792 / 6,
			TickSpellCoefficient: 0.127,
			DebuffID:             InsectSwarmDebuffID,
			SpellID:              SpellIDInsectSwarm,
		},
	}
	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: baseCast,
		},
		SpellHitEffect: effect,
	})
}

// TODO: This might behave weird if we have a moonfire still ticking when we cast one.
//   We could do a check and if the spell is still ticking allocate a new SingleHitSpell?
func (druid *Druid) NewInsectSwarm(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	sf := &druid.InsectSwarmSpell
	druid.insectSwarmCastTemplate.Apply(sf)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	sf.Target = target
	sf.Init(sim)

	return sf
}

func (druid *Druid) ShouldCastInsectSwarm(sim *core.Simulation, target *core.Target, rotation proto.BalanceDruid_Rotation) bool {
	return rotation.InsectSwarm && !druid.InsectSwarmSpell.DotInput.IsTicking(sim)
}
