package warrior

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var DemoralizingShoutActionID = core.ActionID{SpellID: 25203}

func (warrior *Warrior) registerDemoralizingShoutSpell(sim *core.Simulation) {
	warrior.shoutCost = 10.0
	if ItemSetBoldArmor.CharacterHasSetBonus(&warrior.Character, 2) {
		warrior.shoutCost -= 2
	}
	cost := warrior.shoutCost - float64(warrior.Talents.FocusedRage)

	baseEffect := core.SpellEffect{
		ThreatMultiplier: 1,
		FlatThreatBonus:  56,
		OutcomeApplier:   core.OutcomeFuncMagicHit(),
	}

	numHits := sim.GetNumTargets()
	effects := make([]core.SpellEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)

		demoShoutAura := core.DemoralizingShoutAura(effects[i].Target, warrior.Talents.BoomingVoice, warrior.Talents.ImprovedDemoralizingShout)
		if i == 0 {
			warrior.DemoralizingShoutAura = demoShoutAura
		}

		effects[i].OnSpellHit = func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
			if spellEffect.Landed() {
				demoShoutAura.Activate(sim)
			}
		}
	}

	warrior.DemoralizingShout = warrior.RegisterSpell(core.SpellConfig{
		ActionID:    DemoralizingShoutActionID,
		SpellSchool: core.SpellSchoolPhysical,

		ResourceType: stats.Rage,
		BaseCost:     cost,

		Cast: core.CastConfig{
			DefaultCast: core.NewCast{
				Cost: cost,
				GCD:  core.GCDDefault,
			},
			IgnoreHaste: true,
		},

		ApplyEffects: core.ApplyEffectFuncDamageMultiple(effects),
	})
}

func (warrior *Warrior) CanDemoralizingShout(sim *core.Simulation) bool {
	return warrior.CurrentRage() >= warrior.DemoralizingShout.DefaultCast.Cost
}
