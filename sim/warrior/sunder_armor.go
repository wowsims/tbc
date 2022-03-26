package warrior

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var SunderArmorActionID = core.ActionID{SpellID: 25225}

const SunderArmorCost = 15.0

func (warrior *Warrior) newSunderArmorTemplate(_ *core.Simulation) core.SimpleSpellTemplate {
	ability := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            SunderArmorActionID,
				Character:           &warrior.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				SpellSchool:         core.SpellSchoolPhysical,
				GCD:                 core.GCDDefault,
				IgnoreHaste:         true,
				BaseCost: core.ResourceCost{
					Type:  stats.Rage,
					Value: SunderArmorCost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Rage,
					Value: SunderArmorCost,
				},
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ProcMask:         core.ProcMaskMeleeMHSpecial,
				ThreatMultiplier: 1,
				FlatThreatBonus:  301.5,
			},
		},
	}

	refundAmount := SunderArmorCost * 0.8
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		if spellEffect.Landed() {
			target := spellEffect.Target
			// Don't overwrite permanent version of SA
			if target.AuraExpiresAt(core.SunderArmorDebuffID) != core.NeverExpires {
				curStacks := target.NumStacks(core.SunderArmorDebuffID)
				newStacks := core.MinInt32(curStacks+1, 5)
				target.ReplaceAura(sim, core.SunderArmorAura(target, newStacks))
			}
		} else {
			warrior.AddRage(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
		}
	}

	return core.NewSimpleSpellTemplate(ability)
}

func (warrior *Warrior) NewSunderArmor(_ *core.Simulation, target *core.Target) *core.SimpleSpell {
	sa := &warrior.sunderArmor
	warrior.sunderArmorTemplate.Apply(sa)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	sa.Effect.Target = target

	return sa
}

func (warrior *Warrior) CanSunderArmor(sim *core.Simulation, target *core.Target) bool {
	return warrior.CurrentRage() >= SunderArmorCost && !target.HasAura(core.ExposeArmorDebuffID)
}
