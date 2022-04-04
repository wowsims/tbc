package warrior

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var ShieldSlamCooldownID = core.NewCooldownID()
var ShieldSlamActionID = core.ActionID{SpellID: 30356, CooldownID: ShieldSlamCooldownID}

func (warrior *Warrior) registerShieldSlamSpell(_ *core.Simulation) {
	warrior.shieldSlamCost = 20.0 - float64(warrior.Talents.FocusedRage)
	warrior.canShieldSlam = warrior.Talents.ShieldSlam && warrior.Equip[proto.ItemSlot_ItemSlotOffHand].WeaponType == proto.WeaponType_WeaponTypeShield

	ability := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    ShieldSlamActionID,
				Character:   &warrior.Character,
				SpellSchool: core.SpellSchoolPhysical,
				GCD:         core.GCDDefault,
				Cooldown:    time.Second * 6,
				IgnoreHaste: true,
				BaseCost: core.ResourceCost{
					Type:  stats.Rage,
					Value: warrior.shieldSlamCost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Rage,
					Value: warrior.shieldSlamCost,
				},
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategorySpecial,
			CritRollCategory:    core.CritRollCategoryPhysical,
			CritMultiplier:      warrior.critMultiplier(true),
			ProcMask:            core.ProcMaskMeleeMHSpecial, // TODO: Is this right?
			DamageMultiplier:    1,
			ThreatMultiplier:    1,
			FlatThreatBonus:     305,
		},
	}

	damageRollFunc := core.DamageRollFunc(420, 440)
	ability.Effect.BaseDamage = core.BaseDamageConfig{
		Calculator: func(sim *core.Simulation, _ *core.SpellEffect, _ *core.Spell) float64 {
			return damageRollFunc(sim) + warrior.GetStat(stats.BlockValue)
		},
		TargetSpellCoefficient: 1,
	}

	if ItemSetOnslaughtArmor.CharacterHasSetBonus(&warrior.Character, 4) {
		ability.Effect.DamageMultiplier *= 1.1
	}

	refundAmount := warrior.shieldSlamCost * 0.8
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
		if !spellEffect.Landed() {
			warrior.AddRage(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
		}
	}

	warrior.ShieldSlam = warrior.RegisterSpell(core.SpellConfig{
		Template:   ability,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}

func (warrior *Warrior) CanShieldSlam(sim *core.Simulation) bool {
	return warrior.canShieldSlam && warrior.CurrentRage() >= warrior.shieldSlamCost && !warrior.IsOnCD(ShieldSlamCooldownID, sim.CurrentTime)
}
