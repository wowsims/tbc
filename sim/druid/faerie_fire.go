package druid

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var FaerieFireActionID = core.ActionID{SpellID: 26993}

func (druid *Druid) registerFaerieFireSpell(sim *core.Simulation) {
	baseCost := 145.0
	druid.FaerieFireAura = core.FaerieFireAura(sim.GetPrimaryTarget(), druid.Talents.ImprovedFaerieFire)

	druid.FaerieFire = druid.RegisterSpell(core.SpellConfig{
		ActionID:    FaerieFireActionID,
		SpellSchool: core.SpellSchoolNature,

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
					druid.FaerieFireAura.Activate(sim)
				}
			},
		}),
	})
}

func (druid *Druid) ShouldCastFaerieFire(sim *core.Simulation, target *core.Target, rotation proto.BalanceDruid_Rotation) bool {
	return rotation.FaerieFire && !druid.FaerieFireAura.IsActive()
}
