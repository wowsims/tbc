package druid

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (druid *Druid) registerFaerieFireSpell(sim *core.Simulation) {
	druid.FaerieFireAura = core.FaerieFireAura(sim.GetPrimaryTarget(), druid.Talents.ImprovedFaerieFire)

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
	}

	druid.FaerieFire = druid.RegisterSpell(core.SpellConfig{
		Template: template,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ThreatMultiplier: 1,
			FlatThreatBonus:  0, // TODO
			OutcomeApplier:   core.OutcomeFuncMagicHit(),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					druid.FaerieFireAura.Activate(sim)
				}
			},
		}),
	})
}

func (druid *Druid) ShouldCastFaerieFire(sim *core.Simulation, target *core.Target, rotation proto.BalanceDruid_Rotation) bool {
	return rotation.FaerieFire && !druid.FaerieFireAura.IsActive()
}
