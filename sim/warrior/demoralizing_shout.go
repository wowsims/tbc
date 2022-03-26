package warrior

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var DemoralizingShoutActionID = core.ActionID{SpellID: 25203}

const DemoralizingShoutCost = 10.0

func (warrior *Warrior) newDemoralizingShoutTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	ability := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            DemoralizingShoutActionID,
				Character:           &warrior.Character,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				SpellSchool:         core.SpellSchoolPhysical,
				GCD:                 core.GCDDefault,
				IgnoreHaste:         true,
				BaseCost: core.ResourceCost{
					Type:  stats.Rage,
					Value: DemoralizingShoutCost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Rage,
					Value: DemoralizingShoutCost,
				},
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ThreatMultiplier: 1,
				FlatThreatBonus:  56,
			},
		},
	}

	baseEffect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			ThreatMultiplier: 1,
			FlatThreatBonus:  56,
		},
	}

	numHits := sim.GetNumTargets()
	effects := make([]core.SpellHitEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)

		demoShoutAura := core.DemoralizingShoutAura(effects[i].Target, warrior.Talents.BoomingVoice, warrior.Talents.ImprovedDemoralizingShout)
		effects[i].OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
			if spellEffect.Landed() {
				// This needs to be AddAura instead of ReplaceAura, in case a lower rank of demo shout was applied by another warrior.
				spellEffect.Target.AddAura(sim, demoShoutAura)
			}
		}
	}
	ability.Effects = effects

	return core.NewSimpleSpellTemplate(ability)
}

func (warrior *Warrior) NewDemoralizingShout(_ *core.Simulation) *core.SimpleSpell {
	ds := &warrior.demoralizingShout
	warrior.demoralizingShoutTemplate.Apply(ds)

	return ds
}

func (warrior *Warrior) CanDemoralizingShout(sim *core.Simulation) bool {
	return warrior.CurrentRage() >= DemoralizingShoutCost
}
