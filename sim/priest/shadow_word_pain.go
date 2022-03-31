package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDShadowWordPain int32 = 25368

var ShadowWordPainActionID = core.ActionID{SpellID: SpellIDShadowWordPain}

var ShadowWordPainDebuffID = core.NewDebuffID()

func (priest *Priest) registerShadowWordPainSpell(sim *core.Simulation) {
	cost := core.ResourceCost{Type: stats.Mana, Value: 575}
	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    ShadowWordPainActionID,
				Character:   &priest.Character,
				SpellSchool: core.SpellSchoolShadow,
				BaseCost:    cost,
				Cost:        cost,
				CastTime:    0,
				GCD:         core.GCDDefault,
			},
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
		},
		Effect: core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			DotInput: core.DotDamageInput{
				NumberOfTicks:  6,
				TickLength:     time.Second * 3,
				TickBaseDamage: core.DotSnapshotFuncMagic(1236/6, 0.183),
				DebuffID:       ShadowWordPainDebuffID,
			},
		},
	}

	template.Effect.DotInput.NumberOfTicks += int(priest.Talents.ImprovedShadowWordPain) // extra tick per point

	if ItemSetAbsolution.CharacterHasSetBonus(&priest.Character, 2) { // Absolution 2p adds 1 extra tick to swp
		template.Effect.DotInput.NumberOfTicks += 1
	}

	priest.applyTalentsToShadowSpell(&template.SpellCast.Cast, &template.Effect)

	priest.ShadowWordPain = priest.RegisterSpell(core.SpellConfig{
		Template:   template,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}
