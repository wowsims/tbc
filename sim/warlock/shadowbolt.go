package warlock

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

// Starfire spell IDs
const SpellIDSB11 int32 = 27209

var ShadowBolt11ActionID = core.ActionID{SpellID: SpellIDSB11}

// Shadow Bolt	Rank 11
// 420 Mana	30 yd range
// 3 sec cast
// Requires Warlock
// Requires level 69
// Sends a shadowy bolt at the enemy, causing 544 to 607 Shadow damage.

func (warlock *Warlock) newShadowboltSpell(sim *core.Simulation) *core.Spell {
	baseCost := 420.0
	// minBaseDamage := 544.0
	// maxBaseDamage := 607.0
	// spellCoefficient := 0.857

	// // This seems to be unaffected by wrath of cenarius so it needs to come first.
	// bonusFlatDamage := core.TernaryFloat64(druid.Equip[items.ItemSlotRanged].ID == IvoryMoongoddess, 55*spellCoefficient, 0)
	// spellCoefficient += 0.04 * float64(druid.Talents.WrathOfCenarius)

	effect := core.SpellEffect{
		// BonusSpellCritRating: (float64(druid.Talents.FocusedStarlight) * 2 * core.SpellCritRatingPerCritChance) + core.TernaryFloat64(ItemSetThunderheart.CharacterHasSetBonus(&druid.Character, 4), 5*core.SpellCritRatingPerCritChance, 0),
		// DamageMultiplier:     1 + 0.02*float64(druid.Talents.Moonfury),
		// ThreatMultiplier:     1,
		// BaseDamage:           core.BaseDamageConfigMagic(minBaseDamage+bonusFlatDamage, maxBaseDamage+bonusFlatDamage, spellCoefficient),
		// OutcomeApplier:       core.OutcomeFuncMagicHitAndCrit(druid.SpellCritMultiplier(1, 0.2*float64(druid.Talents.Vengeance))),
	}

	// if ItemSetNordrassil.CharacterHasSetBonus(&druid.Character, 4) {
	// 	effect.BaseDamage = core.WrapBaseDamageConfig(effect.BaseDamage, func(oldCalculator core.BaseDamageCalculator) core.BaseDamageCalculator {
	// 		return func(sim *core.Simulation, hitEffect *core.SpellEffect, spell *core.Spell) float64 {
	// 			normalDamage := oldCalculator(sim, hitEffect, spell)

	// 			// Check if moonfire/insectswarm is ticking on the target.
	// 			// TODO: in a raid simulator we need to be able to see which dots are ticking from other druids.
	// 			if druid.MoonfireDot.IsActive() || druid.InsectSwarmDot.IsActive() {
	// 				return normalDamage * 1.1
	// 			} else {
	// 				return normalDamage
	// 			}
	// 		}
	// 	})
	// }

	return warlock.RegisterSpell(core.SpellConfig{
		ActionID:    ShadowBolt11ActionID,
		SpellSchool: core.SpellSchoolShadow,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost:     baseCost, // * (1 - 0.03*float64(druid.Talents.Moonglow)),
				GCD:      core.GCDDefault,
				CastTime: time.Millisecond * 3000, //*3500 - (time.Millisecond * 100 * time.Duration(druid.Talents.StarlightWrath)),
			},

			// ModifyCast: func(_ *core.Simulation, _ *core.Spell, cast *core.Cast) {
			// 	druid.applyNaturesGrace(cast)
			// 	druid.applyNaturesSwiftness(cast)
			// },
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(effect),
	})
}

func (warlock *Warlock) impShadowboltDebuffAura(target *core.Target) *core.Aura {
	return target.GetOrRegisterAura(core.Aura{
		Label:     "Improved Shadow Bolt",
		ActionID:  ShadowBolt11ActionID,
		Duration:  time.Second * 10,
		MaxStacks: 4,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			target.PseudoStats.ShadowDamageTakenMultiplier *= 1.2
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			target.PseudoStats.ShadowDamageTakenMultiplier /= 1.2
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
