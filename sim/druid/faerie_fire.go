package druid

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (druid *Druid) registerFaerieFireSpell(sim *core.Simulation) {
	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    core.ActionID{SpellID: 26993},
				SpellSchool: core.SpellSchoolNature,
				Character:   druid.GetCharacter(),
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 145,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 145,
				},
				GCD: core.GCDDefault,
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			ThreatMultiplier:    1,
			FlatThreatBonus:     0, // TODO
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					return
				}
				// core.FaerieFireAura applies the -armor buff and removes it on expire.
				//  Don't use ReplaceAura or the armor won't be removed.
				spellEffect.Target.AddAura(sim, core.FaerieFireAura(spellEffect.Target, druid.Talents.ImprovedFaerieFire == 3))
			},
		},
	}

	druid.FaerieFire = druid.RegisterSpell(core.SpellConfig{
		Template:   template,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}

func (druid *Druid) ShouldCastFaerieFire(sim *core.Simulation, target *core.Target, rotation proto.BalanceDruid_Rotation) bool {
	return rotation.FaerieFire && !target.HasAura(core.FaerieFireDebuffID)
}
