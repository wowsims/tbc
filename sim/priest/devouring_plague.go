package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDDevouringPlague int32 = 25467

var DevouringPlagueCooldownID = core.NewCooldownID()

func (priest *Priest) registerDevouringPlagueSpell(sim *core.Simulation) {
	cost := core.ResourceCost{Type: stats.Mana, Value: 1145}

	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID: core.ActionID{
					SpellID:    SpellIDDevouringPlague,
					CooldownID: DevouringPlagueCooldownID,
				},
				Character:   &priest.Character,
				SpellSchool: core.SpellSchoolShadow,
				BaseCost:    cost,
				Cost:        cost,
				CastTime:    0,
				GCD:         core.GCDDefault,
				Cooldown:    time.Minute * 3,
			},
		},
		Effect: core.SpellEffect{
			CritRollCategory:    core.CritRollCategoryMagical,
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			DamageMultiplier:    1,
			ThreatMultiplier:    1,
			DotInput: core.DotDamageInput{
				NumberOfTicks:  8,
				TickLength:     time.Second * 3,
				TickBaseDamage: core.DotSnapshotFuncMagic(1216/8, 0.1),
				Aura:           priest.NewDotAura("Devouring Plague", core.ActionID{SpellID: SpellIDDevouringPlague}),
			},
		},
	}

	priest.applyTalentsToShadowSpell(&template.SpellCast.Cast, &template.Effect)

	priest.DevouringPlague = priest.RegisterSpell(core.SpellConfig{
		Template:   template,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}
