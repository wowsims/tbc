package hunter

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

var RaptorStrikeCooldownID = core.NewCooldownID()
var RaptorStrikeActionID = core.ActionID{SpellID: 27014, CooldownID: RaptorStrikeCooldownID}

func (hunter *Hunter) registerRaptorStrikeSpell(sim *core.Simulation) {
	cost := core.ResourceCost{Type: stats.Mana, Value: 120}
	ama := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    RaptorStrikeActionID,
				Character:   &hunter.Character,
				SpellSchool: core.SpellSchoolPhysical,
				Cost:        cost,
				BaseCost:    cost,
				Cooldown:    time.Second * 6,
			},
		},
		Effect: core.SpellEffect{
			OutcomeRollCategory: core.OutcomeRollCategorySpecial,
			CritRollCategory:    core.CritRollCategoryPhysical,
			CritMultiplier:      hunter.critMultiplier(false, sim.GetPrimaryTarget()),
			ProcMask:            core.ProcMaskMeleeMHAuto | core.ProcMaskMeleeMHSpecial,
			DamageMultiplier:    1,
			ThreatMultiplier:    1,
			BaseDamage:          core.BaseDamageConfigMeleeWeapon(core.MainHand, false, 170, 1, true),
		},
	}

	ama.Cost.Value -= 120 * 0.2 * float64(hunter.Talents.Resourcefulness)
	ama.Effect.BonusCritRating += float64(hunter.Talents.SavageStrikes) * 10 * core.MeleeCritRatingPerCritChance

	hunter.raptorStrikeCost = ama.Cost.Value

	hunter.RaptorStrike = hunter.RegisterSpell(core.SpellConfig{
		Template:   ama,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}

// Returns true if the regular melee swing should be used, false otherwise.
func (hunter *Hunter) TryRaptorStrike(sim *core.Simulation) *core.Spell {
	if hunter.Rotation.Weave == proto.Hunter_Rotation_WeaveAutosOnly || hunter.IsOnCD(RaptorStrikeCooldownID, sim.CurrentTime) || hunter.CurrentMana() < hunter.raptorStrikeCost {
		return nil
	}

	return hunter.RaptorStrike
}
