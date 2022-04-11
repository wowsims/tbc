package warrior

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var SunderArmorActionID = core.ActionID{SpellID: 25225}

func (warrior *Warrior) newSunderArmorSpell(sim *core.Simulation, isDevastateEffect bool) *core.Spell {
	warrior.SunderArmorAura = core.SunderArmorAura(sim.GetPrimaryTarget(), 0)
	warrior.ExposeArmorAura = core.ExposeArmorAura(sim.GetPrimaryTarget(), 2)
	refundAmount := warrior.sunderArmorCost * 0.8

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
				SpellExtras: core.SpellExtrasMeleeMetrics,
			},
		},
	}
	if isDevastateEffect {
		ability.Cost.Value = 0
		ability.BaseCost.Value = 0
		ability.GCD = 0
	}

	effect := core.SpellEffect{
		ProcMask: core.ProcMaskMeleeMHSpecial,

		ThreatMultiplier: 1,
		FlatThreatBonus:  301.5,

		OutcomeApplier: core.OutcomeFuncMeleeSpecialHit(),

		OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spellEffect.Landed() {
				warrior.SunderArmorAura.Activate(sim)
				warrior.SunderArmorAura.AddStack(sim)
			} else {
				warrior.AddRage(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
			}
		},
	}
	if isDevastateEffect {
		effect.OutcomeApplier = core.OutcomeFuncAlwaysHit()
		effect.OnInit = func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if warrior.SunderArmorAura.GetStacks() == 5 {
				spellEffect.ThreatMultiplier = 0
			}
		}
	}

	return warrior.RegisterSpell(core.SpellConfig{
		Template:     ability,
		ApplyEffects: core.ApplyEffectFuncDirectDamage(effect),
	})
}

func (warrior *Warrior) CanSunderArmor(sim *core.Simulation, target *core.Target) bool {
	return warrior.CurrentRage() >= warrior.sunderArmorCost && !warrior.ExposeArmorAura.IsActive()
}
