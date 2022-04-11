package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var EviscerateActionID = core.ActionID{SpellID: 26865}

func (rogue *Rogue) registerEviscerateSpell(sim *core.Simulation) {
	rogue.eviscerateEnergyCost = 35
	if ItemSetAssassination.CharacterHasSetBonus(&rogue.Character, 4) {
		rogue.eviscerateEnergyCost -= 10
	}
	refundAmount := 0.4 * float64(rogue.Talents.QuickRecovery)

	basePerComboPoint := 185.0
	if ItemSetDeathmantle.CharacterHasSetBonus(&rogue.Character, 2) {
		basePerComboPoint += 40
	}

	rogue.Eviscerate = rogue.RegisterSpell(core.SpellConfig{
		ActionID:    EviscerateActionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasMeleeMetrics | rogue.finisherFlags(),

		ResourceType: stats.Energy,
		BaseCost:     rogue.eviscerateEnergyCost,

		Cast: core.CastConfig{
			DefaultCast: core.NewCast{
				Cost: rogue.eviscerateEnergyCost,
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
					comboPoints := rogue.ComboPoints()
					base := 60.0 + basePerComboPoint*float64(comboPoints)
					roll := sim.RandomFloat("Eviscerate") * 120.0
					return base + roll + (hitEffect.MeleeAttackPower(spell.Character)*0.03)*float64(comboPoints) + hitEffect.BonusWeaponDamage(spell.Character)
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
