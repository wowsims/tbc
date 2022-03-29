package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDArcaneMissiles int32 = 38699

// Note: AM doesn't charge its mana up-front, instead it charges 1/5 of the mana on each tick.
// This is probably not worth simming since no other spell in the game does this and AM isn't
// even a popular choice for arcane mages.
func (mage *Mage) newArcaneMissilesTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            core.ActionID{SpellID: SpellIDArcaneMissiles},
				Character:           &mage.Character,
				CritRollCategory:    core.CritRollCategoryMagical,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				SpellSchool:         core.SpellSchoolArcane,
				SpellExtras:         SpellFlagMage | core.SpellExtrasChanneled | core.SpellExtrasAlwaysHits,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 740,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 740,
				},
				GCD:            core.GCDDefault,
				CritMultiplier: mage.SpellCritMultiplier(1, 0.25*float64(mage.Talents.SpellPower)),
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				DamageMultiplier: mage.spellDamageMultiplier,
				ThreatMultiplier: 1 - 0.2*float64(mage.Talents.ArcaneSubtlety),
			},
			DotInput: core.DotDamageInput{
				NumberOfTicks:       5,
				TickLength:          time.Second,
				TickBaseDamage:      core.DotSnapshotFuncMagic(265, 1/3.5+0.15*float64(mage.Talents.EmpoweredArcaneMissiles)),
				TicksCanMissAndCrit: true,
				AffectedByCastSpeed: true,

				TicksProcSpellHitEffects: true,
			},
		},
	}

	spell.Effect.BonusSpellHitRating += float64(mage.Talents.ArcaneFocus) * 2 * core.SpellHitRatingPerHitChance
	spell.Cost.Value += spell.BaseCost.Value * float64(mage.Talents.EmpoweredArcaneMissiles) * 0.02

	if ItemSetTempestRegalia.CharacterHasSetBonus(&mage.Character, 4) {
		spell.Effect.DamageMultiplier *= 1.05
	}

	return core.NewSimpleSpellTemplate(spell)
}

func (mage *Mage) NewArcaneMissiles(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	arcaneMissiles := &mage.arcaneMissilesSpell
	mage.arcaneMissilesCastTemplate.Apply(arcaneMissiles)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	arcaneMissiles.Effect.Target = target
	arcaneMissiles.Init(sim)

	return arcaneMissiles
}
