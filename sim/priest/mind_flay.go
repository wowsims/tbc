package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDMindFlay int32 = 25387

const TagMF2 = 2
const TagMF3 = 3

func (priest *Priest) newMindFlaySpell(sim *core.Simulation, numTicks int) *core.SimpleSpellTemplate {
	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID: core.ActionID{
					SpellID: SpellIDMindFlay,
					Tag:     int32(numTicks),
				},
				Character:   &priest.Character,
				SpellSchool: core.SpellSchoolShadow,
				SpellExtras: core.SpellExtrasBinary | core.SpellExtrasChanneled,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 230,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 230,
				},
				CastTime: 0,
				GCD:      core.GCDDefault,
			},
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
		},
		Effect: core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			DotInput: core.DotDamageInput{
				NumberOfTicks:       numTicks,
				TickLength:          time.Second,
				TickBaseDamage:      core.DotSnapshotFuncMagic(528/3, 0.19),
				AffectedByCastSpeed: true,
			},
		},
	}

	priest.applyTalentsToShadowSpell(&template.SpellCast.Cast, &template.Effect)

	if ItemSetIncarnate.CharacterHasSetBonus(&priest.Character, 4) {
		template.Effect.DamageMultiplier *= 1.05
	}

	return priest.RegisterSpell(core.SpellConfig{
		Template: template,
		ModifyCast: func(sim *core.Simulation, target *core.Target, instance *core.SimpleSpell) {
			instance.Effect.Target = target

			// if our channel is longer than GCD it will have human latency to end it beause you can't queue the next spell.
			var wait time.Duration // TODO: I think this got deleted at some point
			gcd := core.MinDuration(core.GCDMin, time.Duration(float64(core.GCDDefault)/priest.CastSpeed()))
			if wait > gcd && priest.Latency > 0 {
				base := priest.Latency * 0.66
				variation := base + sim.RandomFloat("spriest latency")*base // should vary from 0.66 - 1.33 of given latency
				variation = core.MaxFloat(variation, 10)                    // no player can go under XXXms response time
				instance.AfterCastDelay += time.Duration(variation) * time.Millisecond
			}
		},
	})
}
