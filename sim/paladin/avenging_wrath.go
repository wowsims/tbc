package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var AvengingWrathActionID = core.ActionID{SpellID: 31884}

func (paladin *Paladin) registerAvengingWrathCD() {
	aura := paladin.RegisterAura(core.Aura{
		Label:    "Avenging Wrath",
		ActionID: AvengingWrathActionID,
		Duration: time.Second * 20,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.PseudoStats.DamageDealtMultiplier *= 1.3
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.PseudoStats.DamageDealtMultiplier /= 1.3
		},
	})

	baseCost := 236.0

	spell := paladin.RegisterSpell(core.SpellConfig{
		ActionID: AvengingWrathActionID,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost,
			},
			CD: core.Cooldown{
				Timer:    paladin.NewTimer(),
				Duration: time.Minute * 3,
			},
			DisableCallbacks: true,
		},
		ApplyEffects: func(sim *core.Simulation, _ *core.Target, _ *core.Spell) {
			aura.Activate(sim)
		},
	})

	paladin.AddMajorCooldown(core.MajorCooldown{
		Spell: spell,
		Type:  core.CooldownTypeDPS,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			return character.CurrentMana() >= spell.DefaultCast.Cost
		},
		// modify this logic if it should ever not be spammed on CD / maybe should synced with other CDs
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
	})
}
