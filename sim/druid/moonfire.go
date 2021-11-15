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
		ManaCost:       495,
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
		DotInput: core.DotDamageInput{
			Name:             "Moonfire DoT",
			NumberTicks:      4,
			TickLength:       time.Second * 3,
			BaseDamage:       600,
			SpellCoefficient: 0.13,

			// TODO: does druid care about dot ticks?
			// OnDamageTick: func(sim *core.Simulation) {},
		},
	}

	if druid.thunder2p { // Thunderheart 2p adds 1 extra tick to moonfire
		effect.DotInput.NumberTicks += 1
	}

	// Moonfire only talents
	effect.SpellEffect.DamageMultiplier *= 1 + 0.05*float64(druid.Talents.ImprovedMoonfire)
	effect.SpellEffect.BonusSpellCritRating += float64(druid.Talents.ImprovedMoonfire) * 5 * core.SpellCritRatingPerCritChance

	// TODO: Shared talents
	effect.SpellEffect.BonusSpellCritRating += float64(druid.Talents.FocusedStarlight) * 2 * core.SpellCritRatingPerCritChance // 2% crit per point
	baseCast.ManaCost -= baseCast.BaseManaCost * float64(druid.Talents.Moonglow) * 0.03
	effect.SpellEffect.DamageMultiplier *= 1 + 0.02*float64(druid.Talents.Moonfury)
	// Convert to percent, multiply by percent increase, convert back to multiplier by adding 1
	baseCast.CritMultiplier = (baseCast.CritMultiplier-1)*(1+float64(druid.Talents.Vengeance)*0.2) + 1

	// moonfire can proc the on hit but doesn't consume charges (i think)
	effect.OnSpellHit = druid.applyOnHitTalents

	return core.NewDamageOverTimeSpellTemplate(core.DamageOverTimeSpell{
		SpellCast: core.SpellCast{
			Cast: baseCast,
		},
		DamageOverTimeSpellEffect: effect,
	})
}

// TODO: This might behave weird if we have a moonfire still ticking when we cast one.
//   We could do a check and if the spell is still ticking allocate a new DamageOverTimeSpell?
func (druid *Druid) NewMoonfire(sim *core.Simulation, target *core.Target) *core.DamageOverTimeSpell {
	// Initialize cast from precomputed template.
	sf := &druid.MoonfireSpell
	druid.moonfireCastTemplate.Apply(sf)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	sf.Target = target
	sf.Init(sim)

	return sf
}
