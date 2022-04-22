package warlock

import (
	"github.com/wowsims/tbc/sim/core"
)

var LifeTapActionID = core.ActionID{SpellID: 27222}

func (warlock *Warlock) registerLifeTapSpell(sim *core.Simulation) {
	warlock.LifeTap = warlock.RegisterSpell(core.SpellConfig{
		ActionID:    LifeTapActionID,
		SpellSchool: core.SpellSchoolShadow,
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ThreatMultiplier: 1,
			FlatThreatBonus:  0, // TODO
			OutcomeApplier:   core.OutcomeFuncMagicHit(),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				warlock.AddMana(sim, 582, LifeTapActionID, true)
			},
		}),
	})
}
