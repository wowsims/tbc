package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Starfire spell IDs
const SpellIDSF8 int32 = 26986
const SpellIDSF6 int32 = 9876

func (druid *Druid) newStarfireTemplate(sim *core.Simulation, rank int) core.SingleTargetDirectDamageSpellTemplate {
	baseCast := core.Cast{
		Name:           "Starfire",
		CritMultiplier: 1.5,
		SpellSchool:    stats.ArcaneSpellPower,
		Character:      &druid.Character,

		BaseManaCost: 370,
		CastTime:     time.Millisecond * 3500,
		ActionID: core.ActionID{
			SpellID: SpellIDSF8,
		},
	}

	effect := core.DirectDamageSpellEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier: 1,
		},
		DirectDamageSpellInput: core.DirectDamageSpellInput{
			MinBaseDamage:    550,
			MaxBaseDamage:    647,
			SpellCoefficient: 1.0,
		},
	}

	if rank == 6 {
		baseCast.ActionID = core.ActionID{
			SpellID: SpellIDSF6,
		}
		effect.DirectDamageSpellInput.MinBaseDamage = 463
		effect.DirectDamageSpellInput.MaxBaseDamage = 543
		effect.SpellCoefficient = 0.99
	}

	// TODO: Talents
	// effect.SpellEffect.DamageMultiplier *= 1 + 0.01*float64(shaman.Talents.Concussion)
	// effect.SpellEffect.BonusSpellHitRating += float64(shaman.Talents.ElementalPrecision) * 2 * core.SpellHitRatingPerHitChance
	// effect.SpellEffect.BonusSpellHitRating += float64(shaman.Talents.NaturesGuidance) * 1 * core.SpellHitRatingPerHitChance
	// effect.SpellEffect.BonusSpellCritRating += float64(shaman.Talents.TidalMastery) * 1 * core.SpellCritRatingPerCritChance
	// effect.SpellEffect.BonusSpellCritRating += float64(shaman.Talents.CallOfThunder) * 1 * core.SpellCritRatingPerCritChance

	return core.NewSingleTargetDirectDamageSpellTemplate(core.SingleTargetDirectDamageSpell{
		SpellCast: core.SpellCast{
			Cast: baseCast,
		},
		Effect: effect,
	})
}

func (druid *Druid) NewStarfire(sim *core.Simulation, target *core.Target, rank int) *core.SingleTargetDirectDamageSpell {
	// Initialize cast from precomputed template.
	sf := &druid.starfireSpell

	if rank == 8 {
		druid.starfire8CastTemplate.Apply(sf)
	} else if rank == 6 {
		druid.starfire6CastTemplate.Apply(sf)
	}

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	sf.Effect.Target = target
	sf.Init(sim)

	return sf
}
