package priest

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
)

func (priest *Priest) registerStarshardsSpell(sim *core.Simulation) {
	actionID := core.ActionID{SpellID: 25446}
	priest.Starshards = priest.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolArcane,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    priest.NewTimer(),
				Duration: time.Second * 30,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask:         core.ProcMaskEmpty,
			ThreatMultiplier: 1,
			OutcomeApplier:   priest.OutcomeFuncMagicHit(),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					priest.StarshardsDot.Apply(sim)
				}
			},
		}),
	})

	target := sim.GetPrimaryTarget()
	priest.StarshardsDot = core.NewDot(core.Dot{
		Spell: priest.Starshards,
		Aura: target.RegisterAura(core.Aura{
			Label:    "Starshards-" + strconv.Itoa(int(priest.Index)),
			ActionID: actionID,
		}),

		NumberOfTicks: 5,
		TickLength:    time.Second * 3,

		TickEffects: core.TickFuncSnapshot(target, core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			IsPeriodic:       true,
			BaseDamage:       core.BaseDamageConfigMagicNoRoll(785/5, 0.167),
			OutcomeApplier:   priest.OutcomeFuncTick(),
		}),
	})
}
