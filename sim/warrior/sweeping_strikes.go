package warrior

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (warrior *Warrior) registerSweepingStrikesCD() {
	if !warrior.Talents.SweepingStrikes {
		return
	}

	actionID := core.ActionID{SpellID: 12328}

	var curDmg float64
	ssHit := warrior.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasMeleeMetrics | core.SpellExtrasNoOnCastComplete,

		ApplyEffects: core.ApplyEffectFuncDirectDamageTargetModifiersOnly(core.SpellEffect{
			// No proc mask, so it won't proc itself.

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			BaseDamage: core.BaseDamageConfig{
				Calculator: func(_ *core.Simulation, _ *core.SpellEffect, _ *core.Spell) float64 {
					return curDmg
				},
			},
			OutcomeApplier: warrior.OutcomeFuncAlwaysHit(),
		}),
	})

	ssAura := warrior.RegisterAura(core.Aura{
		Label:     "Sweeping Strikes",
		ActionID:  actionID,
		Duration:  core.NeverExpires,
		MaxStacks: 10,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.SetStacks(sim, 10)
		},
		OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if aura.GetStacks() == 0 || spellEffect.Damage == 0 || !spellEffect.ProcMask.Matches(core.ProcMaskMelee) {
				return
			}

			// TODO: If the triggering spell is Execute and 2nd target health > 20%, do a normalized MH hit instead.

			// Undo armor reduction to get the raw damage value.
			curDmg = spellEffect.Damage / warrior.AttackTables[spellEffect.Target.Index].ArmorDamageReduction

			ssHit.Cast(sim, spellEffect.Target.NextTarget(sim))
			ssHit.SpellMetrics[spellEffect.Target.Index].Casts--
			aura.RemoveStack(sim)
		},
	})

	cost := 30.0
	ssCD := warrior.RegisterSpell(core.SpellConfig{
		ActionID:    actionID,
		SpellSchool: core.SpellSchoolPhysical,

		ResourceType: stats.Rage,
		BaseCost:     cost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: cost,
			},
			CD: core.Cooldown{
				Timer:    warrior.NewTimer(),
				Duration: time.Second * 30,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Target, spell *core.Spell) {
			ssAura.Activate(sim)
		},
	})

	warrior.AddMajorCooldown(core.MajorCooldown{
		Spell: ssCD,
		Type:  core.CooldownTypeDPS,
		CanActivate: func(sim *core.Simulation, character *core.Character) bool {
			return sim.GetNumTargets() > 1 && warrior.CurrentRage() >= ssCD.DefaultCast.Cost
		},
		ShouldActivate: func(sim *core.Simulation, character *core.Character) bool {
			return true
		},
	})
}
