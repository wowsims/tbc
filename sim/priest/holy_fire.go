package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDHolyFire int32 = 25384

var HolyFireDebuffID = core.NewDebuffID()

func (priest *Priest) newHolyFireTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	baseCast := core.Cast{
		ActionID:            core.ActionID{SpellID: SpellIDHolyFire},
		Character:           &priest.Character,
		CritRollCategory:    core.CritRollCategoryMagical,
		OutcomeRollCategory: core.OutcomeRollCategoryMagic,
		SpellSchool:         core.SpellSchoolHoly,
		BaseCost: core.ResourceCost{
			Type:  stats.Mana,
			Value: 290,
		},
		Cost: core.ResourceCost{
			Type:  stats.Mana,
			Value: 290,
		},
		CastTime:       time.Millisecond * 3500,
		GCD:            core.GCDDefault,
		CritMultiplier: priest.DefaultSpellCritMultiplier(),
	}

	effect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
			ThreatMultiplier:       1,
		},
		DirectInput: core.DirectDamageInput{
			MinBaseDamage:    426,
			MaxBaseDamage:    537,
			SpellCoefficient: 0.8571,
		},
		DotInput: core.DotDamageInput{
			NumberOfTicks:        5,
			TickLength:           time.Second * 2,
			TickBaseDamage:       33,
			TickSpellCoefficient: 0.17,
			DebuffID:             HolyFireDebuffID,
		},
	}
	
	priest.applyTalentsToHolySpell(&baseCast, &effect)
	
	baseCast.CastTime -= time.Millisecond * 100 * time.Duration(priest.Talents.DivineFury)
	
	effect.DamageMultiplier *= (1 + (0.05 * float64(priest.Talents.SearingLight)))

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: baseCast,
		},
		Effect: effect,
	})
}

func (priest *Priest) NewHolyFire(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	hf := &priest.holyFireSpell
	priest.holyFireCastTemplate.Apply(hf)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	hf.Effect.Target = target
	hf.Init(sim)

	return hf
}

