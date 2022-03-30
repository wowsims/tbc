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
		ActionID:    core.ActionID{SpellID: SpellIDHolyFire},
		Character:   &priest.Character,
		SpellSchool: core.SpellSchoolHoly,
		BaseCost: core.ResourceCost{
			Type:  stats.Mana,
			Value: 290,
		},
		Cost: core.ResourceCost{
			Type:  stats.Mana,
			Value: 290,
		},
		CastTime: time.Millisecond * 3500,
		GCD:      core.GCDDefault,
	}

	effect := core.SpellEffect{
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		BaseDamage:       core.BaseDamageConfigMagic(426, 537, 0.8571),
		DotInput: core.DotDamageInput{
			NumberOfTicks:  5,
			TickLength:     time.Second * 2,
			TickBaseDamage: core.DotSnapshotFuncMagic(33, 0.17),
			DebuffID:       HolyFireDebuffID,
		},
	}

	priest.applyTalentsToHolySpell(&baseCast, &effect)

	baseCast.CastTime -= time.Millisecond * 100 * time.Duration(priest.Talents.DivineFury)

	effect.DamageMultiplier *= (1 + (0.05 * float64(priest.Talents.SearingLight)))

	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast:                baseCast,
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
			CritMultiplier:      priest.DefaultSpellCritMultiplier(),
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
