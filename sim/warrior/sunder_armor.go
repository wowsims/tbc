package warrior

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var SunderArmorActionID = core.ActionID{SpellID: 25225}

func (warrior *Warrior) registerSunderArmorSpell(_ *core.Simulation) {
	warrior.sunderArmorCost = 15.0 - float64(warrior.Talents.ImprovedSunderArmor) - float64(warrior.Talents.FocusedRage)

	ability := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    SunderArmorActionID,
				Character:   &warrior.Character,
				SpellSchool: core.SpellSchoolPhysical,
				GCD:         core.GCDDefault,
				IgnoreHaste: true,
				BaseCost: core.ResourceCost{
					Type:  stats.Rage,
					Value: warrior.sunderArmorCost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Rage,
					Value: warrior.sunderArmorCost,
				},
			},
			OutcomeRollCategory: core.OutcomeRollCategorySpecial,
		},
		Effect: core.SpellEffect{
			ProcMask:         core.ProcMaskMeleeMHSpecial,
			ThreatMultiplier: 1,
			FlatThreatBonus:  301.5,
		},
	}

	refundAmount := warrior.sunderArmorCost * 0.8
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

	warrior.SunderArmor = warrior.RegisterSpell(core.SpellConfig{
		Template:   ability,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}

func (warrior *Warrior) CanSunderArmor(sim *core.Simulation, target *core.Target) bool {
	return warrior.CurrentRage() >= warrior.sunderArmorCost && !target.HasAura(core.ExposeArmorDebuffID)
}
