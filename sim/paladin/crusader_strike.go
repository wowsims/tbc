package paladin

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var CrusaderStrikeCD = core.NewCooldownID()
var CrusaderStrikeActionID = core.ActionID{SpellID: 35395, CooldownID: CrusaderStrikeCD}

// Do some research on the spell fields to make sure I'm doing this right
// Need to add in judgement debuff refreshing feature at some point
func (paladin *Paladin) newCrusaderStrikeTemplate(sim *core.Simulation) core.MeleeAbilityTemplate {
	cs := core.ActiveMeleeAbility{
		MeleeAbility: core.MeleeAbility{
			ActionID:    CrusaderStrikeActionID,
			Character:   &paladin.Character,
			SpellSchool: stats.AttackPower,
			GCD:         core.GCDDefault,
			Cooldown:    time.Second * 6,
			Cost: core.ResourceCost{
				Type:  stats.Mana,
				Value: 236,
			},
			CritMultiplier: paladin.DefaultMeleeCritMultiplier(),
		},
		Effects: []core.AbilityHitEffect{
			core.AbilityHitEffect{
				AbilityEffect: core.AbilityEffect{
					DamageMultiplier:       1, // Need to review to make sure I set these properly
					StaticDamageMultiplier: 1,
					ThreatMultiplier:       1,
				},
				WeaponInput: core.WeaponDamageInput{
					IsOH:             false,
					DamageMultiplier: 1.1,  // maybe this isn't the one that should be set to 1.1
				},
			},
		},
		OnMeleeAttack: func(sim *core.Simulation, ability *core.ActiveMeleeAbility, hitEffect *core.AbilityHitEffect) {
			if !hitEffect.Landed() {
				return
			}
		},
	}

	return core.NewMeleeAbilityTemplate(cs)
}

func (paladin *Paladin) NewCrusaderStrike(sim *core.Simulation, target *core.Target) *core.ActiveMeleeAbility {
	cs := &paladin.crusaderStrikeSpell
	paladin.crusaderStrikeTemplate.Apply(cs)
	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	cs.Effects[0].Target = target
	return cs
}
