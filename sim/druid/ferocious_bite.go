package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/items"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (druid *Druid) registerFerociousBiteSpell() {
	actionID := core.ActionID{SpellID: 24248}
	baseCost := 35.0

	dmgPerComboPoint := 169.0
	if druid.Equip[items.ItemSlotRanged].ID == 25667 { // Idol of the Beast
		dmgPerComboPoint += 14
	}

	druid.FerociousBite = druid.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasMeleeMetrics,

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
				0.03*float64(druid.Talents.FeralAggression) +
				core.TernaryFloat64(ItemSetThunderheartHarness.CharacterHasSetBonus(&druid.Character, 4), 0.15, 0),
			ThreatMultiplier: 1,

			BaseDamage: core.BaseDamageConfig{
				Calculator: func(sim *core.Simulation, hitEffect *core.SpellEffect, spell *core.Spell) float64 {
					comboPoints := float64(druid.ComboPoints())
					excessEnergy := druid.CurrentEnergy() - spell.DefaultCast.Cost
					base := 57.0 + dmgPerComboPoint*comboPoints + 4.1*excessEnergy
					roll := sim.RandomFloat("Ferocious Bite") * 66.0
					return base + roll + hitEffect.MeleeAttackPower(spell.Unit)*0.05*comboPoints
				},
				TargetSpellCoefficient: 1,
			},
			OutcomeApplier: druid.OutcomeFuncMeleeSpecialHitAndCrit(druid.MeleeCritMultiplier()),

			OnSpellHitDealt: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					druid.SpendComboPoints(sim, spell.ActionID)
				}
			},
		}),
	})
}
