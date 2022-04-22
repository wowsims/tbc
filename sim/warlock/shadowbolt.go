package warlock

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDSB11 int32 = 27209

var ShadowBolt11ActionID = core.ActionID{SpellID: SpellIDSB11}

func (warlock *Warlock) registerShadowboltSpell(sim *core.Simulation) {
	baseCost := 420.0
	minBaseDamage := 544.0
	maxBaseDamage := 607.0
	bonusFlatDamage := 0.0
	spellCoefficient := 0.857

	debuffAura := warlock.impShadowboltDebuffAura(sim.GetPrimaryTarget())

	has4pMal := ItemSetMaleficRaiment.CharacterHasSetBonus(&warlock.Character, 4)

	effect := core.SpellEffect{
		BonusSpellCritRating: float64(warlock.Talents.Devastation)*1*core.SpellCritRatingPerCritChance +
			float64(warlock.Talents.Backlash)*1*core.SpellCritRatingPerCritChance,
		DamageMultiplier: 1 * core.TernaryFloat64(has4pMal, 1.06, 1.0),
		ThreatMultiplier: 1 - 0.05*float64(warlock.Talents.DestructiveReach),
		// TODO: so this would mean SB ratio is 1.057?
		BaseDamage:     core.BaseDamageConfigMagic(minBaseDamage+bonusFlatDamage, maxBaseDamage+bonusFlatDamage, spellCoefficient+0.04*float64(warlock.Talents.ShadowAndFlame)),
		OutcomeApplier: core.OutcomeFuncMagicHitAndCrit(warlock.SpellCritMultiplier(1, core.TernaryFloat64(warlock.Talents.Ruin, 1, 0))),
	}
	if warlock.Talents.ImprovedShadowBolt > 0 {
		effect.OnSpellHit = func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if !spellEffect.Landed() || !spellEffect.Outcome.Matches(core.OutcomeCrit) {
				return
			}
			debuffAura.Activate(sim)
			debuffAura.SetStacks(sim, 4)
		}
	}

	warlock.Shadowbolt = warlock.RegisterSpell(core.SpellConfig{
		ActionID:    ShadowBolt11ActionID,
		SpellSchool: core.SpellSchoolShadow,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost:     baseCost * (1 - 0.01*float64(warlock.Talents.Cataclysm)),
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond*3000 - (time.Millisecond * 100 * time.Duration(warlock.Talents.Bane)),
			},
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
			if !spellEffect.Landed() || spellEffect.Damage == 0 {
				return
			}
			aura.RemoveStack(sim)
		},
	})
}
