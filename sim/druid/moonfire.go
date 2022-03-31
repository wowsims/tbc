package druid

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDMoonfire int32 = 26988

var MoonfireDebuffID = core.NewDebuffID()

func (druid *Druid) registerMoonfireSpell(sim *core.Simulation) {
	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    core.ActionID{SpellID: SpellIDMoonfire},
				Character:   &druid.Character,
				SpellSchool: core.SpellSchoolArcane,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 495,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 495,
				},
				GCD: core.GCDDefault,
			},
			OutcomeRollCategory: core.OutcomeRollCategoryMagic,
			CritRollCategory:    core.CritRollCategoryMagical,
			CritMultiplier:      druid.SpellCritMultiplier(1, 0.2*float64(druid.Talents.Vengeance)),
		},
		Effect: core.SpellEffect{
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BaseDamage:       core.BaseDamageConfigMagic(305, 357, 0.15),
			DotInput: core.DotDamageInput{
				NumberOfTicks:  4,
				TickLength:     time.Second * 3,
				TickBaseDamage: core.DotSnapshotFuncMagic(600/4, 0.13),
				DebuffID:       MoonfireDebuffID,
			},
		},
	}

	template.Cost.Value -= template.BaseCost.Value * 0.03 * float64(druid.Talents.Moonglow)

	template.Effect.DamageMultiplier *= 1 + 0.05*float64(druid.Talents.ImprovedMoonfire)
	template.Effect.DamageMultiplier *= 1 + 0.02*float64(druid.Talents.Moonfury)
	template.Effect.BonusSpellCritRating += float64(druid.Talents.ImprovedMoonfire) * 5 * core.SpellCritRatingPerCritChance
	if ItemSetThunderheart.CharacterHasSetBonus(&druid.Character, 2) { // Thunderheart 2p adds 1 extra tick to moonfire
		template.Effect.DotInput.NumberOfTicks += 1
	}

	// moonfire can proc the on hit but doesn't consume charges (i think)
	template.Effect.OnSpellHit = druid.applyOnHitTalents

	druid.Moonfire = druid.RegisterSpell(core.SpellConfig{
		Template:   template,
		ModifyCast: core.ModifyCastAssignTarget,
	})
}

func (druid *Druid) ShouldCastMoonfire(sim *core.Simulation, target *core.Target, rotation proto.BalanceDruid_Rotation) bool {
	return rotation.Moonfire && !druid.Moonfire.Instance.Effect.DotInput.IsTicking(sim)
}
