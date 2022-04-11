package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDBlizzard int32 = 27085

var BlizzardActionID = core.ActionID{SpellID: SpellIDBlizzard}

func (mage *Mage) registerBlizzardSpell(sim *core.Simulation) {
	//AOECap: 3620,
	baseCost := 1645.0

	blizzardDot := core.NewDot(core.Dot{
		Aura: mage.RegisterAura(&core.Aura{
			Label:    "Blizzard",
			ActionID: BlizzardActionID,
		}),
		NumberOfTicks:       8,
		TickLength:          time.Second * 1,
		AffectedByCastSpeed: true,
		TickEffects: core.TickFuncAOESnapshot(sim, core.SpellEffect{
			DamageMultiplier: mage.spellDamageMultiplier *
				(1 + 0.02*float64(mage.Talents.PiercingIce)) *
				(1 + 0.01*float64(mage.Talents.ArcticWinds)),

			ThreatMultiplier: 1 - (0.1/3)*float64(mage.Talents.FrostChanneling),

			BaseDamage:     core.BaseDamageConfigMagicNoRoll(184, 0.119),
			OutcomeApplier: core.OutcomeFuncTick(),
			IsPeriodic:     true,
		}),
	})

	mage.Blizzard = mage.RegisterSpell(core.SpellConfig{
		ActionID:    BlizzardActionID,
		SpellSchool: core.SpellSchoolFrost,
		SpellExtras: SpellFlagMage | core.SpellExtrasChanneled,

		ResourceType: stats.Mana,
		BaseCost:     baseCost,

		Cast: core.CastConfig{
			DefaultCast: core.NewCast{
				Cost: baseCost *
					(1 - 0.05*float64(mage.Talents.FrostChanneling)) *
					(1 - 0.01*float64(mage.Talents.ElementalPrecision)),

				GCD:         core.GCDDefault,
				ChannelTime: time.Second * 8,
			},
		},

		ApplyEffects: core.ApplyEffectFuncDot(blizzardDot),
	})
	blizzardDot.Spell = mage.Blizzard
}
