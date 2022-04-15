package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDConsecration int32 = 27173

var ConsecrationActionID = core.ActionID{SpellID: SpellIDConsecration}

func (paladin *Paladin) registerConsecrationSpell(sim *core.Simulation) {
	baseCost := 660.0

	consecrationDot := core.NewDot(core.Dot{
		Aura: paladin.RegisterAura(core.Aura{
			Label:    "Consecration",
			ActionID: ConsecrationActionID,
		}),
		NumberOfTicks: 8,
		TickLength:    time.Second * 1,
		TickEffects: core.TickFuncAOESnapshot(sim, core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BaseDamage:       core.BaseDamageConfigMagicNoRoll(64, 0.119),
			OutcomeApplier:   core.OutcomeFuncTick(),
			IsPeriodic:       true,
		}),
	})
	// TODO: consecration talents here

	paladin.Consecration = paladin.RegisterSpell(core.SpellConfig{
		ActionID:    ConsecrationActionID,
		SpellSchool: core.SpellSchoolHoly,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost,
				GCD:  core.GCDDefault,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDot(consecrationDot),
	})
	consecrationDot.Spell = paladin.Consecration
}
