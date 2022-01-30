package druid

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func (druid *Druid) newFaerieFireTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:     core.ActionID{SpellID: 26993},
				Character:    druid.GetCharacter(),
				BaseManaCost: 145,
				ManaCost:     145,
				GCD:          core.GCDDefault,
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ThreatMultiplier: 1,
				FlatThreatBonus:  0, // TODO
				OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
					// core.FaerieFireAura applies the -armor buff and removes it on expire.
					//  Don't use ReplaceAura or the armor won't be removed.
					spellEffect.Target.AddAura(sim, core.FaerieFireAura(sim.CurrentTime, spellEffect.Target, druid.Talents.ImprovedFaerieFire == 3))
				},
			},
		},
	})
}

func (druid *Druid) NewFaerieFire(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	// Initialize cast from precomputed template.
	ff := &druid.FaerieFireSpell
	druid.faerieFireCastTemplate.Apply(ff)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	ff.Effect.Target = target
	ff.Init(sim)

	return ff
}

func (druid *Druid) ShouldCastFaerieFire(sim *core.Simulation, target *core.Target, rotation proto.BalanceDruid_Rotation) bool {
	return rotation.FaerieFire && !target.HasAura(core.FaerieFireDebuffID)
}
