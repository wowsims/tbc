package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var EviscerateActionID = core.ActionID{SpellID: 26865}

func (rogue *Rogue) makeEviscerate(comboPoints int32) *core.Spell {
	actionID := EviscerateActionID
	actionID.Tag = comboPoints

	baseDamage := 60.0 + (185+core.TernaryFloat64(ItemSetDeathmantle.CharacterHasSetBonus(&rogue.Character, 2), 40, 0))*float64(comboPoints)
	apRatio := 0.03 * float64(comboPoints)

	cost := 35.0
	if ItemSetAssassination.CharacterHasSetBonus(&rogue.Character, 4) {
		cost -= 10
	}
	refundAmount := 0.4 * float64(rogue.Talents.QuickRecovery)

	return rogue.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasMeleeMetrics | rogue.finisherFlags(),

		ResourceType: stats.Energy,
		BaseCost:     cost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: cost,
				GCD:  time.Second,
			},
			ModifyCast:  rogue.applyDeathmantle,
			IgnoreHaste: true,
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask:         core.ProcMaskMeleeMHSpecial,
			DamageMultiplier: 1 + 0.05*float64(rogue.Talents.ImprovedEviscerate) + 0.02*float64(rogue.Talents.Aggression),
			ThreatMultiplier: 1,
			BaseDamage: core.BaseDamageConfig{
				Calculator: func(sim *core.Simulation, hitEffect *core.SpellEffect, spell *core.Spell) float64 {
					roll := sim.RandomFloat("Eviscerate") * 120.0
					return baseDamage + roll + hitEffect.MeleeAttackPower(spell.Character)*apRatio + hitEffect.BonusWeaponDamage(spell.Character)
				},
				TargetSpellCoefficient: 1,
			},
			OutcomeApplier: core.OutcomeFuncMeleeSpecialHitAndCrit(rogue.critMultiplier(true, false)),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					rogue.ApplyFinisher(sim, spell.ActionID)
				} else {
					if refundAmount > 0 {
						rogue.AddEnergy(sim, spell.CurCast.Cost*refundAmount, core.ActionID{SpellID: 31245})
					}
				}
			},
		}),
	})
}

func (rogue *Rogue) registerEviscerate() {
	rogue.Eviscerate = [6]*core.Spell{
		nil,
		rogue.makeEviscerate(1),
		rogue.makeEviscerate(2),
		rogue.makeEviscerate(3),
		rogue.makeEviscerate(4),
		rogue.makeEviscerate(5),
	}
}
