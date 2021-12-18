package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDSS int32 = 17364

var StormstrikeCD = core.NewCooldownID()

func (shaman *Shaman) newStormstrikeTemplate(sim *core.Simulation) core.MeleeAbilittyTemplate {
	ss := core.ActiveMeleeAbility{
		MeleeAbility: core.MeleeAbility{
			// ID for the action.
			ActionID: core.ActionID{
				SpellID:    SpellIDSS,
				CooldownID: StormstrikeCD,
			},
			Name:     "Stormstrike",
			Cooldown: time.Second * 10,
			Cost: core.ResourceCost{
				Type:  stats.Mana,
				Value: 237,
			},
			CritMultiplier:  2.0,
			ResetSwingTimer: true,
			Character:       &shaman.Character,
		},
		DirectDamageInput: core.DirectDamageInput{
			MinBaseDamage:    0,
			MaxBaseDamage:    0,
			SpellCoefficient: 0,
			FlatDamageBonus:  0,
		},
		WeaponDamageInput: core.WeaponDamageInput{
			MainHand: 1.0,
			Offhand:  1.0,
		},
	}

	// Add weapon % bonus to stormstrike weapons
	ss.MainHand *= 1 + 0.02*float64(shaman.Talents.WeaponMastery)
	ss.Offhand *= 1 + 0.02*float64(shaman.Talents.WeaponMastery)

	return core.NewMeleeAbilittyTemplate(ss)
}

func (shaman *Shaman) NewStormstrike(sim *core.Simulation, target *core.Target) *core.ActiveMeleeAbility {
	ss := &shaman.stormstrikeSpell
	shaman.stormstrikeTemplate.Apply(ss)
	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	ss.Target = target
	return ss
}
