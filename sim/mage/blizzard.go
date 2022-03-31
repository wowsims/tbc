package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDBlizzard int32 = 27085

func (mage *Mage) registerBlizzardSpell(sim *core.Simulation) {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID: core.ActionID{
					SpellID: SpellIDBlizzard,
				},
				SpellSchool: core.SpellSchoolFrost,
				SpellExtras: SpellFlagMage | core.SpellExtrasChanneled | core.SpellExtrasAlwaysHits,
				Character:   &mage.Character,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 1645,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 1645,
				},
				GCD: core.GCDDefault,
			},
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
		},
		AOECap: 3620,
	}

	baseEffect := core.SpellEffect{
		DamageMultiplier: mage.spellDamageMultiplier,
		ThreatMultiplier: 1 - (0.1/3)*float64(mage.Talents.FrostChanneling),
		DotInput: core.DotDamageInput{
			NumberOfTicks:       8,
			TickLength:          time.Second * 1,
			TickBaseDamage:      core.DotSnapshotFuncMagic(184, 0.119),
			AffectedByCastSpeed: true,
		},
	}

	spell.Cost.Value -= spell.BaseCost.Value * float64(mage.Talents.FrostChanneling) * 0.05
	spell.Cost.Value *= 1 - float64(mage.Talents.ElementalPrecision)*0.01
	baseEffect.BonusSpellHitRating += float64(mage.Talents.ElementalPrecision) * 1 * core.SpellHitRatingPerHitChance
	baseEffect.DamageMultiplier *= 1 + 0.02*float64(mage.Talents.PiercingIce)
	baseEffect.DamageMultiplier *= 1 + 0.01*float64(mage.Talents.ArcticWinds)

	numHits := sim.GetNumTargets()
	effects := make([]core.SpellEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)
	}
	spell.Effects = effects

	mage.Blizzard = mage.RegisterSpell(core.SpellConfig{
		Template: spell,
	})
}
