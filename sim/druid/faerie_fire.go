package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

var FaerieFireDebuffID = core.NewDebuffID()

// Returns the time to wait before the next action, or 0 if faerie fire is already active.
func (druid *Druid) TryFaerieFire(sim *core.Simulation, target *core.Target) time.Duration {
	if target.HasAura(FaerieFireDebuffID) {
		return 0
	}

	cast := &core.SimpleSpell{
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
				Target: target,

				OnSpellHit: func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
					target.AddAura(sim, core.Aura{
						ID:      FaerieFireDebuffID,
						Name:    "Faerie Fire",
						Expires: sim.CurrentTime + time.Second*40,
						// TODO: implement increased melee hit
					})
				},
			},
		},
	}
	cast.Init(sim)

	success := cast.Cast(sim)
	if !success {
		regenTime := druid.TimeUntilManaRegen(cast.GetManaCost())
		return sim.CurrentTime + regenTime
	}
	return sim.CurrentTime + druid.GetRemainingCD(core.GCDCooldownID, sim.CurrentTime)
}
