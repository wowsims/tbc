package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func (rogue *Rogue) registerThistleTeaCD() {
	if rogue.Consumes.DefaultConjured != proto.Conjured_ConjuredRogueThistleTea {
		return
	}

	actionID := core.ActionID{ItemID: 7676}

	const energyRegen = 40.0

	thistleTeaSpell := rogue.RegisterSpell(core.SpellConfig{
		ActionID: actionID,

		Cast: core.CastConfig{
			// TODO: Thistle tea might only put conjured on shorter CD
			CD: core.Cooldown{
				Timer:    rogue.GetConjuredCD(),
				Duration: time.Minute * 5,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Target, _ *core.Spell) {
			rogue.AddEnergy(sim, energyRegen, actionID)
		},
	})

	rogue.AddMajorCooldown(core.MajorCooldown{
		Spell: thistleTeaSpell,
		Type:  core.CooldownTypeDPS,
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			// Make sure we have plenty of room so we dont energy cap right after using.
			return rogue.CurrentEnergy() <= 40
		},
	})
}
