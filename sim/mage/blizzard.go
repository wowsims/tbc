package mage

import (
	"time"

	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/stats"
)

const SpellIDBlizzard int32 = 27085

var BlizzardActionID = core.ActionID{SpellID: SpellIDBlizzard}

func (mage *Mage) registerBlizzardSpell(sim *core.Simulation) {
	spell := core.SimpleSpell{
		SpellCast: core.SpellCast{
			Cast: core.Cast{
				ActionID:    BlizzardActionID,
				SpellSchool: core.SpellSchoolFrost,
				SpellExtras: SpellFlagMage | core.SpellExtrasChanneled,
				Character:   &mage.Character,
				BaseCost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 1645,
				},
				Cost: core.ResourceCost{
					Type:  stats.Mana,
					Value: 1645,
				},
				GCD:         core.GCDDefault,
				ChannelTime: time.Second * 8,
			},
		},
		AOECap: 3620,
	}
	spell.Cost.Value -= spell.BaseCost.Value * float64(mage.Talents.FrostChanneling) * 0.05
	spell.Cost.Value *= 1 - float64(mage.Talents.ElementalPrecision)*0.01

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
		Template:     spell,
		ApplyEffects: core.ApplyEffectFuncDot(blizzardDot),
	})
	blizzardDot.Spell = mage.Blizzard
}
