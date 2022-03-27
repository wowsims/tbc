package warrior

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var WhirlwindCooldownID = core.NewCooldownID()
var WhirlwindActionID = core.ActionID{SpellID: 1680, CooldownID: WhirlwindCooldownID}

func (warrior *Warrior) newWhirlwindTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	warrior.whirlwindCost = 25.0 - float64(warrior.Talents.FocusedRage)

	ability := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            WhirlwindActionID,
				Character:           &warrior.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				GCD:                 core.GCDDefault,
				Cooldown:            time.Second*10 - time.Second*time.Duration(warrior.Talents.ImprovedWhirlwind),
				IgnoreHaste:         true,
				BaseCost: core.ResourceCost{
					Type:  stats.Rage,
					Value: warrior.whirlwindCost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Rage,
					Value: warrior.whirlwindCost,
				},
				CritMultiplier: warrior.critMultiplier(true),
			},
		},
	}

	baseEffect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			ProcMask:               core.ProcMaskMeleeMHSpecial,
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
			ThreatMultiplier:       1,
		},
		WeaponInput: core.WeaponDamageInput{
			Normalized:       true,
			DamageMultiplier: 1,
		},
	}

	numHits := core.MinInt32(4, sim.GetNumTargets())
	effects := make([]core.SpellHitEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)
	}
	ability.Effects = effects

	return core.NewSimpleSpellTemplate(ability)
}

func (warrior *Warrior) NewWhirlwind(_ *core.Simulation, target *core.Target) *core.SimpleSpell {
	ww := &warrior.whirlwind
	warrior.whirlwindTemplate.Apply(ww)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	ww.Effect.Target = target

	return ww
}

func (warrior *Warrior) CanWhirlwind(sim *core.Simulation) bool {
	return warrior.CurrentRage() >= warrior.whirlwindCost && !warrior.IsOnCD(WhirlwindCooldownID, sim.CurrentTime)
}
