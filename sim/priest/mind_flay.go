package priest

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDMindFlay int32 = 25387

func (priest *Priest) MindFlayActionID(numTicks int) core.ActionID {
	return core.ActionID{
		SpellID: SpellIDMindFlay,
		Tag:     int32(numTicks),
	}
}

func (priest *Priest) newMindFlaySpell(sim *core.Simulation, numTicks int) *core.Spell {
	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    priest.MindFlayActionID(numTicks),
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
				CastTime:    0,
				ChannelTime: time.Second * time.Duration(numTicks),
				GCD:         core.GCDDefault,
			},
		},
	}
	template.Cost.Value -= template.BaseCost.Value * float64(priest.Talents.FocusedMind) * 0.05

	return priest.RegisterSpell(core.SpellConfig{
		Template: template,
		ModifyCast: func(sim *core.Simulation, target *core.Target, instance *core.SimpleSpell) {
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
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			BonusSpellHitRating: float64(priest.Talents.ShadowFocus) * 2 * core.SpellHitRatingPerHitChance,
			ThreatMultiplier:    1 - 0.08*float64(priest.Talents.ShadowAffinity),
			OutcomeApplier:      core.OutcomeFuncMagicHit(),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					priest.MindFlayDot[numTicks].Apply(sim)
				}
			},
		}),
	})
}

func (priest *Priest) newMindFlayDot(sim *core.Simulation, numTicks int) *core.Dot {
	target := sim.GetPrimaryTarget()
	return core.NewDot(core.Dot{
		Spell: priest.MindFlay[numTicks],
		Aura: target.RegisterAura(&core.Aura{
			Label:    "MindFlay-" + strconv.Itoa(numTicks) + "-" + strconv.Itoa(int(priest.Index)),
			ActionID: priest.MindFlayActionID(numTicks),
		}),

		NumberOfTicks:       numTicks,
		TickLength:          time.Second,
		AffectedByCastSpeed: true,

		TickEffects: core.TickFuncSnapshot(target, core.SpellEffect{
			DamageMultiplier: 1 *
				(1 + float64(priest.Talents.Darkness)*0.02) *
				core.TernaryFloat64(priest.Talents.Shadowform, 1.15, 1) *
				core.TernaryFloat64(ItemSetIncarnate.CharacterHasSetBonus(&priest.Character, 4), 1.05, 1),
			ThreatMultiplier: 1 - 0.08*float64(priest.Talents.ShadowAffinity),
			IsPeriodic:       true,
			BaseDamage:       core.BaseDamageConfigMagicNoRoll(528/3, 0.19),
			OutcomeApplier:   core.OutcomeFuncTick(),
		}),
	})
}
