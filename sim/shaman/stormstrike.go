package shaman

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDSS int32 = 17364

var StormstrikeCD = core.NewCooldownID()
var StormstrikeDebuffID = core.NewDebuffID()

func (shaman *Shaman) newStormstrikeTemplate(sim *core.Simulation) core.MeleeAbilittyTemplate {

	ssDebuffAura := core.Aura{
		ID:     StormstrikeDebuffID,
		Name:   "Stormstrike",
		Stacks: 2,
	}
	ssDebuffAura.OnBeforeSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		if spellCast.SpellSchool != stats.NatureSpellPower {
			return
		}
		spellEffect.DamageMultiplier *= 1.2
		stacks := spellEffect.Target.NumStacks(StormstrikeDebuffID) - 1
		if stacks == 0 {
			spellEffect.Target.RemoveAura(sim, StormstrikeDebuffID)
		} else {
			ssDebuffAura.Stacks = stacks
			spellEffect.Target.ReplaceAura(sim, ssDebuffAura)
		}
	}

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
		OnMeleeAttack: func(sim *core.Simulation, target *core.Target, result core.MeleeHitType, ability *core.ActiveMeleeAbility, isOH bool) {
			ssDebuffAura.Stacks = 2
			target.ReplaceAura(sim, ssDebuffAura)
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
