package warrior

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var HeroicStrikeActionID = core.ActionID{SpellID: 29707}

func (warrior *Warrior) newHeroicStrikeTemplate(_ *core.Simulation) core.SimpleSpellTemplate {
	warrior.heroicStrikeCost = 15.0 - float64(warrior.Talents.ImprovedHeroicStrike) - float64(warrior.Talents.FocusedRage)

	ability := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    HeroicStrikeActionID,
				Character:   &warrior.Character,
				SpellSchool: core.SpellSchoolPhysical,
				BaseCost: core.ResourceCost{
					Type:  stats.Rage,
					Value: warrior.heroicStrikeCost,
				},
				Cost: core.ResourceCost{
					Type:  stats.Rage,
					Value: warrior.heroicStrikeCost,
				},
			},
			OutcomeRollCategory: core.OutcomeRollCategorySpecial,
			CritRollCategory:    core.CritRollCategoryPhysical,
			CritMultiplier:      warrior.critMultiplier(true),
		},
		Effect: core.SpellEffect{
			ProcMask:         core.ProcMaskMeleeMHAuto | core.ProcMaskMeleeMHSpecial,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			FlatThreatBonus:  194,
			BaseDamage:       core.BaseDamageConfigMeleeWeapon(core.MainHand, false, 176, 1, true),
		},
	}

	refundAmount := warrior.heroicStrikeCost * 0.8
	ability.Effect.OnSpellHit = func(sim *core.Simulation, spellCast *core.SpellCast, spellEffect *core.SpellEffect) {
		if !spellEffect.Landed() {
			warrior.AddRage(sim, refundAmount, core.ActionID{OtherID: proto.OtherAction_OtherActionRefund})
		}
	}

	return core.NewSimpleSpellTemplate(ability)
}

func (warrior *Warrior) QueueHeroicStrike(_ *core.Simulation) {
	if warrior.CurrentRage() < warrior.heroicStrikeCost {
		panic("Not enough rage for HS")
	}
	warrior.heroicStrikeQueued = true
}

// Returns true if the regular melee swing should be used, false otherwise.
func (warrior *Warrior) TryHeroicStrike(sim *core.Simulation) *core.SimpleSpell {
	if !warrior.heroicStrikeQueued {
		return nil
	}

	warrior.heroicStrikeQueued = false
	if warrior.CurrentRage() < warrior.heroicStrikeCost {
		return nil
	}

	target := sim.GetPrimaryTarget()
	hs := &warrior.heroicStrike
	warrior.heroicStrikeTemplate.Apply(hs)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	hs.Effect.Target = target

	return hs
}

func (warrior *Warrior) CanHeroicStrike(sim *core.Simulation) bool {
	return warrior.CurrentRage() >= warrior.heroicStrikeCost
}
