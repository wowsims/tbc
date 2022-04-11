package hunter

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (hunter *Hunter) registerScorpidStingSpell(sim *core.Simulation) {
	actionID := core.ActionID{SpellID: 3043}
	cost := core.ResourceCost{Type: stats.Mana, Value: hunter.BaseMana() * 0.09}
	hunter.ScorpidStingAura = core.ScorpidStingAura(sim.GetPrimaryTarget())

	ama := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    actionID,
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

	hunter.ScorpidSting = hunter.RegisterSpell(core.SpellConfig{
		Template: ama,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ThreatMultiplier: 1,
			OutcomeApplier:   core.OutcomeFuncRangedHit(),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					hunter.ScorpidStingAura.Activate(sim)
				}
			},
		}),
	})
}
