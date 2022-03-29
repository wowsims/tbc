package warrior

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

var ThunderClapCooldownID = core.NewCooldownID()
var ThunderClapActionID = core.ActionID{SpellID: 25264, CooldownID: ThunderClapCooldownID}

func (warrior *Warrior) newThunderClapTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	warrior.thunderClapCost = 20.0 - float64(warrior.Talents.FocusedRage)
	impTCDamageMult := 1.0
	if warrior.Talents.ImprovedThunderClap == 1 {
		warrior.thunderClapCost -= 1
		impTCDamageMult = 1.4
	} else if warrior.Talents.ImprovedThunderClap == 2 {
		warrior.thunderClapCost -= 2
		impTCDamageMult = 1.7
	} else if warrior.Talents.ImprovedThunderClap == 3 {
		warrior.thunderClapCost -= 4
		impTCDamageMult = 2
	}

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
					Value: warrior.thunderClapCost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Rage,
					Value: warrior.thunderClapCost,
				},
				CritMultiplier: warrior.spellCritMultiplier(true),
			},
		},
	}

	baseEffect := core.SpellEffect{
		DamageMultiplier: impTCDamageMult,
		ThreatMultiplier: 1.75,
		BaseDamage:       core.BaseDamageConfigFlat(123),
	}

	numHits := core.MinInt32(4, sim.GetNumTargets())
	effects := make([]core.SpellEffect, 0, numHits)
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
	return warrior.StanceMatches(BattleStance|DefensiveStance) && warrior.CurrentRage() >= warrior.thunderClapCost && !warrior.IsOnCD(ThunderClapCooldownID, sim.CurrentTime)
}
