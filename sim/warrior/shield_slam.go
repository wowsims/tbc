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
	refundAmount := warrior.shieldSlamCost * 0.8

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
				SpellExtras: core.SpellExtrasMeleeMetrics,
			},
		},
	}

	damageRollFunc := core.DamageRollFunc(420, 440)

	warrior.ShieldSlam = warrior.RegisterSpell(core.SpellConfig{
		Template: ability,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask: core.ProcMaskMeleeMHSpecial, // TODO: Is this right?

			DamageMultiplier: 1 * core.TernaryFloat64(ItemSetOnslaughtArmor.CharacterHasSetBonus(&warrior.Character, 4), 1.1, 1),
			ThreatMultiplier: 1,
			FlatThreatBonus:  305,

			BaseDamage: core.BaseDamageConfig{
				Calculator: func(sim *core.Simulation, _ *core.SpellEffect, _ *core.Spell) float64 {
					return damageRollFunc(sim) + warrior.GetStat(stats.BlockValue)
				},
				TargetSpellCoefficient: 1,
			},
			OutcomeApplier: core.OutcomeFuncMeleeSpecialHitAndCrit(warrior.critMultiplier(true)),

			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					warrior.AddRage(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
				}
			},
		}),
	})
}

func (warrior *Warrior) CanShieldSlam(sim *core.Simulation) bool {
	return warrior.canShieldSlam && warrior.CurrentRage() >= warrior.shieldSlamCost && !warrior.IsOnCD(ShieldSlamCooldownID, sim.CurrentTime)
}
