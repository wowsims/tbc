package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var RaptorStrikeCooldownID = core.NewCooldownID()
var RaptorStrikeActionID = core.ActionID{SpellID: 27014, CooldownID: RaptorStrikeCooldownID}

func (hunter *Hunter) newRaptorStrikeTemplate(sim *core.Simulation) core.MeleeAbilityTemplate {
	ama := core.ActiveMeleeAbility{
		MeleeAbility: core.MeleeAbility{
			ActionID:    RaptorStrikeActionID,
			Character:   &hunter.Character,
			SpellSchool: stats.AttackPower,
			Cost: core.ResourceCost{
				Type:  stats.Mana,
				Value: 120,
			},
			Cooldown:       time.Second * 6,
			CritMultiplier: hunter.critMultiplier(false, sim.GetPrimaryTarget()),
		},
		Effect: core.AbilityHitEffect{
			AbilityEffect: core.AbilityEffect{
				ProcMask:               core.ProcMaskMeleeMHSpecial,
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
			},
			WeaponInput: core.WeaponDamageInput{
				DamageMultiplier: 1,
				FlatDamageBonus:  170,
			},
		},
	}

	ama.Cost.Value -= 120 * 0.2 * float64(hunter.Talents.Resourcefulness)
	ama.Effect.BonusCritRating += float64(hunter.Talents.SavageStrikes) * 10 * core.MeleeCritRatingPerCritChance

	hunter.raptorStrikeCost = ama.Cost.Value

	return core.NewMeleeAbilityTemplate(ama)
}

func (hunter *Hunter) NewRaptorStrike(sim *core.Simulation, target *core.Target) *core.ActiveMeleeAbility {
	rs := &hunter.raptorStrike
	hunter.raptorStrikeTemplate.Apply(rs)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	rs.Effect.Target = target

	return rs
}

// Returns true if the regular melee swing should be used, false otherwise.
func (hunter *Hunter) TryRaptorStrike(sim *core.Simulation) *core.ActiveMeleeAbility {
	if hunter.Rotation.Weave == proto.Hunter_Rotation_WeaveAutosOnly || hunter.IsOnCD(RaptorStrikeCooldownID, sim.CurrentTime) || hunter.CurrentMana() < hunter.raptorStrikeCost {
		return nil
	}

	return hunter.NewRaptorStrike(sim, sim.GetPrimaryTarget())
}
