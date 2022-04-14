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
	baseCost := 495.0

	druid.Moonfire = druid.RegisterSpell(core.SpellConfig{
		ActionID:    MoonfireActionID,
		SpellSchool: core.SpellSchoolArcane,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				Cost: baseCost * (1 - 0.03*float64(druid.Talents.Moonglow)),
				GCD:  core.GCDDefault,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDirectDamage(core.SpellEffect{
			BonusSpellCritRating: float64(druid.Talents.ImprovedMoonfire) * 5 * core.SpellCritRatingPerCritChance,
			DamageMultiplier:     1 * (1 + 0.05*float64(druid.Talents.ImprovedMoonfire)) * (1 + 0.02*float64(druid.Talents.Moonfury)),
			ThreatMultiplier:     1,
			BaseDamage:           core.BaseDamageConfigMagic(305, 357, 0.15),
			OutcomeApplier:       core.OutcomeFuncMagicHitAndCrit(druid.SpellCritMultiplier(1, 0.2*float64(druid.Talents.Vengeance))),
			OnSpellHit: func(sim *core.Simulation, spell *core.Spell, spellEffect *core.SpellEffect) {
				if spellEffect.Landed() {
					druid.MoonfireDot.Apply(sim)
				}
			},
		}),
	})

	target := sim.GetPrimaryTarget()
	druid.MoonfireDot = core.NewDot(core.Dot{
		Spell: druid.Moonfire,
		Aura: target.RegisterAura(core.Aura{
			Label:    "Moonfire-" + strconv.Itoa(int(druid.Index)),
			ActionID: MoonfireActionID,
		}),
		NumberOfTicks: 4 + core.TernaryInt(ItemSetThunderheart.CharacterHasSetBonus(&druid.Character, 2), 1, 0),
		TickLength:    time.Second * 3,
		TickEffects: core.TickFuncSnapshot(target, core.SpellEffect{
			DamageMultiplier: 1 * (1 + 0.05*float64(druid.Talents.ImprovedMoonfire)) * (1 + 0.02*float64(druid.Talents.Moonfury)),
			ThreatMultiplier: 1,
			BaseDamage:       core.BaseDamageConfigMagicNoRoll(600/4, 0.13),
			OutcomeApplier:   core.OutcomeFuncTick(),
			IsPeriodic:       true,
		}),
	})
}

func (druid *Druid) ShouldCastMoonfire(sim *core.Simulation, target *core.Target, rotation proto.BalanceDruid_Rotation) bool {
	return rotation.Moonfire && !druid.MoonfireDot.IsActive()
}
