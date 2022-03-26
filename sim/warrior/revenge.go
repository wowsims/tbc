package warrior

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var RevengeCooldownID = core.NewCooldownID()
var RevengeActionID = core.ActionID{SpellID: 30357, CooldownID: RevengeCooldownID}

const RevengeCost = 5.0

func (warrior *Warrior) newRevengeTemplate(_ *core.Simulation) core.SimpleSpellTemplate {
	ability := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            RevengeActionID,
				Character:           &warrior.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				GCD:                 core.GCDDefault,
				Cooldown:            time.Second * 5,
				IgnoreHaste:         true,
				BaseCost: core.ResourceCost{
					Type:  stats.Rage,
					Value: RevengeCost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Rage,
					Value: RevengeCost,
				},
				CritMultiplier: warrior.critMultiplier(true),
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ProcMask:               core.ProcMaskMeleeMHSpecial,
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
				FlatThreatBonus:        200,
			},
			DirectInput: core.DirectDamageInput{
				MinBaseDamage: 414,
				MaxBaseDamage: 506,
			},
		},
	}

	refundAmount := RevengeCost * 0.8
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		if !spellEffect.Landed() {
			warrior.AddRage(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
		}
	}

	return core.NewSimpleSpellTemplate(ability)
}

func (warrior *Warrior) NewRevenge(_ *core.Simulation, target *core.Target) *core.SimpleSpell {
	rv := &warrior.revenge
	warrior.revengeTemplate.Apply(rv)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	rv.Effect.Target = target

	return rv
}

func (warrior *Warrior) CanRevenge(sim *core.Simulation) bool {
	return warrior.revengeTriggered && warrior.CurrentRage() >= RevengeCost && !warrior.IsOnCD(RevengeCooldownID, sim.CurrentTime)
}
