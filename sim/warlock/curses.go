package warlock

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var CurseOfElementsActionID = core.ActionID{SpellID: 27228}

func (warlock *Warlock) registerCurseOfElementsSpell(sim *core.Simulation) {
	baseCost := 145.0
	warlock.CurseOfElementsAura = core.CurseOfElementsAura(sim.GetPrimaryTarget(), warlock.Talents.Malediction)
	warlock.CurseOfElementsAura.Duration = time.Minute * 5

	warlock.CurseOfElements = warlock.RegisterSpell(core.SpellConfig{
		ActionID:    CurseOfElementsActionID,
		SpellSchool: core.SpellSchoolShadow,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost,
				GCD:  core.GCDDefault,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ThreatMultiplier: 1,
			FlatThreatBonus:  0, // TODO
			OutcomeApplier:   core.OutcomeFuncMagicHit(),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					warlock.CurseOfElementsAura.Activate(sim)
				}
			},
		}),
	})
}

func (warlock *Warlock) ShouldCastCurseOfElements(sim *core.Simulation, target *core.Target, curse proto.Warlock_Rotation_Curse) bool {
	return curse == proto.Warlock_Rotation_Elements && !warlock.CurseOfElementsAura.IsActive()
}
