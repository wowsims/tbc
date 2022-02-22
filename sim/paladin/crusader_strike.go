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
			ActionID:            CrusaderStrikeActionID,
			Character:           &paladin.Character,
			OutcomeRollCategory: core.OutcomeRollCategorySpecial,
			CritRollCategory:    core.CritRollCategoryPhysical,
			SpellSchool:         core.SpellSchoolPhysical,
			GCD:                 core.GCDDefault,
			Cooldown:            time.Second * 6,
			CritMultiplier:      paladin.DefaultMeleeCritMultiplier(),
			IsPhantom:           true,
			Cost: core.ResourceCost{
				Type:  stats.Mana,
				Value: 236,
			},
		},
		Effect: core.AbilityHitEffect{
			AbilityEffect: core.AbilityEffect{
				ProcMask:               core.ProcMaskMeleeMHSpecial,
				DamageMultiplier:       1, // Need to review to make sure I set these properly
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
			},
			WeaponInput: core.WeaponDamageInput{
				DamageMultiplier: 1.1, // maybe this isn't the one that should be set to 1.1
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

	cs.Effect.Target = target

	return cs
}
