package warrior

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var ThunderClapCooldownID = core.NewCooldownID()
var ThunderClapActionID = core.ActionID{SpellID: 25264, CooldownID: ThunderClapCooldownID}

const ThunderClapCost = 20.0

func (warrior *Warrior) newThunderClapTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	ability := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            ThunderClapActionID,
				Character:           &warrior.Character,
				OutcomeRollCategory: core.OutcomeRollCategoryMagic,
				CritRollCategory:    core.CritRollCategoryMagical,
				SpellSchool:         core.SpellSchoolPhysical,
				SpellExtras:         core.SpellExtrasBinary,
				GCD:                 core.GCDDefault,
				Cooldown:            time.Second * 4,
				IgnoreHaste:         true,
				BaseCost: core.ResourceCost{
					Type:  stats.Rage,
					Value: ThunderClapCost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Rage,
					Value: ThunderClapCost,
				},
				CritMultiplier: warrior.spellCritMultiplier(true),
			},
		},
	}

	baseEffect := core.SpellHitEffect{
		SpellEffect: core.SpellEffect{
			DamageMultiplier:       1,
			StaticDamageMultiplier: 1,
			ThreatMultiplier:       1.75,
		},
		DirectInput: core.DirectDamageInput{
			MinBaseDamage: 123,
			MaxBaseDamage: 123,
		},
	}

	numHits := core.MinInt32(4, sim.GetNumTargets())
	effects := make([]core.SpellHitEffect, 0, numHits)
	for i := int32(0); i < numHits; i++ {
		effects = append(effects, baseEffect)
		effects[i].Target = sim.GetTarget(i)

		tcAura := core.ThunderClapAura(effects[i].Target, warrior.Talents.ImprovedThunderClap)
		effects[i].OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
			if spellEffect.Landed() {
				// This needs to be AddAura instead of ReplaceAura, in case a lower rank of Thunder Clap was applied by another warrior.
				spellEffect.Target.AddAura(sim, tcAura)
			}
		}
	}
	ability.Effects = effects

	return core.NewSimpleSpellTemplate(ability)
}

func (warrior *Warrior) NewThunderClap(_ *core.Simulation) *core.SimpleSpell {
	tc := &warrior.thunderClap
	warrior.thunderClapTemplate.Apply(tc)
	return tc
}

func (warrior *Warrior) CanThunderClap(sim *core.Simulation) bool {
	return warrior.CurrentRage() >= ThunderClapCost && !warrior.IsOnCD(ThunderClapCooldownID, sim.CurrentTime)
}
