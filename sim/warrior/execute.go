package warrior

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (warrior *Warrior) registerExecuteSpell() {
	cost := 15.0 - float64(warrior.Talents.FocusedRage)
	if warrior.Talents.ImprovedExecute == 1 {
		cost -= 2
	} else if warrior.Talents.ImprovedExecute == 2 {
		cost -= 5
	}
	if ItemSetOnslaughtBattlegear.CharacterHasSetBonus(&warrior.Character, 2) {
		cost -= 3
	}
	refundAmount := cost * 0.8

	var extraRage float64

	warrior.Execute = warrior.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 25236},
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
			ModifyCast: func(_ *core.Simulation, spell *core.Spell, cast *core.Cast) {
				cast.Cost = spell.Unit.CurrentRage()
				extraRage = cast.Cost - spell.BaseCost
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask: core.ProcMaskMeleeMHSpecial,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			BaseDamage: core.BaseDamageConfig{
				Calculator: func(sim *core.Simulation, hitEffect *core.SpellEffect, spell *core.Spell) float64 {
					return 925 + 21*extraRage
				},
				TargetSpellCoefficient: 0, // Doesn't scale with +damage on target?
			},
			OutcomeApplier: warrior.OutcomeFuncMeleeSpecialHitAndCrit(warrior.critMultiplier(true)),

			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					warrior.AddRage(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
				}
			},
		}),
	})
}

func (warrior *Warrior) CanExecute() bool {
	return warrior.CurrentRage() >= warrior.Execute.BaseCost
}
