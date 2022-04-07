package druid

import (
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDMoonfire int32 = 26988

var MoonfireActionID = core.ActionID{SpellID: SpellIDMoonfire}

func (druid *Druid) registerMoonfireSpell(sim *core.Simulation) {
	template := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    MoonfireActionID,
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
		},
	}
	template.Cost.Value -= template.BaseCost.Value * 0.03 * float64(druid.Talents.Moonglow)

	baseEffect := core.SpellEffect{
		BonusSpellCritRating: float64(druid.Talents.ImprovedMoonfire) * 5 * core.SpellCritRatingPerCritChance,
		DamageMultiplier:     1 * (1 + 0.05*float64(druid.Talents.ImprovedMoonfire)) * (1 + 0.02*float64(druid.Talents.Moonfury)),
		ThreatMultiplier:     1,
		BaseDamage:           core.BaseDamageConfigMagic(305, 357, 0.15),
		OutcomeApplier:       core.OutcomeFuncMagicHitAndCrit(druid.SpellCritMultiplier(1, 0.2*float64(druid.Talents.Vengeance))),
		OnSpellHit:           druid.applyOnHitTalents,
	}

	druid.Moonfire = druid.RegisterSpell(core.SpellConfig{
		Template:   template,
		ModifyCast: core.ModifyCastAssignTarget,
	})

	target := sim.GetPrimaryTarget()
	debuffAura := target.RegisterAura(&core.Aura{
		Label:    "Moonfire-" + strconv.Itoa(int(druid.Index)),
		ActionID: MoonfireActionID,
	})

	dotEffect := baseEffect
	dotEffect.IsPeriodic = true
	dotEffect.BaseDamage = core.BaseDamageConfigMagicNoRoll(600/4, 0.13)
	dotEffect.OutcomeApplier = core.OutcomeFuncTick()
	dotEffect.OnSpellHit = nil

	numTicks := 4
	if ItemSetThunderheart.CharacterHasSetBonus(&druid.Character, 2) { // Thunderheart 2p adds 1 extra tick to moonfire
		baseEffect.DotInput.NumberOfTicks += 1
	}

	druid.MoonfireDot = core.NewDot(core.Dot{
		Aura:          debuffAura,
		NumberOfTicks: numTicks,
		TickLength:    time.Second * 3,
		TickEffects:   core.TickFuncSnapshot(dotEffect, target),
	})

	druid.Moonfire = druid.RegisterSpell(core.SpellConfig{
		Template:     template,
		ModifyCast:   core.ModifyCastAssignTarget,
		ApplyEffects: core.ApplyEffectFuncDot(baseEffect, druid.MoonfireDot),
	})
	druid.MoonfireDot.Spell = druid.Moonfire
}

func (druid *Druid) ShouldCastMoonfire(sim *core.Simulation, target *core.Target, rotation proto.BalanceDruid_Rotation) bool {
	return rotation.Moonfire && !druid.MoonfireDot.IsActive()
}
