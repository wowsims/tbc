package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDMindBlast int32 = 25375

var MBCooldownID = core.NewCooldownID()

func (priest *Priest) newMindBlastTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	baseCast := core.Cast{
		ActionID: core.ActionID{
			SpellID:    SpellIDMindBlast,
			CooldownID: MBCooldownID,
		},
		Character:           &priest.Character,
		CritRollCategory:    core.CritRollCategoryMagical,
		OutcomeRollCategory: core.OutcomeRollCategoryMagic,
		SpellSchool:         core.SpellSchoolShadow,
		BaseCost: core.ResourceCost{
			Type:  stats.Mana,
			Value: 450,
		},
		Cost: core.ResourceCost{
			Type:  stats.Mana,
			Value: 450,
		},
		CastTime:       time.Millisecond * 1500,
		GCD:            core.GCDDefault,
		Cooldown:       time.Second * 8,
		CritMultiplier: priest.DefaultSpellCritMultiplier(),
	}

	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
			ThreatMultiplier:       1,
		},
		DirectInput: core.DirectDamageInput{
			MinBaseDamage:    711,
			MaxBaseDamage:    752,
			SpellCoefficient: 0.429,
		},
	}

	priest.applyTalentsToShadowSpell(&baseCast, &effect)

	baseCast.Cooldown -= time.Millisecond * 500 * time.Duration(priest.Talents.ImprovedMindBlast)

	effect.BonusHitRating += float64(priest.Talents.FocusedPower) * 2 * core.SpellHitRatingPerHitChance // 2% crit per point

	if ItemSetAbsolution.CharacterHasSetBonus(&priest.Character, 4) { // Absolution 4p adds 10% damage
		effect.StaticDamageMultiplier *= 1.1
	}

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: baseCast,
		},
		Effect: effect,
	})
}

func (priest *Priest) NewMindBlast(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	mf := &priest.mindblastSpell

	priest.mindblastCastTemplate.Apply(mf)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	mf.Effect.Target = target
	mf.Init(sim)

	return mf
}
