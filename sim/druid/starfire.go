package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Starfire spell IDs
const SpellIDSF8 int32 = 26986
const SpellIDSF6 int32 = 9876

// Idol IDs
const IvoryMoongoddess int32 = 27518

func (druid *Druid) newStarfireTemplate(sim *core.Simulation, rank int) core.SimpleSpellTemplate {
	baseCast := core.Cast{
		Name:           "Starfire (Rank 8)",
		CritMultiplier: 1.5,
		SpellSchool:    stats.ArcaneSpellPower,
		Character:      &druid.Character,
		BaseManaCost:   370,
		ManaCost:       370,
		CastTime:       time.Millisecond * 3500,
		ActionID: core.ActionID{
			SpellID: SpellIDSF8,
		},
	}

	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier: 1,
		},
		DirectInput: core.DirectDamageSpellInput{
			MinBaseDamage:    550,
			MaxBaseDamage:    647,
			SpellCoefficient: 1.0,
		},
	}

	if rank == 6 {
		baseCast.Name = "Starfire (Rank 6)"
		baseCast.BaseManaCost = 315
		baseCast.ManaCost = 315
		baseCast.ActionID = core.ActionID{
			SpellID: SpellIDSF6,
		}
		effect.DirectInput.MinBaseDamage = 463
		effect.DirectInput.MaxBaseDamage = 543
		effect.DirectInput.SpellCoefficient = 0.99
	}

	if druid.Equip[items.ItemSlotRanged].ID == IvoryMoongoddess {
		// This seems to be unaffected by wrath of cenarius so it needs to come first.
		effect.DirectInput.FlatDamageBonus += 55 * effect.DirectInput.SpellCoefficient
	}

	effect.DirectInput.SpellCoefficient += 0.04 * float64(druid.Talents.WrathOfCenarius)

	// TODO: Applies to both starfire and moonfire
	baseCast.CastTime -= time.Millisecond * 100 * time.Duration(druid.Talents.StarlightWrath)
	effect.SpellEffect.BonusSpellCritRating += float64(druid.Talents.FocusedStarlight) * 2 * core.SpellCritRatingPerCritChance // 2% crit per point

	// TODO: applies to starfire, wrath and moonfire

	// Convert to percent, multiply by percent increase, convert back to multiplier by adding 1
	baseCast.CritMultiplier = (baseCast.CritMultiplier-1)*(1+float64(druid.Talents.Vengeance)*0.2) + 1
	baseCast.ManaCost -= baseCast.BaseManaCost * float64(druid.Talents.Moonglow) * 0.03
	effect.SpellEffect.DamageMultiplier *= 1 + 0.02*float64(druid.Talents.Moonfury)

	if ItemSetThunderheart.CharacterHasSetBonus(&druid.Character, 4) { // Thunderheart 4p adds 5% crit to starfire
		effect.SpellEffect.BonusSpellCritRating += 5 * core.SpellCritRatingPerCritChance
	}

	effect.OnSpellHit = druid.applyOnHitTalents
	spCast := &core.SpellCast{
		Cast: baseCast,
	}

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast:      *spCast,
		SpellHitEffect: effect,
	})
}

func (druid *Druid) NewStarfire(sim *core.Simulation, target *core.Target, rank int) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	sf := &druid.starfireSpell

	if rank == 8 {
		druid.starfire8CastTemplate.Apply(sf)
	} else if rank == 6 {
		druid.starfire6CastTemplate.Apply(sf)
	}

	// Applies nature's grace cast time reduction if available.
	druid.applyNaturesGrace(&sf.SpellCast)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	sf.SpellHitEffect.Target = target
	sf.Init(sim)

	return sf
}
