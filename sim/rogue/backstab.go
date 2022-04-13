package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var BackstabActionID = core.ActionID{SpellID: 26863}
var BackstabEnergyCost = 60.0

func (rogue *Rogue) registerBackstabSpell(_ *core.Simulation) {
	refundAmount := BackstabEnergyCost * 0.8

	rogue.Backstab = rogue.RegisterSpell(core.SpellConfig{
		ActionID:    BackstabActionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasMeleeMetrics | SpellFlagBuilder,

		ResourceType: stats.Energy,
		BaseCost:     BackstabEnergyCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: BackstabEnergyCost,
				GCD:  time.Second,
			},
			IgnoreHaste: true,
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask:        core.ProcMaskMeleeMHSpecial,
			BonusCritRating: 10 * core.MeleeCritRatingPerCritChance * float64(rogue.Talents.PuncturingWounds),
			// All of these use "Apply Aura: Modifies Damage/Healing Done", and stack additively (up to 142%).
			DamageMultiplier: 1 +
				0.04*float64(rogue.Talents.Opportunity) +
				0.02*float64(rogue.Talents.Aggression) +
				core.TernaryFloat64(rogue.Talents.SurpriseAttacks, 0.1, 0) +
				core.TernaryFloat64(ItemSetSlayers.CharacterHasSetBonus(&rogue.Character, 4), 0.06, 0),
			ThreatMultiplier: 1,
			BaseDamage:       core.BaseDamageConfigMeleeWeapon(core.MainHand, true, 170, 1.5+0.01*float64(rogue.Talents.SinisterCalling), true),
			OutcomeApplier:   core.OutcomeFuncMeleeSpecialHitAndCrit(rogue.critMultiplier(true, true)),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					rogue.AddComboPoints(sim, 1, BackstabActionID)
				} else {
					rogue.AddEnergy(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
				}
			},
		}),
	})
}
