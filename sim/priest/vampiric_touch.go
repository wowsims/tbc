package priest

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var VampiricTouchActionID = core.ActionID{SpellID: 34917}

var VampiricTouchAuraID = core.NewAuraID()

func (priest *Priest) newVampiricTouchSpell(sim *core.Simulation, isAltCast bool) *core.Spell {
	cost := core.ResourceCost{Type: stats.Mana, Value: 425}
	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    VampiricTouchActionID,
				Character:   &priest.Character,
				SpellSchool: core.SpellSchoolShadow,
				BaseCost:    cost,
				Cost:        cost,
				CastTime:    time.Millisecond * 1500,
				GCD:         core.GCDDefault,
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
			DamageMultiplier:    1,
			ThreatMultiplier:    1,
			DotInput: core.DotDamageInput{
				NumberOfTicks:  5,
				TickLength:     time.Second * 3,
				TickBaseDamage: core.DotSnapshotFuncMagic(650/5, 0.2),
				AuraID:         VampiricTouchAuraID,
			},
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if isAltCast {
					priest.CurVTSpell = priest.VampiricTouch2
					priest.NextVTSpell = priest.VampiricTouch
				} else {
					priest.CurVTSpell = priest.VampiricTouch
					priest.NextVTSpell = priest.VampiricTouch2
				}
			},
		},
	}

	priest.applyTalentsToShadowSpell(&template.SpellCast.Cast, &template.Effect)

	return priest.RegisterSpell(core.SpellConfig{
		Template:   template,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}
