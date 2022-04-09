package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDConsecration int32 = 27173

var ConsecrationActionID = core.ActionID{SpellID: SpellIDConsecration}

func (paladin *Paladin) registerConsecrationSpell(sim *core.Simulation) {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    ConsecrationActionID,
				Character:   &paladin.Character,
				SpellSchool: core.SpellSchoolHoly,
				SpellExtras: core.SpellExtrasAlwaysHits,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 660,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 660,
				},
				GCD: core.GCDDefault,
			},
		},
	}

	// TODO: consecration talents here

	consecrationDot := core.NewDot(core.Dot{
		Aura: paladin.RegisterAura(&core.Aura{
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

	paladin.Consecration = paladin.RegisterSpell(core.SpellConfig{
		Template:     spell,
		ApplyEffects: core.ApplyEffectFuncDot(consecrationDot),
	})
	consecrationDot.Spell = paladin.Consecration
}
