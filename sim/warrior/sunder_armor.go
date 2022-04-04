package warrior

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var SunderArmorActionID = core.ActionID{SpellID: 25225}

func (warrior *Warrior) registerSunderArmorSpell(sim *core.Simulation) {
	warrior.SunderArmorAura = core.SunderArmorAura(sim.GetPrimaryTarget(), 0)
	warrior.ExposeArmorAura = core.ExposeArmorAura(sim.GetPrimaryTarget(), 2)

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
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategorySpecial,
			ProcMask:            core.ProcMaskMeleeMHSpecial,
			ThreatMultiplier:    1,
			FlatThreatBonus:     301.5,
		},
	}

	refundAmount := warrior.sunderArmorCost * 0.8
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
		if spellEffect.Landed() {
			warrior.SunderArmorAura.Activate(sim)
			warrior.SunderArmorAura.AddStack(sim)
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
	return warrior.CurrentRage() >= warrior.sunderArmorCost && !warrior.ExposeArmorAura.IsActive()
}
