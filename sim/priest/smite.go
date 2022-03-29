package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDSmite int32 = 25364

var SmiteCooldownID = core.NewCooldownID()

func (priest *Priest) newSmiteTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	baseCast := core.Cast{
		ActionID: core.ActionID{
			SpellID:    SpellIDSmite,
			CooldownID: SmiteCooldownID,
		},
		Character:           &priest.Character,
		CritRollCategory:    core.CritRollCategoryMagical,
		OutcomeRollCategory: core.OutcomeRollCategoryMagic,
		SpellSchool:         core.SpellSchoolHoly,
		BaseCost: core.ResourceCost{
			Type:  stats.Mana,
			Value: 385,
		},
		Cost: core.ResourceCost{
			Type:  stats.Mana,
			Value: 385,
		},
		CastTime:       time.Millisecond * 2500,
		GCD:            core.GCDDefault,
		Cooldown:       time.Second * 0,
		CritMultiplier: priest.DefaultSpellCritMultiplier(),
	}

	effect := core.SpellEffect{
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BaseDamage:       core.BaseDamageConfigMagic(549, 616, 0.7143),
	}

	priest.applyTalentsToHolySpell(&baseCast, &effect)

	baseCast.CastTime -= time.Millisecond * 100 * time.Duration(priest.Talents.DivineFury)

	effect.DamageMultiplier *= (1 + (0.05 * float64(priest.Talents.SearingLight)))

	effect.BonusSpellHitRating += float64(priest.Talents.FocusedPower) * 2 * core.SpellHitRatingPerHitChance // 2% crit per point

	effect.OnSpellHit = priest.applyOnHitTalents

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: baseCast,
		},
		Effect: effect,
	})
}

func (priest *Priest) NewSmite(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	mf := &priest.smiteSpell

	priest.smiteCastTemplate.Apply(mf)

	priest.applySurgeOfLight(&mf.SpellCast)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	mf.Effect.Target = target
	mf.Init(sim)

	return mf
}
