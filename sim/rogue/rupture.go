package rogue

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
)

var RuptureActionID = core.ActionID{SpellID: 26867}
var RuptureDebuffID = core.NewDebuffID()
var RuptureEnergyCost = 25.0

func (rogue *Rogue) registerRuptureSpell(sim *core.Simulation) {
	refundAmount := 0.4 * float64(rogue.Talents.QuickRecovery)

	ability := rogue.newAbility(RuptureActionID, RuptureEnergyCost, SpellFlagFinisher|core.SpellExtrasIgnoreResists, core.ProcMaskMeleeMHSpecial)
	ability.Effect.CritRollCategory = core.CritRollCategoryNone
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		if spellEffect.Landed() {
			rogue.ApplyFinisher(sim, spellCast.ActionID)
		} else {
			if refundAmount > 0 {
				rogue.AddEnergy(sim, spellCast.Cost.Value*refundAmount, core.ActionID{SpellID: 31245})
			}
		}
	}
	ability.Effect.DotInput = core.DotDamageInput{
		NumberOfTicks: 0, // Set dynamically.
		TickLength:    time.Second * 2,
		TickBaseDamage: func(sim *core.Simulation, hitEffect *core.SpellEffect, spell *core.SimpleSpellTemplate) float64 {
			comboPoints := rogue.ComboPoints()
			attackPower := hitEffect.MeleeAttackPower(spell.Character) + hitEffect.MeleeAttackPowerOnTarget()

			return 70 + float64(comboPoints)*11 + attackPower*[]float64{0.01, 0.02, 0.03, 0.03, 0.03}[comboPoints-1]
		},
		DebuffID: RuptureDebuffID,
	}

	ability.Effect.DamageMultiplier += 0.1 * float64(rogue.Talents.SerratedBlades)
	if rogue.Talents.SurpriseAttacks {
		ability.SpellExtras |= core.SpellExtrasCannotBeDodged
	}

	rogue.Rupture = rogue.RegisterSpell(core.SpellConfig{
		Template: ability,
		ModifyCast: func(sim *core.Simulation, target *core.Target, instance *core.SimpleSpell) {
			instance.ActionID.Tag = rogue.ComboPoints()
			instance.Effect.Target = target
			instance.Effect.DotInput.NumberOfTicks = int(rogue.ComboPoints()) + 3
			if rogue.deathmantle4pcProc {
				instance.Cost.Value = 0
				rogue.deathmantle4pcProc = false
			}
		},
	})
}

func (rogue *Rogue) RuptureDuration(comboPoints int32) time.Duration {
	return time.Second*6 + time.Second*2*time.Duration(comboPoints)
}
