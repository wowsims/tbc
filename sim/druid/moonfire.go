package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Starfire spell IDs
const SpellIDMF int32 = 26988

func (druid *Druid) newMoonfireTemplate(sim *core.Simulation) core.DamageOverTimeSpellTemplate {
	baseCast := core.Cast{
		Name:           "Moonfire",
		CritMultiplier: 1.5,
		SpellSchool:    stats.ArcaneSpellPower,
		Character:      &druid.Character,
		BaseManaCost:   495,
		CastTime:       0,
		ActionID: core.ActionID{
			SpellID: SpellIDMF,
		},
	}

	effect := core.DamageOverTimeSpellEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier: 1,
		},
		DirectInput: core.DirectDamageSpellInput{
			MinBaseDamage:    305,
			MaxBaseDamage:    357,
			SpellCoefficient: 0.15,
		},
		// TODO: actually create dot effects
		DotInput: core.DotDamageInput{
			Name:             "Moonfire DoT",
			Duration:         time.Second * 12,
			TickLength:       time.Second * 3,
			BaseDamage:       600 / 4, // base 600 / 4 ticks
			SpellCoefficient: 0.13,

			// TODO: does druid care about dot ticks?
			// OnDamageTick: func(sim *core.Simulation) {},
		},
	}

	// TODO: Talents
	// effect.SpellEffect.DamageMultiplier *= 1 + 0.01*float64(shaman.Talents.Concussion)
	// effect.SpellEffect.BonusSpellHitRating += float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance
	// effect.SpellEffect.BonusSpellHitRating += float64(shaman.Talents.NaturesGuidance) * 1 * core.SpellHitRatingPerHitChance
	// effect.SpellEffect.BonusSpellCritRating += float64(shaman.Talents.TidalMastery) * 1 * core.SpellCritRatingPerCritChance
	// effect.SpellEffect.BonusSpellCritRating += float64(shaman.Talents.CallOfThunder) * 1 * core.SpellCritRatingPerCritChance

	return core.NewDamageOverTimeSpellTemplate(core.DamageOverTimeSpell{
		SpellCast: core.SpellCast{
			Cast: baseCast,
		},
		Effect: effect,
	})
}

func (druid *Druid) NewMoonfire(sim *core.Simulation, target *core.Target) *core.DamageOverTimeSpell {
	// Initialize cast from precomputed template.
	sf := &druid.moonfireSpell
	druid.moonfireTemplate.Apply(sf)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	sf.Effect.Target = target
	sf.Init(sim)

	return sf
}
