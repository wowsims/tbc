package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

var FaerieFireDebuffID = core.NewDebuffID()

func (druid *Druid) newFaerieFireTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	return core.NewSimpleSpellTemplate(core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				Name:         "Faerie Fire",
				ActionID:     core.ActionID{SpellID: 26993},
				Character:    druid.GetCharacter(),
				BaseManaCost: 145,
				ManaCost:     145,
			},
		},
		SpellHitEffect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
					spellEffect.Target.AddAura(sim, core.Aura{
						ID:      FaerieFireDebuffID,
						SpellID: 26993,
						Name:    "Faerie Fire",
						Expires: sim.CurrentTime + time.Second*40,
						// TODO: implement increased melee hit
					})
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
	ff.Target = target
	ff.Init(sim)

	return ff
}

func (druid *Druid) ShouldCastFaerieFire(sim *core.Simulation, target *core.Target, rotation proto.BalanceDruid_Rotation) bool {
	return rotation.FaerieFire && !target.HasAura(FaerieFireDebuffID)
}
