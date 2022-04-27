package warlock

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

func (warlock *Warlock) registerShadowboltSpell(sim *core.Simulation) {
	warlock.ImpShadowboltAura = warlock.impShadowboltDebuffAura(sim.GetPrimaryTarget())
	has4pMal := ItemSetMaleficRaiment.CharacterHasSetBonus(&warlock.Character, 4)

	effect := core.SpellEffect{
		ProcMask: core.ProcMaskSpellDamage,
		BonusSpellCritRating: float64(warlock.Talents.Devastation)*1*core.SpellCritRatingPerCritChance +
			float64(warlock.Talents.Backlash)*1*core.SpellCritRatingPerCritChance,
		DamageMultiplier: 1 * core.TernaryFloat64(has4pMal, 1.06, 1.0) * (1 + 0.02*float64(warlock.Talents.ShadowMastery)),
		ThreatMultiplier: 1 - 0.05*float64(warlock.Talents.DestructiveReach),
		BaseDamage:       core.BaseDamageConfigMagic(544.0, 607.0, 0.857+0.04*float64(warlock.Talents.ShadowAndFlame)),
		OutcomeApplier:   warlock.OutcomeFuncMagicHitAndCrit(warlock.SpellCritMultiplier(1, core.TernaryFloat64(warlock.Talents.Ruin, 1, 0))),
	}
	if warlock.Talents.ImprovedShadowBolt > 0 {
		effect.OnSpellHit = func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.Outcome.Matches(core.OutcomeCrit) {
				return
			}
			warlock.ImpShadowboltAura.Activate(sim)
			warlock.ImpShadowboltAura.SetStacks(sim, 4)
		}
	}

	var modCast func(*core.Simulation, *core.Spell, *core.Cast)

	if warlock.Talents.Nightfall > 0 {
		modCast = func(_ *core.Simulation, _ *core.Spell, cast *core.Cast) {
			warlock.applyNightfall(cast)
		}
	}

	baseCost := 420.0
	warlock.Shadowbolt = warlock.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 27209},
		SpellSchool: core.SpellSchoolShadow,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost:     baseCost * (1 - 0.01*float64(warlock.Talents.Cataclysm)),
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond*3000 - (time.Millisecond * 100 * time.Duration(warlock.Talents.Bane)),
			},
			ModifyCast: modCast,
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(effect),
	})

}

func (warlock *Warlock) impShadowboltDebuffAura(target *core.Target) *core.Aura {
	bonus := 1 + 0.04*float64(warlock.Talents.ImprovedShadowBolt)
	return target.GetOrRegisterAura(core.Aura{
		Label:     "Improved Shadow Bolt",
		ActionID:  core.ActionID{SpellID: 17803},
		Duration:  time.Second * 12,
		MaxStacks: 4,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			target.PseudoStats.ShadowDamageTakenMultiplier *= bonus
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			target.PseudoStats.ShadowDamageTakenMultiplier /= bonus
		},
		OnSpellHit: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spell.SpellSchool != core.SpellSchoolShadow {
				return
			}
			if !spellEffect.Landed() || spellEffect.Damage == 0 || spellEffect.IsPhantom || spellEffect.ProcMask == 0 {
				return
			}
			aura.RemoveStack(sim)
		},
	})
}
