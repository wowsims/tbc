package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Starfire spell IDs
const SpellIDWrath int32 = 26985

const IdolAvenger int32 = 31025

func (druid *Druid) newWrathTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	baseCast := core.Cast{
		CritMultiplier: 1.5,
		SpellSchool:    stats.NatureSpellPower,
		Character:      &druid.Character,
		BaseManaCost:   255,
		ManaCost:       255,
		CastTime:       time.Millisecond * 2000,
		ActionID: core.ActionID{
			SpellID: SpellIDWrath,
		},
	}

	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
		},
		DirectInput: core.DirectDamageInput{
			MinBaseDamage:    383,
			MaxBaseDamage:    432,
			SpellCoefficient: 0.571 + 0.02*float64(druid.Talents.WrathOfCenarius),
		},
	}

	if druid.Equip[items.ItemSlotRanged].ID == IdolAvenger {
		// This seems to be unaffected by wrath of cenarius so it needs to come first.
		effect.DirectInput.FlatDamageBonus += 25 * effect.DirectInput.SpellCoefficient
	}

	baseCast.CastTime -= time.Millisecond * 100 * time.Duration(druid.Talents.StarlightWrath)
	effect.BonusSpellCritRating += float64(druid.Talents.FocusedStarlight) * 2 * core.SpellCritRatingPerCritChance // 2% crit per point

	// Convert to percent, multiply by percent increase, convert back to multiplier by adding 1
	baseCast.CritMultiplier = (baseCast.CritMultiplier-1)*(1+float64(druid.Talents.Vengeance)*0.2) + 1
	baseCast.ManaCost -= baseCast.BaseManaCost * float64(druid.Talents.Moonglow) * 0.03
	effect.StaticDamageMultiplier *= 1 + 0.02*float64(druid.Talents.Moonfury)

	effect.OnSpellHit = druid.applyOnHitTalents
	spCast := &core.SpellCast{
		Cast: baseCast,
	}

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: *spCast,
		Effect:    effect,
	})
}

func (druid *Druid) NewWrath(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	sf := &druid.wrathSpell

	druid.wrathCastTemplate.Apply(sf)

	// Modifies the cast time.
	druid.applyNaturesGrace(&sf.SpellCast)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	sf.Effect.Target = target
	sf.Init(sim)

	return sf
}
