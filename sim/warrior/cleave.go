package warrior

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var CleaveActionID = core.ActionID{SpellID: 25231}

func (warrior *Warrior) registerCleaveSpell(sim *core.Simulation) {
	cost := 20.0 - float64(warrior.Talents.FocusedRage)

	flatDamageBonus := 70 * (1 + 0.4*float64(warrior.Talents.ImprovedCleave))
	baseEffect := core.SpellEffect{
		ProcMask: core.ProcMaskMeleeMHAuto | core.ProcMaskMeleeMHSpecial,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		FlatThreatBonus:  125,

		BaseDamage:     core.BaseDamageConfigMeleeWeapon(core.MainHand, false, flatDamageBonus, 1, true),
		OutcomeApplier: warrior.OutcomeFuncMeleeSpecialHitAndCrit(warrior.critMultiplier(true)),
	}

	numHits := core.MinInt32(2, sim.GetNumTargets())
	effects := make([]core.SpellEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)
	}

	warrior.Cleave = warrior.RegisterSpell(core.SpellConfig{
		ActionID:    CleaveActionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasMeleeMetrics,

		ResourceType: stats.Rage,
		BaseCost:     cost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: cost,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDamageMultiple(effects),
	})
}

func (warrior *Warrior) QueueCleave(sim *core.Simulation) {
	if warrior.CurrentRage() < warrior.Cleave.DefaultCast.Cost {
		panic("Not enough rage for Cleave")
	}
	if warrior.heroicStrikeQueued {
		return
	}
	if sim.Log != nil {
		warrior.Log(sim, "Cleave queued.")
	}
	warrior.heroicStrikeQueued = true
	warrior.PseudoStats.DisableDWMissPenalty = true
}

func (warrior *Warrior) DequeueCleave(sim *core.Simulation) {
	warrior.heroicStrikeQueued = false
	warrior.PseudoStats.DisableDWMissPenalty = false
	if sim.Log != nil {
		warrior.Log(sim, "Cleave dequeued.")
	}
}

// Returns true if the regular melee swing should be used, false otherwise.
func (warrior *Warrior) TryCleave(sim *core.Simulation) *core.Spell {
	if !warrior.heroicStrikeQueued {
		return nil
	}

	warrior.DequeueCleave(sim)
	if warrior.CurrentRage() < warrior.Cleave.DefaultCast.Cost {
		return nil
	}

	return warrior.Cleave
}

func (warrior *Warrior) CanCleave(sim *core.Simulation) bool {
	return warrior.CurrentRage() >= warrior.Cleave.DefaultCast.Cost
}
