package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

const SpellIDStarshards int32 = 25446

var SSCooldownID = core.NewCooldownID()

func (priest *Priest) registerStarshardsSpell(sim *core.Simulation) {
	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID: core.ActionID{
					SpellID:    SpellIDStarshards,
					CooldownID: SSCooldownID,
				},
				Character:   &priest.Character,
				SpellSchool: core.SpellSchoolArcane,
				CastTime:    0,
				GCD:         core.GCDDefault,
				Cooldown:    time.Second * 30,
			},
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
		},
		Effect: core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			DotInput: core.DotDamageInput{
				NumberOfTicks:  5,
				TickLength:     time.Second * 3,
				TickBaseDamage: core.DotSnapshotFuncMagic(785/5, 0.167),
			},
		},
	}

	priest.Starshards = priest.RegisterSpell(core.SpellConfig{
		Template:   template,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}
