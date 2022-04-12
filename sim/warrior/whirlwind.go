package warrior

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var WhirlwindCooldownID = core.NewCooldownID()
var WhirlwindActionID = core.ActionID{SpellID: 1680, CooldownID: WhirlwindCooldownID}

func (warrior *Warrior) registerWhirlwindSpell(sim *core.Simulation) {
	cost := 25.0 - float64(warrior.Talents.FocusedRage)
	if ItemSetWarbringerBattlegear.CharacterHasSetBonus(&warrior.Character, 2) {
		cost -= 5
	}

	baseEffect := core.SpellEffect{
		ProcMask: core.ProcMaskMeleeMHSpecial,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		BaseDamage:     core.BaseDamageConfigMeleeWeapon(core.MainHand, true, 0, 1, true),
		OutcomeApplier: core.OutcomeFuncMeleeSpecialHitAndCrit(warrior.critMultiplier(true)),
	}

	numHits := core.MinInt32(4, sim.GetNumTargets())
	effects := make([]core.SpellEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)
	}

	warrior.Whirlwind = warrior.RegisterSpell(core.SpellConfig{
		ActionID:    WhirlwindActionID,
		SpellSchool: core.SpellSchoolPhysical,
		SpellExtras: core.SpellExtrasMeleeMetrics,

		ResourceType: stats.Rage,
		BaseCost:     cost,

		Cast: core.CastConfig{
			DefaultCast: core.NewCast{
				Cost: cost,
				GCD:  core.GCDDefault,
			},
			IgnoreHaste: true,
			Cooldown:    time.Second*10 - time.Second*time.Duration(warrior.Talents.ImprovedWhirlwind),
		},

		ApplyEffects: core.ApplyEffectFuncDamageMultiple(effects),
	})
}

func (warrior *Warrior) CanWhirlwind(sim *core.Simulation) bool {
	return warrior.CurrentRage() >= warrior.Whirlwind.DefaultCast.Cost && !warrior.IsOnCD(WhirlwindCooldownID, sim.CurrentTime)
}
