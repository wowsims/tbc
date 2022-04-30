package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (rogue *Rogue) registerHemorrhageSpell() {
	actionID := core.ActionID{SpellID: 26864}

	target := rogue.Env.GetPrimaryTarget()
	hemoAura := target.GetOrRegisterAura(core.Aura{
		Label:     "Hemorrhage",
		ActionID:  actionID,
		Duration:  time.Second * 15,
		MaxStacks: 10,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			target.PseudoStats.BonusPhysicalDamageTaken += 42
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			target.PseudoStats.BonusPhysicalDamageTaken -= 42
		},
		OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spell.SpellSchool != core.SpellSchoolPhysical {
				return
			}
			if !spellEffect.Landed() || spellEffect.Damage == 0 {
				return
			}

			aura.RemoveStack(sim)
		},
	})

	baseCost := 35.0
	refundAmount := baseCost * 0.8

	rogue.Hemorrhage = rogue.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasMeleeMetrics | SpellFlagBuilder,

		ResourceType: stats.Energy,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost,
				GCD:  time.Second,
			},
			IgnoreHaste: true,
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask: core.ProcMaskMeleeMHSpecial,
			DamageMultiplier: 1 +
				core.TernaryFloat64(ItemSetSlayers.CharacterHasSetBonus(&rogue.Character, 4), 0.06, 0),
			ThreatMultiplier: 1,
			BaseDamage:       core.BaseDamageConfigMeleeWeapon(core.MainHand, true, 0, 1.1+0.01*float64(rogue.Talents.SinisterCalling), true),
			OutcomeApplier:   rogue.OutcomeFuncMeleeSpecialHitAndCrit(rogue.critMultiplier(true, true)),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					rogue.AddComboPoints(sim, 1, spell.ActionID)

					hemoAura.Activate(sim)
					hemoAura.SetStacks(sim, 10)
				} else {
					rogue.AddEnergy(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
				}
			},
		}),
	})
}
