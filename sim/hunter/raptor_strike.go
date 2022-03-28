package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var RaptorStrikeCooldownID = core.NewCooldownID()
var RaptorStrikeActionID = core.ActionID{SpellID: 27014, CooldownID: RaptorStrikeCooldownID}

func (hunter *Hunter) newRaptorStrikeTemplate(sim *core.Simulation) core.SimpleSpellTemplate {
	cost := core.ResourceCost{Type: stats.Mana, Value: 120}
	ama := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:            RaptorStrikeActionID,
				Character:           &hunter.Character,
				OutcomeRollCategory: core.OutcomeRollCategorySpecial,
				CritRollCategory:    core.CritRollCategoryPhysical,
				SpellSchool:         core.SpellSchoolPhysical,
				Cost:                cost,
				BaseCost:            cost,
				Cooldown:            time.Second * 6,
				CritMultiplier:      hunter.critMultiplier(false, sim.GetPrimaryTarget()),
			},
		},
		Effect: core.SpellHitEffect{
			SpellEffect: core.SpellEffect{
				ProcMask:               core.ProcMaskMeleeMHAuto | core.ProcMaskMeleeMHSpecial,
				DamageMultiplier:       1,
				StaticDamageMultiplier: 1,
				ThreatMultiplier:       1,
			},
			BaseDamage: core.BaseDamageFuncMeleeWeapon(core.MainHand, false, 170, 1, true),
		},
	}

	ama.Cost.Value -= 120 * 0.2 * float64(hunter.Talents.Resourcefulness)
	ama.Effect.BonusCritRating += float64(hunter.Talents.SavageStrikes) * 10 * core.MeleeCritRatingPerCritChance

	hunter.raptorStrikeCost = ama.Cost.Value

	return core.NewSimpleSpellTemplate(ama)
}

func (hunter *Hunter) NewRaptorStrike(sim *core.Simulation, target *core.Target) *core.SimpleSpell {
	rs := &hunter.raptorStrike
	hunter.raptorStrikeTemplate.Apply(rs)

	// Set dynamic fields, i.e. the stuff we couldn't precompute.
	rs.Effect.Target = target

	rs.Init(sim)
	return rs
}

// Returns true if the regular melee swing should be used, false otherwise.
func (hunter *Hunter) TryRaptorStrike(sim *core.Simulation) *core.SimpleSpell {
	if hunter.Rotation.Weave == proto.Hunter_Rotation_WeaveAutosOnly || hunter.IsOnCD(RaptorStrikeCooldownID, sim.CurrentTime) || hunter.CurrentMana() < hunter.raptorStrikeCost {
		return nil
	}

	return hunter.NewRaptorStrike(sim, sim.GetPrimaryTarget())
}
