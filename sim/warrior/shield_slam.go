package warrior

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var ShieldSlamActionID = core.ActionID{SpellID: 30356}

func (warrior *Warrior) registerShieldSlamSpell(_ *core.Simulation, cdTimer *core.Timer) {
	cost := 20.0 - float64(warrior.Talents.FocusedRage)
	refundAmount := cost * 0.8
	warrior.canShieldSlam = warrior.Talents.ShieldSlam && warrior.Equip[proto.ItemSlot_ItemSlotOffHand].WeaponType == proto.WeaponType_WeaponTypeShield

	damageRollFunc := core.DamageRollFunc(420, 440)

	warrior.ShieldSlam = warrior.RegisterSpell(core.SpellConfig{
		ActionID:    ShieldSlamActionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasMeleeMetrics,

		ResourceType: stats.Rage,
		BaseCost:     cost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: cost,
				GCD:  core.GCDDefault,
			},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    cdTimer,
				Duration: time.Second * 6,
			},
		},

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
	return warrior.canShieldSlam && warrior.CurrentRage() >= warrior.ShieldSlam.DefaultCast.Cost && warrior.ShieldSlam.IsReady(sim)
}
