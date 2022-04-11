package hunter

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var SerpentStingActionID = core.ActionID{SpellID: 27016}

func (hunter *Hunter) registerSerpentStingSpell(sim *core.Simulation) {
	cost := core.ResourceCost{Type: stats.Mana, Value: 275}
	ama := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    SerpentStingActionID,
				Character:   &hunter.Character,
				SpellSchool: core.SpellSchoolNature,
				GCD:         core.GCDDefault,
				Cost:        cost,
				BaseCost:    cost,
				IgnoreHaste: true, // Hunter GCD is locked at 1.5s
			},
		},
	}
	ama.Cost.Value *= 1 - 0.02*float64(hunter.Talents.Efficiency)

	hunter.SerpentSting = hunter.RegisterSpell(core.SpellConfig{
		Template: ama,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ThreatMultiplier: 1,
			OutcomeApplier:   core.OutcomeFuncRangedHit(),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					hunter.SerpentStingDot.Apply(sim)
				}
			},
		}),
	})

	target := sim.GetPrimaryTarget()
	hunter.SerpentStingDot = core.NewDot(core.Dot{
		Spell: hunter.SerpentSting,
		Aura: target.RegisterAura(&core.Aura{
			Label:    "SerpentSting-" + strconv.Itoa(int(hunter.Index)),
			ActionID: SerpentStingActionID,
		}),
		NumberOfTicks: 5,
		TickLength:    time.Second * 3,
		TickEffects: core.TickFuncSnapshot(target, core.SpellEffect{
			DamageMultiplier: 1 + 0.06*float64(hunter.Talents.ImprovedStings),
			ThreatMultiplier: 1,
			IsPeriodic:       true,

			BaseDamage: core.BuildBaseDamageConfig(func(sim *core.Simulation, spellEffect *core.SpellEffect, spell *core.Spell) float64 {
				attackPower := spellEffect.RangedAttackPower(spell.Character) + spellEffect.RangedAttackPowerOnTarget()
				return 132 + attackPower*0.02
			}, 0),
			OutcomeApplier: core.OutcomeFuncTick(),
		}),
	})
}
