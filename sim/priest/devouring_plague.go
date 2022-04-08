package priest

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDDevouringPlague int32 = 25467

var DevouringPlagueCooldownID = core.NewCooldownID()
var DevouringPlagueActionID = core.ActionID{SpellID: SpellIDDevouringPlague, CooldownID: DevouringPlagueCooldownID}

func (priest *Priest) registerDevouringPlagueSpell(sim *core.Simulation) {
	cost := core.ResourceCost{Type: stats.Mana, Value: 1145}

	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    DevouringPlagueActionID,
				Character:   &priest.Character,
				SpellSchool: core.SpellSchoolShadow,
				BaseCost:    cost,
				Cost:        cost,
				CastTime:    0,
				GCD:         core.GCDDefault,
				Cooldown:    time.Minute * 3,
			},
		},
	}
	template.Cost.Value -= template.BaseCost.Value * float64(priest.Talents.MentalAgility) * 0.02

	priest.DevouringPlague = priest.RegisterSpell(core.SpellConfig{
		Template: template,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			BonusSpellHitRating: float64(priest.Talents.ShadowFocus) * 2 * core.SpellHitRatingPerHitChance,
			ThreatMultiplier:    1 - 0.08*float64(priest.Talents.ShadowAffinity),
			OutcomeApplier:      core.OutcomeFuncMagicHit(),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					priest.DevouringPlagueDot.Apply(sim)
				}
			},
		}),
	})

	target := sim.GetPrimaryTarget()
	priest.DevouringPlagueDot = core.NewDot(core.Dot{
		Spell: priest.DevouringPlague,
		Aura: target.RegisterAura(&core.Aura{
			Label:    "DevouringPlague-" + strconv.Itoa(int(priest.Index)),
			ActionID: DevouringPlagueActionID,
		}),
		NumberOfTicks: 8,
		TickLength:    time.Second * 3,
		TickEffects: core.TickFuncSnapshot(target, core.SpellEffect{
			DamageMultiplier: 1 *
				(1 + float64(priest.Talents.Darkness)*0.02) *
				core.TernaryFloat64(priest.Talents.Shadowform, 1.15, 1),
			ThreatMultiplier: 1 - 0.08*float64(priest.Talents.ShadowAffinity),
			IsPeriodic:       true,
			BaseDamage:       core.BaseDamageConfigMagicNoRoll(1216/8, 0.1),
			OutcomeApplier:   core.OutcomeFuncTick(),
		}),
	})
}
