package warrior

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var WhirlwindCooldownID = core.NewCooldownID()
var WhirlwindActionID = core.ActionID{SpellID: 1680, CooldownID: WhirlwindCooldownID}

func (warrior *Warrior) registerWhirlwindSpell(sim *core.Simulation) {
	warrior.whirlwindCost = 25.0 - float64(warrior.Talents.FocusedRage)
	if ItemSetWarbringerBattlegear.CharacterHasSetBonus(&warrior.Character, 2) {
		warrior.whirlwindCost -= 5
	}

	ability := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    WhirlwindActionID,
				Character:   &warrior.Character,
				SpellSchool: core.SpellSchoolPhysical,
				GCD:         core.GCDDefault,
				Cooldown:    time.Second*10 - time.Second*time.Duration(warrior.Talents.ImprovedWhirlwind),
				IgnoreHaste: true,
				BaseCost: core.ResourceCost{
					Type:  stats.Rage,
					Value: warrior.whirlwindCost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Rage,
					Value: warrior.whirlwindCost,
				},
				SpellExtras: core.SpellExtrasMeleeMetrics,
			},
		},
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
		Template:     ability,
		ApplyEffects: core.ApplyEffectFuncDamageMultiple(effects),
	})
}

func (warrior *Warrior) CanWhirlwind(sim *core.Simulation) bool {
	return warrior.CurrentRage() >= warrior.whirlwindCost && !warrior.IsOnCD(WhirlwindCooldownID, sim.CurrentTime)
}
