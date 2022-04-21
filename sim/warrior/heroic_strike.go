package warrior

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var HeroicStrikeActionID = core.ActionID{SpellID: 29707}

func (warrior *Warrior) registerHeroicStrikeSpell(_ *core.Simulation) {
	cost := 15.0 - float64(warrior.Talents.ImprovedHeroicStrike) - float64(warrior.Talents.FocusedRage)
	refundAmount := cost * 0.8

	warrior.HeroicStrike = warrior.RegisterSpell(core.SpellConfig{
		ActionID:    HeroicStrikeActionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasMeleeMetrics,

		ResourceType: stats.Rage,
		BaseCost:     cost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: cost,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			ProcMask: core.ProcMaskMeleeMHAuto | core.ProcMaskMeleeMHSpecial,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			FlatThreatBonus:  194,

			BaseDamage:     core.BaseDamageConfigMeleeWeapon(core.MainHand, false, 176, 1, true),
			OutcomeApplier: core.OutcomeFuncMeleeSpecialHitAndCrit(warrior.critMultiplier(true)),

			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if !spellEffect.Landed() {
					warrior.AddRage(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
				}
			},
		}),
	})
}

func (warrior *Warrior) QueueHeroicStrike(sim *core.Simulation) {
	if warrior.CurrentRage() < warrior.HeroicStrike.DefaultCast.Cost {
		panic("Not enough rage for HS")
	}
	if warrior.heroicStrikeQueued {
		return
	}
	if sim.Log != nil {
		warrior.Log(sim, "Heroic strike queued.")
	}
	warrior.heroicStrikeQueued = true
	warrior.PseudoStats.DisableDWMissPenalty = true
}

// Returns true if the regular melee swing should be used, false otherwise.
func (warrior *Warrior) TryHeroicStrike(sim *core.Simulation) *core.Spell {
	if !warrior.heroicStrikeQueued {
		return nil
	}

	warrior.heroicStrikeQueued = false
	warrior.PseudoStats.DisableDWMissPenalty = false
	if sim.Log != nil {
		warrior.Log(sim, "Heroic strike unqueued.")
	}
	if warrior.CurrentRage() < warrior.HeroicStrike.DefaultCast.Cost {
		return nil
	}

	return warrior.HeroicStrike
}

func (warrior *Warrior) CanHeroicStrike(sim *core.Simulation) bool {
	return warrior.CurrentRage() >= warrior.HeroicStrike.DefaultCast.Cost
}
