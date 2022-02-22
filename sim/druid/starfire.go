package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
)

// Starfire spell IDs
const SpellIDSF8 int32 = 26986
const SpellIDSF6 int32 = 9876

// Idol IDs
const IvoryMoongoddess int32 = 27518

func (druid *Druid) newStarfireTemplate(sim *core.Simulation, rank int) core.SimpleSpellTemplate {
	baseCast := core.Cast{
		ActionID:            core.ActionID{SpellID: SpellIDSF8},
		Character:           &druid.Character,
		CritRollCategory:    core.CritRollCategoryMagical,
		OutcomeRollCategory: core.OutcomeRollCategoryMagic,
		SpellSchool:         core.SpellSchoolArcane,
		BaseManaCost:        370,
		ManaCost:            370,
		CastTime:            time.Millisecond * 3500,
		GCD:                 core.GCDDefault,
		CritMultiplier:      druid.SpellCritMultiplier(1, 0.2*float64(druid.Talents.Vengeance)),
	}

	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
			ThreatMultiplier:       1,
		},
		DirectInput: core.DirectDamageInput{
			MinBaseDamage:    550,
			MaxBaseDamage:    647,
			SpellCoefficient: 1.0,
		},
	}

	if rank == 6 {
		baseCast.BaseManaCost = 315
		baseCast.ManaCost = 315
		baseCast.ActionID = core.ActionID{
			SpellID: SpellIDSF6,
		}
		effect.DirectInput.MinBaseDamage = 463
		effect.DirectInput.MaxBaseDamage = 543
		effect.DirectInput.SpellCoefficient = 0.99
	}

	baseCast.ManaCost -= baseCast.BaseManaCost * 0.03 * float64(druid.Talents.Moonglow)
	baseCast.CastTime -= time.Millisecond * 100 * time.Duration(druid.Talents.StarlightWrath)

	effect.StaticDamageMultiplier *= 1 + 0.02*float64(druid.Talents.Moonfury)
	effect.BonusSpellCritRating += float64(druid.Talents.FocusedStarlight) * 2 * core.SpellCritRatingPerCritChance // 2% crit per point
	effect.DirectInput.SpellCoefficient += 0.04 * float64(druid.Talents.WrathOfCenarius)

	if druid.Equip[items.ItemSlotRanged].ID == IvoryMoongoddess {
		// This seems to be unaffected by wrath of cenarius so it needs to come first.
		effect.DirectInput.FlatDamageBonus += 55 * effect.DirectInput.SpellCoefficient
	}
	if ItemSetThunderheart.CharacterHasSetBonus(&druid.Character, 4) { // Thunderheart 4p adds 5% crit to starfire
		effect.BonusSpellCritRating += 5 * core.SpellCritRatingPerCritChance
	}

	effect.OnSpellHit = druid.applyOnHitTalents
	spCast := &core.SpellCast{
		Cast: baseCast,
	}

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: *spCast,
		Effect:    effect,
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
	sf.Effect.Target = target
	sf.Init(sim)

	return sf
}
