package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var EnvenomActionID = core.ActionID{SpellID: 32684}

func (rogue *Rogue) makeEnvenom(comboPoints int32) *core.Spell {
	actionID := EnvenomActionID
	actionID.Tag = comboPoints
	refundAmount := 0.4 * float64(rogue.Talents.QuickRecovery)

	baseDamage := 60.0 + (180+core.TernaryFloat64(ItemSetDeathmantle.CharacterHasSetBonus(&rogue.Character, 2), 40, 0))*float64(comboPoints)
	apRatio := 0.03 * float64(comboPoints)

	cost := 35.0
	if ItemSetAssassination.CharacterHasSetBonus(&rogue.Character, 4) {
		cost -= 10
	}

	return rogue.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolNature,
		SpellExtras: core.SpellExtrasMeleeMetrics | core.SpellExtrasIgnoreResists | rogue.finisherFlags(),

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
			DamageMultiplier: 1 + 0.04*float64(rogue.Talents.VilePoisons),
			ThreatMultiplier: 1,
			BaseDamage: core.BaseDamageConfig{
				Calculator: func(sim *core.Simulation, hitEffect *core.SpellEffect, spell *core.Spell) float64 {
					return baseDamage + apRatio*hitEffect.MeleeAttackPower(spell.Unit)
				},
				TargetSpellCoefficient: 0,
			},
			OutcomeApplier: rogue.OutcomeFuncMeleeSpecialHitAndCrit(rogue.critMultiplier(true, false)),
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

func (rogue *Rogue) registerEnvenom() {
	rogue.Envenom = [6]*core.Spell{
		nil,
		rogue.makeEnvenom(1),
		rogue.makeEnvenom(2),
		rogue.makeEnvenom(3),
		rogue.makeEnvenom(4),
		rogue.makeEnvenom(5),
	}
}
